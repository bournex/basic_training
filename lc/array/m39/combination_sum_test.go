package m39

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		Input  []int
		Target int
		Expect [][]int
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{{2, 2, 3}, {7}},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			[]int{20, 50},
			100,
			[][]int{{20, 20, 20, 20, 20}, {50, 50}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := combinationSum(v.Input, v.Target)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
				return
			}

			// 请无视低效测试代码，时间都用来实现问题算法了，这里其实就是字典序对两二维数组排序即可
			for _, v1 := range res {
				sort.Ints(v1)
				var found bool
				for _, v2 := range v.Expect {
					if len(v1) == len(v2) {
						sort.Ints(v2)
						found = true
						for k, v3 := range v1 {
							if v3 != v2[k] {
								found = false
								break
							}
						}
						if found {
							break
						} else {
							found = false
						}
					}
				}
				if !found {
					t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
					return
				}
			}
		})
	}
}
