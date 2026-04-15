package matcher

import (
	"testing"
)

// テスト用のデータ構造
type MockVtuber struct {
	ID    string
	Attrs map[string]float64
}

func (m MockVtuber) GetID() string                     { return m.ID }
func (m MockVtuber) GetAttributes() map[string]float64 { return m.Attrs }

func TestEngine_Match(t *testing.T) {
	// energy: min=2, max=10 / cute: min=5, max=10 となるデータ
	targets := []Matchable{
		MockVtuber{ID: "Genki_V", Attrs: map[string]float64{"energy": 10, "cute": 5}},
		MockVtuber{ID: "Iyashi_V", Attrs: map[string]float64{"energy": 2, "cute": 10}},
	}

	engine := NewEngine(targets, WithCalculator(&EuclideanCalculator{}))

	userPref := map[string]float64{"energy": 9, "cute": 6}
	report := engine.Match(userPref)

	if report.BestMatch.ID != "Genki_V" {
		t.Errorf("期待した1位: Genki_V, 実際の1位: %s", report.BestMatch.ID)
	}

	if len(report.BestMatch.Contribution) == 0 {
		t.Error("寄与度が計算されていません")
	}

	if report.Differentiator == "" {
		t.Error("1位と2位の差分分析（Differentiator）が空です")
	}

	t.Logf("成功！ 1位: %s, 2位との決め手属性: %s", report.BestMatch.ID, report.Differentiator)
}

func TestEngine_Scaling(t *testing.T) {
	targets := []Matchable{
		MockVtuber{ID: "Huge_V", Attrs: map[string]float64{"subs": 1000000, "age": 1}},
		MockVtuber{ID: "Small_V", Attrs: map[string]float64{"subs": 100, "age": 5}},
	}

	engine := NewEngine(targets)

	userPref := map[string]float64{"subs": 500000, "age": 5}
	report := engine.Match(userPref)

	if report.BestMatch.ID != "Small_V" {
		t.Errorf("スケーリングが機能していません。期待: Small_V, 実際: %s", report.BestMatch.ID)
	}
}

func TestEngine_Validation(t *testing.T) {
	targets := []Matchable{
		MockVtuber{ID: "V1", Attrs: map[string]float64{"energy": 10}},
		MockVtuber{ID: "V2", Attrs: map[string]float64{"energy": 0}},
	}
	engine := NewEngine(targets, WithCalculator(&EuclideanCalculator{}))

	userPrefNeg := map[string]float64{"energy": -5}
	reportNeg := engine.Match(userPrefNeg)
	// energy:0 とみなされるので、V2 (energy:0) にマッチするはず
	if reportNeg.BestMatch.ID != "V2" {
		t.Errorf("負の値の処理失敗: 期待 V2, 実際 %s", reportNeg.BestMatch.ID)
	}

	userPrefUnknown := map[string]float64{"energy": 10, "unknown_attr": 100}
	reportUnknown := engine.Match(userPrefUnknown)
	if reportUnknown.BestMatch.ID != "V1" {
		t.Errorf("未知の属性の処理失敗: 期待 V1, 実際 %s", reportUnknown.BestMatch.ID)
	}

	userPrefMissing := map[string]float64{} // energyが空
	reportMissing := engine.Match(userPrefMissing)
	// energy:0 とみなされるので V2 にマッチするはず
	if reportMissing.BestMatch.ID != "V2" {
		t.Errorf("属性欠落の補完失敗: 期待 V2, 実際 %s", reportMissing.BestMatch.ID)
	}
}
