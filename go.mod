module github.com/aptible/supercronic

go 1.14

require (
	github.com/beorn7/perks v1.0.0
	github.com/certifi/gocertifi v0.0.0-20180118203423-deb3ae2ef261
	github.com/davecgh/go-spew v1.1.1
	github.com/evalphobia/logrus_sentry v0.4.6
	github.com/getsentry/raven-go v0.0.0-20180827214142-a9457d81ec91
	github.com/golang/protobuf v1.3.1
	github.com/gorhill/cronexpr v0.0.0-00010101000000-000000000000
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v0.9.4-0.20190521223130-b46e6ec51bb1
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90
	github.com/prometheus/common v0.4.0
	github.com/prometheus/procfs v0.0.0-20190521135221-be78308d8a4f
	github.com/sirupsen/logrus v1.2.0
	github.com/stretchr/testify v1.2.2
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/sys v0.0.0-20181116152217-5ac8a444bdc5
)

replace github.com/gorhill/cronexpr => github.com/krallin/cronexpr v0.0.0-20180801141017-b648cc9a908c
