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
	/*
			Debug
		log.Println("itens: ", len(data.Words))
		for matter, words := range data.Words {
			for _, word := range words {
				log.Printf("matter: %s - word: %+v\n", matter, word)
			}
		}
	*/
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
