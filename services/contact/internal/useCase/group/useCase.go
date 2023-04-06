package group

import (
	"ol-ilyassov/clean_arch/services/contact/internal/useCase/adapters/storage"
)

// concrete struct implements interface of UseCase/group interface

type UseCase struct {
	adapterStorage storage.Group
	options        Options
}

type Options struct{}

func New(storage storage.Group, options Options) *UseCase {
	var uc = &UseCase{
		adapterStorage: storage,
	}
	uc.SetOptions(options)
	return uc
}

func (uc *UseCase) SetOptions(options Options) {
	if uc.options != options {
		uc.options = options
	}
}
