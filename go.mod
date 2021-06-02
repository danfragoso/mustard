module github.com/danfragoso/mustard

go 1.16

replace github.com/danfragoso/mustard/backend => ./backend

replace github.com/danfragoso/mustard/device => ./device

require (
	github.com/danfragoso/mustard/backend v0.0.0-00010101000000-000000000000
	github.com/danfragoso/mustard/device v0.0.0-00010101000000-000000000000
)
