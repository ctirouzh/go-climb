package http

import (
	"github.com/gorilla/mux"
	"github.com/pkritiotis/go-climb/internal/app"
	"log"
	"net/http"
)

//Server Represents the http server running for this service
type Server struct {
	app    app.App
	router *mux.Router
}

//NewServer HTTP Server constructor
func NewServer(app app.App) *Server {
	httpServer := &Server{app: app}
	httpServer.router = mux.NewRouter()
	httpServer.AddCragHTTPRoutes()
	http.Handle("/", httpServer.router)

	return &Server{app: app}
}

func (httpServer *Server) AddCragHTTPRoutes() {
	const cragsHTTPRoutePath = "/crags"
	//Queries
	httpServer.router.HandleFunc(cragsHTTPRoutePath, NewCragHandler(httpServer.app).GetAllCrags).Methods("GET")
	httpServer.router.HandleFunc(cragsHTTPRoutePath+"/{"+getCragIDURLParam+"}", NewCragHandler(httpServer.app).GetCrag).Methods("GET")

	//Commands
	httpServer.router.HandleFunc(cragsHTTPRoutePath, NewCragHandler(httpServer.app).AddCrag).Methods("POST")
	httpServer.router.HandleFunc(cragsHTTPRoutePath+"/{"+updateCragIDURLParam+"}", NewCragHandler(httpServer.app).UpdateCrag).Methods("PUT")
	httpServer.router.HandleFunc(cragsHTTPRoutePath+"/{"+deleteCragIDURLParam+"}", NewCragHandler(httpServer.app).DeleteCrag).Methods("DELETE")

}

//ListenAndServe Starts listening for requests
func (httpServer *Server) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, nil))
}
