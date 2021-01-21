package libnova

import "testing"

func TestTableHappy(t *testing.T) {
	table := NewTable("TestHappy", 16)
	table.Set("Name", "Nóva")
	table.Set("Best Mod", "PComm")
	t.Logf("Name: %s", table.Get("Name"))
	if table.Get("Name").String() != "Nóva" {
		t.Errorf("Name != Nóva")
	}
	if table.Get("Best Mod").String() != "PComm" {
		t.Errorf("Name != PComm")
	}
}

func TestTableSad(t *testing.T) {
	table := NewTable("TestSad", 8)
	table.Set("Name", "Nóva")
	table.Set("Popsicles", "Hell yeah")
	table.Set("Favorite Color", "Red Plaid")
	table.Set("Best Mod", "PComm")
	if table.Get("Namez").String() == "Nóva" {
		t.Errorf("Name == Nóva")
	}
}
