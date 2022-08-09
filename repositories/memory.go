package repositories

import (
	"errors"
	"sync"

	"github.com/zhangweiii/rank/model"
)

type memoryRepository struct {
	lock       sync.RWMutex
	Scores     []*model.Score
	Tables     [][]*model.Score
	SharedSize int
}

// Delete implements Repository
func (r *memoryRepository) Delete(id string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for i, s := range r.Scores {
		if s.ID == id {
			r.Scores = append(r.Scores[:i], r.Scores[i+1:]...)
			return nil
		}
	}

	return nil
}

// Find implements Repository
func (r *memoryRepository) Find(id string) (*model.Score, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, s := range r.Scores {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, errors.New("not found")
}

// FindAll implements Repository
func (r *memoryRepository) FindAll() ([]*model.Score, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.Scores, nil
}

// Insert implements Repository
func (r *memoryRepository) InsertOrUpdate(score model.Score) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, s := range r.Scores {
		if s.ID == score.ID {
			s.Score = s.Score + score.Score
			return nil
		}
	}

	r.Scores = append(r.Scores, &score)
	return nil
}

// Rank implements Repository
func (r *memoryRepository) Rank(id string, beforeN int, afterN int) ([]*model.Score, error) {
	score, err := r.Find(id)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Score, 0, beforeN+afterN+1)

	// TODO: find the rank of the score

	return result, nil
}

// InsertShared implements Repository
func (r *memoryRepository) InsertShared(scores []*model.Score) error {
	length := len(scores)
	if length == 0 {
		return nil
	}

	rank := 0
	for i := 0; r.SharedSize*i < length; i++ {
		r.Tables = append(r.Tables, make([]*model.Score, 0, r.SharedSize))
		max := r.SharedSize * (i + 1)
		if max > length {
			max = length
		}
		tmpScores := scores[r.SharedSize*i : max]
		for _, s := range tmpScores {
			s.Rank = rank + 1
			rank += 1
			r.Tables[i] = append(r.Tables[i], s)
		}
	}

	return nil
}

func New(opts ...func(*memoryRepository) error) Repository {
	mr := &memoryRepository{}
	for _, opt := range opts {
		opt(mr)
	}

	return mr
}

func WithSharedSize(size int) func(*memoryRepository) error {
	return func(r *memoryRepository) error {
		r.SharedSize = size
		return nil
	}
}
