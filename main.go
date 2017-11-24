package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"text/template"
	"time"

	// "github.com/fatih/structs"
	// "github.com/k0kubun/pp"
	"github.com/MohamedBassem/gormgen"
	"github.com/iancoleman/strcase"
	"github.com/roscopecoltran/aor-gin-swagger/utils"
	"gopkg.in/alecthomas/kingpin.v2"
)

//go:generate go-bindata -pkg main -o bindata.go config.yaml templates

const (
	swaggerTmplConfig       = "config.yaml"
	defaultSwaggerPath      = "./swagger.json"
	defaultOutputPrefixPath = "./generated"
)

var (
	defaultOutputDir = path.Join(defaultOutputPrefixPath, "new-app")
	assetDirFmt      = path.Join(os.TempDir(), "gin-swagger-%d")
	outputDirFmt     string
	config           struct {
		Application        string
		SwaggerPath        string
		OutputDir          string
		DatabaseMiddleware bool
		GenerateServer     bool
		Verbose            bool
		ConvertSwagger     bool
	}
)

func init() {

	kingpin.Flag("application", "Name of the application (passed directly to swagger).").
		Required().Short('A').StringVar(&config.Application)

	kingpin.Flag("verbose", "Verbose generation process.").
		Short('V').BoolVar(&config.Verbose)

	kingpin.Flag("convert-swagger", "Convert Swagger specs format (JSON2YAML, or YAML2JSON).").
		Short('C').BoolVar(&config.ConvertSwagger)

	kingpin.Flag("output-dir", "Output directory for generated files.").
		Short('O').Default(defaultOutputDir).StringVar(&config.OutputDir)

	kingpin.Flag("generate-server", "Generate server command files.").
		Short('S').BoolVar(&config.GenerateServer)

	kingpin.Flag("database-middleware", "Add a gin database middleware.").
		Short('D').BoolVar(&config.DatabaseMiddleware)

	kingpin.Flag("spec", "the spec file to use.").
		Short('f').Default(defaultSwaggerPath).StringVar(&config.SwaggerPath)

	kingpin.Parse()
}

func main() {
	outputDirFmt = path.Join(defaultOutputPrefixPath, strcase.ToSnake(config.Application))
	fmt.Println("defaultOutputPrefixPath: ", defaultOutputPrefixPath)
	fmt.Println("outputDirFmt: ", outputDirFmt)
	fmt.Println("ConvertSwagger: ", config.ConvertSwagger)
	fmt.Println("SwaggerPath: ", config.SwaggerPath)

	if exists, _ := utils.IsDirectory(outputDirFmt); !exists {
		os.MkdirAll(outputDirFmt, 0755)
		fmt.Println("outputDirFmt created, path=", exists)
	}

	if config.ConvertSwagger {
		utils.ConvertSwaggerFile(config.SwaggerPath)
	}

	err := run(config.Application, config.SwaggerPath, config.DatabaseMiddleware)
	if err != nil {
		log.Fatalf("failed to run swagger: %s", err)
	}
}

func run(application, swagger string, databaseMiddleware bool) error {
	tmpDir, err := writeAssets()
	if err != nil {
		return err
	}
	defer func() error {
		return os.RemoveAll(tmpDir)
	}()

	fmt.Println("tmpDir: ", tmpDir)

	cmd := exec.Command("swagger",
		"generate",
		"server",
		"-A",
		application,
		"-t",
		outputDirFmt,
		"-f",
		swagger,
		"-C",
		path.Join(tmpDir, swaggerTmplConfig),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func getWorkDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return wd, nil
}

func writeGormHelpers(wd string) {
	if wd == "" {
		wd, _ = getWorkDir()
	}
	parser := gormgen.NewParser()
	parser.ParseDir(wd)
}

func writeAssets() (string, error) {
	tmp := fmt.Sprintf(assetDirFmt, time.Now().UTC().UnixNano())
	err := RestoreAssets(tmp, "templates")
	if err != nil {
		return "", err
	}

	d, err := Asset(swaggerTmplConfig)
	if err != nil {
		return "", err
	}

	t, err := template.New(swaggerTmplConfig).Parse(string(d))
	if err != nil {
		return "", err
	}

	// UseGorm

	/*
		isStruct := structs.IsStruct(t)
		fmt.Printf("isStruct? %t \n", isStruct)
		if isStruct {
			names := structs.Names(t)
			pp.Println("names: ", names)

			f := structs.Fields(t)
			pp.Println("fields: ", f)

			v := structs.Values(t)
			pp.Println("values: ", v)

			m := structs.Map(t)
			pp.Println("map: ", m)
		}
	*/

	fd, err := os.Create(path.Join(tmp, swaggerTmplConfig))
	if err != nil {
		return "", err
	}
	defer fd.Close()

	w := bufio.NewWriter(fd)

	err = t.Execute(w, map[string]string{"TmpDir": tmp})
	if err != nil {
		return "", err
	}

	err = w.Flush()
	if err != nil {
		return "", err
	}

	return tmp, nil
}

/*

Usage:
  swagger [OPTIONS] generate server [server-OPTIONS]

generate all the files for a server application

Help Options:
  -h, --help                                         Show this help message

[server command options]
      -f, --spec=                                    the spec file to use (default swagger.{json,yml,yaml})
      -a, --api-package=                             the package to save the operations (default: operations)
      -m, --model-package=                           the package to save the models (default: models)
      -s, --server-package=                          the package to save the server specific code (default: restapi)
      -c, --client-package=                          the package to save the client specific code (default: client)
      -t, --target=                                  the base directory for generating the files (default: ./)
      -T, --template-dir=                            alternative template override directory
      -C, --config-file=                             configuration file to use for overriding template options
      -r, --copyright-file=                          copyright file used to add copyright header
          --existing-models=                         use pre-generated models e.g. github.com/foobar/model
      -A, --name=                                    the name of the application, defaults to a mangled value of info.title
      -O, --operation=                               specify an operation to include, repeat for multiple
          --tags=                                    the tags to include, if not specified defaults to all
      -P, --principal=                               the model to use for the security principal
          --default-scheme=                          the default scheme for this API (default: http)
      -M, --model=                                   specify a model to include, repeat for multiple
          --skip-models                              no models will be generated when this flag is specified
          --skip-operations                          no operations will be generated when this flag is specified
          --skip-support                             no supporting files will be generated when this flag is specified
          --exclude-main                             exclude main function, so just generate the library
          --exclude-spec                             don't embed the swagger specification
          --with-context                             handlers get a context as first arg (deprecated)
          --dump-data                                when present dumps the json for the template generator instead of generating files
          --flag-strategy=[go-flags|pflag]           the strategy to provide flags for the server (default: go-flags)
          --compatibility-mode=[modern|intermediate] the compatibility mode for the tls server (default: modern)
          --skip-validation                          skips validation of spec prior to generation
          --skip-flatten                             skips flattening of spec prior to generation

*/
