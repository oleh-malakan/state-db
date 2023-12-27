package db

func Open() (*Entity, error) {
	return &Entity{}, nil
}

type Entity struct {}

func (e *Entity) Close() error {
	return nil
}