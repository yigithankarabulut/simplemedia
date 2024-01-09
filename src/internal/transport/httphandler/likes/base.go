package likeshttphandler

import . "github.com/yigithankarabulut/simplemedia/src/internal/service/likes"

type Handler struct {
	service LikesService
}

func NewHandler(service LikesService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes() {

}
