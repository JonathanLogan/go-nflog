//+build linux

package nflog

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/sys/unix"
)

func pUint8(v uint8) *uint8 {
	return &v
}

func pUint16(v uint16) *uint16 {
	return &v
}

func pUint32(v uint32) *uint32 {
	return &v
}

func pString(v string) *string {
	return &v
}

func pBytes(v []byte) *[]byte {
	return &v
}

func TestExtractAttributes(t *testing.T) {
	tests := map[string]struct {
		data []byte
		a    Attribute
	}{
		"SimplePing": {
			data: []byte{0x02, 0x00, 0x00, 0x64, 0x08, 0x00, 0x01, 0x00, 0x08, 0x00, 0x03, 0x00, 0x05, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x05, 0x00, 0x00, 0x00, 0x00, 0x03, 0x08, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x03, 0xe8, 0x08, 0x00, 0x0e, 0x00, 0x00, 0x00, 0x03, 0xe8, 0x58, 0x00, 0x09, 0x00, 0x45, 0x00, 0x00, 0x54, 0x3d, 0x98, 0x40, 0x00, 0x40, 0x01, 0xf0, 0x52, 0x0a, 0x00, 0x00, 0xbd, 0x01, 0x01, 0x01, 0x01, 0x08, 0x00, 0xfe, 0x4b, 0x46, 0xd2, 0x00, 0x02, 0x4e, 0x01, 0x85, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x1a, 0xb0, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37},
			a: Attribute{Hook: pUint8(0), Prefix: pString(""), HwProtocol: pUint16(unix.ETH_P_IP), UID: pUint32(0x03e8), GID: pUint32(0x03e8), OutDev: pUint32(0x03),
				Payload: pBytes([]byte{0x45, 0x00, 0x00, 0x54, 0x3d, 0x98, 0x40, 0x00, 0x40, 0x01, 0xf0, 0x52, 0x0a, 0x00, 0x00, 0xbd, 0x01, 0x01, 0x01, 0x01, 0x08, 0x00, 0xfe, 0x4b, 0x46, 0xd2, 0x00, 0x02, 0x4e, 0x01, 0x85, 0x5b, 0x00, 0x00, 0x00, 0x00, 0x1a, 0xb0, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37})},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			a, err := extractAttributes(nil, tc.data)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if diff := cmp.Diff(tc.a, a); diff != "" {
				t.Fatalf("unexpected number of request messages (-want +got):\n%s", diff)
			}
		})

	}

}
