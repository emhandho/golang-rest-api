package handler

import (
	"echo/server"
	"fmt"
	_ "mysql-master"
	"net/http"

	"github.com/labstack/echo"
)

type menu struct {
	Id_menu     string
	Name        string
	Description string
	Image_url   string
	Type        string
	Cost        string
	Total_order string
}

var data []menu

func ReadData(c echo.Context) error {
	readMenu()

	return c.JSON(http.StatusOK, data)
}

func ReadPopularMenu(c echo.Context) error {
	popularMenu()

	return c.JSON(http.StatusOK, data)
}

func readMenu() {
	data = nil
	db, err := server.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	items, err := db.Query("SELECT * from menu_table")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for items.Next() {
		item := menu{}
		err := items.Scan(&item.Id_menu, &item.Name, &item.Description, &item.Image_url, &item.Type, &item.Cost)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, item)
	}
	if err = items.Err(); err != nil {
		panic(err.Error())
	}
}

func popularMenu() {
	data = nil
	db, err := server.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	items, err := db.Query("SELECT * from vw_total_order ORDER BY total_order DESC LIMIT 4")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for items.Next() {
		item := menu{}
		err := items.Scan(&item.Id_menu, &item.Name, &item.Description, &item.Image_url, &item.Type, &item.Cost, &item.Total_order)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		data = append(data, item)
	}
	if err = items.Err(); err != nil {
		panic(err.Error())
	}
}

func InputOrder(c echo.Context) error {
	db, err := server.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	id := c.FormValue("id_order")
	name := c.FormValue("order_name")
	phone := c.FormValue("phone_number")
	address := c.FormValue("address")
	total_order := c.FormValue("amount_order")

	_, err = db.Exec("INSERT into order_table values (?,?,?,?,?,?)", nil, id, name, phone, address, total_order)
	if err != nil {
		fmt.Println("Failed to Order!")
		return c.HTML(http.StatusOK, "<strong>Failed to Order the Food!</strong>")
	} else {
		fmt.Println("Successfuly Order!")
		return c.HTML(http.StatusOK, "<script>alert('Successfuly Order the Food, please wait for the food!'); window.location='http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func AddData(c echo.Context) error {
	db, err := server.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	name := c.FormValue("Name")
	description := c.FormValue("Description")
	image_url := c.FormValue("Image_url")
	fdtype := c.FormValue("Type")
	cost := c.FormValue("Cost")

	_, err = db.Exec("INSERT into menu_table values (?,?,?,?,?,?)", nil, name, description, image_url, fdtype, cost)
	if err != nil {
		fmt.Println("Failed to add Menu!")
		return c.JSON(http.StatusOK, "Failed to add Menu!")
	} else {
		fmt.Println("Successful adding Menu!")
		return c.JSON(http.StatusOK, "Successful adding Menu!")
	}
}

func UpdateData(c echo.Context) error {
	db, err := server.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	fdid := c.FormValue("id_menu")
	name := c.FormValue("Name")
	description := c.FormValue("Description")
	image_url := c.FormValue("Image_url")
	fdtype := c.FormValue("Type")
	cost := c.FormValue("Cost")

	_, err = db.Exec("UPDATE into menu_table set name = ?, description = ?, image_url = ?, food_type = ?, cost = ? where id_menu = ?)", name, description, image_url, fdtype, cost, fdid)
	if err != nil {
		fmt.Println("Failed to update Menu!")
		return c.JSON(http.StatusOK, "Failed to update Menu!")
	} else {
		fmt.Println("Successful update Menu!")
		return c.JSON(http.StatusOK, "Successful update Menu!")
	}
}

func DeleteData(c echo.Context) error {
	db, err := server.Connection()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()

	fdid := c.FormValue("id_menu")

	_, err = db.Exec("DELETE from menu_table where id_menu = ?)", fdid)
	if err != nil {
		fmt.Println("Failed to update Menu!")
		return c.JSON(http.StatusOK, "Failed to update Menu!")
	} else {
		fmt.Println("Successful update Menu!")
		return c.JSON(http.StatusOK, "Successful update Menu!")
	}
}
