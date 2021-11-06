package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/pkg/errors"
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
		contacts, err := r.useCases.GetContacts.Do(
			ctx,
			project.ID,
		)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return Forbidden{}, nil
			}

			log.Printf("can't resolve project (id=%d) contacts: %s", project.ID, err)

			return ServerError{Message: "internal server error"}, nil
		}

		return ProjectContactsList{Items: projectContactsToGraphQL(contacts)}, nil
	}
}

func (r *projectContactsResolver) Total(
	ctx context.Context,
	obj *ProjectContacts,
	filter ProjectContactsFilter,
) (ProjectContactsTotalResult, error) {
	panic("implement me")
}

func projectContactsToGraphQL(contacts []*contacts.Contact) []*Contact {
	result := make([]*Contact, 0, len(contacts))

	for _, item := range contacts {
		result = append(result, projectContactToGraphQL(item))
	}

	return result
}

func projectContactToGraphQL(contact *contacts.Contact) *Contact {
	return &Contact{
		ID:         contact.ID,
		FullName:   contact.FullName,
		Photo:      contact.Photo,
		Details:    projectContactDetailsToGraphQL(contact.Details),
		CreatedAt:  contact.CreatedAt,
		ModifiedAt: contact.ModifiedAt,
	}
}

func projectContactDetailsToGraphQL(details []contacts.Details) []*ContactDetails {
	res := make([]*ContactDetails, len(details))

	for i, item := range details {
		res[i] = &ContactDetails{
			Type:  projectContactTypeToGraphQL(item.Type),
			Value: item.Value,
		}
	}

	return res
}

func projectContactTypeToGraphQL(t contacts.Type) ContactType {
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
	default:
		return ""
	}
}

//func projectFileTypeToGraphQL(t store.ProjectFileType) ProjectFileType {
//	switch t {
//	case store.ProjectFileTypeVisualization:
//		return ProjectFileTypeVisualization
//	case store.ProjectFileTypeNone:
//		return ProjectFileTypeNone
//	default:
//		return ProjectFileTypeNone
//	}
//}
