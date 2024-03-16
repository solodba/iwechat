package rest_test

import (
	"testing"

	"github.com/solodba/ichatgpt/apps/audio"
	"github.com/solodba/iwechat/test/tools"
)

func TestCreateSpeech(t *testing.T) {
	req := audio.NewCreateSpeechRequest()
	req.Model = "tts-1-hd"
	req.Input = "大家好，我叫沃尔夫冈，来自德国。你今天要去哪里?"
	req.Voice = "alloy"
	req.FilePath = "audio"
	req.FileName = "voice.mp3"
	speechResp, err := client.CreateSpeech(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(speechResp))
}

func TestTranscription(t *testing.T) {
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
