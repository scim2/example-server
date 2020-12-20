package server

import (
	"github.com/elimity-com/scim"
	"net/http"
)

type usersResourceHandler struct{}

func newUsersResourceHandler() *usersResourceHandler {
	return new(usersResourceHandler)
}

func (u usersResourceHandler) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	panic("implement me")
}

func (u usersResourceHandler) Get(r *http.Request, id string) (scim.Resource, error) {
	panic("implement me")
}

func (u usersResourceHandler) GetAll(r *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	panic("implement me")
}

func (u usersResourceHandler) Replace(r *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	panic("implement me")
}

func (u usersResourceHandler) Delete(r *http.Request, id string) error {
	panic("implement me")
}

func (u usersResourceHandler) Patch(r *http.Request, id string, request scim.PatchRequest) (scim.Resource, error) {
	panic("implement me")
}
