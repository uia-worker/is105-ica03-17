package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {


	w.Write([]byte("<font color=\"green\">Hvordan g\xe5r det, <b>\u16a6</b> ?</font><br/>"))
	w.Write([]byte("\u16a6 - Thurs<br/>"))

}
