package server

import (
	HomeHandler "TechVeritasMultiplex/controllers/home"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//NewServer instance
func NewServer() Server {
	return Server{}
}

//Server struct
type Server struct {
	port string
}

// Init all vals
func (s *Server) Init(port string) {
	log.Println("Initializing server...")
	s.port = ":" + port

}

//NewRouter function
func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)

	return
}

//Router struct
type Router struct {
	Router *mux.Router
}

//Init to initialize router
func (r *Router) Init() {
	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

}

//GetRoutes global
func GetRoutes() Routes {
	return Routes{
		Route{"Home", "GET", "/api/getShows", HomeHandler.GetAllShowDetails()},
		Route{"Home", "GET", "/api/bookShow/{showid}/{seats}", HomeHandler.BookShowSeats()},
		Route{"Home", "GET", "/api/getMovieList", HomeHandler.GetMovieList()},
		Route{"Home", "GET", "/api/getMultiplex", HomeHandler.GetMultiplex()},
	}
}

//Routes struct
type Routes []Route

//Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Start the server
func (s *Server) Start() {

	r := NewRouter()

	r.Init()

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Serving on localhost:8000")
	log.Fatal(newServer.ListenAndServe())

}
