package db

type FieldType interface {
	isFieldType()
}
type FieldTypeInt8 struct{}

func (f FieldTypeInt8) isFieldType() {}

type FieldTypeInt16 struct{}

func (f FieldTypeInt16) isFieldType() {}

type FieldTypeInt32 struct{}

func (f FieldTypeInt32) isFieldType() {}
