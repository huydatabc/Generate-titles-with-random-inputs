package route

import (
	"GenNameFromKey/handler"
	"GenNameFromKey/service"
	"github.com/go-chi/chi"
)

func Route(r *chi.Mux) error {
	var a, b, c, d []string
	genService, err := service.NewGenService(a, b, c, d)
	if err != nil {
		return err
	}
	genHandler := Handler.NewGenHandler(genService)

	genRoute := "/gen"
	r.Post(genRoute, genHandler.Generate)
	return nil
}
