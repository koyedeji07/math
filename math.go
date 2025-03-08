// Package math can be used to perform mathematical operations.
// supported operations:
//   - Additions - [math.Add]
//   - Multiplications - [math.Mul]
//   - Subtractions - [math.Mul]
package math

// Add A method to add numbers together,
// more information about method operation can be found here:  [mathsisfun]
//
// [mathsisfun]: https://www.mathsisfun.com/numbers/addition.html
func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Mul(x, y int) int {
	return x * y
}
