package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

func GetAnswer(client *openai.Client, question string) string {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Use the following JSON data to answer the question below",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.TODO(), req)
	if err != nil {
		fmt.Println("CreateChatCompletion failed: ", err)
	}
	return resp.Choices[0].Message.Content
}

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY"))

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Search> ")
	for scanner.Scan() {
		question := scanner.Text()
		answer := GetAnswer(client, question)
		fmt.Println(answer)
		fmt.Print("Search> ")
	}
}
