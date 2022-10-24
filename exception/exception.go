package exception

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BadRequest(context echo.Context) {
	context.String(http.StatusBadRequest, "400 Bad Request.")
}

func NotFound(context echo.Context) {
	context.String(http.StatusNotFound, "404 Not Found.")
}

func NotAllowed(context echo.Context) {
	context.String(http.StatusMethodNotAllowed, "405 Not Allowed.")
}

func InternalServerError(context echo.Context) {
	context.String(http.StatusInternalServerError, "500 Internal Server Error.")
}
