package checkpoint

import "testing"

func TestOpeningAnotherCheckpointRetainsEarlierState(t *testing.T) {
	coordinator := NewCoordinator()
	first, err := coordinator.Open("audit", []string{"a"})
	if err != nil {
		t.Fatalf("open first: %v", err)
	}
	if _, err := coordinator.Open("metrics", []string{"b"}); err != nil {
		t.Fatalf("open second: %v", err)
	}
	if _, err := coordinator.Snapshot(first.ID); err != nil {
		t.Fatalf("first checkpoint disappeared: %v", err)
	}
}
