package services

import (
	"fmt"
	"io/ioutil"
	"log"
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

func (g *Graph) DrawDeathsGraph(data []Daily) {

	date := 20200315

	var states string
	var y1 string
	var y2 string
	var y3 string
	var y4 string

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

		pendingString := strconv.Itoa(v.Pending)
		y3 += pendingString + ","

		totalString := strconv.Itoa(v.Total)
		y4 += totalString + ","

	}

	states = "[" + strings.TrimSuffix(states, ",") + "]"
	y1 = "[" + strings.TrimSuffix(y1, ",") + "]"
	y2 = "[" + strings.TrimSuffix(y2, ",") + "]"
	y3 = "[" + strings.TrimSuffix(y3, ",") + "]"
	y4 = "[" + strings.TrimSuffix(y4, ",") + "]"

	header := `<!DOCTYPE html>
		<head>
		<script src="http://code.jquery.com/jquery-1.11.3.min.js"></script>
		<script src="http://code.highcharts.com/highcharts.js"></script>
		<script src="http://code.highcharts.com/modules/exporting.js"></script>
		</head>
		<body>
		<div id="container-death" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<div id="container-positive" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<div id="container-pending" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<div id="container-total" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
			`

	str := `
		<script>
		$(function () {
			$('#container-death').highcharts( {
				title:{
					text : 'Deaths'
				},
				chart: {
					type: 'area'
				},
				xAxis: {
					categories: %s
				},
			
				plotOptions: {
					series: {
						fillOpacity: 0.1
					}
				},
			
				series: [{
					data: %s
				}]
			});
			
			$('#container-positive').highcharts( {
				title:{
					text : 'Positive'
				},
					chart: {
						type: 'area'
					},
					xAxis: {
						categories: %s
					},
				
					plotOptions: {
						series: {
							fillOpacity: 0.1
						}
					},
				
					series: [{
						data: %s
					}]
				});
			$('#container-pending').highcharts( {
				title:{
					text : 'Pending'
				},
					chart: {
						type: 'area'
					},
					xAxis: {
						categories: %s
					},
				
					plotOptions: {
						series: {
							fillOpacity: 0.1
						}
					},
				
					series: [{
						data: %s
					}]
				});		
				$('#container-total').highcharts( {
					title:{
						text : 'Total'
					},
					chart: {
						type: 'area'
					},
					xAxis: {
						categories: %s
					},
				
					plotOptions: {
						series: {
							fillOpacity: 0.1
						}
					},
				
					series: [{
						data: %s
					}]
				});			
					});
		</script>
		</body>
	`
	bt := []byte(fmt.Sprintf(header+str, states, y1, states, y2, states, y3, states, y4))

	err := ioutil.WriteFile("graph2.html", bt, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
