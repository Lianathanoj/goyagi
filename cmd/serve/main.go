// cmd/serve/main.go
package main

import (
    "net/http"

    "github.com/Lianathanoj/goyagi/pkg/application"
    "github.com/Lianathanoj/goyagi/pkg/server"
    "github.com/lob/logger-go"
)

func main() {
    log := logger.New()
    app, err := application.New()
    if err != nil {
        log.Err(err).Fatal("failed to initialize application")
    }

    srv := server.New(app)

    log.Info("server started", logger.Data{"port": app.Config.Port})

    err = srv.ListenAndServe()
    if err != nil && err != http.ErrServerClosed {
        log.Err(err).Fatal("server stopped")
    }

    log.Info("server stopped")
}
