package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strings"
)

type User struct {
	Name       string
	Occupation string
}

type Data struct {
	Users []User
}

var data = Data{
	Users: []User{
		{Name: "John Doe", Occupation: "gardener"},
		{Name: "Roger Roe", Occupation: "driver"},
		{Name: "Peter Smith", Occupation: "teacher"},
		{Name: "Jo Do", Occupation: "trader"},
		{Name: "eger Rew", Occupation: "programer"},
		{Name: "Olivia Smith", Occupation: "hairdresser"},
	},
}

var tmp = template.Must(template.ParseFiles("layout.html"))

func main() {

	http.HandleFunc("/get", getProcess)
	http.HandleFunc("/delete", deleteProcess)
	http.HandleFunc("/put", putProcess)
	http.HandleFunc("/patch", patchProcess)
	http.HandleFunc("/post", postProcess)
	http.HandleFunc("/ip", ipProcess)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}

func getProcess(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/get" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		tmp.Execute(w, data)
	default:
		fmt.Fprintf(w, "Method %v Not Allowed.\n", r.Method)
		http.Error(w, "Status Code 405.", http.StatusMethodNotAllowed)
	}
}

func deleteProcess(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/delete" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "DELETE":
		//Removing last User in User Table
		// data.Users = append(data.Users[:len(data.Users)-1])
		// tmp.Execute(w, data)
		fmt.Fprintf(w, "DELETE !!!")
	default:
		fmt.Fprintf(w, "Method %v Not Allowed.\n", r.Method)
		http.Error(w, "Status Code 405.", http.StatusMethodNotAllowed)
	}
}

func patchProcess(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/patch" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "PATCH":
		fmt.Fprintf(w, "PATCH !!!")
	default:
		fmt.Fprintf(w, "Method %v Not Allowed.\n", r.Method)
		http.Error(w, "Status Code 405.", http.StatusMethodNotAllowed)
	}
}

func putProcess(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/put" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "PUT":
		fmt.Fprintf(w, "PUT !!!")
	default:
		fmt.Fprintf(w, "Method %v Not Allowed.\n", r.Method)
		http.Error(w, "Status Code 405.", http.StatusMethodNotAllowed)
	}
}

func postProcess(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/post" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		fmt.Fprintf(w, "POST !!!")
	default:
		fmt.Fprintf(w, "Method %v Not Allowed.\n", r.Method)
		http.Error(w, "Status Code 405.", http.StatusMethodNotAllowed)
	}
}

func ipProcess(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ip" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Your IP = ")
		ip, err := getIP(r)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("No valid ip"))
		}
		w.WriteHeader(200)
		w.Write([]byte(ip))
	default:
		fmt.Fprintf(w, "Method %v Not Allowed.\n", r.Method)
		http.Error(w, "Status Code 405.", http.StatusMethodNotAllowed)
	}

}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}
