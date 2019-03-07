package object

type ObjectManager interface {
	Serial() *uint
}
type ObjManager struct {
	serial uint
}
func (this *ObjManager) Serial() *uint{
	return &this.serial
}
func NewObjectManager() ObjectManager{
	return &ObjManager{0}
}

