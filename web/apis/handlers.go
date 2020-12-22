package apis

import (
	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	registerAccountHandlers(app)
	registerEnvelopeHandlers(app)
}
