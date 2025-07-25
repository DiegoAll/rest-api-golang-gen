# rest-api-golang-gen
Generate a REST API scaffold with Golang from domain model specification.

Soon, it will offer features to create a secure REST API.


<img src="rest-api-golang-gen.gif" align="center"/>


## Requirements

Golang 1.21 or newer and Python 3.7 or later are required.


## Getting started


### Run locally

    git clone https://github.com/diegoall/rest-api-golang-gen.git && cd rest-api-golang-gen

#### Build the project

    go build -o rest-api-golang-gen .  && ./rest-api-golang-gen init --db sqlite --dummy --config inputs/product.json exampleAPI

#### Only run the project

    go run github.com/diegoall/rest-api-golang-gen@v0.1.13 init --db sqlite --dummy --config inputs/product.json exampleAPI
    go run main.go init --db sqlite --dummy --config inputs/product.json exampleAPI

Note: To use the --dummy flag, configure a Gemini API Key. If you don't need this feature or don't have an API Key, you can omit the flag.


### Run from remote repository

    Under construction


<details><summary><code> Input JSON example </code></summary>

### --config parameter

This is the input to define the model to be generated. The path of a JSON file must be assigned with the following structure:

    [
        {
          "tipo": "Product",
          "atributos": {
            "name": {
              "tipoDato": "string"
            },
            "description": {
              "tipoDato": "string"
            },
            "price": {
              "tipoDato": "int"
            },
            "quantity": {
              "tipoDato": "int"
            }
          }
        }
    ]

</summary></details>

## Setup and usage

<details><summary><code> Available Commands   </code></summary>

##
Available commands

    go run github.com/diegoall/rest-api-golang-gen@latest init

    Available Commands:
      completion  Generate the autocompletion script for the specified shell
      help        Help about any command
      init        Inicializa un nuevo proyecto
      rollback    Restaura los archivos genéricos a partir de los archivos base

</summary></details>

-----------------------------------------------------------
<details><summary><code> Create a new project   </code></summary>

##
Options

    go run github.com/diego-all/rest-api-golang-gen@latest init -h

    Flags:
      -c, --config string   Ruta del archivo JSON de configuración
      -d, --db string       Tipo de base de datos (requerido)
      -u, --dummy           Generar Dummy data usando Gemini (Requiere API Key)
      -h, --help            help for init


</summary></details>

-----------------------------------------------------------

<details><summary><code> Generate dummy data using Gemini  </code></summary>

##
Generate Gemini API Key [Gemini](https://aistudio.google.com/app/apikey)

</summary></details>

-----------------------------------------------------------
<details><summary><code> Create the database </code></summary>

## 
You need to have sqlite3 installed on your computer.

    sh create-db.sh

</summary></details>

-----------------------------------------------------------
<details><summary><code> Rollback </code></summary>

##
After using genotype the templates must be reset for next use.

    go run main.go rollback

</summary></details>
    

## Generated project's structure


    ├── cmd
    │   └── api
    │       ├── handlers.go
    │       ├── handlers-Product.go
    │       ├── main.go
    │       ├── routes.go
    │       └── util.go
    ├── database
    │   ├── connection.go
    │   ├── create-db.sh
    │   └── up.sql
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── models.go
    │   └── Products.go
    ├── README.md
    └── requests.md

## Generated API use

    go run ./cmd/api
    go build ./cmd/api
    go build -o exampleAPI ./cmd/api


