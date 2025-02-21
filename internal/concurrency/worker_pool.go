package concurrency

import (
	"customer-api/internal/domain"
	"customer-api/internal/ports"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Customer domain.Customer
}

type JobResult struct {
	Success bool
	Message string
}

type WorkerPool struct {
	jobs    chan Job
	results chan JobResult
	wg      sync.WaitGroup
	service ports.CustomerService
}

func NewWorkerPool(workerCount int, service ports.CustomerService) *WorkerPool {
	pool := &WorkerPool{
		jobs:    make(chan Job, 10),
		results: make(chan JobResult, 10),
		service: service,
	}

	for i := 0; i < workerCount; i++ {
		go pool.worker(i)
	}

	return pool
}

func (p *WorkerPool) worker(workerID int) {
	for job := range p.jobs {
		fmt.Printf("Worker %d procesando cliente: %s\n", workerID, job.Customer.Name)

		responseChan := make(chan string, 2)
		errorChan := make(chan error, 2)
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			err := p.service.ValidateCustomer(job.Customer.ID)
			if err != nil {
				errorChan <- err
			} else {
				responseChan <- fmt.Sprintf("✅ Cliente %s validado correctamente", job.Customer.Name)
			}
		}()

		go func() {
			defer wg.Done()
			err := p.service.CreateCustomer(job.Customer)
			if err != nil {
				errorChan <- err
			} else {
				responseChan <- fmt.Sprintf("💾 Cliente %s guardado en la BD", job.Customer.Name)
			}
		}()

		wg.Wait()
		close(responseChan)
		close(errorChan)

		var finalResult string
		var hasError bool
		for err := range errorChan {
			hasError = true
			finalResult += fmt.Sprintf("❌ Error: %v | ", err)
		}
		for res := range responseChan {
			finalResult += res + " | "
		}

		time.Sleep(1 * time.Second)

		p.results <- JobResult{
			Success: !hasError,
			Message: finalResult,
		}
		p.wg.Done()
	}
}

func (p *WorkerPool) AddJob(job Job) {
	p.wg.Add(1)
	p.jobs <- job
}

func (p *WorkerPool) Results() <-chan JobResult {
	return p.results
}

func (p *WorkerPool) Close() {
	p.wg.Wait()
	close(p.results)
	close(p.jobs)
}
