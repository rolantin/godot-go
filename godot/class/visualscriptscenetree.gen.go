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

//func NewVisualScriptSceneTreeFromPointer(ptr gdnative.Pointer) VisualScriptSceneTree {
func NewVisualScriptSceneTreeFromPointer(ptr gdnative.Pointer) VisualScriptSceneTree {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptSceneTree{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptSceneTree struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptSceneTree) BaseClass() string {
	return "VisualScriptSceneTree"
}