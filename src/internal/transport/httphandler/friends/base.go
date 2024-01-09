package friendshttphandler

import . "github.com/yigithankarabulut/simplemedia/src/internal/service/friends"

type Handler struct {
	service FriendService
}

func NewHandler(service FriendService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) AddRoutes() {

}
