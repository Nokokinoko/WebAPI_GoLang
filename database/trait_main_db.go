package database

import (
	"errors"
	"reflect"
	"strings"
)

// ------ Update ------
func (table abstractTable) Update(record abstractRecord) error {
	usePrimaryKey := 0 < len(table.primaryKey)
	specifyKeys := table.specifyKeys(usePrimaryKey)
	if len(specifyKeys) <= 0 {
		return errors.New("Update empty specify keys. table:" + table.tableName)
	}

	// check specify keys
	reflectRecord := reflect.ValueOf(record)
	for _, specifyKey := range specifyKeys {
		val := reflectRecord.FieldByName(specifyKey)
		if !val.IsValid() {
			return errors.New("Update not found specify key. table:" + table.tableName + ", key:" + specifyKey)
		}
	}

	// create sql
	var sqlWhere []string
	var sqlSets []string
	var sqlPhs []string
	for _, col := range record.columns() {
		// check in array
		var isInclude bool = false
		for _, specifyKey := range specifyKeys {
			if col == specifyKey {
				isInclude = true
				break
			}
		}

		if isInclude {
			sqlWhere = append(sqlWhere, col+" = ?")
		} else {
			sqlSets = append(sqlSets, col+" = ?")
		}
		val := reflectRecord.FieldByName(col)
		sqlPhs = append(sqlPhs, val.String())
	}

	instance := Instance(table.databaseName, true)
	// begin transaction
	tx, _ := instance.db.Begin()
	update, err := instance.db.Prepare("update " + table.tableName + " set " + strings.Join(sqlSets, ", ") + " where " + strings.Join(sqlWhere, " and ") + " limit 1")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer update.Close()

	result, err := update.Exec(sqlPhs)
	if err != nil {
		tx.Rollback()
		return err
	}

	affect, err := result.RowsAffected()
	if affect != 1 || err != nil {
		tx.Rollback()
		return errors.New("Update not found record")
	}

	tx.Commit()
	return nil
}

// ------ Insert ------
func (table abstractTable) Insert(record abstractRecord) (abstractRecord, error) {
	usePrimaryKey := 0 < len(table.primaryKey)
	reflectRecord := reflect.ValueOf(record)
	if usePrimaryKey {
		val := reflectRecord.FieldByName(table.primaryKey)
		if val.IsValid() {
			return record, errors.New("Insert already exists. table:" + table.tableName + ", primary key:" + val.String())
		}
	}

	// create sql
	var sqlKeys []string
	var sqlPhs []string
	for _, col := range record.columns() {
		if col == table.primaryKey {
			continue
		}

		sqlKeys = append(sqlKeys, col)
		val := reflectRecord.FieldByName(col)
		sqlPhs = append(sqlPhs, val.String())
	}

	instance := Instance(table.databaseName, true)
	// begin transaction
	tx, _ := instance.db.Begin()
	insert, err := instance.db.Prepare("insert into " + table.tableName + " ( " + strings.Join(sqlKeys, ", ") + " ) values ( " + strings.Repeat("?, ", len(sqlKeys)-1) + "? )")
	if err != nil {
		tx.Rollback()
		return record, err
	}
	defer insert.Close()

	result, err := insert.Exec(sqlPhs)
	if err != nil {
		tx.Rollback()
		return record, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return record, err
	}

	tx.Commit()
	if usePrimaryKey {
		reflectRecord.FieldByName(table.primaryKey).SetInt(lastInsertId)
	}
	return record, nil
}

// ------ Delete ------
func (table abstractTable) Delete(mapKeyVal map[string]string) error {
	usePrimaryKey := 0 < len(table.primaryKey)
	specifyKeys := table.specifyKeys(usePrimaryKey)
	if len(specifyKeys) <= 0 {
		return errors.New("Delete empty specify keys. table:" + table.tableName)
	}

	// create sql
	var sqlWhere []string
	var sqlPhs []string
	for _, specifyKey := range specifyKeys {
		// check in map
		val, isFound := mapKeyVal[specifyKey]
		if !isFound {
			return errors.New("Delete not found specify key. table:" + table.tableName + ", key:" + specifyKey)
		}

		sqlWhere = append(sqlWhere, specifyKey+" = ?")
		sqlPhs = append(sqlPhs, val)
	}

	instance := Instance(table.databaseName, true)
	// begin transaction
	tx, _ := instance.db.Begin()
	delete, err := instance.db.Prepare("delete from " + table.tableName + " where " + strings.Join(sqlWhere, " and ") + " limit 1")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer delete.Close()

	result, err := delete.Exec(sqlPhs)
	if err != nil {
		tx.Rollback()
		return err
	}

	affect, err := result.RowsAffected()
	if affect != 1 || err != nil {
		tx.Rollback()
		return errors.New("Delete not found record")
	}

	tx.Commit()
	return nil
}
