package demo

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
	"github.com/mlctrez/goapp-mdc/pkg/tab"
)

type TabDemo struct {
	app.Compo
	base.JsUtil
	bar     *tab.Bar
	message *base.Message
	button  *button.Button
}

func (d *TabDemo) Render() app.UI {

	if d.bar == nil {
		bar := &tab.Bar{}
		bar.Tabs = tab.Tabs{}
		bar.Tabs = append(bar.Tabs, &tab.Tab{Label: "no icon"})
		bar.Tabs = append(bar.Tabs, &tab.Tab{Icon: icon.MIApi, Label: "apis"})
		bar.Tabs = append(bar.Tabs, &tab.Tab{Icon: icon.MIFavorite, Label: "favorite"})
		d.bar = bar
		d.bar.Tabs.Select(0)

		d.message = &base.Message{Text: "tab messages appear here"}
		d.button = &button.Button{Label: "index 1", Callback: func(button app.HTMLButton) {
			button.OnClick(func(ctx app.Context, e app.Event) {
				ctx.NewActionWithValue(string(tab.Activate), d.bar, app.T("index", "1"))
			})
		}}

	}

	var cells []app.UI
	cells = append(cells, layout.Cell().Body(d.bar))
	cells = append(cells, layout.Cell().Body(d.message))
	cells = append(cells, layout.Cell().Body(d.button))
	body := layout.Grid().Body(
		layout.Inner().Style("display", "flex").Body(cells...),
	)

	return PageBody(body)
}

func (d *TabDemo) OnMount(ctx app.Context) {
	ctx.Handle(string(tab.Activated), func(context app.Context, action app.Action) {
		if action.Value == d.bar {
			d.message.Text = fmt.Sprintf("action %v", action)
			d.message.Update()
		}
	})

}
