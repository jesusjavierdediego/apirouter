package configuration

import (
	"log"
	"os"
	"reflect"
	"strings"
	"github.com/spf13/viper"
)

var GlobalConfiguration Configuration

type Configuration struct {
	Profile string
	Kafka kafka
	Rest rest
	Grpcclient grpclient
	Grpcresponse grpcresponse
	Gitserver gitserver
	Admindb admindb
}

type admindb struct {
	Host string
	Port int
	Dbname string
	Username string
	Password string
}

type kafka struct {
	Bootstrapserver string
	Groupid string
  	Sessiontimeout int
  	Eventschannelenabled bool
	Rebalanceenabled bool
	Partitioneofenabled bool
	Autooffset string
	Githandlerecordtopic string
	Githandlebatchtopic string
	Gitactionbacktopic string
    Gitdeletetopic string
	Alerttopic string
}


type rest struct {
	Port int
	Mode string  
	Path string
	Apikeyheadername string
}

type grpclient struct {
	Gitreaderhost string
	Gitreaderport int
	Rdbreaderhost string
	Rdbreaderport int
	Timeout int
}

type grpcresponse struct {
	Ok string
	Cancelled string
	Invalid string
	Notfound string
	Permissiondenied string
	Unauthenticated string
	Notimplemented string
	Internal string
	Notavailable string
	Alreadyexists string
}

type gitserver struct {
	Url string
	Authtoken string
	Username string
	Password string
	Email string
	Strategy strategy
}

type strategy struct {
	Timeout int
	Deletebranchaftermerge bool
	Forcemerge bool
}

type ownerMap struct {
	
}

func init() {
	GlobalConfiguration = initConfiguration()
}

/*
InitConfiguration Return the configuration
Read the default configuration from application.yml. If PROFILE=dev then use application.dev.xml

To override default param must be run with system ENV, follow the same structure of of yaml, but points is replace by _
example:

	server:
  		port: 8080
  		name: "API-endpoint"
  		timeout: 5
  		key:
		mode: debug

To override the port SERVER_PORT=8181

SERVER_PORT=8181 go run main.go
*/
func initConfiguration() Configuration {
	var configuration Configuration

	profile := os.Getenv("PROFILE")
	//ENV VARS
	viper.AutomaticEnv()                                   // Automatic binding from enviroment
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // to charge enviroment

	if strings.ToLower(profile) == "dev" {
		viper.SetConfigName("application.dev")
	} else {
		viper.SetConfigName("application")
	}

	viper.SetConfigType("yaml")
	path := calculatePath("resources")

	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ERROR: Error reading config file, %s", err)
	} else {
		err := viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("ERROR: unable to decode into struct, %v", err)
		} else {
			log.Printf("Internal configuration loaded OK")
		}
	}
	return configuration
}
func Reload() {
	GlobalConfiguration = initConfiguration()
}

/*
calculatePath get the configuration path relative to package of configuration and the currentDir of execution
*/
func calculatePath(resourcesPath string) string {

	configurationPatch := reflect.TypeOf(Configuration{}).PkgPath()
	index := strings.LastIndex(configurationPatch, "/")
	configurationPatch = configurationPatch[0:index]

	currentDir, _ := os.Getwd()
	index = strings.LastIndex(currentDir, configurationPatch)
	if index > 0 {
		currentDir = currentDir[0:index]
	}
	currentDir = currentDir + configurationPatch + "/" + resourcesPath
	fileInfo, _ := os.Lstat(currentDir)
	if fileInfo == nil {
		currentDir = "/" + resourcesPath
		fileInfo, _ = os.Lstat(currentDir)
		if fileInfo == nil {
			currentDir = "/"
		}
	}

	return currentDir
}

