/*************************************************************************
> File Name:     client.go
> Author:        程序员Carl
> 微信公众号:    代码随想录
> Created Time:  Wed May 16 16:51:49 2018
> Description:   
 ************************************************************************/

package main

import (
    "fmt"
    "net/http"
    "os"
//    "log"
    "io/ioutil"
)
 func main(){
    upload()
}


func upload() {

    file, err := os.Open("./test.txt")
    if err != nil {

        panic(err)
    }
    defer file.Close()

    res, err := http.Post("http://127.0.0.1:5050/upload", "binary/octet-stream", file)
    if err != nil {

        panic(err)
    }
    defer res.Body.Close()
    message, _ := ioutil.ReadAll(res.Body)
    fmt.Printf(string(message))
}
