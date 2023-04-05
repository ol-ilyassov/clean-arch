package useCase

import (
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

type Group interface {
	Create(domainGroup *group.Group) (*group.Group, error)
	Update(group *group.Group) (*group.Group, error)
	Delete(ID uuid.UUID /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	List(parameter queryParameter.QueryParameter) ([]*group.Group, error)
	ReadByID(ID uuid.UUID) (*group.Group, error)
	Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
}
