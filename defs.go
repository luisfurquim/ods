package ods



type TimeCell struct {
   Type       string `xml:"value-type,attr"`
   Val        string `xml:"p"`
   DisplayVal string `xml:"time-value,attr"`
}

type DateCell struct {
   Type       string `xml:"value-type,attr"`
   Val        string `xml:"p"`
   DisplayVal string `xml:"date-value,attr"`
}

type Cell struct {
   Type       string `xml:"value-type,attr"`
   DisplayVal string `xml:"p"`
   Val        string `xml:"value,attr"`
   DateVal    string `xml:"date-value,attr"`
   TimeVal    string `xml:"time-value,attr"`
}

type Row struct {
   Cells []Cell `xml:"table-cell"`
}

type Document struct {
   Rows  []Row `xml:"body>spreadsheet>table>table-row"`
}

type Document2 struct {
   Row   []Cell `xml:"table-row>table-cell"`
}

