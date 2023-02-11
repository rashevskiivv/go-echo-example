package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/rashevsiivv/echo-example/vo"
	"io"
	"log"
	"net/http"
)

// GetCats http://localhost:8000/cats/json?name=arnold&type=fluffy
func GetCats(c echo.Context) error {
	catName, catType := c.QueryParam("name"), c.QueryParam("type")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your cat name is : %s\nand cat type is : %s\n", catName, catType))
	} else if dataType == "json" {
		cat := vo.Cat{
			Name: catName,
			Type: catType,
		}
		return c.JSON(http.StatusOK, cat)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as Sting or JSON",
		})
	}

}

func AddCat(c echo.Context) error {
	cat := vo.Cat{}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(c.Request().Body)

	err := json.NewDecoder(c.Request().Body).Decode(&cat)
	if err != nil {
		defer log.Fatal("Failed reading the request body", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is yout cat %#v", cat)
	return c.String(http.StatusOK, "We got your Cat!!!")
}
