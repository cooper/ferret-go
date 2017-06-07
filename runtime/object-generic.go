package runtime

import (
	"../utils"
)

type genericObject struct {
	properties     map[string]uint
	weakProperties map[string]bool
	isa            []Object
	lastParent     Object
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
	}
}

// fetch a property and its owner. if this is a computed property
// or lazy-evaluated value, it is NOT evaluated
func (gobj *genericObject) Property(name string) (Object, PropertyValue) {
	owners := append([]Object{gobj.Object()}, gobj.isa...)
	for _, owner := range owners {
		if val := owner.PropertyOwn(name); val != nil {
			return owner, val
		}
	}
	return nil, nil
}

// fetch the object's own property
func (gobj *genericObject) PropertyOwn(name string) PropertyValue {
	if id, ok := gobj.properties[name]; ok {
		return retrieve(id)
	}
	return nil
}

// fetch a property and its owner, always yielding an Object by
// evaluating computed properties
func (gobj *genericObject) PropertyComputed(name string) (Object, Object) {
	obj := gobj
	owner, val := obj.Property(name)
	return owner, computed(name, val, obj, owner)
}

// fetch the object's own property, always yielding an Object by
// evaluating computed properties
func (gobj *genericObject) PropertyOwnComputed(name string) Object {
	obj := gobj
	val := obj.PropertyOwn(name)
	return computed(name, val, obj, obj)
}

// all properties
func (gobj *genericObject) Properties() []string {
	obj := gobj
	props := obj.PropertiesOwn()
	for _, parent := range gobj.isa {
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
func (gobj *genericObject) PropertiesOwn() []string {
	props := make([]string, len(gobj.properties))
	i := 0
	for name := range gobj.properties {
		props[i] = name
		i++
	}
	return props
}

// true if the object has a property by the given name
func (gobj *genericObject) Has(name string) bool {
	obj := gobj
	_, val := obj.Property(name)
	return val != nil
}

// true if the object has its own property by the given name
func (gobj *genericObject) HasOwn(name string) bool {
	obj := gobj
	owner, _ := obj.Property(name)
	return owner == obj
}

// fetch and evaluate the property by the given name
func (gobj *genericObject) Get(name string) Object {
	obj := gobj
	_, val := obj.PropertyComputed(name)
	return val
}

// fetch and evaluate the object's own property by the given name
func (gobj *genericObject) GetOwn(name string) Object {
	obj := gobj
	owner, val := obj.Property(name)
	if owner != obj {
		return nil
	}
	return computed(name, val, obj, obj)
}

// write the given value to the property by the given name
func (gobj *genericObject) Set(name string, value PropertyValue) {
	value = verifyPropertyValue(value)
	
	// check for old
	if existing, ok := gobj.properties[name]; ok {
		
		// unchanged
		if value == existing {
			return
		}
		
		// some other value is there
		if !gobj.weakProperties[name] {
			decrease(existing)
		}
	}

	// set new
	gobj.properties[name] = increase(value)
}

// write the given value to the proprerty by the given name, overwriting an
// existing value on a parent object
func (gobj *genericObject) SetOverwrite(name string, value PropertyValue) {
	obj := gobj
	owner, _ := obj.Property(name)
	if owner != nil {
		owner.Set(name, value)
	}
}

// delete the property by the given name
func (gobj *genericObject) Delete(name string) {
	
	// decrease the value only if it is not weak
	if val, ok := gobj.properties[name]; ok && !gobj.weakProperties[name] {
		decrease(val)
	}
	
	// delete
	delete(gobj.properties, name)
	delete(gobj.weakProperties, name)
}

// delete the property by the given name, even if it is inherited
func (gobj *genericObject) DeleteOverwrite(name string) {
	obj := gobj
	owner, _ := obj.Property(name)
	if owner == nil {
		return
	}
	owner.Delete(name)
}

// weaken the property by the given name
func (gobj *genericObject) Weaken(name string) {
	obj := gobj
	val := obj.PropertyOwn(name)

	// already weak
	if gobj.weakProperties[name] {
		return
	}

	// weaken
	gobj.weakProperties[name] = true
	decrease(val)
}

// weaken the property by the given name, even if it is inherited
func (gobj *genericObject) WeakenOverwrite(name string) {
	obj := gobj
	owner, _ := obj.Property(name)
	if owner == nil {
		return
	}
	owner.Weaken(name)
}

// fetch and evaluate the value at the given index
func (gobj *genericObject) GetIndex(index Object) Object {
	panic("unimplemented")
	return nil
}

// set the value at the given index
func (gobj *genericObject) SetIndex(index Object, value Object) {
	panic("unimplemented")
}

func (gobj *genericObject) Parents() []Object {
	return gobj.isa
}

func (gobj *genericObject) AddParent(p Object) {
	if p == nil {
		return
	}
	gobj.isa = append(gobj.isa, p)
}

func (gobj *genericObject) RemoveParent(p Object) {
	panic("unimplemented")
}

func (gobj *genericObject) HasParent(p Object) bool {
	panic("unimplemented")
	return false
}

// call the object with the given call info, returning an object
func (gobj *genericObject) Call(c Call) Object {
	panic("unimplemented")
	return nil
}

func (gobj *genericObject) GetLastParent() Object {
	return gobj.lastParent
}

func (gobj *genericObject) SetLastParent(p Object) {
	gobj.lastParent = p
}

// return a string description of the object
func (gobj *genericObject) Description(d *DescriptionOption) string {
	obj := gobj

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

func (gobj *genericObject) String() string {
	return gobj.Object().Description(nil)
}

func (gobj *genericObject) Object() Object {
	return gobj
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

func computed(name string, val PropertyValue, obj Object, owner Object) Object {
	switch v := val.(type) {
	case nil:
		return nil
	case Object:
		v.Object().SetLastParent(obj.Object())
		return v
	case ComputedProperty:
		v.code.SetLastParent(obj)
		o := v.code.Call(Call{})
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
