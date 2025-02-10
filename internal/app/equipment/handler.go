package equipment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/utils"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/models"
)

type equipmentHandler struct {
	eqipmentService Service
}

type Handler interface {
	CreateEquipmentHandler(w http.ResponseWriter, r *http.Request)
	ListEquipmentHandler(w http.ResponseWriter, r *http.Request)
	DeleteEquipmentHandler(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service Service) Handler {
	return &equipmentHandler{eqipmentService: service}
}

func (e *equipmentHandler) CreateEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	var equipment models.Equipment

	err := json.NewDecoder(r.Body).Decode(&equipment)

	if err != nil {

		http.Error(w, "Error while Decoding Request Body", http.StatusBadRequest)
	}

	resp, err := e.eqipmentService.CreateEquipment(context.Background(), equipment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.HandleResponse(w, resp, r)
}

func (e *equipmentHandler) ListEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	equipments, err := e.eqipmentService.GetAllEquipment(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	utils.HandleResponse(w, equipments, r)
}

func (e *equipmentHandler) DeleteEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("equipment_id")
	equipmentId, err := strconv.Atoi(id)

	if err != nil {
		resErr := fmt.Errorf("error while converting req param values form string into int: %v", err)
		http.Error(w, resErr.Error(), http.StatusBadRequest)
	}

	err = e.eqipmentService.DeleteEquipmentById(context.Background(), equipmentId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HandleResponse(w, "Equipment Deleted Successfully", r)
}

// func UpdateEquipmentHandler(w http.ResponseWriter, r *http.Request) {

// 	id := r.PathValue("equipment_id")

// 	equipmentId, err := strconv.Atoi(id)

// 	if err != nil {
// 		resErr := fmt.Errorf("error while converting equipment id param form string into int: %v", err)
// 		http.Error(w, resErr.Error(), http.StatusInternalServerError)
// 	}

// 	var equipment models.Equipment

// 	err = json.NewDecoder(r.Body).Decode(&equipment)

// 	if err != nil {

// 		http.Error(w, "Error while Decoding Request Body", http.StatusBadRequest)
// 	}

// 	err = repository.UpdateEquipment(equipmentId, equipment)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	utils.HandleResponse(w, "Updated Equipment Successfully", r)

// }
