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

func TestIfPercentileCanBeExtracted(t *testing.T) {
    percentile := GetPercentileByAgeAndScore(5, 20)
    if percentile != 90 {
        t.Error("Couldn't get percentile from a 5yo that got 20 points")
    }

    percentile = GetPercentileByAgeAndScore(8, 22)
    if percentile != 50 {
        t.Error("Couldn't get percentile from a 8yo that got 22 points")
    }

    percentile = GetPercentileByAgeAndScore(11, 16)
    if percentile != 0 {
        t.Error("Couldn't get percentile from a 11yo that got 16 points")
    }
}
