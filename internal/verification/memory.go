package verification

import (
	"github.com/pkg/errors"
	"sync"
)

type verification struct {
	sync.RWMutex
	ready bool
	items map[uint8]string
}

type MemoryService interface {
	SetPhoneAndCode(phone string, code uint8) error
	Verify(code uint8) (*string, error)
}

func NewService() MemoryService {
	items := make(map[uint8]string, 0)
	s := &verification{
		items: items,
	}
	s.ready = true

	return s
}

func (s *verification) SetPhoneAndCode(phone string, code uint8) error {
	if !s.ready {
		return errors.New("verification service is not ready")
	}
	s.Lock()
	defer s.Unlock()

	s.items[code] = phone
	return nil
}

func (s *verification) Verify(code uint8) (*string, error) {
	if !s.ready {
		return nil, errors.New("verification service is not ready")
	}
	s.RLock()
	defer s.RUnlock()

	if v, ok := s.items[code]; ok {
		return &v, nil
	}

	return nil, errors.New("there is no phone in layer")
}
