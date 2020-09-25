#!/bin/bash

# generate json schema files
mkdir testdata
go run generate_json/gen.go

git clone https://github.com/scim2/tools.git

go run tools/structs/generate/cli/main.go -p="testdata/user.json" -o="resource_user.go" -pkg="server"
go run tools/structs/generate/cli/main.go -p="testdata/group.json" -o="resource_group.go" -pkg="server"
go run tools/structs/generate/cli/main.go -p="testdata/enterprise_user.json" -o="resource_enterpriseUser.go" -pkg="server"

# remove generated json
rm testdata/*.json
rmdir testdata

rm -rf tools
