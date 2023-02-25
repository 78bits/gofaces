package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

var currentPage gofaces.Page

func runws() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Header("content-type", "text/html")
		c.String(200, currentPage.Render())
	})

	r.GET("/menu", func(c *gin.Context) { // this could be loaded in bg later...
		itemidx, err := strconv.Atoi(c.Query("item"))
		if err != nil {
			log.Println(err.Error())
		}
		currentPage.GetMenu().ItemSelected = itemidx
		c.Redirect(301, "/")
	})
	r.Run()
}

func main() {
	// i18n := NewI18n()
	currentPage = gofaces.NewPage("Testpage")

	currentPage.SetMenu(&gofaces.Menu{
		Items: []gofaces.MenuItem{
			{
				Label:   "eins",
				Iconset: ICONSET_FEATHER,
				Icon:    "alert-circle",
			},
			{
				Label:   "zwei",
				Iconset: ICONSET_FEATHER,
				Icon:    "alert-circle",
			},
			{
				Group:   "Eine Gruppe",
				Label:   "drei",
				Iconset: ICONSET_FEATHER,
				Icon:    "alert-circle",
			},
		},
		ItemSelected: 1,
	})

	currentPage.SetContent(&gofaces.Content{
		Children: []gofaces.VC{
			&Panel{Title: "We won the war",
				Children: []gofaces.VC{
					&Text{Text: "Hallo Welt"},
				},
			},
		},
	})

	runws()
}
