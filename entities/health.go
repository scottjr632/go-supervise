package entities

import "log"

// HealthStatus represents the health of a worker
type HealthStatus int

// Status health types
const (
	statusNotSet  = iota
	StatusHealthy = iota
	StatusCloudy  = iota
	StatusStormy  = iota
	StatusDefault = iota
)

var statusMessage = map[HealthStatus]string{
	StatusHealthy: "Worker is healthy",
	StatusCloudy:  "Worker is cloudy",
	StatusStormy:  "Worker is stormy",

	StatusDefault: "No checkups found",
}

// breakpoints for health status
const (
	HealthBreakpointHealthy = .85
	HealthBreakpointCloudy  = 0.50
	HealthBreakpointStormy  = 0
)

type Health struct {
	Worker *Worker
	status HealthStatus

	Checkups []*CheckUp
}

func getHealthyCheckups(checkups []*CheckUp) []*CheckUp {
	var healthyCheckUps []*CheckUp
	for _, checkup := range checkups {
		if checkup.ResponseCode == "200 OK" {
			healthyCheckUps = append(healthyCheckUps, checkup)
		}
	}
	return healthyCheckUps
}

func (h *Health) SetHealthStatus(checkups []*CheckUp) {
	h.Checkups = checkups
	healthyCheckUps := getHealthyCheckups(checkups)
	log.Println(healthyCheckUps)
	log.Println(checkups)
	log.Println(float32(len(healthyCheckUps)) / float32(len(checkups)))
	healthPercentage := float32(len(healthyCheckUps)) / float32(len(checkups))
	switch true {
	case healthPercentage >= HealthBreakpointHealthy:
		h.status = StatusHealthy
		break
	case healthPercentage >= HealthBreakpointCloudy:
		h.status = StatusCloudy
		break
	case healthPercentage >= HealthBreakpointStormy:
		h.status = StatusStormy
		break
	default:
		h.status = StatusHealthy
	}
}

func (h *Health) Status() string {
	if h.status == statusNotSet {
		h.status = StatusDefault
	}
	return statusMessage[h.status]
}
