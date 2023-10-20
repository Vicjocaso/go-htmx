package main

import (
	database "go-htmx/db"
	"go-htmx/handlers"
	"io"
	"log"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {

	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		panic("DB_URL is required")
	}

	err = database.InitDB(db_url)
	if err != nil {
		log.Fatalf("couild not initialize db: %+v", err)
	}

	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
	e.Renderer = renderer
	// template.NewTemplateRender(e, "/public/views/*.html")
	e.GET("/", handlers.Hanlderinvoice)
	e.GET("/invoice", handlers.HanlderAddInvoice)
	e.GET("/invoice/:id", handlers.HandlerInvoiceById)
	e.POST("/add-invoice", handlers.HanlderAddInvoice)

	// e.POST("/add-invoice", handlers.Hanlderinvoice)
	e.Logger.Fatal(e.Start(":3000"))
}
