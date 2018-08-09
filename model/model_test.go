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
