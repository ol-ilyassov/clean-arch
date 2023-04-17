package contact

import (
	"ol-ilyassov/clean_arch/services/contact/internal/useCase/adapters/storage"

	log "ol-ilyassov/clean_arch/pkg/type/logger"

	"go.uber.org/zap"
)

// concrete struct implements interface of UseCase/contact interface

// one struct that implements several interfaces
// realization of CRUD with transactions.
type UseCase struct {
	adapterStorage storage.Contact
	options        Options
}

type Options struct{}

func New(storage storage.Contact, options Options) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
		log.Info("set new options", zap.Any("options", uc.options))
	}
}
