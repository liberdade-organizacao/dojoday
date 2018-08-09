package model

import (
    "os"
    "io/ioutil"
    "strings"
    "strconv"
    "errors"
)

// Just a function to kickstart stuff
func GetTrue() bool {
    return true
}

func GetSourceFolder() string {
    return os.Getenv("DOJODAY_FILES")
}

// Gets the percentile table file path
func getPercentileFilePath() string {
    return GetSourceFolder() + "percentil.csv"
}

// Gets the answers table file path
func getAnswerFilePath() string {
    return GetSourceFolder() + "respostas.csv"
}

// Gets the validity table file path
func getValidityFilePath() string {
    return GetSourceFolder() + "validade.csv"
}

// Gets the percentile list
func GetPercentile() []int {
    outlet := make([]int, 0)
    lines, oops := LoadLines(getPercentileFilePath())
    if oops != nil {
        return outlet
    }
    firstLine := true
    for _, line := range lines {
        if firstLine {
            firstLine = false
        } else {
            fields := strings.Split(line, "\t")
            field, oops := strconv.Atoi(fields[0])
            if oops == nil {
                outlet = append(outlet, field)
            }
        }
    }
    return outlet
}

// Gets minimum score for a percentile based on the age
func GetPercentileScoreByAge(age int) ([]int, error) {
    outlet := make([]int, 0)
    lines, oops := LoadLines(getPercentileFilePath())
    if oops != nil {
        return outlet, oops
    }
    targetColumn := -1
    firstLine := true
    for _, rawLine := range lines {
        line := strings.Trim(rawLine, "\r")
        fields := strings.Split(line, "\t")
        if firstLine {
            for i, field := range fields {
                maybeAge, oops := strconv.Atoi(field)
                if oops == nil {
                    if maybeAge == age {
                        targetColumn = i
                    }
                }
            }
            if targetColumn <= 0 {
                return outlet, errors.New("Invalid age")
            }
            firstLine = false
        } else {
            if len(fields) > targetColumn {
                score, oops := strconv.Atoi(fields[targetColumn])
                if oops == nil {
                    outlet = append(outlet, score)
                }
            }
        }
    }
    return outlet, nil
}

// This function get the each test item's id.
func GetTestItems() []string {
    outlet := make([]string, 0)
    lines, oops := LoadLines(getAnswerFilePath())
    if oops != nil {
        return outlet
    }
    firstLine := true
    for _, line := range lines {
        if !firstLine {
            fields := strings.Split(line, "\t")
            if len(fields) > 1 {
                outlet = append(outlet, fields[0])
            }
        }
        firstLine = false
    }
    return outlet
}

// Gets how many options there are for each item
func GetHowManyOptionsForItem() []int {
    outlet := make([]int, 0)
    lines, oops := LoadLines(getAnswerFilePath())
    if oops != nil {
        return outlet
    }
    firstLine := true
    for _, line := range lines {
        if !firstLine {
            fields := strings.Split(line, "\t")
            if len(fields) > 1 {
                maybe, oops := strconv.Atoi(fields[1])
                if oops == nil {
                    outlet = append(outlet, maybe)
                }
            }
        }
        firstLine = false
    }
    return outlet
}

// Gets the correct answer for each item
func GetCorrectAnswers() []int {
    outlet := make([]int, 0)
    lines, oops := LoadLines(getAnswerFilePath())
    if oops != nil {
        return outlet
    }
    firstLine := true
    for _, line := range lines {
        if !firstLine {
            fields := strings.Split(line, "\t")
            if len(fields) > 1 {
                maybe, oops := strconv.Atoi(fields[2])
                if oops == nil {
                    outlet = append(outlet, maybe)
                }
            }
        }
        firstLine = false
    }
    return outlet
}

// Gets which series the current item is in
func GetItemSeries() []string {
    outlet := make([]string, 0)
    lines, oops := LoadLines(getAnswerFilePath())
    if oops != nil {
        return outlet
    }
    firstLine := true
    for _, rawLine := range lines {
        line := strings.Trim(rawLine, "\r")
        if !firstLine {
            fields := strings.Split(line, "\t")
            if len(fields) > 1 {
                outlet = append(outlet, fields[3])
            }
        }
        firstLine = false
    }
    return outlet
}

// Loads the a file into an array of strings, one for each line
func LoadLines(filepath string) ([]string, error) {
    outlet := make([]string, 0)
    rawContent, oops := ioutil.ReadFile(filepath)
    if oops != nil {
        return outlet, oops
    }
    content := string(rawContent)
    outlet = strings.Split(content, "\n")
    return outlet, nil
}
