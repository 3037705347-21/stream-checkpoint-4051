package checkpoint

import (
	"errors"
	"testing"
)

func TestPublicErrorsAndCheckpointIDsRemainStable(t *testing.T) {
	coordinator := NewCoordinator()
	tests := []struct {
		name string
		run  func() error
		want error
	}{
		{"invalid stream", func() error { _, err := coordinator.Open("", []string{"a"}); return err }, ErrInvalidStream},
		{"no shards", func() error { _, err := coordinator.Open("audit", nil); return err }, ErrNoShards},
		{"unknown checkpoint", func() error { _, err := coordinator.Snapshot("missing"); return err }, ErrUnknownCheckpoint},
		{"unexpected shard", func() error { opened, _ := coordinator.Open("audit", []string{"a"}); _, err := coordinator.Acknowledge(opened.ID, "b", 1); return err }, ErrUnexpectedShard},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.run(); !errors.Is(err, test.want) {
				t.Fatalf("error = %v, want errors.Is(_, %v)", err, test.want)
			}
		})
	}
	opened, err := coordinator.Open("metrics", []string{"a"})
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if opened.ID != "metrics-2" {
		t.Fatalf("checkpoint ID = %q, want metrics-2", opened.ID)
	}
}
