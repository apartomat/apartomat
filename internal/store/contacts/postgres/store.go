package postgres

import (
	"context"
	"github.com/apartomat/apartomat/internal/postgres"
	. "github.com/apartomat/apartomat/internal/store/contacts"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-pg/pg/v10"
	"time"
)

const (
	contactsTableName = `apartomat.contacts`
)

type store struct {
	db *pg.DB
}

func NewStore(db *pg.DB) *store {
	return &store{db}
}

var (
	_ Store = (*store)(nil)
)

func (s *store) List(ctx context.Context, spec Spec, sort Sort, limit, offset int) ([]*Contact, error) {
	qs, err := toSpecQuery(spec)
	if err != nil {
		return nil, err
	}

	expr, err := qs.Expression()
	if err != nil {
		return nil, err
	}

	orderExpr := goqu.I("created_at").Asc()

	q := goqu.From(contactsTableName).Where(expr).Order(orderExpr)

	sql, args, err := q.Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	if err != nil {
		return nil, err
	}

	contacts := make([]*record, 0)

	_, err = s.db.QueryContext(postgres.WithQueryContext(ctx, "contacts.List"), &contacts, sql, args...)
	if err != nil {
		return nil, err
	}

	return fromRecords(contacts), nil
}

func (s *store) Get(ctx context.Context, spec Spec) (*Contact, error) {
	res, err := s.List(ctx, spec, SortDefault, 1, 0)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrContactNotFound
	}

	return res[0], nil
}

func (s *store) Save(ctx context.Context, contacts ...*Contact) error {
	recs := toRecords(contacts)

	_, err := s.db.ModelContext(postgres.WithQueryContext(ctx, "contacts.Save"), &recs).
		Returning("NULL").
		OnConflict("(id) DO UPDATE").
		Insert()

	return err
}

func (s *store) Delete(ctx context.Context, contacts ...*Contact) error {
	var (
		ids = make([]string, len(contacts))
	)

	for i, c := range contacts {
		ids[i] = c.ID
	}

	_, err := s.db.ModelContext(postgres.WithQueryContext(ctx, "contacts.Delete"), (*record)(nil)).
		Where(`id IN (?)`, pg.In(ids)).
		Delete()

	return err
}

type record struct {
	tableName  struct{}       `pg:"apartomat.contacts"`
	ID         string         `pg:"id,pk"`
	FullName   string         `pg:"full_name"`
	Photo      string         `pg:"photo"`
	Details    []detailRecord `pg:"details"`
	CreatedAt  time.Time      `pg:"created_at"`
	ModifiedAt time.Time      `pg:"modified_at"`
	ProjectID  string         `pg:"project_id"`
}

func toRecord(contact *Contact) *record {
	return &record{
		ID:         contact.ID,
		FullName:   contact.FullName,
		Photo:      contact.Photo,
		Details:    toDetailRecord(contact.Details),
		CreatedAt:  contact.CreatedAt,
		ModifiedAt: contact.ModifiedAt,
		ProjectID:  contact.ProjectID,
	}
}

type detailRecord struct {
	Type  Type   `json:"type"`
	Value string `json:"value"`
}

func toDetailRecord(details []Details) []detailRecord {
	res := make([]detailRecord, len(details))

	for i, d := range details {
		res[i] = detailRecord{
			Type:  d.Type,
			Value: d.Value,
		}
	}

	return res
}

func toRecords(contacts []*Contact) []*record {
	var (
		res = make([]*record, len(contacts))
	)

	for i, c := range contacts {
		res[i] = toRecord(c)
	}

	return res
}

func fromRecords(records []*record) []*Contact {
	contacts := make([]*Contact, len(records))

	for i, r := range records {
		contacts[i] = &Contact{
			ID:         r.ID,
			FullName:   r.FullName,
			Photo:      r.Photo,
			Details:    fromDetailRecords(r.Details),
			CreatedAt:  r.CreatedAt,
			ModifiedAt: r.ModifiedAt,
			ProjectID:  r.ProjectID,
		}
	}

	return contacts
}

func fromDetailRecords(records []detailRecord) []Details {
	res := make([]Details, len(records))

	for i, r := range records {
		res[i] = Details{
			Type:  r.Type,
			Value: r.Value,
		}
	}

	return res
}
