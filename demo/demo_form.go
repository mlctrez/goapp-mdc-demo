package demo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/helperline"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/textarea"
	"github.com/mlctrez/goapp-mdc/pkg/textfield"
)

type FormDemo struct {
	app.Compo
}

func id() string {
	return uuid.New().String()
}

func textAreaExample() []app.UI {
	idOne := id()
	taOne := textarea.New(idOne).Size(8, 40).Outlined(true).
		Label("outlined text area").MaxLength(240)
	helpOne := helperline.New(idOne, "textarea help text", "0 / 240")

	return []app.UI{app.Div().Style("display", "inline-block").Body(taOne, helpOne)}
}

func (e *FormDemo) Render() app.UI {

	value := "a"

	in := input.NewInput("myInput", "myInput", input.InputTypeText, &value, func(value app.Value) error {
		if len(value.String()) < 4 {
			return errors.New("an error")
		}
		return nil
	}, input.InputOptMin(2))

	body := layout.Grid().Body(layout.Inner().Style("display", "flex").Body(
		layout.Cell().Body(layout.Inner().Style("display", "flex").Body(
			layout.CellWide().Body(app.H4().Text("Text Area")),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "normal"}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "required", Required: true}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "outlined", Outlined: true}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Label: "outlined required",
				Outlined: true, Required: true}),
			layout.Cell().Body(&textfield.TextField{Id: id(), Placeholder: "placeholder"}),
		)),
		layout.Cell().Body(layout.Inner().Style("display", "flex").Body(
			layout.CellWide().Body(app.H4().Text("Text Field")),
			layout.Cell().Body(textAreaExample()...),
		)),
		layout.Cell().Body(layout.Inner().Style("display", "flex").Body(
			in,
		)),

		//layout.Cell().Body(layout.Inner().Style("display", "flex").Body(
		//	app.Form().OnSubmit(func(ctx app.Context, e app.Event) {
		//		e.PreventDefault()
		//		log.Println("submit")
		//	}, nil).Body(app.Button().Type("submit").Text("submit form")),
		//)),
	))

	_ = body
	return PageBody(body)
}
