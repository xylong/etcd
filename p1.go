package main

import (
	"github.com/gorilla/mux"
	"go-etcd/util"
	"log"
	"net/http"
)

func main() {
	router:=mux.NewRouter()

	router.HandleFunc("/product/{id:\\d+}", func(w http.ResponseWriter,r *http.Request) {
		vars:=mux.Vars(r)
		str:="get product byID:"+vars["id"]
		w.Write([]byte(str))
	})

	s:=util.NewService()
	if err:=s.RegService("p1","productservice","192.168.159.46:8081");err!=nil {
		log.Fatalln(err)
	}

	http.ListenAndServe(":8081",router)
}
