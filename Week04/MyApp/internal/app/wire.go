// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"MyApp/internal/app/api"
	"MyApp/internal/app/biz"
	"MyApp/internal/app/data"
	"MyApp/internal/app/router"

	"github.com/google/wire"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitEntOrm,
		InitGinEngine,
		biz.BizSet,
		api.APISet,
		router.RouterSet,
		InjectorSet,
		data.ModelSet,
	)
	return new(Injector), nil, nil
}
