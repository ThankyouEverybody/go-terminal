package config

import (
	"flag"
	"fmt"
	"github.com/go-terminal-server/common/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

var GlobalConfig *config

type config struct {
	Server *Server
	Sqlite *Sqlite
}

func (c config) String() string {
	return fmt.Sprintf("Server: %v\nSqlite: %v\n", c.Server, c.Sqlite)
}

type Server struct {
	Port int
	DB   string
	Log  *Logger
}

type Sqlite struct {
	File string
}

type Logger struct {
	Level string
	Path  string
}

func (l Logger) String() (s string) {
	s = fmt.Sprintf("Level: %v; Path: %v", l.Level, l.Path)
	return
}

func (s Server) String() string {
	return fmt.Sprintf("Port: %v; DB: %v;\nLogger: %v", s.Port, s.DB, s.Log)
}
func (s Sqlite) String() string {
	return fmt.Sprintf("file: %v;", s.File)
}

func LoadConfig() error {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "configuration file")
	flag.Parse()
	absFilePath, _ := filepath.Abs(configPath)
	//判断absFilePath文件是否存在
	if !strings.HasSuffix(absFilePath, ".yml") {
		return fmt.Errorf("%v: invalid config file %v", "LoadConfig", absFilePath)
	}

	fmt.Println("read config file", absFilePath)
	viper.SetConfigFile(absFilePath)

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var logPath string
	var err error
	pflag.String("server.DB", "sqlite", "DB mode")
	pflag.Int("server.Port", 123, "DB mode")
	pflag.String("server.logger.level", "debug", "logger level")
	pflag.String("server.logger.path", "./data/log", "logger level")

	pflag.String("sqlite.file", "./data/sqlite/go-terminal.db", "DB file")

	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err.(any))
	}

	_ = viper.ReadInConfig()

	db := viper.GetString("server.DB")
	port := viper.GetInt("server.Port")
	logLevel := viper.GetString("server.logger.level")

	if logPath, err = filepath.Abs(viper.GetString("server.logger.path")); err != nil {
		panic(err.(any))
	}
	var sqliteFile string
	if sqliteFile, err = filepath.Abs(viper.GetString("sqlite.file")); err != nil {
		panic(err.(any))
	}
	if db == "sqlite" {

		if !strings.HasSuffix(sqliteFile, ".db") {
			panic("DB file must end with.db")
		}
		sqliteDir := filepath.Dir(sqliteFile)
		sqliteDir, err := homedir.Expand(sqliteDir)
		if err != nil {
			return err
		}
		if err := utils.MkdirP(sqliteDir); err != nil {
			panic(fmt.Sprintf("创建文件夹 %v 失败: %v", sqliteDir, err.Error()))
		}
	}

	GlobalConfig = &config{
		Server: &Server{
			DB:   db,
			Port: port,
			Log: &Logger{
				Level: logLevel,
				Path:  logPath,
			},
		},
		Sqlite: &Sqlite{
			File: sqliteFile,
		},
	}
	return nil
}

func init() {
	err := LoadConfig()
	if err != nil {
		panic(err.(any))
	}
}
