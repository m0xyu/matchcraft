package matcher

// Matchable はマッチングの対象が満たすべき最小限のインターフェースです。
// アプリ側の構造体にこれを実装させることで、どんなデータでも計算可能になります。
type Matchable interface {
	GetID() string
	GetAttributes() map[string]float64
}

// MatchResult は計算された1つのマッチング結果を表します。
type MatchResult struct {
	ID           string             // 対象の識別子
	Score        float64            // 0.0〜1.0 の適合度（1.0に近いほど高マッチ）
	Contribution map[string]float64 // どの属性がどれくらい貢献したかの内訳（%）
}

// AnalysisReport は複数の結果と、その分析データをまとめた最終的な戻り値です。
type AnalysisReport struct {
	BestMatch      MatchResult   // 1位の結果
	Alternatives   []MatchResult // 2位以降の候補
	Differentiator string
}

// Calculator は計算アルゴリズムを入れ替え可能にするためのインターフェースです。
type Calculator interface {
	Calculate(user, target map[string]float64) (score float64, contribution map[string]float64)
}
