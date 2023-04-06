package group

import (
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	panic("implement me")
}
func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) Count() (uint64, error) {
	panic("implement me")
}

// -----------------------------------------

func (uc *UseCase) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	panic("implement me")
}
