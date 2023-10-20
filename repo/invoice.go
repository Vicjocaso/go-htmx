package repo

// func GetAll(c echo.Context) error {
// 	tmpl := template.Must(template.ParseFiles("./views/invoice.html"))
// 	invoices := map[string][]models.Invoice{
// 		"Invoices": {
// 			{Id: "1", Client: "Victor", Price: "100", Status: "Pending", CreateAt: time.Now().Format("2006-01-02")},
// 			{Id: "2", Client: "Jose", Price: "200", Status: "Pending", CreateAt: time.Now().Format("2006-01-02")},
// 			{Id: "3", Client: "Karla", Price: "300", Status: "Paid", CreateAt: time.Now().Format("2006-01-02")},
// 		},
// 	}
// 	tmpl.Execute(c.Response().Writer, invoices)
// 	c.Redirect(200, "../views/invoice.html")
// 	return c.Redirect(200, "../views/invoice.html")
// }

// func Add(w http.ResponseWriter, r *http.Request) {
// 	client := r.PostFormValue("client")
// 	price := r.PostFormValue("price")
// 	status := r.PostFormValue("status")
// 	fmt.Println(client, "Client")
// 	fmt.Println(price, "Price")
// 	fmt.Println(status, "Status")
// 	tmpl := template.Must(template.ParseFiles("./views/invoice.html"))
// 	tmpl.ExecuteTemplate(w, "invoice-list-element", models.Invoice{Id: "4", Client: client, Price: price, Status: status, CreateAt: time.Now().Format("2006-01-02")})
// }
