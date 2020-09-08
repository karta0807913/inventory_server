package main

import (
	"fmt"
	"strings"
)

type MethodSearchParams struct {
	ParsedType Type
	Doc        *Document
	RequireSet *CommaSet
	OptionsSet *CommaSet
	IgnoreSet  *CommaSet
	TagKey     string
}

// default the "index" will be join as options
func MethodSearch(arg MethodSearchParams) *TemplateRoot {
	templateRoot := TemplateRoot{
		FuncName:       *method,
		StructName:     arg.ParsedType.Name,
		Decoder:        "ShouldBindQuery",
		RequiredFields: make([]TemplateField, 0),
		OptionalFields: make([]TemplateField, 0),
	}

	for _, field := range arg.ParsedType.Fields {
		tf := TemplateField{
			Name: field.Name,
			Type: field.Type,
		}
		tags := make([]string, 0)
		decoder, ok := field.Tag.Lookup(arg.TagKey)
		if ok {
			t := strings.Split(decoder, ":")
			if len(t) > 1 {
				tags = append(tags, fmt.Sprintf(`form:"%s"`, arg.TagKey, decoder))
			}
		}
		gormTag, ok := field.Tag.Lookup("gorm")
		//     16       8      4      2        1
		// primaryKey unique index not_null default
		//     0        0      0      0        0
		var flag uint8 = 0
		if ok {
			opt := gormTag
			if strings.Index(opt, "not null") != -1 {
				flag |= 2
			}
			if strings.Index(opt, "primaryKey") != -1 {
				flag |= 16
			}
			if strings.Index(opt, "index") != -1 {
				flag |= 4
			}
			if strings.Index(opt, "unique") != -1 {
				flag |= 8
			}
		}

		// if this field is required
		if arg.IgnoreSet.CheckAndDelete(field.Name) {
			continue
		} else if arg.RequireSet.CheckAndDelete(field.Name) {
			tf.Tag = "`" + strings.Join(
				append(tags, `binding:"required"`), " ") + "`"
			templateRoot.RequiredFields = append(templateRoot.RequiredFields, tf)
		} else if arg.OptionsSet.CheckAndDelete(field.Name) {
			tf.Tag = "`" + strings.Join(tags, " ") + "`"
			templateRoot.OptionalFields = append(templateRoot.OptionalFields, tf)
		} else {
			if flag < 4 {
				continue
			}
			tf.Tag = "`" + strings.Join(tags, " ") + "`"
			templateRoot.OptionalFields = append(templateRoot.OptionalFields, tf)
		}
	}
	return &templateRoot
}
