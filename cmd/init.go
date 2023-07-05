package cmd

import (
	"bytes"
	"errors"
	"go/format"
	"log"
	"strings"
	"text/template"

	"github.com/Arsfiqball/code-generator/cmd/helper"
	"github.com/ettle/strcase"
	"github.com/spf13/cobra"
)

type initConfigData struct {
	Name   string
	Type   string
	Source string
}

type initData struct {
	PkgName         string
	StructName      string
	Configs         []initConfigData
	UseFiber        bool
	UseWorkJob      bool
	UseWatermillSub bool
}

func createMainGo(idata initData) error {
	mainTpl, err := Templates.ReadFile("cmd/templates/main.tpl")
	if err != nil {
		return errors.New(err.Error())
	}

	t, err := template.New("main.go").Parse(string(mainTpl))
	if err != nil {
		return errors.New(err.Error())
	}

	var buf bytes.Buffer

	var data = struct{ initData }{
		initData: idata,
	}

	if err = t.ExecuteTemplate(&buf, "main.go", data); err != nil {
		return errors.New(err.Error())
	}

	fileBytes, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err.Error())
	}

	return helper.SaveFile(fileBytes, "pkg/"+idata.PkgName, "main.go")
}

func createWireGo(idata initData) error {
	mainTpl, err := Templates.ReadFile("cmd/templates/wire.tpl")
	if err != nil {
		return errors.New(err.Error())
	}

	t, err := template.New("wire.go").Parse(string(mainTpl))
	if err != nil {
		return errors.New(err.Error())
	}

	var buf bytes.Buffer

	var data = struct{ initData }{
		initData: idata,
	}

	if err = t.ExecuteTemplate(&buf, "wire.go", data); err != nil {
		return errors.New(err.Error())
	}

	fileBytes, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err.Error())
	}

	return helper.SaveFile(fileBytes, "pkg/"+idata.PkgName, "wire.go")
}

var initCmd = &cobra.Command{
	Use:   "init <name>",
	Short: "Init a new feature",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("name must be provided")
		}

		idata := initData{
			StructName: strcase.ToPascal(args[0]),
			PkgName:    strings.ToLower(strcase.ToPascal(args[0])),
			Configs: []initConfigData{
				{
					Name:   "Tracer",
					Type:   "trace.Tracer",
					Source: "go.opentelemetry.io/otel/trace",
				},
			},
		}

		flags := cmd.Flags()

		drivenLibs := map[string]initConfigData{
			"zap": {
				Name:   "Logger",
				Type:   "*zap.Logger",
				Source: "go.uber.org/zap",
			},
			"viper": {
				Name:   "Config",
				Type:   "*viper.Viper",
				Source: "github.com/spf13/viper",
			},
			"gorm": {
				Name:   "Database",
				Type:   "*gorm.DB",
				Source: "gorm.io/gorm",
			},
			"mongo": {
				Name:   "MongoClient",
				Type:   "*mongo.Client",
				Source: "go.mongodb.org/mongo-driver/mongo",
			},
			"wmpub": {
				Name:   "Publisher",
				Type:   "message.Publisher",
				Source: "github.com/ThreeDotsLabs/watermill/message",
			},
			"worken": {
				Name:   "Enqueuer",
				Type:   "*work.Enqueuer",
				Source: "github.com/gocraft/work",
			},
		}

		for f, lib := range drivenLibs {
			use, err := flags.GetBool(f)
			if err != nil {
				return err
			}

			if use {
				idata.Configs = append(idata.Configs, lib)
			}
		}

		useFiber, err := flags.GetBool("fiber")
		if err != nil {
			return err
		}

		useWorkJob, err := flags.GetBool("workjob")
		if err != nil {
			return err
		}

		useWatermillSub, err := flags.GetBool("wmsub")
		if err != nil {
			return err
		}

		idata.UseFiber = useFiber
		idata.UseWorkJob = useWorkJob
		idata.UseWatermillSub = useWatermillSub

		if err := createMainGo(idata); err != nil {
			return err
		}

		if err := createWireGo(idata); err != nil {
			return err
		}

		if err := helper.GoModTidy(); err != nil {
			return err
		}

		if err := helper.Wire("./pkg/" + idata.PkgName); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Driven dependency
	initCmd.Flags().BoolP("zap", "", false, "Add zap logger")
	initCmd.Flags().BoolP("viper", "", false, "Add viper config")
	initCmd.Flags().BoolP("gorm", "", false, "Add gorm database driver")
	initCmd.Flags().BoolP("mongo", "", false, "Add mongo database client")
	initCmd.Flags().BoolP("wmpub", "", false, "Add watermill publisher")
	initCmd.Flags().BoolP("worken", "", false, "Add work background task enqueuer")

	// Driver dependency
	initCmd.Flags().BoolP("fiber", "", false, "Add fiber router code")
	initCmd.Flags().BoolP("wmsub", "", false, "Add watermill subscriber code")
	initCmd.Flags().BoolP("workjob", "", false, "Add work background task job handler code")
}
