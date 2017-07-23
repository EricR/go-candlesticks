package candlestick_chart

import "time"

type Candle struct {
	Time   time.Time
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
}

type CandlestickChart struct {
	Candles          []*Candle
	Resolution       time.Duration
	TimeSeries       map[time.Time]*Candle
	LastCandle       *Candle
	CurrentCandle    *Candle
	CurrentCandleNew bool
	StartTime        time.Time
	EndTime          time.Time
}

func NewCandlestickChart(res time.Duration) *CandlestickChart {
	return &CandlestickChart{
		Resolution: res,
		Candles:    make([]*Candle, 0),
		TimeSeries: map[time.Time]*Candle{},
	}
}

func NewCandle(ti time.Time, value float64, volume float64) *Candle {
	return &Candle{
		Time:   ti,
		High:   value,
		Low:    value,
		Open:   value,
		Close:  value,
		Volume: volume,
	}
}

func (chart *CandlestickChart) AddCandle(candle *Candle) {
	chart.CurrentCandle = candle
	chart.Candles = append(chart.Candles, candle)
	chart.TimeSeries[candle.Time] = candle

	if candle.Time.Before(chart.StartTime) {
		chart.StartTime = candle.Time
	} else if candle.Time.After(chart.EndTime) {
		chart.EndTime = candle.Time
	}
}

func (chart *CandlestickChart) AddTrade(ti time.Time, value float64, volume float64) {
	var x = ti.Truncate(chart.Resolution)
	var candle = chart.TimeSeries[x]

	if candle != nil {
		candle.Add(value, volume)
		chart.CurrentCandleNew = false
	} else {
		candle = NewCandle(x, value, volume)
		chart.CurrentCandleNew = true
		chart.setLastCandle(candle)

		if chart.LastCandle != nil && x.After(chart.LastCandle.Time.Add(chart.Resolution)) {
			chart.backfill(candle.Time, chart.LastCandle.Close)
		}
		chart.AddCandle(candle)
	}
}

func (chart *CandlestickChart) backfill(x time.Time, value float64) {
	var flatCandle *Candle

	for ti := x; !ti.Equal(chart.LastCandle.Time); ti = ti.Add(-chart.Resolution) {
		if chart.TimeSeries[x] == nil {
			flatCandle = NewCandle(x, value, 0)
			chart.Candles = append(chart.Candles, flatCandle)
			chart.TimeSeries[x] = flatCandle
		}
	}
}

func (chart *CandlestickChart) setLastCandle(candle *Candle) {
	if chart.CurrentCandle == nil {
		chart.LastCandle = candle
	} else {
		chart.LastCandle = chart.CurrentCandle
	}
}

func (candle *Candle) Add(value float64, volume float64) {
	if value > candle.High {
		candle.High = value
	} else if value < candle.Low {
		candle.Low = value
	}

	candle.Volume += volume
	candle.Close = value
}
