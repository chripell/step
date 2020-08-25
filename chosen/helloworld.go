// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START gae_go111_app]

// Sample helloworld is an App Engine app.
package main

// [START import]
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

// [END import]
// [START main_func]

func main() {
	// http.HandleFunc("/", indexHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
	// [END setting_port]
	r := mux.NewRouter()
	r.HandleFunc("/login", GetLoginURL)
	r.HandleFunc("/who", who)
	r.HandleFunc("/", indexHandler)
	http.Handle("/", r)
	go http.ListenAndServe(":"+port, nil)
	appengine.Main()
}

// [END main_func]

// [START indexHandler]

func who(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	w.Write([]byte(fmt.Sprintf(`email: %s authdomain: %s id: %s cid: %s`, u.Email, u.AuthDomain, u.ID, u.ClientID)))
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, Gorilla World!")
}

func GetLoginURL(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	url, err := user.LoginURL(ctx, "/")

	fmt.Println(url, err)

	w.Write([]byte(fmt.Sprintf(`<a href="%s">Login</a>`, url)))
}

// [END indexHandler]
// [END gae_go111_app]