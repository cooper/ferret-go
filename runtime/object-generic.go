package runtime

import "github.com/cooper/ferret-go/utils"

type genericObject struct {
	properties     map[string]uint
	weakProperties map[string]bool
	isa            []Object
	lastParent     Object
	object         Object
}

func NewObject() Object {
	return objectBase()
}

func objectBase() *genericObject {
	return &genericObject{
		make(map[string]uint),
		make(map[string]bool),
		make([]Object, 0),
		nil,
		nil,
	}
}

// fetch a property and its owner. if this is a computed property
// or lazy-evaluated value, it is NOT evaluated
func (gob *genericObject) Property(name string) (Object, PropertyValue) {
	owners := append([]Object{gob.Object()}, gob.isa...)
	for _, owner := range owners {
		if val := owner.PropertyOwn(name); val != nil {
			return owner, val
		}
	}
	return nil, nil
}

// fetch the object's own property
func (gob *genericObject) PropertyOwn(name string) PropertyValue {
	if id, ok := gob.properties[name]; ok {
		return retrieve(id)
	}
	return nil
}

// fetch a property and its owner, always yielding an Object by
// evaluating computed properties
func (gob *genericObject) PropertyComputed(name string) (Object, Object) {
	obj := gob.Object()
	owner, val := obj.Property(name)
	return owner, computed(name, val, obj, owner)
}

// fetch the object's own property, always yielding an Object by
// evaluating computed properties
func (gob *genericObject) PropertyOwnComputed(name string) Object {
	obj := gob.Object()
	val := obj.PropertyOwn(name)
	return computed(name, val, obj, obj)
}

// all properties
func (gob *genericObject) Properties() []string {
	obj := gob.Object()
	props := obj.PropertiesOwn()
	for _, parent := range gob.isa {
		props = append(props, parent.PropertiesOwn()...)
	}
	seen := make(map[string]bool, len(props))
	uniq := make([]string, len(props))
	i := 0
	for _, name := range props {
		if seen[name] {
			continue
		}
		seen[name] = true
		uniq[i] = name
		i++
	}
	return uniq[:i]
}

// own properties
func (gob *genericObject) PropertiesOwn() []string {
	props := make([]string, len(gob.properties))
	i := 0
	for name := range gob.properties {
		props[i] = name
		i++
	}
	return props
}

// true if the object has a property by the given name
func (gob *genericObject) Has(name string) bool {
	obj := gob.Object()
	_, val := obj.Property(name)
	return val != nil
}

// true if the object has its own property by the given name
func (gob *genericObject) HasOwn(name string) bool {
	obj := gob
	owner, _ := obj.Property(name)
	return owner == obj
}

// fetch and evaluate the property by the given name
func (gob *genericObject) Get(name string) Object {
	obj := gob.Object()
	_, val := obj.PropertyComputed(name)
	return val
}

// fetch and evaluate the object's own property by the given name
func (gob *genericObject) GetOwn(name string) Object {
	obj := gob.Object()
	owner, val := obj.Property(name)
	if owner != obj {
		return nil
	}
	return computed(name, val, obj, obj)
}

// write the given value to the property by the given name
func (gob *genericObject) Set(name string, value PropertyValue) {
	value = verifyPropertyValue(value)

	// check for old
	if existing, ok := gob.properties[name]; ok {

		// unchanged
		if value == existing {
			return
		}

		// some other value is there
		if !gob.weakProperties[name] {
			decrease(existing)
		}
	}

	// set new
	gob.properties[name] = increase(value)
}

// write the given value to the proprerty by the given name, overwriting an
// existing value on a parent object
func (gob *genericObject) SetOverwrite(name string, value PropertyValue) {
	obj := gob.Object()
	owner, _ := obj.Property(name)
	if owner != nil {
		owner.Set(name, value)
	}
}

// delete the property by the given name
func (gob *genericObject) Delete(name string) {

	// decrease the value only if it is not weak
	if val, ok := gob.properties[name]; ok && !gob.weakProperties[name] {
		decrease(val)
	}

	// delete
	delete(gob.properties, name)
	delete(gob.weakProperties, name)
}

// delete the property by the given name, even if it is inherited
func (gob *genericObject) DeleteOverwrite(name string) {
	obj := gob.Object()
	owner, _ := obj.Property(name)
	if owner == nil {
		return
	}
	owner.Delete(name)
}

// weaken the property by the given name
func (gob *genericObject) Weaken(name string) {
	obj := gob.Object()
	val := obj.PropertyOwn(name)

	// already weak
	if gob.weakProperties[name] {
		return
	}

	// weaken
	gob.weakProperties[name] = true
	decrease(val)
}

// weaken the property by the given name, even if it is inherited
func (gob *genericObject) WeakenOverwrite(name string) {
	obj := gob.Object()
	owner, _ := obj.Property(name)
	if owner == nil {
		return
	}
	owner.Weaken(name)
}

// fetch and evaluate the value at the given index
func (gob *genericObject) GetIndex(index Object) Object {
	panic("unimplemented")
	return nil
}

// set the value at the given index
func (gob *genericObject) SetIndex(index Object, value Object) {
	panic("unimplemented")
}

func (gob *genericObject) Parents() []Object {
	return gob.isa
}

func (gob *genericObject) AddParent(p Object) {
	if p == nil {
		return
	}
	gob.isa = append(gob.isa, p)
}

func (gob *genericObject) RemoveParent(p Object) {
	panic("unimplemented")
}

func (gob *genericObject) HasParent(p Object) bool {
	panic("unimplemented")
	return false
}

// call the object with the given call info, returning an object
func (gob *genericObject) Call(c Call) Object {
	panic("unimplemented")
	return nil
}

func (gob *genericObject) GetLastParent() Object {
	return gob.lastParent
}

func (gob *genericObject) SetLastParent(p Object) {
	gob.lastParent = p
}

// return a string description of the object
func (gob *genericObject) Description(d *DescriptionOption) string {
	obj := gob.Object()

	// initial options
	if d == nil {
		d = &DescriptionOption{ignore: make(map[Object]uint)}
	}

	// properties
	s := "("
	for i, propName := range obj.Properties() {
		owner, value := obj.Property(propName)
		if i == 0 {
			s += "\n"
		}
		s += "    "

		// property name
		if owner == obj {
			s += propName
		} else {
			s += "(" + propName + ")"
		}

		// value
		s += " = " + valueStr(value, d) + "\n"
	}
	s += ")"
	return s
}

func (gob *genericObject) String() string {
	return gob.Object().Description(nil)
}

func (gob *genericObject) Object() Object {
	if gob.object != nil {
		return gob.object
	}
	return gob
}

func valueStr(value PropertyValue, d *DescriptionOption) string {
	switch v := value.(type) {
	case nil:
		return "undefined"
	case Object:
		if d.ignore[v] != 0 {
			return "(recursion)"
		} else {
			d.ignore[v]++
			return utils.Indent(4, v.Description(d))
		}
	case LazyEvaluatedValue, ComputedProperty:
		return "(computed)"
	default:
		return "(unknown)"
	}
}

func computed(name string, val PropertyValue, obj, owner Object) Object {
	switch v := val.(type) {
	case nil:
		return nil
	case Object:
		v = v.Object()
		v.SetLastParent(obj)
		return v
	case ComputedProperty:
		v.code.SetLastParent(obj)
		o := v.code.Call(Call{}).Object()
		if v.lazy {
			owner.Set(name, o)
		}
		return computed(name, o, obj, owner)
	case LazyEvaluatedValue:
		return computed(name, v(obj, owner), obj, owner)
	default:
		return nil
	}
}
