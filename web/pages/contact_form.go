package pages

import "github.com/rohanthewiz/element"

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
