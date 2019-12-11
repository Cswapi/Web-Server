package service

import (
	"net/http"

	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var pagelen = 5

// api root 处理：得到一个API目录结构.
func apiRootHandler(formatter *render.Render) negroni.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		formatter.JSON(w, http.StatusOK, struct {
			Films     string `json:"films"`
			People    string `json:"people"`
			Planets   string `json:"planets"`
			Species   string `json:"species"`
			Starships string `json:"starships"`
			Vehicles  string `json:"vehicles"`
		}{Films: "https://swapi.co/api/films/",
			People:    "https://swapi.co/api/people/",
			Planets:   "https://swapi.co/api/planets/",
			Species:   "https://swapi.co/api/species/",
			Starships: "https://swapi.co/api/starships/",
			Vehicles:  "https://swapi.co/api/vehicles/"})
	}
}
