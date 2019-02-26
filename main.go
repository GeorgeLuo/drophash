package main

import (
   "github.com/labstack/echo"
   "github.com/labstack/echo/middleware"
   "github.com/sirupsen/logrus"

   "fmt"
   "encoding/json"
   "io/ioutil"
   "os"
   "strconv"
   "github.com/GeorgeLuo/drophash/environment"
)

var log = logrus.New()
var PortNoteChannelMap = make(map[string]string)

var Configuration DropHashConfiguration

/**
   initialize DropHash API
**/

var DROPHASH_DOCKER_CONFIGURATION_PATH = "/go/src/github.com/GeorgeLuo/drophash/config/drophash_docker_config.json"
var DROPHASH_LOCAL_CONFIGURATION_PATH = "/usr/local/conf/drophash_local_config.json"

func init() {
   ReadConfigurations(DROPHASH_DOCKER_CONFIGURATION_PATH)

   /** 
      Uncomment below to run server on localized host. 
      Either place provided drophash_local_config.json file from drophash/config
      into /usr/local/conf/ or define your own DROPHASH_LOCAL_CONFIGURATION_PATH above
   **/

   // ReadConfigurations(DROPHASH_LOCAL_CONFIGURATION_PATH)
   
   if Configuration.LocalMemMode == 0 {
      err := environment.InitializeDatabaseClient(Configuration.DatabasePath, 
         Configuration.DatabaseName, Configuration.CollectionName)
      if err != nil {
         fmt.Println("error initializing database session, using local memory")
         environment.LocalMemoryMode = true
      } else {
         environment.LocalMemoryMode = false
      }
   } else {
      fmt.Println("initialized in local_mem_mode")
      environment.LocalMemoryMode = true
   }

	// initialize logger
   log.Formatter = new(logrus.JSONFormatter)
   log.Formatter = new(logrus.TextFormatter) 
   log.Level = logrus.DebugLevel

}

func main() {
   e := echo.New()
   e.Use(middleware.Logger())
   e.Use(middleware.Recover())


   // Define api paths
   e.POST("/messages", environment.PostMessageAndReturnHash)
   e.GET("/messages/:hash", environment.GetMessageFromHash)

   // start listening
   e.Logger.Fatal(e.Start(":" + strconv.Itoa(Configuration.ServicePort)))
}

func ReadConfigurations(mapping_config_file string) {

   jsonFile, err := os.Open(mapping_config_file)
   if err != nil {
      fmt.Println(err)
   }

   defer jsonFile.Close()
   byteValue, _ := ioutil.ReadAll(jsonFile)
   json.Unmarshal(byteValue, &Configuration)
}

type DropHashConfiguration struct {
   DatabasePath string `json:"database_path"`
   ServicePort int `json:"service_port"`
   DatabaseName string `json:"database_name"`
   CollectionName string `json:"collection_name"`
   LocalMemMode int `json:"local_mem_mode"`
}