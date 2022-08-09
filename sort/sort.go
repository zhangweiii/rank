package sort

import (
	ossort "sort"

	"github.com/zhangweiii/rank/model"
)

func Sort(scores []*model.Score) {
	ossort.Slice(scores, func(i, j int) bool {
		if scores[i].Score == scores[j].Score {
			return scores[i].UpdatedAt < scores[j].UpdatedAt
		}
		return scores[i].Score < scores[j].Score
	})
}
