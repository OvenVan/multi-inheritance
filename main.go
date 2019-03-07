package main

import (
	"github.com/ovenvan/multi-inheritance/object"
	"github.com/ovenvan/multi-inheritance/object/creature/monster"
)

func main() {
	crtrManager:=object.NewObjectManager()
	mstr:=monster.NewMonster(crtrManager,false)
	mstr1:=monster.NewMonster(crtrManager,true)
	mstr.Hatch()
	mstr1.Hatch()
}


