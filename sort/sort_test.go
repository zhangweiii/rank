package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/zhangweiii/rank/model"
)

func TestSort(t *testing.T) {
	Convey("TestSort", t, func() {
		scores := []*model.Score{
			{ID: "1", Score: 100, UpdatedAt: 2},
			{ID: "4", Score: 101, UpdatedAt: 1},
			{ID: "2", Score: 100, UpdatedAt: 1},
			{ID: "3", Score: 77, UpdatedAt: 1},
		}
		Sort(scores)
		So(scores, ShouldResemble, []*model.Score{
			{ID: "3", Score: 77, UpdatedAt: 1},
			{ID: "2", Score: 100, UpdatedAt: 1},
			{ID: "1", Score: 100, UpdatedAt: 2},
			{ID: "4", Score: 101, UpdatedAt: 1},
		})
	})
	// 没 benchmark 是因为 sort 是原地排序，没法准备 b.N 个待排序数组
	Convey("Test 1000000 record spend time", t, func() {
		scores := []*model.Score{}
		for i := 0; i < 1000000; i++ {
			scores = append(scores, &model.Score{
				ID:        fmt.Sprintf("%d", i),
				Score:     rand.Intn(10000),
				UpdatedAt: uint32(time.Now().Add(time.Second * time.Duration(rand.Intn(100))).Unix()),
			})
		}
		start := time.Now()
		Sort(scores)
		t.Log("spare time:", time.Since(start).Milliseconds())
	})
}
