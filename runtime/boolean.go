package runtime

var True = Boolean(true)
var False = Boolean(false)

type Boolean bool

var _ = Object(True)

// fetch a property and its owner. if this is a computed property
// or lazy-evaluated value, it is NOT evaluated
func (b Boolean) Property(name string) (Object, PropertyValue) {
	return nil, nil
}

// fetch the object's own property
func (b Boolean) PropertyOwn(name string) PropertyValue {
	return nil
}

// fetch a property and its owner, always yielding an Object by
// evaluating computed properties
func (b Boolean) PropertyComputed(name string) (Object, Object) {
	return nil, nil
}

// fetch the object's own property, always yielding an Object by
// evaluating computed properties
func (b Boolean) PropertyOwnComputed(name string) Object {
	return nil
}

// all properties
func (b Boolean) Properties() []string {
	return nil
}

// own properties
func (b Boolean) PropertiesOwn() []string {
	return nil
}

// true if the object has a property by the given name
func (b Boolean) Has(name string) bool {
	return false
}

// true if the object has its own property by the given name
func (b Boolean) HasOwn(name string) bool {
	return false
}

// fetch and evaluate the property by the given name
func (b Boolean) Get(name string) Object {
	return nil
}

// fetch and evaluate the object's own property by the given name
func (b Boolean) GetOwn(name string) Object {
	return nil
}

// write the given value to the property by the given name
func (b Boolean) Set(name string, value PropertyValue) {
	panic("attempted to set property of boolean")
}

// write the given value to the proprerty by the given name, overwriting an
// existing value on a parent object
func (b Boolean) SetOverwrite(name string, value PropertyValue) {
	panic("attempted to set property of boolean")
}

// delete the property by the given name
func (b Boolean) Delete(name string) {
	panic("attempted to delete property of boolean")
}

// delete the property by the given name, even if it is inherited
func (b Boolean) DeleteOverwrite(name string) {
	panic("attempted to delete property of boolean")
}

// weaken the property by the given name
func (b Boolean) Weaken(name string) {
	panic("attempted to weaken property of boolean")
}

// weaken the property by the given name, even if it is inherited
func (b Boolean) WeakenOverwrite(name string) {
	panic("attempted to weaken property of boolean")
}

// fetch and evaluate the value at the given index
func (b Boolean) GetIndex(index Object) Object {
	panic("attempted to access index of boolean")
}

// set the value at the given index
func (b Boolean) SetIndex(index Object, value Object) {
	panic("attempted to set index of boolean")
}

func (b Boolean) Parents() []Object {
	return nil
}

func (b Boolean) AddParent(p Object) {
	panic("attempted to modify isa of boolean")
}

func (b Boolean) RemoveParent(p Object) {
	panic("attempted to modify isa of boolean")
}

func (b Boolean) HasParent(p Object) bool {
	return false
}

func (b Boolean) SetLastParent(p Object) {
}

func (b Boolean) GetLastParent() Object {
	return nil
}

// call the object with the given call info, returning an object
func (b Boolean) Call(c Call) Object {
	panic("attempted to call boolean")
	return nil
}

func (b Boolean) Description(d *DescriptionOption) string {
	if b {
		return "true"
	}
	return "false"
}

func (b Boolean) String() string {
	return b.Description(nil)
}

func (b Boolean) Object() Object {
	return b
}
