package jsonparser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

type Role struct {
	Name        string 			`json:"name"`
	Assets map[string][]string 	`json:"assets"`
}

type Assets struct {
	Name   string 				`json:"name"`
	Permissions []string 		`json:"permissions"`
}

type User struct {
	Id	string 					`json:"id"`
	Roles	[]string 			`json:"roles"`
}



// ReadRoles reads Roles from the json file.
func ReadRoles() ([]Role, error) {
	// Open the file.
	file, err := os.Open("pkg/schema/roles.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	// Parse the JSON.
	var roles []Role
	err = json.Unmarshal(data, &roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func ReadUsers()([]User, error) {
	file, err := os.Open("pkg/schema/users.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func ReadUser(Id string) (User, error) {
	users, err := ReadUsers()
	if err != nil {
		return User{}, err
	}
	for _, user := range users {
		if user.Id == Id {
			return user, nil
		}
	}
	return User{}, nil
}

// WriteRoles writes Roles to the JSON file.
func WriteRoles(roles []byte) error {
	err := ValidateRoles(roles)
	if err != nil {
		return err
	}
    // Write the JSON.
    err = os.WriteFile("pkg/schema/roles.json", roles, 0644)
    if err != nil {
        return err
    }
    return nil
}

//for PUT method
func AddRoles(roles []byte) error {
    err := ValidateRoles(roles)
    if err != nil {
        return err
    }
    existingRoles, err := ReadRoles()
    if err != nil {
        return err
    }
    var newRoles []Role
    err = json.Unmarshal(roles, &newRoles)
    if err != nil {
        return err
    }
    //check for duplicates, if found, return error saying role already exists, else append to existing Roles
    for _, newRole := range newRoles {
        for _, existingRole := range existingRoles {
            if newRole.Name == existingRole.Name {
                return fmt.Errorf("Role %s already exists", newRole.Name)
            }
        }
    }
    existingRoles = append(existingRoles, newRoles...)
    roles, err = json.Marshal(existingRoles)
    if err != nil {
        return err
    }
    err = WriteRoles(roles)
    if err != nil {
        return err
    }
    return nil
}

//for PATCH method
func UpdateRoles(roles []byte) error {
	err := ValidateRoles(roles)
	if err != nil {
		return err
	}
	existingRoles, err := ReadRoles()
	if err != nil {
		return err
	}
	var newRoles []Role
	err = json.Unmarshal(roles, &newRoles)
	if err != nil {
		return err
	}

	//check for existing Roles, if found, update, else return error saying role does not exist
	for _, newRole := range newRoles {
		for i, existingRole := range existingRoles {
			if newRole.Name == existingRole.Name {
				existingRoles[i] = newRole
				break
			}
		}
	}

	roles, err = json.Marshal(existingRoles)
	if err != nil {
		return err
	}
	err = WriteRoles(roles)
	if err != nil {
		return err
	}
	return nil
}

//for DELETE method
func DeleteRoles(RolesToDelete []string) error {
	existingRoles, err := ReadRoles()
	if err != nil {
		return err
	}
	//check for existing Roles, if found, delete, else return error saying role does not exist
	for _, roleToDelete := range RolesToDelete {
		for i, existingRole := range existingRoles {
			if roleToDelete == existingRole.Name {
				existingRoles = append(existingRoles[:i], existingRoles[i+1:]...)
				break
			}
		}
	}

	roles, err := json.Marshal(existingRoles)
	if err != nil {
		return err
	}
	err = WriteRoles(roles)
	if err != nil {
		return err
	}
	return nil
}

func WriteUsers(users []byte) error {
	err := ValidateUsers(users)
	if err != nil {
		return err
	}
	// Write the JSON.
	err = os.WriteFile("pkg/schema/users.json", users, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AddUsers(users []byte) error {
	err := ValidateUsers(users)
	if err != nil {
		return err
	}
	existingUsers, err := ReadUsers()
	if err != nil {
		return err
	}
	var newUsers []User
	err = json.Unmarshal(users, &newUsers)
	if err != nil {
		return err
	}
	//check for duplicates, if found, return error saying user already exists, else append to existing users
	for _, newUser := range newUsers {
		for _, existingUser := range existingUsers {
			if newUser.Id == existingUser.Id {
				return fmt.Errorf("User %s already exists", newUser.Id)
			}
		}
	}
	existingUsers = append(existingUsers, newUsers...)
	users, err = json.Marshal(existingUsers)
	if err != nil {
		return err
	}
	err = WriteUsers(users)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUsers(users []byte) error {
	err := ValidateUsers(users)
	if err != nil {
		return err
	}
	existingUsers, err := ReadUsers()
	if err != nil {
		return err
	}
	var newUsers []User
	err = json.Unmarshal(users, &newUsers)
	if err != nil {
		return err
	}

	//check for existing users, if found, update, else return error saying user does not exist
	for _, newUser := range newUsers {
		for i, existingUser := range existingUsers {
			if newUser.Id == existingUser.Id {
				existingUsers[i] = newUser
				break
			}
		}
	}

	users, err = json.Marshal(existingUsers)
	if err != nil {
		return err
	}
	err = WriteUsers(users)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUsers(usersToDelete []string) error {
	existingUsers, err := ReadUsers()
	if err != nil {
		return err
	}
	//check for existing users, if found, delete, else return error saying user does not exist
	for _, userToDelete := range usersToDelete {
		for i, existingUser := range existingUsers {
			if userToDelete == existingUser.Id {
				existingUsers = append(existingUsers[:i], existingUsers[i+1:]...)
				break
			}
		}
	}

	users, err := json.Marshal(existingUsers)
	if err != nil {
		return err
	}
	err = WriteUsers(users)
	if err != nil {
		return err
	}
	return nil
}

func ValidateRoles(roles []byte) error {
	schemaLoader := gojsonschema.NewStringLoader(`{
		"$schema": "http://json-schema.org/draft-07/schema#",
		"type": "array",
		"items": {
			"type": "object",
			"required": ["name", "assets"],
			"properties": {
				"Name": {"type": "string"},
				"Assets": {
					"type": "object",
					"patternProperties": {
						"^[a-zA-Z0-9_-]+$": {
							"type": "array",
							"items": {"type": "string"}
						}
					}
				}
			}
		}
	}`)
	schema, err := gojsonschema.NewSchema(schemaLoader)
    if err != nil {
        return err
    }
	dataLoader := gojsonschema.NewBytesLoader(roles)
	result, err := schema.Validate(dataLoader)
    if err != nil {
        return err
    }

    // Print the validation result.
    if result.Valid() {
        return nil
    } else {
		errorMessage := ""
        for _, desc := range result.Errors() {
            errorMessage += fmt.Sprintf("- %s\n", desc)
        }
		return fmt.Errorf("JSON data is not valId:\n%s", errorMessage)
		
    }
}

func ValidateUsers(users []byte) error {
    schemaLoader := gojsonschema.NewStringLoader(`{
        "$schema": "http://json-schema.org/draft-07/schema#",
        "type": "array",
        "items": {
            "type": "object",
            "required": ["id", "roles"],
            "properties": {
                "Id": {"type": "string"},
                "Roles": {
                    "type": "array",
                    "items": {"type": "string"}
                }
            }
        }
    }`)
	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		return err
	}
	dataLoader := gojsonschema.NewBytesLoader(users)
	result, err := schema.Validate(dataLoader)
	if err != nil {
		return err
	}

	// Print the valIdation result.
	if result.Valid() {
		return nil
	} else {
		errorMessage := ""
		for _, desc := range result.Errors() {
			errorMessage += fmt.Sprintf("- %s\n", desc)
		}
		return fmt.Errorf("JSON data is not valId:\n%s", errorMessage)
	}
}


