package services

type PageFramework interface {
	GenerateHeader() string
	GenerateImports() string
	GenerateChart(data ...string) string
	BuildPage() error
	GenerateStyle() string
	GenerateBody() string
}
