package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nodes [][3]int
var reader = bufio.NewReader(os.Stdin)

func walk(parent, min int) int {
	if parent < 0 {
		return min
	}
	min = walk(nodes[parent][1], min)
	nodes[parent][0] = min
	min = walk(nodes[parent][2], min+1)
	return min
}

func main() {
	var n, index, left, right int
	fmt.Scanf("%d\n", &n)
	nodes = make([][3]int, n)
	for i := 0; i < n; i++ {
		scan(&index, &left, &right)
		nodes[index-1][1] = left - 1
		nodes[index-1][2] = right - 1
		if left != -1 {
			nodes[left-1][0] = index
		}
		if right != -1 {
			nodes[right-1][0] = index
		}
	}
	for index = 0; nodes[index][0] != 0; index++ {
	}
	nodes[index][0] = 9
	walk(index, 1)
	for _, v := range nodes {
		fmt.Printf("%v ", v[0])
	}
}

func scan(index, left, right *int) {
	aMembers, _ := reader.ReadString('\n')
	aMembers = strings.TrimSuffix(aMembers, "\n")
	numbers := strings.Split(aMembers, " ")
	*index, _ = strconv.Atoi(numbers[0])
	*left, _ = strconv.Atoi(numbers[1])
	*right, _ = strconv.Atoi(numbers[2])
}
