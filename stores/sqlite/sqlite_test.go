package sqlite

import (
	"os"
	"testing"

	"github.com/juliankoehn/mchurl/stores/shared"
)

var (
	testDbName = "test.db"
	testId     = "123s7a"
)

func TestSqlite(t *testing.T) {
	store, err := New("test.db")
	if err != nil {
		t.Error(err)
	}

	if _, err := store.CreateEntry(shared.Entry{
		URL:        "https://www.google.com",
		RemoteAddr: "0.0.0.0:443",
	}, testId); err != nil {
		t.Error(err)
	}

	// read the entry in db
	entry, err := store.GetEntryByID(testId)
	if err != nil {
		t.Error(err)
	}

	if entry.URL != "https://www.google.com" {
		t.Error("invalid Entry")
	}

	// test 404 entry
	_, err = store.GetEntryByID("a0a0a0")
	if err == nil {
		t.Error(err)
	}

	// increase visit counter
	if err := store.IncreaseVisitCounter(&shared.Entry{
		ID: testId,
	}); err != nil {
		t.Error(err)
	}

	if err := store.DeleteEntry(testId); err != nil {
		t.Error(err)
	}

	// clean up db file from disk after test
	os.Remove(testDbName)
}
