module github.com/whileW/enze-global

go 1.13

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/etcd-io/etcd v3.3.25+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.3.3
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/jinzhu/gorm v1.9.15
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	google.golang.org/genproto v0.0.0-20200911024640-645f7a48b24f // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
