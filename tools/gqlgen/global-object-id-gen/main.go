package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	idPackage    = "github.com/sky0621/wht/adapter/web/gqlmodel."
	templateFile = "./tools/gqlgen/global-object-id-gen/scalar_id_go.tmpl"
	outputFile   = "./src/adapter/web/gqlmodel/scalar_%s.go"
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("gqlgen")
	viper.AddConfigPath("./src")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	t := template.Must(template.ParseFiles(templateFile))

	for _, id := range targetIDPackages() {
		if err := exec(id, t); err != nil {
			log.Printf("failed to create for id:%s ... %+v", id, err)
			continue
		}
	}
}

func targetIDPackages() []string {
	idPackages := viper.Sub("models").Sub("id").GetStringSlice("model")

	var results []string
	for _, idPkg := range idPackages {
		fmt.Println(idPkg)
		if strings.Contains(idPkg, idPackage) {
			results = append(results, strings.Trim(idPkg, idPackage))
		}
	}

	return results
}

func exec(id string, t *template.Template) error {
	f, err := os.OpenFile(fmt.Sprintf(outputFile, id), os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer func() {
		if f != nil {
			if err := f.Close(); err != nil {
				log.Printf("failed to close file: %+v", err)
			}
		}
	}()

	if err := t.Execute(f, id); err != nil {
		return err
	}

	return nil
}
