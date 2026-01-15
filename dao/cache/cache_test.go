package cache

import (
	"errors"
	"testing"

	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/fields"
)

type testRecord struct {
	ID    string
	AltID string
	Name  string
}

func resetCacheForTest() {
	Cache.cache = make(map[string]entrys)
	Cache.indices = make(map[string][]fields.Field)
	Cache.key = make(map[string]fields.Field)
	Cache.tablesActive = make(map[string]bool)
}

func TestCache_Get_WhenCacheDoesNotExist(t *testing.T) {
	resetCacheForTest()

	_, err := Get(testRecord{}, "any")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, ce.ErrCacheDoesNotExist) {
		t.Fatalf("expected ErrCacheDoesNotExist, got %v", err)
	}
}

func TestCache_AddKey_WhenNotEnabled(t *testing.T) {
	resetCacheForTest()

	_ = Enable(testRecord{})
	// Not initialised/enabled yet.
	err := AddKey(testRecord{}, fields.Field("ID"))
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, ce.ErrCacheNotEnabled) {
		t.Fatalf("expected ErrCacheNotEnabled, got %v", err)
	}
}

func TestCache_Add_WhenNoKeyDefined(t *testing.T) {
	resetCacheForTest()

	_ = Enable(testRecord{})
	_ = Initialise(testRecord{})

	err := Add(testRecord{ID: "A", Name: "Alpha"})
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, ce.ErrCacheNoKeyDefined) {
		t.Fatalf("expected ErrCacheNoKeyDefined, got %v", err)
	}
}

func TestCache_AddGetRemove_HappyPath(t *testing.T) {
	resetCacheForTest()

	_ = Enable(testRecord{})
	_ = Initialise(testRecord{})
	if err := AddKey(testRecord{}, fields.Field("ID")); err != nil {
		t.Fatalf("AddKey: %v", err)
	}

	rec := testRecord{ID: "A", AltID: "X", Name: "Alpha"}
	if err := Add(rec); err != nil {
		t.Fatalf("Add: %v", err)
	}

	gotAny, err := Get(testRecord{}, "A")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	got, ok := gotAny.(testRecord)
	if !ok {
		t.Fatalf("expected testRecord, got %T", gotAny)
	}
	if got != rec {
		t.Fatalf("expected %+v, got %+v", rec, got)
	}

	if c, err := Count(testRecord{}); err != nil || c != 1 {
		t.Fatalf("Count expected 1, got %d (err=%v)", c, err)
	}

	all, err := GetAll(testRecord{})
	if err != nil {
		t.Fatalf("GetAll: %v", err)
	}
	if len(all) != 1 {
		t.Fatalf("expected 1 record, got %d", len(all))
	}

	if err := Remove(rec); err != nil {
		t.Fatalf("Remove: %v", err)
	}

	_, err = Get(testRecord{}, "A")
	if err == nil {
		t.Fatalf("expected error")
	}
	if !errors.Is(err, ce.ErrCacheRecordNotFound) {
		t.Fatalf("expected ErrCacheRecordNotFound, got %v", err)
	}
}

func TestCache_RemoveByKey(t *testing.T) {
	resetCacheForTest()

	_ = Enable(testRecord{})
	_ = Initialise(testRecord{})
	if err := AddKey(testRecord{}, fields.Field("ID")); err != nil {
		t.Fatalf("AddKey: %v", err)
	}

	_ = Add(testRecord{ID: "A"})
	_ = Add(testRecord{ID: "B"})

	if err := RemoveByKey(testRecord{}, "A"); err != nil {
		t.Fatalf("RemoveByKey: %v", err)
	}
	if c, err := Count(testRecord{}); err != nil || c != 1 {
		t.Fatalf("Count expected 1, got %d (err=%v)", c, err)
	}
}

func TestCache_AddIndex_RemoveIndex(t *testing.T) {
	resetCacheForTest()

	_ = Enable(testRecord{})
	_ = Initialise(testRecord{})

	if err := AddIndex(testRecord{}, fields.Field("AltID")); err != nil {
		t.Fatalf("AddIndex: %v", err)
	}
	// Duplicate should be a no-op.
	if err := AddIndex(testRecord{}, fields.Field("AltID")); err != nil {
		t.Fatalf("AddIndex duplicate: %v", err)
	}

	table := GetStructType(testRecord{})
	if got := len(Cache.indices[table]); got != 1 {
		t.Fatalf("expected 1 index, got %d", got)
	}

	if err := RemoveIndex(testRecord{}, fields.Field("AltID")); err != nil {
		t.Fatalf("RemoveIndex: %v", err)
	}
	if got := len(Cache.indices[table]); got != 0 {
		t.Fatalf("expected 0 indices, got %d", got)
	}

	// Removing non-existent index should not error.
	if err := RemoveIndex(testRecord{}, fields.Field("AltID")); err != nil {
		t.Fatalf("RemoveIndex non-existent: %v", err)
	}
}
