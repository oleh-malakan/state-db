package db

func Open(filePath string) (*Entity, error) {
	return &Entity{}, nil
}

type Entity struct{}

func (e *Entity) ByIndex(index string) (ok bool) {
	return
}

func (e *Entity) First() (ok bool) {
	return
}

func (e *Entity) Next() (ok bool) {
	return
}

func (e *Entity) Prev() (ok bool) {
	return
}

func (e *Entity) Last() (ok bool) {
	return
}

func (e *Entity) Data(fieldName ...string) *Data {
	return &Data{}
}

// New entity, construct new fields
func (e *Entity) New() {

}

// New entity, with clone fields and default values
func (e *Entity) Clone(fieldName ...string) *Data {
	return &Data{}
}

func (e *Entity) Begin() (*Tx, error) {
	return &Tx{}, nil
}

func (e *Entity) Close() error {
	return nil
}

func (e *Entity) Connect() (*Entity, error) {
	return &Entity{}, nil
}

type Tx struct{}

func (t *Tx) ByIndex(index string) (ok bool) {
	return
}

func (t *Tx) First() (ok bool) {
	return
}

func (t *Tx) Next() (ok bool) {
	return
}

func (t *Tx) Prev() (ok bool) {
	return
}

func (t *Tx) Last() (ok bool) {
	return
}

func (t *Tx) Data(fieldName ...string) *Data {
	return &Data{}
}

// New entity, construct new fields
func (t *Tx) New() {

}

// New entity, with clone fields and default values
func (t *Tx) Clone(fieldName ...string) *Data {
	return &Data{}
}

func (t *Tx) Commit() error {
	return nil
}

func (t *Tx) Rollback() error {
	return nil
}

type Data struct{}

func (d *Data) Scan(v ...any) error {
	return nil
}

func (d *Data) Update(v ...any) error {
	return nil
}
