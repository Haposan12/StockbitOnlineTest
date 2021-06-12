package movie

import "errors"

//param
const (
	SearchWordParam string = "search"
	PaginationParam string = "pagination"
	MovieTitleParam string = "title"
)

//search type
const (
	PaginationSearch string = "pagination"
	DetailSearch string = "detail"
)

//error
var (
	CheckParamError error = errors.New("check your parameter")
)
