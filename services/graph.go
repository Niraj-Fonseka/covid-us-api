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

	dateString := strconv.Itoa(date)
	dateTitle := fmt.Sprintf("%s-%s-%s", dateString[:4], dateString[4:6], dateString[6:])
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
		<style>
			#date-title{
				text-align: center;
			}
		</style>
		<body>
		<div id="date-title">
			<h1>%s</h1>
		</div>
		<div id="container-death" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<hr>
		<div id="container-positive" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<hr>
		<div id="container-pending" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<hr>
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
					title:{
						text: 'Deaths'
					},
					categories: %s
				},
			
				plotOptions: {
					series: {
						fillOpacity: 0.1
					}
				},
			
				series: [{
					name: "deaths",
					data: %s
				}]
			});
			
			$('#container-positive').highcharts( {
				title:{
					text : 'Positive Cases'
				},
					chart: {
						type: 'column'
					},
					xAxis: {
						title:{
							text: 'Positive Cases'
						},
						categories: %s
					},

					plotOptions: {
						series: {
							fillOpacity: 0.1
						}
					},
				
					series: [{
						name: "positive",
						data: %s
					}]
				});
			$('#container-pending').highcharts( {
				title:{
					text : 'Pending Results'
				},
					chart: {
						type: 'column'
					},
					xAxis: {
						title:{
							text: 'Pending Results'
						},
						categories: %s
					},
					
				
					plotOptions: {
						series: {
							fillOpacity: 0.1
						}
					},
				
					series: [{
						name: "pending",
						data: %s
					}]
				});		
				$('#container-total').highcharts( {
					title:{
						text : 'Total Cases'
					},
					chart: {
						type: 'column'
					},
					xAxis: {
						title:{
							text: 'Total Cases'
						},
						categories: %s
					},
				
					plotOptions: {
						series: {
							fillOpacity: 0.1
						}
					},
				
					series: [{
						name: "total",
						data: %s
					}]
				});			
					});
		</script>
		</body>
	`
	bt := []byte(fmt.Sprintf(header+str, dateTitle, states, y1, states, y2, states, y3, states, y4))

	err := ioutil.WriteFile("covid.html", bt, 0644)

	if err != nil {
		log.Fatal(err)
	}

	var s3Manageer S3Manager

	s3Manageer.UploadFile("covid-19-us-dataset", "covid.html")

}
