package postgres

type ModelError string

func (e ModelError) Error() string {
	return string(e)
}
