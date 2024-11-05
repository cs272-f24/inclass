package main

import (
	"context"
	"log"
	"os"

	"github.com/amikos-tech/chroma-go/openai"
)

func main() {

	efunc, err := openai.NewOpenAIEmbeddingFunction(os.Getenv("OPENAI_PROJECT_KEY"))
	if err != nil {
		log.Fatalln("NewEmbeddingFunction failed: ", err)
	}

	docs := []string{"who is teaching CS 272 at USF this semester?"}
	embeddings, err := efunc.EmbedDocuments(context.TODO(), docs)
	if err != nil {
		log.Fatalln("EmbedDocuments failed: ", err)
	}
	for _, e := range embeddings {
		log.Println(e)
	}
}
