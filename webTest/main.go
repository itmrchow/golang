package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	routeTest2()
}

/*路由*/

func webHelloWorld() {
	//於根目錄 / 登記handler處理器
	//當網站發出request時，app會使用handler所作的內容來回應response
	http.HandleFunc("/", handler)

	//啟動網頁伺服器程式
	//handler 則傳入 ServeMux，ServeMux 是 HTTP request multiplexer
	//mux 統籌處理該網頁程式有登記的路徑及其處理器 (handler)
	//將傳入的路徑指向相對應的處理器，也就是整個網頁程式的路由 (router)
	//在本例中我們傳入 nil，代表使用內建的 DefaultServeMux
	http.ListenAndServe(":8080", nil)
}

func routeTest() {
	//mux獨立出來
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	//server物件
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}

//動態路由
func routeTest2() {
	mux := httprouter.New()
	mux.GET("/hello/:name", handler2)
	// Custom 404 page
	mux.NotFound = http.HandlerFunc(notFound)
	// Custom 500 page
	mux.PanicHandler = errorHandler

	//server物件
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	//Run the server
	server.ListenAndServe()
}

/*handler*/

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}

func handler2(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	fmt.Fprintf(writer, "Hello world")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Page Not Found")
}

func errorHandler(w http.ResponseWriter, r *http.Request, p interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal Server Error")
}
