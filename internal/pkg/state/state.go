package state

import (
	"errors"
	"fmt"
	"tew4fs/trivia-backend/internal/pkg/action"
	"tew4fs/trivia-backend/internal/pkg/constant"
	"tew4fs/trivia-backend/internal/pkg/model"
)

type State struct {
	CurrentQuestion model.Question
	PlayerScores    map[string]int
	PlayersAnswered map[string]bool
	GameStarted     bool
}

func NewState() *State {
	s := &State{
		CurrentQuestion: *action.GetQuestion(),
		PlayerScores:    map[string]int{},
		PlayersAnswered: map[string]bool{},
		GameStarted:     false,
	}
	return s
}

func (s *State) StartGame() {
	s.GameStarted = true
}

func (s *State) NextQuestion() {
	s.CurrentQuestion = *action.GetQuestion()
	for user, _ := range s.PlayersAnswered {
		s.PlayersAnswered[user] = false
	}
}

func (s *State) AddPlayer(user string) error {
	if _, has := s.PlayerScores[user]; has {
		return errors.New(fmt.Sprintf("player with name %s has already been added", user))
	}
	s.PlayerScores[user] = 0
	s.PlayersAnswered[user] = false
	return nil
}

func (s *State) UpdateScore(user string, isCorrect bool) error {
	if _, has := s.PlayerScores[user]; !has {
		return errors.New("player has not been added yet")
	}

	if s.PlayersAnswered[user] {
		return errors.New(fmt.Sprintf("%s has already answered this question", user))
	}
	s.PlayersAnswered[user] = true

	if isCorrect {
		s.PlayerScores[user] += constant.QuestionPoints
	}
	return nil
}
