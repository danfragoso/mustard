module github.com/danfragoso/mustard

go 1.16

replace github.com/danfragoso/mustard/backend => ./backend

replace github.com/danfragoso/mustard/device => ./device

require (
	github.com/danfragoso/mustard/backend v0.0.0-20210602153524-cd31edb6ea95
	github.com/danfragoso/mustard/device v0.0.0-20210602153524-cd31edb6ea95
)
