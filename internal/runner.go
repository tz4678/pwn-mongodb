package internal

import (
	"bufio"
	"context"
	"fmt"
	"pwn-mongodb/pkg/utils"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/ratelimit"
)

// Run .
func Run(cfg *Config) error {
	r := bufio.NewReader(cfg.Input)
	lines, err := utils.ReadLines(r)
	if err != nil {
		return err
	}
	jobs := make(chan string, cfg.Concurrency)
	go func() {
		defer close(jobs)
		for line := range lines {
			jobs <- line
		}
	}()
	results := make(chan string)
	var wg sync.WaitGroup
	rl := ratelimit.New(cfg.RateLimit)
	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go Worker(jobs, results, &wg, rl, cfg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	w := bufio.NewWriter(cfg.Output)
	for result := range results {
		w.WriteString(fmt.Sprintf("%s\n", result))
		w.Flush()
	}
	return nil
}

// Worker .
func Worker(
	jobs <-chan string,
	results chan<- string,
	wg *sync.WaitGroup,
	rl ratelimit.Limiter,
	cfg *Config,
) {
	defer wg.Done()
	for hostname := range jobs {
		func(h string) {
			rl.Take()
			ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
			defer cancel()
			client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+h))
			if err != nil {
				return
			}
			defer client.Disconnect(ctx)
			_, err = client.ListDatabaseNames(ctx, bson.M{})
			if err == nil {
				results <- hostname
			}
		}(hostname)
	}
}
