package pages

import "github.com/rohanthewiz/element"

// Home Page component
type Home struct {
	Heading string
}

func (h Home) Render() (out string) {
	b := element.NewBuilder()

	b.Body("style", "background-color:tan").R(
		b.H1("style", "color:maroon;background-color:#dfc673").T(h.Heading),

		element.RenderComponents(b, ContactForm{}, Footer{}),
		// ContactForm{}.Render(b), Footer{}.Render(b),
	)

	// ...

	return b.String()
}
