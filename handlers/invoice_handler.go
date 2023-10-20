package handlers

import (
	"fmt"
	database "go-htmx/db"
	"go-htmx/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hanlderinvoice(c echo.Context) error {
	invoices := []models.Invoice{{
		Id:     "1",
		Client: "Victor",
		Price:  100,
		Status: "Paid",
	}, {
		Id:     "2",
		Client: "Jose",
		Price:  200,
		Status: "Pending",
	}, {
		Id:     "3",
		Client: "Karla",
		Price:  100,
		Status: "Paid",
	}}
	fmt.Println("GG")

	return c.Render(http.StatusOK, "index.html", invoices)
}

func HanlderAddInvoice(c echo.Context) error {
	fmt.Println("client")

	if database.Db == nil {
		return c.String(http.StatusInternalServerError, "database")
	}
	id := 90
	client := c.FormValue("client")
	price := c.FormValue("price")
	status := c.FormValue("status")
	fmt.Println(client)
	fmt.Println(price)
	fmt.Println(status)

	_, err := database.Db.Exec("INSERT INTO invoice (Client, price, status) VALUES (?, ?, ?,?)", id, client, price, status)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(200, "../views/invoice.html")
}

func HandlerInvoicesByClient(c echo.Context) error {
	invoices := []models.Invoice{}
	if database.Db == nil {
		return c.String(http.StatusInternalServerError, "database")
	}
	client := c.FormValue("client")

	res := database.Db.QueryRow("SELECT * FROM invoice WHERE client = ?", client)

	err := res.Scan(&invoices)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	fmt.Println(invoices, "invoices")

	return c.Redirect(200, "../views/invoice.html")
}

func HandlerInvoiceById(c echo.Context) error {
	invoices := []models.Invoice{}
	fmt.Println(database.Db)
	if database.Db == nil {
		return c.String(http.StatusInternalServerError, "database")
	}
	id := c.Param("id")
	fmt.Println(id, "id")

	res, err := database.Db.Query("SELECT * FROM invoice WHERE id = ?", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	for res.Next() {
		invoice := models.Invoice{}
		err := res.Scan(&invoice.Id, &invoice.Client, &invoice.Price, &invoice.Status)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		invoices = append(invoices, invoice)
	}

	fmt.Println(invoices, "invoices")
	// return c.Redirect(200,"/invoice")
	return c.Render(http.StatusOK, "invoice.html", invoices)
}

func HandlerDeleteInvoice(c echo.Context) error {
	if database.Db == nil {
		return c.String(http.StatusInternalServerError, "database")
	}
	id := c.Param("id")

	res, err := database.Db.Exec("DELETE FROM invoice WHERE id = ?", id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	rowsAff, err := res.RowsAffected()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if rowsAff > 0 {
		return c.String(http.StatusOK, "Invoice deleted")
	}

	return c.Render(200, "invoice.html", nil)
}

func HanlderAddinvoice(c echo.Context) error {
	if database.Db == nil {
		return c.String(http.StatusInternalServerError, "database")
	}

	return c.Render(200, "create-invoice.html", nil)
}
