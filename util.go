package server

import (
	"github.com/elimity-com/scim"
	"reflect"
)

var mapType = reflect.TypeOf(map[string]interface{}{})

func rAttr2Map(attributes scim.ResourceAttributes) map[string]interface{} {
	return reflect.
		ValueOf(attributes).
		Convert(mapType).
		Interface().
	(map[string]interface{})
}
