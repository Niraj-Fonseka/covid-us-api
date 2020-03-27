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

	countyData, err := c.CovidService.GetCountyLevelDataRefactor()
	if err != nil {
		return err
	}

	for state, stateValues := range generatedData.StateData {
		generatedImports := fmt.Sprintf(imports, strings.ToLower(state))

		normalizedData, err := c.GenerateData(stateValues, generatedData.LastUpdated, countyData)
		if err != nil {
			log.Println(err)
			continue
		}
		bodyDataInjected := fmt.Sprintf(upperBody, state, normalizedData["lastUpdated"])
		chartScriptDataInjected := fmt.Sprintf(chartScript, strings.ToLower(state), normalizedData[strings.ToLower(state)+"-death"], "'%y-%m-%e'", normalizedData["deathsArray"], strings.ToLower(state), normalizedData[strings.ToLower(state)+"-positive"], "'%y-%m-%e'", normalizedData["positiveArray"])
		page := fmt.Sprintf(header, generatedImports, styles, bodyDataInjected+chartScriptDataInjected)
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

func generateStateCode(state, code string) string {
	state = strings.ToLower(state)
	shortenedCode := code[2:]

	return fmt.Sprintf("us-%s-%s", state, shortenedCode)
}

func generateCountyName(county, state string) string {
	return fmt.Sprintf("%s , %s", strings.TrimSpace(county), strings.TrimSpace(state))
}

func (c *StatePage) GenerateData(daily []Daily, lastUpdated string, countyData USCountyAll) (map[string]interface{}, error) {

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

	for state, countyList := range countyData.CountyData {
		stateLower := strings.ToLower(state)
		var generatedDataConfirmed string
		var generatedDataDeath string
		for _, county := range countyList {
			if len(county.CountyFIPS) == 2 {
				continue
			}
			code := generateStateCode(stateLower, county.CountyFIPS)
			//generatedName := generateCountyName(county.County, state)

			if county.Confirmed[len(county.Confirmed)-1] > 0 {
				generatedDataConfirmed += fmt.Sprintf("['%s', %d ],", code, county.Confirmed[len(county.Confirmed)-1])
			}

			if county.Deaths[len(county.Deaths)-1] > 0 {
				generatedDataDeath += fmt.Sprintf("['%s', %d ],", code, county.Deaths[len(county.Deaths)-1])
			}

			// generatedDataDeath = append(generatedDataDeath, CountyRecord{Code: code, Name: generatedName, Value: county.Deaths[len(county.Deaths)-1]})
			// generatedDataConfirmed = append(generatedDataConfirmed, CountyRecord{Code: code, Name: generatedName, Value: county.Confirmed[len(county.Confirmed)-1]})

		}
		generatedDataConfirmed = "[" + strings.TrimSuffix(generatedDataConfirmed, ",") + "]"
		generatedDataDeath = "[" + strings.TrimSuffix(generatedDataDeath, ",") + "]"

		dataStore[strings.ToLower(state)+"-death"] = generatedDataDeath
		dataStore[strings.ToLower(state)+"-positive"] = generatedDataConfirmed

	}

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
	<script src="https://code.highcharts.com/maps/highmaps.js"></script>
	<script src="https://code.highcharts.com/highcharts.js"></script>
	<script src="https://code.highcharts.com/modules/exporting.js"></script>
	<script src="https://code.highcharts.com/themes/dark-unica.js"></script>
	<script src="https://code.highcharts.com/maps/modules/offline-exporting.js"></script>
	<script src="https://code.highcharts.com/mapdata/countries/us/us-%s-all.js"></script>
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
		$('#state-trending-deaths-map').highcharts('Map', {
			title:{
				text : 'Deaths by County'
			},
			chart: {
				map: 'countries/us/us-%s-all',
				backgroundColor: 'transparent'
			},
		
			mapNavigation: {
				enabled: true,
				buttonOptions: {
					verticalAlign: 'bottom'
				}
			},
		
			colorAxis: {
				min: 0
			},
		
			series: [{
				data: %s,
				name: 'Deaths',
				states: {
					hover: {
						color: '#BADA55'
					}
				},
				dataLabels: {
					enabled: true,
					format: '{point.name}'
				}
			}]
		});

		$('#state-trending-deaths').highcharts( {
			title:{
				text : 'Deaths'
			},
			chart: {
				type: 'line',
				backgroundColor: '#2A2D34',
				backgroundColor: 'transparent'
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

			yAxis: {
				title:{
					text: 'deaths'
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


		$('#state-trending-positive-map').highcharts('Map', {
			title:{
				text : 'Positive cases by County'
			},
			chart: {
				map: 'countries/us/us-%s-all',
				backgroundColor: 'transparent'
			},
		
			mapNavigation: {
				enabled: true,
				buttonOptions: {
					verticalAlign: 'bottom'
				}
			},
		
			colorAxis: {
				min: 0
			},
		
			series: [{
				data: %s,
				name: 'Positive',
				states: {
					hover: {
						color: '#BADA55'
					}
				},
				dataLabels: {
					enabled: true,
					format: '{point.name}'
				}
			}]
		});

		$('#state-trending-positive').highcharts( {
			title:{
				text : 'Positive Cases'
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

			yAxis: {
				title:{
					text: 'cases'
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
		
		<div id="state-trending-deaths-map" style="min-width: 310px; height: 500px; margin: 0 auto"></div>

		<div id="state-trending-deaths" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<hr>
		<div id="state-trending-positive-map" style="min-width: 310px; height: 500px; margin: 0 auto"></div>

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
		background-color:white;
		width: 200px;
		height: 40px;
		display: block;
		font-size: 15px;
		box-shadow: 3px 3px #888888;
		border: none;
		margin: 0 auto;
		padding-left: 5px;
		border-radius: 5px;
		text-align: center;
		font-weight: 600;
	}


	#goback{
		border-radius: 5px;
		width: 70px;
		height: 30px;
		border: none;
		font-size: 15px;
		background-color: gainsboro;
		box-shadow: 3px 3px #888888;
		cursor:pointer;
		display: block;
		padding-top: 10px;
		margin: 0 auto;
		text-align: center;
		font-weight: 600;
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
