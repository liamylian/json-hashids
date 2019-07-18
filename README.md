# json-hashids

A [json iterator](https://github.com/json-iterator/go) extension that marshal integer to short unique id

## Usage

100% compatibility with standard lib

Replace
```go
import "encoding/json"

json.Marshal(&data)
json.Unmarshal(input, &data)
```

with
```go
import "github.com/liamylian/jsonhashids"

var json = NewConfigWithHashIDs("abcdefg", 10)

json.Marshal(&data)
json.Unmarshal(input, &data)
```


## Example

```go
package main

import(
	"fmt"
	"github.com/liamylian/json-hashids"
	"time"
)

var json = jsonhashids.NewConfigWithHashIDs("abcdefg", 10)

type Book struct {
	Id    int    `json:"id" hashids:"true"`
	Name  string `json:"name"`
}

func main() {
	book := Book {
		Id:          1,
		Name:        "Jane Eyre",
	}
	
	bytes, _ := json.Marshal(book)
	
	// output: {"id":"gYEL5rKBnd","name":"Jane Eyre"}
	fmt.Printf("%s", bytes)
}

```
