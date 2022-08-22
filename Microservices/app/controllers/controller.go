package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sQUARys/GO-pumping/app/services"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	Service services.Service
}

func New(service *services.Service) *Controller {
	return &Controller{
		Service: *service,
	}
}

func (ctr *Controller) GetOrderById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	idString := vars["id"]

	idInt, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("Error strconv in controller level : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order, err := ctr.Service.GetOrderById(idInt)
	if err != nil {
		log.Println("Error GetOrder in controller level : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println("Error json in controller level : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(orderJSON)

}
