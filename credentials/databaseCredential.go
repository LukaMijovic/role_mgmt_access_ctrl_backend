package creds

type DBCredential struct {
	DBType   string `binding:required`
	User     string `binding:required`
	Password string `binding:required`
	DBName   string `binding:required`
	SSLMode  string
}

func NewDBCredential(dbType string, user string, password string, dbName string, sslMode string) *DBCredential {
	return &DBCredential{
		DBType:   dbType,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode: func() string {
			if sslMode == "" {
				return "disable"
			}
			return sslMode
		}(),
	}
}
