package functions

import (
	"math"
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	sample := []float64{2, 4, 6}
	expected := 4.0
	result := Average(sample)
	roundresult := (math.Round(result))
	if !reflect.DeepEqual(roundresult, expected) {
		t.Errorf("TestFailed, expected: %v,found :%v ", expected, result)
	}
}

func TestMedian(t *testing.T) {
	sample := []float64{2, 4, 6}
	expected := 4.0
	result := Median(sample)
	roundresult := (math.Round(result))
	if !reflect.DeepEqual(roundresult, expected) {
		t.Errorf("TestFailed, expected: %v,found :%v ", expected, result)
	}
}

func TestVariance(t *testing.T) {
	sample := []float64{2, 4, 6}
	expected := 3.0
	result := Variance(sample)
	roundedResult := math.Round(result)
	if roundedResult != expected {
		t.Errorf("Test failed, expected: %v, found: %v", expected, roundedResult)
	}
}

func TestStandardDev(t *testing.T) {
	sample := []float64{2, 4, 6}
	expected := 2.0
	result := StandardDev(sample)
	roundedResult := math.Round(result)
	if roundedResult != expected {
		t.Errorf("TestFailed, expected: %v, found: %v ", expected, result)
	}
}
