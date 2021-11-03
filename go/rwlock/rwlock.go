package rwlock

import "info/go/semaphore"

type RWLock struct {
	lock      semaphore.Semaphore // binary semaphore: protects RWLock data structure (guard)
	writeLock semaphore.Semaphore // binary semaphore: exclude writers when reading and exclude readers when writing
	readers   int                 // count of readers reading in critical section
}

func New() *RWLock {
	return &RWLock{
		lock:      *semaphore.New(1),
		writeLock: *semaphore.New(1),
	}
}

func (rw *RWLock) AcquireReadLock() {
	rw.lock.Wait()
	rw.readers++
	if rw.readers == 1 {
		// first reader acquires write lock
		rw.writeLock.Wait()
	}
	rw.lock.Post()
}

func (rw *RWLock) ReleaseReadLock() {
	rw.lock.Wait()
	rw.readers--
	if rw.readers == 0 {
		// last reader release write lock
		rw.writeLock.Post()
	}
	rw.lock.Post()
}

func (rw *RWLock) AcquireWriteLock() {
	rw.writeLock.Wait()
}

func (rw *RWLock) ReleaseWriteLock() {
	rw.writeLock.Post()
}
