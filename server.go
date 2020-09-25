package server

import (
	"github.com/elimity-com/scim"
	"github.com/elimity-com/scim/optional"
	"github.com/elimity-com/scim/schema"
)

var Server = scim.Server{
	Config: scim.ServiceProviderConfig{
		DocumentationURI: optional.NewString("https://example.com/scim/docs"),
	},
	ResourceTypes: []scim.ResourceType{
		{
			ID:          optional.NewString("User"),
			Name:        "User",
			Endpoint:    "/Users",
			Description: optional.NewString("User Account"),
			Schema:      schema.CoreUserSchema(),
			SchemaExtensions: []scim.SchemaExtension{
				{Schema: schema.ExtensionEnterpriseUser()},
			},
			Handler: newUsersResourceHandler(),
		},
	},
}
