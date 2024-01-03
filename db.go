package state

import (
	"time"
)

func Open(filePath string) (*Entity, error) {
	return &Entity{}, nil
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

func (e *Entity) Cut() *Clipboard {
	return &Clipboard{}
}

func (e *Entity) Copy() *Clipboard {
	return &Clipboard{}
}

func (e *Entity) Paste(clipboard ...Clipboard) error {
	return nil
}

func (e *Entity) New(template *Template) error {
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

func (t *Tx) Cut() *Clipboard {
	return &Clipboard{}
}

func (t *Tx) Copy() *Clipboard {
	return &Clipboard{}
}

func (t *Tx) Paste(clipboard ...Clipboard) error {
	return nil
}

func (t *Tx) New(template *Template) error {
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

type Clipboard struct{}

type Template struct{}
