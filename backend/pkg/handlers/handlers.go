package handlers

import (
	"net/http"
    "project1/pkg/render"
)


func Home(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w,"home.page.tmpl")
}

func AboutUs(w http.ResponseWriter, r *http.Request){
    render.RenderTemplate(w,"about.page.tmpl")
    
}

