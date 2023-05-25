package authorization

import (
	"github.com/hakan-einarsson/permitbridge/pkg/jsonparser"
)

func AuthorizeHandler(roleName string, asset string, permission string) (bool, error) {
	roles, err := jsonparser.ReadRoles()
	if err != nil {
		return false, err
	}
	for _, role := range roles {
		if role.Name == roleName {
			permissions := role.Assets[asset]
			if stringIncluded(permissions, permission) {
				return true, nil
			}
		}
	}
	return false, nil
}

func PermissionsHandler(roles []string, asset string) ([]string, error) {
	permissions, err := findPermissions(roles, asset)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func AssetsHandler(roleNames []string, permissions []string) ([]string, error) { 
	assets, err := findAssets(roleNames, permissions)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func findPermissions(roleName []string, asset string) ([]string, error) {
	roles, err := jsonparser.ReadRoles()
	if err != nil {
		return nil, err
	}
	var permissions []string
	for _, role := range roles {
		if fullPermissions(permissions){
			return permissions, nil
		}
		if stringIncluded(roleName, role.Name) {
			globalPermissions, ok := role.Assets["global"]
			if ok {
				permissions = addToPermissions(permissions, globalPermissions)
			}
			newPermissions, ok := role.Assets[asset]
			if ok {
				permissions = addToPermissions(permissions, newPermissions)
			}
			break
		}
	}
	return permissions, nil
}

func fullPermissions(permissions []string) bool {
	var fullPermissions = []string{"read", "write", "delete"}
	for _, permission := range fullPermissions {
		if !stringIncluded(permissions, permission) {
			return false
		}
	}
	return true
}

func addToPermissions(permissions []string, newPermissions []string) []string {
	for _, permission := range newPermissions {
		if !stringIncluded(permissions, permission) {
			permissions = append(permissions, permission)
		}
	}
	return permissions
}

func stringIncluded(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func containsRequiredPermissions(perissionsRequired []string, permissionsProvided []string) bool {
	for _, permission := range perissionsRequired {
		if !stringIncluded(permissionsProvided, permission) {
			return false
		}
	}
	return true
}

func findAssets(roleNames []string, permissions []string) ([]string, error) {
	roles, err := jsonparser.ReadRoles()
	if err != nil {
		return nil, err
	}
	var assets []string
	for _, role := range roles {
		if(stringIncluded(roleNames, role.Name)){
			for assetName, assetPermissions := range role.Assets {
				if(assetName == "global" && containsRequiredPermissions(permissions, assetPermissions)){
					assets = []string{"global"}
					return assets, nil
				}
				if(containsRequiredPermissions(permissions, assetPermissions)){
					assets = append(assets, assetName)
				}
			}
		}
	}
	return assets, nil
}