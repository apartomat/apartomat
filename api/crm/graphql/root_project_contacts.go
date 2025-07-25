package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/contacts"
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
		slog.ErrorContext(ctx, "can't resolve project contacts", slog.String("err", "unknown project"))

		return serverError()
	} else {
		items, err := r.crm.GetContacts(
			ctx,
			project.ID,
			limit,
			offset,
		)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(
				ctx,
				"can't resolve project contacts",
				slog.String("project", project.ID),
				slog.Any("err", err),
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
