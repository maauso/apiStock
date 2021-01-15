package getjsonresponse

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//GetResponse Json response from API
func GetResponse(domain, endpoint, company, apikey string) []byte {

	url := domain + endpoint + company + "?&apikey=" + apikey
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData

}
