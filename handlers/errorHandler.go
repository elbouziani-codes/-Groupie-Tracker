package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

type ErrorData struct {
	Error  string
	Status int
}

func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	errorData := ErrorData{
		Error:  message,
		Status: statusCode,
	}

	var buff bytes.Buffer
	err = tmpl.Execute(&buff, errorData)
	if err != nil {
		http.Error(w, message, statusCode)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8") 
	w.WriteHeader(statusCode)                             
	
	w.Write(buff.Bytes())
}