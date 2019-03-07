package creature

import (
	"fmt"
	"github.com/ovenvan/multi-inheritance/object"
)

type Crea struct {
	object.Obj		//construct struct by embedding interface.
	// struct 中绑定interface和struct的区别？
	// Object只实现了一个Obj实例，这个实例的作用是被继承，提供父类的代码，因此应该继承Obj，而非Object
}

func (t *Crea) Create(){
	fmt.Println("This is a Base Create Method")
	/*do something*/
}

func (t *Crea)GetID() uint{
	fmt.Println("Override GetID from Creature")
	return t.Obj.GetID()
	//t.GetID()  it's a recursive call
}
func (t *Crea)Register(m object.ObjectManager) uint{
	fmt.Println("Override Register from Creature")
	return t.Obj.Register(m)
	//t.GetID()  it's a recursive call
}

func NewCrea () Creature{
	return &Crea{}
}