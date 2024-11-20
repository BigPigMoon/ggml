package ggml

import (
	"math"
	"testing"
)

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

func TestVector2Length(t *testing.T) {
	vec := Vector2{3, 4}

	got := vec.Length()
	wish := 5.0

	if got != float32(wish) {
		t.Errorf("expect vec %.2f but got %.2f", got, wish)
	}
}

func TestNegativeVector2Length(t *testing.T) {
	vec := Vector2{-3, -4}

	got := vec.Length()
	wish := 5.0

	if got != float32(wish) {
		t.Errorf("expect vec %.2f but got %.2f", got, wish)
	}
}

func TestVector2LengthSquared(t *testing.T) {
	vec := Vector2{3, 4}

	got := vec.LengthSquared()
	wish := 25.0

	if got != float32(wish) {
		t.Errorf("expect vec %.2f but got %.2f", got, wish)
	}
}

func TestVector2Normalized(t *testing.T) {
	vec := Vector2{0, 2}

	got := vec.Normalized()
	wish := Vector2{0, 1}

	assertVector2Equal(t, got, wish)
}

func TestZeroVector2Normalized(t *testing.T) {
	vec := Vector2{0, 0}

	got := vec.Normalized()
	wish := Vector2{0, 0}

	assertVector2Equal(t, got, wish)
}

func TestVector2Angle(t *testing.T) {
	vec := Right

	got := vec.Angle()
	wish := float32(0.0)

	if got != wish {
		t.Errorf("expect vec %.2f but got %.2f", got, wish)
	}
}

func TestUpVector2Angle(t *testing.T) {
	vec := Up

	got := vec.Angle()
	wish := float32(math.Pi / 2.0)

	if got != wish {
		t.Errorf("expect vec %.2f but got %.2f", got, wish)
	}
}

func TestVector2MoveToward(t *testing.T) {
	vec := Right

	got := vec.MoveToward(Right, 10)
	wish := Vector2{11, 0}

	assertVector2Equal(t, got, wish)
}

func TestUpVector2MoveToward(t *testing.T) {
	vec := Zero

	got := vec.MoveToward(Vector2{10, 10}, 15)
	wish := Vector2{10, 10}

	assertVector2Equal(t, got, wish)
}

func assertVector2Equal(t testing.TB, got, wish Vector2) {
	t.Helper()

	if got.X != wish.X && got.Y != wish.Y {
		t.Errorf("expect '%s', but got '%s'", wish, got)
	}
}
