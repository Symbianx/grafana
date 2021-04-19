// +build integration

package sqlstore

import (
	"encoding/json"
	"testing"

	"github.com/grafana/grafana/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestLiveMessage(t *testing.T) {
	ss := InitTestDB(t)

	getQuery := &models.GetLiveChannelQuery{
		OrgId:   1,
		Channel: "test_channel",
	}
	_, ok, err := ss.GetLiveChannel(getQuery)
	require.NoError(t, err)
	require.False(t, ok)

	saveQuery := &models.SaveLiveChannelDataQuery{
		OrgId:   1,
		Channel: "test_channel",
		Data:    []byte(`{}`),
	}
	err = ss.SaveLiveChannelData(saveQuery)
	require.NoError(t, err)

	msg, ok, err := ss.GetLiveChannel(getQuery)
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, int64(1), msg.OrgId)
	require.Equal(t, "test_channel", msg.Channel)
	require.Equal(t, json.RawMessage(`{}`), msg.Data)
	require.NotZero(t, msg.Created)

	// try saving again, should be replaced.
	saveQuery2 := &models.SaveLiveChannelDataQuery{
		OrgId:   1,
		Channel: "test_channel",
		Data:    []byte(`{"input": "hello"}`),
	}
	err = ss.SaveLiveChannelData(saveQuery2)
	require.NoError(t, err)

	getQuery2 := &models.GetLiveChannelQuery{
		OrgId:   1,
		Channel: "test_channel",
	}
	msg2, ok, err := ss.GetLiveChannel(getQuery2)
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, int64(1), msg2.OrgId)
	require.Equal(t, "test_channel", msg2.Channel)
	require.Equal(t, json.RawMessage(`{"input": "hello"}`), msg2.Data)
	require.NotZero(t, msg2.Created)
}
