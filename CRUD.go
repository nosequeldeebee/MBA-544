package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type Province struct {
	Name   string
	Region string
}

var (
	provinces = make(map[string]Province)
	mu        sync.Mutex
)

var tmpl = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
		<title>Canadian Provinces</title>
</head>
<body>
		<h1>Canadian Provinces</h1>
		<form method="POST" action="/add">
				<input type="text" name="name" placeholder="Province name" required>
				<input type="submit" value="Add Province">
		</form>
		<ul>
		{{range $name, $province := .}}
				<li>
						{{$name}} - {{$province.Region}}
						<form method="POST" action="/update" style="display:inline;">
								<input type="hidden" name="name" value="{{$name}}">
								<select name="region">
										<option value="Western" {{if eq $province.Region "Western"}}selected{{end}}>Western</option>
										<option value="Eastern" {{if eq $province.Region "Eastern"}}selected{{end}}>Eastern</option>
								</select>
								<input type="submit" value="Update">
						</form>
						<form method="POST" action="/delete" style="display:inline;">
								<input type="hidden" name="name" value="{{$name}}">
								<input type="submit" value="Delete">
						</form>
				</li>
		{{end}}
		</ul>
</body>
</html>
`))

func main() {
	http.HandleFunc("/", listProvinces)
	http.HandleFunc("/add", addProvince)
	http.HandleFunc("/update", updateProvince)
	http.HandleFunc("/delete", deleteProvince)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func listProvinces(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, provinces)
}

func addProvince(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	region := "Not an Eastern or Western province!"
	if name == "British Columbia" || name == "Alberta" || name == "Saskatchewan" || name == "Manitoba" {
		region = "Western"
	}
	if name == "Ontario" || name == "Quebec" || name == "Prince Edward Island" || name == "Newfoundland" || name == "New Brunswick" || name == "Nova Scotia"  {
		region = "Eastern"
	}

	mu.Lock()
	provinces[name] = Province{Name: name, Region: region}
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateProvince(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	region := r.FormValue("region")

	mu.Lock()
	if p, exists := provinces[name]; exists {
		p.Region = region
		provinces[name] = p
	}
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteProvince(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")

	mu.Lock()
	delete(provinces, name)
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
