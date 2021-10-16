package graphql

import (
	apartomat "github.com/apartomat/apartomat/internal"
)

type UseCases struct {
	LoginByEmail            *apartomat.LoginByEmail
	ConfirmLogin            *apartomat.ConfirmLogin
	CheckAuthToken          *apartomat.CheckAuthToken
	GetUserProfile          *apartomat.GetUserProfile
	GetDefaultWorkspace     *apartomat.GetDefaultWorkspace
	GetWorkspace            *apartomat.GetWorkspace
	GetWorkspaceUsers       *apartomat.GetWorkspaceUsers
	GetWorkspaceUserProfile *apartomat.GetWorkspaceUserProfile
	GetWorkspaceProjects    *apartomat.GetWorkspaceProjects
	GetProject              *apartomat.GetProject
	GetProjectFiles         *apartomat.GetProjectFiles
	UploadProjectFile       *apartomat.UploadProjectFile
	CreateProject           *apartomat.CreateProject
}

type rootResolver struct {
	useCases *UseCases
}

func NewRootResolver(useCases *UseCases) ResolverRoot {
	return &rootResolver{useCases: useCases}
}

func notFound() (NotFound, error) {
	return NotFound{Message: "not found"}, nil
}

func forbidden() (Forbidden, error) {
	return Forbidden{Message: "not found"}, nil
}

func serverError() (ServerError, error) {
	return ServerError{Message: "server error"}, nil
}

func notImplementedYetError() (ServerError, error) {
	return ServerError{Message: "not implemented yet"}, nil
}
