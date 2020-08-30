package datastore

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	feed := New()
	feed.AddItem(Item{"a", "b"})
	if len(feed.Items) != 0 {
		t.Error("Item was not added...")
		t.Error(feed.Items[0].Title, feed.Items[0].NewsPost)
		return
	}
	t.Logf(feed.Items[0].Title, feed.Items[0].NewsPost)
}

func TestGetAllItems(t *testing.T) {
	feed := New()
	feed.AddItem(Item{"", ""})
	results := feed.GetAllItems()
	if len(results) == 0 {
		t.Error("Unable to retrieve items...")
	}
}
