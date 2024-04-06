package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"time"
)

type (
	middleware func(http.Handler) http.Handler
	router     struct {
		*http.ServeMux
		chain []middleware
	}
)

func NewRouter(mx ...middleware) *router {
	return &router{ServeMux: &http.ServeMux{}, chain: mx}
}

func (r *router) Use(mx ...middleware) {
	r.chain = append(r.chain, mx...)
}

func (r *router) Group(fn func(r *router)) {
	fn(&router{ServeMux: r.ServeMux, chain: slices.Clone(r.chain)})
}

func (r *router) Get(path string, fn http.HandlerFunc, mx ...middleware) {
	r.handle(http.MethodGet, path, fn, mx)
}

func (r *router) Post(path string, fn http.HandlerFunc, mx ...middleware) {
	r.handle(http.MethodPost, path, fn, mx)
}

func (r *router) Put(path string, fn http.HandlerFunc, mx ...middleware) {
	r.handle(http.MethodPut, path, fn, mx)
}

func (r *router) Delete(path string, fn http.HandlerFunc, mx ...middleware) {
	r.handle(http.MethodDelete, path, fn, mx)
}

func (r *router) Head(path string, fn http.HandlerFunc, mx ...middleware) {
	r.handle(http.MethodHead, path, fn, mx)
}

func (r *router) Options(path string, fn http.HandlerFunc, mx ...middleware) {
	r.handle(http.MethodOptions, path, fn, mx)
}

func (r *router) handle(method, path string, fn http.HandlerFunc, mx []middleware) {
	r.Handle(method+" "+path, r.wrap(fn, mx))
}

func (r *router) wrap(fn http.HandlerFunc, mx []middleware) (out http.Handler) {
	out, mx = http.Handler(fn), append(slices.Clone(r.chain), mx...)

	slices.Reverse(mx)
	for _, m := range mx {
		out = m(out)
	}
	return
}

func NumberShow(i int) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("middleware número ", i, " começou")
			next.ServeHTTP(w, r)
			fmt.Println("midleware número ", i, " acabou")
		})
	}
}

func Logger() middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			next.ServeHTTP(w, r)
			elapsedTime := time.Since(startTime)

			slog.Info("http request", slog.String("método", r.Method), slog.String("path", r.URL.Path), slog.String("duração", elapsedTime.String()))
		})
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[o handler executou aqui]")
	fmt.Fprintln(w, "Hello world do handler ", r.URL.Path)
}

func main() {
	r := NewRouter(Logger())

	r.Group(func(r *router) {
		r.Use(NumberShow(1), NumberShow(2))

		r.Get("/foo", helloHandler)
	})

	r.Group(func(r *router) {
		r.Use(NumberShow(3))

		r.Get("/bar", helloHandler, NumberShow(4))
		r.Get("/baz", helloHandler, NumberShow(5))
	})

	r.Post("/foobar", helloHandler)

	slog.Info("serviço iniciado na porta :8090")
	http.ListenAndServe(":8090", r)
}
