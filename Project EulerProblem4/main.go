package main

import "fmt"

func main() {
	var (
		i  int
		j  int
		t1 int
		t2 int
		t3 int
	)
	for i = 100; i < 1000; i++ {
		for j = 100; j < 1000; j++ {
			t1 = i * j
			if reverse_int(t1) == t1 {
				t3=t1
				if t3>t2 {
					t2=t3
				}

			}
		}
	}
	fmt.Println(t2)

}
func reverse_int(n int) int {
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}