package models

type ModuleAttachment struct {
	ID        string `db:"id"`
	ModuleID  string `db:"id_module"`
	FileName  string `db:"c_file_name"`
	Path      string `db:"c_path"`
	CreatedAt string `db:"c_created_at"`
	Visible   bool   `db:"c_visible"`
}
