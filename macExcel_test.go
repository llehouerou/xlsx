package xlsx

import (
	. "gopkg.in/check.v1"
)

type MacExcelSuite struct{}

var _ = Suite(&MacExcelSuite{})

// Test that we can successfully read an XLSX file generated by
// Microsoft Excel for Mac.  In particular this requires that we
// respect the contents of workbook.xml.rels, which maps the sheet IDs
// to their internal file names.
func (m *MacExcelSuite) TestMacExcel(c *C) {
	xlsxFile, err := OpenFile("macExcelTest.xlsx")
	c.Assert(err, IsNil)
	c.Assert(xlsxFile, NotNil)
	s := xlsxFile.Sheets[0].Cell(0, 0).String()
	c.Assert(s, Equals, "编号")
}
