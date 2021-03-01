package tasker

import (
	"sync"
)

type worker struct {
	mx      sync.RWMutex
	id      int
	tasks   <-chan Task
	running bool
}

func (w *worker) run() {
	defer w.setRuning(false)
	w.setRuning(true)

	for t := range w.tasks {
		if err := t.Execute(); err != nil {
			t.HandleError(err)
		}
	}
}

func (w *worker) isDown() bool {
	w.mx.RLock()
	defer w.mx.RUnlock()
	return !w.running
}

func (w *worker) setRuning(s bool) {
	w.mx.Lock()
	defer w.mx.Unlock()
	w.running = s
}
