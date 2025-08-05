package jsonproc

import (
	"context"
	"sync"
)

func Run(inpath string, outpath string, noofworkers int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan RawJson)
	results := make(chan string)
	errCh := make(chan error, 1)

}
