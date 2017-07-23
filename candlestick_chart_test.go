package candlestick_chart

import (
	"testing"
	"time"
)

func TestCandlestickChart(t *testing.T) {
	var chart = NewCandlestickChart(time.Minute)
	var start = time.Date(2009, time.November, 10, 23, 30, 5, 0, time.UTC)

	chart.AddTrade(start, 5, 1)
	chart.AddTrade(start.Add(5*time.Second), 25, 1)
	chart.AddTrade(start.Add(25*time.Second), 3, 1)
	var c1 = chart.Candles[0]

	chart.AddTrade(start.Add(60*time.Second), 12, 5)
	chart.AddTrade(start.Add(70*time.Second), 13, 2)
	var c2 = chart.Candles[1]

	// Intentionally empty data series included here, to test flat candles
	chart.AddTrade(start.Add(240*time.Second), 15, 6)
	var c3 = chart.Candles[2]
	var c4 = chart.Candles[3]

	if !(c1.Volume == 3 && c1.Open == 5 && c1.Close == 3 &&
		c1.High == 25 && c1.Low == 3) {
		t.Log("Got wrong val: %s", c1)
		t.Fail()
	}

	if !(c2.Volume == 7 && c2.Open == 12 && c2.Close == 13 &&
		c2.High == 13 && c2.Low == 12) {
		t.Log("Got wrong val: %s", c2)
		t.Fail()
	}

	if !(c3.Volume == 0 && c3.Open == 13 && c3.Close == 13 &&
		c3.High == 13 && c3.Low == 13) {
		t.Log("Got wrong val: %s", c3)
		t.Fail()
	}

	if !(c4.Volume == 6 && c4.Open == 15 && c4.Close == 15 &&
		c4.High == 15 && c4.Low == 15) {
		t.Log("Got wrong val: %s", c4)
		t.Fail()
	}

	if len(chart.Candles) != 4 {
		t.Log("Got wrong len")
		t.Fail()
	}
}
