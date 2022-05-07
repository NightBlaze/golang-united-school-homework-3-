package homework

func average(input [15]float32) (result float32) {
	for _, item := range input {
		result += item
	}
	result /= float32(len(input))
	return
}
