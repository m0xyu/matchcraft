package matcher

import "math"

type EuclideanCalculator struct{}

func (c *EuclideanCalculator) Calculate(user, target map[string]float64) (float64, map[string]float64) {
	var totalDist float64
	diffs := make(map[string]float64)

	// 各属性の距離（差の2乗）を計算
	for key, userVal := range user {
		targetVal := target[key]
		diff := userVal - targetVal
		dist := diff * diff
		diffs[key] = dist
		totalDist += dist
	}

	// 適合度スコアに変換（距離が近いほど1.0に近づく）
	finalScore := 1.0 / (1.0 + math.Sqrt(totalDist))

	// 寄与度の計算（どの属性がマッチを邪魔しなかったか）
	contribution := make(map[string]float64)
	if totalDist == 0 {
		for key := range user {
			contribution[key] = 1.0 / float64(len(user))
		}
	} else {
		if len(user) <= 1 {
			for key := range user {
				contribution[key] = 1.0
			}
		} else {
			for key, dist := range diffs {
				// 差が小さいほど「マッチに貢献した」とみなす
				contribution[key] = (totalDist - dist) / (totalDist * float64(len(user)-1))
			}
		}
	}

	return finalScore, contribution
}
