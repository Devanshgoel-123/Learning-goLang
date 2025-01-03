package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// here we keep the files which can't be imported by others
type HTTPServer struct{
	Address string `yaml:"address" env-required:"true"`
}

type Config struct{ 
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}


func MustLoad() *Config{
	var configPath string
	configPath=os.Getenv("CONFIG_PATH")

	if configPath==""{
		flags:=flag.String("config","","path to the configuration file")
		flag.Parse();
		log.Output(0,*flags)
		configPath=*flags
		if configPath==""{
			log.Fatal("Config path is not set")
		}
	}
    //check if the file is available at this path
	if _,err :=os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("config file doesnt exist at the given path: %s",configPath);
	}

	var cfg Config
	err:=cleanenv.ReadConfig(configPath,&cfg)
	if err!=nil{
		log.Fatalf("can not read config file %s",err.Error())
	}

	return &cfg

}