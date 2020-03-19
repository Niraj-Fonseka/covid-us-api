package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
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
		<script src="https://code.highcharts.com/themes/dark-unica.js"></script>

		</head>
		<style>
			#date-title{
				text-align: center;
				color: #E0E0E3;
			}

			#dropdown{
				background-color:gainsboro;
				width: 120px;
				height: 30px;
				display: block;
				font-size: 15px;
				border: none;
				margin: 0 auto;
				border-radius: 5px;
			}

			#dropdownwrapper{
				padding: 10px;
			}
		</style>
		<body style="background-color:#2A2D34;">
		<div id="date-title">
			<h1>%s</h1>
		</div>
		<div id="dropdownwrapper">
			<select id="dropdown" onchange="javascript:handleSelect(this)">
				<option>select state</option>
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
			location.href = value+'.html'
		});
	
		$(function () {
			$('#container-death').highcharts( {
				title:{
					text : 'Deaths'
				},
				chart: {
					type: 'column',
					backgroundColor: '#2A2D34'
				},
				xAxis: {
					title:{
						text: 'Deaths'
					},
					categories: %s
				},
			
				plotOptions: {
					column:{
						borderColor: null
					},
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
					type: 'column',
					backgroundColor: '#2A2D34'
				},
					xAxis: {
						title:{
							text: 'Positive Cases'
						},
						categories: %s
					},

					plotOptions: {
						column:{
							borderColor: null
						},
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
					type: 'column',
					backgroundColor: '#2A2D34'
				},
					xAxis: {
						title:{
							text: 'Pending Results'
						},
						categories: %s
					},
					
				
					plotOptions: {
						column:{
							borderColor: null
						},
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
						type: 'column',
						backgroundColor: '#2A2D34'
					},
					xAxis: {
						title:{
							text: 'Total Cases'
						},
						categories: %s
					},
				
					plotOptions: {
						column:{
							borderColor: null
						},
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
		<script src="https://code.highcharts.com/themes/dark-unica.js"></script>

		</head>
		<style>
			#state-title{
				text-align: center;
				color: #E0E0E3;
			}

			#dropdown{
				background-color:gainsboro;
				width: 120px;
				height: 30px;
				display: block;
				font-size: 15px;
				border: none;
				margin: 0 auto;
				border-radius: 5px;
			}


			#goback{
				border-radius: 5px;
				width: 50px;
				height: 30px;
				border: none;
				font-size: 15px;
				background-color: gainsboro;
				display: block;
				padding-left: 10px;
				padding-top: 10px;
				margin: 0 auto;
			}

			#gobackwrapper{
				padding: 10px;
			}

			#dropdownwrapper{
				padding: 10px;
			}
		</style>
		
		<body style="background-color:#2A2D34;">

		

		<div id="state-title">
			<h1> Currently viewing data for : %s</h1>
		</div>


		<div id="gobackwrapper"> 
			<div id="goback">Home</div>
		</div>


		<div id="dropdownwrapper">
			<select id="dropdown" onchange="javascript:handleSelect(this)">
				<option>select state</option>
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
			location.href = value+'.html'
		});

		$('#goback').click(function(){
			location.href = 'https://covid-19-us-dataset.s3.amazonaws.com/covid.html'
		});
	
		$(function () {
			$('#state-trending-deaths').highcharts( {
				title:{
					text : 'Deaths Trending'
				},
				chart: {
					type: 'line',
					backgroundColor: '#2A2D34'
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
					line: {
            			dataLabels: {
                			enabled: true
            			},
            		enableMouseTracking: false
					},
					
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
					type: 'line',
					backgroundColor: '#2A2D34'
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
					line: {
            			dataLabels: {
                			enabled: true
            			},
            		enableMouseTracking: false
        			},
					
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
	bt := []byte(fmt.Sprintf(str, stateID, "'%y-%m-%e'", y1, "'%y-%m-%e'", y2))

	err := ioutil.WriteFile(fmt.Sprintf("states/%s.html", stateID), bt, 0644)

	if err != nil {
		log.Fatal(err)
	}

	var s3Manageer S3Manager

	s3Manageer.UploadFile("covid-19-us-dataset", fmt.Sprintf("states/%s.html", stateID))
	time.Sleep(2 * time.Second)
}

func (g *Graph) DrawUSMapGraph(data []Daily, summary []Summary) {

	fmt.Println("generating us maps")

	type StateDatapoint struct {
		Value float64 `json:"value"`
		Code  string  `json:"code"`
	}

	var deaths []StateDatapoint
	var positive []StateDatapoint
	var negative []StateDatapoint
	var pending []StateDatapoint
	var total []StateDatapoint

	for _, v := range data {

		date := data[0].Date

		if v.Date != date {
			break
		}

		var deathVal float64
		if v.Death == 0 {
			deathVal = 0.00001
		} else {
			deathVal = float64(v.Death)
		}

		deaths = append(deaths, StateDatapoint{Value: deathVal, Code: v.State})

		var posVal float64
		if v.Positive == 0 {
			posVal = 0.00001
		} else {
			posVal = float64(v.Positive)
		}

		positive = append(positive, StateDatapoint{Value: posVal, Code: v.State})

		var negVal float64
		if v.Death == 0 {
			negVal = 0.00001
		} else {
			negVal = float64(v.Negative)
		}

		negative = append(negative, StateDatapoint{Value: negVal, Code: v.State})

		var penVal float64
		if v.Pending == 0 {
			penVal = 0.00001
		} else {
			penVal = float64(v.Pending)
		}

		pending = append(pending, StateDatapoint{Value: penVal, Code: v.State})

		var totalVal float64
		if v.Total == 0 {
			totalVal = 0.00001
		} else {
			totalVal = float64(v.Total)
		}

		total = append(total, StateDatapoint{Value: totalVal, Code: strings.ToLower(v.State)})

	}

	deathsJson, err := json.Marshal(&deaths)
	if err != nil {
		log.Fatal(err)
	}
	positiveJson, err := json.Marshal(&positive)
	if err != nil {
		log.Fatal(err)
	}
	// negaitveJson, err := json.Marshal(&negative)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pendingJson, err := json.Marshal(&pending)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// totalJson, err := json.Marshal(&total)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	date := data[0].Date

	dateString := strconv.Itoa(date)
	dateTitle := fmt.Sprintf("%s-%s-%s", dateString[:4], dateString[4:6], dateString[6:])

	str := `<!DOCTYPE html>
		<head>
		<script src="https://code.jquery.com/jquery-1.11.3.min.js"></script>
		<script src="https://code.highcharts.com/maps/highmaps.js"></script>
		<script src="https://code.highcharts.com/maps/modules/data.js"></script>
		<script src="https://code.highcharts.com/maps/modules/exporting.js"></script>
		<script src="https://code.highcharts.com/themes/dark-unica.js"></script>
		<script src="https://code.highcharts.com/maps/modules/offline-exporting.js"></script>
		<script src="https://code.highcharts.com/mapdata/countries/us/us-all.js"></script>

		</head>
		<style>
		#date-title{
			text-align: center;
			color: #E0E0E3;
		}

		#dropdown{
			background-color:gainsboro;
			width: 120px;
			height: 30px;
			display: block;
			font-size: 15px;
			border: none;
			margin: 0 auto;
			border-radius: 5px;
		}

		#dropdownwrapper{
			padding: 10px;
		}

		#banner{
			display: flex; /* equal height of the children */
			margin-top: 10px;
			margin-bottom: 10px;
		}

		#positive{
			flex: 1; /* additionally, equal width */
			padding: 1em;
			border:1px solid black;
			height: 100px;
			border-radius: 2px;
			text-align: center;
			border-color: #393e46;
		}

		#category{
			font-size: 40px;
			color: #00adb5;
		}
		#negative{
			display: inline-block;
		}

		#pending{
			display: inline-block;
		}

		#death{
			display: inline-block;
		}

		#total{
			display: inline-block;
		}

		#number{
			flex: 1; /* additionally, equal width */
			font-size: 30px;
			margin-top: 10px;
			color: #eeeeee;

		}
		</style>

		
		<body style="background-color:#2A2D34;">
		<div id="date-title">
			<h1>%s</h1>
		</div>
		<div id="dropdownwrapper">
			<select id="dropdown" onchange="javascript:handleSelect(this)">
				<option>select state</option>
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
		<div id="banner">
		<div id="positive">
		   <div id="category">
			   Positive
		   </div>
		   <div id="number">
				%d
		   </div>
		</div>

		<div id="positive">
			<div id="category">
				Negative
			</div>
			<div id="number">
				%d
			</div>
		</div>

		<div id="positive">
			<div id="category">
				Pending
			</div>
			<div id="number">
				%d
			</div>
		</div>

		<div id="positive">
			<div id="category">
				Deaths
			</div>
			<div id="number">
				%d
			</div>
		</div>

		<div id="positive">
			<div id="category">
				Total
			</div>
			<div id="number">
				%d
			</div>
		</div>
	</div>
		<div id="container-death" style="min-width: 310px; height: 600px; margin: 0 auto"></div>
		<div id="container-positive" style="min-width: 310px; height: 600px; margin: 0 auto"></div>
		<script>


		$('#dropdown').change(function(){
			var value = $(this).children("option:selected").val();
			location.href = value+'.html'
		});

		$('#goback').click(function(){
			location.href = 'https://covid-19-us-dataset.s3.amazonaws.com/covid.html'
		});

		$(function () {
			$('#container-death').highcharts('Map',{

				chart: {
					map: 'countries/us/us-all',
					borderWidth: 1
				},
		
				title: {
					text: 'Deaths'
				},
		
				exporting: {
					sourceWidth: 600,
					sourceHeight: 500
				},
		
				legend: {
					layout: 'horizontal',
					borderWidth: 0,
					backgroundColor: '#2a2a2b',
					itemStyle: {
						color: '#FFF',
						fontSize: '12px'
        			},
					borderRadius: 5,
					floating: true,
					verticalAlign: 'top',
					y: 25
				},
		
				mapNavigation: {
					enabled: true
				},
		
				colorAxis: {
					min: 1,
					type: 'logarithmic',
					minColor: '#EEEEFF',
					maxColor: '#000022',
					stops: [
						[0, '#EFEFFF'],
						[0.67, '#4444FF'],
						[1, '#000022']
					]
				},
		
				series: [{
					animation: {
						duration: 1000
					},
					data: %s,
					joinBy: ['postal-code', 'code'],
					dataLabels: {
						enabled: true,
						color: '#FFFFFF',
						format: '{point.code}'
					},
					name: 'deaths',
					tooltip: {
						headerFormat: '',
						pointFormat: '<span style="font-size:23px">{point.name}: <b style="font-size:30px">{point.value:.1f} </b></span>',
					}
				}]
			});

			$('#container-positive').highcharts('Map',{

				chart: {
					map: 'countries/us/us-all',
					borderWidth: 1
				},
		
				title: {
					text: 'Positive cases'
				},
		
				exporting: {
					sourceWidth: 600,
					sourceHeight: 500
				},
		
				legend: {
					layout: 'horizontal',
					borderWidth: 0,
					backgroundColor: '#2a2a2b',
					borderRadius: 5,
					floating: true,
					verticalAlign: 'top',
					y: 25
				},
		
				mapNavigation: {
					enabled: true
				},
		
				colorAxis: {
					min: 1,
					type: 'logarithmic',
					minColor: '#EEEEFF',
					maxColor: '#000022',
					stops: [
						[0, '#EFEFFF'],
						[0.67, '#4444FF'],
						[1, '#000022']
					]
				},
		
				series: [{
					animation: {
						duration: 1000
					},
					data: %s,
					joinBy: ['postal-code', 'code'],
					dataLabels: {
						enabled: true,
						color: '#FFFFFF',
						format: '{point.code}'
					},
					name: 'positive',
					tooltip: {
						headerFormat: '',
						pointFormat: '<span style="font-size:23px">{point.name}: <b style="font-size:30px">{point.value:.1f} </b></span>',
					}
				}]
			});
			});
		</script>
		</body>
	`

	bt := []byte(fmt.Sprintf(str, dateTitle, summary[0].Positive, summary[0].Negative, summary[0].Pending, summary[0].Death, summary[0].Total, deathsJson, positiveJson))

	err = ioutil.WriteFile("covid.html", bt, 0644)

	if err != nil {
		log.Fatal(err)
	}

	var s3Manageer S3Manager

	s3Manageer.UploadFile("covid-19-us-dataset", "covid.html")
}
