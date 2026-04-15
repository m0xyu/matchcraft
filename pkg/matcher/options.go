package matcher

type Option func(*Engine)

// WithCalculator は計算アルゴリズムを指定するためのオプション
func WithCalculator(c Calculator) Option {
	return func(e *Engine) {
		e.calculator = c
	}
}
