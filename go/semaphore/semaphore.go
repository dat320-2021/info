package semaphore

type Semaphore struct {
	semChan chan struct{}
}

func New(n int) *Semaphore {
	return &Semaphore{
		semChan: make(chan struct{}, n),
	}
}

func (s *Semaphore) Acquire() {
	s.semChan <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.semChan
}

func (s *Semaphore) Wait() {
	s.semChan <- struct{}{}
}

func (s *Semaphore) Post() {
	<-s.semChan
}
