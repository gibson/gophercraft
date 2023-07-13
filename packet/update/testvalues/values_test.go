package testvalues

import (
	"fmt"
	"testing"

	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/packet/update"

	_ "github.com/gibson/gophercraft/packet/update/descriptorsupport"
)

func TestValues(t *testing.T) {
	vb, err := update.NewValuesBlock(5875, guid.TypeMaskPlayer|guid.TypeMaskUnit|guid.TypeMaskObject)
	if err != nil {
		t.Fatal(err)
	}
	vb.SetUint32("Type", 69)

	fmt.Println("level", vb.Get("Level").ChangeMaskOffset)

	vb.Get("Level").SetUint32(255)

	fmt.Println("slots", vb.Get("InventorySlots").ChangeMaskOffset)

	// var data bytes.Buffer

	// enc, err := update.NewEncoder(5875, &data, 1)
	// if err != nil {
	// 	panic(err)
	// }

	// enc.AddBlock()
}
