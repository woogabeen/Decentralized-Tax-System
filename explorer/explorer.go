package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/WoodoCoin/blockchain"
	"github.com/gorilla/mux"
)

var templates *template.Template

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := HomeData{"Home", nil}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		blockchain.Blockchain().AddBlock()
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	router := mux.NewRouter()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	router.HandleFunc("/", home)
	router.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
