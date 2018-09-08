package infinitescroll

// Checkbox represents the state of the select all checkbox
type Checkbox int

const (
	// Checked is when the select all box is selected.
	Checked Checkbox = iota
	// Unchecked is when the select all box is deselected.
	Unchecked
	// Indeterminate is when the select all box is greyed out.
	Indeterminate
)

type table struct {
	currentState           Checkbox
	sites                  map[string]bool
	totalCount, clickCount uint64
}

// Table represents an infinite scroll list in a webpage.
type Table interface {
	AddSites(sites []string)
	ToggleSelectAll()
	CurrentState() Checkbox
	IsClicked(site string) bool
	Click(site string)
}

// NewTable creates a new instance of an Table
func NewTable(totalCount uint64) Table {
	return &table{sites: make(map[string]bool),
		currentState: Unchecked, totalCount: totalCount}
}

func (ist *table) AddSites(sites []string) {
	click := ist.CurrentState() == Checked
	for _, site := range sites {
		ist.sites[site] = click
	}
	if click {
		ist.clickCount += uint64(len(sites))
	}
}

func (ist *table) ToggleSelectAll() {
	if ist.currentState == Checked {
		ist.currentState = Unchecked
		toggleAll(ist.sites, false)
		ist.clickCount = 0
	} else {
		ist.currentState = Checked
		toggleAll(ist.sites, true)
		ist.clickCount = uint64(len(ist.sites))
	}
}

func toggleAll(sites map[string]bool, click bool) {
	for site := range sites {
		sites[site] = click
	}
}

func (ist *table) CurrentState() Checkbox { return ist.currentState }

func (ist *table) IsClicked(site string) bool {
	if ist.currentState == Checked {
		return true
	} else if ist.currentState == Unchecked {
		return false
	}
	return ist.sites[site]
}

func (ist *table) Click(site string) {
	if state := ist.CurrentState(); state == Unchecked || state == Checked {
		ist.currentState = Indeterminate
	}
	clicked := ist.sites[site]
	ist.sites[site] = !clicked
	if clicked {
		ist.clickCount--
	} else {
		ist.clickCount++
	}
	if ist.clickCount == 0 {
		ist.currentState = Unchecked
	}
	if ist.clickCount == ist.totalCount {
		ist.currentState = Checked
	}
}
