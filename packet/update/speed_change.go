package update

import (
	"fmt"

	"github.com/gibson/gophercraft/guid"
	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type SpeedOpcodePair struct {
	Force  packet.WorldType
	Spline packet.WorldType
}

var (
	SpeedCodes = map[SpeedType]SpeedOpcodePair{
		Walk:           {packet.SMSG_FORCE_WALK_SPEED_CHANGE, packet.SMSG_SPLINE_SET_WALK_SPEED},
		Run:            {packet.SMSG_FORCE_RUN_SPEED_CHANGE, packet.SMSG_SPLINE_SET_RUN_SPEED},
		RunBackward:    {packet.SMSG_FORCE_RUN_BACK_SPEED_CHANGE, packet.SMSG_SPLINE_SET_RUN_BACK_SPEED},
		Swim:           {packet.SMSG_FORCE_SWIM_SPEED_CHANGE, packet.SMSG_SPLINE_SET_SWIM_SPEED},
		SwimBackward:   {packet.SMSG_FORCE_SWIM_BACK_SPEED_CHANGE, packet.SMSG_SPLINE_SET_SWIM_BACK_SPEED},
		Turn:           {packet.SMSG_FORCE_TURN_RATE_CHANGE, packet.SMSG_SPLINE_SET_TURN_RATE},
		Flight:         {packet.SMSG_FORCE_FLIGHT_SPEED_CHANGE, packet.SMSG_SPLINE_SET_FLIGHT_SPEED},
		FlightBackward: {packet.SMSG_FORCE_FLIGHT_BACK_SPEED_CHANGE, packet.SMSG_SPLINE_SET_FLIGHT_BACK_SPEED},
	}
)

type ForceSpeedChange struct {
	Type     SpeedType
	ID       guid.GUID
	Counter  uint32
	NewSpeed float32
}

func (fs *ForceSpeedChange) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	code, ok := SpeedCodes[fs.Type]
	if !ok {
		err = fmt.Errorf("%d", fs.Type)
		return
	}

	out.Type = code.Force
	fs.ID.EncodePacked(build, out)
	out.WriteUint32(fs.Counter)

	if fs.Type == Run {
		out.WriteByte(0)
	}

	out.WriteFloat32(fs.NewSpeed)
	return nil
}

type SplineSpeedChange struct {
	Type     SpeedType
	ID       guid.GUID
	NewSpeed float32
}

func (ss *SplineSpeedChange) Encode(build vsn.Build, out *packet.WorldPacket) (err error) {
	code, ok := SpeedCodes[ss.Type]
	if !ok {
		err = fmt.Errorf("%d", ss.Type)
		return
	}

	out.Type = code.Spline
	ss.ID.EncodePacked(build, out)
	out.WriteFloat32(ss.NewSpeed)
	return nil
}
