package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ctrl Controller) Version(context echo.Context) error {
	message := "1.0" // TODO: parameter store
	return context.String(http.StatusOK, message)
}

func (ctrl Controller) Log(context echo.Context) error {
	prm, err := bindJson(context)
	if err != nil {
		return err
	}

	context.Echo().Logger.Error("[" + ctrl.identification() + "] Log:" + prm.convString())
	return context.String(http.StatusOK, "")
}
