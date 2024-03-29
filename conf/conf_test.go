package conf_test

import (
	"testing"

	"github.com/solodba/iwechat/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App.Name)
	t.Log(conf.C().WeChat.RemarkNameList)
	t.Log(conf.C().WeChat.VoiceFilePath)
	t.Log(conf.C().WeChat.FileTuningFilePath)
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().App.Name)
	t.Log(conf.C().WeChat.RemarkNameList)
	t.Log(conf.C().WeChat.VoiceFilePath)
	t.Log(conf.C().WeChat.FileTuningFilePath)
}
