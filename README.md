# fv-mapper

fv-mapper makes map[string]string from form values of http.Request.

### Installation

```
go get github.com/kunihiko-t/fv-mapper
```

### Usage

```golang
package main

import (
  "github.com/kunihiko-t/fv-mapper"
  "net/http"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  r.ParseForm()

  //Params: {"name_1": "foo", "name_2": "bar", "another":"buzz"}

  //Fetch All params
  m := fvm.GetMap(r)
  //Result: map[name_1:foo name_2:bar another:buzz]

  //For sequential params
  ms := fvm.GetMapSequential("name",r)
  //Result: map[name_1:foo name_2:bar]


  //Params: { "name_test": "foo","nameA": "bar" }

  //With camel key
  camel := fvm.GetCamelMap(false,r)
  //Result: map[nameTest:foo nameA:bar]

  //With snake key
  snake := fvm.GetSnakeMap(r)
  //Result: map[name_test:foo name_a:bar]

}

```

## License
MIT
