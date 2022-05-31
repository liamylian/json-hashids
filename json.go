package jsonhashids

import (
	"reflect"
	"unsafe"

	"errors"
	"github.com/json-iterator/go"
	"github.com/speps/go-hashids"
)

var ErrNotInteger = errors.New("not integer")

func NewConfigWithHashIDs(salt string, minLength int) jsoniter.API {

	e := NewHashIDsExtension(salt, minLength)
	config := jsoniter.ConfigCompatibleWithStandardLibrary
	config.RegisterExtension(e)

	return config
}

func NewHashIDsExtension(salt string, minLength int) *HashIDsExtension {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, _ := hashids.NewWithData(hd)
	return &HashIDsExtension{
		hashid: h,
	}
}

type HashIDsExtension struct {
	hashid *hashids.HashID
	jsoniter.DummyExtension
}

func (extension *HashIDsExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {

	for _, binding := range structDescriptor.Fields {

		tag := binding.Field.Tag().Get("hashids")
		if tag != "true" {
			continue
		}

		switch binding.Field.Type().Kind() {
		case reflect.Int:
		case reflect.Uint:
		case reflect.Int8:
		case reflect.Uint8:
		case reflect.Int16:
		case reflect.Uint16:
		case reflect.Int32:
		case reflect.Uint32:
		case reflect.Int64:
		case reflect.Uint64:
		default:
			continue
		}

		typeName := binding.Field.Type().String()
		binding.Encoder = &funcEncoder{fun: func(ptr unsafe.Pointer, stream *jsoniter.Stream) {

			i, err := int64Val(typeName, ptr)
			if err != nil {
				stream.Error = err
				return
			}

			hashed, err := extension.hashid.EncodeInt64([]int64{i})
			if err != nil {
				stream.Error = err
				return
			}
			stream.WriteString(hashed)
		}}

		binding.Decoder = &funcDecoder{fun: func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {

			str := iter.ReadString()
			if str == "" {
				return
			}

			// See https://github.com/liamylian/json-hashids/issues/1
			ints, err := extension.hashid.DecodeInt64WithError(str)
			if len(ints) != 1 {
				iter.Error = err
				return
			}

			if err = setIntValue(typeName, ptr, ints[0]); err != nil {
				iter.Error = err
				return
			}
		}}
	}
}

type funcDecoder struct {
	fun jsoniter.DecoderFunc
}

func (decoder *funcDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	decoder.fun(ptr, iter)
}

type funcEncoder struct {
	fun         jsoniter.EncoderFunc
	isEmptyFunc func(ptr unsafe.Pointer) bool
}

func (encoder *funcEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	encoder.fun(ptr, stream)
}

func (encoder *funcEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	if encoder.isEmptyFunc == nil {
		return false
	}
	return encoder.isEmptyFunc(ptr)
}

func int64Val(typeName string, ptr unsafe.Pointer) (int64, error) {

	var i int64

	switch typeName {
	case "int":
		ip := (*int)(ptr)
		i = int64(*ip)
	case "uint":
		ip := (*uint)(ptr)
		i = int64(*ip)
	case "int8":
		ip := (*int8)(ptr)
		i = int64(*ip)
	case "uint8":
		ip := (*uint8)(ptr)
		i = int64(*ip)
	case "int16":
		ip := (*int16)(ptr)
		i = int64(*ip)
	case "uint16":
		ip := (*uint16)(ptr)
		i = int64(*ip)
	case "int32":
		ip := (*int32)(ptr)
		i = int64(*ip)
	case "uint32":
		ip := (*uint32)(ptr)
		i = int64(*ip)
	case "int64":
		ip := (*int64)(ptr)
		i = *ip
	case "uint64":
		ip := (*uint64)(ptr)
		i = int64(*ip)
	default:
		return 0, ErrNotInteger
	}

	return i, nil
}

func setIntValue(typeName string, ptr unsafe.Pointer, val int64) error {
	switch typeName {
	case "int":
		ip := (*int)(ptr)
		*ip = int(val)
	case "uint":
		ip := (*uint)(ptr)
		*ip = uint(val)
	case "int8":
		ip := (*int8)(ptr)
		*ip = int8(val)
	case "uint8":
		ip := (*uint8)(ptr)
		*ip = uint8(val)
	case "int16":
		ip := (*int16)(ptr)
		*ip = int16(val)
	case "uint16":
		ip := (*uint16)(ptr)
		*ip = uint16(val)
	case "int32":
		ip := (*int32)(ptr)
		*ip = int32(val)
	case "uint32":
		ip := (*uint32)(ptr)
		*ip = uint32(val)
	case "int64":
		ip := (*int64)(ptr)
		*ip = int64(val)
	case "uint64":
		ip := (*uint64)(ptr)
		*ip = uint64(val)
	default:
		return ErrNotInteger
	}

	return nil
}
