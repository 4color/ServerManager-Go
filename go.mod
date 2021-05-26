module servermanager

go 1.16

require (
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.1
	github.com/hpcloud/tail v1.0.0
	github.com/maoqide/melody v0.0.0-20200117052833-c8712ffeaea5
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.1
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect

	//找不到时，用这个代替	"github.com/fsnotify/fsnotify"
)
