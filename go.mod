module github.com/orzkratos/errkratos

go 1.22.8

require (
	github.com/go-kratos/kratos/v2 v2.8.4
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.10.0
	github.com/yyle88/erero v1.0.23
	github.com/yyle88/must v0.0.25
	github.com/yyle88/zaplog v0.0.26
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/yyle88/mutexmap v1.0.14 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.71.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract [v0.0.0, v0.0.14] // rename repo-name
