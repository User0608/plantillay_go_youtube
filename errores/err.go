package errores

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func logErrs(err error) {
	log.Println(err.Error())
}
func GenErrResponseEcho(c echo.Context, err error) error {
	go logErrs(err)
	code := http.StatusBadRequest
	message := ""
	if errors.Is(err, ErrJson) {
		message = "mensaje ............."
	} else {
		message = ""
	}
	/// .....
	return c.JSON(code, message)
}
