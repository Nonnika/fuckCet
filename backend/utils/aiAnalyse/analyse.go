package aianalyse

// 这个文件是用来分析单词的，使用了DeepSeek的API来进行分析

import "github.com/sashabaranov/go-openai"

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

//TODO: 这里需要实现一个方法，来分析单词的意思，输入是单词，输出是单词的意思
// prompt 我已经在qwen测试了
