package graphql

type ProjectFiles struct {
	Project *ID
	List    ProjectFilesListResult  `json:"list"`
	Total   ProjectFilesTotalResult `json:"total"`
}

func (ProjectFiles) IsProjectFilesResult() {}
