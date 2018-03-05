package class

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

//func NewJoint2DFromPointer(ptr gdnative.Pointer) Joint2D {
func NewJoint2DFromPointer(ptr gdnative.Pointer) Joint2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Joint2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base node for all joint constraints in 2D physics. Joints take 2 bodies and apply a custom constraint.
*/
type Joint2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Joint2D) BaseClass() string {
	return "Joint2D"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Joint2D) GetBias() gdnative.Float {
	//log.Println("Calling Joint2D.GetBias()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "get_bias")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Joint2D) GetExcludeNodesFromCollision() gdnative.Bool {
	//log.Println("Calling Joint2D.GetExcludeNodesFromCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "get_exclude_nodes_from_collision")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *Joint2D) GetNodeA() gdnative.NodePath {
	//log.Println("Calling Joint2D.GetNodeA()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "get_node_a")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *Joint2D) GetNodeB() gdnative.NodePath {
	//log.Println("Calling Joint2D.GetNodeB()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "get_node_b")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false bias float}], Returns: void
*/
func (o *Joint2D) SetBias(bias gdnative.Float) {
	//log.Println("Calling Joint2D.SetBias()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(bias)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "set_bias")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Joint2D) SetExcludeNodesFromCollision(enable gdnative.Bool) {
	//log.Println("Calling Joint2D.SetExcludeNodesFromCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "set_exclude_nodes_from_collision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false node NodePath}], Returns: void
*/
func (o *Joint2D) SetNodeA(node gdnative.NodePath) {
	//log.Println("Calling Joint2D.SetNodeA()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(node)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "set_node_a")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false node NodePath}], Returns: void
*/
func (o *Joint2D) SetNodeB(node gdnative.NodePath) {
	//log.Println("Calling Joint2D.SetNodeB()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(node)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Joint2D", "set_node_b")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}