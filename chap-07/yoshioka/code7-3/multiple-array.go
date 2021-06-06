package main

func MultipleArray(array1, array2 []int) int {
	if len(array1) != len(array2) {
		panic("Wrong array arguments")
	}
	n := len(array1)

	sum := 0
	for i := n - 1; i >= 0; i-- {
		array1[i] += sum // Add the results until last process
		mod := array1[i] % array2[i]

		d := 0 //number of press the button
		if mod != 0 {
			d = array2[i] - mod //difference to the target num = number to press the button
		}
		sum += d

	}
	return sum

}
