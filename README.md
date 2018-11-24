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


## Use with [jsontime](https://github.com/liamylian/jsontime)

```go
package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/liamylian/json-hashids"
	"github.com/liamylian/jsontime"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	json.RegisterExtension(&jsontime.CustomTimeExtension{})
	e := jsonhashids.NewHashIDsExtension("abcdef", 10)
	json.RegisterExtension(e)
}

type Book struct {
	Id        int       `json:"id" hashids:"true"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" time_format:"sql_datetime"`
}

func main() {
	book := Book{
		Id:        1,
		Name:      "Jane Eyre",
		CreatedAt: time.Now(),
	}

	bytes, _ := json.Marshal(book)
	
	// output: {"id":"qO3NnD68BX","name":"Jane Eyre","created_at":"2018-11-24 16:59:43"}
	fmt.Printf("%s", bytes)
}
```