package jsonparser

import (
	"encoding/json"
	"io"
	"os"
)

type Role struct {
	Name        string
	Assets map[string][]string
}

type Assets struct {
	Name   string
	Permissions []string
}

type RolesWrapper struct {
	Roles []Role `json:"roles"`
}

// ReadRoles reads roles from the json file.
func ReadRoles() ([]Role, error) {
	// Open the file.
	file, err := os.Open("pkg/schema/schema.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	// Parse the JSON.
	var wrapper RolesWrapper
	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Roles, nil
}

// WriteRoles writes roles to the JSON file.
func WriteRoles(roles []byte) error {
    // First unmarshal the roles to validate the input
    var wrapper RolesWrapper
    err := json.Unmarshal(roles, &wrapper)
    if err != nil {
        return err
    }

    // Write the JSON.
    err = os.WriteFile("pkg/schema/schema.json", roles, 0644)
    if err != nil {
        return err
    }
    return nil
}


