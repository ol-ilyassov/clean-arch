package contact

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"ol-ilyassov/clean_arch/pkg/type/email"
	"ol-ilyassov/clean_arch/pkg/type/gender"
	"ol-ilyassov/clean_arch/pkg/type/phoneNumber"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/age"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/name"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/patronymic"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/surname"
)

var (
	ErrPhoneNumberRequired = errors.New("phone number is required")
)

// custom types could be in order to store specific behaviour and checks.
// These custom types are defined by Domain Driven Design.
// the use of simple type primitives is able.
// in case of validation simple types, the constructor should store the validation code.

// name, surname, patronymic, age - (business case validators) unique fields (types) connected only with contacts.
// phoneNumber, email, gender - universal fields (types) that could be used in other domains.

// one more approach - more complex:
// use interfaces{} as types. These types would have concrete structs that implements interfaces and could be used of interface values.

// Small Summary:
// Type #1: Primitive type values. Validation code on constructor.
// Type #2: Domain Driven Design (Custom types).
// Type #3: Types are stored in interfaces with implementation on concrete struct.

// Addition: Code Designer said, that it is not valid to use json tags for further maybe use on the side.
// It is valid to use tags in case of external data validators.

type Contact struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	phoneNumber phoneNumber.PhoneNumber
	email       email.Email

	name       name.Name
	surname    surname.Surname
	patronymic patronymic.Patronymic

	age    age.Age
	gender gender.Gender
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	phoneNumber phoneNumber.PhoneNumber,
	email email.Email,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {
	if phoneNumber.IsEmpty() {
		return nil, ErrPhoneNumberRequired
	}
	if id == uuid.Nil {
		id = uuid.New()
	}
	return &Contact{
		id:          id,
		createdAt:   createdAt.UTC(),
		modifiedAt:  modifiedAt.UTC(),
		phoneNumber: phoneNumber,
		email:       email,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
	}, nil
}

func New(
	phoneNumber phoneNumber.PhoneNumber,
	email email.Email,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {
	if phoneNumber.IsEmpty() {
		return nil, ErrPhoneNumberRequired
	}

	var timeNow = time.Now().UTC()
	return &Contact{
		id:          uuid.New(),
		createdAt:   timeNow,
		modifiedAt:  timeNow,
		phoneNumber: phoneNumber,
		email:       email,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
	}, nil
}

func (c Contact) ID() uuid.UUID {
	return c.id
}

func (c Contact) CreatedAt() time.Time {
	return c.createdAt
}

func (c Contact) ModifiedAt() time.Time {
	return c.modifiedAt
}

func (c Contact) Email() email.Email {
	return c.email
}

func (c Contact) PhoneNumber() phoneNumber.PhoneNumber {
	return c.phoneNumber
}

func (c Contact) Name() name.Name {
	return c.name
}

func (c Contact) Surname() surname.Surname {
	return c.surname
}

func (c Contact) Patronymic() patronymic.Patronymic {
	return c.patronymic
}

func (c Contact) FullName() string {
	return fmt.Sprintf("%s %s %s", c.surname, c.name, c.patronymic)
}

func (c Contact) Age() age.Age {
	return c.age
}

func (c Contact) Gender() gender.Gender {
	return c.gender
}

func (c Contact) Equal(contact Contact) bool {
	return c.id == contact.id
}
