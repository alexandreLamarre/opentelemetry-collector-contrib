package prometheusremotewriteexporter

import (
	"encoding/binary"
	"errors"

	"github.com/gogo/protobuf/proto"
	"github.com/prometheus/prometheus/prompb"
)

// ! TenantIds are expected to have length < math.MaxUint32, which is reasonable
// because otherwise we have bigger problems
type TenantWritePb struct {
	TenantId string
	WriteReq *prompb.WriteRequest
}

func (t *TenantWritePb) Marshal() ([]byte, error) {
	protoBytes, err := proto.Marshal(t.WriteReq)
	if err != nil {
		return nil, err
	}
	nameBytes := []byte(t.TenantId)
	buf := make([]byte, 4+len(nameBytes)+len(protoBytes))
	binary.LittleEndian.PutUint32(buf[:4], uint32(len(nameBytes)))
	copy(buf[4:], nameBytes)
	copy(buf[4+len(nameBytes):], protoBytes)
	return buf, nil
}

func (t *TenantWritePb) Unmarshal(data []byte) error {
	if len(data) < 4 {
		return errors.New("invalid data")
	}
	nameLen := int(binary.LittleEndian.Uint32(data[:4]))
	if len(data) < 4+nameLen {
		return errors.New("invalid data")
	}
	t.TenantId = string(data[4 : 4+nameLen])
	protoBytes := data[4+nameLen:]
	pb := new(prompb.WriteRequest)
	err := proto.Unmarshal(protoBytes, pb)
	if err != nil {
		return err
	}
	t.WriteReq = pb
	return nil
}
