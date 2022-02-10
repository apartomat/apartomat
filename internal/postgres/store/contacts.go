package store

import (
	"context"
	"errors"
	. "github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	contactsTableName = `apartomat.contacts`
)

type contactsStore struct {
	db *pg.DB
}

func NewContactsStore(db *pg.DB) *contactsStore {
	return &contactsStore{db}
}

var (
	_ Store = (*contactsStore)(nil)
)

func (s *contactsStore) Save(ctx context.Context, contact *Contact) (*Contact, error) {
	if contact.CreatedAt.IsZero() {
		contact.CreatedAt = time.Now()
	}

	if contact.ModifiedAt.IsZero() {
		contact.ModifiedAt = contact.CreatedAt
	}

	rec := toContactsRecord(contact)

	_, err := s.db.ModelContext(ctx, rec).Returning("NULL").Insert()
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *contactsStore) Delete(ctx context.Context, contact *Contact) error {
	_, err := s.db.ModelContext(ctx, (*contactsRecord)(nil)).Where(`id = ?`, contact.ID).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (s *contactsStore) List(ctx context.Context, spec Spec, limit, offset int) ([]*Contact, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	sql, args, err := goqu.From(contactsTableName).Where(expr).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	contacts := make([]*contactsRecord, 0)

	_, err = s.db.QueryContext(ctx, &contacts, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromContactsRecords(contacts), nil
}

type contactsRecord struct {
	tableName  struct{}              `pg:"apartomat.contacts,alias:contacts"`
	ID         string                `pg:"id,pk"`
	FullName   string                `pg:"full_name"`
	Photo      string                `pg:"photo"`
	Details    []contactRecordDetail `pg:"details"`
	CreatedAt  time.Time             `pg:"created_at"`
	ModifiedAt time.Time             `pg:"modified_at"`
	ProjectID  int                   `pg:"project_id"`
}

func toContactsRecord(contact *Contact) *contactsRecord {
	return &contactsRecord{
		ID:         contact.ID,
		FullName:   contact.FullName,
		Photo:      contact.Photo,
		Details:    toContactRecordDetails(contact.Details),
		CreatedAt:  contact.CreatedAt,
		ModifiedAt: contact.ModifiedAt,
		ProjectID:  contact.ProjectID,
	}
}

func fromContactsRecords(records []*contactsRecord) []*Contact {
	contacts := make([]*Contact, len(records))

	for i, r := range records {
		contacts[i] = &Contact{
			ID:         r.ID,
			FullName:   r.FullName,
			Photo:      r.Photo,
			Details:    fromContactRecordsDetails(r.Details),
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.ModifiedAt,
			ProjectID:  r.ProjectID,
		}
	}

	return contacts
}

type contactRecordDetail struct {
	Type  Type   `json:"type"`
	Value string `json:"value"`
}

func toContactRecordDetails(details []Details) []contactRecordDetail {
	res := make([]contactRecordDetail, len(details))

	for i, d := range details {
		res[i] = contactRecordDetail{
			Type:  d.Type,
			Value: d.Value,
		}
	}

	return res
}

func fromContactRecordsDetails(records []contactRecordDetail) []Details {
	res := make([]Details, len(records))

	for i, r := range records {
		res[i] = Details{
			Type:  r.Type,
			Value: r.Value,
		}
	}

	return res
}

//

type specQuery interface {
	Expression() (goqu.Expression, error)
}

func toSpecQuery(spec Spec) (specQuery, error) {
	if s, ok := spec.(specQuery); ok {
		return s, nil
	}

	switch s := spec.(type) {
	case IDInSpec:
		return contactIDInSpecQuery{s}, nil
	case ProjectIDInSpec:
		return contactProjectIDInSpecQuery{s}, nil
	case AndSpec:
		return andSpecQuery{spec: s}, nil
	case OrSpec:
		return orSpecQuery{spec: s}, nil
	}

	return nil, errors.New("unknown spec")
}

//
type andSpecQuery struct {
	spec AndSpec
}

func (s andSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toSpecQuery(spec); err != nil {
			return nil, err
		} else {
			expr, err := ps.Expression()
			if err != nil {
				return nil, err
			}

			exs = append(exs, expr)
		}
	}

	return goqu.And(exs...), nil
}

//
type orSpecQuery struct {
	spec OrSpec
}

func (s orSpecQuery) Expression() (goqu.Expression, error) {
	exs := make([]goqu.Expression, 0, len(s.spec.Specs))

	for _, spec := range s.spec.Specs {
		if ps, err := toSpecQuery(spec); err != nil {
			return nil, err
		} else {
			expr, err := ps.Expression()
			if err != nil {
				return nil, err
			}

			exs = append(exs, expr)
		}
	}

	return goqu.Or(exs...), nil
}

//
type contactIDInSpecQuery struct {
	spec IDInSpec
}

func (s contactIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"id": s.spec.IDs}, nil
}

//
type contactProjectIDInSpecQuery struct {
	spec ProjectIDInSpec
}

func (s contactProjectIDInSpecQuery) Expression() (goqu.Expression, error) {
	return goqu.Ex{"project_id": s.spec.IDs}, nil
}
