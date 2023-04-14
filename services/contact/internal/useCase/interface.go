package useCase

import (
	"ol-ilyassov/clean_arch/pkg/type/context"
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

// Create func with the option of creating several contacts is Further Dev option: example = import/export contacts.
// Delivery will have several implementations (create one contact/several contacts), But UseCase will have only one correct interface for implementations.

// The struct queryParameter could have store filter option.
// Filter should be separated struct and could be used for Count methods, etc...

type Contact interface {
	Create(c context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error)
	Update(c context.Context, contact contact.Contact) (*contact.Contact, error)

	// addition of filter use in Delete method, will be good option.
	Delete(c context.Context, ID uuid.UUID /*Тут можно передавать фильтр*/) error

	ContactReader
}

type ContactReader interface {
	List(c context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error)
	ReadByID(c context.Context, ID uuid.UUID) (response *contact.Contact, err error)
	Count(c context.Context /*Тут можно передавать фильтр*/) (uint64, error)
}

// -------------------

type Group interface {
	Create(c context.Context, domainGroup *group.Group) (*group.Group, error)
	Update(c context.Context, group *group.Group) (*group.Group, error)
	Delete(c context.Context, ID uuid.UUID /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	List(c context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error)
	ReadByID(c context.Context, ID uuid.UUID) (*group.Group, error)
	Count(c context.Context /*Тут можно передавать фильтр*/) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(c context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(c context.Context, groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(c context.Context, groupID, contactID uuid.UUID) error
}
