package graphql

import "context"

func (r *queryResolver) Enums(ctx context.Context) (*Enums, error) {
	return &Enums{
		Project: &ProjectEnums{
			Status: &ProjectStatusEnum{
				Items: []*ProjectStatusEnumItem{
					{ProjectStatusNew, "Новый"},
					{ProjectStatusInProgress, "В работе"},
					{ProjectStatusDone, "Завершен"},
					{ProjectStatusCanceled, "Отменен"},
				},
			},
		},
	}, nil
}
