package pkg

import (
	"sync"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/magaldima/bizday/holiday"
)

type holidayRegistry struct {
	source   plugin.ClientProtocol
	mu       sync.Mutex
	holidays map[string]holiday.Holiday
}

func (s *server) getHoliday(name string) (holiday.Holiday, error) {
	s.holidayRegistry.mu.Lock()
	defer s.holidayRegistry.mu.Unlock()
	if h, ok := s.holidayRegistry.holidays[name]; ok {
		return h, nil
	}
	raw, err := s.holidayRegistry.source.Dispense(name)
	if err != nil {
		return nil, err
	}
	h := raw.(holiday.Holiday)
	// save it into the registry
	s.holidayRegistry.holidays[name] = h
	return h, nil
}
