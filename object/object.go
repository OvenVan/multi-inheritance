package object

type Object interface {
	GetID() uint
	Register(manager ObjectManager) uint
}

type Obj struct {	//need to be public
	id uint
}
func (this *Obj) GetID() uint{
	return this.id
}
func (this *Obj) Register(m ObjectManager) uint{
	this.id = *m.Serial()
	*m.Serial()++
	return this.id
}