package main

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"

	"encoding/json"
)

func fetchETFInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fetchResult := `{
		"ErrorMessage": null,
		"DiagnosticInfoJson": null,
		"Data": [
			{
				"S": {
					"OFST020000": "IE00BYWQWR46",
					"OFST900016": "VanEck Vec. UCITS ETFs plc - Vid. Gam. and eSports UCITS ETF",
					"OFST020540": "USD",
					"OFST020262": "no",
					"OFST020050": "A",
					"OFST020400": "accumulating",
					"OFST6030CH": "YES",
					"OFST6031CH": "YES",
					"OFST900025": "5d24c92f-5cc1-474f-869e-b2ce2f639607",
					"OFST452000": "0.0055",
					"OFST452001": "2020-06-24",
					"OFST452200": "0.0055",
					"OFST452220": "2020-06-24",
					"OFST452100": "0.0055",
					"OFST452110": "2020-06-24",
					"OFST452120": "0.0055",
					"OFST452130": "2020-06-24",
					"OFST023200": "MVIS Global Video Gaming and eSports Index",
					"OFST900000": "dffe04b2-8bfc-48d1-86bb-befd09125a47",
					"OFST001000": "VanEck",
					"OFST900017": "56ae1c2a-8009-42f8-8e62-eb65b775bb2e",
					"OFST010410": "USD",
					"OFDY901035": "41.4697|2021-03-23|USD",
					"OFDY100110": "",
					"OFDY901041": "-0.019257|2021-03-22|USD",
					"OFDY908009": "989051487.44|2021-03-23|USD",
					"OFDY908011": "989051487.44|2021-03-23|USD",
					"OFDY908104": "-0.00802|2021-02-26|USD",
					"OFDY908106": "0.100536|2021-02-26|USD",
					"OFDY908108": "0.163148|2021-02-26|USD",
					"OFDY908110": "0.872911|2021-02-26|USD",
					"OFDY908112": "",
					"OFDY908114": "",
					"FIF": "0",
					"MEMBER": "0"
				},
				"D": {
					"AR": [
						{
							"DocType": "AR",
							"Label": "AR",
							"Url": "https://api.fundinfo.com/document/ab79736cce5ea957eae4a77323acf975_1670867/AR_ES_en_IE00BYWQWR46_YES_2020-04-30.pdf?apiKey=b9934aa2-1a83-4286-b11b-c8415da9e581",
							"Date": "2020-04-30",
							"Language": "EN",
							"Active": true,
							"Country": "ES",
							"DocumentConstraint": "YES"
						}
					],
					"CON": [
						{
							"DocType": "CON",
							"Label": "CON",
							"Url": "https://api.fundinfo.com/document/8222c2f3f53d602e74291716d1071f90_652349/CON_ES_en_IE00BYWQWR46_YES_2016-04-25.pdf?apiKey=b9934aa2-1a83-4286-b11b-c8415da9e581",
							"Date": "2016-04-25",
							"Language": "EN",
							"Active": true,
							"Country": "ES",
							"DocumentConstraint": "YES"
						}
					],
					"INF": [
						{
							"DocType": "INF",
							"Label": "INF",
							"Url": "https://api.fundinfo.com/document/51543241c3c643f61008c405656d98ff_133289/INF_ES_en_IE00BYWQWR46_YES_2020-11-27.pdf?apiKey=b9934aa2-1a83-4286-b11b-c8415da9e581",
							"Date": "2020-11-27",
							"Language": "EN",
							"Active": true,
							"Country": "ES",
							"DocumentConstraint": "YES"
						}
					],
					"KID": [
						{
							"DocType": "KID",
							"Label": "KIID",
							"Url": "https://api.fundinfo.com/document/13e29b12da01aed1288d016ac779be74_130966/KID_ES_es_IE00BYWQWR46_YES_2021-02-11.pdf?apiKey=b9934aa2-1a83-4286-b11b-c8415da9e581",
							"Date": "2021-02-11",
							"Language": "ES",
							"Active": false,
							"Country": "ES",
							"DocumentConstraint": "YES"
						}
					],
					"MR": [
						{
							"DocType": "MR",
							"Label": "MR",
							"Url": "https://api.fundinfo.com/document/0600b34c4193f467b7ca09902b0653ac_99717/MR_ES_en_IE00BYWQWR46_YES_2021-01-31.pdf?apiKey=b9934aa2-1a83-4286-b11b-c8415da9e581",
							"Date": "2021-01-31",
							"Language": "EN",
							"Active": true,
							"Country": "ES",
							"DocumentConstraint": "YES"
						}
					],
					"PR": [
						{
							"DocType": "PR",
							"Label": "PR",
							"Url": "https://api.fundinfo.com/document/87d5d3d458430265b539a6a9eb0fc793_2172951/PR_ES_en_IE00BYWQWR46_YES_2020-12-22.pdf?apiKey=b9934aa2-1a83-4286-b11b-c8415da9e581",
							"Date": "2020-12-22",
							"Language": "EN",
							"Active": true,
							"Country": "ES",
							"DocumentConstraint": "YES"
						}
					]
				},
				"V": null,
				"N": null,
				"R": null,
				"LNA": false,
				"LNL": null,
				"Member": false
			}
		],
		"Total": 1,
		"TotalFunds": 1
	}`

	var result map[string]interface{}

	err := json.Unmarshal([]byte(fetchResult), &result)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := result["Data"].([]interface{})
	first := data[0].(map[string]interface{})
	s := first["S"].(map[string]interface{})

	json.NewEncoder(w).Encode(s)

	fmt.Printf("ETF name --> %s - %s\n", s["OFST001000"], s["OFST023200"])
	fmt.Printf("Stock price|date|currency --> %s\n", s["OFDY901035"])
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/fetch/etf/{id}", fetchETFInfo).Methods("GET")

	fmt.Printf("Listening on port 8000\n")
	http.ListenAndServe(":8000", router)
}
