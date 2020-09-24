package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	negronilogrus "github.com/meatballhat/negroni-logrus"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {
	routeTest3()
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

//加入negroni來紀錄系統
func routeTest3() {
	host := "127.0.0.1"
	port := "8080"
	output := ""

	args := os.Args[1:]

	for {
		if len(args) < 2 {
			break
		} else if args[0] == "-h" || args[0] == "--host" {
			host = args[1]
			args = args[2:]
		} else if args[0] == "-p" || args[0] == "--port" {
			port = args[1]
			args = args[2:]
		} else if args[0] == "-l" || args[0] == "--log" {
			output = args[1]
			args = args[2:]
		} else {
			log.Fatalln(fmt.Sprintf("Unknown parameter: %s", args[0]))
		}
	}

	mux := httprouter.New()
	//加載靜態資源 *filepath 的部分會由實際的資料夾和檔案名稱來取代
	// Listen to CSS assets
	mux.ServeFiles("/css/*filepath", http.Dir("public/css"))
	// Listen to JavaScript assets
	mux.ServeFiles("/js/*filepath", http.Dir("public/js"))

	mux.GET("/", index)
	// Custom 404 page
	mux.NotFound = http.HandlerFunc(notFound)
	// Custom 500 page
	mux.PanicHandler = errorHandler

	//create a new logger
	log := log.New()

	var f *os.File
	var err error

	if output != "" {
		f, err = os.Create(output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		log.SetOutput(f)
	}

	//set the logger of the server
	negroni := negroni.New()
	negroni.Use(negronilogrus.NewMiddlewareFromLogger(log, "web"))
	negroni.UseHandler(mux)

	//set the parameters for a http server
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: negroni,
	}

	//run
	log.Println(fmt.Sprintf("Run the web server at %s:%s", host, port))
	log.Fatal(server.ListenAndServe())
}

/*handler*/

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}

func handler2(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	fmt.Fprintf(writer, "Hello world")
}

func index(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	//加載模板
	//tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl := template.Must(template.ParseFiles("views/layout.html", "views/index.html", "views/head.html"))

	//資料
	data := struct {
		Title string
		Names []string
	}{
		"My Awesome Site",
		[]string{"Tony", "James", "Amy"},
	}

	//啟動
	err := tmpl.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		http.Error(writer, err.Error(), 500)
	}

}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Page Not Found")
}

func errorHandler(w http.ResponseWriter, r *http.Request, p interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal Server Error")
}
