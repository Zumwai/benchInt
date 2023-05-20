package main

func funnel64_1() {
	var steps int = 256 * 1024 * 1024
	var arr []int64 = []int64{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[0]++
		arr[2]++
		arr[2]++
	}
}

func funnel64_2() {
	var steps int = 256 * 1024 * 1024
	var arr []int64 = []int64{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[1]++
		arr[2]++
		arr[3]++
	}
}

func funnel32_1() {
	var steps int = 256 * 1024 * 1024
	var arr []int32 = []int32{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[0]++
		arr[2]++
		arr[2]++
	}
}

func funnel32_2() {
	var steps int = 256 * 1024 * 1024
	var arr []int32 = []int32{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[1]++
		arr[2]++
		arr[3]++
	}
}

func funnel16_1() {
	var steps int = 256 * 1024 * 1024
	var arr []int16 = []int16{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[0]++
		arr[2]++
		arr[2]++
	}
}

func funnel16_2() {
	var steps int = 256 * 1024 * 1024
	var arr []int16 = []int16{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[1]++
		arr[2]++
		arr[3]++
	}
}

func funnel8_1() {
	var steps int = 256 * 1024 * 1024
	var arr []int8 = []int8{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[0]++
		arr[2]++
		arr[2]++
	}
}

func funnel8_2() {
	var steps int = 256 * 1024 * 1024
	var arr []int8 = []int8{0, 0, 0, 0}
	for i := 0; i < steps; i++ {
		arr[0]++
		arr[1]++
		arr[2]++
		arr[3]++
	}
}

func main() {
}
