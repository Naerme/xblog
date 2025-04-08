package ai_service

import (
	"bufio"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type Choice struct {
	Index int `json:"index"`
	Delta struct {
		Content string `json:"content"`
	} `json:"delta"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason interface{} `json:"finish_reason"`
}

type StreamData struct {
	Id                string      `json:"id"`
	Choices           []Choice    `json:"choices"`
	Created           int         `json:"created"`
	Model             string      `json:"model"`
	Object            string      `json:"object"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

func ChatStream(content string) (msgChan chan string, err error) {
	msgChan = make(chan string)
	r := Request{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: "你是一个叫枫枫知道的人工智能助手",
			},
			{
				Role:    "user",
				Content: content,
			},
		},
		Stream: true,
	}
	res, err := BaseRequest(r)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(res.Body)
	// 按行分割
	scanner.Split(bufio.ScanLines)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				continue
			}
			data := text[6:]
			if data == "[DONE]" {
				close(msgChan)
				return
			}
			var item StreamData
			err = json.Unmarshal([]byte(data), &item)
			if err != nil {
				logrus.Errorf("解析失败 %s %s", err, data)
				continue
			}
			if len(item.Choices) > 0 && item.Choices[0].Delta.Content != "" {
				msgChan <- item.Choices[0].Delta.Content
			}

		}
	}()

	return
}
