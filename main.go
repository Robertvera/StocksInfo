package main

import (
	"io/ioutil"

	"log"

	"net/http"

	"time"

	"fmt"

	"github.com/gorilla/mux"

	"encoding/json"
)

func fetchETFInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	isin := params["id"]

	url := fmt.Sprintf("https://www.etfinfo.com/es/DE-priv/LandingPage/Data?query=%s", isin)

	etfClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "etfInfo")

	res, getErr := etfClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	fetchResult, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var result map[string]interface{}

	jsonErr := json.Unmarshal([]byte(fetchResult), &result)

	if jsonErr != nil {
		fmt.Println(err)
		return
	}

	data := result["Data"].([]interface{})
	first := data[0].(map[string]interface{})
	s := first["S"].(map[string]interface{})

	json.NewEncoder(w).Encode(s)

	fmt.Printf("ETF name --> %s\n", s["OFST900016"])
	fmt.Printf("Stock price|date|currency --> %s\n", s["OFDY908007"])
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/fetch/etf/{id}", fetchETFInfo).Methods("GET")

	fmt.Printf("Listening on port 8000\n")
	http.ListenAndServe(":8000", router)
}
