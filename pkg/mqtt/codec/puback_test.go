package mqttcodec

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPubackPacket(t *testing.T) {
	a := assert.New(t)
	packet := NewControlPacket(PUBACK).(*PubackPacket)
	a.Equal(PUBACK, packet.MessageType)
	a.Equal(MqttMessageTypeNames[packet.MessageType], packet.Name())
	t.Log(packet)
}

func TestPubackPacketCodec(t *testing.T) {
	tests := []struct {
		name       string
		encodedHex string
		packet     *PubackPacket
	}{
		{
			name:       "message 1",
			encodedHex: "40020001",
			packet: &PubackPacket{
				FixedHeader: FixedHeader{
					MessageType:     PUBACK,
					RemainingLength: 2,
				},
				MessageID: 1,
			},
		},
		{
			name:       "message 1024",
			encodedHex: "40020400",
			packet: &PubackPacket{
				FixedHeader: FixedHeader{
					MessageType:     PUBACK,
					RemainingLength: 2,
				},
				MessageID: 1024,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := assert.New(t)
			t.Log(tc.packet)

			// decode
			encodedBytes, err := hex.DecodeString(tc.encodedHex)
			if err != nil {
				t.Fatal(err)
			}
			r := bytes.NewReader(encodedBytes)
			decoded, err := ReadPacket(r)
			if err != nil {
				t.Fatal(err)
			}
			packet := decoded.(*PubackPacket)
			a.Equal(*tc.packet, *packet)

			// encode
			var output bytes.Buffer
			err = packet.Write(&output)
			if err != nil {
				t.Fatal(err)
			}
			encodedBytes = output.Bytes()
			a.Equal(tc.encodedHex, hex.EncodeToString(encodedBytes))
		})
	}
}
