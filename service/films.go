package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/Cswapi/Web-Server/database/database"
)

// 分页查询资源
func filmsHandler(r *render.Render) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Access Origin", "*")
		req.ParseForm()
		page := 1
		res.Write([]byte("{\"result\" : \n["))
		if req.Form["page"] != nil {
			page, _ = strconv.Atoi(req.Form["page"][0])
		}
		// 页数量
		count := 0
		for i := 1; ; i++ {
			item := database.GetValue([]byte("films"), []byte(strconv.Itoa(i)))
			if len(item) != 0 {
				count++
				if count > pagelen*(page-1) {
					res.Write([]byte(item))
					if count >= pagelen*page || count >= database.GetBucketCount([]byte("films")) {
						break
					}
					res.Write([]byte(", \n"))
				}
			}
		}
		res.Write([]byte("]\n}"))
	}
}

// 按ID查询资源的API的处理函数
func getFilmsById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	res.WriteHeader(http.StatusOK)
	// 调用database中的GetValue函数获取键（ID）对应的值
	value := database.GetValue([]byte("films"), []byte(vars["id"]))
	res.Write([]byte(value))
}

// 查询页数的API的处理函数
func filmsPagesHandler(res http.ResponseWriter, req *http.Request) {
	// 直接调用database中的GetBucketCount函获取每种资源数据库中的桶的数量
	counts := database.GetBucketCount([]byte("films"))
	res.Write([]byte(strconv.Itoa(counts)))
}
