package erros

type Errors struct {
	Error string `json:"error,omitempty"`
}

type DBLoadError struct {
	message string
}

func (dbe *DBLoadError) Error() string {
	return dbe.message
}
