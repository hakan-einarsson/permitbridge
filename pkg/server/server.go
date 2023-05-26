package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hakan-einarsson/permitbridge/pkg/authorization"
	"github.com/hakan-einarsson/permitbridge/pkg/jsonparser"
)

func IndexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "views/index.html")
}

func AuthorizeHandler(responseWriter http.ResponseWriter, request *http.Request){
	role := request.URL.Query().Get("role")
	asset := request.URL.Query().Get("asset")
	permission := request.URL.Query().Get("permission")
	if role == "" || asset == "" || permission == "" {
		http.Error(responseWriter, "role, asset or permission not found", http.StatusBadRequest)
		return
	}
	authorized, err := authorization.AuthorizeHandler(role, asset, permission)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	if authorized {
		io.WriteString(responseWriter, "true")
	} else {
		io.WriteString(responseWriter, "false")
	}

}

func AssetsPermissionsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	asset := request.URL.Query().Get("asset")
	if asset == "" {
        http.Error(responseWriter, "'asset' query parameter is missing", http.StatusBadRequest)
        return
    }
	roles, ok := request.URL.Query()["role"]
	if !ok || len(roles) < 1 {
		http.Error(responseWriter, "role not found", http.StatusBadRequest)
		return
	}
	permissions, err := authorization.PermissionsHandler(roles, asset)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(responseWriter).Encode(permissions)
    if err != nil {
        http.Error(responseWriter, "Error encoding the response into JSON: "+err.Error(), http.StatusInternalServerError)
    }
}

func RolesAssetsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	roleNames, ok := request.URL.Query()["role"]
	if !ok || len(roleNames) < 1 {
		http.Error(responseWriter, "role not found", http.StatusBadRequest)
		return
	}
	permissions, ok := request.URL.Query()["permission"]
	if !ok || len(permissions) < 1 {
		permissions = []string{"read"}
	}
	assets, err := authorization.AssetsHandler(roleNames, permissions)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(responseWriter).Encode(assets)
}

func RolesSchemaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		roles, err := jsonparser.ReadRoles()
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(responseWriter).Encode(roles)
	} else if request.Method == "POST" {
		roles, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
	
		err = jsonparser.WriteRoles(roles)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	}
}

func UsersSchemaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		users, err := jsonparser.ReadUsers()
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(responseWriter).Encode(users)
	} else if request.Method == "POST" {
		users, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
	
		err = jsonparser.WriteUsers(users)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	}
}