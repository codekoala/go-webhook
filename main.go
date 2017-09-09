package main

import (
	"fmt"
	"net/http"

	"github.com/json-iterator/go"
	"github.com/labstack/echo"

	"github.com/codekoala/go-webhook/gitlab"
)

func main() {
	e := echo.New()

	e.POST("/gitlab", func(c echo.Context) (err error) {
		var (
			kind string
			hook interface{}
		)

		r := c.Request()
		defer r.Body.Close()

		kind, hook, err = gitlab.ParseReader(r.Body)
		if err != nil {
			return
		}

		out, _ := jsoniter.MarshalIndent(hook, "", "  ")
		fmt.Printf("%s: %s\n", kind, out)

		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
