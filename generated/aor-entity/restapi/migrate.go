package database

import (

	// DefaultImports
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/models"

	"github.com/roscopecoltran/admin-on-rest-server/server/models"
)

var (
	DefaultTables = []interface{}{
		&models.ListUsingGETOKBody{},
		&models.JwtAuthenticationRequest{},
		&models.FindEntityFieldsUsingGETOKBody{},
		&models.GrantedAuthority{},
		&models.JwtUserAuthorities{},
		&models.User{},
		&models.Apply{},
		&models.GetRolesOKBody{},
		&models.FindAllFieldsUsingGETOKBody{},
		&models.ListUsersGETOKBody{},
		&models.RoleUsers{},
		&models.UserRoles{},
		&models.ChoiceItem{},
		&models.JwtUser{},
		&models.Permission{},
		&models.Entity{},
		&models.IChoiceItem{},
		&models.GetSchemasUsingGETOKBody{},
		&models.DataSource{},
		&models.ResponseEntity{},
		&models.Role{},
		&models.Field{},
		&models.EntityFields{},
		&models.ListUsingGET1OKBody{},
	}
)
