package mqttcodec

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDisconnectPacket(t *testing.T) {
	a := assert.New(t)
	packet := NewControlPacket(DISCONNECT).(*DisconnectPacket)
	a.Equal(DISCONNECT, packet.MessageType)
	a.Equal(MqttMessageTypeNames[packet.MessageType], packet.Name())
	t.Log(packet)
}

func TestDisconnectPacketCodec(t *testing.T) {
	tests := []struct {
		name       string
		encodedHex string
		packet     *DisconnectPacket
	}{
		{
			name:       "disconnect",
			encodedHex: "e000",
			packet: &DisconnectPacket{
				FixedHeader: FixedHeader{
					MessageType:     DISCONNECT,
					RemainingLength: 0,
				},
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
			packet := decoded.(*DisconnectPacket)
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
