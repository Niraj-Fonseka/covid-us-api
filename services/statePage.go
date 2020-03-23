package services

import (
	"covid-us-api/file"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type StatePage struct {
	CovidService *Covid
	CacheService *Cache
}

func (c *StatePage) BuildPage() error {

	header := c.GenerateHeader()
	imports := c.GenerateImports()
	upperBody := c.GenerateBody()
	chartScript := c.GenerateChart("data")
	styles := c.GenerateStyle()

	generatedData, err := c.CacheService.GetDailyStateRecords()

	if err != nil {
		return err
	}

	for state, stateValues := range generatedData.StateData {
		normalizedData, err := c.GenerateData(stateValues, generatedData.LastUpdated)
		if err != nil {
			log.Println(err)
			continue
		}
		bodyDataInjected := fmt.Sprintf(upperBody, state, normalizedData["lastUpdated"])
		chartScriptDataInjected := fmt.Sprintf(chartScript, "'%y-%m-%e'", normalizedData["deathsArray"], "'%y-%m-%e'", normalizedData["positiveArray"])
		page := fmt.Sprintf(header, imports, styles, bodyDataInjected, chartScriptDataInjected)
		fileName := fmt.Sprintf("%s.html", state)
		err = file.SaveFile(fileName, "states", []byte(page))
		if err != nil {
			log.Printf("Unable to generate page for state : %s , err : %s ", state, err.Error())
			continue
		}

		log.Printf("Generated page for state : %s ", state)

	}

	return nil
}

func (c *StatePage) GenerateData(daily []Daily, lastUpdated string) (map[string]interface{}, error) {

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

	dataStore := make(map[string]interface{})

	dataStore["deathsArray"] = y1
	dataStore["positiveArray"] = y2
	dataStore["datesArray"] = dates
	dataStore["lastUpdated"] = lastUpdated

	return dataStore, nil
}

func (c *StatePage) GenerateHeader() string {

	header := `<!DOCTYPE html>
	<head>
	%s
	</head>

	<body style="background-color:#2A2D34;">
	%s
	%s
	</body>
	`

	return header
}

func (c *StatePage) GenerateImports() string {
	imports := `
	<script src="https://code.jquery.com/jquery-1.11.3.min.js"></script>
	<script src="https://code.highcharts.com/highcharts.js"></script>
	<script src="https://code.highcharts.com/modules/exporting.js"></script>
	<script src="https://code.highcharts.com/themes/dark-unica.js"></script>
	`
	return imports
}

func (c *StatePage) GenerateChart(data ...string) string {

	chart := `
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
	</script>`
	return chart
}

func (c *StatePage) GenerateBody() string {

	body := `
		<div id="state-title">
			<h1> Currently viewing data for : %s</h1>
		</div>

		<div id="date-title">
			<h1> Last updated : %s</h1>
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
	
	`
	return body
}

func (c *StatePage) GenerateStyle() string {
	styles := `<style>

	#date-title{
		text-align: center;
		color: #E0E0E3;
	}
	
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
	`
	return styles
}
