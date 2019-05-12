// Package mocksqldb is a GoMock implementation of the sqldb interface
package mocksqldb

// Only the Queryable is generated. Connection and Tx both extend the generated
// Queryable struct in order to be able to write helpers that works with the
// 3 classes
//go:generate mockgen -destination queryable.go -package mocksqldb github.com/Nivl/go-sqldb Queryable
