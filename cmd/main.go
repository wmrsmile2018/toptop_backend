package main

import (
	"log"
	awesomeProject "top-top"
	"top-top/pkg/handler"
	"top-top/pkg/repository"
	"top-top/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(awesomeProject.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

//import (
//	"crypto/tls"
//	"encoding/json"
//	"fmt"
//	"log"
//	"mime"
//	"net/http"
//	"github.com/twilio/twilio-go"
//	api "github.com/twilio/twilio-go/rest/api/v2010"
//)
//)
//
//func taskHandler(w http.ResponseWriter, req *http.Request) {
//	if req.URL.Path == "/api/v1/auth" {
//		if req.Method == http.MethodPost {
//			createTaskHandler(w, req)
//		} else {
//			http.Error(w, fmt.Sprintf("expect method POST at /auth/, got %v", req.Method), http.StatusMethodNotAllowed)
//			return
//		}
//	} else {
//		http.Error(w, "Bad Request", http.StatusBadRequest)
//		return
//	}
//	//if req.URL.Path == "/api/v1/remote-execution" {
//	//	if req.Method == http.MethodPost {
//	//		createTaskHandler(w, req)
//	//	} else {
//	//		http.Error(w, fmt.Sprintf("expect method POST at /task/, got %v", req.Method), http.StatusMethodNotAllowed)
//	//		return
//	//	}
//	//} else {
//	//	http.Error(w, "Bad Request", http.StatusBadRequest)
//	//	return
//	//}
//}
//
//func authHandler(w http.ResponseWriter, req *http.Request) {
//	log.Printf("handling auth at %s\n", &req.URL.Path)
//
//	type RequestAuth struct {
//		Phone string `json:"phone"`
//	}
//
//	//Check if content is json
//	contentType := req.Header.Get("Content-Type")
//	mediatype, _, err := mime.ParseMediaType(contentType)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	if mediatype != "application/json" {
//		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
//		return
//	}
//
//	//decode json
//	dec := json.NewDecoder(req.Body)
//	dec.DisallowUnknownFields()
//	var rt RequestAuth
//	if err := dec.Decode(&rt); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//}
//
//func createTaskHandler(w http.ResponseWriter, req *http.Request) {
//	log.Printf("handling task create at %s\n", req.URL.Path)
//
//	type RequestTask struct {
//		Cmd   string `json:"cmd"`
//		Os    string `json:"os"`
//		Stdin string `json:"stdin"`
//	}
//
//	contentType := req.Header.Get("Content-Type")
//	mediatype, _, err := mime.ParseMediaType(contentType)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	if mediatype != "application/json" {
//		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
//		return
//	}
//	dec := json.NewDecoder(req.Body)
//	dec.DisallowUnknownFields()
//	var rt []RequestTask
//	if err := dec.Decode(&rt); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	type Response struct {
//		Stdout string `json:"stdout"`
//		Stderr string `json:"stderr"`
//	}
//	var response Response
//	if rt[0].Os == "windows" {
//		response.Stdout, response.Stderr = WinShellExe(rt[0].Cmd)
//	} else {
//		response.Stderr = "Incorrect OS"
//	}
//	js, _ := json.Marshal(response)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(js)
//}
//
//func init() {
//	err := godotenv.Load(".env")
//}
//
//func main() {
//	cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
//	s := &http.Server{
//		Addr:    ":8085",
//		Handler: nil,
//		TLSConfig: &tls.Config{
//			Certificates: []tls.Certificate{cert},
//		},
//	}
//	log.Printf("Server started")
//	http.HandleFunc("/api/v1/remote-execution", taskHandler)
//	log.Fatal(s.ListenAndServeTLS("", ""))
//}
