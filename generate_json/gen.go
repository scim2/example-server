package main

import (
	"github.com/elimity-com/scim/schema"
	"io/ioutil"
)

func main() {
	u, _ := schema.CoreUserSchema().MarshalJSON()
	_ = ioutil.WriteFile("testdata/user.json", u, 0644)

	g, _ := schema.CoreGroupSchema().MarshalJSON()
	_ = ioutil.WriteFile("testdata/group.json", g, 0644)

	e, _ := schema.ExtensionEnterpriseUser().MarshalJSON()
	_ = ioutil.WriteFile("testdata/enterprise_user.json", e, 0644)
}
