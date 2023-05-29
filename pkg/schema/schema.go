package schema

import (
	"github.com/hakan-einarsson/permitbridge/pkg/jsonparser"
)

func SaveRolesSchema(roles []byte)(error) {
	//add validation
	return jsonparser.WriteRoles(roles)
}

func AddRolesSchema(roles []byte)(error) {
	//add validation
	return jsonparser.AddRoles(roles)
}

func UpdateRolesSchema(roles []byte)(error) {
	//add validation
	return jsonparser.UpdateRoles(roles)
}

func DeleteFromRolesSchema(roles []string)(error) {
	return jsonparser.DeleteRoles(roles)
}

func SaveUsersSchema(users []byte)(error) {
	//add validation
	return jsonparser.WriteUsers(users)
}

func AddUsersSchema(users []byte)(error) {
	//add validation
	return jsonparser.AddUsers(users)
}

func UpdateUsersSchema(users []byte)(error) {
	//add validation
	return jsonparser.UpdateUsers(users)
}

func DeleteFromUsersSchema(users []string)(error) {
	//add validation
	return jsonparser.DeleteUsers(users)
}