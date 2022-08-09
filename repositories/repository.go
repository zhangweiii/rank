package repositories

import "github.com/zhangweiii/rank/model"

type Repository interface {
	InsertOrUpdate(score model.Score) error
	Find(id string) (*model.Score, error)
	FindAll() ([]*model.Score, error)
	Delete(id string) error
	Rank(id string, beforeN int, afterN int) ([]*model.Score, error)
	InsertShared(scores []*model.Score) error
}
