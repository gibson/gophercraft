package dbdefs

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cybriq/gotiny"
	"github.com/gibson/gophercraft/format/dbc/dbd"
)

var (
	All map[string]*dbd.Definition
)

func init() {
	reader := bytes.NewReader(packedDefinitions)

	zr, err := zlib.NewReader(reader)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(zr)
	if err != nil {
		panic(err)
	}

	gotiny.Unmarshal(b, &All)
}

func Lookup(name string) (*dbd.Definition, error) {
	index := strings.ToLower(name)
	def, ok := All[index]
	if !ok {
		return nil, fmt.Errorf("dbdefs: could not find %s definition", name)
	}

	return def, nil
}

func Register(def *dbd.Definition) (alreadyRegistered bool) {
	index := strings.ToLower(def.Name)

	_, alreadyRegistered = All[index]
	All[index] = def
	return
}
