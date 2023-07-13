package ex

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gibson/gophercraft/utils/log"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Dumper struct {
	Ex         *Extractor
	File       *os.File
	Descriptor *descriptorpb.FileDescriptorProto
}

func (ex *Extractor) NewDumper(desc *descriptorpb.FileDescriptorProto) *Dumper {
	d := new(Dumper)
	d.Ex = ex
	d.Descriptor = desc
	return d
}

func (d *Dumper) printf(fomt string, args ...any) {
	fmt.Fprintf(d.File, fomt, args...)
}

func (d *Dumper) Dump() (err error) {
	desc := d.Descriptor
	path := desc.GetName()

	if path == "" {
		err = fmt.Errorf("no name")
		return
	}

	realPath := filepath.Join(d.Ex.Dir, path)

	realDir := filepath.Dir(realPath)

	log.Println("making a real dir", realDir)

	if err = os.MkdirAll(realDir, 0700); err != nil {
		return
	}

	d.File, err = os.OpenFile(realPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0700)
	if err != nil {
		return
	}

	syntax := "proto2"

	if desc.Syntax != nil {
		syntax = desc.GetSyntax()
	}

	d.printf("syntax = \"%s\";\n\n", syntax)

	if desc.Package != nil {
		d.printf("package %s;\n\n", desc.GetPackage())
	}

	d.printf("// DO NOT EDIT: generated by Gophercraft\n\n")

	for i, dep := range desc.Dependency {
		index := int32(i)

		prefix := ""

		for _, idx := range desc.WeakDependency {
			if index == idx {
				prefix = " weak"
				break
			}
		}

		for _, idx := range desc.PublicDependency {
			if index == idx {
				prefix = " public"
				break
			}
		}

		d.printf("import%s \"%s\"\n", prefix, dep)
	}

	if len(desc.Dependency) > 0 {
		d.printf("\n")
	}

	return d.File.Close()
}
