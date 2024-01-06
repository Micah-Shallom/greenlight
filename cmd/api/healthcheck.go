package main

import (
	"net/http"
)

//handler writes a plain-text response with the information about the application status, os and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)

	// js := `{"status": "available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)

	data := map[string]string{
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}

	//marshal data into an encoded json object
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}