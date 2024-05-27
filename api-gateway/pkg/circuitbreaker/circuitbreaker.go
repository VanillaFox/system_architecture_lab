package circuitbreaker

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type CircuitBreaker struct {
	maxFailures          int
	minSuccess           int
	openDuration         time.Duration
	untilFailureDuration time.Duration
	state                State
	failures             int
	success              int
	mutex                sync.Mutex
}

type State int

const (
	Opened State = iota
	HalfOpened
	Closed
)

func NewCircuitBreaker(maxFailures, minSuccess int, openDuration, untilFailureDuration time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:          maxFailures,
		minSuccess:           minSuccess,
		openDuration:         openDuration,
		untilFailureDuration: untilFailureDuration,
		state:                Closed,
	}
}

func (cb *CircuitBreaker) RoundTrip(req *http.Request) (*http.Response, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	switch cb.state {
	case Opened:
		log.Info("Circuit breaker is Opened")

		return nil, errors.New("circuit breaker is opened")
	case HalfOpened:
		resp, err := http.DefaultTransport.RoundTrip(req)
		if err != nil || resp.StatusCode == http.StatusInternalServerError {
			log.Info("Circuit Breaker is (Half Opened -> Opened)")
			cb.state = Opened
			cb.success = 0
			cb.failures++
			go cb.halfOpenAfterTimeout()

			return nil, err
		} else {
			cb.success++

			if cb.success >= cb.minSuccess {
				log.Info("Circuit Breaker is (Half Opened -> Closed)")
				cb.state = Closed
			} else {
				log.Info("Circuit Breaker is Half Opened")
			}

			cb.failures = 0

			return resp, nil
		}
	case Closed:
		resp, err := http.DefaultTransport.RoundTrip(req)
		if err != nil || resp.StatusCode == http.StatusInternalServerError {
			if cb.failures == 0 {
				go cb.failuresResetTimeout()
			}

			cb.failures++
			if cb.failures >= cb.maxFailures {
				log.Info("Circuit Breaker is (Closed -> Opened)")
				cb.state = Opened
				cb.success = 0

				go cb.halfOpenAfterTimeout()
			} else {
				log.Info("Circuit Breaker is Closed")
			}

			return nil, err
		} else {
			log.Info("Circuit Breaker is Closed")

			return resp, nil
		}
	}

	return nil, nil
}

func (cb *CircuitBreaker) halfOpenAfterTimeout() {
	time.Sleep(cb.openDuration)
	log.Info("Circuit Breaker is (Opened -> Half Opened)")
	cb.state = HalfOpened

	cb.mutex.Lock()
	cb.success = 0
	cb.mutex.Unlock()
}

func (cb *CircuitBreaker) failuresResetTimeout() {
	time.Sleep(cb.untilFailureDuration)
	cb.mutex.Lock()
	cb.failures = 0
	cb.mutex.Unlock()
}
