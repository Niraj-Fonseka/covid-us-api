package services

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
}

type record struct {
	Date string
	Open float64
}

func makeRecord(row []string) record {
	open, _ := strconv.ParseFloat(row[1], 64)
	return record{
		Date: row[0],
		Open: open,
	}
}

func (g *Graph) DrawGraph() {
	f, err := os.Open("/home/hungryotter/go/src/covid-us-api/services/table.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	header := `<!DOCTYPE html>
<head>
<script src="http://code.jquery.com/jquery-1.11.3.min.js"></script>
<script src="http://code.highcharts.com/highcharts.js"></script>
<script src="http://code.highcharts.com/modules/exporting.js"></script>
</head>
<body>
<div id="container" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
  <table>
    <thead>
      <tr>
        <th>Date</th>
        <th>Open</th>
      </tr>
    </thead>
    <tbody>
    `

	openValues := []string{}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		record := makeRecord(row)
		fmt.Println(`
      <tr>
        <td>` + record.Date + `</td>
        <td>` + fmt.Sprintf("%.2f", record.Open) + `</td>
      </tr>
      `)

		openValues = append(openValues, fmt.Sprintf("%.2f", record.Open))

	}

	str := `
    </tbody>
  </table>
  <script>
  $(function () {
      $('#container').highcharts({
          title: {
              text: 'Monthly Average Temperature',
              x: -20 //center
          },
          subtitle: {
              text: 'Source: WorldClimate.com',
              x: -20
          },
          xAxis: {
              categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
                  'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
          },
          yAxis: {
              title: {
                  text: 'Temperature (°C)'
              },
              plotLines: [{
                  value: 0,
                  width: 1,
                  color: '#808080'
              }]
          },
          tooltip: {
              valueSuffix: '°C'
          },
          legend: {
              layout: 'vertical',
              align: 'right',
              verticalAlign: 'middle',
              borderWidth: 0
          },
          series: [{
              name: 'Tokyo',
              data: [
` + strings.Join(openValues, ",") + `
              ]
          }]
      });
  });
  </script>
</body>
	`

	bt := []byte(header + str)

	err = ioutil.WriteFile("graph.html", bt, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func (g *Graph) DrawGraphTwo(data []Daily) {

	date := 20200315

	var states string
	var y1 string
	var y2 string

	for _, v := range data {
		if v.Date != date {
			break
		}
		stateString := v.State
		states += fmt.Sprintf("'%s'", stateString) + ","

		deathsString := strconv.Itoa(v.Death)
		y1 += deathsString + ","

		positiveString := strconv.Itoa(v.Positive)
		y2 += positiveString + ","
	}

	states = "[" + strings.TrimSuffix(states, ",") + "]"
	y1 = "[" + strings.TrimSuffix(y1, ",") + "]"
	y2 = "[" + strings.TrimSuffix(y2, ",") + "]"

	fmt.Println(states)
	fmt.Println(y1)
	fmt.Println(y2)

	header := `<!DOCTYPE html>
<head>
<script src="http://code.jquery.com/jquery-1.11.3.min.js"></script>
<script src="http://code.highcharts.com/highcharts.js"></script>
<script src="http://code.highcharts.com/modules/exporting.js"></script>
</head>
<body>
<div id="container" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
    `

	str := `
  <script>
  $(function () {
      $('#container').highcharts({
		chart: {
			zoomType: 'xy'
		},
		title: {
			text: 'Average Monthly Weather Data for Tokyo',
			align: 'left'
		},
		subtitle: {
			text: 'Source: WorldClimate.com',
			align: 'left'
		},
		xAxis: [{
			categories: %s,
			crosshair: true
		}],
		yAxis: [{ // Primary yAxis
			labels: {
				format: '{value}°C',
				style: {
					color: Highcharts.getOptions().colors[2]
				}
			},
			title: {
				text: 'Temperature',
				style: {
					color: Highcharts.getOptions().colors[2]
				}
			},
			opposite: true
	
		}, { // Secondary yAxis
			gridLineWidth: 0,
			title: {
				text: 'Rainfall',
				style: {
					color: Highcharts.getOptions().colors[0]
				}
			},
			labels: {
				format: '{value} mm',
				style: {
					color: Highcharts.getOptions().colors[0]
				}
			}
	
		}],
		tooltip: {
			shared: true
		},
		legend: {
			layout: 'vertical',
			align: 'left',
			x: 80,
			verticalAlign: 'top',
			y: 55,
			floating: true,
			backgroundColor:
				Highcharts.defaultOptions.legend.backgroundColor || // theme
				'rgba(255,255,255,0.25)'
		},
		series: [{
			name: 'Rainfall',
			type: 'column',
			yAxis: 1,
			data: %s,
			tooltip: {
				valueSuffix: 'deaths'
			}
	
		},
		{
			name: 'Temperature',
			type: 'spline',
			data: %s,
			tooltip: {
				valueSuffix: ' °C'
			}
		}],
		responsive: {
			rules: [{
				condition: {
					maxWidth: 500
				},
				chartOptions: {
					legend: {
						floating: false,
						layout: 'horizontal',
						align: 'center',
						verticalAlign: 'bottom',
						x: 0,
						y: 0
					},
					yAxis: [{
						labels: {
							align: 'right',
							x: 0,
							y: -6
						},
						showLastLabel: false
					}, {
						labels: {
							align: 'left',
							x: 0,
							y: -6
						},
						showLastLabel: false
					}, {
						visible: false
					}]
				}
			}]
		}
	});	
  </script>
</body>
	`
	bt := []byte(fmt.Sprintf(header+str, states, y1, y2))

	err := ioutil.WriteFile("graph2.html", bt, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
