package matcher

// Scaler is responsible for scaling attributes to a range of 0.0 to 1.0 based on target data.
type Scaler struct {
	Min map[string]float64
	Max map[string]float64
}

// NewScaler calculates the minimum and maximum values for each attribute from the target data.
func NewScaler(targets []Matchable) *Scaler {
	s := &Scaler{
		Min: make(map[string]float64),
		Max: make(map[string]float64),
	}

	for _, t := range targets {
		for attr, val := range t.GetAttributes() {
			if currMin, ok := s.Min[attr]; !ok || val < currMin {
				s.Min[attr] = val
			}
			if currMax, ok := s.Max[attr]; !ok || val > currMax {
				s.Max[attr] = val
			}
		}
	}
	return s
}

// Transform scales the given attributes to a range of 0.0 to 1.0 based on target data.
func (s *Scaler) Transform(attrs map[string]float64) map[string]float64 {
	scaled := make(map[string]float64)
	for attr, val := range attrs {
		vMin, okMin := s.Min[attr]
		vMax, okMax := s.Max[attr]

		// 対象データに存在しない属性は無視する（スケーリングできないため）
		if !okMin || !okMax {
			continue
		}

		if vMax-vMin == 0 {
			// 全て同じ値の場合はスケーリングできないので、0
			scaled[attr] = 0
		} else {
			// 範囲外の値が来ても 0.0〜1.0 に収まるようにスケーリング
			res := (val - vMin) / (vMax - vMin)
			if res < 0 {
				res = 0
			}
			if res > 1 {
				res = 1
			}
			scaled[attr] = res
		}
	}
	return scaled
}
