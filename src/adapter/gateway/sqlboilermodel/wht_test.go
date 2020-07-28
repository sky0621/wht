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

func testWHTS(t *testing.T) {
	t.Parallel()

	query := WHTS()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWHTSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
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

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWHTSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := WHTS().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWHTSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WHTSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWHTSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WHTExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if WHT exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WHTExists to return true, but got false.")
	}
}

func testWHTSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	whtFound, err := FindWHT(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if whtFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWHTSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = WHTS().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWHTSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := WHTS().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWHTSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	whtOne := &WHT{}
	whtTwo := &WHT{}
	if err = randomize.Struct(seed, whtOne, whtDBTypes, false, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}
	if err = randomize.Struct(seed, whtTwo, whtDBTypes, false, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = whtOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = whtTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WHTS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWHTSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	whtOne := &WHT{}
	whtTwo := &WHT{}
	if err = randomize.Struct(seed, whtOne, whtDBTypes, false, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}
	if err = randomize.Struct(seed, whtTwo, whtDBTypes, false, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = whtOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = whtTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func whtBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func whtAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WHT) error {
	*o = WHT{}
	return nil
}

func testWHTSHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &WHT{}
	o := &WHT{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, whtDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WHT object: %s", err)
	}

	AddWHTHook(boil.BeforeInsertHook, whtBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	whtBeforeInsertHooks = []WHTHook{}

	AddWHTHook(boil.AfterInsertHook, whtAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	whtAfterInsertHooks = []WHTHook{}

	AddWHTHook(boil.AfterSelectHook, whtAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	whtAfterSelectHooks = []WHTHook{}

	AddWHTHook(boil.BeforeUpdateHook, whtBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	whtBeforeUpdateHooks = []WHTHook{}

	AddWHTHook(boil.AfterUpdateHook, whtAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	whtAfterUpdateHooks = []WHTHook{}

	AddWHTHook(boil.BeforeDeleteHook, whtBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	whtBeforeDeleteHooks = []WHTHook{}

	AddWHTHook(boil.AfterDeleteHook, whtAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	whtAfterDeleteHooks = []WHTHook{}

	AddWHTHook(boil.BeforeUpsertHook, whtBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	whtBeforeUpsertHooks = []WHTHook{}

	AddWHTHook(boil.AfterUpsertHook, whtAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	whtAfterUpsertHooks = []WHTHook{}
}

func testWHTSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWHTSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(whtColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWHTToManyContentImages(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WHT
	var b, c ContentImage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, contentImageDBTypes, false, contentImageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contentImageDBTypes, false, contentImageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.WHTID = a.ID
	c.WHTID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ContentImages().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.WHTID == b.WHTID {
			bFound = true
		}
		if v.WHTID == c.WHTID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WHTSlice{&a}
	if err = a.L.LoadContentImages(ctx, tx, false, (*[]*WHT)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentImages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ContentImages = nil
	if err = a.L.LoadContentImages(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentImages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWHTToManyContentTexts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WHT
	var b, c ContentText

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, contentTextDBTypes, false, contentTextColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, contentTextDBTypes, false, contentTextColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.WHTID = a.ID
	c.WHTID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ContentTexts().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.WHTID == b.WHTID {
			bFound = true
		}
		if v.WHTID == c.WHTID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WHTSlice{&a}
	if err = a.L.LoadContentTexts(ctx, tx, false, (*[]*WHT)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentTexts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ContentTexts = nil
	if err = a.L.LoadContentTexts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentTexts); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWHTToManyAddOpContentImages(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WHT
	var b, c, d, e ContentImage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, whtDBTypes, false, strmangle.SetComplement(whtPrimaryKeyColumns, whtColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ContentImage{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, contentImageDBTypes, false, strmangle.SetComplement(contentImagePrimaryKeyColumns, contentImageColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ContentImage{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddContentImages(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.WHTID {
			t.Error("foreign key was wrong value", a.ID, first.WHTID)
		}
		if a.ID != second.WHTID {
			t.Error("foreign key was wrong value", a.ID, second.WHTID)
		}

		if first.R.WHT != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.WHT != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ContentImages[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ContentImages[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ContentImages().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testWHTToManyAddOpContentTexts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WHT
	var b, c, d, e ContentText

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, whtDBTypes, false, strmangle.SetComplement(whtPrimaryKeyColumns, whtColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ContentText{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, contentTextDBTypes, false, strmangle.SetComplement(contentTextPrimaryKeyColumns, contentTextColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ContentText{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddContentTexts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.WHTID {
			t.Error("foreign key was wrong value", a.ID, first.WHTID)
		}
		if a.ID != second.WHTID {
			t.Error("foreign key was wrong value", a.ID, second.WHTID)
		}

		if first.R.WHT != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.WHT != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ContentTexts[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ContentTexts[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ContentTexts().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testWHTSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
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

func testWHTSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WHTSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWHTSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WHTS().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	whtDBTypes = map[string]string{`ID`: `bigint`, `RecordDate`: `timestamp without time zone`, `Title`: `character varying`, `CreatedBy`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedBy`: `character varying`, `UpdatedAt`: `timestamp with time zone`}
	_          = bytes.MinRead
)

func testWHTSUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(whtPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(whtAllColumns) == len(whtPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, whtDBTypes, true, whtPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWHTSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(whtAllColumns) == len(whtPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WHT{}
	if err = randomize.Struct(seed, o, whtDBTypes, true, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, whtDBTypes, true, whtPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(whtAllColumns, whtPrimaryKeyColumns) {
		fields = whtAllColumns
	} else {
		fields = strmangle.SetComplement(
			whtAllColumns,
			whtPrimaryKeyColumns,
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

	slice := WHTSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testWHTSUpsert(t *testing.T) {
	t.Parallel()

	if len(whtAllColumns) == len(whtPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := WHT{}
	if err = randomize.Struct(seed, &o, whtDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert WHT: %s", err)
	}

	count, err := WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, whtDBTypes, false, whtPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert WHT: %s", err)
	}

	count, err = WHTS().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}