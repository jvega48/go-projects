package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", moneyFuntion)
	http.HandleFunc("/time_stamp", timeStamp)
	http.ListenAndServe(":8080", nil)
}

func moneyFuntion(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world"))
}

func timeStamp(a http.ResponseWriter, r *http.Request) {
	fmt.Println("The time route", time.Now())
}
