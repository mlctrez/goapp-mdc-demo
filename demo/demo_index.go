package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type Index struct {
	app.Compo
}

func (i *Index) Render() app.UI {

	content := layout.Grid().Body(
		layout.Inner().Body(
			layout.Cell().Body(
				app.H3().Text("goapp mdc"),

				app.Text("The goal of this is to demonstrate the goapp-mdc components in a similar layout as "),
				app.A().Href("https://material-components.github.io/material-components-web-catalog/#/").Text("Material Components Web Catalog"),
				app.Text("."),
			)))

	return PageBody(content)

}
