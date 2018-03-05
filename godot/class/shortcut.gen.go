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

//func NewShortCutFromPointer(ptr gdnative.Pointer) ShortCut {
func NewShortCutFromPointer(ptr gdnative.Pointer) ShortCut {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ShortCut{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A shortcut for binding input. Shortcuts are commonly used for interacting with a [Control] element from a [InputEvent].
*/
type ShortCut struct {
	Resource
	owner gdnative.Object
}

func (o *ShortCut) BaseClass() string {
	return "ShortCut"
}

/*
        Returns the Shortcut's [InputEvent] as a [String].
	Args: [], Returns: String
*/
func (o *ShortCut) GetAsText() gdnative.String {
	//log.Println("Calling ShortCut.GetAsText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShortCut", "get_as_text")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: InputEvent
*/
func (o *ShortCut) GetShortcut() InputEvent {
	//log.Println("Calling ShortCut.GetShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShortCut", "get_shortcut")

	// Call the parent method.
	// InputEvent
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewInputEventFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns [code]true[/code] if the Shortcut's [InputEvent] equals [code]event[/code].
	Args: [{ false event InputEvent}], Returns: bool
*/
func (o *ShortCut) IsShortcut(event InputEvent) gdnative.Bool {
	//log.Println("Calling ShortCut.IsShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShortCut", "is_shortcut")

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
        If [code]true[/code] this Shortcut is valid.
	Args: [], Returns: bool
*/
func (o *ShortCut) IsValid() gdnative.Bool {
	//log.Println("Calling ShortCut.IsValid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShortCut", "is_valid")

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
	Args: [{ false event InputEvent}], Returns: void
*/
func (o *ShortCut) SetShortcut(event InputEvent) {
	//log.Println("Calling ShortCut.SetShortcut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShortCut", "set_shortcut")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}