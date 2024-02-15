package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// HandlerWithMethod stores an HTTP handler along with its method.
type HandlerWithMethod struct {
	Method  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string][]HandlerWithMethod // Map[path][]HandlerWithMethod
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string][]HandlerWithMethod),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers[path] = append(s.Handlers[path], HandlerWithMethod{Method: method, Handler: handler})
	//fmt.Println("Added", method, "handler for", path)
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handlers := range s.Handlers {
		for _, hwMethod := range handlers {
			//fmt.Println("Adding", hwMethod.Method, "handler for", path)
			switch hwMethod.Method {
			case http.MethodGet:
				s.Router.Get(path, hwMethod.Handler)
			case http.MethodPost:
				s.Router.Post(path, hwMethod.Handler)
				// Add cases for other HTTP methods as needed
			}
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
