package runtime

import "../weakref"

type Object interface {

	// PROPERTIES

	// fetch a property and its owner. if this is a computed property
	// or lazy-evaluated value, it is NOT evaluated
	Property(name string) (owner Object, value PropertyValue)

	// fetch the object's own property
	PropertyOwn(name string) PropertyValue

	// fetch a property and its owner, always yielding an Object by
	// evaluating computed properties
	PropertyComputed(name string) (owner Object, value Object)

	// fetch the object's own property, always yielding an Object by
	// evaluating computed properties
	PropertyOwnComputed(name string) Object

	// slice of property names, optionally including inherited
	Properties() []string
	PropertiesOwn() []string

	// true if the object has a property by the given name
	Has(name string) bool
	HasOwn(name string) bool

	// fetch and evaluate the property by the given name
	Get(name string) Object
	GetOwn(name string) Object

	// write the given value to the property by the given name, optionally
	// overwriting an inherited property
	Set(name string, value PropertyValue)
	SetOverwrite(name string, value PropertyValue)

	// delete the property by the given name, optionally taking into
	// consideration inherited properties
	Delete(name string)
	DeleteOverwrite(name string)

	// weaken the property by the given name, optionally taking into
	// consideration inherited properties
	Weaken(name string)
	WeakenOverwrite(name string)

	// INDICES

	// fetch and evaluate or set the value at the given index
	GetIndex(index Object) Object
	SetIndex(index Object, value Object)

	// INHERITANCE

	// fetch parent objects
	Parents() []Object

	// add parent object
	AddParent(p Object)

	// remove parent object
	RemoveParent(p Object)

	// object has this parent
	HasParent(p Object) bool

	// OTHER

	// call the object with the given call info, returning an object
	Call(c Call) Object

	// set and fetch the object's most recent parent
	GetLastParent() Object
	SetLastParent(p Object)

	// return a string description of the object
	Description(d *DescriptionOption) string
	String() string

	// returns the object itself
	Object() Object
}

// a value which can be stored as a property
type PropertyValue interface{}

// called in Set methods to verify that a property is valid
func verifyPropertyValue(v PropertyValue) PropertyValue {
	switch v.(type) {

	// nothing
	case nil:

	// normal property values
	case Object:
	case LazyEvaluatedValue:
	case ComputedProperty:
	case *weakref.WeakRef:

	// conversion from go builtin types
	case string:
		return Fstring(v)
	case float64, float32,
		int64, int32, int16, int8, int,
		uint64, uint32, uint16, uint8, uint:
		return Fnum(v)
	case bool:
		return Fbool(v)

	default:
		panic("invalid property value")
	}
	return v
}

type Code interface {
	SetLastParent(p Object)
	Signature() *Signature
	Call(c Call) Object
}

// a PropertyValue representing a computed property
type ComputedProperty struct {
	code Code // function or event
	lazy bool // set the value after first evaluation
}

// a PropertyValue representing a lazy-evaluated value.
// unlike ComputedProperty, this does not have to be associated with an actual
// function object
type LazyEvaluatedValue func(object Object, owner Object) Object

// options passed to all Description methods
type DescriptionOption struct {
	ignore map[Object]uint
}
