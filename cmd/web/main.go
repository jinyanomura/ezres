package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jinyanomura/ezres-web/pkg/config"
	"github.com/jinyanomura/ezres-web/pkg/driver"
	"github.com/jinyanomura/ezres-web/pkg/handlers"
	"github.com/jinyanomura/ezres-web/pkg/helpers"
	"github.com/jinyanomura/ezres-web/pkg/render"
)

var (
	app config.AppConfig
	session *scs.SessionManager
	port string
)

func main() {
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbHost := flag.String("dbhost", "localhost", "Database host")
	dbName := flag.String("dbname", "", "Database name")
	dbUser := flag.String("dbuser", "", "Database user")
	dbPass := flag.String("dbpass", "", "Database password")
	dbPort := flag.String("dbport", "5432", "Database port")
	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require))")

	flag.Parse()

	app.InProduction = *inProduction
	app.UseCache = *useCache

	if !app.InProduction && (*dbName == "" || *dbUser == "") {
		log.Fatal("Missing required flags")
	}

	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	var connectionString string
	if app.InProduction {
		connectionString = os.Getenv("DATABASE_URL")
	} else {
		connectionString = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	}

	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	log.Println("Connected to database")

	c, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = c

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
	render.SetNewTemplates(&app)

	if app.InProduction {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":8080"
	}

	fmt.Printf("Starting server on port %s\n", port)
	srv := &http.Server{
		Addr: port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("cannot connect to server", err)
	}
}