module github.com/{{.repository.owner}}/{{.repository.name}}/build

go 1.20

require (
	github.com/golangci/golangci-lint v1.52.2
	golang.org/x/tools v0.7.0
	mvdan.cc/gofumpt v0.4.0
)

