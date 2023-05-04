package prometheusremotewriteexporter

import (
	"testing"

	"github.com/prometheus/prometheus/prompb"
	"github.com/stretchr/testify/require"
)

func TestTenantMarshal(t *testing.T) {
	w := &TenantWritePb{
		TenantId: "test",
		WriteReq: &prompb.WriteRequest{},
	}
	data, err := w.Marshal()
	require.NoError(t, err)
	require.NotNil(t, data)
	w2 := &TenantWritePb{}
	err = w2.Unmarshal(data)
	require.NoError(t, err)
	require.Equal(t, w.TenantId, w2.TenantId)
	require.Equal(t, w, w2)
}
