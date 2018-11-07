package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

type Bowl struct {
  Name    string `json:"name,omitempty"`
  Ingredients []Ingredient
}

type Cake struct {
  Name    string `json:"name,omitempty"`
  Ingredients []Ingredient
  CookedAtTemperature int `json:"cookedattemperature,omitempty"`
}

type Ingredient struct {
  Name     string `json:"name,omitempty"`
  Quantity string `json:"quantity,omitempty"`
}

type Message struct {
  Message     string `json:"message,omitempty"`
}

type OvenTemp struct {
  Temperature int `json:"temperature,omitempty"`
}

var Temperature = 0

func main() {
  log.Println("Starting go_cake_api_bake")
  router := mux.NewRouter()
  router.HandleFunc("/heatoven", PostHeatOven).Methods("POST")
  router.HandleFunc("/bake", PostBake).Methods("POST")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func PostHeatOven(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var t OvenTemp
  err := decoder.Decode(&t)
  message := ""
  if err != nil {
    message =  "Couldn't parse form data"
    rw.WriteHeader(http.StatusInternalServerError)
  } else {
    message = "Oven heating"
    Temperature = t.Temperature
  }
  json.NewEncoder(rw).Encode(Message{Message: message})
  log.Println(message)
}

func PostBake(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var t Bowl
  err := decoder.Decode(&t)
  message := ""
  if err != nil {
    message =  "Couldn't parse form data"
    rw.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(rw).Encode(Message{Message: message})
  } else {
    message = "Baking Cake"
    newCake := Cake{Name: t.Name, Ingredients: t.Ingredients, CookedAtTemperature: Temperature}
    rw.WriteHeader(http.StatusCreated)
    json.NewEncoder(rw).Encode(newCake)
  }
  log.Println(message)
}
