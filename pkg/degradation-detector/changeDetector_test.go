package degradation_detector

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestChangeDetector(t *testing.T) {
  data := []int{5691, 5855, 5720, 6339, 5829, 5496, 5427, 5586, 5859, 5603, 5868, 5761, 5440, 5590, 5870, 5781, 5632, 6092, 5636, 5849, 5730, 5639, 5678, 5857, 5655, 5486, 5877, 5639, 5668, 5864, 5602, 5855, 6049, 5741, 5794, 5822, 5704, 5707, 6167, 5923, 5765, 5648, 5775, 5578, 5541, 5919, 5498, 5436, 5857, 5508, 5739, 5820, 5662, 5582, 5565, 5708, 5587, 5813, 5618, 5796, 5682, 5778, 5848, 6034, 5847, 5653, 5783, 6006, 5647, 5509, 5869, 5738, 5709, 5762, 5793, 5607, 5620, 5580, 5710, 5641, 5673, 5794, 5937, 5708, 5705, 5747, 5679, 5963, 6240, 5958, 5915, 5737, 6000, 5747, 5529, 5562, 5909, 5713, 5680, 5729, 5656, 5820, 5670, 5884, 5686, 5662, 5848, 5710, 5707, 5821, 5564, 6029, 6045, 5765, 5727, 5653, 5766, 5784, 5893, 5755, 5756, 5836, 5652, 5971, 6000, 5689, 6110, 5953, 6102, 5747, 5872, 5808, 5891, 5839, 5719, 5865, 6114, 5811, 5687, 5834, 5759, 5873, 6114, 6314, 5757, 5849, 5901, 5993, 5226, 5338, 5356, 5299, 5426, 5479, 5687, 5594, 5497, 5735, 19094, 19217, 19264, 18976, 19040, 19348, 19092, 5777, 5810, 5636, 5600, 5681, 5528, 5573, 5494, 5613, 5603, 5509, 5455, 5552, 5773, 5903, 5385}
  indexes := GetChangePointIndexes(data, 1)
  assert.Equal(t, []int{111, 148, 153, 158, 165}, indexes)
}

func TestWithoutDegradations(t *testing.T) {
  data := []int{99, 85, 68, 70, 67, 82, 87, 68, 93, 72, 85, 86, 84, 71, 95, 67, 85, 97, 86, 87, 85, 97, 99, 101, 75, 86, 107, 87, 97, 120, 75, 80, 70, 78, 75, 92, 97, 98, 98, 115, 85, 110, 120, 85, 70, 80, 105, 104, 74, 78, 100, 109, 98, 107, 74, 72, 69, 66, 86, 76, 66, 75, 84, 85, 93, 74, 95, 98, 71, 70, 74, 98, 109, 75, 120, 72, 86, 72, 73, 80, 102, 85, 86, 74, 75, 74, 98, 100, 110, 85, 90, 105, 95, 113, 71, 98, 111, 76, 69, 106, 97, 78, 67, 86, 109, 88, 94, 111, 86, 71, 72, 85, 74, 66, 84, 86, 67, 67, 75, 90, 78, 101, 67, 75, 81, 87, 74, 84, 95, 89, 96, 74, 74, 71, 114, 98, 87, 87, 107, 111, 86, 71, 75, 74, 88, 83, 86, 62, 77, 98, 74, 77, 87, 76, 90, 110, 78, 112, 113, 100, 101, 86, 102, 85, 69, 67, 74, 87, 98, 98, 94, 71, 72, 73, 85, 85, 99, 97, 96, 72, 94, 85, 110, 75, 69, 89, 73, 85, 90, 108, 76, 73, 94, 105, 103, 86, 96, 67, 102, 95, 112, 69, 89, 96, 71, 69, 89, 84, 74, 88, 87, 78, 110, 97, 72, 64, 99, 98, 110, 118, 74, 77, 109, 73, 74, 102, 71, 88, 87, 99, 98, 73, 101, 105, 85, 69, 69, 95, 73, 79, 89, 73, 101, 71, 86, 100, 86, 69, 71, 66, 71, 84, 100, 66, 71, 96, 71, 96, 101, 87, 86, 85, 65, 99, 81, 95, 98, 98, 75, 70, 72, 73, 74, 71, 84, 73, 95, 85, 106, 70, 71, 67, 69, 85, 74, 84, 108, 100, 95, 84, 67, 85, 86, 88, 71, 91, 71, 101, 70, 69, 68, 76, 73, 72, 86, 87, 73, 86, 73, 96, 87, 71, 88, 70, 95, 67, 98, 75, 72, 72, 95, 95, 68, 76, 84, 96, 73, 95, 97, 95, 63, 85, 100, 70, 110, 85, 100, 73, 97, 99, 100, 94, 83, 93, 85, 72, 102, 76, 70, 96, 107, 82, 78, 74, 98, 113, 71, 85, 87, 70, 101, 74, 98, 72, 98, 98, 73, 87, 102, 88, 86, 99, 70, 86, 75, 86, 74, 67, 111, 72, 96, 99, 75, 75, 70, 70, 71, 95, 70, 86, 92, 109, 96, 91, 88, 70, 97, 69, 74, 64, 103, 97, 84, 71, 97, 102, 80, 76, 112, 76, 99, 69, 74, 69, 90, 77, 86, 107, 96, 68, 99, 89, 89, 85, 73, 88, 85, 84, 75, 82, 95, 68, 98, 90, 94, 85, 86, 84, 85, 73, 94, 97, 95, 74, 85, 73, 107, 99, 72, 70, 75, 88, 87, 85, 98, 97, 84, 91, 71, 70, 75, 88, 97, 70, 95, 77, 66, 76, 109, 74, 84, 69, 81, 76, 87, 72, 97, 101, 109, 85, 98, 84, 97, 75, 112, 108, 73, 96, 73, 84, 84, 73, 86, 70, 69, 71, 73, 85, 67, 101, 97, 91, 74, 75, 97, 82, 73, 73, 85, 97, 61, 70, 72, 85, 88, 71, 67, 85, 65, 68, 68, 98, 84, 73, 87, 71, 80, 77, 90, 100, 71, 120, 69, 85, 87, 82, 85, 96, 128, 86, 71, 69, 102, 85, 85, 110, 71, 74, 73, 88, 72, 94, 97, 85, 74, 67, 111, 73, 96, 84, 94, 94, 66, 66, 75, 75, 74, 74, 94, 94, 74, 74, 69, 69, 73, 73, 86, 86, 66, 66, 95, 95, 73, 73, 88, 88, 97, 97, 86, 86, 86, 86, 88, 88, 68, 68, 69, 69, 86, 86, 85, 85, 73, 73, 73, 73, 103, 103, 99, 99, 97, 97, 72, 72, 83, 83, 70, 70, 98, 98, 96, 96, 106, 106, 85, 85, 108, 108, 114, 114, 86, 86, 96, 96, 77, 77, 97, 97, 102, 102, 89, 89, 107, 107, 99, 99, 101, 101, 96, 96, 113, 113, 112, 112, 98, 98, 83, 83, 72, 72, 74, 74, 96, 96, 84, 84, 91, 91, 97, 97, 73, 73, 88, 88, 83, 83, 97, 97, 114, 114, 84, 84, 75, 75, 89, 89, 75, 75, 94, 94, 87, 87, 86, 86, 101, 101, 96, 96, 138, 138, 104, 104, 121, 121, 71, 71, 99, 99, 72, 72, 72, 72, 87, 87, 95, 95, 73, 73, 71, 71, 82, 82, 74, 74, 88, 88, 106, 106, 74, 74, 85, 85, 74, 74, 85, 85, 86, 86, 88, 88, 103, 103, 99, 99, 91, 91, 101, 101, 100, 100, 85, 85, 73, 73, 83, 83, 95, 95, 71, 71, 73, 73, 75, 75, 87, 87, 74, 74, 77, 77, 102, 102, 91, 91, 81, 81, 79, 79, 99, 99, 79, 79, 94, 94, 76, 76, 78, 78, 84, 84, 81, 81, 90, 90, 93, 93, 95, 95, 79, 79, 81, 81, 85, 85, 87, 87, 87, 87, 85, 85, 83, 83, 81, 81, 88, 88, 88, 88, 75, 75, 88, 88, 84, 84, 74, 74, 74, 113, 74, 84, 75, 99, 85, 97, 85, 73, 92, 72, 86, 86, 109, 83, 101, 68, 73, 73, 86, 68, 101, 108, 79, 96, 93, 97, 71, 93, 88, 94, 78, 95, 83, 82, 93, 92, 86, 101, 81, 81, 81, 148, 81, 77, 92, 81, 94, 83, 91, 93, 89, 82, 97, 83, 95, 87, 77, 106, 94, 90, 84, 98, 93, 97, 97, 92}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := InferDegradations(data, builds, times, CommonAnalysisSettings{})
  assert.Equal(t, 0, len(degradations))
}
func TestWithoutDegradations2(t *testing.T) {
  data := []int{2940, 2633, 2758, 2648, 2884, 2920, 3205, 2936, 2868, 3212, 2324, 2290, 2474, 3000, 2740, 2737, 2413, 2873, 3105, 3049, 2521, 3185, 2950, 2696, 2725, 3188, 3781, 2493, 2241, 2528, 2680, 3126, 3126, 2649, 2649, 2623, 2623, 2820, 2820, 2693, 2693, 2812, 2812, 2527, 2527, 3031, 3031, 2571, 2571, 3066, 3066, 2670, 2670, 2699, 2699, 3106, 2782, 2617}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := InferDegradations(data, builds, times, CommonAnalysisSettings{})
  assert.Equal(t, 0, len(degradations))
}

func TestWithDegradations(t *testing.T) {
  data := []int{100, 100, 100, 100, 100, 200, 200, 200}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := InferDegradations(data, builds, times, CommonAnalysisSettings{MinimumSegmentLength: 3})
  assert.Equal(t, 1, len(degradations))
  assert.Equal(t, int64(5), degradations[0].timestamp)
}

func TestWithoutDegradations3(t *testing.T) {
  data := []int{100, 100, 100, 100, 100, 200, 100, 100, 100}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := InferDegradations(data, builds, times, CommonAnalysisSettings{MinimumSegmentLength: 3})
  assert.Equal(t, 0, len(degradations))
}

func TestWithDegradations2(t *testing.T) {
  data := []int{100, 100, 100, 100, 100, 200, 300, 300, 300}
  builds := make([]string, len(data))

  times := make([]int64, len(data))
  for i := range times {
    times[i] = int64(i)
  }
  degradations := InferDegradations(data, builds, times, CommonAnalysisSettings{MinimumSegmentLength: 3})
  assert.Equal(t, 1, len(degradations))
  assert.Equal(t, int64(5), degradations[0].timestamp)
}
