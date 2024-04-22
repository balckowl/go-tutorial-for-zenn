package main

import (
	"fmt"
)

const (
	PI = 3.14
)

func familyName(fname string, age int){
	fmt.Println(fname, age)
}

func main() {

	var student1 string = "Jhon"
	var student2 = "Jane"
	x := 2

	//コメント〜
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	r := 5.0

	fmt.Println(r * r * PI)

	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}
	arr3 := [10]int{1: 3}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(len(arr1))
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	fmt.Println(cap(arr3))

	// 元となる配列を作成
	originalArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// スライスを作成 (最初の3要素だけを参照)
	sliceOfArray := originalArray[:3]

	// スライスの長さと容量を出力
	fmt.Println("Length of slice:", len(sliceOfArray))   // 3
	fmt.Println("Capacity of slice:", cap(sliceOfArray)) // 10

	myslice1 := make([]int, 5, 10)

	fmt.Println(myslice1)

	k := 10
	m := 100

	if k > m {
		fmt.Println("Hello")
	} else {
		fmt.Println("Not")
	}

	nn := "hello"

	switch nn {
	case "hello":
		fmt.Println("Yes")
	default:
		fmt.Println("yes")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fruits := [3]string{"apple", "banana", "orange"}
	

	for index, fruit := range fruits{
		fmt.Println(index, fruit)
	}

	familyName("Jany", 18)
}
