package influx

import (
	"github.com/zgwit/iot-master/v3/pkg/config"
	"github.com/zgwit/iot-master/v3/pkg/env"
)

type Options struct {
	Url    string `yaml:"url"`
	Org    string `yaml:"org"`
	Bucket string `yaml:"bucket"`
	Token  string `yaml:"token"`
}

func Default() Options {
	return Options{}
}

var options Options = Default()
var configure = config.AppName() + ".influxdb.yaml"

const ENV = "IOT_MASTER_INFLUXDB_"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func init() {
	options.FromEnv()
}

func (options *Options) FromEnv() {
	options.Url = env.Get(ENV+"URL", options.Url)
	options.Org = env.Get(ENV+"ORG", options.Org)
	options.Bucket = env.Get(ENV+"BUCKET", options.Bucket)
	options.Token = env.Get(ENV+"TOKEN", options.Token)
}

func (options *Options) ToEnv() []string {
	ret := []string{
		ENV + "URL=" + options.Url,
		ENV + "ORG=" + options.Org,
		ENV + "BUCKET=" + options.Bucket,
		ENV + "TOKEN=" + options.Token,
	}
	return ret
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
