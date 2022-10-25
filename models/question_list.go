package models

import (
	"fmt"

	"github.com/atedesch1/leet-tui/api"
	tea "github.com/charmbracelet/bubbletea"
)

type questionListModel struct {
	cursor            int
	totalNumQuestions int
	questions         []api.QuestionInfo
	skip              int
	limit             int
	categorySlug      string
}

func newQuestionListModel() *questionListModel {
	return &questionListModel{
		cursor:            0,
		totalNumQuestions: 0,
		questions:         []api.QuestionInfo{},
		skip:              0,
		limit:             10,
		categorySlug:      "",
	}
}

func (m *questionListModel) getProblemsetQuestionList() tea.Msg {
	problemsetQuestionList, _ := api.GetProblemsetQuestionList(m.categorySlug, m.skip, m.limit)
	return problemsetQuestionList
}

type SelectedQuestionMsg struct {
	question api.QuestionInfo
}

func (m *questionListModel) selectQuestionCmd() tea.Msg {
	return SelectedQuestionMsg{
		question: m.questions[m.skip+m.cursor],
	}
}

func (m *questionListModel) Init() tea.Cmd {
	return m.getProblemsetQuestionList
}

func (m *questionListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case api.ProblemsetQuestionList:
		m.totalNumQuestions = msg.Total
		m.questions = append(m.questions, msg.Questions...)

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.questions)-1 {
				m.cursor++
			}

		case "right", "h":
			if m.skip+2*m.limit > m.totalNumQuestions {
				break
			}
			m.skip += m.limit
			if len(m.questions) < m.skip+m.limit {
				return m, m.getProblemsetQuestionList
			}

		case "left", "l":
			if m.skip >= 10 {
				m.skip -= m.limit
			}

		case "enter", " ":
			return m, m.selectQuestionCmd
		}

	}
	return m, nil
}

func (m *questionListModel) View() string {
	if len(m.questions) < m.skip+m.limit {
		return "Loading...\n\n"
	}

	s := "Choose the question!\n\n"

	for i, choice := range m.questions[m.skip : m.skip+m.limit] {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	s += fmt.Sprintf("\n\nPage %d/%d", m.skip/m.limit+1, m.totalNumQuestions/m.limit)

	return s
}
