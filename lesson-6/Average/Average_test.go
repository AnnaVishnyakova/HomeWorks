package Average

import "testing"

type TestCase1 struct {
	Name          string
	InputValue    []float64
	ExpectedValue float64
}
type TestCase2 struct {
	Name       string
	InputValue []float64
	SumValue   float64
}

func TestAverage(t *testing.T) {
	testCases1 := []TestCase1{
		{
			Name:          "TestCase1",
			InputValue:    []float64{4, 4, 4, 4, 4, 4},
			ExpectedValue: 4,
		},
		{
			Name:          "TestCase2",
			InputValue:    []float64{1, 2, 1, 2},
			ExpectedValue: 1.5,
		},
	}
	testCases2 := []TestCase2{
		{
			Name:       "Testsum1",
			InputValue: []float64{4, 4, 4, 4, 4, 4},
			SumValue:   24,
		},
		{
			Name:       "Testsum2",
			InputValue: []float64{1, 2, 1, 2},
			SumValue:   6,
		},
	}
	for _, testCase1 := range testCases1 {
		resultValue := Average(testCase1.InputValue)
		if resultValue != testCase1.ExpectedValue {
			t.Errorf("Failed! Test case: %s; expected value %.2f != actual value %.2f", testCase1.Name, testCase1.ExpectedValue, resultValue)
		}

	}
	for _, testCase2 := range testCases2 {
		resultValue2 := Sum(testCase2.InputValue)
		if resultValue2 != testCase2.SumValue {
			t.Errorf("Failed! Test case: %s; sum value %.2f != actual value %.2f", testCase2.Name, testCase2.SumValue, resultValue2)
		}

	}

}
