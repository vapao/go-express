# go-express
like nodejs express

```go
package main

import (
	"github.com/Yooke/go-express"
	"github.com/Yooke/go-logger"
	"net/http"
	"time"
)

func init() {
	express.NewFilter(logFilter).Register("log").Doc("Global logging")
}

func main() {
	router := express.NewRouter()
	router.TopFilter("log")
	router.AddModule(module())
	logger.Fatal(http.ListenAndServe(":80", router))
}

func logFilter(w *express.Response, r *express.Request, f *express.Channel) {
	now := time.Now()
	f.Handle(w, r)
	logger.Infof("%s %s %d %s\n", r.Request.Method, r.Request.URL.Path, w.StatusCode, time.Since(now))
}

func module() *express.Module {
	mde := express.NewModule()
	mde.POST("/login", login)
	return mde
}

func login(w *express.Response, r *express.Request) {
	w.Send("ok")
}
```
