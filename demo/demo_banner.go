package demo

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/banner"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/button"
	"github.com/mlctrez/goapp-mdc/pkg/checkbox"
	"github.com/mlctrez/goapp-mdc/pkg/formfield"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type BannerDemo struct {
	app.Compo
	base.JsUtil
	floating *banner.Banner
	fixed    *banner.Banner
	message  *base.Message
}


func (c *BannerDemo) Render() app.UI {

	if c.floating == nil {
		c.floating = &banner.Banner{
			Id: "normalBanner", Text: "This is the banner text for a normal banner",
			Buttons: []app.UI{
				&button.Button{Id: c.UUID(), Label: "Primary", Banner: true, BannerAction: "primary"},
				&button.Button{Id: c.UUID(), Label: "Secondary", Banner: true, BannerAction: "secondary"},
			},
		}
		c.fixed = &banner.Banner{
			Id: "fixedBanner", Text: "This is the banner text for a fixed banner", Fixed: true,
			Buttons: []app.UI{
				&button.Button{Id: c.UUID(), Label: "Primary", Banner: true, BannerAction: "primary"},
				&button.Button{Id: c.UUID(), Label: "Secondary", Banner: true, BannerAction: "secondary"},
			},
		}
		c.message = &base.Message{Text: "banner events will appear here"}
	}
	openFloating := &button.Button{Id: c.UUID(), Label: "floating", Callback: func(button app.HTMLButton) {
		button.OnClick(func(ctx app.Context, e app.Event) {
			ctx.NewActionWithValue(string(banner.Open), c.floating)
		})
	}}
	openFixed := &button.Button{Id: c.UUID(), Label: "fixed", Callback: func(button app.HTMLButton) {
		button.OnClick(func(ctx app.Context, e app.Event) {
			ctx.NewActionWithValue(string(banner.Open), c.fixed)
		})
	}}

	// Label: "centered"

	checkBox := &checkbox.Checkbox{Id: c.UUID(), Callback: func(input app.HTMLInput) {
		input.OnClick(func(ctx app.Context, e app.Event) {
			centeredValue := ctx.JSSrc().Get("checked").Bool()
			c.floating.Centered = centeredValue
			c.floating.Update()
			c.fixed.Centered = centeredValue
			c.fixed.Update()
		})
	}}

	centered := &formfield.FormField{Component: checkBox, Label: "centered"}

	body := app.Div().Body(
		c.floating, c.fixed,
		layout.Grid().Body(layout.Inner().Body(
			layout.Cell().Body(openFloating, openFixed, centered),
			layout.CellModified("middle", 12).Body(c.message),
		)))
	return PageBody(body)

}

func (c *BannerDemo) OnMount(ctx app.Context) {
	// handle all banner events
	for _, n := range []banner.EventType{banner.Opening, banner.Opened, banner.Closing, banner.Closed} {
		ctx.Handle(string(n), c.actionHandler)
	}
}

func (c *BannerDemo) actionHandler(ctx app.Context, action app.Action) {
	if !(action.Value == c.fixed || action.Value == c.floating) {
		return
	}
	if b, ok := action.Value.(*banner.Banner); ok {
		c.message.Text = fmt.Sprintf("message from banner %q: Event=%25s Tags=%v", b.Id, action.Name, action.Tags)
		c.message.Update()
	}
}
