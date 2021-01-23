// +build ignore

// This program generates the resource_*.go files that contain the structures
// for every SCIM resource type.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	core "github.com/elimity-com/scim/schema"
	gen "github.com/scim2/tools/generate"
	"github.com/scim2/tools/schema"
)

func main() {
	_ = os.MkdirAll("testdata", 0777)

	us, _ := core.CoreUserSchema().MarshalJSON()
	_ = ioutil.WriteFile("testdata/user.json", us, 0644)

	gs, _ := core.CoreGroupSchema().MarshalJSON()
	_ = ioutil.WriteFile("testdata/group.json", gs, 0644)

	e, _ := core.ExtensionEnterpriseUser().MarshalJSON()
	_ = ioutil.WriteFile("testdata/enterprise_user.json", e, 0644)

	generate("testdata/user.json", "resource_user.go", "server", "testdata/enterprise_user.json")
	generate("testdata/group.json", "resource_group.go", "server")

	_ = os.RemoveAll("testdata")
}

func generate(path, out, pkgName string, extensionPaths ...string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var s schema.ReferenceSchema
	if err := json.Unmarshal(raw, &s); err != nil {
		panic(err)
	}

	var e []schema.ReferenceSchema
	for _, path := range extensionPaths {
		raw, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		var extension schema.ReferenceSchema
		if err := json.Unmarshal(raw, &extension); err != nil {
			panic(err)
		}
		e = append(e, extension)
	}
	g, err := gen.NewStructGenerator(s, e...)
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
