package useCase

import (
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"

	"github.com/google/uuid"
)

// Create func with the option of creating several contacts is Further Dev option: example = import/export contacts.
// Delivery will have several implementations (create one contact/several contacts), But UseCase will have only one correct interface for implementations.

// The struct queryParameter could have store filter option.
// Filter should be separated struct and could be used for Count methods, etc...

type Contact interface {
	Create(contacts ...*contact.Contact) ([]*contact.Contact, error)
	Update(contact contact.Contact) (*contact.Contact, error)

	// addition of filter use in Delete method, will be good option.
	Delete(ID uuid.UUID /*Тут можно передавать фильтр*/) error

	ContactReader
}

type ContactReader interface {
	List(parameter queryParameter.QueryParameter) ([]*contact.Contact, error)
	ReadByID(ID uuid.UUID) (response *contact.Contact, err error)
	Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
