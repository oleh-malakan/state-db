package db

import (
	"time"
)

func MakeState() *State {
	return &State{}
}

type State struct{}

func (s *State) Open(filePath string) (*Entity, error) {
	return &Entity{}, nil
}

func (s *State) Close() error {
	return nil
}

type Entity struct{}

func (e *Entity) First() (ok bool) {
	return
}

func (e *Entity) Next() (ok bool) {
	return
}

func (e *Entity) NextStep(step int) (ok bool) {
	return
}

func (e *Entity) Prev() (ok bool) {
	return
}

func (e *Entity) PrevStep(step int) (ok bool) {
	return
}

func (e *Entity) Last() (ok bool) {
	return
}

func (e *Entity) Data(fieldName ...string) *Data {
	return &Data{}
}

func (e *Entity) Goto(link *Link) (ok bool) {
	return
}

func (e *Entity) Link() *Link {
	return &Link{}
}

func (e *Entity) Cut() Clipboard {
	return nil
}

func (e *Entity) CutFrom(link *Link) (Clipboard, error) {
	return nil, nil
}

func (e *Entity) CutTo(link *Link) (Clipboard, error) {
	return nil, nil
}

func (e *Entity) Copy() Clipboard {
	return nil
}

func (e *Entity) CopyFrom(link *Link) (Clipboard, error) {
	return nil, nil
}

func (e *Entity) CopyTo(link *Link) (Clipboard, error) {
	return nil, nil
}

func (e *Entity) Paste(clipboard ...Clipboard) error {
	return nil
}

func (e *Entity) Delete() error {
	return nil
}

func (e *Entity) DeleteFrom(link *Link) error {
	return nil
}

func (e *Entity) DeleteTo(link *Link) error {
	return nil
}

func (e *Entity) New(template *Template) error {
	return nil
}

func (e *Entity) Apply(template *Template) error {
	return nil
}

func (e *Entity) Template() *Template {
	return &Template{}
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

func (t *Tx) First() (ok bool) {
	return
}

func (t *Tx) Next() (ok bool) {
	return
}

func (t *Tx) NextStep(step int) (ok bool) {
	return
}

func (t *Tx) Prev() (ok bool) {
	return
}

func (t *Tx) PrevStep(step int) (ok bool) {
	return
}

func (t *Tx) Last() (ok bool) {
	return
}

func (t *Tx) Data(fieldName ...string) *Data {
	return &Data{}
}

func (t *Tx) Goto(link *Link) (ok bool) {
	return
}

func (t *Tx) Link() *Link {
	return &Link{}
}

func (t *Tx) Cut() Clipboard {
	return nil
}

func (t *Tx) CutFrom(link *Link) (Clipboard, error) {
	return nil, nil
}

func (t *Tx) CutTo(link *Link) (Clipboard, error) {
	return nil, nil
}

func (t *Tx) Copy() Clipboard {
	return nil
}

func (t *Tx) CopyFrom(link *Link) (Clipboard, error) {
	return nil, nil
}

func (t *Tx) CopyTo(link *Link) (Clipboard, error) {
	return nil, nil
}

func (t *Tx) Paste(clipboard ...Clipboard) error {
	return nil
}

func (t *Tx) Delete() error {
	return nil
}

func (t *Tx) DeleteFrom(link *Link) error {
	return nil
}

func (t *Tx) DeleteTo(link *Link) error {
	return nil
}

func (t *Tx) New(template *Template) error {
	return nil
}

func (t *Tx) Apply(template *Template) error {
	return nil
}

func (t *Tx) Template() *Template {
	return &Template{}
}

func (t *Tx) Commit() error {
	return nil
}

func (t *Tx) Rollback() error {
	return nil
}

type NullInt64 struct {
	Int64 int64
	Valid bool
}

type Data struct{}

func (d *Data) Scan(dest ...any) error {
	for _, d := range dest {
		switch v := d.(type) {
		case *Entity:
			*v = Entity{}
		case *Link:
			*v = Link{}
		case *Bookmarks:
			*v = Bookmarks{}
		case *Template:
			*v = Template{}
		case *NullInt64:
		}
	}
	return nil
}

func (d *Data) Update(args ...any) error {
	return nil
}

type Link struct{}

type Bookmarks struct{}

func (b *Bookmarks) Set(key string, link *Link) {}

func (b *Bookmarks) SetTTL(key string, link *Link, ttl time.Duration) {}

func (b *Bookmarks) Get(key string) *Link {
	return &Link{}
}

func (b *Bookmarks) Contains(key string) bool {
	return false
}

func (b *Bookmarks) Delete(key string) {}

type Clipboard interface{}

func CutSlice(from *Link, to *Link) (Clipboard, error) {
	return nil, nil
}

func CopySlice(from *Link, to *Link) (Clipboard, error) {
	return nil, nil
}

func DeleteSlice(from *Link, to *Link) error {
	return nil
}

func Cut(link *Link) (Clipboard, error) {
	return nil, nil
}

func Copy(link *Link) (Clipboard, error) {
	return nil, nil
}

func Paste(link *Link, clipboard ...Clipboard) error {
	return nil
}

func Delete(link *Link) error {
	return nil
}
