// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     union.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

type UnionBytesStringTypeEnum int

const (
	UnionBytesStringTypeEnumBytes UnionBytesStringTypeEnum = 0

	UnionBytesStringTypeEnumString UnionBytesStringTypeEnum = 1
)

type UnionBytesString struct {
	Bytes     []byte
	String    string
	UnionType UnionBytesStringTypeEnum
}

func writeUnionBytesString(r *UnionBytesString, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionBytesStringTypeEnumBytes:
		return vm.WriteBytes(r.Bytes, w)
	case UnionBytesStringTypeEnumString:
		return vm.WriteString(r.String, w)
	}
	return fmt.Errorf("invalid value for *UnionBytesString")
}

func NewUnionBytesString() *UnionBytesString {
	return &UnionBytesString{}
}

func (_ *UnionBytesString) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionBytesString) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionBytesString) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionBytesString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionBytesString) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionBytesString) SetString(v string)  { panic("Unsupported operation") }
func (r *UnionBytesString) SetLong(v int64) {
	r.UnionType = (UnionBytesStringTypeEnum)(v)
}
func (r *UnionBytesString) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Bytes{Target: (&r.Bytes)}
	case 1:
		return &types.String{Target: (&r.String)}
	}
	panic("Unknown field index")
}
func (_ *UnionBytesString) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionBytesString) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionBytesString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionBytesString) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionBytesString) Finalize()                        {}

func (r *UnionBytesString) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte("null"), nil
	}
	switch r.UnionType {
	case UnionBytesStringTypeEnumBytes:
		return json.Marshal(map[string]interface{}{"bytes": r.Bytes})
	case UnionBytesStringTypeEnumString:
		return json.Marshal(map[string]interface{}{"string": r.String})
	}
	return nil, fmt.Errorf("invalid value for *UnionBytesString")
}

func (r *UnionBytesString) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if value, ok := fields["bytes"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.Bytes)
	}
	if value, ok := fields["string"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.String)
	}
	return fmt.Errorf("invalid value for *UnionBytesString")
}
