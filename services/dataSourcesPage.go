package services

import (
	"covid-us-api/file"
	"fmt"
)

type DataSourcesPage struct {
	CovidService *Covid
	CacheService *Cache
}

func (c *DataSourcesPage) BuildPage() error {

	header := c.GenerateHeader()
	imports := c.GenerateImports()
	upperBody := c.GenerateBody()
	styles := c.GenerateStyle()
	chart := c.GenerateChart()

	//imports , style , upper body , bodyscript
	page := fmt.Sprintf(header, imports, styles, upperBody+chart)

	return file.SaveFile("datasources.html", "", []byte(page))
}

func (c *DataSourcesPage) GenerateHeader() string {

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

func (c *DataSourcesPage) GenerateData(data DailyAll, summary SummaryAll) (map[string]interface{}, error) {

	return nil, nil

}

func (c *DataSourcesPage) GenerateImports() string {
	return `
	<script src="https://code.jquery.com/jquery-1.11.3.min.js"></script>
	`
}

func (c *DataSourcesPage) GenerateChart(data ...string) string {

	return `
	<script>
		$('#goback').click(function(){
			location.href = 'https://covid-19-us-dataset.s3.amazonaws.com/covid.html'
		});
	</script>
	`
}

func (c *DataSourcesPage) GenerateBody() string {

	body := `
	<div id="container"> 
		<div id="gobackwrapper"> 
			<div id="goback">Home</div>
		</div>

		<div id="data-sources-title">
			Data Sources
		<div>
		<hr>
		<div id="datasource">
			<a href="https://github.com/CSSEGISandData/COVID-19"> Johns Hopkins dataset </a>
		</div>
		<div id="datasource">
			<a href="https://covidtracking.com/"> covidtracking </a>
		</div>
		<div id="datasource">
			<a href="https://www.nytimes.com/interactive/2020/us/coronavirus-us-cases.html#g-cases-by-county"> nytimes </a>
		</div>
	</div>`

	return body
}

func (c *DataSourcesPage) GenerateStyle() string {
	styles := `<style>
		#container {
			padding : 30px;
		}

		#data-sources-title {
			font-size : 35px;
			color: #00adb5;
			padding : 30px;
		}

		a {
			color: inherit;
			text-decoration: none; /* no underline */
		}

		#datasource {
			padding : 10px;
			color: #E0E0E3;
			text-decoration: none;
			font-size : 20px;
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
	</style>
	`
	return styles
}
