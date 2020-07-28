// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboilermodel

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("ContentImages", testContentImages)
	t.Run("ContentTexts", testContentTexts)
	t.Run("Migrations", testMigrations)
	t.Run("WHTS", testWHTS)
}

func TestDelete(t *testing.T) {
	t.Run("ContentImages", testContentImagesDelete)
	t.Run("ContentTexts", testContentTextsDelete)
	t.Run("Migrations", testMigrationsDelete)
	t.Run("WHTS", testWHTSDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("ContentImages", testContentImagesQueryDeleteAll)
	t.Run("ContentTexts", testContentTextsQueryDeleteAll)
	t.Run("Migrations", testMigrationsQueryDeleteAll)
	t.Run("WHTS", testWHTSQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("ContentImages", testContentImagesSliceDeleteAll)
	t.Run("ContentTexts", testContentTextsSliceDeleteAll)
	t.Run("Migrations", testMigrationsSliceDeleteAll)
	t.Run("WHTS", testWHTSSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("ContentImages", testContentImagesExists)
	t.Run("ContentTexts", testContentTextsExists)
	t.Run("Migrations", testMigrationsExists)
	t.Run("WHTS", testWHTSExists)
}

func TestFind(t *testing.T) {
	t.Run("ContentImages", testContentImagesFind)
	t.Run("ContentTexts", testContentTextsFind)
	t.Run("Migrations", testMigrationsFind)
	t.Run("WHTS", testWHTSFind)
}

func TestBind(t *testing.T) {
	t.Run("ContentImages", testContentImagesBind)
	t.Run("ContentTexts", testContentTextsBind)
	t.Run("Migrations", testMigrationsBind)
	t.Run("WHTS", testWHTSBind)
}

func TestOne(t *testing.T) {
	t.Run("ContentImages", testContentImagesOne)
	t.Run("ContentTexts", testContentTextsOne)
	t.Run("Migrations", testMigrationsOne)
	t.Run("WHTS", testWHTSOne)
}

func TestAll(t *testing.T) {
	t.Run("ContentImages", testContentImagesAll)
	t.Run("ContentTexts", testContentTextsAll)
	t.Run("Migrations", testMigrationsAll)
	t.Run("WHTS", testWHTSAll)
}

func TestCount(t *testing.T) {
	t.Run("ContentImages", testContentImagesCount)
	t.Run("ContentTexts", testContentTextsCount)
	t.Run("Migrations", testMigrationsCount)
	t.Run("WHTS", testWHTSCount)
}

func TestHooks(t *testing.T) {
	t.Run("ContentImages", testContentImagesHooks)
	t.Run("ContentTexts", testContentTextsHooks)
	t.Run("Migrations", testMigrationsHooks)
	t.Run("WHTS", testWHTSHooks)
}

func TestInsert(t *testing.T) {
	t.Run("ContentImages", testContentImagesInsert)
	t.Run("ContentImages", testContentImagesInsertWhitelist)
	t.Run("ContentTexts", testContentTextsInsert)
	t.Run("ContentTexts", testContentTextsInsertWhitelist)
	t.Run("Migrations", testMigrationsInsert)
	t.Run("Migrations", testMigrationsInsertWhitelist)
	t.Run("WHTS", testWHTSInsert)
	t.Run("WHTS", testWHTSInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ContentImageToWHTUsingWHT", testContentImageToOneWHTUsingWHT)
	t.Run("ContentTextToWHTUsingWHT", testContentTextToOneWHTUsingWHT)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("WHTToContentImages", testWHTToManyContentImages)
	t.Run("WHTToContentTexts", testWHTToManyContentTexts)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ContentImageToWHTUsingContentImages", testContentImageToOneSetOpWHTUsingWHT)
	t.Run("ContentTextToWHTUsingContentTexts", testContentTextToOneSetOpWHTUsingWHT)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("WHTToContentImages", testWHTToManyAddOpContentImages)
	t.Run("WHTToContentTexts", testWHTToManyAddOpContentTexts)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("ContentImages", testContentImagesReload)
	t.Run("ContentTexts", testContentTextsReload)
	t.Run("Migrations", testMigrationsReload)
	t.Run("WHTS", testWHTSReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("ContentImages", testContentImagesReloadAll)
	t.Run("ContentTexts", testContentTextsReloadAll)
	t.Run("Migrations", testMigrationsReloadAll)
	t.Run("WHTS", testWHTSReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("ContentImages", testContentImagesSelect)
	t.Run("ContentTexts", testContentTextsSelect)
	t.Run("Migrations", testMigrationsSelect)
	t.Run("WHTS", testWHTSSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("ContentImages", testContentImagesUpdate)
	t.Run("ContentTexts", testContentTextsUpdate)
	t.Run("Migrations", testMigrationsUpdate)
	t.Run("WHTS", testWHTSUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("ContentImages", testContentImagesSliceUpdateAll)
	t.Run("ContentTexts", testContentTextsSliceUpdateAll)
	t.Run("Migrations", testMigrationsSliceUpdateAll)
	t.Run("WHTS", testWHTSSliceUpdateAll)
}
