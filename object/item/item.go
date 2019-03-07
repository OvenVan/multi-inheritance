package item

import "github.com/ovenvan/multi-inheritance/object"

type Item interface {
	object.Object
	ItemXXX()		//own method
}