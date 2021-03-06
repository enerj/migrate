package stub

import (
	"github.com/orkusinc/migrate"
	"github.com/orkusinc/migrate/source"
	"github.com/orkusinc/migrate/source/stub"
	"testing"

	dt "github.com/orkusinc/migrate/database/testing"
)

func Test(t *testing.T) {
	s := &Stub{}
	d, err := s.Open("")
	if err != nil {
		t.Fatal(err)
	}
	dt.Test(t, d, []byte("/* foobar migration */"))
}

func TestMigrate(t *testing.T) {
	s := &Stub{}
	d, err := s.Open("")
	if err != nil {
		t.Fatal(err)
	}

	stubMigrations := source.NewMigrations()
	stubMigrations.Append(&source.Migration{Version: 1, Direction: source.Up, Identifier: "CREATE 1"})
	stubMigrations.Append(&source.Migration{Version: 1, Direction: source.Down, Identifier: "DROP 1"})
	src := &stub.Stub{}
	srcDrv, err := src.Open("")
	srcDrv.(*stub.Stub).Migrations = stubMigrations
	m, err := migrate.NewWithInstance("stub", srcDrv, "", d)
	if err != nil {
		t.Fatalf("%v", err)
	}

	dt.TestMigrate(t, m, []byte("/* foobar migration */"))
}
