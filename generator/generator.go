package generator

import (
	"fmt"

	"github.com/diegoall/rest-api-golang-gen/models"
)

type Attribute struct {
	TipoDato string `json:"tipoDato"`
}

type Entity struct {
	Tipo      string               `json:"tipo"`
	Atributos map[string]Attribute `json:"atributos"`
}

func Generate(projectName string, dbType string, config models.Config, dummy bool) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)
	fmt.Print("\n")

	// fmt.Println("Config from Generate (output python): ", config)

	// for _, trin := range config.MatrizAtributos {
	// 	fmt.Println(trin)
	// }

	//fmt.Println(config.Tipo)

	// The NaturalID used in generateEntityModels is calculated
	tiposGenerados := generateHandlers(config.Tipo, config.MatrizAtributos)
	//fmt.Println("Longitud de tiposGenerados: (generator/Generate)", len(tiposGenerados))

	//generatedDatabaseDDL := generateDatabaseDDL(config.Tipo, config.MatrizAtributos, dummy)
	//fmt.Println("El DDL es: \n", generatedDatabaseDDL)
	generateDatabaseDDL(config.Tipo, config.MatrizAtributos, dummy)

	//generatedModels := generateEntityModels(config.Tipo, config.MatrizAtributos)
	generateEntityModels(config.Tipo, config.MatrizAtributos)
	//fmt.Println("Generated Models es: ", generatedModels)

	modifyBaseTemplates(tiposGenerados)

	// Generate folder structure
	createFolderStructure(projectName, config.Tipo, config.MatrizAtributos)

}
