package contact

import (
	"ol-ilyassov/clean_arch/pkg/type/context"
	"time"

	"github.com/google/uuid"

	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
)

func (uc *UseCase) Create(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return uc.adapterStorage.CreateContact(ctx, contacts...)
}

func (uc *UseCase) Update(ctx context.Context, contactUpdate contact.Contact) (*contact.Contact, error) {
	return uc.adapterStorage.UpdateContact(ctx, contactUpdate.ID(), func(oldContact *contact.Contact) (*contact.Contact, error) {
		return contact.NewWithID(
			oldContact.ID(),
			oldContact.CreatedAt(),
			time.Now().UTC(),
			contactUpdate.PhoneNumber(),
			contactUpdate.Email(),
			contactUpdate.Name(),
			contactUpdate.Surname(),
			contactUpdate.Patronymic(),
			contactUpdate.Age(),
			contactUpdate.Gender(),
		)
	})
}

func (uc *UseCase) Delete(ctx context.Context, ID uuid.UUID) error {
	return uc.adapterStorage.DeleteContact(ctx, ID)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	return uc.adapterStorage.ListContact(ctx, parameter)
}

func (uc *UseCase) ReadByID(ctx context.Context, ID uuid.UUID) (response *contact.Contact, err error) {
	return uc.adapterStorage.ReadContactByID(ctx, ID)
}

func (uc *UseCase) Count(ctx context.Context) (uint64, error) {
	return uc.adapterStorage.CountContact(ctx)
}
