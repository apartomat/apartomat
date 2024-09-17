package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"go.uber.org/zap"
)

func (r *rootResolver) ProjectContacts() ProjectContactsResolver {
	return &projectContactsResolver{r}
}

type projectContactsResolver struct {
	*rootResolver
}

func (r *projectContactsResolver) List(
	ctx context.Context,
	obj *ProjectContacts,
	filter ProjectContactsFilter,
	limit int,
	offset int,
) (ProjectContactsListResult, error) {
	if project, ok := graphql.GetFieldContext(ctx).Parent.Parent.Result.(*Project); !ok {
		r.logger.Error("can't resolve project contacts", zap.Error(errors.New("unknown project")))

		return serverError()
	} else {
		items, err := r.useCases.GetContacts(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			r.logger.Error(
				"can't resolve project contacts",
				zap.String("project", project.ID),
				zap.Error(err),
			)

			return serverError()
		}

		return ProjectContactsList{Items: contactsToGraphQL(items)}, nil
	}
}

func (r *projectContactsResolver) Total(
	ctx context.Context,
	obj *ProjectContacts,
	filter ProjectContactsFilter,
) (ProjectContactsTotalResult, error) {
	return notImplementedYetError()
}

func contactsToGraphQL(contacts []*contacts.Contact) []*Contact {
	result := make([]*Contact, 0, len(contacts))

	for _, item := range contacts {
		result = append(result, contactToGraphQL(item))
	}

	return result
}

func contactToGraphQL(contact *contacts.Contact) *Contact {
	return &Contact{
		ID:         contact.ID,
		FullName:   contact.FullName,
		Photo:      contact.Photo,
		Details:    contactDetailsToGraphQL(contact.Details),
		CreatedAt:  contact.CreatedAt,
		ModifiedAt: contact.ModifiedAt,
	}
}

func contactDetailsToGraphQL(details []contacts.Details) []*ContactDetails {
	res := make([]*ContactDetails, len(details))

	for i, item := range details {
		res[i] = &ContactDetails{
			Type:  contactTypeToGraphQL(item.Type),
			Value: item.Value,
		}
	}

	return res
}

func contactTypeToGraphQL(t contacts.Type) ContactType {
	switch t {
	case contacts.TypeInstagram:
		return ContactTypeInstagram
	case contacts.TypePhone:
		return ContactTypePhone
	case contacts.TypeEmail:
		return ContactTypeEmail
	case contacts.TypeWhatsApp:
		return ContactTypeWhatsapp
	case contacts.TypeTelegram:
		return ContactTypeTelegram
	case contacts.TypeUnknown:
		return ContactTypeUnknown
	default:
		return ""
	}
}

func contactTypeFromGraphQL(t ContactType) contacts.Type {
	switch t {
	case ContactTypeInstagram:
		return contacts.TypeInstagram
	case ContactTypePhone:
		return contacts.TypePhone
	case ContactTypeEmail:
		return contacts.TypeEmail
	case ContactTypeWhatsapp:
		return contacts.TypeWhatsApp
	case ContactTypeTelegram:
		return contacts.TypeTelegram
	case ContactTypeUnknown:
		return contacts.TypeUnknown
	default:
		return ""
	}
}
