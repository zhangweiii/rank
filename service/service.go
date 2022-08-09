package service

import (
	"fmt"

	"github.com/zhangweiii/rank/model"
	"github.com/zhangweiii/rank/repositories"
	"github.com/zhangweiii/rank/sort"
)

type Service struct {
	repo repositories.Repository
}

// NewService create a new service
func NewService(repo repositories.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// BeginRank begin rank when call
// 1. find all scores
// 2. sort scores
// 3. insert scores to shared db
func (s *Service) BeginRank() error {
	scores, err := s.repo.FindAll()
	if err != nil {
		return fmt.Errorf("find all scores error: %w", err)
	}

	sort.Sort(scores)

	err = s.repo.InsertShared(scores)
	if err != nil {
		return fmt.Errorf("insert shared error: %w", err)
	}
	return nil
}

func (s *Service) Rank(id string, beforeN int, afterN int) ([]*model.Score, error) {
	return s.repo.Rank(id, beforeN, afterN)
}
