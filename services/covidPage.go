package services

import (
	"covid-us-api/file"
	"encoding/json"
	"fmt"
	"strings"
)

type CovidPage struct {
	CovidService *Covid
	CacheService *Cache
}

func (c *CovidPage) BuildPage() error {

	header := c.GenerateHeader()
	imports := c.GenerateImports()
	upperBody := c.GenerateBody()
	chartScript := c.GenerateChart("data")
	styles := c.GenerateStyle()

	dailyData, err := c.CacheService.GetDailyRecordsByDate(1)
	fmt.Println("GetDailyRecordsByDate", err)
	if err != nil {
		return err
	}

	summaryData, err := c.CacheService.GetOverallRecords()
	fmt.Println("GetOverallRecords", err)

	if err != nil {
		return err
	}

	generatedData, err := c.GenerateData(dailyData, summaryData)
	fmt.Println("GenerateData", err)

	if err != nil {
		return err
	}

	bodyDataInjected := fmt.Sprintf(upperBody, dailyData.LastUpdated, generatedData["summaryPositive"], generatedData["summaryNegative"], generatedData["summaryHospitalized"], generatedData["summaryDeaths"], generatedData["total"])
	chartScriptDataInjected := fmt.Sprintf(chartScript, generatedData["deathsJSON"], generatedData["positiveJSON"])
	//imports , style , upper body , bodyscript
	page := fmt.Sprintf(header, imports, styles, bodyDataInjected+chartScriptDataInjected)

	return file.SaveFile("covid.html", "", []byte(page))
}

func (c *CovidPage) GenerateHeader() string {

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

func (c *CovidPage) GenerateData(data DailyAll, summary SummaryAll) (map[string]interface{}, error) {

	type StateDatapoint struct {
		Value float64 `json:"value"`
		Code  string  `json:"code"`
	}

	var deaths []StateDatapoint
	var positive []StateDatapoint
	var negative []StateDatapoint
	var pending []StateDatapoint
	var total []StateDatapoint

	for _, v := range data.Daily {

		date := data.Daily[0].Date

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
		return nil, err
	}
	positiveJson, err := json.Marshal(&positive)
	if err != nil {
		return nil, err
	}

	lastUpdatedTime := LastUpdated()

	dataStore := make(map[string]interface{})

	dataStore["deathsJSON"] = deathsJson
	dataStore["positiveJSON"] = positiveJson
	dataStore["lastUpdated"] = []byte(lastUpdatedTime)
	dataStore["summaryDeaths"] = summary.Summary[0].Death
	dataStore["summaryPositive"] = summary.Summary[0].Positive
	dataStore["summaryHospitalized"] = summary.Summary[0].Hospitalized
	dataStore["summaryNegative"] = summary.Summary[0].Negative
	dataStore["total"] = summary.Summary[0].Total

	return dataStore, nil

}

func (c *CovidPage) GenerateImports() string {
	imports := `
	<script src="https://code.jquery.com/jquery-1.11.3.min.js"></script>
	<script src="https://code.highcharts.com/maps/highmaps.js"></script>
	<script src="https://code.highcharts.com/maps/modules/data.js"></script>
	<script src="https://code.highcharts.com/maps/modules/exporting.js"></script>
	<script src="https://code.highcharts.com/themes/dark-unica.js"></script>
	<script src="https://code.highcharts.com/maps/modules/offline-exporting.js"></script>
	<script src="https://code.highcharts.com/mapdata/countries/us/us-all.js"></script>
	`

	return imports
}

func (c *CovidPage) GenerateChart(data ...string) string {

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
		$('#container-death').highcharts('Map',{

			chart: {
				map: 'countries/us/us-all',
				backgroundColor: 'transparent'
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
				backgroundColor: 'transparent'
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
	</script>`
	return chart
}

func (c *CovidPage) GenerateBody() string {

	body := `
	<div id="date-title">
		<h1> Last updated : %s</h1>
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
				Hospitalized
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
	<hr>
	<div id="container-positive" style="min-width: 310px; height: 600px; margin: 0 auto"></div>`

	return body
}

func (c *CovidPage) GenerateStyle() string {
	styles := `<style>
	#date-title{
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
		height: 100px;
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
	`
	return styles
}
