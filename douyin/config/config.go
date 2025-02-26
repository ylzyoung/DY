package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Mysql struct {
	Host      string
	Port      int
	Database  string
	Username  string
	Password  string
	Charset   string
	Loc       string
	ParseTime bool
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

type Server struct {
	IP   string
	Port int
}

type Path struct {
	FfmpegPath       string `toml:"ffmpeg_path"`
	StaticSourcePath string `toml:"static_source_path"`
}

type Config struct {
	DB     Mysql `toml:"mysql"`
	RDB    Redis `toml:"redis"`
	Server `toml:"server"`
	Path   `toml:"path"`
}

var Global Config

func ensurePathValid() {
	var err error
	if _, err = os.Stat(Global.Path.FfmpegPath); os.IsNotExist(err) {
		if err = os.MkdirAll(Global.Path.FfmpegPath, 0755); err != nil {
			log.Fatalf("mkdir error:path %s", Global.StaticSourcePath)
		}
	}
	if _, err = os.Stat(Global.FfmpegPath); os.IsNotExist(err) {
		if _, err = exec.Command("ffmpeg", "-version").Output(); err != nil {
			log.Fatalf("ffmpeg not valid %s", Global.FfmpegPath)
		} else {
			Global.FfmpegPath = "ffmpeg"
		}
	} else {
		Global.FfmpegPath, err = filepath.Abs(Global.FfmpegPath)
		if err != nil {
			log.Fatalln("get abs path failed:", Global.FfmpegPath)
		}
	}
	//把资源路径转化为绝对路径，防止调用ffmpeg命令失效
	Global.StaticSourcePath, err = filepath.Abs(Global.StaticSourcePath)
	if err != nil {
		log.Fatalln("get abs path failed:", Global.StaticSourcePath)
	}
}

func init() {
	if _, err := toml.DecodeFile("E:\\GOproject\\douyin\\config/config.toml", &Global); err != nil {
		panic(err)
	}
	//去除左右的空格
	Global.Server.IP = strings.Trim(Global.Server.IP, " ")
	Global.RDB.IP = strings.Trim(Global.RDB.IP, " ")
	Global.DB.Host = strings.Trim(Global.DB.Host, " ")
	//保证路径正常
	ensurePathValid()
}

func DBConnectString() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Global.DB.Username, Global.DB.Password, Global.DB.Host, Global.DB.Port, Global.DB.Database,
		Global.DB.Charset, Global.DB.ParseTime, Global.DB.Loc)
	log.Println(arg)
	log.Println(Global.DB.ParseTime)
	return arg
}
