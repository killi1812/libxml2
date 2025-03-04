// Package xsd contains some of the tools available from libxml2
// that allows you to validate your XML against an XSD
//
// This is basically all you need to do:
//
//	schema, err := relaxng.Parse(xsdsrc)
//	if err != nil {
//	    panic(err)
//	}
//	defer schema.Free()
//	if err := schema.Validate(doc); err != nil{
//	    for _, e := range err.(SchemaValidationErr).Error() {
//	         println(e.Error())
//	    }
//	}
package relaxng

import (
	"github.com/killi1812/libxml2/clib"
	"github.com/killi1812/libxml2/types"
	"github.com/pkg/errors"
)

const ValueVCCreate = 1

// Parse is used to parse an XML Schema Document to produce a
// Schema instance. Make sure to call Free() on the instance
// when you are done with it.
// ==IMPORTANT== still in development
// TODO: finish transfering to relaxNG
func Parse(buf []byte, options ...types.Option) (types.Schema, error) {
	sptr, err := clib.XMLSchemaParse(buf, options...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse input")
	}

	return &relaxNGSchema{ptr: sptr}, nil
}

// ParseFromFile is used to parse an XML schema using only the file path.
// Make sure to call Free() on the instance when you are done with it.
func ParseFromFile(path string) (types.Schema, error) {
	sptr, err := clib.XMLRelaxNGSchemaParseFromFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse input from file")
	}

	return &relaxNGSchema{ptr: sptr}, nil
}

// Pointer returns the underlying C struct
func (s *relaxNGSchema) Pointer() uintptr {
	return s.ptr
}

// Free frees the underlying C struct
func (s *relaxNGSchema) Free() {
	if err := clib.XMLRelaxNGFree(s); err != nil {
		return
	}
	s.ptr = 0
}

// Validate takes in a XML document and validates it against
// the schema. If there are any problems, and error is
// returned.
func (s *relaxNGSchema) Validate(d types.Document, options ...int) error {
	errs := clib.XMLRelaxNGValidateDocument(s, d, options...)
	if errs == nil {
		return nil
	}

	return types.SchemaValidationError{Errors: errs}
}
