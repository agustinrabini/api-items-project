# External clients example Jopit Go Toolkit

### Build url with QueryParams

```go
package main
import (
	...
)

func getWithQueryParams() {
	var response *rest.Response

	query := clients.Query()
	query.Add("caller.id", strconv.Itoa(sellerID))
	query.Add("attributes", itemAttributes)

	uri, err := clients.BuildURL([]string{"/items", itemID}, query)
    response = c.restClient.Get(uri, rest.Context(ctx))
}
```

### Build url with Body

```go
package main
import (
	...
)

func getWithBody() {
	var response *rest.Response

	url := fmt.Sprintf("%s/points/reference", config.ConfMap.APIBaseEndpoint)
	response = c.restClient.Post(url, userCoordinate, rest.Context(ctx))
}
```