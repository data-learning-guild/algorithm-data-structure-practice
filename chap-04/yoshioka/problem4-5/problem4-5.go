package main

func SevenFiveThree(input_number int) int {
	output, count_3, count_5, count_7 := count_number(input_number, 0, 0, 0)

	if output != 0 {
		panic("Unknown result")
	}

	min_num := Min(count_3, Min(count_5, count_7))
	return min_num

}

func count_number(input int, count_3 int, count_5 int, count_7 int) (int, int, int, int) {
	if input > 0 {
		target_int := input % 10
		remain_int := input / 10

		switch target_int {
		case 3:
			count_3++
		case 5:
			count_5++
		case 7:
			count_7++
		default:
			panic("Unknown integer")

		}

		return count_number(remain_int, count_3, count_5, count_7)
	} else {
		return input, count_3, count_5, count_7
	}

}
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func main() {

}
