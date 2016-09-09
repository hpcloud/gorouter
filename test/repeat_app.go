package test

import (
	"io"
	"io/ioutil"
	"net/http"

	"code.cloudfoundry.org/gorouter/route"
	"code.cloudfoundry.org/gorouter/test/common"
	"github.com/nats-io/nats"
)

func NewRepeatApp(urls []route.Uri, rPort uint16, mbusClient *nats.Conn, tags map[string]string) *common.TestApp {
	app := common.NewTestApp(urls, rPort, mbusClient, tags, "")
	app.AddHandler("/", greetHandler)

	return app
}

func repeatHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, "Failed to read req body")
	}
	io.WriteString(w, string(reqBody))
}
