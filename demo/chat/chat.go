package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetAnswer(client *openai.Client, question string) string {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a snarky and sarcastic chatbot",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	}
	resp, err := client.CreateChatCompletion(
		context.TODO(),
		req,
	)
	if err != nil {
		log.Println("CreateChatCompletion failed: ", err)
	}

	//fmt.Printf("%v\n", resp)
	return resp.Choices[0].Message.Content
}

func main() {
	client := openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY"))

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nWhat? >")
	for scanner.Scan() {
		question := scanner.Text()
		answer := GetAnswer(client, question)
		fmt.Println(answer)
		fmt.Print("\nWhat? >")
	}
}
