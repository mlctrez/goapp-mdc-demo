package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/formfield"
	"github.com/mlctrez/goapp-mdc/pkg/radio"
)

type RadioDemo struct {
	app.Compo
}

func (d *RadioDemo) Render() app.UI {

	radioOne := &formfield.FormField{Label: "radio one",
		Component: &radio.Radio{Id: "radioOne", Name: "radios", Value: "radioOne"}}
	radioTwo := &formfield.FormField{Label: "radio two",
		Component: &radio.Radio{Id: "radioTwo", Name: "radios", Value: "radioTwo"}}

	return PageBody(radioOne, radioTwo)

}
