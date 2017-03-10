package runtime

import "strings"

type genericObject struct {
	properties map[string]PropertyValue
	isa        []Object
}

func NewObject() Object {
	return objectBase()
}

func objectBase() *genericObject {
	return &genericObject{
		make(map[string]PropertyValue),
		make([]Object, 0),
	}
}

// fetch a property and its owner. if this is a computed property
// or lazy-evaluated value, it is NOT evaluated
func (obj *genericObject) Property(name string) (Object, PropertyValue) {
	owners := append([]Object{obj}, obj.isa...)
	for _, owner := range owners {
		if val := owner.PropertyOwn(name); val != nil {
			return owner, val
		}
	}
	return nil, nil
}

// fetch the object's own property
func (obj *genericObject) PropertyOwn(name string) PropertyValue {
	return obj.properties[name]
}

// fetch a property and its owner, always yielding an Object by
// evaluating computed properties
func (obj *genericObject) PropertyComputed(name string) (Object, Object) {
	owner, val := obj.Property(name)
	return owner, computed(val, obj, owner)
}

// fetch the object's own property, always yielding an Object by
// evaluating computed properties
func (obj *genericObject) PropertyOwnComputed(name string) Object {
	val := obj.PropertyOwn(name)
	return computed(val, obj, obj)
}

// all properties
func (obj *genericObject) Properties() []string {
	props := obj.PropertiesOwn()
	for _, parent := range obj.isa {
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
func (obj *genericObject) PropertiesOwn() []string {
	props := make([]string, len(obj.properties))
	i := 0
	for name := range obj.properties {
		props[i] = name
		i++
	}
	return props
}

// true if the object has a property by the given name
func (obj *genericObject) Has(name string) bool {
	_, val := obj.Property(name)
	return val != nil
}

// true if the object has its own property by the given name
func (obj *genericObject) HasOwn(name string) bool {
	owner, _ := obj.Property(name)
	return owner == obj
}

// fetch and evaluate the property by the given name
func (obj *genericObject) Get(name string) Object {
	_, val := obj.PropertyComputed(name)
	return val
}

// fetch and evaluate the object's own property by the given name
func (obj *genericObject) GetOwn(name string) Object {
	owner, val := obj.Property(name)
	if owner != obj {
		return nil
	}
	return computed(val, obj, obj)
}

// write the given value to the property by the given name
func (obj *genericObject) Set(name string, value PropertyValue) {
	obj.properties[name] = value
}

// write the given value to the proprerty by the given name, overwriting an
// existing value on a parent object
func (obj *genericObject) SetOverwrite(name string, value PropertyValue) {
	owner, _ := obj.Property(name)
	if owner != nil {
		owner.Set(name, value)
	}
}

// delete the property by the given name
func (obj *genericObject) Delete(name string) {
	delete(obj.properties, name)
}

// delete the property by the given name, even if it is inherited
func (obj *genericObject) DeleteOverwrite(name string) {
	owner, _ := obj.Property(name)
	if owner == nil {
		return
	}
	owner.Delete(name)
}

// fetch and evaluate the value at the given index
func (obj *genericObject) GetIndex(index Object) Object {
	panic("unimplemented")
	return nil
}

// set the value at the given index
func (obj *genericObject) SetIndex(index Object, value Object) {
	panic("unimplemented")
}

func (obj *genericObject) Parents() []Object {
	return obj.isa
}

func (obj *genericObject) AddParent(p Object) {
	obj.isa = append(obj.isa, p)
}

func (obj *genericObject) RemoveParent(p Object) {
	panic("unimplemented")
}

func (obj *genericObject) HasParent(p Object) bool {
	panic("unimplemented")
	return false
}

// call the object with the given call info, returning an object
func (obj *genericObject) Call(c Call) Object {
	panic("unimplemented")
	return nil
}

// return a string description of the object
func (obj *genericObject) Description(d *DescriptionOption) string {

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
		var valueStr string
		switch v := value.(type) {
		case nil:
			valueStr = "undefined"
		case Object:
			if d.ignore[v] != 0 {
				valueStr = "(recursion)"
			} else {
				valueStr = v.Description(d)
				d.ignore[v]++
			}
		// TODO: computed
		default:
			valueStr = "(unknown)"
		}

		s += " = " + valueStr + "\n"
	}
	s += ")"
	return s
}

func indent(n int, s string) string {
	pfx := strings.Repeat(" ", n)
	lines := strings.Split(s, "\n")
	indented := make([]string, len(lines))
	for _, line := range lines {
		line = pfx + line
		indented = append(indented, line)
	}
	return strings.Join(indented, "\n")
}

func computed(val PropertyValue, obj Object, owner Object) Object {
	switch v := val.(type) {
	case nil:
		return nil
	case Object:
		return v
	case ComputedProperty:
		c := Call{} // TODO
		return v.function.Call(c)
	case LazyEvaluatedValue:
		return v(obj, owner)
	default:
		return nil
	}
	return nil
}
