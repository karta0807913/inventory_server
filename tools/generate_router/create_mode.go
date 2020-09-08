package main

// create mode default every fields (without primaryKey) are required

import (
	"fmt"
	"strings"
)

type MethodCreateParams struct {
	ParsedType Type
	Doc        *Document
	RequireSet *CommaSet
	OptionsSet *CommaSet
	IgnoreSet  *CommaSet
	TagKey     string
}

// Create *TemplateRoot and update *Document
func MethodCreate(arg MethodCreateParams) *TemplateRoot {
	templateRoot := TemplateRoot{
		FuncName:       *method,
		StructName:     arg.ParsedType.Name,
		Decoder:        "ShouldBind" + strings.ToUpper(arg.TagKey),
		RequiredFields: make([]TemplateField, 0),
		OptionalFields: make([]TemplateField, 0),
		Mode:           "Create",
	}

	for _, field := range arg.ParsedType.Fields {
		tf := TemplateField{
			Name: field.Name,
			Type: field.Type,
		}
		tags := make([]string, 0)
		decoder, ok := field.Tag.Lookup(arg.TagKey)
		if ok {
			tags = append(tags, fmt.Sprintf(`%s:"%s"`, arg.TagKey, decoder))
		}
		gormTag, ok := field.Tag.Lookup("gorm")
		//     16       8      4      2        1
		// primaryKey unique index not_null default
		//     0        0      0      0        0
		var join uint8 = 0
		if ok {
			opt := gormTag
			if strings.Index(opt, "not null") != -1 {
				join |= 2
			}
			if strings.Index(opt, "primaryKey") != -1 {
				join |= 16
			}
			if strings.Index(opt, "index") != -1 {
				join |= 4
			}
			if strings.Index(opt, "unique") != -1 {
				join |= 8
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
			if join >= 16 {
				continue
			}
			tf.Tag = "`" + strings.Join(
				append(tags, `binding:"required"`), " ") + "`"
			templateRoot.RequiredFields = append(templateRoot.RequiredFields, tf)
		}
	}
	return &templateRoot
}
