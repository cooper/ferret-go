package runtime

var valuesByIdentifier = make(map[uint]PropertyValue)
var identifiersByValue = make(map[PropertyValue]uint)
var refcount = make(map[uint]uint)
var currentIdentifier uint

func increase(value PropertyValue) uint {
	id, ok := identifiersByValue[value]
	if !ok {
		id = currentIdentifier
		currentIdentifier++
		valuesByIdentifier[id] = value
		identifiersByValue[value] = id
	}
	refcount[id]++
	return id
}

func decrease(value PropertyValue) bool {
	id := identifiersByValue[value]
	refcount[id]--
	if refcount[id] <= 0 {
		delete(identifiersByValue, value)
		delete(valuesByIdentifier, id)
		delete(refcount, id)
		return true
	}
	return false
}

func retrieve(id uint) PropertyValue {
	val, ok := valuesByIdentifier[id]
	if !ok {
		return nil
	}
	return val
}
