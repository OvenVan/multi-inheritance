package human

import "github.com/ovenvan/multi-inheritance/object/creature"

type Human interface {
	creature.Creature
	DNOW()
}