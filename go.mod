module thingue-launcher

go 1.21

toolchain go1.21.3

require (
	github.com/bluele/gcache v0.0.2
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.3.1
	github.com/gorilla/websocket v1.5.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0
	github.com/jaypipes/ghw v0.12.0
	github.com/jhump/grpctunnel v0.3.0
	github.com/jinzhu/copier v0.3.5
	github.com/labstack/gommon v0.4.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mochi-mqtt/server/v2 v2.4.5
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/soheilhy/cmux v0.1.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.16.0
	github.com/wailsapp/wails/v2 v2.7.1
	go.uber.org/zap v1.21.0
	golang.org/x/sys v0.15.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/sqlite v1.5.4
	gorm.io/gorm v1.25.5
	k8s.io/apimachinery v0.29.0
)

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/bep/debounce v1.2.1 // indirect
	github.com/bytedance/sonic v1.10.0-rc2 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fullstorydev/grpchan v1.1.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jaypipes/pcidb v1.0.0 // indirect
	github.com/jchv/go-winloader v0.0.0-20210711035445-715c2860da7e // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/labstack/echo/v4 v4.10.2 // indirect
	github.com/leaanthony/go-ansi-parser v1.6.0 // indirect
	github.com/leaanthony/gosod v1.0.3 // indirect
	github.com/leaanthony/slicer v1.6.0 // indirect
	github.com/leaanthony/u v1.1.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tkrajina/go-reflector v0.5.6 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/wailsapp/go-webview2 v1.0.10 // indirect
	github.com/wailsapp/mimetype v1.4.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	howett.net/plist v1.0.0 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
)

// replace github.com/wailsapp/wails/v2 v2.5.1 => C:\Users\yangwei\go\pkg\mod
