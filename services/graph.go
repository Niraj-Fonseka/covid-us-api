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

func (g *Graph) DrawDeathsGraph(data []Daily) {

	//date, _ := strconv.Atoi(time.Now().Format("20060102"))
	date := data[0].Date

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
		<script src="https://code.jquery.com/jquery-1.11.3.min.js"></script>
		<script src="https://code.highcharts.com/highcharts.js"></script>
		<script src="https://code.highcharts.com/modules/exporting.js"></script>
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
					type: 'column'
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
						type: 'column'
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
						type: 'column'
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
						type: 'column'
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

	err := ioutil.WriteFile("covid.html", bt, 0644)

	if err != nil {
		log.Fatal(err)
	}

	var s3Manageer S3Manager

	s3Manageer.UploadFile("covid-19-us-dataset", "covid.html")

}
