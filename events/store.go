package events

const initialCapacity = 100

// StoreInterface store interface
type StoreInterface interface {
	Append(event Event) error
	GetPatientEvents(patientID string) ([]Event, error)
}

// ArrayStore array store
type ArrayStore struct {
	events map[string][]Event
}

// NewArrayStore create array store
func NewArrayStore() StoreInterface {
	return &ArrayStore{
		make(map[string][]Event),
	}
}

// Append append event
func (store *ArrayStore) Append(event Event) error {
	patientEvents, ok := store.events[*event.PatientID]
	if !ok {
		patientEvents = make([]Event, 0, initialCapacity)
	}
	store.events[*event.PatientID] = append(patientEvents, event)
	return nil
}

// GetPatientEvents get all events for patient id
func (store *ArrayStore) GetPatientEvents(patientID string) ([]Event, error) {
	if patientEvents, ok := store.events[patientID]; ok {
		return patientEvents, nil
	}
	return make([]Event, 0), nil
}
