package tmpl

import "registry-test/tmpl"
import "fmt"

type jettyTmpl struct {
}

func (j *jettyTmpl) Parser(s string) string {
	return fmt.Sprintf("jetty-template: %v", s)
}

func init() {
	tmpl.TmplRegistry["jetty"] = func() tmpl.Tmpl {
		return &jettyTmpl{}
	}
}
