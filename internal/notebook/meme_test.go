package notebook

import (
	"testing"
)

func TestGetMemeID(t *testing.T) {
	nb := loadTestNotebook(t)

	if id := GetMemeID(nb.Cells[0]); id != "" {
		t.Errorf("cell 0: expected empty meme, got %q", id)
	}
	if id := GetMemeID(nb.Cells[2]); id != "aaaa-bbbb-cccc" {
		t.Errorf("cell 2: expected 'aaaa-bbbb-cccc', got %q", id)
	}
	if id := GetMemeID(nb.Cells[4]); id != "aaaa-bbbb-cccc-1-abcd" {
		t.Errorf("cell 4: expected 'aaaa-bbbb-cccc-1-abcd', got %q", id)
	}
}

func TestGetMemeIDNoMetadata(t *testing.T) {
	c := Cell{CellType: "code", Source: NBSource{"x = 1"}}
	if id := GetMemeID(c); id != "" {
		t.Errorf("expected empty meme for cell without metadata, got %q", id)
	}
}

func TestGetMemePrevious(t *testing.T) {
	nb := loadTestNotebook(t)

	if prev := GetMemePrevious(nb.Cells[0]); prev != "" {
		t.Errorf("cell 0: expected empty previous, got %q", prev)
	}
	if prev := GetMemePrevious(nb.Cells[2]); prev != "prev-1111" {
		t.Errorf("cell 2: expected 'prev-1111', got %q", prev)
	}
}

func TestGetMemeNext(t *testing.T) {
	nb := loadTestNotebook(t)

	if next := GetMemeNext(nb.Cells[0]); next != "" {
		t.Errorf("cell 0: expected empty next, got %q", next)
	}
	if next := GetMemeNext(nb.Cells[2]); next != "next-2222" {
		t.Errorf("cell 2: expected 'next-2222', got %q", next)
	}
}
