package demo

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-mdc/pkg/base"
	"github.com/mlctrez/goapp-mdc/pkg/icon"
	"github.com/mlctrez/goapp-mdc/pkg/layout"
)

type IconDemo struct {
	app.Compo
	base.JsUtil
}

func (d *IconDemo) Render() app.UI {

	toggleOne := &icon.Toggle{Id: d.UUID(), Icon: icon.MIFavorite,
		IconOff: icon.MIFavoriteBorder, AriaOn: "remove from favorites", AriaOff: "add to favorites"}
	toggleTwo := &icon.Toggle{Id: d.UUID(), Icon: icon.MIFavorite,
		IconOff: icon.MIFavoriteBorder, AriaOn: "remove from favorites", AriaOff: "add to favorites"}

	body := layout.Grid().Body(
		layout.Inner().Style("display", "flex").Body(
			layout.Cell().Body(
				&icon.Button{Id: d.UUID(), Icon: icon.MIBookmark, AriaLabel: "bookmark this"}),
			layout.Cell().Body(toggleOne, toggleTwo),
		),
	)

	return PageBody(body)
}
