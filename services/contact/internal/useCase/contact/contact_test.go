package contact

import (
	"os"
	"testing"

	"github.com/google/uuid"
	// libraries for mock and testing
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ol-ilyassov/clean_arch/pkg/type/context"
	"ol-ilyassov/clean_arch/pkg/type/email"
	"ol-ilyassov/clean_arch/pkg/type/gender"
	"ol-ilyassov/clean_arch/pkg/type/phoneNumber"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/age"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/name"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/patronymic"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/surname"
	mockStorage "ol-ilyassov/clean_arch/services/contact/internal/repository/storage/mock"
	"ol-ilyassov/clean_arch/services/contact/internal/useCase"
)

var (
	storageRepository = new(mockStorage.Contact)
	ucDialog          *UseCase
	data              = make(map[uuid.UUID]*contact.Contact)
	createContacts    []*contact.Contact
)

func TestMain(m *testing.M) {

	contactAge, _ := age.New(42)
	contactName, _ := name.New("Иван")
	contactSurname, _ := surname.New("Иванов")
	contactPatronymic, _ := patronymic.New("Иванович")
	contactEmail, _ := email.New("ivanii@gmail.com")
	createContact, _ := contact.New(
		*phoneNumber.New("88002002020"),
		contactEmail,
		*contactName,
		*contactSurname,
		*contactPatronymic,
		*contactAge,
		gender.MALE,
	)
	createContacts = append(createContacts, createContact)
	os.Exit(m.Run())
}

func initTestUseCaseContact(t *testing.T) {
	assertion := assert.New(t)
	storageRepository.On("CreateContact",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, contacts ...*contact.Contact) []*contact.Contact {
			assertion.Equal(contacts, createContacts)
			for _, c := range contacts {
				data[c.ID()] = c
			}
			return contacts
		}, func(ctx context.Context, contacts ...*contact.Contact) error {
			return nil
		})

	storageRepository.On("ReadContactByID",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, ID uuid.UUID) *contact.Contact {
			if c, ok := data[ID]; ok {
				return c
			}
			return nil
		}, func(ctx context.Context, ID uuid.UUID) error {
			if _, ok := data[ID]; !ok {
				return useCase.ErrContactNotFound
			}
			return nil
		})

	storageRepository.On("UpdateContact",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) *contact.Contact {
			return nil
		}, func(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) error {
			return nil
		})
}

func TestContact(t *testing.T) {

	initTestUseCaseContact(t)
	ucDialog = New(storageRepository, Options{})

	assertion := assert.New(t)
	t.Run("create contact", func(t *testing.T) {
		var ctx = context.Empty()

		result, err := ucDialog.Create(ctx, createContacts...)
		assertion.NoError(err)
		assertion.Equal(result, createContacts)
	})

	t.Run("get contact", func(t *testing.T) {
		var ctx = context.Empty()

		result, err := ucDialog.ReadByID(ctx, createContacts[0].ID())
		assertion.NoError(err)
		assertion.Equal(result, createContacts[0])
	})
}
