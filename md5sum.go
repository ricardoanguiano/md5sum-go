package main

import (
       "fmt"
       "io"
       "os"
       "log"
       "crypto/md5"
)

func main() {
     // Argument handling
     args := os.Args;
     if len(args) != 2 {
     	log.Printf("Error: wrong number of arguments")
     	log.Fatal("Usage: md5sum filename")
     }
     f, err := os.Open(args[1])
     if err != nil {
     	log.Fatal(err)
     }
     h := md5.New()
     io.Copy(h, f)
     fmt.Printf("%x  %s\n", h.Sum(nil), args[1])
}
