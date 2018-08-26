package main

import (
	"github.com/kataras/iris"
)

var page = struct {
	Title string
}{"Welcome"}

func newApp() *iris.Application {
	app := iris.New()
	app.RegisterView(iris.HTML("./static", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Page", page)
		ctx.View("index.html")
	})

	// or just serve index.html as it is:
	// app.Get("/{f:path}", func(ctx iris.Context) {
	// 	ctx.ServeFile("index.html", false)
	// })

	assetHandler := app.StaticHandler("./static", false, false)
	// as an alternative of SPA you can take a look at the /routing/dynamic-path/root-wildcard
	// example too
	app.SPA(assetHandler)

	return app

}
func main() {

	app := newApp()
	app.Run(iris.Addr(":8000"))

}
