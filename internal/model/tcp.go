package model

import (
	"CrabGo/internal/consts"
	"errors"
	"fmt"
	"strconv"

	"github.com/gogf/gf/v2/encoding/gbinary"
	"github.com/gogf/gf/v2/encoding/gjson"
)

// Packet /*
type Packet struct {
	Raw      []byte             `description:"原始报文"`
	Header   uint32             `description:"请求头"`
	SerialNo int32              `description:"请求报文序列号"`
	Type     consts.MessageType `description:"报文类型"`
	Length   int                `description:"数据包长度"`
	Data     any                `description:"报文内容"`
	Code     uint32             `description:"报文校验码"`
	Footer   uint32             `description:"报文尾"`

	DataBuffer []byte
}

// NewPacket 创建Packet对象
// raw 请求包源数据字节数组
// 返回值:指向新Packet的指针
func NewPacket(raw []byte) *Packet {
	return &Packet{
		Raw: raw,
	}
}

// Deserialize 解析报文到对象中
func (p *Packet) Deserialize() error {

	err := p.IsValid()
	if err != nil {
		return err
	}

	err = gbinary.Decode(p.Raw, &p.Header, &p.SerialNo, &p.Type, &p.Length)

	if err != nil {
		return errors.New("报文解析失败")
	}

	index := 16
	var dataBytes = p.Raw[index:p.Length]
	p.DataBuffer = dataBytes

	dataJson, err := gjson.EncodeString(dataBytes)

	fmt.Print(dataJson)

	index += p.Length

	dataBytes = p.Raw[index:]

	err = gbinary.Decode(dataBytes, &p.Code, &p.Footer)

	if err != nil {
		return errors.New("解析Code失败")
	}

	return nil
}

// CreatePacket 创建发送包,报文
func CreatePacket(pType consts.MessageType, data any) (packet *Packet, err error) {

	dataBuffer, err := gjson.Encode(data)

	if err != nil {
		return nil, err
	}

	packet.Header = consts.Header
	packet.SerialNo = 1
	packet.Type = pType
	packet.Length = len(dataBuffer)
	packet.Data = data
	packet.Code = 1
	packet.Footer = consts.Footer

	raw := gbinary.Encode(packet.Header, packet.SerialNo, packet.Type, packet.Length, dataBuffer, packet.Code, packet.Footer)

	packet.Raw = raw

	return packet, nil
}

func (p *Packet) IsValid() error {

	if p.Raw == nil {
		return errors.New("原始报文为空")
	}

	if len(p.Raw) < consts.MinLength {
		return errors.New("报文长度不足" + strconv.Itoa(consts.MinLength))
	}

	return nil
}
