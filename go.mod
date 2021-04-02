module gin-server

go 1.16

require (
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/ilibs/gosql/v2 v2.1.0
	github.com/maotan/go-truffle v1.1.6
	github.com/sirupsen/logrus v1.7.0
	gopkg.in/resty.v1 v1.12.0 // indirect
)

replace github.com/maotan/go-truffle => ../go-truffle
