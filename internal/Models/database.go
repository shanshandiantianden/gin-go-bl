package Models

type DatabaseConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Database string `ini:"database"`
	Charset  string `ini:"charset"`
	Loc      string `ini:"loc"`
}
type Config struct {
	DatabaseConfig `ini:"database"`
}
