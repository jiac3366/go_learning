package main


//写法1
func add(a int, b int) (int, int) {
	return a + b, a - b
}

//写法2
func add2(a int, b int) (sum int,  res int) {
	sum = a + b
	res = a - b
	return
}