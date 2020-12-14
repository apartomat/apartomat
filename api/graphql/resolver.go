package graphql

import (
	apartomat "github.com/ztsu/apartomat/internal"
)

type UseCases struct {
	LoginByEmail        *apartomat.LoginByEmail
	ConfirmLogin        *apartomat.ConfirmLogin
	CheckAuthToken      *apartomat.CheckAuthToken
	GetUserProfile      *apartomat.GetUserProfile
	GetDefaultWorkspace *apartomat.GetDefaultWorkspace
}

type rootResolver struct {
	useCases *UseCases
}

func NewRootResolver(
	useCases *UseCases,
) ResolverRoot {
	return &rootResolver{
		useCases: useCases,
	}
}

func (r *rootResolver) Query() QueryResolver { return &queryResolver{r} }

func (r *rootResolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *rootResolver) ShoppingList() ShoppingListResolver { return &shoppingListResolver{r} }

func (r *rootResolver) UserProfile() UserProfileResolver { return &userProfileResolver{r} }
