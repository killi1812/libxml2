package dom

import (
	"github.com/killi1812/libxml2/clib"
)

func (n *CDataSection) Literal() (string, error) {
	return clib.XMLNodeValue(n)
}

// Data returns the content associated with this node
func (n *Text) Data() string {
	return clib.XMLTextData(n)
}
