package main

import (
	"html/template"
	"net/http"
)


func (app *app) home(w http.ResponseWriter, r *http.Request){

	// if this is the post request then update the count:
	//  if r.Method== http.MethodPost{
	// 	 app.count++;
	//  }

	  ts,err:=template.ParseFiles("./views/index.html")

		if err!=nil{
			app.logger.Error("error parsing template","err",err.Error())
			http.Error(w,"internal server error",http.StatusInternalServerError)
			return
		}

    // data:= struct{
		// 	Count int
		// }{
		// 	Count: app.count,
		// }

		// Check if the request is an HTMX request
	// if r.Header.Get("HX-Request") == "true" {
	// 	// Render only the "count" template for HTMX requests
	// 	err := ts.ExecuteTemplate(w, "count", data)
	// 	if err != nil {
	// 		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	// 	}
	// 	return
	// }


	// for the get request and if i dont click the count button
	// then this is showing

		err = ts.ExecuteTemplate(w,"index",nil)

		if err!=nil{
			app.logger.Error("error executing template","err",err.Error())
			http.Error(w,"internal server error",http.StatusInternalServerError)
			return
		}
}









func (app *app) contacts(w http.ResponseWriter, r *http.Request){
	
	 err:=r.ParseForm()

	 if err!=nil{
		 app.logger.Error("error parsing form","err",err.Error())
		 http.Error(w,"internal server error",http.StatusInternalServerError)
		 return
	 }

	 name:= r.PostForm.Get("name")
	 email:= r.PostForm.Get("email")


	 	  ts,err:=template.ParseFiles("./views/index.html")

		if err!=nil{
			app.logger.Error("error parsing template","err",err.Error())
			http.Error(w,"internal server error",http.StatusInternalServerError)
			return
		}
		// lock the mutex:
		app.mu.Lock()

		app.Contacts = append(app.Contacts,contact{
			Name: name,
			Email: email,
		})

		// unlock the mutex:
		app.mu.Unlock()

		data:= struct{
			Contacts []contact
		}{
			Contacts: app.Contacts,
		}


		ts.ExecuteTemplate(w,"contacts",data)


}