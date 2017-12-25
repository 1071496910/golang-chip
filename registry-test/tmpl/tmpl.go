package tmpl

import "fmt"

type Tmpl interface {
	Parser(s string) string
}

var TmplRegistry map[string]func() Tmpl = make(map[string]func() Tmpl)

func Parser(tmplName string, content string) string {
	if factoryFunc, ok := TmplRegistry[tmplName]; ok {
		return factoryFunc().Parser(content)
	}
	return defaultParser(content)
}

func defaultParser(s string) string {
	return fmt.Sprintf("default: %v", s)
}
