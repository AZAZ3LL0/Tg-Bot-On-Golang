package services

import (
	"sync"
)

type UserStateService struct {
	userStates map[int64]int // userID -> state
	mu         sync.Mutex
}

const (
	InitialState = iota
	WaitingForAuthor
	WaitingForBookID
	// ... добавьте другие состояния, если необходимо
)

func NewUserStateService() *UserStateService {
	return &UserStateService{
		userStates: make(map[int64]int),
	}
}

func (s *UserStateService) SetUserState(userID int64, state int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.userStates[userID] = state
}

func (s *UserStateService) GetUserState(userID int64) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	state, ok := s.userStates[userID]
	if !ok {
		return InitialState
	}
	return state
}

func (s *UserStateService) ResetUserState(userID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.userStates, userID)
}
