package configs

type Configs struct {
	PostgreSQL PostgreSQL
	App        Fiber
	MySql
}

type Fiber struct {
	Host string
	Port string
}
type PostgreSQL struct {
	Host      string
	Port      string
	Protocol  string
	Username  string
	Password  string
	Database  string
	SSLMode   string
	ParseTime bool
}

type MySql struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Database  string
	ParseTime bool // Relevant for MySQL

}
