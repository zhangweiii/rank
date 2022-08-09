package repositories

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/zhangweiii/rank/model"
)

func TestMemory(t *testing.T) {
	Convey("Find one after insert it", t, func() {
		repo := New()
		score := model.Score{
			ID:        "1",
			Score:     100,
			UpdatedAt: 2,
		}
		err := repo.InsertOrUpdate(score)
		So(err, ShouldBeNil)
		s, err := repo.Find("1")
		So(err, ShouldBeNil)
		So(s, ShouldResemble, &score)
	})

	Convey("Repeat insert one will be conver", t, func() {
		repo := New()
		score := model.Score{
			ID:        "1",
			Score:     100,
			UpdatedAt: 2,
		}
		err := repo.InsertOrUpdate(score)
		So(err, ShouldBeNil)
		s, err := repo.Find("1")
		So(err, ShouldBeNil)
		So(s, ShouldResemble, &score)

		score.Score = 3
		err = repo.InsertOrUpdate(score)
		So(err, ShouldBeNil)
		s, err = repo.Find("1")
		So(err, ShouldBeNil)
		So(s.Score, ShouldEqual, 103)

		scores, err := repo.FindAll()
		So(err, ShouldBeNil)
		So(len(scores), ShouldEqual, 1)
	})

	Convey("Delete one", t, func() {
		repo := New()
		score := model.Score{
			ID:        "1",
			Score:     100,
			UpdatedAt: 2,
		}
		err := repo.InsertOrUpdate(score)
		So(err, ShouldBeNil)
		s, err := repo.Find("1")
		So(err, ShouldBeNil)
		So(s, ShouldResemble, &score)

		err = repo.Delete("1")
		So(err, ShouldBeNil)
		s, err = repo.Find("1")
		So(err, ShouldResemble, errors.New("not found"))
		So(s, ShouldBeNil)
	})

	Convey("Insert shared and search rank", t, func() {
		size := 3
		repo := New(WithSharedSize(size))
		scores := []*model.Score{
			{ID: "1", Score: 100, UpdatedAt: 2},
			{ID: "2", Score: 101, UpdatedAt: 2},
			{ID: "3", Score: 102, UpdatedAt: 2},

			{ID: "4", Score: 103, UpdatedAt: 2},
			{ID: "5", Score: 104, UpdatedAt: 2},
			{ID: "6", Score: 105, UpdatedAt: 2},

			{ID: "7", Score: 106, UpdatedAt: 2},
			{ID: "8", Score: 107, UpdatedAt: 2},
		}

		err := repo.InsertShared(scores)
		So(err, ShouldBeNil)
		So((repo.(*memoryRepository)).Tables, ShouldHaveLength, 3)
		So((repo.(*memoryRepository)).Tables[0], ShouldResemble, []*model.Score{
			{ID: "1", Score: 100, UpdatedAt: 2, Rank: 1},
			{ID: "2", Score: 101, UpdatedAt: 2, Rank: 2},
			{ID: "3", Score: 102, UpdatedAt: 2, Rank: 3},
		})
		So((repo.(*memoryRepository)).Tables[1], ShouldResemble, []*model.Score{
			{ID: "4", Score: 103, UpdatedAt: 2, Rank: 4},
			{ID: "5", Score: 104, UpdatedAt: 2, Rank: 5},
			{ID: "6", Score: 105, UpdatedAt: 2, Rank: 6},
		})
		So((repo.(*memoryRepository)).Tables[2], ShouldResemble, []*model.Score{
			{ID: "7", Score: 106, UpdatedAt: 2, Rank: 7},
			{ID: "8", Score: 107, UpdatedAt: 2, Rank: 8},
		})

		for _, s := range scores {
			repo.InsertOrUpdate(*s)
		}
		rank2, err := repo.Rank("5", 1, 1)
		So(err, ShouldBeNil)
		So(rank2, ShouldResemble, []*model.Score{
			{ID: "4", Score: 103, UpdatedAt: 2, Rank: 4},
			{ID: "5", Score: 104, UpdatedAt: 2, Rank: 5},
			{ID: "6", Score: 105, UpdatedAt: 2, Rank: 6},
		})
		rank1, err := repo.Rank("4", 1, 1)
		So(err, ShouldBeNil)
		So(rank1, ShouldResemble, []*model.Score{
			{ID: "3", Score: 102, UpdatedAt: 2},
			{ID: "4", Score: 103, UpdatedAt: 2},
			{ID: "5", Score: 104, UpdatedAt: 2},
		})
	})
}
