package registry

import (
	"github.com/gmisail/fontman-client/pkg/api"
	"github.com/gmisail/fontman-client/pkg/model"
)

func UploadRegistryFile(file model.RegistryFile, baseUrl string) error {
	return api.UploadRegistryFile(file, baseUrl)
}
