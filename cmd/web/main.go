package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// go run	 ./cmd/web -addr=":40000"

type application struct {
	logger *slog.Logger
}

func main() {
	// :4000 will be default value if no flag
	addr := flag.String("addr", ":4000", "HTTP net addr")
	flag.Parse() //this must be before using `addr`

	// nil when u want to use the default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("srv up on ", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
