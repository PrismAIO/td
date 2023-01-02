package main

const stylingTmpl = `// Code generated by mkentity, DO NOT EDIT.
package {{ $.PackageName }}

import (
	"context"

	"github.com/PrismAIO/td/tg"
)

var (
	_ = tg.Invoker(nil)
	_ = context.Context(nil)
)

{{- /*gotype: github.com/PrismAIO/td/telegram/message/internal/mkrun.Config*/ -}}
{{- range $typ := $.Data }}
{{ $helperName := trimPrefix ( trimPrefix $typ.Name "Input" ) "MessageEntity" -}}
// {{ $helperName }} formats text as {{ $helperName }} entity.
//
// See https://core.telegram.org/constructor/{{ $typ.SchemaType.Name }}.
func {{ $helperName }}(s string,
{{- range $f := $typ.Fields }}{{ lowerFirst $f.Name }} {{ $f.Type }}{{- end }}) StyledTextOption {
	return styledTextOption(s, func(b *textBuilder) error {
		b.{{ $helperName }}(s, {{- range $f := $typ.Fields }}{{ lowerFirst $f.Name }},{{- end }})
		return nil
	})
}
{{- end }}
`
