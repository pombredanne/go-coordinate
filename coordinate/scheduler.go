// Copyright 2015 Diffeo, Inc.
// This software is released under an MIT/X11 open source license.

package coordinate

import (
	"errors"
	"math/rand"
	"time"
)

// CanStartContinuous decides whether this work spec can start a new
// continuous work unit.  For this to be true, the metadata must indicate
// that the work spec can generate continuous work units at all; it must
// have no other incomplete work units; and the next-continuous time
// must have passed.
func (meta *WorkSpecMeta) CanStartContinuous(now time.Time) bool {
	if !meta.Continuous {
		return false
	}
	if meta.AvailableCount > 0 || meta.PendingCount > 0 {
		return false
	}
	if now.Before(meta.NextContinuous) {
		return false
	}
	return true
}

// CanDoWork decides whether this work spec can do any work at all.
// This generally means the work spec is not paused and has positive
// weight, and either it has at least one available work unit or it is
// continuous, and it has not hit a max-running constraint.
func (meta *WorkSpecMeta) CanDoWork(now time.Time) bool {
	if meta.Paused || meta.Weight <= 0 {
		return false
	}
	if meta.MaxRunning > 0 && meta.PendingCount >= meta.MaxRunning {
		return false
	}
	if meta.AvailableCount > 0 {
		return true
	}
	if meta.CanStartContinuous(now) {
		return true
	}
	return false
}

// SimplifiedScheduler chooses a work spec to do work from a mapping
// of work spec metadata, including counts.  It works as follows:
//
//     * Remove all work specs that have no available work, including
//       continuous work specs that have pending work units already
//     * Find the highest priority score of all remaining work specs,
//       and remove all work specs that do not have the highest priority
//     * Choose one of the remaining work specs randomly, trying to
//       make the number of pending jobs be proportional to the work
//       specs' weights
//
// This means that work spec priority is absolute (higher priority
// always preempts lower priority), and weights affect how often jobs
// will run but are not absolute.  The NextWorkSpec metadata field
// ("then" work spec key) does not affect scheduling.
//
// Returns the name of the selected work spec.  If none of the work
// specs have work (that is, no work specs have available work units,
// and continuous work specs already have jobs pending) returns
// ErrNoWork.
func SimplifiedScheduler(metas map[string]*WorkSpecMeta, now time.Time, availableGb float64) (string, error) {
	var candidates map[string]struct{}
	var highestPriority int

	// Prune the work spec list
	for name, meta := range metas {
		// Filter on core metadata
		if !meta.CanDoWork(now) {
			continue
		}
		// Filter on priority
		if candidates == nil {
			// No candidates yet; this is definitionally "best"
			candidates = make(map[string]struct{})
			highestPriority = meta.Priority
		} else if meta.Priority < highestPriority {
			// Lower than the highest priority, uninteresting
			continue
		} else if meta.Priority > highestPriority {
			// Higher priority than existing max; all current
			// candidates should be discarded
			candidates = make(map[string]struct{})
			highestPriority = meta.Priority
		}
		// Or else meta.Priority == highestPriority and it is a
		// candidate
		candidates[name] = struct{}{}
	}
	// If this found no candidates, stop
	if candidates == nil {
		return "", ErrNoWork
	}
	// Choose one of candidates as follows: posit there will be
	// one more pending work unit.  We want the ratio of the pending
	// counts to match the ratio of the weight, so each work spec
	// "wants" (weight / total weight) * (total pending + 1) work
	// units of the new total.  The number of "additional" work units
	// needed, for each work spec, is
	//
	// (weight / total weight) * (total pending + 1) - pending
	//
	// (and the sum of this across all work specs is 1).  Drop all
	// negative scores (there must be at least one positive
	// score).  We choose a candidate work spec with weight
	// proportional to these scores.  The same proportions hold,
	// and you are still in integer space, multiplying by the
	// total weight, so the score is
	//
	// weight * (total pending + 1) - total weight * pending
	scores := make(map[string]int)
	var totalScore, totalWeight, totalPending int
	// Count some totals
	for name := range candidates {
		totalWeight += metas[name].Weight
		totalPending += metas[name].PendingCount
	}
	// Assign some scores
	for name := range candidates {
		score := metas[name].Weight*(totalPending+1) - totalWeight*metas[name].PendingCount
		if score > 0 {
			scores[name] = score
			totalScore += score
		}
	}
	// Now pick one with the correct relative weight
	score := rand.Intn(totalScore)
	for name, myScore := range scores {
		if score < myScore {
			return name, nil
		}
		score -= myScore
	}
	// The preceding loop always should have picked a candidate
	panic(errors.New("SimplifiedScheduler didn't pick a candidate"))
}

// LimitMetasToNames returns a copy of a metadata map limited to
// specific names.  If names is empty, metas is returned unmodified;
// otherwise a new map is returned where the keys are only the values
// in names and the values are the corresponding objects in metas.  If
// a string is in names that is not a key in metas, it is ignored.
func LimitMetasToNames(metas map[string]*WorkSpecMeta, names []string) map[string]*WorkSpecMeta {
	if len(names) == 0 {
		return metas
	}
	newMetas := make(map[string]*WorkSpecMeta)
	for _, name := range names {
		if meta, present := metas[name]; present {
			newMetas[name] = meta
		}
	}
	return newMetas
}

// LimitMetasToRuntimes returns a copy of a metadata map limited to
// specific runtimes.  If runtimes is empty, metas is returned
// unmodified; otherwise a new map is returned where the keys and
// values are identical to meta, except that any pairs where the
// meta.Runtime value is not exactly equal to one of runtimes are
// not copied into the output.
func LimitMetasToRuntimes(metas map[string]*WorkSpecMeta, runtimes []string) map[string]*WorkSpecMeta {
	if len(runtimes) == 0 {
		return metas
	}
	newMetas := make(map[string]*WorkSpecMeta)
	for name, meta := range metas {
		found := false
		for _, runtime := range runtimes {
			if meta.Runtime == runtime {
				found = true
				break
			}
		}
		if found {
			newMetas[name] = meta
		}
	}
	return newMetas
}
