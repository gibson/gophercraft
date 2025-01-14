package dbc

import (
	"fmt"

	"github.com/gibson/gophercraft/format/dbc/dbd"
	"github.com/gibson/gophercraft/format/dbc/dbdefs"
	"github.com/gibson/gophercraft/vsn"
)

type DB struct {
	Build vsn.Build
	Table []*Table
}

func NewDB(v vsn.Build) *DB {
	return &DB{
		Build: v,
	}
}

func (db *DB) lookupDef(name string) (*dbd.Definition, error) {
	return dbdefs.Lookup(name)
}

func (db *DB) detectLayout(table *Table) error {
	def, err := db.lookupDef(table.Name)
	if err != nil {
		return err
	}

	table.Definition = def

	for i := range def.Layouts {
		layout := &def.Layouts[i]
		// Search for exact builds
		for _, exact := range layout.VerifiedBuilds {
			if exact == db.Build {
				table.Layout = layout
				return nil
			}
		}
		// Failing that, see if there is a match within build ranges
		for _, rng := range layout.BuildRanges {
			if rng.Contains(db.Build) {
				table.Layout = layout
				return nil
			}
		}
	}

	return fmt.Errorf("dbc: table found, but could not find layouts matching build %s", db.Build)
}
