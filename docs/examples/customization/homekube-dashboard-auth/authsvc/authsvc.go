/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	//	"fmt"
	"net/http"
	"strconv"
)

// Sample authentication service returning several HTTP headers in response
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Query()
		//		if strings.ContainsAny(r.Header.Get("User"), "internal") {
		var perms = r.URL.Query().Get("perms")
		log.Print("Received perms: " + perms)
		if perms == "view" {
			w.Header().Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6Ii1KYTVyV25XNlB3RWI1YWozOVgwZDNGb3pTbklqLWdNUGVncWE2MDU3bWcifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi11c2VyLXRva2VuLTc0OWdxIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI5ZGI5ZjE1My00ZTE2LTQwMzYtYjFiZS1hZDE3Y2U4OWU5YzUiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZXJuZXRlcy1kYXNoYm9hcmQ6YWRtaW4tdXNlciJ9.TGbd-Ru-fZmYtcnaVkQa0JAbsFGzd7aZJspkcf1wCOTqk_HnjO48JQGq7VTQItPIbP-YNcLF1xBrUKmv4iCGuupTgI9gsZaEarxXDiHtdPWK2VbQ3zWZYi9Q8xdcirqRBtPTBhK_3yK758pKzE7CnIF9lfmfUeNJdsuQ6ThNSMlZSVZFCxJunZSa2x7Fcffd6KO8k_RKQLJUoYZRSv1kKKtyQLFkd09UYWJfLeGn3FAGSH39wwCZv3sNnwzZ-ixAU7xaYWo1-yQcW8pD3MQiK3WRuYpJPEr78JwS86TxXyUSwaQQZSDt88KLGtsr9BWxRp0CPypRQa-hnevyScJCBQ")
			//			fmt.Fprint(w, "ok")
			log.Print("Authorized " + perms)
		} else {
			rc := http.StatusForbidden
			if c := r.URL.Query().Get("code"); len(c) > 0 {
				c, _ := strconv.Atoi(c)
				if c > 0 && c < 600 {
					rc = c
				}
			}

			w.WriteHeader(rc)
			log.Print("Unauthorized " + r.URL.Path)
			//			fmt.Fprint(w, "unauthorized")
		}
	})
	http.ListenAndServe(":8080", nil)
}
