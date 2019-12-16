package entities

// HealthStatus represents the health of a worker
type HealthStatus int

// Status health types
const (
	StatusHealthy = iota
	StatusCloudy  = iota
	StatusStormy  = iota
)

var statusMessage = map[HealthStatus]string{
	StatusHealthy: "Worker is healthy",
	StatusCloudy:  "Worker is cloudy",
	StatusStormy:  "Worker is stormy",
}

// breakpoints for health status
const (
	HealthBreakpointHealthy = .10
	HealthBreakpointCloudy  = 0.50
	HealthBreakpointStormy  = 1
)

type Health struct {
	Worker *Worker
	status HealthStatus
}

func getHealthyCheckups(checkups []*CheckUp) []*CheckUp {
	var healthyCheckUps []*CheckUp
	for _, checkup := range checkups {
		if checkup.ResponseCode == "200" {
			healthyCheckUps = append(healthyCheckUps, checkup)
		}
	}
	return healthyCheckUps
}

func (h *Health) SetHealthStatus(checkups []*CheckUp) {
	healthyCheckUps := getHealthyCheckups(checkups)
	healthPercentage := float32(len(healthyCheckUps)) / float32(len(checkups))
	switch true {
	case healthPercentage >= HealthBreakpointHealthy:
		h.status = StatusHealthy
	case healthPercentage >= HealthBreakpointCloudy:
		h.status = StatusCloudy
	case healthPercentage >= HealthBreakpointStormy:
		h.status = StatusStormy
	default:
		h.status = StatusHealthy
	}
}

func (h *Health) Status() string {
	if h.status == 0 {
		h.status = StatusHealthy
	}
	return statusMessage[h.status]
}
