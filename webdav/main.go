package main

import (
    "flag"
    "os"
    "fmt"
    "log"
    "net/http"
    "golang.org/x/net/webdav"
)

func webdavLogger(r *http.Request, err error) {
    if err != nil {
        log.Printf("WEBDAV [%s]: %s, ERROR: %s\n", r.Method, r.URL, err)
    } else {
        log.Printf("WEBDAV [%s]: %s \n", r.Method, r.URL)
    }
}

func main() {
    var (
        addr = flag.String("addr", ":3001", "address to listen")
        root = flag.String("root", "data", "root folder")
    )
    flag.Parse()

    if _, err := os.Stat(*root); os.IsNotExist(err) {
        os.MkdirAll(*root, 0755)
    }

    fmt.Println("addr=", *addr, ", root=", *root)

    http.ListenAndServe(*addr, &webdav.Handler {
        FileSystem: webdav.Dir(*root),
        LockSystem: webdav.NewMemLS(),
        Logger: webdavLogger,
    })
}