package environment

import (
   "net/http"
   "github.com/labstack/echo"
)

// TODO handle errors

func PostMessageAndReturnHash(c echo.Context) error {
	msg := new(PostMessage)
	c.Bind(msg)
	var hashString string
	if HashAndStoreMessage(msg.Message, &hashString) {
		response := &PostMessageResponse{Digest:hashString}
		return c.JSON(http.StatusOK, response)
	}

    return echo.NewHTTPError(500, CATCH_ALL_ERROR)
}

func GetMessageFromHash(c echo.Context) error{
	hash := c.Param("hash")
	var message string
	if LookupHash(hash, &message) {
		response := &GetMessageResponse{Message:message}
		return c.JSON(http.StatusOK, response)
	}

	return echo.NewHTTPError(404, MSG_NOT_FOUND)
}