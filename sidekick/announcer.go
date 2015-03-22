package sidekick

// Presence interface defines actions for handling a service being removed or
// added by the monitoring service.
type Presence interface {
	// WasAdded will be called once the monitoring service affirms the new service
	// is available.
	WasAdded() error
	// WasRemoved will be called either right before this unit exits, or the
	// monitoring service receives an invalid response.
	WasRemoved() error
}

func AnnounceServiceWasAdded(presenceService Presence) error {
	return presenceService.WasAdded()
}
