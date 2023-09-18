package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompression(t *testing.T) {

	compressed, err := Compress(uncompressed)
	require.NoError(t, err)

	decompressed, err := Decompress(compressed)
	require.NoError(t, err)

	assert.Equal(t, uncompressed, decompressed)

	f, err := os.Create("zsrdpayload.txt")
	assert.NoError(t, err)

	_, err = f.Write(compressed)
	assert.NoError(t, err)

	f.Close()

}

var uncompressed = []byte(`[
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "start app"
	},
	{
	  "client_time": "2020-12-01 23:59:01",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 2,
	  "event": "app_load",
	  "param_int": 10,
	  "param_str": "loading some data"
	},
	{
	  "client_time": "2020-12-01 23:59:04",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 3,
	  "event": "app_get",
	  "param_int": 0,
	  "param_str": "requested data from https://not.hehe?cat=1"
	},
	{
	  "client_time": "2020-12-02 00:00:01",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 4,
	  "event": "app_killpid",
	  "param_int": 6686,
	  "param_str": "kill long playing job"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 5,
	  "event": "app_close",
	  "param_int": 1,
	  "param_str": "close application"
	},
	{
	  "client_time": "2020-12-02 00:00:05",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 6,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "start app"
	},
	{
	  "client_time": "2020-12-02 00:00:06",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 7,
	  "event": "app_get",
	  "param_int": 80,
	  "param_str": "get data from resource"
	},
	{
	  "client_time": "2020-12-02 00:00:10",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 8,
	  "event": "app_post",
	  "param_int": 0,
	  "param_str": "sent some data to some host"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	},
	{
	  "client_time": "2020-12-02 00:00:03",
	  "device_id": "0287D9AA-4ADF-4B37-A60F-3E9E645C821E",
	  "device_os": "IOS 13.5.1",
	  "session": "ybuRi8mAUypxjbxQ",
	  "sequence": 1,
	  "event": "app_start",
	  "param_int": 0,
	  "param_str": "some text"
	}
  ]`)
