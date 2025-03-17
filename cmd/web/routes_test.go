package main

import (
	"fmt"
	"testing"

	"github.com/Anahita-ghloo/foyer_international/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("The type is not *chi.Mux, it is of type %T", v))
	}
}
