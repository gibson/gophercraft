package wizard

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/gibson/gophercraft/datapack"
	"github.com/gibson/gophercraft/format/dbc"
	"github.com/gibson/gophercraft/format/dbc/dbdefs"
	"github.com/gibson/gophercraft/utils/log"
	"github.com/gibson/gophercraft/utils/text"
	"github.com/gibson/gophercraft/vsn"
)

const DBPackName = "!db.zip"

type ExDb struct {
	Name string
	Type reflect.Type
}

// Return the list of DB files required to produce a base datapack
func (ex *Extractor) neededDBs() []ExDb {
	var need []ExDb

	need = []ExDb{
		{"AreaTable", reflect.TypeOf(dbdefs.Ent_AreaTable{})},
		{"AreaTrigger", reflect.TypeOf(dbdefs.Ent_AreaTrigger{})},
		{"ChrRaces", reflect.TypeOf(dbdefs.Ent_ChrRaces{})},
		{"ChrClasses", reflect.TypeOf(dbdefs.Ent_ChrClasses{})},
		{"CharStartOutfit", reflect.TypeOf(dbdefs.Ent_CharStartOutfit{})},
		{"CreatureFamily", reflect.TypeOf(dbdefs.Ent_CreatureFamily{})},
		{"EmotesText", reflect.TypeOf(dbdefs.Ent_EmotesText{})},
		{"Map", reflect.TypeOf(dbdefs.Ent_Map{})},
		{"Spell", reflect.TypeOf(dbdefs.Ent_Spell{})},
		{"SpellCastTimes", reflect.TypeOf(dbdefs.Ent_SpellCastTimes{})},
		{"SpellDuration", reflect.TypeOf(dbdefs.Ent_SpellDuration{})},
		{"SpellVisual", reflect.TypeOf(dbdefs.Ent_SpellVisual{})},
	}

	return need
}

func (ex *Extractor) ExtractDatabases() error {
	const tempPackDir = "x-db"

	if ex.packExists(DBPackName) {
		ex.removePack(DBPackName)
	}

	exAuthor := fmt.Sprintf("Gophercraft Wizard %s", vsn.GophercraftVersion)

	pack, err := ex.AuthorPack(tempPackDir, datapack.PackConfig{
		Name:        "Gophercraft Base Databases",
		Author:      exAuthor,
		Description: "The base databases required by Gophercraft Core " + ex.generationNotice(),
		Version:     "1.0",
		Depends:     ex.dependencies(),
	})

	if err != nil {
		return err
	}

	need := ex.neededDBs()

	pb := log.NewIntProgressBar("Extracting databases...", 0, int64(len(need)))
	log.StartProgressBar(pb)

	for i := range need {
		err := ex.extractDB(pack, &need[i])
		if err != nil {
			return err
		}
		pb.SetInt(int64(i))
	}

	pb.Complete()

	if err := pack.ZipToFile(filepath.Join(ex.Dir, DBPackName)); err != nil {
		ex.removePack(tempPackDir)
		return err
	}

	return ex.removePack(tempPackDir)
}

func (ex *Extractor) extractDB(pack *datapack.Pack, exDB *ExDb) error {
	prefix := "DBFilesClient\\"
	suffix := ".dbc"

	path := prefix + exDB.Name + suffix

	data, err := ex.Source.ReadFile(path)
	if err != nil {
		return err
	}

	db := dbc.NewDB(ex.Source.Build())
	table, err := db.Open(exDB.Name, bytes.NewReader(data))
	if err != nil {
		return err
	}

	numRecords := int(table.Header.RecordCount)

	file, err := pack.WriteFile(fmt.Sprintf("DB/Ent_%s.txt", exDB.Name))
	if err != nil {
		return err
	}

	encoder := text.NewEncoder(file)

	zero := reflect.New(exDB.Type)
	cursor := reflect.New(exDB.Type)

	progress := log.NewIntProgressBar(
		fmt.Sprintf("Loading %s", exDB.Name),
		0, int64(numRecords),
	)

	log.StartProgressBar(progress)

	for i := 0; i < numRecords; i++ {
		// set cursor to zero value
		cursor.Elem().Set(zero.Elem())

		// read dbc file into struct
		if err := table.Index(i, cursor.Interface()); err != nil {
			return err
		}

		// encode struct to text file
		if err := encoder.Encode(cursor.Interface()); err != nil {
			return err
		}

		progress.SetInt(int64(i))
	}

	progress.Complete()

	return file.Close()
}
