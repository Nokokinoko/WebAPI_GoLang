package database

// ------ Select ------
func (table abstractTable) Rows() []abstractRecord {
	var records []abstractRecord
	return records
}

func (table abstractTable) Row() abstractRecord {
	var record abstractRecord
	return record
}

func (table abstractTable) Count() int {
	return 0
}
