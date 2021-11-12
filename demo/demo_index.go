package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Index struct {
	app.Compo
}

func (i *Index) Render() app.UI {

	homeBody := app.Div()

	paragraph := func() app.HTMLP {
		t := "lorem ipsum lorem ipsum lorem ipsum "
		for i := 0; i < 10; i++ {
			t += t
		}
		return app.P().Text(t)
	}

	homeBody.Body(

		app.H3().Text("go-app mdc"),
		paragraph(),
		app.Hr(),
		paragraph(),
		app.Hr(),
		paragraph(),
		app.Hr(),
		paragraph(),
	)

	return PageBody(homeBody)

}
