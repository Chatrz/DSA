package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	var m, n int
	intScan(&n, &m)
	table := make([][]int, n+2)

	table[0] = make([]int, m+2)
	table[n+1] = make([]int, m+2)
	for i := 1; i <= n; i++ {
		table[i] = make([]int, m+2)
		var str string
		strScan(&str)
		for j := 1; j <= m; j++ {

			if str[j-1] == '#' {
				table[i][j] = -1
			} else if str[j-1] == '.' {
				table[i][j] = 1
			}
		}
	}
	var q int
	intScan(&q)
	res := make([]string, q)
	questions := make([][]int, q)

	for i := 0; i < q; i++ {
		questions[i] = make([]int, 4)
		intScan(&questions[i][0], &questions[i][1], &questions[i][2], &questions[i][3])

	}
	for i := 0; i < q; i++ {
		if bfs(table, m*n, questions[i][0], questions[i][1], questions[i][2], questions[i][3], i+2) {
			res[i] = "YES"
		} else {
			res[i] = "NO"
		}
	}
	for i := 0; i < q; i++ {
		fmt.Println(res[i])
	}
}

func bfs(table [][]int, nm, r1, c1, r2, c2, i int) bool {
	// if r1 == r2 && c1 == c2 {
	// 	return true
	// }
	if (table[r1][c1] == -1) || (table[r2][c2] == -1) || (table[r1][c1] == 0) || (table[r2][c2] == 0) {
		return false
	}
	if (table[r1][c1] > 1 && table[r2][c2] == 1) || table[r2][c2] > 1 && table[r1][c1] == 1 {
		return false
	}
	if table[r1][c1] > 1 && table[r2][c2] > 1 {
		if table[r1][c1] == table[r2][c2] {
			return true
		}
		return false
	}
	queue := make([][]int, 0)
	pair := make([]int, 2)
	pair[0], pair[1] = r2, c2
	queue = append(queue, pair)
	counter := 0
	for len(queue) != 0 && counter < nm {
		top := queue[0]
		queue = queue[1:]
		if table[top[0]-1][top[1]] == 1 {
			table[top[0]-1][top[1]] = i
			tempPair := make([]int, 2)
			tempPair[0], tempPair[1] = top[0]-1, top[1]
			queue = append(queue, tempPair)
		}
		if table[top[0]+1][top[1]] == 1 {
			table[top[0]+1][top[1]] = i
			tempPair := make([]int, 2)
			tempPair[0], tempPair[1] = top[0]+1, top[1]
			queue = append(queue, tempPair)
		}
		if table[top[0]][top[1]-1] == 1 {
			table[top[0]][top[1]-1] = i
			tempPair := make([]int, 2)
			tempPair[0], tempPair[1] = top[0], top[1]-1
			queue = append(queue, tempPair)
		}
		if table[top[0]][top[1]+1] == 1 {
			table[top[0]][top[1]+1] = i
			tempPair := make([]int, 2)
			tempPair[0], tempPair[1] = top[0], top[1]+1
			queue = append(queue, tempPair)
		}
		counter++
	}
	if counter == nm {
		return false
	}
	return (table[r1][c1] == i)
}

func strScan(str *string) {
	aMembers, _ := reader.ReadString('\n')
	aMembers = strings.TrimSuffix(aMembers, "\n")
	*str = aMembers
}

func intScan(ints ...*int) {
	aMembers, _ := reader.ReadString('\n')
	aMembers = strings.TrimSuffix(aMembers, "\n")
	numbers := strings.Split(aMembers, " ")
	for k, _ := range numbers {
		*ints[k], _ = strconv.Atoi(numbers[k])
	}
}
