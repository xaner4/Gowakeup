package alias

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type alias struct {
	Name string `json:name`
	Mac  string `json:mac`
}

// Aliases is a slice of all the Aliases that was loaded from alias.json
var Aliases []alias

func init() {
	Load()
}

// Load all Aliases from file
func Load() (err error) {

	file, err := ioutil.ReadFile("alias.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Aliases)
	if err != nil {
		return err
	}

	return nil
}

// Write writes the Aliases from Aliases to disk
func Write() (err error) {
	data, err := json.MarshalIndent(Aliases, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("alias.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Exists(n string, m string) (exists bool, order int) {
	if len(Aliases) <= 0 {
		return false, 0
	}

	for i, a := range Aliases {
		if a.Name == n || a.Mac == m {
			return true, i
		} else {
			continue
		}
	}
	return false, 0
}

// Add adds a new alias
func Add(n string, m string) (err error) {
	if err, _ := Exists(n, m); err {
		return fmt.Errorf("alias %s or mac address %s already exists as an alias", n, m)
	}
	Aliases = append(Aliases, alias{
		Name: n,
		Mac:  m,
	})
	return nil
}

// Remove alias from file
func Remove(n string, m string) (err error) {
	exsist, order := Exists(n, m)
	if !exsist {
		return fmt.Errorf("The alias '%s' or mac '%s' does not exists ", n, m)
	}

	Aliases = append(Aliases[:order], Aliases[order+1:]...)
	return nil
}
