package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hakan-einarsson/permitbridge/pkg/authorization"
	"github.com/hakan-einarsson/permitbridge/pkg/jsonparser"
	"github.com/hakan-einarsson/permitbridge/pkg/schema"
)

type AuthorizeData struct {
	UserId string `json:"userid"`
	Asset string `json:"asset"`
	Action string `json:"action"`
}

func IndexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	//log request
	fmt.Println(request.Method, request.URL.Path, request.RemoteAddr, request.UserAgent())
	http.ServeFile(responseWriter, request, "./views/index.html")
}

func AuthorizeHandler(responseWriter http.ResponseWriter, request *http.Request){
	//get user id, asset and action from post request
	var authorizeData AuthorizeData
	err := json.NewDecoder(request.Body).Decode(&authorizeData)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}
	
	authorized, err := authorization.AuthorizeHandler(authorizeData.UserId, authorizeData.Asset, authorizeData.Action)
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

func RolesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	roles, err := jsonparser.ReadRoles()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	userId := request.URL.Query().Get("userid")
	rolesList := authorization.RolesHandler(roles, userId)
	err = json.NewEncoder(responseWriter).Encode(rolesList)
	if err != nil {
		http.Error(responseWriter, "Error encoding the response into JSON: "+err.Error(), http.StatusInternalServerError)
	}

}

func UsersHandler(responseWriter http.ResponseWriter, request *http.Request) {
	users, err := jsonparser.ReadUsers()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	role := request.URL.Query().Get("role")
	userId := request.URL.Query().Get("userid")
	err = json.NewEncoder(responseWriter).Encode(authorization.UsersHandler(users, role, userId))
	if err != nil {
		http.Error(responseWriter, "Error encoding the response into JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func AssetsHandler(responseWriter http.ResponseWriter, request *http.Request) {
	roles, err := jsonparser.ReadRoles()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	role := request.URL.Query().Get("role")
	userId := request.URL.Query().Get("userid")
	assets, err := authorization.AssetsHandler(roles, role, userId)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(responseWriter).Encode(assets)
	if err != nil {
		http.Error(responseWriter, "Error encoding the response into JSON: "+err.Error(), http.StatusInternalServerError)
	}
}

func RolesSchemaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		roles, err := jsonparser.ReadRoles()
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(responseWriter).Encode(roles)
	case "POST":
		roles, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()
		err = schema.SaveRolesSchema(roles)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	case "PUT":
		roles, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()
		err = schema.AddRolesSchema(roles)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	case "PATCH":
		roles, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()
		err = schema.UpdateRolesSchema(roles)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	case "DELETE":
		rolesData, err := io.ReadAll(request.Body)
		//roles as a string array from rolesData
		var roles []string
		err = json.Unmarshal(rolesData, &roles)

		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()
		err = schema.DeleteFromRolesSchema(roles)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UsersSchemaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		users, err := jsonparser.ReadUsers()
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(responseWriter).Encode(users)
	case "POST":
		users, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
	
		err = schema.SaveUsersSchema(users)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	case "PUT":
		users, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		err  = schema.AddUsersSchema(users)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	case "PATCH":
		users, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusBadRequest)
			return
		}
		defer request.Body.Close()
		err = schema.UpdateUsersSchema(users)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
	case "DELETE":
		usersData, err := io.ReadAll(request.Body)
		//users as a string array from usersData
		var users []string
		err = json.Unmarshal(usersData, &users)

		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()
		err = schema.DeleteFromUsersSchema(users)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}