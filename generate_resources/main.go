package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	core "github.com/elimity-com/scim/schema"
	"github.com/scim2/tools/schema"
	gen "github.com/scim2/tools/structs/generate"
)

func main() {
	_ = os.MkdirAll("testdata", 0777)

	us, _ := core.CoreUserSchema().MarshalJSON()
	_ = ioutil.WriteFile("testdata/user.json", us, 0644)

	gs, _ := core.CoreGroupSchema().MarshalJSON()
	_ = ioutil.WriteFile("testdata/group.json", gs, 0644)

	e, _ := core.ExtensionEnterpriseUser().MarshalJSON()
	_ = ioutil.WriteFile("testdata/enterprise_user.json", e, 0644)

	generate("testdata/user.json", "resource_user.go", "server")
	generate("testdata/group.json", "resource_group.go", "server")
	generate("testdata/enterprise_user.json", "resource_enterpriseUser.go", "server")

	_ = os.RemoveAll("testdata")
}

func generate(path, out, pkgName string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var s schema.ReferenceSchema
	if err := json.Unmarshal(raw, &s); err != nil {
		panic(err)
	}

	g, err := gen.NewStructGenerator(s)
	if err != nil {
		panic(err)
	}

	b := &bytes.Buffer{}
	b.WriteString("// Do not edit. This file is auto-generated.\n")
	b.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	_, _ = g.Generate().WriteTo(b)

	if err := ioutil.WriteFile(out, b.Bytes(), 0644); err != nil {
		panic(err)
	}
}
