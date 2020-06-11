package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var serviceURLs = make(map[string]string)

// Get URL values from JSON file
func getServiceURLs() {
	jsonFile, err := os.Open("serviceURLs.json")
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal([]byte(byteValue), &serviceURLs)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(serviceURLs)
}

// Get the url for a given proxy condition
func getProxyURL(operator string) (url string, found bool) {
	operator = strings.ToUpper(operator)

	log.Println(serviceURLs[operator])

	url, found = serviceURLs[operator]
	if !found {
		log.Println("No operator " + operator + " found")
	}
	return
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	res.Header().Set("Access-Control-Allow-Origin", "*")

	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func getOperator(res http.ResponseWriter, req *http.Request) {
	operator := mux.Vars(req)["operator"]
	log.Println(operator)

	if url, found := getProxyURL(operator); found {
		serveReverseProxy(url, res, req)
	} else {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, "404 operator does not exist")
	}
}

func main() {
	// Setup values
	getServiceURLs()
	port := serviceURLs["PROXY_PORT"]

	// start server
	router := mux.NewRouter()
	router.HandleFunc("/{operator:[a-z]+}", getOperator).Methods("GET")
	router.HandleFunc("/{operator:[a-z]+}/", getOperator).Methods("GET") //Thanks PHP

	log.Println(port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalln(err)
	}
}
