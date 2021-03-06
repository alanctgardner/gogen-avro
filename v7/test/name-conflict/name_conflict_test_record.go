// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type NameConflictTestRecord struct {
	Field_Schema bool `json:"Schema"`

	Field_Serialize bool `json:"Serialize"`

	Field_SchemaName bool `json:"SchemaName"`

	Field_MarshalJSON bool `json:"MarshalJSON"`

	Field_UnmarshalJSON bool `json:"UnmarshalJSON"`

	Field_AvroCRC64Fingerprint bool `json:"AvroCRC64Fingerprint"`

	Field_SetBoolean bool `json:"SetBoolean"`

	Field_SetInt bool `json:"SetInt"`

	Field_SetLong bool `json:"SetLong"`

	Field_SetFloat bool `json:"SetFloat"`

	Field_SetDouble bool `json:"SetDouble"`

	Field_SetBytes bool `json:"SetBytes"`

	Field_SetString bool `json:"SetString"`

	Field_Get bool `json:"Get"`

	Field_SetDefault bool `json:"SetDefault"`

	Field_AppendMap bool `json:"AppendMap"`

	Field_AppendArray bool `json:"AppendArray"`

	Field_NullField bool `json:"NullField"`

	Field_Finalize bool `json:"Finalize"`
}

const NameConflictTestRecordAvroCRC64Fingerprint = "\xcesIh9\x9f&\b"

func NewNameConflictTestRecord() *NameConflictTestRecord {
	return &NameConflictTestRecord{}
}

func DeserializeNameConflictTestRecord(r io.Reader) (*NameConflictTestRecord, error) {
	t := NewNameConflictTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeNameConflictTestRecordFromSchema(r io.Reader, schema string) (*NameConflictTestRecord, error) {
	t := NewNameConflictTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeNameConflictTestRecord(r *NameConflictTestRecord, w io.Writer) error {
	var err error
	err = vm.WriteBool(r.Field_Schema, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_Serialize, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SchemaName, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_MarshalJSON, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_UnmarshalJSON, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_AvroCRC64Fingerprint, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetBoolean, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetInt, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetLong, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetFloat, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetDouble, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetBytes, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetString, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_Get, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetDefault, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_AppendMap, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_AppendArray, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_NullField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_Finalize, w)
	if err != nil {
		return err
	}
	return err
}

func (r *NameConflictTestRecord) Serialize(w io.Writer) error {
	return writeNameConflictTestRecord(r, w)
}

func (r *NameConflictTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"Schema\",\"type\":\"boolean\"},{\"name\":\"Serialize\",\"type\":\"boolean\"},{\"name\":\"SchemaName\",\"type\":\"boolean\"},{\"name\":\"MarshalJSON\",\"type\":\"boolean\"},{\"name\":\"UnmarshalJSON\",\"type\":\"boolean\"},{\"name\":\"AvroCRC64Fingerprint\",\"type\":\"boolean\"},{\"name\":\"SetBoolean\",\"type\":\"boolean\"},{\"name\":\"SetInt\",\"type\":\"boolean\"},{\"name\":\"SetLong\",\"type\":\"boolean\"},{\"name\":\"SetFloat\",\"type\":\"boolean\"},{\"name\":\"SetDouble\",\"type\":\"boolean\"},{\"name\":\"SetBytes\",\"type\":\"boolean\"},{\"name\":\"SetString\",\"type\":\"boolean\"},{\"name\":\"Get\",\"type\":\"boolean\"},{\"name\":\"SetDefault\",\"type\":\"boolean\"},{\"name\":\"AppendMap\",\"type\":\"boolean\"},{\"name\":\"AppendArray\",\"type\":\"boolean\"},{\"name\":\"NullField\",\"type\":\"boolean\"},{\"name\":\"Finalize\",\"type\":\"boolean\"}],\"name\":\"NameConflictTestRecord\",\"type\":\"record\"}"
}

func (r *NameConflictTestRecord) SchemaName() string {
	return "NameConflictTestRecord"
}

func (_ *NameConflictTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NameConflictTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Boolean{Target: &r.Field_Schema}
	case 1:
		return &types.Boolean{Target: &r.Field_Serialize}
	case 2:
		return &types.Boolean{Target: &r.Field_SchemaName}
	case 3:
		return &types.Boolean{Target: &r.Field_MarshalJSON}
	case 4:
		return &types.Boolean{Target: &r.Field_UnmarshalJSON}
	case 5:
		return &types.Boolean{Target: &r.Field_AvroCRC64Fingerprint}
	case 6:
		return &types.Boolean{Target: &r.Field_SetBoolean}
	case 7:
		return &types.Boolean{Target: &r.Field_SetInt}
	case 8:
		return &types.Boolean{Target: &r.Field_SetLong}
	case 9:
		return &types.Boolean{Target: &r.Field_SetFloat}
	case 10:
		return &types.Boolean{Target: &r.Field_SetDouble}
	case 11:
		return &types.Boolean{Target: &r.Field_SetBytes}
	case 12:
		return &types.Boolean{Target: &r.Field_SetString}
	case 13:
		return &types.Boolean{Target: &r.Field_Get}
	case 14:
		return &types.Boolean{Target: &r.Field_SetDefault}
	case 15:
		return &types.Boolean{Target: &r.Field_AppendMap}
	case 16:
		return &types.Boolean{Target: &r.Field_AppendArray}
	case 17:
		return &types.Boolean{Target: &r.Field_NullField}
	case 18:
		return &types.Boolean{Target: &r.Field_Finalize}
	}
	panic("Unknown field index")
}

func (r *NameConflictTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NameConflictTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *NameConflictTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *NameConflictTestRecord) Finalize()                        {}

func (_ *NameConflictTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NameConflictTestRecordAvroCRC64Fingerprint)
}
