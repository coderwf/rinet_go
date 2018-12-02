package log

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T){
	LogTo(STDOUT,DEBUG)
	logger := NewPrefixedLogger("test")
	logger.Debug("i love you %s ","fzz")
	time.Sleep(1 * time.Second)
}//
