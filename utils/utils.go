package utils

// import (
// 	"fmt"
// 	"math"
// )

// func MergeBytes(h, l uint8) uint16 {
// 	return uint16(h)<<8 + uint16(l)
// }

// func SumBytes(data []float64) float64 {
// 	var sum float64 = 0
// 	for _, v := range data {
// 		sum += v
// 	}

// 	return sum
// }

// func MultipleBytes(data []float64) float64 {
// 	var product float64 = 1.0

// 	for _, v := range data {
// 		product *= v
// 	}

// 	return product
// }

// func Average(data []float64) float64 {
// 	sum := SumBytes(data)

// 	return sum / float64(len(data))
// 	//return 1
// }

// func Stdev(data []float64) float64 {
// 	diff := make([]float64, len(data), cap(data))
// 	avg := Average(data)
// 	fmt.Println("avg=", avg)
// 	for i, v := range data {
// 		diff[i] = math.Pow((v - avg), 2)
// 	}

// 	return math.Sqrt(Average(diff))
// }
