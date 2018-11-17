package stat

import "testing"

func TestStat_Min_1_3_5_return_1(t *testing.T) {
	r := Min([]float64{1, 3, 5})
	if r != 1 {
		t.Error("Expected 1 got", r)
	}
}

func TestStat_Min_5_5_4_return_4(t *testing.T) {
	r := Min([]float64{5, 5, 4})
	if r != 4 {
		t.Error("Expected 4 got", r)
	}
}
func TestStat_Min_5_3_4_return_3(t *testing.T) {
	r := Min([]float64{5, 3, 4})
	if r != 3 {
		t.Error("Expected 3 got", r)
	}
}

func TestStat_Max_1_3_5_return_5(t *testing.T) {
	r := Max([]float64{1, 3, 5})
	if r != 5 {
		t.Error("Expected 5 got", r)
	}
}

func TestStat_Max_1_6_3_return_6(t *testing.T) {
	r := Max([]float64{1, 6, 3})
	if r != 6 {
		t.Error("Expected 6 got", r)
	}
}

func TestStat_Max_7_6_3_return_7(t *testing.T) {
	r := Max([]float64{7, 6, 3})
	if r != 7 {
		t.Error("Expected 7 got", r)
	}
}

func TestStat_Average_3_2_4_return_3(t *testing.T) {
	r := Average([]float64{3, 2, 4})
	if r != 3 {
		t.Error("Expected 3 got", r)
	}
}

func TestStat_MakeReport_3_2_4_return_2_4_3(t *testing.T) {
	r := MakeReport([]float64{3, 2, 4})
	tempReport := Report{
		Min:     2,
		Max:     4,
		Average: 3,
	}
	if r != tempReport {
		t.Error("Expected 2,4,3 got", r)
	}
}
