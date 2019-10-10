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
	flags := c.generateSubMessageRequestFlags("reqArgs", "", d, file, types)
	initialize := c.generateRequestInitialization(d, file, types)
	return initialize, flags
}

func (c *client) generateSubMessageRequestFlags(objectName, flagPrefix string, d *pb.DescriptorProto, file *generator.FileDescriptor, types protoTypeCache) []string {
	out := make([]string, 0, len(d.Field))

	for _, f := range d.Field {
		fieldName := goFieldName(f)
		fieldFlagName := strings.ToLower(fieldName)
		if f.GetLabel() == pb.FieldDescriptorProto_LABEL_REPEATED {
			// TODO
			out = append(out, fmt.Sprintf(`.PersistentFlags() // Warning: list flags are not yet supported (field %q)`, fieldName))
			continue
		}

		switch f.GetType() {
		// Field is a complex type (another message, or an enum)
		case pb.FieldDescriptorProto_TYPE_MESSAGE:
			// if both type and name are set, descriptor must be either a message or enum
			_, _, ttype := inputNames(f.GetTypeName())
			if fdesc, found, _ := types.byName(file.MessageType, ttype, noop /*prefix("// ", c.P)*/); found {
				if fdesc.GetOptions().GetMapEntry() {
					// TODO
					return []string{fmt.Sprintf(`.PersistentFlags() // Warning: map flags are not yet supported (message %q)`, d.GetName())}
				}

				flags := c.generateSubMessageRequestFlags(objectName+"."+fieldName, flagPrefix+fieldFlagName+"-", fdesc, file, types)
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
	return out
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
	initialize := genReqInit(d, file, types, "", false, debug, noop /*prefix("// ", c.P)*/)
	// c.P(debug.String())
	return initialize
}

func genReqInit(d *pb.DescriptorProto, file *generator.FileDescriptor, types protoTypeCache, typePrefix string, repeated bool, w io.Writer, log func(...interface{})) string {
	if repeated {
		// if we're repeated, we only want to compute the type then bail, we won't figure out if we're trying to create an instance
		out := fmt.Sprintf("[]*%s%s{}", typePrefix, d.GetName())
		fmt.Fprintf(w, "// computed %q\n", out)
		return out
	}

	fields := make(map[string]string)
	fmt.Fprintf(w, "// generating initialization for %s with prefix %q which has %d fields\n", d.GetName(), typePrefix, len(d.Field))
	for _, f := range d.Field {
		switch f.GetType() {
		case pb.FieldDescriptorProto_TYPE_MESSAGE:
			_, _, ttype := inputNames(f.GetTypeName())
			desc, found, nested := types.byName(file.MessageType, ttype, log)
			fmt.Fprintf(w, "// searching for type %q with ttype %q for field %q\n", f.GetTypeName(), ttype, f.GetName())
			if !found {
				fmt.Fprint(w, "// not found, skipping\n")
				continue
			}

			if desc.GetOptions().GetMapEntry() {
				fmt.Fprintf(w, "// skipping map fields, which do not need to be initialized")
				continue
			}

			prefix := typePrefix
			if nested {
				prefix += d.GetName() + "_"
			}

			fmt.Fprintf(w, "// found, recursing with %q\n", desc.GetName())
			m := genReqInit(desc, file, types, prefix, listField(f), w, log)
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
	values := "{}"
	if len(vals) > 0 {
		values = fmt.Sprintf("{\n%s,\n}", strings.Join(vals, ",\n"))
	}

	prefix := fmt.Sprintf("&%s%s", typePrefix, d.GetName())

	out := prefix + values
	fmt.Fprintf(w, "// computed %q\n", out)
	return out
}

func listField(d *pb.FieldDescriptorProto) bool {
	return d.GetLabel() == pb.FieldDescriptorProto_LABEL_REPEATED
}
