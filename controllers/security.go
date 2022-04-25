package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"stocksinfo.biz/app"
)

type SecurityController struct {
	*app.App
}

func NewSecurityController(a *app.App) *SecurityController {
	return &SecurityController{a}
}

func (security *SecurityController) FetchSecurityInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	isin := params["id"]

	baseUrl, err := url.Parse("https://tools.morningstar.co.uk/api/rest.svc/klr5zyak8x/security/screener")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	urlParams := url.Values{}
	urlParams.Add("outputType", "json")
	urlParams.Add("securityDataPoints", "SecId|Name|PriceCurrency|LegalName|ClosePrice|CategoryName")
	urlParams.Add("term", isin)

	baseUrl.RawQuery = urlParams.Encode()
	

	url := baseUrl.String()

	securityClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "securityInfo")

	res, getErr := securityClient.Do(req)
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

	rows := result["rows"].([]interface{})

	json.NewEncoder(w).Encode(rows)

	fmt.Printf("---- info --> %s\n", rows)
}

func (security *SecurityController) SaveSecurityInfo(w http.ResponseWriter, r *http.Request) {
	
}
