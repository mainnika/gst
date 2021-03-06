package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"github.com/lidouf/glib"
	"unsafe"
	//"fmt"
)

type ElementFactory struct {
	GstObj
}

func (f *ElementFactory) g() *C.GstElementFactory {
	return (*C.GstElementFactory)(f.GetPtr())
}

func (f *ElementFactory) AsElementFactory() *ElementFactory {
	return f
}

// Get the GType for elements managed by this factory. The type can only be retrieved if the element factory is loaded,
// which can be assured with gst_plugin_feature_load().
// Returns
// the GType for elements managed by this factory or 0 if the factory is not loaded.
func (f *ElementFactory) GetElementType() glib.Type {
	return glib.Type(C.gst_element_factory_get_element_type(f.g()))
}

func (f *ElementFactory) Create(name string) *Element {
	n := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(n))

	ge := C.gst_element_factory_create(f.g(), n)
	if ge == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(glib.Pointer(ge))
	return e
}

/**
Get metadata from ElementFactory
*/
func (f *ElementFactory) GetMetadataKeys() []string {
	k := C.gst_element_factory_get_metadata_keys(f.g())
	if k == nil {
		return []string{}
	}
	defer C.free(unsafe.Pointer(k))

	return convertToGoSlice(k, 6)
}

func (f *ElementFactory) GetLongName() string {
	return C.GoString((*C.char)(C.gst_element_factory_get_metadata(f.g(), (*C.gchar)(C.CString(ELEMENT_METADATA_LONGNAME)))))
}

func (f *ElementFactory) GetKlass() string {
	return C.GoString((*C.char)(C.gst_element_factory_get_metadata(f.g(), (*C.gchar)(C.CString(ELEMENT_METADATA_KLASS)))))
}

func (f *ElementFactory) GetDescription() string {
	return C.GoString((*C.char)(C.gst_element_factory_get_metadata(f.g(), (*C.gchar)(C.CString(ELEMENT_METADATA_DESCRIPTION)))))
}

func (f *ElementFactory) GetAuthor() string {
	return C.GoString((*C.char)(C.gst_element_factory_get_metadata(f.g(), (*C.gchar)(C.CString(ELEMENT_METADATA_AUTHOR)))))
}

func (f *ElementFactory) GetDocumentationUri() string {
	return C.GoString((*C.char)(C.gst_element_factory_get_metadata(f.g(), (*C.gchar)(C.CString(ELEMENT_METADATA_DOC_URI)))))
}

func (f *ElementFactory) GetIconName() string {
	return C.GoString((*C.char)(C.gst_element_factory_get_metadata(f.g(), (*C.gchar)(C.CString(ELEMENT_METADATA_ICON_NAME)))))
}

func (f *ElementFactory) GetNumPadTemplates() uint {
	return uint(C.gst_element_factory_get_num_pad_templates(f.g()))
}

func (f *ElementFactory) GetStaticPadTemplates() []*StaticPadTemplate {
	clist := C.gst_element_factory_get_static_pad_templates(f.g())

	padTemplates := make([]*StaticPadTemplate, 0)
	list := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	for ; list != nil && list.Data() != nil; list = list.Next() {
		padTemplates = append(padTemplates, (*StaticPadTemplate)(list.Data().(unsafe.Pointer)))
	}

	return padTemplates
}

func ElementFactoryMake(factory_name, name string) *Element {
	fn := (*C.gchar)(C.CString(factory_name))
	defer C.free(unsafe.Pointer(fn))
	n := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(n))
	ge := C.gst_element_factory_make(fn, n)
	if ge == nil {
		return nil
	}
	e := new(Element)
	e.SetPtr(glib.Pointer(ge))
	return e
}

func ElementFactoryFind(name string) *ElementFactory {
	n := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(n))

	gf := C.gst_element_factory_find(n)
	if gf == nil {
		return nil
	}

	f := new(ElementFactory)
	f.SetPtr(glib.Pointer(gf))

	return f
}
