package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"ol-ilyassov/clean_arch/pkg/tools/converter"
	"ol-ilyassov/clean_arch/pkg/type/pagination"
	"ol-ilyassov/clean_arch/pkg/type/phoneNumber"
	"ol-ilyassov/clean_arch/pkg/type/query"
	"ol-ilyassov/clean_arch/pkg/type/queryParameter"
	jsonContact "ol-ilyassov/clean_arch/services/contact/internal/delivery/http/contact"
	jsonGroup "ol-ilyassov/clean_arch/services/contact/internal/delivery/http/group"
	domainContact "ol-ilyassov/clean_arch/services/contact/internal/domain/contact"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/age"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/name"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/patronymic"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/contact/surname"
	domainGroup "ol-ilyassov/clean_arch/services/contact/internal/domain/group"
	"ol-ilyassov/clean_arch/services/contact/internal/domain/group/description"

	// i have added alias, because group name and contact name had conflicts, same package names.
	gName "ol-ilyassov/clean_arch/services/contact/internal/domain/group/name"
)

var mappingSortsGroup = query.SortsOptions{
	"id":           {},
	"name":         {},
	"description":  {},
	"contactCount": {},
}

func (d *Delivery) CreateGroup(c *gin.Context) {
	var group = &jsonGroup.ShortGroup{}

	if err := c.ShouldBindJSON(&group); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groupName, err := gName.New(group.Name)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	groupDescription, err := description.New(group.Description)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	newGroup, err := d.ucGroup.Create(domainGroup.New(
		groupName,
		groupDescription,
	))
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jsonGroup.GroupResponse{
		ID:         newGroup.ID().String(),
		CreatedAt:  newGroup.CreatedAt(),
		ModifiedAt: newGroup.ModifiedAt(),
		Group: jsonGroup.Group{
			ShortGroup: jsonGroup.ShortGroup{
				Name:        newGroup.Name().Value(),
				Description: newGroup.Description().Value(),
			},
			ContactsAmount: newGroup.ContactCount(),
		},
	})
}

func (d *Delivery) UpdateGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	group := jsonGroup.ShortGroup{}
	if err := c.ShouldBindJSON(&group); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groupName, err := gName.New(group.Name)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	groupDescription, err := description.New(group.Description)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	response, err := d.ucGroup.Update(domainGroup.NewWithID(
		converter.StringToUUID(id.Value),
		time.Now().UTC(),
		time.Now().UTC(),
		groupName,
		groupDescription,
		0,
	))
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, jsonGroup.ProtoToGroupResponse(response))
}

func (d *Delivery) DeleteGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	if err := d.ucGroup.Delete(converter.StringToUUID(id.Value)); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (d *Delivery) ListGroup(c *gin.Context) {
	params, err := query.ParseQuery(c, query.Options{
		Sorts: mappingSortsGroup,
	})

	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groups, err := d.ucGroup.List(queryParameter.QueryParameter{
		Sorts: params.Sorts,
		Pagination: pagination.Pagination{
			Limit:  params.Limit,
			Offset: params.Offset,
		},
	})
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	count, err := d.ucContact.Count()
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	var list = make([]*jsonGroup.GroupResponse, len(groups))

	for i, elem := range groups {
		list[i] = jsonGroup.ProtoToGroupResponse(elem)
	}

	c.JSON(http.StatusOK, jsonGroup.GroupList{
		Total:  count,
		Limit:  params.Limit,
		Offset: params.Offset,
		List:   list,
	})
}

func (d *Delivery) ReadGroupByID(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	response, err := d.ucGroup.ReadByID(converter.StringToUUID(id.Value))
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jsonGroup.ProtoToGroupResponse(response))
}

func (d *Delivery) CreateContactIntoGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contact := jsonContact.ShortContact{}
	if err := c.ShouldBindJSON(&contact); err != nil {
		SetError(c, http.StatusBadRequest, fmt.Errorf("payload is not correct, Error: %w", err))
		return
	}

	contactAge, err := age.New(contact.Age)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contactName, err := name.New(contact.Name)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contactSurname, err := surname.New(contact.Surname)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contactPatronymic, err := patronymic.New(contact.Patronymic)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	dContact, err := domainContact.New(
		*phoneNumber.New(contact.PhoneNumber),
		contact.Email,
		*contactName,
		*contactSurname,
		*contactPatronymic,
		*contactAge,
		contact.Gender,
	)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contacts, err := d.ucGroup.CreateContactIntoGroup(converter.StringToUUID(id.Value), dContact)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	var list = []*jsonContact.ContactResponse{}
	for _, value := range contacts {
		list = append(list, jsonContact.ToContactResponse(value))
	}

	c.JSON(http.StatusOK, list)
}

func (d *Delivery) AddContactToGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	var contactID jsonGroup.ContactID
	if err := c.ShouldBindUri(&contactID); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	if err := d.ucGroup.AddContactToGroup(converter.StringToUUID(id.Value), converter.StringToUUID(contactID.Value)); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (d *Delivery) DeleteContactFromGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	var contactID jsonGroup.ContactID
	if err := c.ShouldBindUri(&contactID); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	if err := d.ucGroup.DeleteContactFromGroup(converter.StringToUUID(id.Value), converter.StringToUUID(contactID.Value)); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
