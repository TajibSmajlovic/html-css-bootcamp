package main

import (
	    "fmt"
    	"net/http"
    	"google.golang.org/appengine"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	appengine.Main()
}