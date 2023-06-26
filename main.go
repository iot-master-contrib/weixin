package weixin

import (
	"embed"
	"encoding/json"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
	"weixin/api"
	"weixin/influx"
	"weixin/internal"
)

func App() *model.App {
	return &model.App{
		Id:   "weixin",
		Name: "weixin",
		Icon: "",
		Entries: []model.AppEntry{
			{
				Path: "app/weixin/history",
				Name: "历史",
			},
			{
				Path: "app/weixin/setting",
				Name: "配置",
			},
		},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
	}
}

//go:embed all:app/weixin
var wwwFiles embed.FS

// @title 微信接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /app/weixin/api/
// @query.collection.format multi
func main() {}
func Startup(app *web.Engine) error {
	influx.Open()
	internal.SubscribeProperty(mqtt.Client)
	api.RegisterRoutes(app.Group("/app/weixin/api"))
	web.RegisterRoutes(app.Group("/app/weixin"), "weixin")
	return nil
}
func Register() error {
	payload, _ := json.Marshal(App())
	return mqtt.Publish("", payload, false, 0)
}
func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("/app/weixin", http.FS(wwwFiles), "", "app/weixin/index.html")
}
func Shutdown() error {
	return nil
}
