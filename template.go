package db

type FieldType interface {
	isFieldType()
}

type Template struct{}

func (t *Template) Add(fieldName string, fieldType FieldType) error {
	return nil
}

func (t *Template) Get(fieldName string) FieldType {
	return nil
}

func (t *Template) All() map[string]FieldType {
	return nil
}

func (t *Template) Delete(fieldName string) error {
	return nil
}

type FieldTypeInt8 struct{}

func (f FieldTypeInt8) isFieldType() {}

type FieldTypeInt16 struct{}

func (f FieldTypeInt16) isFieldType() {}

type FieldTypeInt32 struct{}

func (f FieldTypeInt32) isFieldType() {}
