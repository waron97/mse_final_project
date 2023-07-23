package core

import "encoding/gob"

func Init() {
	gob.Register(Vector{})
	gob.Register(Document{})
}
