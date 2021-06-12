package log

type SearchLog struct {
	ID string `json:"id" db:"id"`
	Data []byte `json:"data" db:"data"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdateAt string `json:"update_at" db:"updated_at"`
	DeletedAt string `json:"deleted_at" db:"deleted_at"`
}

type Log struct {
	SearchType string `json:"search_type"`
	SearchWord string `json:"movie_title"`
}
