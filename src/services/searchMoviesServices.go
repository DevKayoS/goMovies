package services

import (
	"net/http"

	"github.com/DevKayoS/goMovies/src/omdb"
	"github.com/DevKayoS/goMovies/src/utils"
)

func HandleSearchMovie(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		result, err := omdb.Search(apiKey, search)

		if err != nil {
			utils.SendJson(w, utils.Response{Error: "something wrong with omdb"}, http.StatusBadGateway)
			return
		}

		utils.SendJson(w, utils.Response{Data: result}, http.StatusOK)
	}
}
