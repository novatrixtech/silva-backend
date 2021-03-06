package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/novatrixtech/silva-backend/data"
	"github.com/novatrixtech/silva-backend/handler"
	"github.com/novatrixtech/silva-backend/model"
)

func init() {
	data.KnowledgeBase = make(map[string][]model.Word, 0)
	data.KnowledgeBase["rios"] = data.LoadDataFromFiles("Silva-Conhecimento - rios.csv")
	data.KnowledgeBase["capitaisbrasil"] = data.LoadDataFromFiles("Silva-Conhecimento - capitaisbrasil.csv")
	data.KnowledgeBase["estadosbrasil"] = data.LoadDataFromFiles("Silva-Conhecimento - estadosbrasil.csv")
	data.KnowledgeBase["paises"] = data.LoadDataFromFiles("Silva-Conhecimento - paises.csv")
	data.KnowledgeBase["frutas"] = data.LoadDataFromFiles("Silva-Conhecimento - frutas.csv")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/newgame", handler.NewGame)
	e.GET("/knowledge", handler.ListKnowledges)
	e.Logger.Fatal(e.Start(":" + port()))
}

// configure http port
func port() (httpPort string) {
	httpPort = "5000"
	if os.Getenv("PORT") == "" {
		httpPort = os.Getenv("PORT")
	}
	return
}
