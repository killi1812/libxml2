package dom

// Auto-generated by internal/cmd/genwrapnode/genwrapnode.go. DO NOT EDIT!

import (
	"fmt"

	"github.com/killi1812/libxml2/clib"
	"github.com/killi1812/libxml2/types"
)

func wrapNamespaceNode(ptr uintptr) *Namespace {
	var n Namespace
	n.ptr = ptr
	return &n
}

func wrapAttributeNode(ptr uintptr) *Attribute {
	var n Attribute
	n.ptr = ptr
	return &n
}

func wrapCDataSectionNode(ptr uintptr) *CDataSection {
	var n CDataSection
	n.ptr = ptr
	return &n
}

func wrapCommentNode(ptr uintptr) *Comment {
	var n Comment
	n.ptr = ptr
	return &n
}

func wrapElementNode(ptr uintptr) *Element {
	var n Element
	n.ptr = ptr
	return &n
}

func wrapTextNode(ptr uintptr) *Text {
	var n Text
	n.ptr = ptr
	return &n
}

func wrapPiNode(ptr uintptr) *Pi {
	var n Pi
	n.ptr = ptr
	return &n
}

// WrapNode is a function created with the sole purpose of allowing
// go-libxml2 consumers that can generate a C.xmlNode pointer to
// create libxml2.Node types, e.g. go-xmlsec.
func WrapNode(n uintptr) (types.Node, error) {
	switch typ := clib.XMLGetNodeTypeRaw(n); typ {
	case clib.AttributeNode:
		return wrapAttributeNode(n), nil
	case clib.CDataSectionNode:
		return wrapCDataSectionNode(n), nil
	case clib.CommentNode:
		return wrapCommentNode(n), nil
	case clib.ElementNode:
		return wrapElementNode(n), nil
	case clib.TextNode:
		return wrapTextNode(n), nil
	case clib.PiNode:
		return wrapPiNode(n), nil
	default:
		return nil, fmt.Errorf("unknown node: %d", typ)
	}
}
