// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboilermodel

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testContentTexts(t *testing.T) {
	t.Parallel()

	query := ContentTexts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testContentTextsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentTextsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ContentTexts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentTextsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContentTextSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentTextsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ContentTextExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ContentText exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContentTextExists to return true, but got false.")
	}
}

func testContentTextsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	contentTextFound, err := FindContentText(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if contentTextFound == nil {
		t.Error("want a record, got nil")
	}
}

func testContentTextsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ContentTexts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testContentTextsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ContentTexts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContentTextsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentTextOne := &ContentText{}
	contentTextTwo := &ContentText{}
	if err = randomize.Struct(seed, contentTextOne, contentTextDBTypes, false, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}
	if err = randomize.Struct(seed, contentTextTwo, contentTextDBTypes, false, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contentTextOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contentTextTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ContentTexts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContentTextsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	contentTextOne := &ContentText{}
	contentTextTwo := &ContentText{}
	if err = randomize.Struct(seed, contentTextOne, contentTextDBTypes, false, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}
	if err = randomize.Struct(seed, contentTextTwo, contentTextDBTypes, false, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contentTextOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contentTextTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func contentTextBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func contentTextAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentText) error {
	*o = ContentText{}
	return nil
}

func testContentTextsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ContentText{}
	o := &ContentText{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, contentTextDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ContentText object: %s", err)
	}

	AddContentTextHook(boil.BeforeInsertHook, contentTextBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	contentTextBeforeInsertHooks = []ContentTextHook{}

	AddContentTextHook(boil.AfterInsertHook, contentTextAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	contentTextAfterInsertHooks = []ContentTextHook{}

	AddContentTextHook(boil.AfterSelectHook, contentTextAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	contentTextAfterSelectHooks = []ContentTextHook{}

	AddContentTextHook(boil.BeforeUpdateHook, contentTextBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	contentTextBeforeUpdateHooks = []ContentTextHook{}

	AddContentTextHook(boil.AfterUpdateHook, contentTextAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	contentTextAfterUpdateHooks = []ContentTextHook{}

	AddContentTextHook(boil.BeforeDeleteHook, contentTextBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	contentTextBeforeDeleteHooks = []ContentTextHook{}

	AddContentTextHook(boil.AfterDeleteHook, contentTextAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	contentTextAfterDeleteHooks = []ContentTextHook{}

	AddContentTextHook(boil.BeforeUpsertHook, contentTextBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	contentTextBeforeUpsertHooks = []ContentTextHook{}

	AddContentTextHook(boil.AfterUpsertHook, contentTextAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	contentTextAfterUpsertHooks = []ContentTextHook{}
}

func testContentTextsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentTextsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(contentTextColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentTextToOneWHTUsingWHT(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ContentText
	var foreign WHT

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contentTextDBTypes, false, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, whtDBTypes, false, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.WHTID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.WHT().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ContentTextSlice{&local}
	if err = local.L.LoadWHT(ctx, tx, false, (*[]*ContentText)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WHT == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.WHT = nil
	if err = local.L.LoadWHT(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WHT == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContentTextToOneSetOpWHTUsingWHT(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ContentText
	var b, c WHT

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentTextDBTypes, false, strmangle.SetComplement(contentTextPrimaryKeyColumns, contentTextColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, whtDBTypes, false, strmangle.SetComplement(whtPrimaryKeyColumns, whtColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, whtDBTypes, false, strmangle.SetComplement(whtPrimaryKeyColumns, whtColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*WHT{&b, &c} {
		err = a.SetWHT(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.WHT != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ContentTexts[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.WHTID != x.ID {
			t.Error("foreign key was wrong value", a.WHTID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.WHTID))
		reflect.Indirect(reflect.ValueOf(&a.WHTID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.WHTID != x.ID {
			t.Error("foreign key was wrong value", a.WHTID, x.ID)
		}
	}
}

func testContentTextsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testContentTextsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContentTextSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testContentTextsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ContentTexts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	contentTextDBTypes = map[string]string{`ID`: `bigint`, `WHTID`: `bigint`, `Name`: `character varying`, `Text`: `text`, `CreatedBy`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedBy`: `character varying`, `UpdatedAt`: `timestamp with time zone`}
	_                  = bytes.MinRead
)

func testContentTextsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(contentTextPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(contentTextAllColumns) == len(contentTextPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testContentTextsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(contentTextAllColumns) == len(contentTextPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ContentText{}
	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contentTextDBTypes, true, contentTextPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(contentTextAllColumns, contentTextPrimaryKeyColumns) {
		fields = contentTextAllColumns
	} else {
		fields = strmangle.SetComplement(
			contentTextAllColumns,
			contentTextPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ContentTextSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testContentTextsUpsert(t *testing.T) {
	t.Parallel()

	if len(contentTextAllColumns) == len(contentTextPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ContentText{}
	if err = randomize.Struct(seed, &o, contentTextDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ContentText: %s", err)
	}

	count, err := ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, contentTextDBTypes, false, contentTextPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentText struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ContentText: %s", err)
	}

	count, err = ContentTexts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
