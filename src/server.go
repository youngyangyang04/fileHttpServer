/*************************************************************************
> File Name:     main.go
> Author:        sunxiuyang
> Mail:          sunxiuyang@baidu.com
> Created Time:  Wed May 16 14:30:07 2018
> Description:
 ************************************************************************/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
    "strings"
    "strconv"
)
const PORT string = "8085"
const IPADDR string = "106.12.97.46"
const DOWNLOADURL string  = "wget http://"+ IPADDR +":"+ PORT + "/"
const UPLOADFILEDIR string = "./upload/"
const FILENAMEERR string = "Get file name error !"
const COVERVALUE string = "noCover"
const COVERKEY string = "cover"
const PARAMNUMCHECK int = 3

//check the file name exist or not
func checkFileName(fileName string) (bool,string){
    bExist, err := exists(UPLOADFILEDIR + fileName)
    if err != nil{
        log.Fatal(err)
    }
    if bExist == true{
        i := 1
        version := strconv.Itoa(i)
        for true{
            bfExist, _ := exists(UPLOADFILEDIR + fileName + "_" + version)
            if bfExist == false{
                fileName += "_"+version 
                break
            }
            i++
            version = strconv.Itoa(i)
        }
    }
    return bExist, fileName
}
func postMethod(w http.ResponseWriter, r *http.Request){
    //The whole request body is parsed and up to a total of maxMemory bytes of its file parts are stored in memory
    err := r.ParseMultipartForm(100000)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    r.ParseForm()
    sCover := r.Form.Get(COVERKEY)
    fmt.Printf(COVERKEY+":%s\n", sCover)
    //get a ref to the parsed multipart form
    m := r.MultipartForm
    files := m.File["uploadfile"]
    for i, _ := range files {
        bNewFileName := false
        newFileName :=files[i].Filename;
        if sCover == COVERVALUE{
            bNewFileName,newFileName = checkFileName(files[i].Filename)
        }
        fmt.Printf("fileNname[%d]:"+files[i].Filename+", newFileName:" + newFileName +"\n", i)
        
        file, err := files[i].Open()
        defer file.Close()
        targetFile, err := os.Create(UPLOADFILEDIR + newFileName)
        defer targetFile.Close()
        if err != nil {
            panic(err)
        }
        n, err := io.Copy(targetFile, file)
        if err != nil {
            panic(err)
        }
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        
        if bNewFileName == true{
            w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n%s already exists, new file name:%s\nGet object way: " + DOWNLOADURL + "%s\n", 
            n,files[i].Filename,newFileName, newFileName)))
        }else{
            w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\nGet object way: " + DOWNLOADURL + "%s\n", 
            n,newFileName)))
        }
    }
}
func putMethod(w http.ResponseWriter, r *http.Request){
        fileName, fileNameErr := getFileNameFromURL(r.URL.Path)
        if fileNameErr !=""{
            log.Fatal(fileNameErr)
        }
        fmt.Printf("getHandle URL:%s, filename:%s\n", r.URL.Path,fileName)

        targetFile, err := os.Create(UPLOADFILEDIR + fileName)
        defer targetFile.Close()
        if err != nil {
            panic(err)
        }
        file := r.Body
        n, err := io.Copy(targetFile, file)
        if err != nil {
            panic(err)
        }
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\nGet object way: "+ DOWNLOADURL +"%s\n", n, fileName)))
}
// upload object
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("method:" + r.Method + "\n")
    if r.Method == "PUT"{
        putMethod(w, r)
    }
    if r.Method == "POST" {
        postMethod(w, r)
    };
}
// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
//get object
func getHandle(w http.ResponseWriter, r *http.Request) {
    fileName, err := getFileNameFromURL(r.URL.Path)
    if err != ""{
        log.Fatal(err)
    }
    fmt.Printf("getHandle URL:%s, filename:%s\n", r.URL.Path,fileName)
    objectPath := UPLOADFILEDIR + fileName
    http.ServeFile(w, r, objectPath)
}
//get filename from url
func getFileNameFromURL(URL string) (string, string){
    params := strings.Split(URL,"/")
    fileName := params[len(params)-1]
    fmt.Printf("getFileNameFromURL, len:%d slice=%v\n",len(params),params)
    var err string
    if len(params) >PARAMNUMCHECK{
        err = FILENAMEERR
    }
    return fileName,err
}
func main() {
    http.HandleFunc("/", getHandle)
    http.HandleFunc("/upload/", uploadHandler)
    err := http.ListenAndServe(":" + PORT, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err.Error())
    }
}
