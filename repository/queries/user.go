package queries

import _ "embed"

var (
	//go:embed user/select.sql
	ListUsers string
	//go:embed user/insert.sql
	InsertUsers string
	//go:embed user/update.sql
	UpdateUsers string
	//go:embed user/delete.sql
	DeleteUsers string
	//go:embed user/select_email.sql
	SelectEmail string
	//go:embed user/select_id.sql
	SelectID string
)
