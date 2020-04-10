package configfx

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"strings"
)

// Initialize viper for reading configurations
func InitConfig(svcName, cfgFile, cfgPath string) fx.Option {
	return fx.Invoke(func() {
		fmt.Printf("[ConfigFx] Load config %s\n", svcName)
		loadConfig(svcName, cfgFile, cfgPath)
	})
}

// loadConfig loads configurations from config file/environment
func loadConfig(svcName, cfgFile, cfgPath string) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()         // Tells viper to look at the Environment Variables.
	viper.SetConfigName(cfgFile) // name of config file (without extension)
	//viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(cfgPath)                           // path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", svcName)) // call multiple times to add many search paths
	viper.AddConfigPath(".")                               // optionally look for config in the working directory
	err := viper.ReadInConfig()                            // Find and read the config file
	if err != nil {                                        // Handle errors reading the config file
		log.Panicf("[ConfigFx] Can't load configuration from local file system, error: %v", err)
		return
	}
}
