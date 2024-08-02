package utils

import (
	"blog/utils"
	"testing"
)

func TestLogger(t *testing.T) {
	utils.InitLogger("log")
	utils.LogRus.Info("test log")

	utils.LogRus.Error("test error log")

	utils.LogRus.Debug("test debug log")

	utils.LogRus.Warn("test warn log")

}
