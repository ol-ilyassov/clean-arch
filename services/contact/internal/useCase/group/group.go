package group

import (
	"ol-ilyassov/clean_arch/pkg/type/context"
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group"
	"time"

	"github.com/google/uuid"
)

func (uc *UseCase) Create(ctx context.Context, groupCreate *group.Group) (*group.Group, error) {
	return uc.adapterStorage.CreateGroup(ctx, groupCreate)
}
func (uc *UseCase) Update(ctx context.Context, groupUpdate *group.Group) (*group.Group, error) {
	return uc.adapterStorage.UpdateGroup(ctx, groupUpdate.ID(), func(oldGroup *group.Group) (*group.Group, error) {
		return group.NewWithID(oldGroup.ID(), oldGroup.CreatedAt(), time.Now().UTC(), groupUpdate.Name(), groupUpdate.Description(), oldGroup.ContactCount()), nil
	})
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteGroup(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	return uc.adapterStorage.ListGroup(ctx, parameter)
}

func (uc *UseCase) ReadByID(ctx context.Context, ID uuid.UUID) (*group.Group, error) {
	return uc.adapterStorage.ReadGroupByID(ctx, ID)
}

func (uc *UseCase) Count(ctx context.Context) (uint64, error) {
	return uc.adapterStorage.CountGroup(ctx)
}

// -----------------------------------------

func (uc *UseCase) CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return uc.adapterStorage.CreateContactIntoGroup(ctx, groupID, contacts...)
}

func (uc *UseCase) AddContactToGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	return uc.adapterStorage.AddContactsToGroup(ctx, groupID, contactID)
}

func (uc *UseCase) DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	return uc.adapterStorage.DeleteContactFromGroup(ctx, groupID, contactID)
}
