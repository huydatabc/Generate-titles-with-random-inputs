package Handler

import (
	"GenNameFromKey/model"
	"GenNameFromKey/service"
	"encoding/json"
	"net/http"
)

type UserController struct {
	service *service.GenService
}

func NewGenHandler(h *service.GenService) *UserController {
	return &UserController{h}
}

func (c *UserController) Generate(w http.ResponseWriter, r *http.Request) {
	var req model.Input
	er1 := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.service.Gen([]byte(req.Id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
