package client

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"

	"github.com/tetratelabs/protoc-gen-cobra/generator"
)

type protoTypeCache map[string]entry
type entry struct {
	d    *pb.DescriptorProto
	f, n bool
}

func (p protoTypeCache) byName(desc []*pb.DescriptorProto, name string, log func(...interface{})) (*pb.DescriptorProto, bool, bool) {
	return byName(p, desc, name, false, log)
}

func byName(p protoTypeCache, desc []*pb.DescriptorProto, name string, nested bool, log func(...interface{})) (*pb.DescriptorProto, bool, bool) {
	log("searching for ", name)
	if entry, found := p[name]; found {
		log("* found ", entry.d.GetName(), "in cache ", fmt.Sprintf("%v", p))
		return entry.d, entry.f, entry.n
	}

	for _, d := range desc {
		if d.GetName() == name {
			p[name] = entry{d, true, nested}
			log("* comparing against ", d.GetName(), " inserting into cache: \n// ", fmt.Sprintf("%v", p))
			return d, true, nested
		} else {
			log("  comparing against ", d.GetName())
		}
		if desc, found, _ := byName(p, d.NestedType, name, true, prefix("    ", log)); found {
			return desc, found, true
		}
	}
	return nil, false, false
}

func prefix(pre string, l func(...interface{})) func(...interface{}) {
	return func(i ...interface{}) { l(append([]interface{}{pre}, i...)...) }
}

func noop(...interface{}) {}

// first return is the instantiation of the struct and fields that are messages; second is the set of
// flag declarations using the fields of the struct to receive values
func (c *client) generateRequestFlags(file *generator.FileDescriptor, d *pb.DescriptorProto, types protoTypeCache) (string, []string) {
	if d == nil {
		return "", []string{}
	}
	return c.generateSubMessageRequestFlags("reqArgs", "", d, file, types)
}

func (c *client) generateSubMessageRequestFlags(objectName, flagPrefix string, d *pb.DescriptorProto, file *generator.FileDescriptor, types protoTypeCache) (s string, ss []string) {
	out := make([]string, 0, len(d.Field))
	for _, f := range d.Field {
		fieldName := goFieldName(f)
		fieldFlagName := strings.ToLower(fieldName)

		switch f.GetType() {
		// Field is a complex type (another message, or an enum)
		case pb.FieldDescriptorProto_TYPE_MESSAGE:
			// if both type and name are set, descriptor must be either a message or enum
			_, _, ttype := inputNames(f.GetTypeName())
			if fdesc, found, _ := types.byName(file.MessageType, ttype, noop /*prefix("// ", c.P)*/); found {
				_, flags := c.generateSubMessageRequestFlags(objectName+"."+fieldName, flagPrefix+fieldFlagName+"-", fdesc, file, types)
				out = append(out, flags...)
			}
		case pb.FieldDescriptorProto_TYPE_ENUM:
			// TODO
		case pb.FieldDescriptorProto_TYPE_STRING:
			out = append(out, fmt.Sprintf(`.PersistentFlags().StringVar(&%s.%s, "%s%s", "", "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_BYTES:
			f.GetJsonName()
			out = append(out, fmt.Sprintf(`.PersistentFlags().BytesBase64Var(&%s.%s, "%s%s", []byte{}, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_BOOL:
			out = append(out, fmt.Sprintf(`.PersistentFlags().BoolVar(&%s.%s, "%s%s", false, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_FLOAT:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Float32Var(&%s.%s, "%s%s", 0, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_DOUBLE:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Float64Var(&%s.%s, "%s%s", 0, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_INT64:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int64Var(&%s.%s, "%s%s", 0, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_UINT64:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Uint64Var(&%s.%s, "%s%s", 0, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_INT32:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int32Var(&%s.%s, "%s%s", 0, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_FIXED64:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int64Var(&%s.%s, "%s%s", 0, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_FIXED32:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int32Var(&%s.%s, "%s%s", 0, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_UINT32:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Uint32Var(&%s.%s, "%s%s", false, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_SFIXED32:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int32Var(&%s.%s, "%s%s", false, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_SFIXED64:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int64Var(&%s.%s, "%s%s", false, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_SINT32:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int32Var(&%s.%s, "%s%s", false, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_SINT64:
			out = append(out, fmt.Sprintf(`.PersistentFlags().Int64Var(&%s.%s, "%s%s", false, "%s")`,
				objectName, fieldName, flagPrefix, fieldFlagName, "get-comment-from-proto"))
		case pb.FieldDescriptorProto_TYPE_GROUP:
		default:
		}
	}

	initialize := c.generateRequestInitialization(d, file, types)
	return initialize, out
}

func goFieldName(f *pb.FieldDescriptorProto) string {
	fieldName := f.GetJsonName()
	if fieldName != "" {
		fieldName = strings.ToUpper(string(fieldName[0])) + fieldName[1:]
	}
	return fieldName
}

func (c *client) generateRequestInitialization(d *pb.DescriptorProto, file *generator.FileDescriptor, types protoTypeCache) string {
	debug := &bytes.Buffer{}
	initialize := genReqInit(d, file, types, "", debug, 0, noop /*prefix("// ", c.P)*/)
	// c.P(debug.String())
	return initialize
}

func genReqInit(d *pb.DescriptorProto, file *generator.FileDescriptor, types protoTypeCache, typePrefix string, w io.Writer, count int, log func(...interface{})) string {
	if count > 10 {
		log("recursed too many times")
		return ""
	}
	fields := make(map[string]string)
	fmt.Fprintf(w, "// generating initialization for %s with prefix %q which has %d fields\n", d.GetName(), typePrefix, len(d.Field))
	for _, f := range d.Field {
		switch f.GetType() {
		case pb.FieldDescriptorProto_TYPE_MESSAGE:
			_, _, ttype := inputNames(f.GetTypeName())
			fieldDesc, found, nested := types.byName(file.MessageType, ttype, log)
			fmt.Fprintf(w, "// searching for type %q with ttype %q for field %q\n", f.GetTypeName(), ttype, f.GetName())
			if !found {
				fmt.Fprint(w, "// not found, skipping\n")
				continue
			}

			prefix := typePrefix
			if nested {
				prefix += d.GetName() + "_"
			}

			fmt.Fprintf(w, "// found, recursing with %q\n", fieldDesc.GetName())
			m := genReqInit(fieldDesc, file, types, prefix, w, count+1, log)
			fmt.Fprintf(w, "// found field %q which we'll initialize with %q\n", goFieldName(f), m)
			fields[goFieldName(f)] = m
		default:
			fmt.Fprintf(w, "// found non-message field %q\n", f.GetName())
		}
	}

	vals := make([]string, 0, len(fields))
	for n, v := range fields {
		vals = append(vals, n+": "+v)
	}
	out := fmt.Sprintf("&%s%s{}", typePrefix, d.GetName())
	if len(vals) > 0 {
		out = fmt.Sprintf("&%s%s{\n%s,\n}", typePrefix, d.GetName(), strings.Join(vals, ",\n"))
	}
	fmt.Fprintf(w, "// computed %q\n", out)
	return out
}
