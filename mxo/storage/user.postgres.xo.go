// Package postgres contains the types for schema 'public'.
package storage

import (
	"errors"
)

func (s *PostgresStorage) UserByID(db XODB, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, subject, created_date, changed_date, deleted_date ` +
		`FROM public.user ` +
		`WHERE id = $1`

	// run query
	xoLog(sqlstr, id)
	u := User{}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Subject, &u.CreatedDate, &u.ChangedDate, &u.DeletedDate)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Insert inserts the User to the database.
func (s *PostgresStorage) InsertUser(db XODB, u *User) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.user (` +
		`subject, created_date, changed_date, deleted_date` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	xoLog(sqlstr, u.Subject, u.CreatedDate, u.ChangedDate, u.DeletedDate)
	err = db.QueryRow(sqlstr, u.Subject, u.CreatedDate, u.ChangedDate, u.DeletedDate).Scan(&u.ID)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}