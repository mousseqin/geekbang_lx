package main

import "fmt"

func main(){
	//s1 := []int{1,2,3,4}
	//fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	//
	//s2 := make([]int,3,4)
	//fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	//s2 = append(s2,7)
	//fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	//s2 = append(s2,8)
	//fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
	//
	//s3 := make([]int,4)
	//fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))
	//// 按下标索引
	//fmt.Printf("s3[2]: %d", s3[2])

	//SubSlice()
	ShareSlice()


}

func ShareSlice() {

	s1 := []int{1,2,3,4}
	s2 := s1[2:]
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s2[1] = 1999
	fmt.Printf("s1: %v, len %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))
}

func SubSlice() {
	s1 := []int{2,4,6,8,10}
	s2 := s1[1:3]
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4: %v, len %d, cap: %d \n", s4, len(s4), cap(s4))
}

















