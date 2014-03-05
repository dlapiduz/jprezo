package main

import (
	"github.com/codegangsta/martini"
	"github.com/eaigner/hood"
	"github.com/martini-contrib/render"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	hd, err := hood.Load("db/config.json", "development")
	if err != nil {
		panic(err)
	}

	m.Get("/", func(r render.Render) {
		var results []Presentation
		err = hd.OrderBy("title").Find(&results)
		if err != nil {
			panic(err)
		}

		context := map[string]interface{}{
			"results": results,
			"title":   "My presentations",
		}

		r.HTML(200, "presentations/index", context)
	})

	m.Get("/edit/:id", func(r render.Render, params martini.Params) {
		var presentation []Presentation
		err = hd.Where("id", "=", params["id"]).Limit(1).Find(&presentation)
		if err != nil {
			panic(err)
		}

		context := map[string]interface{}{
			"presentation": presentation,
			"title":        "Edit Presentation",
		}

		r.HTML(200, "presentations/edit", context)
	})

	m.Run()
}
