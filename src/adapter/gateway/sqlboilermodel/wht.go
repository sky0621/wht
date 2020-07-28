// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboilermodel

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// WHT is an object representing the database table.
type WHT struct {
	ID         int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	RecordDate time.Time   `boil:"record_date" json:"record_date" toml:"record_date" yaml:"record_date"`
	Title      null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	CreatedBy  null.String `boil:"created_by" json:"created_by,omitempty" toml:"created_by" yaml:"created_by,omitempty"`
	CreatedAt  time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedBy  null.String `boil:"updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
	UpdatedAt  time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *whtR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L whtL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WHTColumns = struct {
	ID         string
	RecordDate string
	Title      string
	CreatedBy  string
	CreatedAt  string
	UpdatedBy  string
	UpdatedAt  string
}{
	ID:         "id",
	RecordDate: "record_date",
	Title:      "title",
	CreatedBy:  "created_by",
	CreatedAt:  "created_at",
	UpdatedBy:  "updated_by",
	UpdatedAt:  "updated_at",
}

// Generated where

var WHTWhere = struct {
	ID         whereHelperint64
	RecordDate whereHelpertime_Time
	Title      whereHelpernull_String
	CreatedBy  whereHelpernull_String
	CreatedAt  whereHelpertime_Time
	UpdatedBy  whereHelpernull_String
	UpdatedAt  whereHelpertime_Time
}{
	ID:         whereHelperint64{field: "\"wht\".\"id\""},
	RecordDate: whereHelpertime_Time{field: "\"wht\".\"record_date\""},
	Title:      whereHelpernull_String{field: "\"wht\".\"title\""},
	CreatedBy:  whereHelpernull_String{field: "\"wht\".\"created_by\""},
	CreatedAt:  whereHelpertime_Time{field: "\"wht\".\"created_at\""},
	UpdatedBy:  whereHelpernull_String{field: "\"wht\".\"updated_by\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"wht\".\"updated_at\""},
}

// WHTRels is where relationship names are stored.
var WHTRels = struct {
	ContentImages string
	ContentTexts  string
}{
	ContentImages: "ContentImages",
	ContentTexts:  "ContentTexts",
}

// whtR is where relationships are stored.
type whtR struct {
	ContentImages ContentImageSlice `boil:"ContentImages" json:"ContentImages" toml:"ContentImages" yaml:"ContentImages"`
	ContentTexts  ContentTextSlice  `boil:"ContentTexts" json:"ContentTexts" toml:"ContentTexts" yaml:"ContentTexts"`
}

// NewStruct creates a new relationship struct
func (*whtR) NewStruct() *whtR {
	return &whtR{}
}

// whtL is where Load methods for each relationship are stored.
type whtL struct{}

var (
	whtAllColumns            = []string{"id", "record_date", "title", "created_by", "created_at", "updated_by", "updated_at"}
	whtColumnsWithoutDefault = []string{"record_date", "title", "created_by", "updated_by"}
	whtColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	whtPrimaryKeyColumns     = []string{"id"}
)

type (
	// WHTSlice is an alias for a slice of pointers to WHT.
	// This should generally be used opposed to []WHT.
	WHTSlice []*WHT
	// WHTHook is the signature for custom WHT hook methods
	WHTHook func(context.Context, boil.ContextExecutor, *WHT) error

	whtQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	whtType                 = reflect.TypeOf(&WHT{})
	whtMapping              = queries.MakeStructMapping(whtType)
	whtPrimaryKeyMapping, _ = queries.BindMapping(whtType, whtMapping, whtPrimaryKeyColumns)
	whtInsertCacheMut       sync.RWMutex
	whtInsertCache          = make(map[string]insertCache)
	whtUpdateCacheMut       sync.RWMutex
	whtUpdateCache          = make(map[string]updateCache)
	whtUpsertCacheMut       sync.RWMutex
	whtUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var whtBeforeInsertHooks []WHTHook
var whtBeforeUpdateHooks []WHTHook
var whtBeforeDeleteHooks []WHTHook
var whtBeforeUpsertHooks []WHTHook

var whtAfterInsertHooks []WHTHook
var whtAfterSelectHooks []WHTHook
var whtAfterUpdateHooks []WHTHook
var whtAfterDeleteHooks []WHTHook
var whtAfterUpsertHooks []WHTHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *WHT) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *WHT) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *WHT) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *WHT) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *WHT) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *WHT) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *WHT) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *WHT) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *WHT) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range whtAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddWHTHook registers your hook function for all future operations.
func AddWHTHook(hookPoint boil.HookPoint, whtHook WHTHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		whtBeforeInsertHooks = append(whtBeforeInsertHooks, whtHook)
	case boil.BeforeUpdateHook:
		whtBeforeUpdateHooks = append(whtBeforeUpdateHooks, whtHook)
	case boil.BeforeDeleteHook:
		whtBeforeDeleteHooks = append(whtBeforeDeleteHooks, whtHook)
	case boil.BeforeUpsertHook:
		whtBeforeUpsertHooks = append(whtBeforeUpsertHooks, whtHook)
	case boil.AfterInsertHook:
		whtAfterInsertHooks = append(whtAfterInsertHooks, whtHook)
	case boil.AfterSelectHook:
		whtAfterSelectHooks = append(whtAfterSelectHooks, whtHook)
	case boil.AfterUpdateHook:
		whtAfterUpdateHooks = append(whtAfterUpdateHooks, whtHook)
	case boil.AfterDeleteHook:
		whtAfterDeleteHooks = append(whtAfterDeleteHooks, whtHook)
	case boil.AfterUpsertHook:
		whtAfterUpsertHooks = append(whtAfterUpsertHooks, whtHook)
	}
}

// One returns a single wht record from the query.
func (q whtQuery) One(ctx context.Context, exec boil.ContextExecutor) (*WHT, error) {
	o := &WHT{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboilermodel: failed to execute a one query for wht")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all WHT records from the query.
func (q whtQuery) All(ctx context.Context, exec boil.ContextExecutor) (WHTSlice, error) {
	var o []*WHT

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboilermodel: failed to assign all query results to WHT slice")
	}

	if len(whtAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all WHT records in the query.
func (q whtQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to count wht rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q whtQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboilermodel: failed to check if wht exists")
	}

	return count > 0, nil
}

// ContentImages retrieves all the content_image's ContentImages with an executor.
func (o *WHT) ContentImages(mods ...qm.QueryMod) contentImageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"content_image\".\"wht_id\"=?", o.ID),
	)

	query := ContentImages(queryMods...)
	queries.SetFrom(query.Query, "\"content_image\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"content_image\".*"})
	}

	return query
}

// ContentTexts retrieves all the content_text's ContentTexts with an executor.
func (o *WHT) ContentTexts(mods ...qm.QueryMod) contentTextQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"content_text\".\"wht_id\"=?", o.ID),
	)

	query := ContentTexts(queryMods...)
	queries.SetFrom(query.Query, "\"content_text\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"content_text\".*"})
	}

	return query
}

// LoadContentImages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (whtL) LoadContentImages(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWHT interface{}, mods queries.Applicator) error {
	var slice []*WHT
	var object *WHT

	if singular {
		object = maybeWHT.(*WHT)
	} else {
		slice = *maybeWHT.(*[]*WHT)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &whtR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &whtR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`content_image`),
		qm.WhereIn(`content_image.wht_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load content_image")
	}

	var resultSlice []*ContentImage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice content_image")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on content_image")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for content_image")
	}

	if len(contentImageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ContentImages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &contentImageR{}
			}
			foreign.R.WHT = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.WHTID {
				local.R.ContentImages = append(local.R.ContentImages, foreign)
				if foreign.R == nil {
					foreign.R = &contentImageR{}
				}
				foreign.R.WHT = local
				break
			}
		}
	}

	return nil
}

// LoadContentTexts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (whtL) LoadContentTexts(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWHT interface{}, mods queries.Applicator) error {
	var slice []*WHT
	var object *WHT

	if singular {
		object = maybeWHT.(*WHT)
	} else {
		slice = *maybeWHT.(*[]*WHT)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &whtR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &whtR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`content_text`),
		qm.WhereIn(`content_text.wht_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load content_text")
	}

	var resultSlice []*ContentText
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice content_text")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on content_text")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for content_text")
	}

	if len(contentTextAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ContentTexts = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &contentTextR{}
			}
			foreign.R.WHT = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.WHTID {
				local.R.ContentTexts = append(local.R.ContentTexts, foreign)
				if foreign.R == nil {
					foreign.R = &contentTextR{}
				}
				foreign.R.WHT = local
				break
			}
		}
	}

	return nil
}

// AddContentImages adds the given related objects to the existing relationships
// of the wht, optionally inserting them as new records.
// Appends related to o.R.ContentImages.
// Sets related.R.WHT appropriately.
func (o *WHT) AddContentImages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ContentImage) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.WHTID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"content_image\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"wht_id"}),
				strmangle.WhereClause("\"", "\"", 2, contentImagePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.WHTID = o.ID
		}
	}

	if o.R == nil {
		o.R = &whtR{
			ContentImages: related,
		}
	} else {
		o.R.ContentImages = append(o.R.ContentImages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &contentImageR{
				WHT: o,
			}
		} else {
			rel.R.WHT = o
		}
	}
	return nil
}

// AddContentTexts adds the given related objects to the existing relationships
// of the wht, optionally inserting them as new records.
// Appends related to o.R.ContentTexts.
// Sets related.R.WHT appropriately.
func (o *WHT) AddContentTexts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ContentText) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.WHTID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"content_text\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"wht_id"}),
				strmangle.WhereClause("\"", "\"", 2, contentTextPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.WHTID = o.ID
		}
	}

	if o.R == nil {
		o.R = &whtR{
			ContentTexts: related,
		}
	} else {
		o.R.ContentTexts = append(o.R.ContentTexts, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &contentTextR{
				WHT: o,
			}
		} else {
			rel.R.WHT = o
		}
	}
	return nil
}

// WHTS retrieves all the records using an executor.
func WHTS(mods ...qm.QueryMod) whtQuery {
	mods = append(mods, qm.From("\"wht\""))
	return whtQuery{NewQuery(mods...)}
}

// FindWHT retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWHT(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*WHT, error) {
	whtObj := &WHT{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"wht\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, whtObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboilermodel: unable to select from wht")
	}

	return whtObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *WHT) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboilermodel: no wht provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(whtColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	whtInsertCacheMut.RLock()
	cache, cached := whtInsertCache[key]
	whtInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			whtAllColumns,
			whtColumnsWithDefault,
			whtColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(whtType, whtMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(whtType, whtMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"wht\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"wht\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "sqlboilermodel: unable to insert into wht")
	}

	if !cached {
		whtInsertCacheMut.Lock()
		whtInsertCache[key] = cache
		whtInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the WHT.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *WHT) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	whtUpdateCacheMut.RLock()
	cache, cached := whtUpdateCache[key]
	whtUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			whtAllColumns,
			whtPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sqlboilermodel: unable to update wht, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"wht\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, whtPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(whtType, whtMapping, append(wl, whtPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to update wht row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by update for wht")
	}

	if !cached {
		whtUpdateCacheMut.Lock()
		whtUpdateCache[key] = cache
		whtUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q whtQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to update all for wht")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to retrieve rows affected for wht")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WHTSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlboilermodel: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), whtPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"wht\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, whtPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to update all in wht slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to retrieve rows affected all in update all wht")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *WHT) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboilermodel: no wht provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(whtColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	whtUpsertCacheMut.RLock()
	cache, cached := whtUpsertCache[key]
	whtUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			whtAllColumns,
			whtColumnsWithDefault,
			whtColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			whtAllColumns,
			whtPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sqlboilermodel: unable to upsert wht, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(whtPrimaryKeyColumns))
			copy(conflict, whtPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"wht\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(whtType, whtMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(whtType, whtMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "sqlboilermodel: unable to upsert wht")
	}

	if !cached {
		whtUpsertCacheMut.Lock()
		whtUpsertCache[key] = cache
		whtUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single WHT record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *WHT) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlboilermodel: no WHT provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), whtPrimaryKeyMapping)
	sql := "DELETE FROM \"wht\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to delete from wht")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by delete for wht")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q whtQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlboilermodel: no whtQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to delete all from wht")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by deleteall for wht")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WHTSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(whtBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), whtPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"wht\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, whtPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to delete all from wht slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by deleteall for wht")
	}

	if len(whtAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *WHT) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWHT(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WHTSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WHTSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), whtPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"wht\".* FROM \"wht\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, whtPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlboilermodel: unable to reload all in WHTSlice")
	}

	*o = slice

	return nil
}

// WHTExists checks if the WHT row exists.
func WHTExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"wht\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboilermodel: unable to check if wht exists")
	}

	return exists, nil
}
