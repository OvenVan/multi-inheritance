package creature

import (
	"github.com/ovenvan/multi-inheritance/object"
)

type Creature interface {
	object.Object
	Create()
}


