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
<html lang="en">
<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Canadian Provinces</title>
		<style>
				body {
						font-family: Arial, sans-serif;
						line-height: 1.6;
						color: #333;
						max-width: 800px;
						margin: 0 auto;
						padding: 20px;
						background-color: #f4f4f4;
				}
				h1 {
						color: #2c3e50;
						text-align: center;
				}
				form {
						background-color: #fff;
						padding: 20px;
						border-radius: 5px;
						box-shadow: 0 2px 5px rgba(0,0,0,0.1);
						margin-bottom: 20px;
				}
				input[type="text"], select {
						width: 100%;
						padding: 10px;
						margin-bottom: 10px;
						border: 1px solid #ddd;
						border-radius: 4px;
				}
				input[type="submit"] {
						background-color: #3498db;
						color: #fff;
						border: none;
						padding: 10px 20px;
						cursor: pointer;
						border-radius: 4px;
				}
				input[type="submit"]:hover {
						background-color: #2980b9;
				}
				ul {
						list-style-type: none;
						padding: 0;
				}
				li {
						background-color: #fff;
						margin-bottom: 10px;
						padding: 15px;
						border-radius: 5px;
						box-shadow: 0 2px 5px rgba(0,0,0,0.1);
				}
				.province-name {
						font-weight: bold;
						color: #2c3e50;
				}
				.region {
						color: #7f8c8d;
				}
				.update-form, .delete-form {
						display: inline-block;
						margin-left: 10px;
				}
				img {
						max-width: 100%;
						height: auto;
						display: block;
						margin: 20px auto;
						border-radius: 5px;
						box-shadow: 0 2px 5px rgba(0,0,0,0.1);
				}
		</style>
</head>
<body>
		<h1>Canadian Provinces</h1>
		<img src="../images/map.svg" alt="Map of Canada">
		<form method="POST" action="/add">
				<input type="text" name="name" placeholder="Province name" required>
				<input type="submit" value="Add Province">
		</form>
		<ul>
		{{range $name, $province := .}}
				<li>
						<span class="province-name">{{$name}}</span> - <span class="region">{{$province.Region}}</span>
						<form class="update-form" method="POST" action="/update">
								<input type="hidden" name="name" value="{{$name}}">
								<select name="region">
										<option value="Western" {{if eq $province.Region "Western"}}selected{{end}}>Western</option>
										<option value="Eastern" {{if eq $province.Region "Eastern"}}selected{{end}}>Eastern</option>
								</select>
								<input type="submit" value="Update">
						</form>
						<form class="delete-form" method="POST" action="/delete">
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
	// create a directory called images then put the map.svg image in it
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

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
	region := "Eastern"
	if name == "British Columbia" || name == "Alberta" || name == "Saskatchewan" || name == "Manitoba" {
		region = "Western"
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
