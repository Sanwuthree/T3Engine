package T3Engine

import "fmt"

type IDisplayObject interface {
	SetVisible(NewVisible bool)
	RemoveFromParent()
}
type IDisplayContainer interface {
	AddChild()
	RemoveChild()
}
type TObject struct {
}
type TActor struct {
	TObject
}

func (a *TActor) BeginPlay() {
	fmt.Println("Begin Play")
}

type TPawn struct {
	TActor
}

type TCharacter struct {
	TPawn
}
