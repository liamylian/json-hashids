package jsonhashids

import (
	"testing"
)

var json = NewConfigWithHashIDs("abcdefg", 10)

type Book struct {
	Id    int    `json:"id" hashids:"true"`
	Idu   uint   `json:"idu" hashids:"true"`
	Id8   int8   `json:"id8" hashids:"true"`
	Idu8  uint   `json:"idu8" hashids:"true"`
	Id16  int16  `json:"id16" hashids:"true"`
	Idu16 uint   `json:"idu16" hashids:"true"`
	Id32  int32  `json:"id32" hashids:"true"`
	Idu32 uint   `json:"idu32" hashids:"true"`
	Id64  int64  `json:"id64" hashids:"true"`
	Idu64 uint   `json:"idu64" hashids:"true"`
	Name  string `json:"name"`
}

func TestMarshalFormat(t *testing.T) {
	book := Book{
		Id:    1,
		Idu:   1,
		Id8:   1,
		Idu8:  1,
		Id16:  1,
		Idu16: 1,
		Id32:  1,
		Idu32: 1,
		Id64:  1,
		Idu64: 1,
		Name:  "Jane Eyre",
	}

	if bytes, err := json.Marshal(book); err != nil {
		t.Error(err)
	} else if string(bytes) != `{"id":"gYEL5rKBnd","idu":"gYEL5rKBnd","id8":"gYEL5rKBnd","idu8":"gYEL5rKBnd","id16":"gYEL5rKBnd","idu16":"gYEL5rKBnd","id32":"gYEL5rKBnd","idu32":"gYEL5rKBnd","id64":"gYEL5rKBnd","idu64":"gYEL5rKBnd","name":"Jane Eyre"}` {
		t.Errorf("unexpected: %s\n", bytes)
	}

}

func TestUnmarshalFormat(t *testing.T) {
	bytes := []byte(`{"id":"gYEL5rKBnd","idu":"gYEL5rKBnd","id8":"gYEL5rKBnd","idu8":"gYEL5rKBnd","id16":"gYEL5rKBnd","idu16":"gYEL5rKBnd","id32":"gYEL5rKBnd","idu32":"gYEL5rKBnd","id64":"gYEL5rKBnd","idu64":"gYEL5rKBnd","name":"Jane Eyre"}`)

	book := Book{}
	if err := json.Unmarshal(bytes, &book); err != nil {
		t.Error(err)
	} else if book.Id != 1 || book.Idu != 1 ||
		book.Id8 != 1 || book.Idu8 != 1 ||
		book.Id16 != 1 || book.Idu16 != 1 ||
		book.Id32 != 1 || book.Idu32 != 1 ||
		book.Id64 != 1 || book.Idu64 != 1 ||
		book.Name != "Jane Eyre" {
		t.Errorf("unexpected: %v", book)
	}
}
