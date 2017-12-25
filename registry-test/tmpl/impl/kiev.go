package tmpl

import "registry-test/tmpl"
import "fmt"

type kievTmpl struct {
}

func (j *kievTmpl) Parser(s string) string {
	return fmt.Sprintf("kiev-template: %v", s)
}

func init() {
	tmpl.TmplRegistry["kiev"] = func() tmpl.Tmpl {
		return &kievTmpl{}
	}
}
