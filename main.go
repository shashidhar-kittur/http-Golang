package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)
func main(){
    fmt.Println("hello")
    res,err:=http.Get("https://jsonplaceholder.typicode.com/posts/1") //res contains an *http.Response struct (pointer to an http.Response). The body is a stream (io.ReadCloser), not already read into memory
    if err != nil {
        log.Fatalln(err)
    }
    body, err:= io.ReadAll(res.Body) //res.Body returns bytes
    if err != nil {
        log.Fatalln(err)
    }
    res.Body.Close() //prevents mem leaks
    fmt.Println(string(body))

    

}