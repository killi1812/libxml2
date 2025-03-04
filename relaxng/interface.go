package relaxng

// relaxNGSchema represents a RelaxNG XML schema.
type relaxNGSchema struct {
	ptr uintptr // *C.xmlSchema
}
