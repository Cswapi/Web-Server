package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/Cswapi/Web-Server/database/database"
)

func starshipsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access Origin", "*")
		req.ParseForm()
		page := 1
		w.Write([]byte("{\"result\" : \n["))
		if req.Form["page"] != nil {
			page, _ = strconv.Atoi(req.Form["page"][0])
		}
		count := 0
		for i := 1; ; i++ {
			item := database.GetValue([]byte("starships"), []byte(strconv.Itoa(i)))
			if len(item) != 0 {
				count++
				if count > pagelen*(page-1) {
					w.Write([]byte(item))
					if count >= pagelen*page || count >= database.GetBucketCount([]byte("starships")) {
						break
					}
					w.Write([]byte(", \n"))
				}
			}
		}
		w.Write([]byte("]\n}"))
	}
}

func getStarshipsById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	value := database.GetValue([]byte("starships"), []byte(vars["id"]))
	w.Write([]byte(value))
}

func starshipsPagesHandler(w http.ResponseWriter, req *http.Request) {
	counts := database.GetBucketCount([]byte("starships"))
	w.Write([]byte(strconv.Itoa(counts)))
}
