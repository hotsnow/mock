module app

replace github.com/hotsnow/api => ./src/github.com/hotsnow/api

replace github.com/hotsnow/mocks => ./src/github.com/hotsnow/mocks

go 1.13

require (
	github.com/golang/mock v1.3.1
	github.com/hotsnow/api v0.0.0-00010101000000-000000000000
	github.com/hotsnow/mocks v0.0.0-00010101000000-000000000000
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
)
