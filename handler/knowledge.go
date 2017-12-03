package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/novatrixtech/silva-backend/data"
	"github.com/novatrixtech/silva-backend/model"
)

//ListKnowledges handles request to list knowledges known by the Bot
func ListKnowledges(c echo.Context) (err error) {

	resp := []model.Knowledge{}
	for knowledge := range data.KnowledgeBase {
		tmp := model.Knowledge{}
		tmp.Name = knowledge
		resp = append(resp, tmp)
	}
	return c.JSON(http.StatusOK, resp)
}
