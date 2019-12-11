package service

import (
	"fmt"
	"os"

	"github.com/Cswapi/Web-Server/models"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"

	"github.com/Cswapi/Web-Server/database/database"
)

// Negroni库开发中间件
func NewServer() *negroni.Negroni {

	r := render.New(render.Options{
		Directory:  "templates",       // 指定从哪个路径加载模板
		Extensions: []string{".html"}, // 指定要加载模板的扩展名为html
		IndentJSON: true,              // 输出可读的JSON
	})
	/* 实例化一个negroni
	 * 添加3个组件NewRecovery()（Panic Recovery）,
	 * NewLogger()（日志处理）
	 * NewStatic（静态文件服务器）
	 */
	n := negroni.Classic()
	// 实例化一个mux.Router
	router := mux.NewRouter()
	// registerRouter将处理函数（Handler）注册到路由中
	registerRouter(router, r)
	// 添加中间件到处理链中
	n.UseHandler(router)
	return n
}

func registerRouter(router *mux.Router, r *render.Render) {
	// API root获取API目录结构
	apiRoot := os.Getenv("WEBROOT")
	if len(apiRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("The directory could not be detected!")
		} else {
			apiRoot = root
			fmt.Println(root)
		}
	}
	// 开启数据库
	database.Start("database/database/Cswapi.db")

	router.Handle("/api/", negroni.New(
		negroni.HandlerFunc(models.ValidateMid),
		negroni.HandlerFunc(apiRootHandler(r)),
	))
	// 将处理函数注册到路由中，并添加http方法
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/register", registerHandler).Methods("POST")

	router.HandleFunc("/api/films/", filmsHandler(r)).Methods("GET")
	router.HandleFunc("/api/films/pages", filmsPagesHandler).Methods("GET")
	filmsOne := router.PathPrefix("/api/films").Subrouter()
	filmsOne.HandleFunc("/{id:[0-9]+}", getFilmsById).Methods("GET")

	router.HandleFunc("/api/people/", peopleHandler(r)).Methods("GET")
	router.HandleFunc("/api/people/pages", peoplePagesHandler).Methods("GET")
	peopleOne := router.PathPrefix("/api/people").Subrouter()
	peopleOne.HandleFunc("/{id:[0-9]+}", getPeopleById).Methods("GET")

	router.HandleFunc("/api/planets/", planetsHandler(r)).Methods("GET")
	router.HandleFunc("/api/planets/pages", planetsPagesHandler).Methods("GET")
	planetsOne := router.PathPrefix("/api/planets").Subrouter()
	planetsOne.HandleFunc("/{id:[0-9]+}", getPlanetsById).Methods("GET")

	router.HandleFunc("/api/species/", speciesHandler(r)).Methods("GET")
	router.HandleFunc("/api/species/pages", speciesPagesHandler).Methods("GET")
	speciesOne := router.PathPrefix("/api/species").Subrouter()
	speciesOne.HandleFunc("/{id:[0-9]+}", getSpeciesById).Methods("GET")

	router.HandleFunc("/api/starships/", starshipsHandler(r)).Methods("GET")
	router.HandleFunc("/api/starships/pages", starshipsPagesHandler).Methods("GET")
	starshipsOne := router.PathPrefix("/api/starships").Subrouter()
	starshipsOne.HandleFunc("/{id:[0-9]+}", getStarshipsById).Methods("GET")

	router.HandleFunc("/api/vehicles/", vehiclesHandler(r)).Methods("GET")
	router.HandleFunc("/api/vehicles/pages", vehiclesPagesHandler).Methods("GET")
	vehiclesOne := router.PathPrefix("/api/vehicles").Subrouter()
	vehiclesOne.HandleFunc("/{id:[0-9]+}", getVehiclesById).Methods("GET")
}
