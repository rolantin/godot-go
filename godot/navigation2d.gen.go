package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewNavigation2DFromPointer(ptr gdnative.Pointer) Navigation2D {
func newNavigation2DFromPointer(ptr gdnative.Pointer) Navigation2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Navigation2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type Navigation2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Navigation2D) BaseClass() string {
	return "Navigation2D"
}

/*

	Args: [{ false to_point Vector2}], Returns: Vector2
*/
func (o *Navigation2D) GetClosestPoint(toPoint gdnative.Vector2) gdnative.Vector2 {
	//log.Println("Calling Navigation2D.GetClosestPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(toPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation2D", "get_closest_point")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false to_point Vector2}], Returns: Object
*/
func (o *Navigation2D) GetClosestPointOwner(toPoint gdnative.Vector2) ObjectImplementer {
	//log.Println("Calling Navigation2D.GetClosestPointOwner()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(toPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation2D", "get_closest_point_owner")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*

	Args: [{ false start Vector2} { false end Vector2} {True true optimize bool}], Returns: PoolVector2Array
*/
func (o *Navigation2D) GetSimplePath(start gdnative.Vector2, end gdnative.Vector2, optimize gdnative.Bool) gdnative.PoolVector2Array {
	//log.Println("Calling Navigation2D.GetSimplePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromVector2(start)
	ptrArguments[1] = gdnative.NewPointerFromVector2(end)
	ptrArguments[2] = gdnative.NewPointerFromBool(optimize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation2D", "get_simple_path")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false mesh NavigationPolygon} { false xform Transform2D} {Null true owner Object}], Returns: int
*/
func (o *Navigation2D) NavpolyAdd(mesh NavigationPolygonImplementer, xform gdnative.Transform2D, owner ObjectImplementer) gdnative.Int {
	//log.Println("Calling Navigation2D.NavpolyAdd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(mesh.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromTransform2D(xform)
	ptrArguments[2] = gdnative.NewPointerFromObject(owner.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation2D", "navpoly_add")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false id int}], Returns: void
*/
func (o *Navigation2D) NavpolyRemove(id gdnative.Int) {
	//log.Println("Calling Navigation2D.NavpolyRemove()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation2D", "navpoly_remove")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false id int} { false xform Transform2D}], Returns: void
*/
func (o *Navigation2D) NavpolySetTransform(id gdnative.Int, xform gdnative.Transform2D) {
	//log.Println("Calling Navigation2D.NavpolySetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromTransform2D(xform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation2D", "navpoly_set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Navigation2DImplementer is an interface that implements the methods
// of the Navigation2D class.
type Navigation2DImplementer interface {
	Node2DImplementer
	GetClosestPoint(toPoint gdnative.Vector2) gdnative.Vector2
	GetClosestPointOwner(toPoint gdnative.Vector2) ObjectImplementer
	GetSimplePath(start gdnative.Vector2, end gdnative.Vector2, optimize gdnative.Bool) gdnative.PoolVector2Array
	NavpolyAdd(mesh NavigationPolygonImplementer, xform gdnative.Transform2D, owner ObjectImplementer) gdnative.Int
	NavpolyRemove(id gdnative.Int)
	NavpolySetTransform(id gdnative.Int, xform gdnative.Transform2D)
}
