package main

import "fmt"

func Create() {
	s1 := make([]int, 0)
	s2 := make([]int, 1)
	s3 := make([]int, 1, 3)
	s4 := make([]int, 3, 5)
	s5 := []int{}
	s6 := []int{1, 2, 3}
	s7 := []int{1: 2, 3}
	s8 := s6[1:2]
	s9 := s6[:2]
	s15 := s6[:2:2]
	s10 := s6[1:]
	s11 := s6[:]
	s12 := s6[3:] // 可以使用切片s6的长度3来生成新切片，超过将会报错误
	// s12 := s3[3:] // compile error
	// s12 := s4[5:] // compile error
	// s12 := s6[5:] // compile error
	s13 := s6[1:2:2]
	s14 := s6[1:2:3]

	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice1", s1, len(s1), cap(s1))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice2", s2, len(s2), cap(s2))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice3", s3, len(s3), cap(s3))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice4", s4, len(s4), cap(s4))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice5", s5, len(s5), cap(s5))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice6", s6, len(s6), cap(s6))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice7", s7, len(s7), cap(s7))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice8", s8, len(s8), cap(s8))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice9", s9, len(s9), cap(s9))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice10", s10, len(s10), cap(s10))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice11", s11, len(s11), cap(s11))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice12", s12, len(s12), cap(s12))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice13", s13, len(s13), cap(s13))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice14", s14, len(s14), cap(s14))
	fmt.Printf("%s = %v,\t len = %d, cap = %d\n", "slice15", s15, len(s15), cap(s15))
}

func main() {
	Create()
}
