package matcher

import "math"

type CosineCalculator struct{}

func (c *CosineCalculator) Calculate(user, target map[string]float64) (score float64, contribution map[string]float64) {
	var dotProduct, normUser, normTarget float64
	contribution = make(map[string]float64)

	for key, uVal := range user {
		tVal := target[key]
		dotProduct += uVal * tVal
		normUser += uVal * uVal
		normTarget += tVal * tVal

		// 暫定的な寄与度：各次元の積がどれだけ全体に貢献したか
		contribution[key] = (uVal * tVal)
	}

	if normUser == 0 || normTarget == 0 {
		return 0, contribution
	}

	// コサイン類似度 = (A・B) / (|A|*|B|)
	score = dotProduct / (math.Sqrt(normUser) * math.Sqrt(normTarget))

	// 寄与度を正規化
	for key := range contribution {
		contribution[key] /= dotProduct
	}

	return score, contribution
}
