package restapi

import (
	"github.com/roscopecoltran/gin-swagger/generated/aor-entity/models"
)

var (
	RestApiTables = []interface{}{
		&models.FindAllFieldsUsingGETOKBody{},
		&models.ListUsingGET1OKBody{},
		&models.FindEntityFieldsUsingGETOKBody{},
		&models.IChoiceItem{},
		&models.JwtAuthenticationRequest{},
		&models.ListUsingGETOKBody{},
		&models.Apply{},
		&models.ListUsersGETOKBody{},
		&models.Field{},
		&models.JwtUser{},
		&models.Permission{},
		&models.ResponseEntity{},
		&models.GetSchemasUsingGETOKBody{},
		&models.Entity{},
		&models.UserRoles{},
		&models.DataSource{},
		&models.Role{},
		&models.User{},
		&models.ChoiceItem{},
		&models.JwtUserAuthorities{},
		&models.EntityFields{},
		&models.GrantedAuthority{},
		&models.RoleUsers{},
		&models.GetRolesOKBody{},
	}
)
