package postgres

import (
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) UpdateGroup(ID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) DeleteGroup(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListGroup(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	panic("implement me")
}

func (r *Repository) ReadGroupByID(ID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) CountGroup() (uint64, error) {
	panic("implement me")
}

// ---------------

func (r *Repository) CreateContactIntoGroup(groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) AddContactsToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	panic("implement me")
}
