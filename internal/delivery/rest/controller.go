package rest

import (
	"encoding/json"
	"log"
	"net/http"
	request "testDeployment/internal/delivery/http"
	"testDeployment/pkg/Bot"
	"testDeployment/internal/usecase"
)

type controller struct {
	usecase usecase.Usecase
	bot     Bot.Bot
	http request.CustomJSONRequester
}

func NewController(usecase usecase.Usecase, bot Bot.Bot,request request.CustomJSONRequester ) *controller {
	return &controller{usecase: usecase,
		bot: bot,
		http:request,
	}
}

func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	user := c.usecase.GetAll()
	log.Println(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}
