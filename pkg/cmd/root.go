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
	version   = "v0.0.4"
)

func init() {
	var err error
	initConfig()

	persistentFlags := map[string]string{
		"apiKey":   "API Key",
		"apiUser":  "API User",
		"clientIP": "Client IP",
	}

	for k, v := range persistentFlags {
		rootCmd.PersistentFlags().String(k, "", v)
		err = rootCmd.MarkPersistentFlagRequired(k)
		if err != nil {
			panic(err)
		}

		err = viper.BindPFlag(k, rootCmd.PersistentFlags().Lookup(k))
		if err != nil {
			panic(err)
		}
	}

	rootCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
			err = rootCmd.PersistentFlags().Set(f.Name, viper.GetString(f.Name))
			if err != nil {
				panic(err)
			}
		}
	})

	rootCmd.PersistentFlags().Bool("sandbox", false, "")
	err = rootCmd.PersistentFlags().MarkHidden("sandbox")
	if err != nil {
		panic(err)
	}

	cobra.OnInitialize(func() {
		url := "https://api.namecheap.com/xml.response"

		sandbox, err := rootCmd.PersistentFlags().GetBool("sandbox")
		if err != nil {
			panic(err)
		} else if sandbox {
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

func addCommand(parent *cobra.Command, fn func() (*cobra.Command, error)) {
	cmd, err := fn()
	if err != nil {
		panic(err)
	}

	parent.AddCommand(cmd)
}

func Execute() {
	_ = rootCmd.Execute()
}
