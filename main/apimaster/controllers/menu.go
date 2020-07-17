package controllers

import (
	"encoding/json"
	"gomux/main/apimaster/models"
	"gomux/main/apimaster/usecases"
	"gomux/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//MenuHandler MenuHandler
type MenuHandler struct {
	MenuUseCase usecases.MenuUseCase
}

//MenuController MenuController
func MenuController(r *mux.Router, service usecases.MenuUseCase) {
	MenuHandler := MenuHandler{service}
	r.HandleFunc("/allmenus", MenuHandler.AllMenus).Methods(http.MethodGet)
	r.HandleFunc("/menu/{id}", MenuHandler.MenuByID).Methods(http.MethodGet)
	r.HandleFunc("/menu", MenuHandler.AddMenu).Methods(http.MethodPost)
	r.HandleFunc("/menu/{id}", MenuHandler.UpdateMenu).Methods(http.MethodPut)
	r.HandleFunc("/menu/{id}", MenuHandler.DeleteMenu).Methods(http.MethodDelete)
}

//AllMenus AllMenus
func (s MenuHandler) AllMenus(w http.ResponseWriter, r *http.Request) {
	allMenus, err := s.MenuUseCase.GetAllMenus()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data Not Found")
	}
	menuResponse := utils.Response{Status: http.StatusOK, Messages: "Data of Menu", Data: allMenus}
	byteOfMenuResponse, err := json.Marshal(menuResponse)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
		log.Println("Oops something when wrong")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(byteOfMenuResponse)
}

//AddMenu InsertMenu
func (s MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	var newMenu models.Menu
	_ = json.NewDecoder(r.Body).Decode(&newMenu)
	err := s.MenuUseCase.AddMenu(newMenu)
	if err != nil {
		w.Write([]byte("Insert Failed Cannot null"))
		log.Print(err)
	} else {
		w.Write([]byte("Insert Success"))
		log.Println("Insert Success")
	}
}

//MenuByID MenuById
func (s MenuHandler) MenuByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	menu, err := s.MenuUseCase.GetMenuByID(idMenu)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data not found")
	}
	menuResponse := utils.Response{Status: http.StatusOK, Messages: "Data of Menu", Data: menu}
	byteOfMenuResponseByID, err2 := json.Marshal(menuResponse)
	if err2 != nil {
		w.Write([]byte("Oops something when wrong"))
		log.Println("Oops something when wrong")
	} else if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfMenuResponseByID)
	}
}

// UpdateMenu UpdateMenu
func (s MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	var changeMenu models.Menu
	_ = json.NewDecoder(r.Body).Decode(&changeMenu)
	err := s.MenuUseCase.UpdateMenusByID(idMenu, changeMenu)
	if err != nil {
		w.Write([]byte("Id Not Found"))
		log.Println("Data not found")
	} else {
		w.Write([]byte("Data Updated"))
		log.Println("Data Updated")
	}

}

//DeleteMenu DeleteById
func (s MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&s)
	err := s.MenuUseCase.DeleteMenusByID(idMenu)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data not found")
	} else {
		w.Write([]byte("Data Deleted"))
		log.Println("Data Deleted")
	}
}