package group

import (
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group"
	"time"

	"github.com/google/uuid"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	return uc.adapterStorage.CreateGroup(groupCreate)
}
func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	return uc.adapterStorage.UpdateGroup(groupUpdate.ID(), func(oldGroup *group.Group) (*group.Group, error) {
		return group.NewWithID(oldGroup.ID(), oldGroup.CreatedAt(), time.Now().UTC(), groupUpdate.Name(), groupUpdate.Description(), oldGroup.ContactCount()), nil
	})
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteGroup(ID)
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	return uc.adapterStorage.ListGroup(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	return uc.adapterStorage.ReadGroupByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountGroup()
}

// -----------------------------------------

func (uc *UseCase) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return uc.adapterStorage.CreateContactIntoGroup(groupID, contacts...)
}

func (uc *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	return uc.adapterStorage.AddContactsToGroup(groupID, contactID)
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	return uc.adapterStorage.DeleteContactFromGroup(groupID, contactID)
}
