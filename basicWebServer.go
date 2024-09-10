package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a template
		tmpl, err := template.New("hello").Parse(`
<!DOCTYPE html>
<html>
<head>
		<title>{{.Title}}</title>
</head>
<body>
		<h1>{{.Message}}</h1>
</body>
</html>
`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Prepare data for the template
		data := PageData{
			Title:   "Hello Page",
			Message: "Hello, World!",
		}

		// Execute the template
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the server
	println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
