package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var (
	baseURL = "http://localhost:1323"
	method1 = "GET"
)

func HomeHandler(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	var datax, err = getData()
	if err != nil {
		fmt.Println(err.Error())
	}

	popData, errpop := getPopData()
	if errpop != nil {
		fmt.Println(errpop.Error())
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":    "HOME",
		"msg":     "Framework Echo Worked!",
		"data":    datax,
		"popData": popData,
	})
}

func getData() ([]menu, error) {
	client := &http.Client{}
	var data []menu
	var err error

	req, err := http.NewRequest(method1, baseURL+"/read_menu", nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // dont forget to close the response portal

	// to decode data from byte to json
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, err

}

func getPopData() ([]menu, error) {
	client := &http.Client{}
	var data []menu
	var err error

	req, err := http.NewRequest(method1, baseURL+"/popular_menu", nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // dont forget to close the response portal

	// to decode data from byte to json
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, err

}
