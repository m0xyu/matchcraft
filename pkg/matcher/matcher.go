package matcher

import (
	"math"
	"sort"
)

// Engine is the core matching engine that handles scaling and score calculation.
type Engine struct {
	targets       []Matchable
	calculator    Calculator
	scaler        *Scaler
	cachedTargets []map[string]float64
}

// NewEngine initializes a new matching engine with the provided targets and options.
func NewEngine(targets []Matchable, opts ...Option) *Engine {
	e := &Engine{
		targets:    targets,
		calculator: &EuclideanCalculator{},
	}

	// オプションの適用
	for _, opt := range opts {
		opt(e)
	}

	// スケーラーを初期化して、対象データをスケーリングする
	e.scaler = NewScaler(targets)

	// スケーリングされた対象データをキャッシュしておく
	for _, t := range targets {
		e.cachedTargets = append(e.cachedTargets, e.scaler.Transform(t.GetAttributes()))
	}

	return e
}

// sanitizeInputs 入力を検査し、定義されていない属性や負の値を0に置き換えます。
func (e *Engine) sanitizeInputs(prefs map[string]float64) map[string]float64 {
	sanitized := make(map[string]float64)

	for attr := range e.scaler.Min {
		val, ok := prefs[attr]
		if !ok || val < 0 {
			// 定義されていない、または負の値が来た場合は 0 としてカウント
			sanitized[attr] = 0
		} else {
			sanitized[attr] = val
		}
	}

	return sanitized
}

// Match はユーザーの入力をもとに、全対象を評価・ソートしてレポートを返します
func (e *Engine) Match(userPrefs map[string]float64) AnalysisReport {
	sanitizedPrefs := e.sanitizeInputs(userPrefs)
	scaledUser := e.scaler.Transform(sanitizedPrefs)
	var results []MatchResult

	// 各対象に対してスコアと貢献度を計算
	for i, target := range e.targets {
		score, contrib := e.calculator.Calculate(scaledUser, e.cachedTargets[i])
		results = append(results, MatchResult{
			ID:           target.GetID(),
			Score:        score,
			Contribution: contrib,
		})
	}

	if len(results) == 0 {
		return AnalysisReport{}
	}

	// スコアでソート（降順）
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})

	report := AnalysisReport{
		BestMatch: results[0],
	}
	if len(results) > 1 {
		report.Alternatives = results[1:]

		first := results[0]
		second := results[1]
		maxDiff := -1.0
		diffAttr := ""

		// 全属性を走査して、差が最大の部分を特定
		for attr, val := range first.Contribution {
			diff := math.Abs(val - second.Contribution[attr])
			if diff > maxDiff {
				maxDiff = diff
				diffAttr = attr
			}
		}
		report.Differentiator = diffAttr
	}

	return report
}
