// Package main is an entry point for the coffee recommendation example.
package main

import (
	"fmt"
	"log"

	"github.com/m0xyu/matchcraft/pkg/matcher"
)

type CoffeeBean struct {
	Name   string
	Flavor map[string]float64
}

func (c CoffeeBean) GetID() string                     { return c.Name }
func (c CoffeeBean) GetAttributes() map[string]float64 { return c.Flavor }

func main() {
	//nolint:gocritic
	// targets := []matcher.Matchable{
	// 	CoffeeBean{"イタリアンロースト", map[string]float64{"bitter": 10, "acid": 1, "aroma": 5}},
	// 	CoffeeBean{"エチオピア・シダモ", map[string]float64{"bitter": 2, "acid": 9, "aroma": 10}},
	// }

	targets, err := matcher.LoadFromJSON("examples/coffee/data.json")
	if err != nil {
		log.Fatalf("データの読み込みに失敗しました: %v", err)
	}

	// 単位がバラバラでも自動スケーリングで安心
	engine := matcher.NewEngine(targets)

	// ユーザーの好み
	prefs := map[string]float64{"bitter": 1, "acid": 8}
	report := engine.Match(prefs)

	if report.BestMatch.ID != "" {
		fmt.Printf("☕️ あなたにおすすめの豆: %s\n", report.BestMatch.ID)
		fmt.Printf("💡 決め手となった属性: %s\n", report.Differentiator)
	}

	if len(report.Alternatives) > 0 {
		fmt.Println("🔍 他の候補:")
		for i, alt := range report.Alternatives {
			fmt.Printf("  %d. %s\n", i+2, alt.ID)
		}
	}
}
