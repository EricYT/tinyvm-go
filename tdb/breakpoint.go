package tdb

const (
	TDB_BREAKPOINT_LABLE int = 0x1
	TDB_BREAKPOINT_ADDRESS
)

type BreakPoint struct {
	address int
	label   string
	typ     int
}
