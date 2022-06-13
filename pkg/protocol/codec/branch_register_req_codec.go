package codec

import (
	"github.com/fagongzi/goetty"
)

import (
	model2 "github.com/seata/seata-go/pkg/protocol/branch"
	"github.com/seata/seata-go/pkg/protocol/message"
)

func init() {
	GetCodecManager().RegisterCodec(CodeTypeSeata, &BranchRegisterRequestCodec{})
}

type BranchRegisterRequestCodec struct {
}

func (g *BranchRegisterRequestCodec) Decode(in []byte) interface{} {
	buf := goetty.NewByteBuf(len(in))
	buf.Write(in)
	msg := message.BranchRegisterRequest{}

	length := ReadUInt16(buf)
	if length > 0 {
		bytes := make([]byte, length)
		msg.Xid = string(Read(buf, bytes))
	}

	msg.BranchType = model2.BranchType(ReadByte(buf))

	length = ReadUInt16(buf)
	if length > 0 {
		bytes := make([]byte, length)
		msg.ResourceId = string(Read(buf, bytes))
	}

	length32 := ReadUInt32(buf)
	if length > 0 {
		bytes := make([]byte, length32)
		msg.LockKey = string(Read(buf, bytes))
	}

	length32 = ReadUInt32(buf)
	if length > 0 {
		bytes := make([]byte, length32)
		msg.ApplicationData = Read(buf, bytes)
	}

	return msg
}

func (c *BranchRegisterRequestCodec) Encode(in interface{}) []byte {
	buf := goetty.NewByteBuf(0)
	req, _ := in.(message.BranchRegisterRequest)

	Write16String(req.Xid, buf)
	buf.WriteByte(byte(req.BranchType))
	Write16String(req.ResourceId, buf)
	Write32String(req.LockKey, buf)
	Write32String(string(req.ApplicationData), buf)

	return buf.RawBuf()
}

func (g *BranchRegisterRequestCodec) GetMessageType() message.MessageType {
	return message.MessageType_BranchRegister
}