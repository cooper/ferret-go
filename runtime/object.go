package runtime

type Object interface {

    // fetch a property and its owner. if this is a computed property
    // or lazy-evaluated value, it is NOT evaluated
    Property(name string) (owner Object, value PropertyValue)

    // fetch a property and its own, always yielding an Object by
    // evaluating computed properties
    PropertyComputed(name string) (owner Object, value Object)

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
    // consideration inherited properties. return the deleted object
    Delete(name string) Object
    DeleteOverwrite(name string) Object

    // fetch and evaluate or set the value at the given index
    GetIndex(index Object) Object
    SetIndex(index Object, value Object)

    // call the object with the given call info, returning an object
    Call(c Call) Object

    // return a string description of the object
    Description() string
}

type PropertyValue interface{}

type Call struct {

}
