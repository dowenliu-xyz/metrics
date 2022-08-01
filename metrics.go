package metrics

// Counter is metrics counter.
type Counter interface {
	With(labels map[string]string) Counter
	Inc()
	Add(delta float64)
}

// Gauge is metrics gauge.
type Gauge interface {
	With(labels map[string]string) Gauge
	Set(value float64)
	Add(delta float64)
	Sub(delta float64)
}

// Observer is metrics observer.
type Observer interface {
	With(labels map[string]string) Observer
	Observe(float64)
}
