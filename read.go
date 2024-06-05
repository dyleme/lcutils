package leetcode

import (
	"strconv"
	"strings"
)

func ReadSlice(input string) []int {
	input = strings.TrimPrefix(input, "[")
	input = strings.TrimSuffix(input, "]")
	split := strings.Split(input, ",")
	result := make([]int, 0, len(split))
	for i := 0; i < len(split); i++ {
		k, err := strconv.Atoi(split[i])
		if err != nil {
			panic(err)
		}

		result = append(result, k)
	}

	return result
}

func ReadMatrix(input string) [][]int {
	input = strings.TrimPrefix(input, "[[")
	input = strings.TrimSuffix(input, "]]")
	rows := strings.Split(input, "],[")
	matrix := make([][]int, 0, len(rows))
	for _, r := range rows {
		vals := strings.Split(r, ",")
		row := make([]int, 0, len(vals))
		for _, v := range vals {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			row = append(row, i)
		}

		matrix = append(matrix, row)
	}

	return matrix
}
