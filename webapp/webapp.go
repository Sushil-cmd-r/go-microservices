package webapp

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Webapp struct {
	routers []*Router
}

func NewApp() *Webapp {
	return &Webapp{}
}

func (w *Webapp) Use(path string, router *Router) {
	router.basePath = path
	regexPath := path + `(\/?$|\/[a-zA-Z0-9]+)`
	router.pathRegex = regexp.MustCompile(regexPath)

	w.routers = append(w.routers, router)
}

func (w *Webapp) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	url := TrimSuffix(request.URL.Path, "/")
	routers := w.routers
	for _, router := range routers {
		group := router.pathRegex.FindStringSubmatch(url)
		if len(group) > 0 {
			router.handle(writer, request, request.Method)
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprintf(writer, "404 Not Found\n")
	if err != nil {
		log.Fatal("error: Unable to write to client: ", err)
	}
}

func (w *Webapp) Listen(addr string) {
	err := http.ListenAndServe(addr, w)
	if err != nil {
		log.Fatal("error: Unable to Listen on given Port: ", err)
	}
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
