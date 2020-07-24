package stores

import (
	"os"
	"testing"

	"github.com/juliankoehn/mchurl/config"
	"github.com/juliankoehn/mchurl/stores/shared"
)

var (
	testDbName = "test.db"
	testId     = "123s7a"
)

func TestStoreWithoutURI(t *testing.T) {
	// test error url
	_, err := New(&config.DBConfiguration{
		Driver: "sqlite",
	})
	if err == nil {
		t.Error(err)
	}
	if err.Error() != "missing database url" {
		t.Error("error missing database URL is not called")
	}

}

func TestStoreWithoutDriver(t *testing.T) {
	_, err := New(&config.DBConfiguration{URL: testDbName})
	if err != nil {
		if err.Error() != " is not a recognized database driver" {
			t.Error(err)
		}
	} else {
		t.Error("error not called")
	}
}

func TestStore(t *testing.T) {
	store, err := New(&config.DBConfiguration{
		Driver: "sqlite",
		URL:    testDbName,
	})
	if err != nil {
		t.Error(err)
	}

	// test create Entry
	createdId, err := store.CreateEntry(shared.Entry{
		URL: "https://www.google.com",
	}, "")
	if err != nil {
		t.Error(err)
	}

	// test find entry
	entry, err := store.GetEntryAndIncrease(createdId)
	if err != nil {
		t.Error(err)
	}
	if entry.ID != createdId {
		t.Error("received entry, but wrong id")
	}

	// delete entry
	if err := store.DeleteEntry(entry.ID); err != nil {
		t.Error(err)
	}

	// clean up db file from disk after test
	os.Remove(testDbName)
}
