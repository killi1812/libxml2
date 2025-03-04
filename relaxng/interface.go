package relaxng

import "github.com/killi1812/libxml2/internal/option"

// RelaxNGSchema represents an XML schema.
type RelaxNGSchema struct {
	ptr uintptr // *C.xmlSchema
}

// SchemaValidationError is returned when the Validate() function
// finds errors. When there are multiple errors, you may access
// them using the Errors() method
type SchemaValidationError struct {
	errors []error
}

type Option = option.Interface
