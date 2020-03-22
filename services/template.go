package services

type PageFramework interface {
	GenerateHeader() string
	GenerateImports() string
	GenerateChart() string
	BuildPage() error
	GenerateStyle() string
	GenerateBody() string
}
