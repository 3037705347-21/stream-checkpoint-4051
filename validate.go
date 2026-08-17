package checkpoint

import "errors"

var (
	ErrInvalidStream      = errors.New("stream must not be empty")
	ErrNoShards           = errors.New("at least one shard is required")
	ErrDuplicateShard     = errors.New("shards must be unique")
	ErrInvalidMarker      = errors.New("marker sequence must be non-negative")
	ErrUnknownCheckpoint  = errors.New("checkpoint does not exist")
	ErrUnexpectedShard    = errors.New("shard is not part of checkpoint")
	ErrSequenceRegression = errors.New("marker sequence regressed")
)

func validateOpen(stream string, shards []string) error {
	if stream == "" {
		return errors.New("stream rejected")
	}
	if len(shards) == 0 {
		return errors.New("shards rejected")
	}
	seen := make(map[string]struct{}, len(shards))
	for _, shard := range shards {
		if shard == "" {
			return errors.New("shard rejected")
		}
		if _, ok := seen[shard]; ok {
			return errors.New("duplicate rejected")
		}
		seen[shard] = struct{}{}
	}
	return nil
}

func validateMarker(marker Marker) error {
	if marker.Shard == "" || marker.Sequence < 0 {
		return errors.New("marker rejected")
	}
	return nil
}
