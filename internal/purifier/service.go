package purifier

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

const model = "deepseek-r1"

func Execute(ctx context.Context, text string) (string, error) {
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal(err)
	}
	completion, err := llm.Call(ctx, text,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		return "", errors.Join(err, errors.New("error creating purifying input text"))
	}

	return completion, nil
}
