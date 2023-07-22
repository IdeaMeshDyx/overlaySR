module overlaysr/client

go 1.18

require (
	github.com/cilium/cilium v1.8.10
	github.com/gorilla/websocket v1.5.0
	github.com/pkg/errors v0.8.1
)

require (
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/fsnotify/fsnotify v1.4.10-0.20200417215612-7f4cf4dd2b52 // indirect
	github.com/go-openapi/analysis v0.19.10 // indirect
	github.com/go-openapi/errors v0.19.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.3 // indirect
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/loads v0.19.5 // indirect
	github.com/go-openapi/runtime v0.19.4 // indirect
	github.com/go-openapi/spec v0.19.6 // indirect
	github.com/go-openapi/strfmt v0.19.4 // indirect
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/go-openapi/validate v0.19.5 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.3.0 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sys v0.0.0-20201231184435-2d18734c6014 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	k8s.io/klog v1.0.0 // indirect
)

require (
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496 // indirect
	github.com/fsnotify/fsnotify v1.4.10-0.20200417215612-7f4cf4dd2b52 // indirect
	github.com/go-openapi/analysis v0.19.10 // indirect
	github.com/go-openapi/errors v0.19.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.3 // indirect
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/loads v0.19.5 // indirect
	github.com/go-openapi/runtime v0.19.4 // indirect
	github.com/go-openapi/spec v0.19.6 // indirect
	github.com/go-openapi/strfmt v0.19.4 // indirect
	github.com/go-openapi/swag v0.19.7 // indirect
	github.com/go-openapi/validate v0.19.5 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/spf13/jwalterweatherman v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	go.mongodb.org/mongo-driver v1.3.0 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sys v0.0.0-20201231184435-2d18734c6014 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	k8s.io/klog v1.0.0 // indirect
)

replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	// Using cilium/netlink until XFRM patches merged upstream
	github.com/vishvananda/netlink => github.com/cilium/netlink v0.0.0-20210223023818-d826f2a4c934
	k8s.io/client-go => github.com/cilium/client-go v0.0.0-20210417023617-aeb4c6f1b557
)
