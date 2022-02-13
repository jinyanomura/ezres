package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jinyanomura/ezres/pkg/config"
	"github.com/jinyanomura/ezres/pkg/handlers"
	"github.com/jinyanomura/ezres/pkg/helpers"
	"github.com/jinyanomura/ezres/pkg/render"
)

var (
	app config.AppConfig
	session *scs.SessionManager
	infoLog *log.Logger
	errorLog *log.Logger
	port string
)

func main() {
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")

	flag.Parse()

	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	c, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = c

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
	render.SetNewTemplates(&app)

	if app.InProduction {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":8080"
	}

	fmt.Println("starting server on port 8080...")
	srv := &http.Server{
		Addr: port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("cannot connect to server", err)
	}
}