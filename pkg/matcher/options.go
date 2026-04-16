package matcher

// Option defines a functional option for configuring the matching engine.
type Option func(*Engine)

// WithCalculator specifies the calculation algorithm to use.
func WithCalculator(c Calculator) Option {
	return func(e *Engine) {
		e.calculator = c
	}
}
