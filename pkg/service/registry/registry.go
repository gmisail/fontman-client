package registry

import (
	"fontman/client/pkg/api"
	"fontman/client/pkg/model"
)

func UploadRegistryFile(file model.RegistryFile, baseUrl string) error {
	return api.UploadRegistryFile(file, baseUrl)
}
