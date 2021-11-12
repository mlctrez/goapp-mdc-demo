package demo

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/switchm"
)

type SwitchDemo struct {
	app.Compo
}

func (s *SwitchDemo) Render() app.UI {

	var rows []app.UI

	for _, selected := range []bool{false, true} {
		for _, disabled := range []bool{false, true} {
			label := fmt.Sprintf("Selected:%t Disabled:%t", selected, disabled)
			sel := &switchm.MDCSwitch{Selected: selected,Disabled: disabled}
			rows = append(rows, GridRow(label, sel))
		}
	}

	return PageBody(layout.Grid().Body(rows...))
}
