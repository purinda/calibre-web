package models

// A thin Model interface.  Implementing this interface will allow a number
// of simplifications to be executed on that model, like applying default
// ordering, atomatically creating indexes, etc.  Though not necessary to
// use the interface, you should add an empty instance of your model to
// db.Models, which will auto-register indexes at connection time.
type Model interface {
	Collection() string
	Indexes() [][]string
}
