// . . www.phoenix-valley.com - Explore - Website App

package main

// .

import (
  "os"
  "log"
		
  "text/template"
  "net/http"
		
  "cloud.google.com/go"
		
  "cloud.google.com/go/firestore"
  "google.golang.org/api/iterator"
)



type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}


