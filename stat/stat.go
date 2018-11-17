package stat

// Min of nums
func Min(nums []float64) float64 {
	var min float64
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			min = nums[i]
			continue
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

// Max of nums
func Max(nums []float64) float64 {
	var max float64
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = nums[i]
			continue
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// Average of nums
func Average(nums []float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	avg := sum / float64(len(nums))
	return avg
}

// Report sturcture
type Report struct {
	Min     float64
	Max     float64
	Average float64
}

// MakeReport use for return min, max and average of nums
func MakeReport(nums []float64) Report {
	min := Min(nums)
	max := Max(nums)
	average := Average(nums)
	report := Report{
		Min:     min,
		Max:     max,
		Average: average,
	}
	return report
}
