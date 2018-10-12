package excelize

import "math"

const (
	EMUPerInch = 914400
)

// UnitsToPixelsMap should describe a function that accepts excel units and outputs the mapped pixels as integer
type UnitsToPixelsMap = func(units float64) int

// Display describes the display configuration for the excel file. It can derive
type Display struct {
	ColMap           UnitsToPixelsMap
	ColUnitsDefault  float64
	colPixelsDefault int
	RowMap           UnitsToPixelsMap
	RowUnitsDefault  float64
	rowPixelsDefault int
	DPI              int
	emuPerPixel      int
}

// defaultDisplay returns the default display configuration for a windows-based excel viewer
func defaultDisplay() Display {
	result := Display{
		ColMap: func(units float64) int {
			if units <= 0 {
				return 0
			}

			if units <= 1 {
				return int(math.Ceil(units * 13))
			}

			return int(math.Ceil(8*units+5)) // TODO: Check these on a windows machine
		},
		ColUnitsDefault: 10,
		RowMap: func(units float64) int {
			if units <= 0 {
				return 0
			}

			return int(math.Ceil(units*4/3))
		},
		RowUnitsDefault: 15,
		DPI:             96,
		emuPerPixel:     EMUPerInch / 96,
	}

	result.colPixelsDefault = result.ColMap(result.ColUnitsDefault)
	result.rowPixelsDefault = result.RowMap(result.RowUnitsDefault)
	return result
}

// SetDisplay modifies the internal mapping from units to pixels and from pixels to EMUs. This function should be called
// to properly configure the display before any width/height modifying functions are called.
func (f *File) SetDisplay(display Display) {
	display.colPixelsDefault = display.ColMap(display.ColUnitsDefault)
	display.rowPixelsDefault = display.RowMap(display.RowUnitsDefault)
	display.emuPerPixel = EMUPerInch / display.DPI
	f.display = display
}
