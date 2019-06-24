// pkg/movies/routes_test.go
package movies

import (
    "testing"

    "github.com/Lianathanoj/goyagi/internal/test"
    "github.com/Lianathanoj/goyagi/pkg/application"
    "github.com/labstack/echo"
    "github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
    e := echo.New()
    app := application.App{}

    RegisterRoutes(e, app)

    routes := test.FilterRoutes(e.Routes())
    assert.Len(t, routes, 3)
    assert.Contains(t, routes, "GET /movies")
    assert.Contains(t, routes, "GET /movies/:id")
    assert.Contains(t, routes, "POST /movies")
}
