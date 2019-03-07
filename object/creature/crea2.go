package creature

import (
	"fmt"
	"github.com/ovenvan/multi-inheritance/object"
)

type Crea2 struct {
	object.Obj		//struct 中绑定interface和struct的区别？
	// Object只实现了一个Obj实例，这个实例的作用是被继承，提供父类的代码，因此应该继承Obj，而非Object
}

func (t *Crea2) Create(){
	fmt.Println("This is a Base Create Method from Crea2")
	/*do something*/
}

func NewCrea2 () Creature{
	return &Crea2{}
}