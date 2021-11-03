package main

import (
	"flag"
	"fmt"
	"strings"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func stderr(format string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, format, params...)
}

func main() {
	var (
		flags        flag.FlagSet
		plugins      = flags.String("plugins", "", "list of plugins to enable (supported values: grpc)")
		importPrefix = flags.String("import_prefix", "", "prefix to prepend to import paths")
	)
	importRewriteFunc := func(importPath protogen.GoImportPath) protogen.GoImportPath {
		switch importPath {
		case "context", "fmt", "math":
			return importPath
		}
		if *importPrefix != "" {
			return protogen.GoImportPath(*importPrefix) + importPath
		}
		return importPath
	}
	protogen.Options{
		ParamFunc:         flags.Set,
		ImportRewriteFunc: importRewriteFunc,
	}.Run(func(gen *protogen.Plugin) error {
		for _, plugin := range strings.Split(*plugins, ",") {
			switch plugin {
			case "":
				// TODO(dape): what does this do? looks strange? bug in lib?
			default:
				return fmt.Errorf("protoc-gen-go: unknown plugin %q", plugin)
			}
		}
		for _, f := range gen.Files {
			stderr("%+v", f)
			
			if !f.Generate {
				continue
			}
			
			filename := f.GeneratedFilenamePrefix + ".bson.go"
			gf := gen.NewGeneratedFile(filename, "")
			gf.P("package", " ", f.GoPackageName)
			
			gf.P()

			gf.Import("go.mongodb.org/mongo-driver/bson")
			
			for _, msg := range f.Messages {
				emitToBsonFuncs(gf, msg)
			}
		}
		// TODO(dape): what does this do? needed? or ok to skip?
		//gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}

func createFieldAccessor(field *protogen.Field) string {	
	if field.Message == nil {
		// Primitive field; such as int and string, use the GetXyz() method
		// since it will return the correct default value.
		return "msg.Get" + field.GoName + "()" + "//a"
	}
	
	prefix := field.Message.GoIdent.GoName
	if field.Desc.IsList() {
		// TODO(dape): create a bson.A{} and then populate it.
		panic("should never get here; list")
		return prefix + "ArrayToBSON(msg." + field.GoName + ")"
	}

	if field.Desc.IsMap() {
		// TODO(dape): create a bson.M{} and then populate it.
		panic("should never get here; map")
		return prefix + "ToBSONMap(msg." + field.GoName + ")"
	}
	
	return field.Message.GoIdent.GoName + fmt.Sprintf("ToBSON(msg.%s)", field.GoName)
}

// TODO(dape): change into a more generic structure that is more generally useful.
type bsonPackage struct {
	dest *protogen.GeneratedFile
}

func (b bsonPackage) M() string {
	return b.dest.QualifiedGoIdent(protogen.GoIdent{"M", "go.mongodb.org/mongo-driver/bson"})
}

func (b bsonPackage) A() string {
	return b.dest.QualifiedGoIdent(protogen.GoIdent{"A", "go.mongodb.org/mongo-driver/bson"})
}

// TODO(dape): change to emit per type functions to avoid the problem of unique naming using prefixes of the type names.

func emitListConversionFunc(gf *protogen.GeneratedFile, field *protogen.Field, msg *protogen.Message) string {
	name := fmt.Sprintf("%s_%s_ArrayToBSON", msg.GoIdent.GoName, field.GoName)
	ftype := field.Desc.Kind().String()
	if field.Message != nil {
		ftype = "*" + field.Message.GoIdent.GoName
	}
	
	gf.P("func ", name, "(values []", ftype, ") bson.A {")
	gf.P("res := bson.A{}")
	gf.P("for _, v := range values {")
	// TODO(dape): refactor the concat of Message and ToBSON to a function to avoid duplication.
	if field.Message != nil {
		gf.P("res = append(res, ", field.Message.GoIdent.GoName, "ToBSON(v))")
	} else {
		gf.P("res = append(res, v)")
	}
	gf.P("}")
	gf.P("return res")
	gf.P("}")
	gf.P()

	return name
}

func emitMapConversionFunc(gf *protogen.GeneratedFile, field *protogen.Field, msg *protogen.Message) string {
	name := fmt.Sprintf("%s_%s_MapToBSON", msg.GoIdent.GoName, field.GoName)
	stderr("\n\n%s:%s", field.Desc.MapKey().Kind(), field.Desc.MapValue().Kind())

	ktype := field.Desc.MapKey().Kind().String()
	vkind := field.Desc.MapValue().Kind()
	vtype := vkind.String()
	if vkind == protoreflect.MessageKind {
		panic("unsupported feature")
	}

	gf.P("func ", name, "(values map[", ktype, "]", vtype, ") bson.M {")
	gf.P("res := bson.M{}")
	gf.P("for k, v := range values {")
	gf.P("res[k] = v")
	gf.P("}")
	gf.P("return res")
	gf.P("}")
	gf.P()

	return name
}

func emitToBsonFuncs(gf *protogen.GeneratedFile, msg *protogen.Message) {	
	// Since we emit the conversion methods and types into the same
	// package as the normal package uses we know that the names of
	// the messages are unique.
	funcName := msg.GoIdent.GoName + "ToBSON"
	
	bson := bsonPackage{gf}

	conversionFuncs := map[string]string{}
	for _, field := range msg.Fields {
		if field.Desc.IsList() {
			conversionFuncs[field.GoName] = emitListConversionFunc(gf, field, msg)
		} else if field.Desc.IsMap() {
			conversionFuncs[field.GoName] = emitMapConversionFunc(gf, field, msg)
		}
	}

	// the Message type func
	// TODO(dape): refactor the concat of Message and ToBSON to a function to avoid duplication.
	gf.P("func ", funcName, "(msg *", msg.GoIdent.GoName, ") ", bson.M(), " {")
	
	gf.P("res := ", bson.M(), "{}")
	gf.P("if msg == nil {")
	gf.P("msg = &" + msg.GoIdent.GoName + "{}")
	gf.P("}")
	
	for _, field := range msg.Fields {
		accessor, ok := conversionFuncs[field.GoName]
		if ok {
			accessor = accessor + "(msg.Get" + field.GoName + "())"
		} else {
			accessor = createFieldAccessor(field)
		}

		gf.P("res[\"", field.GoName, "\"]=", accessor)
	}
	gf.P("return res")
	gf.P("}")
}
