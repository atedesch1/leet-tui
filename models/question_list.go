package models

import (
	"fmt"

	"github.com/atedesch1/leet-tui/api"
	tea "github.com/charmbracelet/bubbletea"
)

type questionListModel struct {
	totalNumQuestions int
	questions         []api.QuestionInfo
	skip              int
	limit             int
	categorySlug      string

	cursor  int
	loading bool
}

func newQuestionListModel() *questionListModel {
	return &questionListModel{
		totalNumQuestions: 0,
		questions:         []api.QuestionInfo{},
		skip:              0,
		limit:             10,
		categorySlug:      "",

		cursor:  0,
		loading: true,
	}
}

type problemsetQuestionListResponseMsg struct {
	questionList api.ProblemsetQuestionList
}

func (m *questionListModel) getProblemsetQuestionListCmd() tea.Msg {
	problemsetQuestionList, _ := api.GetProblemsetQuestionList(m.categorySlug, m.skip, m.limit)
	return problemsetQuestionListResponseMsg{
		questionList: problemsetQuestionList,
	}
}

func (m *questionListModel) goToPreviousPage() tea.Cmd {
	if m.skip < m.limit {
		return nil
	}

	m.skip -= m.limit
	return nil
}

func (m *questionListModel) goToNextPage() tea.Cmd {
	if m.skip+2*m.limit > m.totalNumQuestions {
		return nil
	}

	m.skip += m.limit
	if len(m.questions) < m.skip+m.limit {
		m.loading = true
		return m.getProblemsetQuestionListCmd
	}
	return nil
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
	return m.getProblemsetQuestionListCmd
}

func (m *questionListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case problemsetQuestionListResponseMsg:
		m.loading = false
		m.totalNumQuestions = msg.questionList.Total
		m.questions = append(m.questions, msg.questionList.Questions...)

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else if m.cursor == 0 && m.skip >= m.limit {
				cmd = m.goToPreviousPage()
				m.cursor = m.limit - 1
			}

		case "down", "j":
			if m.cursor < m.limit-1 {
				m.cursor++
			} else if m.cursor == m.limit-1 && m.skip+2*m.limit <= m.totalNumQuestions {
				cmd = m.goToNextPage()
				m.cursor = 0
			}

		case "right", "h":
			cmd = m.goToNextPage()

		case "left", "l":
			cmd = m.goToPreviousPage()

		case "enter", " ":
			return m, m.selectQuestionCmd
		}

	}
	return m, cmd
}

func (m *questionListModel) View() string {
	s := "Choose a problem from below:\n\n"

	if m.loading {
		s += "  Loading..."
		return s
	}

	for i, choice := range m.questions[m.skip : m.skip+m.limit] {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice.Title)
	}

	s += fmt.Sprintf("\nPage %d/%d", m.skip/m.limit+1, m.totalNumQuestions/m.limit)

	return s
}
