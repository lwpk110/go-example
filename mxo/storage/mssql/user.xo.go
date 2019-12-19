// Package mssql contains the types for schema 'dbo'.
package mssql

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// User represents a row from 'dbo.user'.
type User struct {
	ID          int       `json:"id"`           // id
	Subject     string    `json:"subject"`      // subject
	CreatedDate time.Time `json:"created_date"` // created_date
	ChangedDate time.Time `json:"changed_date"` // changed_date
	DeletedDate time.Time `json:"deleted_date"` // deleted_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by identity
	const sqlstr = `INSERT INTO dbo.user (` +
		`subject, created_date, changed_date, deleted_date` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`

	// run query
	XOLog(sqlstr, u.Subject, u.CreatedDate, u.ChangedDate, u.DeletedDate)
	res, err := db.Exec(sqlstr, u.Subject, u.CreatedDate, u.ChangedDate, u.DeletedDate)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = int(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE dbo.user SET ` +
		`subject = $1, created_date = $2, changed_date = $3, deleted_date = $4` +
		` WHERE id = $5`

	// run query
	XOLog(sqlstr, u.Subject, u.CreatedDate, u.ChangedDate, u.DeletedDate, u.ID)
	_, err = db.Exec(sqlstr, u.Subject, u.CreatedDate, u.ChangedDate, u.DeletedDate, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM dbo.user WHERE id = $1`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserByID retrieves a row from 'dbo.user' as a User.
//
// Generated from index 'PK__user__3213E83F79A696F8'.
func UserByID(db XODB, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, subject, created_date, changed_date, deleted_date ` +
		`FROM dbo.user ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Subject, &u.CreatedDate, &u.ChangedDate, &u.DeletedDate)
	if err != nil {
		return nil, err
	}

	return &u, nil
}