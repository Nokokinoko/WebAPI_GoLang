package database

import (
	"errors"
	"strings"
)

// ------ Select ------
func (table abstractTable) Rows(mapKeyVal map[string]string) ([]abstractRecord, error) {
	// create sql
	sql := "select * from " + table.tableName
	var sqlPhs []string
	if 0 < len(mapKeyVal) {
		var sqlWhere []string
		for key, val := range mapKeyVal {
			sqlWhere = append(sqlWhere, key+" = ?")
			sqlPhs = append(sqlPhs, val)
		}
		sql += " where " + strings.Join(sqlWhere, " and ")
	}

	var records []abstractRecord
	instance := Instance(table.tableName, false)
	stmt, err := instance.db.Prepare(sql)
	if err != nil {
		return records, err
	}

	rows, err := stmt.Query(sqlPhs)
	if err != nil {
		return records, err
	}
	defer rows.Close()

	// assign rows
	for rows.Next() {
		var record abstractRecord
		if err := record.assignRows(rows); err != nil {
			return records, err
		}
		records = append(records, record)
	}
	return records, nil
}

func (table abstractTable) Row(usePrimaryKey bool, mapKeyVal map[string]string) (abstractRecord, error) {
	specifyKeys := table.specifyKeys(usePrimaryKey)
	if len(specifyKeys) <= 0 {
		return abstractRecord{}, errors.New("Row empty specify keys. table:" + table.tableName)
	}

	// create sql
	var sqlWhere []string
	var sqlPhs []string
	for _, specifyKey := range specifyKeys {
		// check in map
		val, isFound := mapKeyVal[specifyKey]
		if !isFound {
			return abstractRecord{}, errors.New("Row not found specify key. table:" + table.tableName + ", key:" + specifyKey)
		}

		sqlWhere = append(sqlWhere, specifyKey+" = ?")
		sqlPhs = append(sqlPhs, val)
	}

	instance := Instance(table.tableName, false)
	stmt, err := instance.db.Prepare("select * from " + table.tableName + " where " + strings.Join(sqlWhere, " and ") + " limit 1")
	if err != nil {
		return abstractRecord{}, err
	}

	// assign row
	var record abstractRecord
	row := stmt.QueryRow(sqlPhs)
	if err := record.assignRow(row); err != nil {
		return record, err
	}
	return record, nil
}

func (table abstractTable) Count(mapKeyVal map[string]string) (int, error) {
	// create sql
	sql := "select count(1) as count from " + table.tableName
	var sqlPhs []string
	if 0 < len(mapKeyVal) {
		var sqlWhere []string
		for key, val := range mapKeyVal {
			sqlWhere = append(sqlWhere, key+" = ?")
			sqlPhs = append(sqlPhs, val)
		}
		sql += " where " + strings.Join(sqlWhere, " and ")
	}

	instance := Instance(table.tableName, false)
	stmt, err := instance.db.Prepare(sql)
	if err != nil {
		return 0, err
	}

	// assign row
	var count int
	if err := stmt.QueryRow(sqlPhs).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
