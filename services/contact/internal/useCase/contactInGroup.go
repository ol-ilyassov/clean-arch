package useCase

import (
	"github.com/google/uuid"

	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
)

type ContactInGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
