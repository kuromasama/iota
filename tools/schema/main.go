// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/iotaledger/wasp/tools/schema/generator"
	"gopkg.in/yaml.v2"
)

var (
	disabledFlag = false
	flagCore     = flag.Bool("core", false, "generate core contract interface")
	flagForce    = flag.Bool("force", false, "force code generation")
	flagGo       = flag.Bool("go", false, "generate Go code")
	flagInit     = flag.String("init", "", "generate new schema file for smart contract named <string>")
	flagJava     = &disabledFlag // flag.Bool("java", false, "generate Java code <outdated>")
	flagRust     = flag.Bool("rust", false, "generate Rust code")
	flagTs       = flag.Bool("ts", false, "generate TypScript code")
	flagType     = flag.String("type", "yaml", "type of schema file that will be generated. Values(yaml,json)")
)

func init() {
	flag.Parse()
}

func main() {
	err := generator.FindModulePath()
	if err != nil && *flagGo {
		fmt.Println(err)
		return
	}

	if *flagCore {
		generateCoreInterfaces()
		return
	}

	file, err := os.Open("schema.yaml")
	if err != nil {
		file, err = os.Open("schema.json")
	}
	if err == nil {
		defer file.Close()
		if *flagInit != "" {
			fmt.Println("schema definition file already exists")
			return
		}
		err = generateSchema(file)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if *flagInit != "" {
		err = generateSchemaNew()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	flag.Usage()
}

func generateCoreInterfaces() {
	err := filepath.WalkDir("interfaces", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".yaml") {
			return generateCoreCoreInterface(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func generateCoreCoreInterface(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return generateSchema(file)
}

func generateSchema(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	schemaTime := info.ModTime()

	s, err := loadSchema(file)
	if err != nil {
		return err
	}
	s.CoreContracts = *flagCore

	if *flagTs {
		g := generator.NewTypeScriptGenerator(s)
		err = setFolder(s, "ts", true)
		if err != nil {
			return err
		}
		info, err = os.Stat(s.Folder + "consts.ts")
		if err == nil && info.ModTime().After(schemaTime) && !*flagForce {
			fmt.Println("skipping AssemblyScript code generation")
		} else {
			fmt.Println("generating AssemblyScript code")
			err = g.GenerateTs()
			if err != nil {
				return err
			}
			if !s.CoreContracts {
				err = g.GenerateGoTests()
				if err != nil {
					return err
				}
			}
		}
	}

	if *flagGo {
		g := generator.NewGoGenerator(s)
		err = setFolder(s, "go", true)
		if err != nil {
			return err
		}
		info, err = os.Stat(s.Folder + "consts.go")
		if err == nil && info.ModTime().After(schemaTime) && !*flagForce {
			fmt.Println("skipping Go code generation")
		} else {
			fmt.Println("generating Go code")
			err = g.GenerateGo()
			if err != nil {
				return err
			}
			if !s.CoreContracts {
				err = g.GenerateGoTests()
				if err != nil {
					return err
				}
			}
		}
	}

	if *flagJava {
		fmt.Println("generating Java code")
		err = s.GenerateJava()
		if err != nil {
			return err
		}
	}

	if *flagRust {
		g := generator.NewRustGenerator(s)
		err = setFolder(s, "src", false)
		if err != nil {
			return err
		}
		info, err = os.Stat(s.Folder + "consts.rs")
		if err == nil && info.ModTime().After(schemaTime) && !*flagForce {
			fmt.Println("skipping Rust code generation")
		} else {
			fmt.Println("generating Rust code")
			err = g.GenerateRust()
			if err != nil {
				return err
			}
			if !s.CoreContracts {
				err = g.GenerateGoTests()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func setFolder(s *generator.Schema, root string, moduleName bool) error {
	s.Folder = root + "/"
	if moduleName {
		module := strings.ReplaceAll(generator.ModuleCwd, "\\", "/")
		module = module[strings.LastIndex(module, "/")+1:]
		s.Folder += module + "/"
	}
	if s.CoreContracts {
		s.Folder += s.Name + "/"
	}
	return os.MkdirAll(s.Folder, 0o755)
}

func generateSchemaNew() error {
	name := *flagInit
	fmt.Println("initializing " + name)

	subfolder := strings.ToLower(name)
	err := os.Mkdir(subfolder, 0o755)
	if err != nil {
		return err
	}
	err = os.Chdir(subfolder)
	if err != nil {
		return err
	}

	schemaDef := &generator.SchemaDef{}
	schemaDef.Name = name
	schemaDef.Description = name + " description"
	schemaDef.Structs = make(generator.StringMapMap)
	schemaDef.Typedefs = make(generator.StringMap)
	schemaDef.State = make(generator.StringMap)
	schemaDef.State["owner"] = "AgentID // current owner of this smart contract"
	schemaDef.Funcs = make(generator.FuncDefMap)
	schemaDef.Views = make(generator.FuncDefMap)

	funcInit := &generator.FuncDef{}
	funcInit.Params = make(generator.StringMap)
	funcInit.Params["owner"] = "AgentID? // optional owner of this smart contract"
	schemaDef.Funcs["init"] = funcInit

	funcSetOwner := &generator.FuncDef{}
	funcSetOwner.Access = "owner // current owner of this smart contract"
	funcSetOwner.Params = make(generator.StringMap)
	funcSetOwner.Params["owner"] = "AgentID // new owner of this smart contract"
	schemaDef.Funcs["setOwner"] = funcSetOwner

	viewGetOwner := &generator.FuncDef{}
	viewGetOwner.Results = make(generator.StringMap)
	viewGetOwner.Results["owner"] = "AgentID // current owner of this smart contract"
	schemaDef.Views["getOwner"] = viewGetOwner
	switch *flagType {
	case "json":
		return WriteJSONSchema(schemaDef)
	case "yaml":
		return WriteYAMLSchema(schemaDef)
	}
	return errors.New("invalid schema type: " + *flagType)
}

func WriteJSONSchema(schemaDef *generator.SchemaDef) error {
	file, err := os.Create("schema.json")
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.Marshal(schemaDef)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return err
	}

	_, err = out.WriteTo(file)
	return err
}

func WriteYAMLSchema(schemaDef *generator.SchemaDef) error {
	file, err := os.Create("schema.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := yaml.Marshal(schemaDef)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}

func loadSchema(file *os.File) (s *generator.Schema, err error) {
	fmt.Println("loading " + file.Name())
	schemaDef := &generator.SchemaDef{}
	switch filepath.Ext(file.Name()) {
	case ".json":
		err = json.NewDecoder(file).Decode(schemaDef)
		if err == nil && *flagType == "convert" {
			err = WriteYAMLSchema(schemaDef)
		}
	case ".yaml":
		fileByteArray, _ := ioutil.ReadAll(file)
		err = yaml.Unmarshal(fileByteArray, schemaDef)
		if err == nil && *flagType == "convert" {
			err = WriteJSONSchema(schemaDef)
		}
	default:
		err = errors.New("unexpected file type: " + file.Name())
	}
	if err != nil {
		return nil, err
	}

	s = generator.NewSchema()
	err = s.Compile(schemaDef)
	if err != nil {
		return nil, err
	}
	return s, nil
}
