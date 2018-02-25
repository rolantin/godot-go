package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include "gdnative.gen.h"
#include <gdnative/variant.h>
*/
import "C"

type Variant struct {
	base *C.godot_variant
}

func (gdt Variant) getBase() *C.godot_variant {
	return gdt.base
}

// AsBool godot_variant_as_bool [[const godot_variant * p_self]] godot_bool
func (gdt *Variant) AsBool() Bool {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_bool(GDNative.api, arg0)

	return Bool(ret)
}

// AsUint godot_variant_as_uint [[const godot_variant * p_self]] uint64_t
func (gdt *Variant) AsUint() Uint64T {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_uint(GDNative.api, arg0)

	return Uint64T(ret)
}

// AsInt godot_variant_as_int [[const godot_variant * p_self]] int64_t
func (gdt *Variant) AsInt() Int64T {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_int(GDNative.api, arg0)

	return Int64T{base: ret}
}

// AsReal godot_variant_as_real [[const godot_variant * p_self]] double
func (gdt *Variant) AsReal() Double {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_real(GDNative.api, arg0)

	return Double{base: ret}
}

// AsString godot_variant_as_string [[const godot_variant * p_self]] godot_string
func (gdt *Variant) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_string(GDNative.api, arg0)

	return String{base: ret}
}

// AsVector2 godot_variant_as_vector2 [[const godot_variant * p_self]] godot_vector2
func (gdt *Variant) AsVector2() Vector2 {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_vector2(GDNative.api, arg0)

	return Vector2{base: ret}
}

// AsRect2 godot_variant_as_rect2 [[const godot_variant * p_self]] godot_rect2
func (gdt *Variant) AsRect2() Rect2 {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_rect2(GDNative.api, arg0)

	return Rect2{base: ret}
}

// AsVector3 godot_variant_as_vector3 [[const godot_variant * p_self]] godot_vector3
func (gdt *Variant) AsVector3() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_vector3(GDNative.api, arg0)

	return Vector3{base: ret}
}

// AsTransform2D godot_variant_as_transform2d [[const godot_variant * p_self]] godot_transform2d
func (gdt *Variant) AsTransform2D() Transform2D {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_transform2d(GDNative.api, arg0)

	return Transform2D{base: ret}
}

// AsPlane godot_variant_as_plane [[const godot_variant * p_self]] godot_plane
func (gdt *Variant) AsPlane() Plane {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_plane(GDNative.api, arg0)

	return Plane{base: ret}
}

// AsQuat godot_variant_as_quat [[const godot_variant * p_self]] godot_quat
func (gdt *Variant) AsQuat() Quat {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_quat(GDNative.api, arg0)

	return Quat{base: ret}
}

// AsAabb godot_variant_as_aabb [[const godot_variant * p_self]] godot_aabb
func (gdt *Variant) AsAabb() Aabb {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_aabb(GDNative.api, arg0)

	return Aabb{base: ret}
}

// AsBasis godot_variant_as_basis [[const godot_variant * p_self]] godot_basis
func (gdt *Variant) AsBasis() Basis {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_basis(GDNative.api, arg0)

	return Basis{base: ret}
}

// AsTransform godot_variant_as_transform [[const godot_variant * p_self]] godot_transform
func (gdt *Variant) AsTransform() Transform {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_transform(GDNative.api, arg0)

	return Transform{base: ret}
}

// AsColor godot_variant_as_color [[const godot_variant * p_self]] godot_color
func (gdt *Variant) AsColor() Color {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_color(GDNative.api, arg0)

	return Color{base: ret}
}

// AsNodePath godot_variant_as_node_path [[const godot_variant * p_self]] godot_node_path
func (gdt *Variant) AsNodePath() NodePath {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_node_path(GDNative.api, arg0)

	return NodePath{base: ret}
}

// AsRid godot_variant_as_rid [[const godot_variant * p_self]] godot_rid
func (gdt *Variant) AsRid() Rid {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_rid(GDNative.api, arg0)

	return Rid{base: ret}
}

// AsObject godot_variant_as_object [[const godot_variant * p_self]] godot_object *
func (gdt *Variant) AsObject() Object {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_object(GDNative.api, arg0)

	return Object{base: ret}
}

// AsDictionary godot_variant_as_dictionary [[const godot_variant * p_self]] godot_dictionary
func (gdt *Variant) AsDictionary() Dictionary {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_dictionary(GDNative.api, arg0)

	return Dictionary{base: ret}
}

// AsArray godot_variant_as_array [[const godot_variant * p_self]] godot_array
func (gdt *Variant) AsArray() Array {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_array(GDNative.api, arg0)

	return Array{base: ret}
}

// AsPoolByteArray godot_variant_as_pool_byte_array [[const godot_variant * p_self]] godot_pool_byte_array
func (gdt *Variant) AsPoolByteArray() PoolByteArray {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_pool_byte_array(GDNative.api, arg0)

	return PoolByteArray{base: ret}
}

// AsPoolIntArray godot_variant_as_pool_int_array [[const godot_variant * p_self]] godot_pool_int_array
func (gdt *Variant) AsPoolIntArray() PoolIntArray {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_pool_int_array(GDNative.api, arg0)

	return PoolIntArray{base: ret}
}

// AsPoolRealArray godot_variant_as_pool_real_array [[const godot_variant * p_self]] godot_pool_real_array
func (gdt *Variant) AsPoolRealArray() PoolRealArray {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_pool_real_array(GDNative.api, arg0)

	return PoolRealArray{base: ret}
}

// AsPoolStringArray godot_variant_as_pool_string_array [[const godot_variant * p_self]] godot_pool_string_array
func (gdt *Variant) AsPoolStringArray() PoolStringArray {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_pool_string_array(GDNative.api, arg0)

	return PoolStringArray{base: ret}
}

// AsPoolVector2Array godot_variant_as_pool_vector2_array [[const godot_variant * p_self]] godot_pool_vector2_array
func (gdt *Variant) AsPoolVector2Array() PoolVector2Array {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_pool_vector2_array(GDNative.api, arg0)

	return PoolVector2Array{base: ret}
}

// AsPoolVector3Array godot_variant_as_pool_vector3_array [[const godot_variant * p_self]] godot_pool_vector3_array
func (gdt *Variant) AsPoolVector3Array() PoolVector3Array {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_pool_vector3_array(GDNative.api, arg0)

	return PoolVector3Array{base: ret}
}

// AsPoolColorArray godot_variant_as_pool_color_array [[const godot_variant * p_self]] godot_pool_color_array
func (gdt *Variant) AsPoolColorArray() PoolColorArray {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_as_pool_color_array(GDNative.api, arg0)

	return PoolColorArray{base: ret}
}

// Call godot_variant_call [[godot_variant * p_self] [const godot_string * p_method] [const godot_variant ** p_args] [const godot_int p_argcount] [godot_variant_call_error * r_error]] godot_variant
func (gdt *Variant) Call(method String, args []Variant, argcount Int, error VariantCallError) Variant {
	arg0 := gdt.getBase()
	arg1 := method.getBase()
	arg2 := args.getBase()
	arg3 := argcount.getBase()
	arg4 := error.getBase()

	ret := C.go_godot_variant_call(GDNative.api, arg0, arg1, arg2, arg3, arg4)

	return Variant{base: ret}
}

// HasMethod godot_variant_has_method [[const godot_variant * p_self] [const godot_string * p_method]] godot_bool
func (gdt *Variant) HasMethod(method String) Bool {
	arg0 := gdt.getBase()
	arg1 := method.getBase()

	ret := C.go_godot_variant_has_method(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorEqual godot_variant_operator_equal [[const godot_variant * p_self] [const godot_variant * p_other]] godot_bool
func (gdt *Variant) OperatorEqual(other Variant) Bool {
	arg0 := gdt.getBase()
	arg1 := other.getBase()

	ret := C.go_godot_variant_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorLess godot_variant_operator_less [[const godot_variant * p_self] [const godot_variant * p_other]] godot_bool
func (gdt *Variant) OperatorLess(other Variant) Bool {
	arg0 := gdt.getBase()
	arg1 := other.getBase()

	ret := C.go_godot_variant_operator_less(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// HashCompare godot_variant_hash_compare [[const godot_variant * p_self] [const godot_variant * p_other]] godot_bool
func (gdt *Variant) HashCompare(other Variant) Bool {
	arg0 := gdt.getBase()
	arg1 := other.getBase()

	ret := C.go_godot_variant_hash_compare(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// Booleanize godot_variant_booleanize [[const godot_variant * p_self]] godot_bool
func (gdt *Variant) Booleanize() Bool {
	arg0 := gdt.getBase()

	ret := C.go_godot_variant_booleanize(GDNative.api, arg0)

	return Bool(ret)
}

// Destroy godot_variant_destroy [[godot_variant * p_self]] void
func (gdt *Variant) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_variant_destroy(GDNative.api, arg0)
}

type VariantType int

func (e VariantType) getBase() C.godot_variant_type {
	return C.godot_variant_type(e)
}

const (
	VariantTypeNil VariantType = iota

	VariantTypeBool
	VariantTypeInt
	VariantTypeReal
	VariantTypeString

	VariantTypeVector2
	VariantTypeRect2
	VariantTypeVector3
	VariantTypeTransform2D
	VariantTypePlane
	VariantTypeQuat
	VariantTypeAabb
	VariantTypeBasis
	VariantTypeTransform

	VariantTypeColor
	VariantTypeNodePath
	VariantTypeRid
	VariantTypeObject
	VariantTypeDictionary
	VariantTypeArray

	VariantTypePoolByteArray
	VariantTypePoolIntArray
	VariantTypePoolRealArray
	VariantTypePoolStringArray
	VariantTypePoolVector2Array
	VariantTypePoolVector3Array
	VariantTypePoolColorArray
)

type VariantCallErrorError int

func (e VariantCallErrorError) getBase() C.godot_variant_call_error_error {
	return C.godot_variant_call_error_error(e)
}

const (
	CallErrorCallOk VariantCallErrorError = iota
	CallErrorCallErrorInvalidMethod
	CallErrorCallErrorInvalidArgument
	CallErrorCallErrorTooManyArguments
	CallErrorCallErrorTooFewArguments
	CallErrorCallErrorInstanceIsNull
)

type VariantCallError struct {
	base *C.godot_variant_call_error

	Error    VariantCallErrorError
	Argument Int
	Expected VariantType
}

func (gdt VariantCallError) getBase() *C.godot_variant_call_error {
	return gdt.base
}