package cmd

import (
	"github.com/jsgv/namecheap-go/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	rootCmd   = newRootCmd()
	apiClient *api.Client
	version   = "v0.0.2"
)

func init() {
	initConfig()

	persistentFlags := map[string]string{
		"apiKey":   "API Key",
		"apiUser":  "API User",
		"clientIP": "Client IP",
	}

	for k, v := range persistentFlags {
		rootCmd.PersistentFlags().String(k, "", v)
		rootCmd.MarkPersistentFlagRequired(k)
		viper.BindPFlag(k, rootCmd.PersistentFlags().Lookup(k))
	}

	rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
			rootCmd.PersistentFlags().Set(f.Name, viper.GetString(f.Name))
		}
	})

	rootCmd.PersistentFlags().Bool("sandbox", false, "")
	rootCmd.PersistentFlags().MarkHidden("sandbox")

	cobra.OnInitialize(func() {
		url := "https://api.namecheap.com/xml.response"

		sandbox, err := rootCmd.PersistentFlags().GetBool("sandbox")
		if err != nil {
			panic(err)
		} else if sandbox == true {
			url = "https://api.sandbox.namecheap.com/xml.response"
		}

		apiClient = api.NewClient(
			url,
			viper.GetString("apiKey"),
			viper.GetString("apiUser"),
			viper.GetString("clientIP"),
		)
	})

	rootCmd.AddCommand(newCmdDomains())
	rootCmd.AddCommand(newCmdVersion())
}

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "namecheap",
		Short: "A CLI to help you integrate with Namecheap",
	}
}

func initConfig() {
	viper.AddConfigPath("$HOME/.config/namecheap")
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func addCommand(parent *cobra.Command, child *cobra.Command) {
	parent.AddCommand(child)
}

func Execute() {
	rootCmd.Execute()
}
