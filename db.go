package db

func Open(filePath string) (*Entity, error) {
	return &Entity{}, nil
}

type Entity struct{}

func (e *Entity) Next() (ok bool) {
	return
}

func (e *Entity) Data(fieldName ...string) *Data {
	return &Data{}
}

func (e *Entity) Close() error {
	return nil
}

func (e *Entity) Connect() (*Entity, error) {
	return &Entity{}, nil
}

type Data struct{}

func (d *Data) Scan(v ...any) error {
	return nil
}

func (d *Data) Update(v ...any) error {
	return nil
}
