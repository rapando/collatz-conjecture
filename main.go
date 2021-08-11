package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/rapando/collatz-conjecture/counter"
)

func main() {
	var n uint64
	var steps, nValues []uint64

	start := time.Now()
	defer func() {
		fmt.Printf("---\nSteps : %4d \tTime %v", len(steps), time.Since(start))
	}()

	fmt.Printf("Enter a positive non-zero integer... ")
	n = 10000000
	steps, nValues = counter.CountSteps(n)

	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "A Line chart to show the collatz-conjecturek",
		Subtitle: "It's extremely easy to use, right?",
	}))

	bar.SetXAxis(steps).
		AddSeries("N-values", generateLineItems(nValues))

	// actually generate bar graph
	f, err := os.Create("./results.html")
	if err != nil {
		log.Println(err)
		return
	}
	bar.Render(f)
}

func generateLineItems(values []uint64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, n := range values {
		items = append(items, opts.LineData{Value: n})
	}
	return items
}
