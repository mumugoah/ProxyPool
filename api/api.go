package api

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/mumugoah/ProxyPool/storage"
	"github.com/mumugoah/ProxyPool/util"
)

// VERSION for this program
const VERSION = "/v1"

// Run for request
func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc(VERSION+"/ip", ProxyHandler)
	mux.HandleFunc(VERSION+"/https", FindHandler)
	log.Println("Starting server at port", util.NewConfig().Port)
	http.ListenAndServe(fmt.Sprintf(":%s", util.NewConfig().Port), mux)
}

// ProxyHandler .
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyRandom())
		if err != nil {
			return
		}
		w.Write(b)
	}
}

// FindHandler .
func FindHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		b, err := json.Marshal(storage.ProxyFind("https"))
		if err != nil {
			return
		}
		w.Write(b)
	}
}
