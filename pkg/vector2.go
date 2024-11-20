package yagml

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X, Y float32
}

var (
	Zero     = Vector2{0, 0}
	One      = Vector2{1, 1}
	Negative = Vector2{-1, -1}
	Up       = Vector2{0, 1}
	Down     = Vector2{0, -1}
	Right    = Vector2{1, 0}
	Left     = Vector2{-1, 0}
)

func (v Vector2) String() string {
	return fmt.Sprintf("vec: %.2f, %.2f", v.X, v.Y)
}

// Return vector with all components is absolute values (positive)
func (v Vector2) Abs() Vector2 {
	return Vector2{float32(math.Abs(float64(v.X))), float32(math.Abs(float64(v.Y)))}
}

// Retrun angle of vector in radian
func (v Vector2) Angle() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}

// Retrun lenght of vector
func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Return squared length, faster then Length
func (v Vector2) LengthSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

// Return normalized vector
func (v Vector2) Normalized() Vector2 {
	length := v.LengthSquared()
	if length != 0 {
		length = float32(math.Sqrt(float64(length)))

		v.X /= length
		v.Y /= length
	}
	return v
}

// Return distance between two vectors
func (v Vector2) DistanceTo(other Vector2) float32 {
	return float32(math.Sqrt(float64((v.X-other.X)*(v.X-other.X) + (v.Y-other.Y)*(v.Y-other.Y))))
}

// Return distance between two vectors
func (v Vector2) DistanceSquaredTo(other Vector2) float32 {
	return (v.X-other.X)*(v.X-other.X) + (v.Y-other.Y)*(v.Y-other.Y)
}

// Retrun dot product of vectors
func (v Vector2) Dot(other Vector2) float32 {
	return v.X*other.X + v.Y*other.Y
}

// Retrun cross product of vectors
func (v Vector2) Cross(other Vector2) float32 {
	return v.X*other.X - v.Y*other.Y
}

func (v Vector2) Sign() Vector2 {
	return Vector2{sign(v.X), sign(v.Y)}
}

func (v Vector2) Floor() Vector2 {
	return Vector2{float32(math.Floor(float64(v.X))), float32(math.Floor(float64(v.Y)))}
}

func (v Vector2) Ceil() Vector2 {
	return Vector2{float32(math.Ceil(float64(v.X))), float32(math.Ceil(float64(v.Y)))}
}

func (v Vector2) Round() Vector2 {
	return Vector2{float32(math.Round(float64(v.X))), float32(math.Round(float64(v.Y)))}
}

// Retrun rotated vector on angle in radian
func (v Vector2) Rotated(angle float32) Vector2 {
	sine := float32(math.Sin(float64(angle)))
	cose := float32(math.Cos(float64(angle)))

	return Vector2{v.X*cose - v.Y*sine, v.X*sine + v.Y*cose}
}

// Returns the result of the linear interpolation
func (v Vector2) Lerp(to Vector2, weight float32) Vector2 {
	return Vector2{lerp(v.X, to.X, weight), lerp(v.Y, to.Y, weight)}
}

// Retruns the component-wise maximum of this and with
func (v Vector2) Max(with Vector2) Vector2 {
	return Vector2{float32(math.Max(float64(v.X), float64(with.X))), float32(math.Max(float64(v.Y), float64(with.Y)))}
}

// Retruns the component-wise minimum of this and with
func (v Vector2) Min(with Vector2) Vector2 {
	return Vector2{float32(math.Min(float64(v.X), float64(with.X))), float32(math.Min(float64(v.Y), float64(with.Y)))}
}

// Retruns the component-wise maximum of this and with
func (v Vector2) Maxf(with float32) Vector2 {
	return Vector2{float32(math.Max(float64(v.X), float64(with))), float32(math.Max(float64(v.Y), float64(with)))}
}

// Retruns the component-wise minimum of this and with
func (v Vector2) Minf(with float32) Vector2 {
	return Vector2{float32(math.Min(float64(v.X), float64(with))), float32(math.Min(float64(v.Y), float64(with)))}
}

// Returns a new vector moved toward to by the fixed delta amount.
// Will not go past the final value
func (v Vector2) MoveToward(to Vector2, delta float32) Vector2 {
	vd := to.Sub(v)
	l := vd.Length()

	if l <= delta || l < 0.00001 {
		return to
	} else {
		dir := vd.DivValue(l).MulValue(delta)
		return v.Add(dir)
	}
}

// Sum two vectors
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}

// Subdivide two vectors
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}

// Multiply two vectors
func (v Vector2) Mul(other Vector2) Vector2 {
	return Vector2{v.X * other.X, v.Y * other.Y}
}

// Divide two vectors
func (v Vector2) Div(other Vector2) Vector2 {
	return Vector2{v.X / other.X, v.Y / other.Y}
}

// Add scalar to vector, return changed vector
func (v Vector2) AddValue(value float32) Vector2 {
	return Vector2{v.X + value, v.Y + value}
}

// Sub scalar to vector, return changed vector
func (v Vector2) SubValue(value float32) Vector2 {
	return Vector2{v.X - value, v.Y - value}
}

// Mul scalar to vector, return changed vector
func (v Vector2) MulValue(value float32) Vector2 {
	return Vector2{v.X * value, v.Y * value}
}

// Div scalar to vector, return changed vector
func (v Vector2) DivValue(value float32) Vector2 {
	return Vector2{v.X / value, v.Y / value}
}

func sign(x float32) float32 {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func lerp(from, to, weight float32) float32 {
	return from*(1-weight) + to*weight
}
