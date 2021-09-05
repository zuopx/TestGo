package algorithm_test

import (
	"fmt"
	"sort"
	"testing"
	"touchgo/utils"
)

func TestSorting(t *testing.T) {
	sortDemo1()
	utils.Hello()
}

// < 正序；> 逆序
func sortDemo1() {
	t1 := Team{Score: 1}
	t2 := Team{Score: 2}
	t3 := Team{Score: 3}
	ts := []Team{t2, t1, t3}
	fmt.Println(ts)

	sort.Slice(ts, func(i, j int) bool {
		return ts[i].Score > ts[j].Score
	})
	fmt.Println(ts)
}

type Team struct {
	Score int
}
