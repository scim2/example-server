package server

import (
	"github.com/elimity-com/scim"
	"github.com/elimity-com/scim/optional"
)

//go:generate go run gen.go

var server = scim.Server{
	Config: scim.ServiceProviderConfig{
		DocumentationURI: optional.NewString("https://example.com/scim/docs"),
	},
	ResourceTypes: []scim.ResourceType{
		userType,
	},
}
