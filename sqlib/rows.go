package sqlib

type TableRows struct {
	rowslist map[string]string
}

func (r *TableRows) AddRow(name string, datatype string) {
	r.rowslist[name] = datatype
}

func NewRowsList() TableRows {
	return TableRows{
		rowslist: make(map[string]string),
	}
}

func (r *TableRows) ToString() string {
	var line string

	for k, v := range r.rowslist {
		line += k + " " + v + ", "
	}

	if len(line) > 0 {
		line = line[:len(line)-2]
	}

	return line
}
