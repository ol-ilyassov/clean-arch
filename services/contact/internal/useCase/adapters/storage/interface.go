package storage

import (
	"ol-ilyassov/clean_arch/pkg/type/context"
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

// By use on the base of CRUD, they could be similar to UseCase interface.
// But, there could be a difference. Example: UpdateContact, UpdateGroup
// UpdateContact func:
// In one transaction checks the existence of the Contact, and on level of useCase,
// set which fields will be updated of this contact.

// Overall, Storage interface will be used to send all these methods in once.
type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error)
	UpdateContact(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error)
	DeleteContact(ctx context.Context, ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	ListContact(ctx context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error)
	ReadContactByID(ctx context.Context, ID uuid.UUID) (response *contact.Contact, err error)
	CountContact(ctx context.Context /* Тут можно передавать фильтр */) (uint64, error)
}

type Group interface {
	CreateGroup(ctx context.Context, group *group.Group) (*group.Group, error)
	UpdateGroup(ctx context.Context, ID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error)
	DeleteGroup(ctx context.Context, ID uuid.UUID /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	ListGroup(ctx context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error)
	ReadGroupByID(ctx context.Context, ID uuid.UUID) (*group.Group, error)
	CountGroup(ctx context.Context /*Тут можно передавать фильтр*/) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error)
	DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error
	AddContactsToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error
}
