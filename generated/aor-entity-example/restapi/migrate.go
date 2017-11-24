package restapi

import (
	"github.com/roscopecoltran/aor-gin-swagger/generated/aor-entity-example/models"
)

var (
	RestApiTables = []interface{}{
		&models.Apply{},
		&models.Entity{},
		&models.JwtUser{},
		&models.JwtAuthenticationRequest{},
		&models.Field{},
		&models.User{},
		&models.ResponseEntity{},
		&models.FindAllFieldsUsingGETOKBody{},
		&models.Role{},
		&models.ListUsersGETOKBody{},
		&models.IChoiceItem{},
		&models.Permission{},
		&models.JwtUserAuthorities{},
		&models.UserRoles{},
		&models.DataSource{},
		&models.GetRolesOKBody{},
		&models.ListUsingGETOKBody{},
		&models.GetSchemasUsingGETOKBody{},
		&models.ChoiceItem{},
		&models.ListUsingGET1OKBody{},
		&models.FindEntityFieldsUsingGETOKBody{},
		&models.GrantedAuthority{},
		&models.EntityFields{},
		&models.RoleUsers{},
	}
)
