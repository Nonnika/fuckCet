package aianalyse

// 这个文件是用来分析单词的

import (
	"context"
	"fmt"
	"fuckCet/backend/source"

	"github.com/sashabaranov/go-openai"
)

type Agent struct {
	Model  string              // 模型名称，例如 "gpt-3.5-turbo"
	config openai.ClientConfig // 这里存储了模型的配置，包括API密钥和基础URL
}

func NewAgent(model, key, baseURL string) *Agent {
	config := openai.DefaultConfig(key)
	config.BaseURL = baseURL

	return &Agent{
		Model:  model,
		config: config,
	}
}

// TODO: 这里需要实现一个方法，来分析单词的意思，输入是单词，输出是单词的意思
// prompt 我已经在qwen测试了

func (a *Agent) AnalyseWord(word string) (string, error) {
	// 这里我们需要构造一个prompt，来让模型分析单词的意思

	prompt := fmt.Sprintf(source.WordAnalyseSystemPrompt, word)

	client := openai.NewClientWithConfig(a.config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: a.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser, // 直接输出不需要系统提示
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil

}
