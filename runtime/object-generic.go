package runtime

type genericObject struct {

}

// fetch a property and its owner. if this is a computed property
// or lazy-evaluated value, it is NOT evaluated
func (obj *genericObject) Property(name string) (Object, PropertyValue) {
    return nil, nil
}

// fetch a property and its own, always yielding an Object by
// evaluating computed properties
func (obj *genericObject) PropertyComputed(name string) (Object, Object) {
    return nil, nil
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
    owner, val := obj.PropertyComputed(name)
    if owner != obj {
        return nil
    }
    return val
}

// write the given value to the property by the given name
func (obj *genericObject) Set(name string, value PropertyValue) {

}

// write the given value to the proprerty by the given name, overwriting an
// existing value on a parent object
func (obj *genericObject) SetOverwrite(name string, value PropertyValue) {
    owner, _ := obj.Property(name)
    if owner != nil {
        owner.Set(name, value)
    }
}

// delete the property by the given name. return the deleted object
func (obj *genericObject) Delete(name string) Object {
    return nil
}

// delete the property by the given name, even if it is inherited.
// return the deleted object
func (obj *genericObject) DeleteOverwrite(name string) Object {
    owner, _ := obj.Property(name)
    if owner == nil {
        return nil
    }
    return owner.Delete(name)
}

// fetch and evaluate the value at the given index
func (obj *genericObject) GetIndex(index Object) Object {
    return nil
}

// set the value at the given index
func (obj *genericObject) SetIndex(index Object, value Object) {
    
}

// call the object with the given call info, returning an object
func (obj *genericObject) Call(c Call) Object {
    return nil
}

// return a string description of the object
func (obj *genericObject) Description() string {
    return ""
}
