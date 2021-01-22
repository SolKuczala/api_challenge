package db

type DBClient struct {
}

func NewDBClient(conString string) (*DBClient, error) {
	return &DBClient{}, nil
}

func (db *DBClient) Save() (string, error) {
	return "new connection", nil
}

func (db *DBClient) Close() (string, error) {
	return "new connection", nil
}

func (db *DBClient) Update() (string, error) {
	return "new connection", nil
}
