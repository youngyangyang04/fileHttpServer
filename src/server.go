/*************************************************************************
> File Name:     main.go
> Author:        sunxiuyang
> Mail:          sunxiuyang@baidu.com 
> Created Time:  Wed May 16 14:30:07 2018
> Description:   
 ************************************************************************/
package main
import (
    "io"
    "log"
    "net/http"
    "os"
    "fmt"
)
func helloWorldHandler(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "Hello world")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

    file, err := os.Create("../upload/result")
    if err != nil {

        panic(err)
    }
    n, err := io.Copy(file, r.Body)
    if err != nil {

        panic(err)
    }

    w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}
func main(){
//    http.HandleFunc("/", helloWorldHandler)
    http.HandleFunc("/upload", uploadHandler)
    err := http.ListenAndServe(":5050", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err.Error())
    }
}
