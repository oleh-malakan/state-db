package db

func Open(path string) (*Entity, error) {
	return &Entity{}, nil
}

type Entity struct{}

func (e *Entity) Next() (ok bool) {
	return
}

func (e *Entity) Close() error {
	return nil
}

func (e *Entity) Connect() (*Entity, error) {
	return &Entity{}, nil
}
