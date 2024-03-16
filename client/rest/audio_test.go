package rest_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/iwechat/test/tools"
)

func TestCreateSpeech(t *testing.T) {
	req := audio.NewCreateSpeechRequest()
	req.Model = "tts-1-hd"
	req.Input = "你好, 我的名字是韩梅梅"
	req.Voice = "alloy"
	req.FilePath = "audio"
	req.FileName = "voice.mp3"
	speechResp, err := client.CreateSpeech(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(speechResp))
}

func TestCreateTranscription(t *testing.T) {
	req := audio.NewCreateTranscriptionRequest()
	req.Model = "whisper-1"
	req.FilePath = "audio"
	req.FileName = "voice.mp3"
	req.Language = "zh"
	req.ResponseFormat = "json"
	transcriptionResp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(transcriptionResp))
}

func TestCreateTranslation(t *testing.T) {
	req := audio.NewCreateTranslationRequest()
	req.Model = "whisper-1"
	req.FileName = "voice.mp3"
	req.FilePath = "audio"
	translationResp, err := client.CreateTranslation(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(translationResp))
}
