package pkg

import (
	"sync"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/magaldima/bizday/holidays/shared"
)

type holidayCalendarRegistry struct {
	source      plugin.ClientProtocol
	mu          sync.Mutex
	holidayCals map[string]shared.Holiday
}

func (s *server) getHolidayCalendar(name string) (shared.Holiday, error) {
	s.holidayRegistry.mu.Lock()
	defer s.holidayRegistry.mu.Unlock()
	if h, ok := s.holidayRegistry.holidayCals[name]; ok {
		return h, nil
	}
	raw, err := s.holidayRegistry.source.Dispense(name)
	if err != nil {
		return nil, err
	}
	h := raw.(shared.Holiday)
	// save it into the registry
	s.holidayRegistry.holidayCals[name] = h
	return h, nil
}
