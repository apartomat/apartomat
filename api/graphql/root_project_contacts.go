package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"log"
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
		log.Printf("can't resolve project contacts: %s", errors.New("unknown project"))

		return serverError()
	} else {
		res, err := r.useCases.GetContacts(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%s) contacts: %s", project.ID, err)

			return ServerError{Message: "internal server error"}, nil
		}

		return ProjectContactsList{Items: contactsToGraphQL(res)}, nil
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
