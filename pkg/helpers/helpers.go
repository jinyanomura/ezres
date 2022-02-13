package helpers

import "github.com/jinyanomura/ezres/pkg/config"

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}