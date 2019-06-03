module microbook

go 1.12

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

require (
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.3.0
	github.com/micro/go-web v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20190531175056-4c3a928424d2 // indirect
)
