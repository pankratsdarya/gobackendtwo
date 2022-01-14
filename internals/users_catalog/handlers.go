package users_catalog

import (
	"net/http"

	"github.com/pankratsdarya/gobackendtwo/internals/users_catalog/model"
)

// NewUserHandler adds a new user
func newUserHandler(item model.CatalogItem) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

// NewSettingHandler adds a new setting
func newSettingHandler(item model.CatalogItem) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
