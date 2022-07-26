package handlers

import (
	"net/http"
)


func Home(w http.ResponseWriter, r *http.Request) {
    RenderTemplate(w,"home.page.tmpl")
}

func AboutUs(w http.ResponseWriter, r *http.Request){
    RenderTemplate(w,"about.page.tmpl")
    
}

