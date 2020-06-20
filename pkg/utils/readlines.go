package utils

import (
	"bufio"
	"io"
)

// ReadLines .
func ReadLines(r io.Reader) (<-chan string, error) {
	s := bufio.NewScanner(r)
	if err := s.Err(); err != nil {
		return nil, err
	}
	ch := make(chan string)
	go func() {
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch, nil
}
