package config

import (
	"github.com/unknwon/goconfig"
	"log"
	"os"
	"strings"
)

var defaultPath = "config" // 配置文件路径
var defaultName = "config" // 默认配置文件名称
var defaultExt = ".ini"    // 配置文件扩展名

type Config struct {
	Fetch           *goconfig.ConfigFile
	Path            string
	Filename        string
	Ext             string
	DefaultCallback func()
	ErrorCallback   func()
}

// LoadConfigFile 加载配置文件
func (zs *Config) LoadConfigFile(f ...func()) {
	var err error // error handle
	var file *goconfig.ConfigFile
	var filepath string

	if zs.Path == "" {
		zs.Path = defaultPath
	}

	if zs.Ext == "" {
		zs.Ext = defaultExt
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		zs.ErrorCallback()
	}

	dir = strings.Replace(dir, "\\", "/", -1) + "/" + zs.Path + "/"

	// 判断默认配置文件是否存在
	defaultFilepath := dir + defaultName + zs.Ext // 构造默认配置文件
	_, err = os.Stat(defaultFilepath)             // 判断是否存在自定义配置文件
	if err != nil {
		zs.DefaultCallback()
	}

	// 使用自定义配置文件
	if zs.Filename != "" {
		filepath = dir + "/" + zs.Filename + zs.Ext
		_, err = os.Stat(filepath) // 判断是否存在自定义配置文件
		if err == nil {
			file, err = goconfig.LoadConfigFile(defaultFilepath, filepath)
		}
	}

	// 未指定自定义配置文件或者自定义配置文件加载失败
	if zs.Filename == "" || err != nil {
		file, err = goconfig.LoadConfigFile(defaultFilepath)
	}

	// 配置文件加载失败
	if err != nil {
		log.Println(err)
		zs.ErrorCallback()
	} else {
		zs.Fetch = file
	}
	if len(f) > 0 && f[0] != nil {
		f[0]()
	}
}
