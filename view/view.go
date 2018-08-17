package view

import (
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
            // TODO Store answers
            view.CurrentItem += 1
            if view.CurrentItem >= controller.GetNumberOfItems() {
                view.State = EndingState
                // TODO Display end result
                view.Output = append(view.Output, "Deseja tentar novamente? [s/N]")
            } else {
                // TODO Load new items
                view.Output = append(view.Output, "Mais uma questão...")
            }
        } else {
            view.Output = append(view.Output, "Digite uma idade válida")
        }
        break
    case EndingState:
        endMe := true
        if len(line) > 0 {
            if (line[0] == 's') || (line[0] == 'S') {
                view.State = TestState
                view.CurrentItem = 0
                view.Output = append(view.Output, "Item A1:")
                endMe = false
            }
        }
        view.Working = !endMe
    }
}
