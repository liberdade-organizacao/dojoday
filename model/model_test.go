package model

import (
    "testing"
)

func TestCanActuallyTest(t *testing.T) {
    isTrue := GetTrue()
    if isTrue != true {
        t.Error("true is not true? wtf??")
    }
}

func TestContainRequiredEnvironmentVariable(t *testing.T) {
    sourceFolder := GetSourceFolder()
    if len(sourceFolder) == 0 {
        t.Error("DOJODAY_FILES environment variable not set")
    }
}

func TestCanLoadPercentileTable(t *testing.T) {
    percentile := GetPercentile()
    if (percentile[0] != 99) || (percentile[len(percentile)-1] != 1) {
        t.Error("percentile table not loaded correctly")
    }
}

func TestCanLoadPercentileScoresByAge(t *testing.T) {
    _, oops := GetPercentileScoreByAge(13)
    if oops == nil {
        t.Error("Loaded invalid percentile scores by age")
    }

    scores, oops := GetPercentileScoreByAge(11)
    if (scores[0] != 35) || (scores[len(scores)-1] != 17) {
        t.Error("Loaded wrong score list for that age")
    }
}

func TestCanLoadTestItems(t *testing.T) {
    items := GetTestItems()
    if (items[0] != "A1") || (items[len(items)-1] != "B12") {
        t.Error("Couldn't load test items")
    }
}

func TestCanLoadOptionQuantityForEachItem(t *testing.T) {
    optionQty := GetHowManyOptionsForItem()
    for _, option := range optionQty {
        if option != 6 {
            t.Error("Wrong loading of option quantity")
            return
        }
    }

    if len(optionQty) != 36 {
        t.Error("Couldn't load all items")
    }
}

func TestCanLoadCorrectAnswers(t *testing.T) {
    correctAnswers := GetCorrectAnswers()
    if (correctAnswers[0] != 4) || (correctAnswers[len(correctAnswers)-1] != 5) {
        t.Error("Wrong loading of correct answers")
    }

    if len(correctAnswers) != 36 {
        t.Error("Couldn't load all items")
    }
}

func TestCanLoadItemSeries(t *testing.T) {
    series := GetItemSeries()
    correctSeries := []string { "a", "ab", "b" }
    for i := 0; i < 36; i++ {
        if series[i] != correctSeries[i/12] {
            t.Error("Couldn't load item series")
        }
    }
}

func TestCanLoadValidities(t *testing.T) {
    validities := GetValiditiesForSeries("a")
    if _, ok := validities[5]; ok == true {
        t.Error("Loaded validities from invalid score")
    }
    if baseScore := validities[10]; baseScore != 5 {
        t.Error("Couldn't load first validity score for A the series")
    }
    if baseScore := validities[35]; baseScore != 12 {
        t.Error("Couldn't load last validity score for A series")
    }

    validities = GetValiditiesForSeries("b")
    if baseScore := validities[10]; baseScore != 2 {
        t.Error("Couldn't load first validity score for B the series")
    }
    if baseScore := validities[35]; baseScore != 11 {
        t.Error("Couldn't load last validity score for B series")
    }
}
