package dom

import (
	"github.com/killi1812/libxml2/clib"
)

// URI returns the namespace URL
func (n *Namespace) URI() string {
	return clib.XMLNamespaceHref(n)
}

// Prefix returns the prefix for this namespace
func (n *Namespace) Prefix() string {
	return clib.XMLNamespacePrefix(n)
}

// Free releases the underlying C struct
func (n *Namespace) Free() {
	clib.XMLNamespaceFree(n)
	n.ptr = 0
}

// String returns the stringified Namespace
func (n *Namespace) String() string {
	prefix := n.Prefix()
	if prefix != "" {
		prefix = ":" + prefix
	}
	return "xmlns" + prefix + `="` + n.URI() + `"`
}
