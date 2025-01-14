package chat

import (
	"fmt"

	"github.com/gibson/gophercraft/packet"
	"github.com/gibson/gophercraft/vsn"
)

type ClientMessage struct {
	Type     MsgType
	Language Language
	To       string
	Body     string
}

func (msg *ClientMessage) Decode(build vsn.Build, in *packet.WorldPacket) error {
	code := in.ReadUint32()

	var err error
	msg.Type, err = ResolveMsgType(build, uint32(code))
	if err != nil {
		return err
	}

	msg.Language = Language(in.ReadUint32())

	switch msg.Type {
	case MsgSay, MsgEmote, MsgYell,
		MsgParty, MsgOfficer, MsgRaid,
		MsgRaidLeader, MsgRaidWarning,
		MsgAFK, MsgDND:
		msg.Body = in.ReadCString()
	case MsgWhisper:
		msg.To = in.ReadCString()
		msg.Body = in.ReadCString()
	case MsgChannel:
		msg.To = in.ReadCString()
		msg.Body = in.ReadCString()
	default:
		return fmt.Errorf("chat: unrecognized type in client message: %v", msg.Type)
	}

	return nil
}

func (msg *ClientMessage) Encode(build vsn.Build, out *packet.WorldPacket) error {
	out.Type = packet.CMSG_MESSAGECHAT

	code, err := ConvertMsgType(build, msg.Type)
	if err != nil {
		return err
	}

	out.WriteUint32(code)
	out.WriteUint32(uint32(msg.Language))

	switch msg.Type {
	case MsgSay, MsgEmote, MsgYell,
		MsgParty, MsgOfficer, MsgRaid,
		MsgRaidLeader, MsgRaidWarning,
		MsgAFK, MsgDND:
		out.WriteCString(msg.Body)
	case MsgWhisper:
		out.WriteCString(msg.To)
		out.WriteCString(msg.Body)
	case MsgChannel:
		out.WriteCString(msg.To)
		out.WriteCString(msg.Body)
	default:
		return fmt.Errorf("chat: unrecognized type in client message: %v", msg.Type)
	}

	return nil
}
