package main

func add(a, b int) int {
	return a + b
}

func scaleAndAdd(x int) int {
	return add(x*2, 5)
}

func main() {
	_ = scaleAndAdd(21)
}
