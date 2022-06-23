package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Stock struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability int     `json:"availability"`
	IsActive     bool    `json:"is_active"`
}

var inventory = map[string]Stock{}

func ss() {
	router := httprouter.New()

	router.POST("/stock", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var stock Stock
		err := decoder.Decode(&stock)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{
				"error": %s
			}`, err.Error())))
			return
		}

		stockID := uuid.New().String()
		inventory[stockID] = stock

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{
			"id": "%s"
		}`, stockID)))
	})

	router.GET("/stock/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		stockID := params.ByName("id")
		if len(stockID) < 1 {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{
				"error": "must specify id"
			}`))
			return
		}

		stock := inventory[stockID]
		if (stock == Stock{}) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{
				"error": "stock not found"
			}`))
			return
		}

		responseBody, err := json.Marshal(stock)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{
				"error": "%s"
			}`, err.Error())))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	})

	http.ListenAndServe(":8080", router)
}
