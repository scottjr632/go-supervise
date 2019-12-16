package entities

// Worker represents a worker that has subscribed
type Worker struct {
	WorkerID         string `json:"workerId"`
	Name             string `json:"name"`
	CheckUpURI       string `json:"checkUpUri"`
	ExpectedResponse string `json:"expectedRespone"`
}

// WorkerBuilder a builder for worker
func WorkerBuilder() *Worker {
	return &Worker{}
}

// SetWorkerID ...
func (w *Worker) SetWorkerID(workerID string) *Worker {
	w.WorkerID = workerID
	return w
}

// SetName ...
func (w *Worker) SetName(name string) *Worker {
	w.Name = name
	return w
}

// SetCheckUpURI ...
func (w *Worker) SetCheckUpURI(checkUpURI string) *Worker {
	w.CheckUpURI = checkUpURI
	return w
}

// SetExpectedResponse ...
func (w *Worker) SetExpectedResponse(expectedResponse string) *Worker {
	w.ExpectedResponse = expectedResponse
	return w
}
