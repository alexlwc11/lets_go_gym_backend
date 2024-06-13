package configs

import (
	"bytes"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Credentials struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"credentials"`
		Net       string `yaml:"net"`
		Addr      string `yaml:"addr"`
		DBName    string `yaml:"db_name"`
		Charset   string `yaml:"charset"`
		Loc       string `yaml:"loc"`
		ParseTime bool   `yaml:"parse_time"`
	} `yaml:"database"`
}

func LoadConfig() Config {
	f, err := os.Open("config_files/api/private_config.yaml")
	if err != nil {
		log.Fatal("Failed to load database config")
	}

	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Failed to decode database config")
	}

	return cfg
}

// referencing the Config.FormatDSN() method in [github.com/go-sql-driver/mysql]
// just added the params we are using, may need to add some more later
// more details: [https://github.com/go-sql-driver/mysql#dsn-data-source-name]
func (cfg *Config) GetDSNString() string {
	dbCfg := cfg.Database

	var buf bytes.Buffer

	// [username[:password]@]
	if dbCfg.Credentials.Username != "" {
		buf.WriteString(dbCfg.Credentials.Username)
		if dbCfg.Credentials.Password != "" {
			buf.WriteByte(':')
			buf.WriteString(dbCfg.Credentials.Password)
		}
		buf.WriteByte('@')
	}

	// [protocol[(address)]]
	if dbCfg.Net != "" {
		buf.WriteString(dbCfg.Net)
		if dbCfg.Addr != "" {
			buf.WriteByte('(')
			buf.WriteString(dbCfg.Addr)
			buf.WriteByte(')')
		}
	}

	// db name
	buf.WriteByte('/')
	buf.WriteString(dbCfg.DBName)

	hasParam := false

	if dbCfg.Charset != "" {
		writeDSNParam(&buf, &hasParam, "charset", dbCfg.Charset)
	}

	if dbCfg.Loc != "" {
		writeDSNParam(&buf, &hasParam, "loc", dbCfg.Loc)
	}

	if dbCfg.ParseTime {
		writeDSNParam(&buf, &hasParam, "parseTime", "true")
	}

	return buf.String()
}

func writeDSNParam(buf *bytes.Buffer, hasParam *bool, key string, value string) {
	if !*hasParam {
		*hasParam = true
		buf.WriteByte('?')
	} else {
		buf.WriteByte('&')
	}
	buf.WriteString(key)
	buf.WriteByte('=')
	buf.WriteString(value)
}
