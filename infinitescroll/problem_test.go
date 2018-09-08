package infinitescroll

import "testing"

func TestClickThenUnclick(t *testing.T) {
	t.Parallel()
	table := NewTable(3000)
	table.AddSites([]string{"google", "yahoo", "microsoft", "amazon"})
	table.Click("amazon")
	if result := table.CurrentState(); result != Indeterminate {
		t.Errorf("Expected Indeterminate but got %v", result)
	}
	table.Click("amazon")
	if result := table.CurrentState(); result != Unchecked {
		t.Errorf("Expected Unchecked but got %v", result)
	}
}

func TestIsClickForSeenAndUnseen(t *testing.T) {
	t.Parallel()
	table := NewTable(5)
	table.AddSites([]string{"google", "yahoo", "microsoft", "amazon"})
	if clicked := table.IsClicked("unseen"); clicked {
		t.Errorf("Should be unclicked")
	}
	if selectAll := table.CurrentState(); selectAll != Unchecked {
		t.Errorf("Table should be unchecked for all here")
	}
	table.Click("amazon")
	if selectAll := table.CurrentState(); selectAll != Indeterminate {
		t.Errorf("Table should be indeterminate because 1 site is clicked")
	}
	if clicked := table.IsClicked("unseen"); clicked {
		t.Errorf("Should be unclicked")
	}
}

func TestToggleAll(t *testing.T) {
	t.Parallel()
	table := NewTable(3)
	table.AddSites([]string{"google", "yahoo"})
	table.ToggleSelectAll()
	if clicked := table.IsClicked("google"); !clicked {
		t.Errorf("Google should be clicked because of the toggle select all")
	}
	if clicked := table.IsClicked("yahoo"); !clicked {
		t.Errorf("Yahoo should be clicked because of the toggle select all")
	}
	if clicked := table.IsClicked("amazon"); !clicked {
		t.Errorf("Unseen Amazon should be clicked because of the toggle select all")
	}
	table.AddSites([]string{"amazon"})
	table.ToggleSelectAll()
	if clicked := table.IsClicked("google"); clicked {
		t.Errorf("Google should be unclicked because of the toggle select all")
	}
	if clicked := table.IsClicked("yahoo"); clicked {
		t.Errorf("Yahoo should be unclicked because of the toggle select all")
	}
	if clicked := table.IsClicked("amazon"); clicked {
		t.Errorf("Amazon should be unclicked because of the toggle select all")
	}
}

func TestToggleAllWithClicks(t *testing.T) {
	t.Parallel()
	table := NewTable(3)
	table.ToggleSelectAll()
	table.AddSites([]string{"google", "yahoo"})
	if state := table.CurrentState(); state != Checked {
		t.Errorf("Table should be checked for select all")
	}
	table.Click("google")
	if state := table.CurrentState(); state != Indeterminate {
		t.Errorf("Table should be indeterminate after a deselect click")
	}
	table.ToggleSelectAll()
	if state := table.CurrentState(); state != Checked {
		t.Errorf("Table should be select all")
	}
}

func TestManuallySelectTotalSites(t *testing.T) {
	t.Parallel()
	table := NewTable(3)
	table.AddSites([]string{"google", "yahoo", "amazon"})
	if state := table.CurrentState(); state != Unchecked {
		t.Errorf("Table should be unchecked")
	}
	table.Click("google")
	table.Click("yahoo")
	if state := table.CurrentState(); state != Indeterminate {
		t.Errorf("Table should be indeterminate after two clicks")
	}
	table.Click("amazon")
	if state := table.CurrentState(); state != Checked {
		t.Errorf("Everything manually checked up to the total size")
	}
}
