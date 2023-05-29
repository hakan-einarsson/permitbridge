package authorization

import (
	"github.com/hakan-einarsson/permitbridge/pkg/jsonparser"
)

func AuthorizeHandler(userId string, asset string, action string) (bool, error) {
	roles, err := jsonparser.ReadRoles()
	if err != nil {
		return false, err
	}
	user, err := jsonparser.ReadUser(userId)
	if err != nil {
		return false, err
	}
	filteredRoles := filterRoles(roles, user.Roles)
	return authorize(filteredRoles, asset, action)
}

func isAssetInRole(role jsonparser.Role, asset string) bool {
	_, exists := role.Assets[asset]
	return exists
}

func filterRoles(roles []jsonparser.Role, roleNames []string) []jsonparser.Role {
	var filteredRoles []jsonparser.Role
	for _, role := range roles {
		if stringIncluded(roleNames, role.Name) {
			filteredRoles = append(filteredRoles, role)
		}
	}
	return filteredRoles
}


func authorize(roles []jsonparser.Role, asset string, action string) (bool, error) {
	for _, role := range roles {
		if permissions, exists := role.Assets["global"]; exists && stringIncluded(permissions, action) {
			return true, nil
		}
		if isAssetInRole(role, asset) {
			if stringIncluded(role.Assets[asset], action) {
				return true, nil
			}
		}
	}
	return false, nil
}

func RolesHandler(roles []jsonparser.Role, userId string) []jsonparser.Role {
	if userId != "" {
		user, err := jsonparser.ReadUser(userId)
		if err != nil {
			return nil
		}
		return filterRoles(roles, user.Roles)
	}
	return roles
}

func UsersHandler(users []jsonparser.User, role string, userId string) []jsonparser.User {
	if userId != "" {
		user, err := jsonparser.ReadUser(userId)
		if err != nil {
			return nil
		}
		return []jsonparser.User{user}
	}
	if role != "" {
		var filteredUsers []jsonparser.User
		for _, user := range users {
			if stringIncluded(user.Roles, role) {
				filteredUsers = append(filteredUsers, user)
			}
		}
		return filteredUsers
	}
	return users
}

func AssetsHandler(roles []jsonparser.Role, role string, userid string)([]string, error) {
	if userid != "" {
		return readAssetsForUserId(roles, userid)
	}
	if role != "" {
		return readAssetsForRole(roles, role)
	}
	return readAssets(roles)
}

func readAssetsForRole(roles []jsonparser.Role, roleName string) ([]string, error) {
	var assets []string
	for _, role := range roles {
		if role.Name == roleName {
			for assetName := range role.Assets {
				if(!slizeContains(assets, assetName)) {
					assets = append(assets, assetName)
				}
			}
		}
	}
	return assets, nil
}

func readAssets(roles []jsonparser.Role) ([]string, error) {
	var assets []string
	for _, role := range roles {
		for assetName := range role.Assets {
			if(!slizeContains(assets, assetName)) {
				assets = append(assets, assetName)
			}
		}
	}
	return assets, nil
}

func readAssetsForUserId(roles []jsonparser.Role, userId string)([]string, error) {
	user, err := jsonparser.ReadUser(userId)
	if err != nil {
		return nil, err
	}
	var filteredRoles = filterRoles(roles, user.Roles)
	var assets []string
	for _, role := range filteredRoles {
		for assetName := range role.Assets {
			if(!slizeContains(assets, assetName)) {
				assets = append(assets, assetName)
			}
		}
	}
	return assets, nil
}

func slizeContains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func stringIncluded(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
