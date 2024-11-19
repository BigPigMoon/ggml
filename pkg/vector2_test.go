package yagml

import "testing"

func TestVector2Zero(t *testing.T) {
	vec := Vector2Zero()

	if vec.X != 0 && vec.Y != 0 {
		t.Errorf("expect vec %s but got %s", Vector2{0, 0}, vec)
	}
}

func TestVector2Add(t *testing.T) {
	vec1 := Vector2{0, 2}
	vec2 := Vector2{2, 0}

	got := vec1.Add(vec2)
	wish := Vector2{2, 2}

	if got.X != wish.X && got.Y != wish.Y {
		t.Errorf("expect vec %s but got %s", Vector2{2, 2}, got)
	}
}
