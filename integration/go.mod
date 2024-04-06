module github.com/AlekSi/submodules-test/v2/integration

go 1.22.2

replace github.com/AlekSi/submodules-test/v2 => ../

require (
	github.com/AlekSi/submodules-test/v2 v2.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
