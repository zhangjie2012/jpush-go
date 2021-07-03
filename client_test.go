package jpush

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testRegistrationId = "your_rid"
	key                = "your_key"
	secret             = "your_secret"
)

func TestMain(m *testing.M) {
	Init(key, secret)
	os.Exit(m.Run())
}

func TestPushIOS(t *testing.T) {
	b := NewBasicPushBody("ios", testRegistrationId, false, "test push")
	resp, err := Push(b)
	require.Nil(t, err)
	t.Log(resp)
}
