package runtime

type UndefinedObject string

var Undefined = UndefinedObject("(undefined)")
var _ = Object(Undefined)

// fetch a property and its owner. if this is a computed property
// or lazy-evaluated value, it is NOT evaluated
func (u UndefinedObject) Property(name string) (Object, PropertyValue) {
	panic("attempted to access property of undefined")
}

// fetch the object's own property
func (u UndefinedObject) PropertyOwn(name string) PropertyValue {
	panic("attempted to access property of undefined")
}

// fetch a property and its owner, always yielding an Object by
// evaluating computed properties
func (u UndefinedObject) PropertyComputed(name string) (Object, Object) {
	panic("attempted to access property of undefined")
}

// fetch the object's own property, always yielding an Object by
// evaluating computed properties
func (u UndefinedObject) PropertyOwnComputed(name string) Object {
	panic("attempted to access property of undefined")
}

// all properties
func (u UndefinedObject) Properties() []string {
	return nil
}

// own properties
func (u UndefinedObject) PropertiesOwn() []string {
	return nil
}

// true if the object has a property by the given name
func (u UndefinedObject) Has(name string) bool {
	return false
}

// true if the object has its own property by the given name
func (u UndefinedObject) HasOwn(name string) bool {
	return false
}

// fetch and evaluate the property by the given name
func (u UndefinedObject) Get(name string) Object {
	panic("attempted to access property of undefined")
}

// fetch and evaluate the object's own property by the given name
func (u UndefinedObject) GetOwn(name string) Object {
	panic("attempted to access property of undefined")
}

// write the given value to the property by the given name
func (u UndefinedObject) Set(name string, value PropertyValue) {
	panic("attempted to set property of undefined")
}

// write the given value to the proprerty by the given name, overwriting an
// existing value on a parent object
func (u UndefinedObject) SetOverwrite(name string, value PropertyValue) {
	panic("attempted to set property of undefined")
}

// delete the property by the given name
func (u UndefinedObject) Delete(name string) {
	panic("attempted to delete property of undefined")
}

// delete the property by the given name, even if it is inherited
func (u UndefinedObject) DeleteOverwrite(name string) {
	panic("attempted to delete property of undefined")
}

// weaken the property by the given name
func (u UndefinedObject) Weaken(name string) {
	panic("attempted to weaken property of undefined")
}

// weaken the property by the given name, even if it is inherited
func (u UndefinedObject) WeakenOverwrite(name string) {
	panic("attempted to weaken property of undefined")
}

// fetch and evaluate the value at the given index
func (u UndefinedObject) GetIndex(index Object) Object {
	panic("attempted to access index of undefined")
}

// set the value at the given index
func (u UndefinedObject) SetIndex(index Object, value Object) {
	panic("attempted to set index of undefined")
}

func (u UndefinedObject) Parents() []Object {
	return nil
}

func (u UndefinedObject) AddParent(p Object) {
	panic("attempted to modify isa of undefined")
}

func (u UndefinedObject) RemoveParent(p Object) {
	panic("attempted to modify isa of undefined")
}

func (u UndefinedObject) HasParent(p Object) bool {
	return false
}

// call the object with the given call info, returning an object
func (u UndefinedObject) Call(c Call) Object {
	panic("attempted to call undefined")
	return nil
}

func (u UndefinedObject) SetLastParent(p Object) {
}

func (u UndefinedObject) GetLastParent() Object {
	return nil
}

func (u UndefinedObject) Description(d *DescriptionOption) string {
	return "undefined"
}

func (u UndefinedObject) String() string {
	return u.Description(nil)
}

func IsUndefined(o Object) bool {
	return o == Undefined
}

func (u UndefinedObject) Object() Object {
	return u
}
