package main

import (
	"flag"
	data "movies_api/data"
	"movies_api/internal"
	"movies_api/internal/mailer"
	"net/http"
	"syscall"
)

type Config struct {
	Domain string
	Port   string
	DB     struct {
		DbHost     string
		Dbname     string
		Dbuser     string
		Dbport     string
		Dbpassword string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type Application struct {
	config *Config
	models data.Models
	mailer mailer.Mailer
	logger *internal.ConsoleLoger
}

func main() {
	var config Config

	flag.StringVar(&config.Domain, "domain", "localhost", "the http port the server will listen to")
	flag.StringVar(&config.Port, "port", ":8080", "the http port the server will listen to")

	flag.StringVar(&config.DB.DbHost, "dbHost", "localhost", "the the database host")
	flag.StringVar(&config.DB.Dbname, "dbname", "movies_api", "the name of the database used")
	flag.StringVar(&config.DB.Dbuser, "dbuser", "postgres", "the name of the database user")
	flag.StringVar(&config.DB.Dbport, "dbport", "5432", "the database port")
	flag.StringVar(&config.DB.Dbpassword, "dbpassword", "postgres", "the database user's password")

	flag.StringVar(&config.smtp.host, "smtp-host", "<your-smpt-host>", "SMTP host")
	flag.StringVar(&config.smtp.username, "smtp-username", "<your-smpt-username>", "SMTP username")
	flag.StringVar(&config.smtp.password, "smtp-password", "<your-smpt-password>", "SMTP password")
	flag.StringVar(&config.smtp.sender, "smtp-sender", "<the-sender-email>", "SMTP sender")
	flag.IntVar(&config.smtp.port, "smtp-port", 587, "SMTP port")

	flag.Parse()

	dbConn, err := config.OpenConnection()
	if err != nil {
		dbConn.Close()
		syscall.Exit(1)
	}
	defer dbConn.Close()

	logger := internal.NewConsoleLogger()

	app := Application{
		config: &config,
		models: data.NewModels(dbConn),
		mailer: mailer.New(config.smtp.host, config.smtp.port, config.smtp.username, config.smtp.password, config.smtp.sender),
		logger: logger,
	}

	server := &http.Server{
		Addr:    config.Port,
		Handler: app.routes(),
	}

	if err := server.ListenAndServe(); err != nil {
		app.logger.LogError.Println(err)
	}

}
