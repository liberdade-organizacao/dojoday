package view

import (
    "fmt"
    "strconv"
    "strings"
    "github.com/liberdade-organizacao/dojoday/controller"
)

type StateEnum int
const (
    AgeQuestionState StateEnum = 1 << iota
    TestState
    EndingState
)

type View struct {
    Working bool
    CurrentItem int
    Age int
    State StateEnum
    Output []string
    Answers []int
}

func NewView() View {
    return View {
        Working: true,
        CurrentItem: 0,
        Age: 10,
        State: AgeQuestionState,
        Output: []string {
            "Digite a sua idade:",
        },
        Answers: make([]int, 0),
    }
}

func (view *View) Write() string {
    output := ""
    for _, line := range view.Output {
        output += line + "\n"
    }
    view.Output = make([]string, 0)
    return output
}

func (view *View) Read(line string) {
    // TODO Implement the rest of me!
    switch view.State {
    case AgeQuestionState:
        age, oops := strconv.Atoi(strings.Trim(line, "\r\n"))
        if oops != nil {
            view.Output = append(view.Output, "Digite um valor válido!")
        } else if controller.ValidateAge(age) {
            view.Age = age
            view.State = TestState
            view.CurrentItem = 0
            view.Answers = make([]int, 0)
            view.Output = append(view.Output, "Item A1:")
        } else {
            view.Output = append(view.Output, "Digite uma idade válida")
        }
        break
    case TestState:
        answer, oops := strconv.Atoi(strings.Trim(line, "\r\n"))
        if oops != nil {
            view.Output = append(view.Output, "Digite um valor válido!")
        } else if controller.ValidateAnswer(view.CurrentItem, answer) {
            view.Answers = append(view.Answers, answer)
            view.CurrentItem += 1
            if view.CurrentItem >= controller.GetNumberOfItems() {
                view.State = EndingState
                feedback := "Seu resultado é inconclusivo"
                if percentile := controller.ValidateApplication(view.Age, view.Answers); percentile >= 0 {
                     feedback = fmt.Sprintf("Você está entre os %d%% mais inteligentes da população com %d anos", 100-percentile, view.Age)
                }
                view.Output = append(view.Output, feedback)
                view.Output = append(view.Output, "Deseja tentar novamente? [s/N]")
            } else {
                item := fmt.Sprintf("Item %s", controller.GetNthItem(view.CurrentItem))
                view.Output = append(view.Output, item)
            }
        } else {
            view.Output = append(view.Output, "Digite uma idade válida")
        }
        break
    case EndingState:
        endMe := true
        if len(line) > 0 {
            if (line[0] == 's') || (line[0] == 'S') {
                view.State = AgeQuestionState
                view.Output = append(view.Output, "Digite a sua idade:")
                endMe = false
            }
        }
        view.Working = !endMe
    }
}
