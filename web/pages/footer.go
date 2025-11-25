package pages

import "github.com/rohanthewiz/element"

type Footer struct{}

func (f Footer) Render(b *element.Builder) (dontCare any) {
	b.Div("style", "background-color:lightgray").R(
		b.P("style", "color:gray").T("Copyright &copy; 2025"),
	)
	return
}
