package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pedrosantosbr/proto-hornex/infra/services"
)

type UserService interface {
	Create(ctx context.Context, params services.UserCreateParams) (internal.User, error)
}

type UserHandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(r *chi.Mux) {
	r.Post("api/users", h.create)
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	h.svc.Create()
}
