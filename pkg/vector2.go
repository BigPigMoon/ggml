package yagml

import "fmt"

type Vector2 struct {
	X, Y float32
}

func Vector2Zero() Vector2 {
	return Vector2{0, 0}
}

func (v Vector2) String() string {
	return fmt.Sprintf("vec: %.2f, %.2f", v.X, v.Y)
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}
