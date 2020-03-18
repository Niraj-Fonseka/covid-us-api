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

	//date, _ := strconv.Atoi(time.Now().Format("2006-01-02"))
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
		states += fmt.Sprintf("'%s'\n", stateString) + ","

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

	str := `<!DOCTYPE html>
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
		<div>
			<select id="dropdown" onchange="javascript:handleSelect(this)">
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AK">AK</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AL">AL</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AR">AR</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AZ">AZ</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/CA">CA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/CO">CO</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/CT">CT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/DC">DC</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/DE">DE</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/FL">FL</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/GA">GA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/HI">HI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/IA">IA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/ID">ID</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/IL">IL</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/IN">IN</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/KS">KS</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/KY">KY</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/LA">LA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MA">MA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MD">MD</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/ME">ME</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MI">MI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MN">MN</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MO">MO</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MS">MS</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MT">MT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NC">NC</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/ND">ND</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NE">NE</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NH">NH</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NJ">NJ</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NM">NM</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NV">NV</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NY">NY</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/OH">OH</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/OK">OK</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/OR">OR</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/PA">PA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/RI">RI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/SC">SC</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/SD">SD</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/TN">TN</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/TX">TX</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/UT">UT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/VA">VA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/VT">VT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WA">WA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WI">WI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WV">WV</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WY">WY</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/PR">PR</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/VI">VI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/GU">GU</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MP">MP</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AS">AS</option>
			</select>
		</div> 
		
		<div id="container-death" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<hr>
		<div id="container-positive" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<hr>
		<div id="container-pending" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<hr>
		<div id="container-total" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		
		<script>
		$('#dropdown').change(function(){
			var value = $(this).children("option:selected").val();
			window.open(value+'.html','statename')
		});
	
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
	bt := []byte(fmt.Sprintf(str, dateTitle, states, y1, states, y2, states, y3, states, y4))

	err := ioutil.WriteFile("covid.html", bt, 0644)

	if err != nil {
		log.Fatal(err)
	}

	var s3Manageer S3Manager

	s3Manageer.UploadFile("covid-19-us-dataset", "covid.html")

}

func (g *Graph) RenderStatePage(stateID string, daily []Daily) {

	var dates string
	var y1 string
	var y2 string

	for _, v := range daily {

		dateString := strconv.Itoa(v.Date)

		formatMonthToDateUTC, _ := strconv.Atoi(dateString[4:6]) //because apparently the month is n -1 ?
		dateFormatted := fmt.Sprintf("Date.UTC(%s, %d, %s)", dateString[:4], formatMonthToDateUTC-1, dateString[6:])

		deathsString := strconv.Itoa(v.Death)
		y1 = ",[" + dateFormatted + "," + deathsString + "]" + y1

		positiveString := strconv.Itoa(v.Positive)
		y2 = ",[" + dateFormatted + "," + positiveString + "]" + y2

	}

	y1 = "[" + strings.TrimPrefix(y1, ",") + "]"
	y2 = "[" + strings.TrimPrefix(y2, ",") + "]"
	dates = "[" + strings.TrimPrefix(dates, ",") + "]"

	str := `<!DOCTYPE html>
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

		<div>
			<select id="dropdown" onchange="javascript:handleSelect(this)">
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AK">AK</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AL">AL</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AR">AR</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AZ">AZ</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/CA">CA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/CO">CO</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/CT">CT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/DC">DC</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/DE">DE</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/FL">FL</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/GA">GA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/HI">HI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/IA">IA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/ID">ID</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/IL">IL</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/IN">IN</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/KS">KS</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/KY">KY</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/LA">LA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MA">MA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MD">MD</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/ME">ME</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MI">MI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MN">MN</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MO">MO</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MS">MS</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MT">MT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NC">NC</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/ND">ND</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NE">NE</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NH">NH</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NJ">NJ</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NM">NM</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NV">NV</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/NY">NY</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/OH">OH</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/OK">OK</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/OR">OR</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/PA">PA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/RI">RI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/SC">SC</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/SD">SD</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/TN">TN</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/TX">TX</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/UT">UT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/VA">VA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/VT">VT</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WA">WA</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WI">WI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WV">WV</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/WY">WY</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/PR">PR</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/VI">VI</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/GU">GU</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/MP">MP</option>
				<option value="https://covid-19-us-dataset.s3.amazonaws.com/states/AS">AS</option>
			</select>
		</div> 
		
		<div id="state-trending-deaths" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		
		<div id="state-trending-positive" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
	
		<script>

		$('#dropdown').change(function(){
			var value = $(this).children("option:selected").val();
			window.open(value+'.html','statename')
		});
	
		$(function () {
			$('#state-trending-deaths').highcharts( {
				title:{
					text : 'Deaths Trending'
				},
				chart: {
					type: 'line'
				},
				xAxis: {
					type: 'datetime',
					dateTimeLabelFormats: {
						day: %s
					},
					title:{
						text: 'dates'
					},
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
			$('#state-trending-positive').highcharts( {
				title:{
					text : 'Positive Trending'
				},
				chart: {
					type: 'line'
				},
				xAxis: {
					type: 'datetime',
					dateTimeLabelFormats: {
						day: %s
					},
					title:{
						text: 'dates'
					},
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
			});
		</script>
		</body>
	`
	bt := []byte(fmt.Sprintf(str, "'%y-%m-%e'", y1, "'%y-%m-%e'", y2))

	err := ioutil.WriteFile(fmt.Sprintf("states/%s.html", stateID), bt, 0644)

	if err != nil {
		log.Fatal(err)
	}

	var s3Manageer S3Manager

	s3Manageer.UploadFile("covid-19-us-dataset", fmt.Sprintf("states/%s.html", stateID))

}
