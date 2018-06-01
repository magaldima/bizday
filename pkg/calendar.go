package pkg

import (
	"sync"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/magaldima/bizday/calendar"
)

type calendarRegistry struct {
	source    plugin.ClientProtocol
	mu        sync.Mutex
	calendars map[string]calendar.Calendar
}

func (s *server) getCalendar(name string) (calendar.Calendar, error) {
	s.holidayRegistry.mu.Lock()
	defer s.holidayRegistry.mu.Unlock()
	if h, ok := s.calendarRegistry.calendars[name]; ok {
		return h, nil
	}
	raw, err := s.calendarRegistry.source.Dispense(name)
	if err != nil {
		return nil, err
	}
	c := raw.(calendar.Calendar)
	// save it into the registry
	s.calendarRegistry.calendars[name] = c
	return c, nil
}
