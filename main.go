package main

import (
	"DRSCheck/DRS"
	"net/http"
)

func main() {
	http.HandleFunc("/", DRS.DRSHandler)
	http.ListenAndServe(":8585", nil)
}
