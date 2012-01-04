package gform

import (
    "w32"
)

type Controller interface {
    Caption() string
    Enabled() bool
    Focus()
    Handle() w32.HWND
    Invalidate(erase bool)
    Parent() Controller
    Pos() (x, y int)
    Size() (w, h int)
    Visible() bool
    Bounds() *Rect
    ClientRect() *Rect
    SetCaption(s string)
    SetEnabled(b bool)
    SetPos(x, y int)
    SetSize(w, h int)
    SetDragAcceptFilesEnabled(b bool)
    Show()
    Hide()
    Font() *Font
    SetFont(font *Font)
    InvokeRequired() bool
    PreTranslateMessage(msg *w32.MSG) bool
    WndProc(msg uint, wparam, lparam uintptr) uintptr

    //General events
    OnCreate() *EventManager

    // Focus events
    OnKillFocus() *EventManager
    OnSetFocus() *EventManager

    //Drag and drop events
    OnDropFiles() *EventManager

    //Mouse events
    OnLBDown() *EventManager
    OnLBUp() *EventManager
    OnMBDown() *EventManager
    OnMBUp() *EventManager
    OnRBDown() *EventManager
    OnRBUp() *EventManager

    OnMouseHover() *EventManager
    OnMouseLeave() *EventManager

    //Keyboard events
    OnKeyUp() *EventManager

    //Paint events
    OnPaint() *EventManager
}
