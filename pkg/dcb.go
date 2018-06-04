package pkg

import (
	"sync"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/magaldima/bizday/dcb"
)

type dcbRegistry struct {
	source plugin.ClientProtocol
	mu     sync.Mutex
	dcbs   map[string]dcb.DayCountBasis
}

func (s *server) getDCB(name string) (dcb.DayCountBasis, error) {
	s.holidayRegistry.mu.Lock()
	defer s.holidayRegistry.mu.Unlock()
	if h, ok := s.dcbRegistry.dcbs[name]; ok {
		return h, nil
	}
	raw, err := s.dcbRegistry.source.Dispense(name)
	if err != nil {
		return nil, err
	}
	c := raw.(dcb.DayCountBasis)
	// save it into the registry
	s.dcbRegistry.dcbs[name] = c
	return c, nil
}
