package main

import (
	"fmt"
	"net/http"

	"github.com/hakan-einarsson/permitbridge/pkg/server"
)

func main() {
	http.HandleFunc("/", server.IndexHandler)
	http.HandleFunc("/authorize", server.AuthorizeHandler)
	http.HandleFunc("/authorize/permissions", server.AssetsPermissionsHandler)
	http.HandleFunc("/authorize/assets", server.RolesAssetsHandler)
	// http.HandleFunc("/authorize/rolse", server.AuthorizeHandler)
	http.HandleFunc("/schema", server.EditSchemaHandler)
	var port = "8000"
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}