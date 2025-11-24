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

// ContactForm component renders a contact form using element package
// TODO: Put contact form and footer logic in separate files
type ContactForm struct{}

func (c ContactForm) Render(b *element.Builder) (dontCare any) {
	// Build the contact form with proper attributes and structure
	b.Form("action", "/contact", "method", "POST").R(
		b.Input("type", "text", "name", "name", "placeholder", "Name"),
		b.Input("type", "email", "name", "email", "placeholder", "Email"),
		b.TextArea("name", "message", "placeholder", "Message").R(),
		b.Button("type", "submit").T("Send"),
	)

	return
}

type Footer struct{}

func (f Footer) Render(b *element.Builder) (dontCare any) {
	b.Div("style", "background-color:lightgray").R(
		b.P("style", "color:gray").T("Copyright &copy; 2025"),
	)
	return
}
