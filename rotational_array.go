package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5,6,7}
	n := 2
	fmt.Println(zero(a,n))
}
func zero(a []int,n int) []int {
	var temp1 int
	for j:=1 ; j<=n ;j++ {
		temp1 = a[0]
		for i := 0; i < len(a)-1; i++ {
			a[i] = a[i+1]
		}
		a[len(a)-1] = temp1
	}
	return(a)
}
