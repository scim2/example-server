# Template: SCIM Server Implementation

### Example Resource Handler
```go
package server

import (
	"github.com/elimity-com/scim"
	"net/http"
)

type handler struct{}

func (h handler) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	panic("implement me")
}

func (h handler) Get(r *http.Request, id string) (scim.Resource, error) {
	panic("implement me")
}

func (h handler) GetAll(r *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	panic("implement me")
}

func (h handler) Replace(r *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	panic("implement me")
}

func (h handler) Delete(r *http.Request, id string) error {
	panic("implement me")
}

func (h handler) Patch(r *http.Request, id string, request scim.PatchRequest) (scim.Resource, error) {
	panic("implement me")
}
```
