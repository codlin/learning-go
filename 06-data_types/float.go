package main

import (
	"fmt"
	"math"
)

/*
GO语言提供了两种精度的浮点数：float32和float64
常量math.MaxFloat32: ~3.4e38，math.SmallestNonzeroFloat32: ~1.4e-45
常量math.MaxFloat64: ~1.8e308，math.SmallestNonzeroFloat64: ~4.9e-324
*/

func float_maxmin() {
	fmt.Printf("float32 max: %e, min: %e\n", math.MaxFloat32, math.SmallestNonzeroFloat32)
	fmt.Printf("float32 max: %e, min: %e\n", math.MaxFloat64, math.SmallestNonzeroFloat64)
}

/* float32类型可以提供大约6个十进制数的精度（即小数点后6位）*/
func float32_f() {
	const e float32 = 3.1415926
	fmt.Printf("float32: %v, %.7f, %g\n", e, e, e)
	const r float32 = 31415926.
	fmt.Printf("float32: %v, %.7f, %g\n", r, r, r)
}

/* float64类型可以提供大约15个十进制数的精度（即小数点后15位）*/
func float64_f() {
	const e float64 = 3.1415926141592614159261415926
	fmt.Printf("float32: %v, %.17f, %g\n", e, e, e)
	const r float64 = 31415926141592614159261415926.
	fmt.Printf("float32: %v, %.17f, %g\n", r, r, r)
}

func main() {
	float_maxmin()
	float32_f()
	float64_f()
}
