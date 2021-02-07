package financialmodelingprep

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// getResponse Json response from API
func getResponse(domain, endpoint, company, apikey string) []byte {
	url := domain + endpoint + company + "?&apikey=" + apikey
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
	}
	return responseData
}
