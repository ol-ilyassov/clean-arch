package postgres

import (
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"

	"github.com/google/uuid"
)

func (r *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) UpdateContact(ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListContact(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error) {
	panic("implement me")
}

func (r *Repository) CountContact() (uint64, error) {
	panic("implement me")
}
