package monster

import (
	"fmt"
	"github.com/ovenvan/multi-inheritance/object"
	"github.com/ovenvan/multi-inheritance/object/creature"
)

type Monster interface {
	creature.Creature
	Hatch()
}

type Monster_s struct {
	creature.Creature		//construct struct by embedding interface
	alive bool
}
func (t *Monster_s) Hatch(){
	t.Create()
	fmt.Println("After created, i was hatched by index",t.GetID())
	fmt.Println()
}

func (t *Monster_s) Create(){
	t.Creature.Create()
	fmt.Println("This is an Override Create Method from monster")
}

func (t *Monster_s)GetID() uint{
	fmt.Println("Override GetID from Monster")
	return t.Creature.GetID()
	//return t.Obj.GetID()		//可以调用父类的GetID方法，也可以直接调用父类的父类的方法
	//t.GetID()  recursive call
}

func NewMonster (m object.ObjectManager,skipcreature bool) Monster{
	mstr:=Monster_s{}
	if skipcreature==false{mstr.Creature = creature.NewCrea()}else{mstr.Creature = creature.NewCrea2()}
	mstr.Register(m)
	return &mstr
}