package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/signup.html")
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		//generate totp secret
		key, _ := totp.Generate(totp.GenerateOpts{
			Issuer:      "MyApp",
			AccountName: username,
		})

		user := User{
			Username: username,
			Password: password,
			Secret:   key.Secret(),
		}

		// load user
		var users []User
		file, _ := os.ReadFile("users.json")
		json.Unmarshal(file, &users)

		//append
		users = append(users, user)
		data, _ := json.MarshalIndent(users, "", " ")
		os.WriteFile("users.json", data, 0644)

		qrCodeFile := "static/qrcode.png"
		_ = qrcode.WriteFile(key.URL(), qrcode.Medium, 256, qrCodeFile)

		http.Redirect(w, r, "/qrcode", http.StatusSeeOther)
	}
}

func qrCodeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/qr_code.html")
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/login.html")
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		//load users from json
		file, _ := os.ReadFile("users.json")
		var users []User
		json.Unmarshal(file, &users)

		for _, user := range users {
			if user.Username == username && user.Password == password {
				tmpl, _ := template.ParseFiles("templates/otp.html")
				tmpl.Execute(w, user)
				return
			}
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func otpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		otp := r.FormValue("otp")

		//load users from json
		file, _ := os.ReadFile("users.json")
		var users []User
		json.Unmarshal(file, &users)

		for _, user := range users {
			if user.Username == username && totp.Validate(otp, user.Secret) {
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			}
		}

		http.Redirect(w, r, "/otp-error", http.StatusSeeOther)
	}
}

func otpErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/otp_error.html")
	tmpl.Execute(w, nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/dashboard.html")
	tmpl.Execute(w, nil)
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/qrcode", qrCodeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/otp", otpHandler)
	http.HandleFunc("/otp-error", otpErrorHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	fmt.Println("server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
