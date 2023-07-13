package social

import (
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type Add struct {
	Type packet.WorldType
	Name string
}

func (add *Add) Decode(build vsn.Build, in *packet.WorldPacket) (err error) {
	add.Type = in.Type
	add.Name = in.ReadCString()
	return
}

func (add *Add) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	out.Type = add.Type
	out.WriteCString(add.Name)
	return
}
