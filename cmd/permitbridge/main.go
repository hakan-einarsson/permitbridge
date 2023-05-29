package main

import (
	"fmt"
	"net/http"

	"github.com/hakan-einarsson/permitbridge/pkg/server"
)

func main() {
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/authorize", server.AuthorizeHandler)
	http.HandleFunc("/roles", server.RolesHandler)
	http.HandleFunc("/users", server.UsersHandler)
	http.HandleFunc("/assets", server.AssetsHandler)
	http.HandleFunc("/schema/roles", server.RolesSchemaHandler)
	http.HandleFunc("/schema/users", server.UsersSchemaHandler)
	var port = "8000"
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}