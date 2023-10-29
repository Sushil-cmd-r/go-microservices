package webapp

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Param struct {
	key string
	val string
}

type Route struct {
	method    string
	pathRegex *regexp.Regexp
	handler   http.Handler
	params    []Param
}

type Router struct {
	basePath  string
	pathRegex *regexp.Regexp
	routes    []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) handle(writer http.ResponseWriter, request *http.Request, httpMethod string) {
	url := request.URL.Path
	url = strings.TrimPrefix(url, r.basePath)
	for _, route := range r.routes {
		group := route.pathRegex.FindStringSubmatch(url)
		fmt.Println(group)
		if len(group) != len(route.params)+1 {
			log.Fatal("error: Something went wrong")
		}

		group = group[1:]

		for i := 0; i < len(group); i++ {
			route.params[i].val = group[i]
		}
	}

	writer.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(writer, "200 OK\n")
	if err != nil {
		log.Fatal("error: Unable to write to client: ", err)
	}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	route := Route{}
	if strings.HasPrefix(path, "/") == false {
		log.Fatal("error: Path should start with /")
	}
	regx := regexp.MustCompile(`:([^/]+)`)
	groups := regx.FindAllStringSubmatch(path, -1)

	for _, group := range groups {
		p := Param{key: group[1], val: ""}
		route.params = append(route.params, p)
	}

	str := regx.ReplaceAllString(path, "([a-zA-Z0-9]+)")
	route.pathRegex = regexp.MustCompile(str)
	route.method = http.MethodGet
	route.handler = handler

	r.routes = append(r.routes, route)
}
