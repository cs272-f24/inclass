package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func MakeTool() openai.Tool {
	// describe the function & its inputs
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"instructor": {
				Type:        jsonschema.String,
				Description: "The name of the instructor, e.g. John Smith",
			},
			"semester": {
				Type: jsonschema.String,
				Enum: []string{"Fall 2024", "Spring 2025"},
			},
		},
		Required: []string{"instructor"},
	}
	f := openai.FunctionDefinition{
		Name:        "get_instructor_courses",
		Description: "Get the courses taught by the given instructor",
		Parameters:  params,
	}
	t := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &f,
	}
	return t
}

func main() {
	ctx := context.Background()
	client := openai.NewClient(os.Getenv("OPENAI_PROJECT_KEY"))

	getInstCourses := MakeTool()

	// simulate user asking a question that requires the function
	dialogue := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: "What courses is Greg Benson teaching?"},
	}
	fmt.Printf("Asking OpenAI '%v' and providing it a '%v()' function...\n",
		dialogue[0].Content, getInstCourses.Function.Name)
	resp, err := client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4TurboPreview,
			Messages: dialogue,
			Tools:    []openai.Tool{getInstCourses},
		},
	)
	if err != nil || len(resp.Choices) != 1 {
		fmt.Printf("Completion error: err:%v len(choices):%v\n", err,
			len(resp.Choices))
		return
	}
	msg := resp.Choices[0].Message
	if len(msg.ToolCalls) != 1 {
		fmt.Printf("Completion error: len(toolcalls): %v\n", len(msg.ToolCalls))
		return
	}

	// simulate calling the function & responding to OpenAI
	dialogue = append(dialogue, msg)
	fmt.Printf("OpenAI called us back wanting to invoke our function '%v' with params '%v'\n",
		msg.ToolCalls[0].Function.Name, msg.ToolCalls[0].Function.Arguments)
	dialogue = append(dialogue, openai.ChatCompletionMessage{
		Role:       openai.ChatMessageRoleTool,
		Content:    "CS 315 and CS 315L",
		Name:       msg.ToolCalls[0].Function.Name,
		ToolCallID: msg.ToolCalls[0].ID,
	})
	fmt.Printf("Sending OpenAI our '%v()' function's response and requesting the reply to the original question...\n",
		getInstCourses.Function.Name)
	resp, err = client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4TurboPreview,
			Messages: dialogue,
			Tools:    []openai.Tool{getInstCourses},
		},
	)
	if err != nil || len(resp.Choices) != 1 {
		fmt.Printf("2nd completion error: err:%v len(choices):%v\n", err,
			len(resp.Choices))
		return
	}

	// display OpenAI's response to the original question utilizing our function
	msg = resp.Choices[0].Message
	fmt.Printf("OpenAI answered the original request with: %v\n",
		msg.Content)
}
