package main
import (
	"net/http"
	"io"
	"fmt"
)
func hello(res http.ResponseWriter, req*http.Request) {
	res.Header().Set("Content-Type", "text/html",)
	io.WriteString( res,
		`<doctype html>
			<html>
			<head>
			<title>Hello World</title>
			</head>
			<body>
			Hello World!
			</body>
		</html>`,)
}

func goodbye(res http.ResponseWriter, req*http.Request) {
	res.Header().Set("Content-Type", "text/html")
	fmt.Println(req)
	io.WriteString(res, "Goodbye")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.ListenAndServe(":9000", nil)
}
