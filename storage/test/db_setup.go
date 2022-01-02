package test

import (
	"github.com/koepto/gotrue/conf"
	"github.com/koepto/gotrue/storage"
)

func SetupDBConnection(globalConfig *conf.GlobalConfiguration) (*storage.Connection, error) {
	return storage.Dial(globalConfig)
}
