package verification

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewService__isNotNil(t *testing.T) {
	srv := NewService()
	require.NotNil(t, srv)
}

func TestVerification_SetPhoneAndCode(t *testing.T) {
	srv := NewService()
	test := struct{
		phone string
		code int
	}{
		phone: "+7(999)8887766",
		code: 1234,
	}
	err := srv.SetPhoneAndCode(test.phone, uint8(test.code))
	require.NoError(t, err)
}

func TestVerification_Verify(t *testing.T) {
	srv := NewService()
	test := struct{
		phone string
		code int
	}{
		phone: "+7(999)8887766",
		code: 1234,
	}
	err := srv.SetPhoneAndCode(test.phone, uint8(test.code))
	require.NoError(t, err)

	result, err :=srv.Verify(uint8(test.code))
	require.NoError(t, err)
	require.Equal(t, test.phone, *result)
}
