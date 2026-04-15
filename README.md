# MatchCraft

MatchCraftは、Go言語で実装された軽量かつ堅牢な多次元データマッチングエンジンです。複数の属性を持つデータ（例：商品、キャラクター、ユーザープロファイル）を比較し、最適なマッチングを算出します。

## 特徴

- **自動スケーリング (Normalization)**: 異なる単位（例：登録者数と活動年数）を持つデータを自動的に `0.0` 〜 `1.0` の範囲に正規化し、公平に評価します。
- **アルゴリズムの切り替え (Strategy Pattern)**: ユークリッド距離やコサイン類似度など、用途に合わせて計算ロジックをオプションで容易に変更できます。
- **堅牢なバリデーション**: 入力データに欠損があったり、負の値が含まれていたりしても、内部で適切にサニタイズして計算を継続します。
- **分析レポート**: 単なる順位付けだけでなく、各属性がマッチングにどれだけ貢献したか（Contribution）や、1位と2位の違い（Differentiator）を提示します。

## インストール

```bash
go get github.com/m0xyu/matchcraft
```

## クイックスタート

以下は、コーヒー豆の属性データを使用して最適な一杯を見つける例です。

```go
package main

import (
	"fmt"
	"log"
	"github.com/m0xyu/matchcraft/pkg/matcher"
)

func main() {
	// JSONからターゲットデータを読み込み
	targets, err := matcher.LoadFromJSON("examples/coffee/data.json")
	if err != nil {
		log.Fatal(err)
	}

	// エンジンの初期化
	engine := matcher.NewEngine(targets)

	// ユーザーの好みを定義
	userPrefs := map[string]float64{
		"bitter": 1, // 苦味は控えめ
		"acid":   8, // 酸味は強め
	}

	// マッチングの実行
	report := engine.Match(userPrefs)

	// 結果の表示
	fmt.Printf("🏆 おすすめ: %s (マッチ度: %.1f%%)\n",
		report.BestMatch.ID, report.BestMatch.Score*100)
	fmt.Printf("💡 決め手となった属性: %s\n", report.Differentiator)
}
```

## プロジェクト構造

```text
.
├── examples/           # 具体的な使用例
│   └── coffee/         # コーヒー豆診断のデモ
├── pkg/
│   └── matcher/        # ライブラリのコアロジック
│       ├── cosine.go   # コサイン類似度計算
│       ├── euclidean.go # ユークリッド距離計算
│       ├── loader.go    # JSONロード機能
│       ├── matcher.go   # メインエンジン
│       ├── scaler.go    # データ正規化ロジック
│       └── types.go     # インターフェース定義
└── go.mod              # モジュール定義
```

## 動作環境

- Go 1.26.1 以上
