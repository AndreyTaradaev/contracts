package log

import (
	//"reflect"
	"testing"
	//"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{struct{ name string }{"1"}} // TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); got == nil {
				t.Errorf("New() = %v", got)
			} else {
				defer Close()
			}
		})
	}
}

func TestInitConfig(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		logger := New()
		defer Close()
		InitConfig("", true)
		logger.Info("test debug enable")
	})
	t.Run("1", func(t *testing.T) {
		logger := New()
		defer Close()
		InitConfig("./log.log", false)
		logger.Info("test")
	})

}
