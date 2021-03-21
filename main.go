package main

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"

	"encoding/json"
)

type Stock struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Currency string `json:"currency"`
	Price    string `json:"price"`
}

var stocks []Stock

func getGreeting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello: %v!\n", vars["name"])
	fmt.Printf("Greeting served\n")
}

// func getStockInfo(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, stock := range stocks {
// 		if stock.ID == params["id"] {
// 			json.NewEncoder(w).Encode(stock)
// 			fmt.Printf("Stock info --> %s <-- served\n", stock.Name)
// 			break
// 		}
// 		return
// 	}
// }

func fetchETFInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fetchResult := `[{
		"ID": "0000",
		"name": "Fake Stock Inc",
		"currency": "BTC",
		"price": "20"
	},{
		"ID": "0001",
		"name": "Fake Stock Inc 2",
		"currency": "ARS",
		"price": "20"
	},{
		"ID": "0002",
		"name": "Fake Stock Inc 3",
		"currency": "USD",
		"price": "20"
	},{
		"ID": "0003",
		"name": "Fake Stock Inc 4",
		"currency": "EUR",
		"price": "20"
	}]`

	// var fetchedStocks []Stock

	err := json.Unmarshal([]byte(fetchResult), &stocks)

	if err != nil {
		fmt.Println(err)
		return
	}

	// stocks = append(stocks, fetchedStocks)

	for _, stock := range stocks {
		if stock.ID == params["id"] {
			json.NewEncoder(w).Encode(stock)
			fmt.Printf("Stock info --> %s <-- served\n", stock.Name)
			break
		}
		return
	}

}

func main() {
	router := mux.NewRouter()

	// stocks = append(stocks, Stock{ID: "0001", Name: "Fake Stock Inc", Currency: "BTC", Price: "20 btc"})

	router.HandleFunc("/greetings/{name}", getGreeting).Methods("GET")
	// router.HandleFunc("/stock/{id}", getStockInfo).Methods("GET")
	router.HandleFunc("/fetch/etf/{id}", fetchETFInfo).Methods("GET")

	fmt.Printf("Listening on port 8000\n")
	http.ListenAndServe(":8000", router)
}
