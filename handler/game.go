package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/novatrixtech/silva-backend/data"
	"github.com/novatrixtech/silva-backend/model"
)

//NewGame handles new game request
func NewGame(c echo.Context) (err error) {
	req := new(model.Request)
	if err = c.Bind(req); err != nil {
		objReq := c.Request()
		body, errBody := ioutil.ReadAll(objReq.Body)
		if errBody != nil {
			log.Printf("Error reading body: %v", errBody)
			return
		}
		log.Println("[NewGame] Error during bind data to struct process. Body: ", body, " Error: ", err.Error())
		return
	}

	var wordsSet []model.Word

	for knowledge, words := range data.KnowledgeBase {
		if knowledge == strings.ToLower(strings.TrimSpace(req.Knowledge.Name)) {
			wordsSet = words
		}
	}
	resp := model.Response{}
	resp.Possibilities = data.EvaluateCondition(&wordsSet, req.Text, req.NonExistentChars)
	return c.JSON(http.StatusOK, resp)
}
