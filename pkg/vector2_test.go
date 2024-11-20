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

	assertVector2Equal(t, got, wish)
}

func TestVector2Multiply(t *testing.T) {
	vec1 := Vector2{0, 2}
	vec2 := Vector2{2, 0}

	got := vec1.Mul(vec2)
	wish := Vector2{0, 0}

	assertVector2Equal(t, got, wish)
}

func TestVector2Substract(t *testing.T) {
	vec1 := Vector2{2, 2}
	vec2 := Vector2{2, 2}

	got := vec1.Sub(vec2)
	wish := Vector2{0, 0}

	assertVector2Equal(t, got, wish)
}

func TestVector2Divide(t *testing.T) {
	vec1 := Vector2{2, 2}
	vec2 := Vector2{2, 2}

	got := vec1.Div(vec2)
	wish := Vector2{1, 1}

	assertVector2Equal(t, got, wish)
}

func assertVector2Equal(t testing.TB, got, wish Vector2) {
	t.Helper()

	if got.X != wish.X && got.Y != wish.Y {
		t.Errorf("expect vec %s but got %s", Vector2{2, 2}, got)
	}
}
