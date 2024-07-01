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



// . appHandler
func appHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,

  pageTitle := "www.Phoenix-Valley.com - Explore - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
  
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  pageFilePath := template.Must(template.ParseFiles("main_layout.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler

  
  
  
  
  
}