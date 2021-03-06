// Copyright 2015-2016 Diffeo, Inc.
// This software is released under an MIT/X11 open source license.

// Package coordinate defines an abstract API to Coordinate.
//
// In most cases, applications will know of specific implementations of
// this API and will get an implementation of Coordinate or Namespace
// from that implementation.
//
// In general, objects here have a small amount of immutable data
// (a WorkUnit.Name() never changes, for instance) and the accessors
// of these return the value directly.  Accessors to mutable data return
// the value and an error.
package coordinate

import "time"

// Coordinate is the principal interface to the Coordinate system.
// Implementations of this interface provide a specific database backend,
// RPC system, or other way to interact with Coordinate.
type Coordinate interface {
	Summarizable

	// Namespace retrieves a Namespace object for some name.  If
	// no namespace already exists with that name, creates one.
	// Coordinate implementations such as the Python one that do
	// not use namespaces pass an empty string here.
	Namespace(namespace string) (Namespace, error)

	// Namespaces retrieves a map of all known namespaces.
	Namespaces() (map[string]Namespace, error)
}

// Namespace is a single application's state within Coordinate.  A
// namespace has an immutable name, and a collection of work specs.  A
// namespace is tied to a single Coordinate backend.  Most
// applications will only need to interact with a single namespace.
type Namespace interface {
	Summarizable

	// Name returns the name of this namespace.  This may be an
	// empty string.
	Name() string

	// Destroy destroys this namespace and all state associated
	// with it.  There is no recovery from this.  There is no
	// confirmation in the API.  This generally should not be
	// called outside of test code.
	//
	// If other goroutines or processes are using this namespace
	// or an equivalent object, operations on it will likely fail
	// (or, depending on database constraints, this operation may
	// itself fail).  You probably don't want to be in this state.
	Destroy() error

	// SetWorkSpec creates or updates a work spec.  The map may
	// have any string keys and any values, except that it must
	// contain a key "name" with a string value.  You cannot
	// rename an existing work spec, but changing any other keys
	// will change an existing work spec.  On success returns the
	// created (or modified) WorkSpec object.
	SetWorkSpec(workSpec map[string]interface{}) (WorkSpec, error)

	// WorkSpec retrieves a work spec by its name.  If no work
	// spec exists with that name, returns an instance of
	// ErrNoSuchWorkSpec as an error.
	WorkSpec(name string) (WorkSpec, error)

	// DestroyWorkSpec irretreviably destroys a work spec, all
	// work units associated with it, and all attempts to perform
	// those work units.  If workers are currently working on any
	// of these work units, the attempts will also be removed from
	// their active lists, and calls to complete or update those
	// attempts will result in errors.  If the named work spec
	// does not exist, returns an instance of ErrNoSuchWorkSpec.
	DestroyWorkSpec(name string) error

	// WorkSpecNames returns the names of all of the work specs in
	// this namespace.  This may be an empty slice if there are no
	// work specs.  Unless one of the work specs is destroyed,
	// calling GetWorkSpec on one of these names will retrieve the
	// corresponding WorkSpec object.
	WorkSpecNames() ([]string, error)

	// Worker retrieves or creates a Worker object by its name.
	// Every Worker in this Namespace has a nominally unique but
	// client-provided name.  If no Worker exists yet with the
	// requested name, returns a new one with no parent.
	Worker(name string) (Worker, error)

	// Workers retrieves a complete map of worker IDs to worker
	// objects, including parent, child, active, and inactive workers.
	//
	// TODO(dmaze): This interface is likely to change.  The
	// Python coordinate worker system generates about one worker
	// per 10 CPU-seconds (8640 per CPU-day, over a quarter
	// million per day for a 32-core box).  The object hierarchy
	// here requires keeping every worker object that is
	// associated with an attempt.  The most obvious change is to
	// add calls returning active (root) workers and summary
	// statistics, but not have a single "fetch everything" call
	// like this.  Another fairly obvious change is to add
	// (start,limit) windowing like elsewhere.
	Workers() (map[string]Worker, error)
}

// WorkSpecMeta defines control data for a work spec.  This information
// is used to influence the work spec scheduler.
type WorkSpecMeta struct {
	// Priority specifies the absolute priority of this work spec:
	// if this work spec has higher priority than another then this
	// work spec will always run before that other one.  Default
	// priority is the "priority" field in the work spec data, or 0.
	Priority int `json:"priority"`

	// Weight specifies the relative weight of this work spec: if
	// this work spec's priority is twice another one's, then the
	// scheduler will try to arrange for twice as many work units
	// of this work spec to be pending as the other.  Default
	// weight is the "weight" field in the work spec data, or 20
	// minus the "nice" field in the work spec data, defaulting
	// to 20 if neither field is specified at all.
	Weight int `json:"weight"`

	// Paused indicates whether this work unit can generate more
	// work units.  Default is the value of the work spec
	// "disabled" flag, if set, otherwise false.
	Paused bool `json:"paused"`

	// Continuous indicates whether the system can generate new
	// artificial work units for this work spec if there is no
	// other work to do.  If the work spec data does not set the
	// "continuous" flag to true, setting this field has no effect.
	// Defaults to the value of CanBeContinuous.
	Continuous bool `json:"continuous"`

	// CanBeContinuous indicates whether the work spec allows
	// continuous work unit generation.  This is directly set from
	// the "continuous" flag in the work spec data, and
	// WorkSpec.SetMeta() will not change this.
	CanBeContinuous bool `json:"can_be_continuous"`

	// MinMemoryGb specifies the minimum memory required to run
	// this job (or 0 for unlimited).
	MinMemoryGb float64 `json:"min_memory_gb,omitempty"`

	// Interval specifies the minimum time duration before
	// generating another continuous work unit.  Setting this in
	// isolation will affect the second continuous work unit
	// generated; the next one will be not before NextContinuous.
	// Defaults to the value of the "interval" field in the data
	// in seconds, or 0 (i.e., generate more continuous work units
	// immediately) if absent.
	Interval time.Duration `json:"interval"`

	// NextContinuous specifies the earliest time a new continuous
	// work unit could be generated.  This is updated every time
	// a new continuous work unit is produced.  Defaults to a zero
	// time, which should always mean continuous work units could be
	// immediately generated on startup.
	NextContinuous time.Time `json:"next_continuous"`

	// MaxRunning specifies the maximum number of concurrent work
	// units of this work spec that are allowed to execute across
	// the entire system.  If MaxRunning is greater than or equal
	// to PendingCount, no new work units will be allowed.
	// Defaults to the value of the "max_running" field in the
	// work spec data, or 0.  A zero value is interpreted as
	// "unlimited".
	MaxRunning int `json:"max_running"`

	// MaxAttemptsReturned specifies the maximum number of
	// attempts that can be produced by Worker.RequestAttempts().
	// In any case, that function will never return more than
	// MaxRunning work units.  Defaults to the value of the
	// "max_getwork" field in the work spec data, or 0.  A zero
	// value is interpreted as "unlimited".
	MaxAttemptsReturned int `json:"max_attempts_returned"`

	// MaxRetries specifies the maximum number of attempts that
	// can exist for a work unit.  If non-zero, then when
	// Worker.RequestAttempts() produces attempts, it will
	// immediately fail any that have more than this many attempts
	// already.  Defaults to the value of the "max_retries" field
	// in the work spec data, or 0.  A zero value is interpreted
	// as "unlimited".
	MaxRetries int `json:"max_retries"`

	// NextWorkSpecName gives the name of a work spec that runs
	// after this one.  If this is a non-empty string, then when
	// an attempt completes successfully, if the updated work unit
	// data contains a key "outputs", creates work units in this
	// work spec.  WorkSpec.SetMeta() ignores this field.
	// Defaults to the value of the "then" field in the work spec
	// data, or empty string.
	NextWorkSpecName string `json:"next_work_spec_name"`

	// AvailableCount indicates the number of work units in this
	// work spec that could be returned from a
	// Worker.RequestAttempts() call.  These are work units that
	// do not have an active attempt, or that do but it is either
	// expired or retryable.  WorkSpec.Meta() only returns this
	// field if its "withCounts" parameter is true.
	// WorkSpec.SetMeta() ignores this field.
	AvailableCount int `json:"available_count"`

	// PendingCount indicates the number of work units in this
	// work spec that are currently have an active attempt that is
	// in "pending" state, meaning there is a worker performing
	// this work unit.  WorkSpec.Meta() only returns this field
	// if its "withCounts" parameter is true.  WorkSpec.SetMeta()
	// ignores this field.
	PendingCount int `json:"pending_count"`

	// Runtime names the runtime environment this work spec
	// expects to have.  This should generally be a short
	// description such as "python_2", "go", or "java_1.7".
	// WorkSpec.SetMeta() ignores this field.  Defaults to the
	// value of the "runtime" field in the work spec data, or
	// empty string.  For backwards compatibility, empty string
	// should be interpreted as "python_2".
	Runtime string `json:"runtime"`
}

// WorkUnitStatus defines a high-level status of a work unit.
type WorkUnitStatus int

const (
	// AnyStatus is not a real work unit status, but in queries
	// specifies that any status is acceptable.
	AnyStatus WorkUnitStatus = iota

	// AvailableUnit corresponds to work units that do not have
	// active attempts, or if they do have active attempts, they are
	// either Expired or Retryable.  These are work units that
	// Worker.RequestAttempts can return.
	AvailableUnit

	// PendingUnit corresponds to work units that have an active
	// attempt, where that attempt is Pending.  A worker is
	// currently working on these work units.
	PendingUnit

	// FinishedUnit corresponds to work units that have an active
	// attempt, where that attempt is Finished.  The work units
	// have completed successfully.
	FinishedUnit

	// FailedUnit corresponds to work units that have an active
	// attempt, where that attempt is Failed.  The work units have
	// completed unsuccessfully.
	FailedUnit

	// DelayedUnit corresponds to work units that do not have
	// active attempts, but do have a not-before start time that
	// has not yet been reached.
	DelayedUnit
)

// String provides the name of the enum
func (wus *WorkUnitStatus) String() string {
	var s string
	switch *wus {
	case AnyStatus:
		s = "AnyStatus"
	case AvailableUnit:
		s = "AvailableUnit"
	case PendingUnit:
		s = "PendingUnit"
	case FinishedUnit:
		s = "FinishedUnit"
	case FailedUnit:
		s = "FailedUnit"
	case DelayedUnit:
		s = "DelayedUnit"
	default:
		panic("unreachable WorkUnitStatus")
	}
	return s
}

// WorkUnitQuery defines terms to select some subset of the work units
// in a single work spec.  Its zero value selects all work units.
type WorkUnitQuery struct {
	// Names specifies the names of specific work units.  If
	// non-nil, only these work units will be retrieved, provided
	// they meet other criteria.  Specifying the name of a work
	// unit that does not exist is not an error, that work unit
	// will just not be returned.
	Names []string

	// Status specifies high-level status(es).  If non-nil, any
	// status is acceptable.  No work unit whose status is not
	// in this list will be retrieved.
	Statuses []WorkUnitStatus

	// PreviousName specifies the name of the last work unit in a
	// previous query.  This name is lexicographically less than
	// the names of all selected work units.  If empty string,
	// there is no constraint.
	PreviousName string

	// Limit specifies the maximum number of work units to select.
	// If the possible work unit keys are sorted
	// lexicographically, the first Limit keys will be returned.
	Limit int
}

// A WorkSpec defines a collection of related jobs.  For instance, a
// work spec could define a specific function to call, and its work units
// give parameters to that function.  A work spec has a string-keyed
// metadata map, where some keys (e.g., "name") have well-known types
// (string) and meanings.  A work spec also has any number of WorkUnit
// associated with it.
type WorkSpec interface {
	Summarizable

	// Name returns the name of this work spec.
	Name() string

	// Data returns the definition of this work spec.
	Data() (map[string]interface{}, error)

	// SetData changes the definition of this work spec.  It is an
	// error to change "name".  This will also reset fields in the
	// work spec metadata that are derived from the data
	// dictionary: any field in WorkSpecMeta that is specified to
	// default to a field from the data dictionary is reset to
	// that value if present and its specified default otherwise.
	// This in turn means every field in WorkSpecMeta will be
	// reset, except if the NextContinuous time is set the next
	// continuous work unit will still not be generated until that
	// time.
	//
	// Returns ErrNoWorkSpecName if "name" is not in data,
	// ErrBadWorkSpecName if it is not a string, and
	// ErrChangedName if it is different from the existing name.
	// Type errors in other fields (for instance, "weight" is a
	// string) are ignored.
	SetData(data map[string]interface{}) error

	// Meta returns the WorkSpecMeta options for this work spec.
	// If withCounts is true, the WorkSpecMeta.AvailableCount and
	// WorkSpecMeta.PendingCount fields will be filled in; this
	// may be more expensive than other operations.
	Meta(withCounts bool) (WorkSpecMeta, error)

	// SetMeta sets the WorkSpecMeta options for this work spec.
	// The WorkSpecMeta.PendingCount field is ignored.
	SetMeta(WorkSpecMeta) error

	// AddWorkUnit adds a single work unit to this work spec.  If
	// a work unit already exists with the specified name, it is
	// overridden.
	AddWorkUnit(name string, data map[string]interface{}, meta WorkUnitMeta) (WorkUnit, error)

	// WorkUnit retrieves a single work unit by name.  If it does
	// not exist, return ErrNoSuchWorkUnit.
	WorkUnit(name string) (WorkUnit, error)

	// WorkUnits retrieves any number of work units by a query.
	// See the definition of WorkUnitQuery to see what will be
	// retrieved.  This could return an empty map if nothing
	// will be selected.
	WorkUnits(WorkUnitQuery) (map[string]WorkUnit, error)

	// CountWorkUnitStatus retrieves the number of work units in
	// each status in this work spec.  This is mostly useful as an
	// administrator's tool.  It is expected to typically be
	// faster than calling WorkUnits() and iterating through the
	// results.
	CountWorkUnitStatus() (map[WorkUnitStatus]int, error)

	// SetWorkUnitPriorities updates the priorities of multiple
	// work units to all have the same value.
	SetWorkUnitPriorities(WorkUnitQuery, float64) error

	// AdjustWorkUnitPriorities adds a given amount to the
	// priorities of multiple work units.
	AdjustWorkUnitPriorities(WorkUnitQuery, float64) error

	// DeleteWorkUnits deletes work units selected by a query.  If
	// a zero WorkUnitQuery is passed, this deletes all work units
	// in this work spec.  Deleting a work unit also deletes all
	// attempts associated with it, which in turn causes those
	// attempts to not be reported by Worker object queries.
	// Deleting a PendingUnit work unit will not proactively
	// terminate its worker, but the corresponding attempt will no
	// longer appear in either the worker's attempt list or its
	// active attempt list.
	//
	// On success, returns the number of work units actually deleted.
	DeleteWorkUnits(WorkUnitQuery) (int, error)
}

// WorkUnitMeta defines control data for a work unit.  This information
// determines what order work units run in and if they need to be delayed.
type WorkUnitMeta struct {
	// Priority is an arbitrary priority score for this work unit.
	// Higher priority executes sooner.  Zero is the default
	// priority score.
	Priority float64 `json:"priority"`

	// NotBefore specifies the earliest time this work unit is
	// allowed to run.  A zero time allows the work unit to run
	// immediately.
	NotBefore time.Time `json:"not_before"`
}

// A WorkUnit is a single job to perform.  It is associated with a
// specific WorkSpec.  It could be a map entry, and has a name (key)
// and a data map.
type WorkUnit interface {
	// Name returns the name (key) of this work unit.
	Name() string

	// Data returns the data map of this work unit.
	Data() (map[string]interface{}, error)

	// WorkSpec returns the associated work spec.
	WorkSpec() WorkSpec

	// Status gets a high-level status of this work unit.
	// This information is derived from ActiveAttempt().
	Status() (WorkUnitStatus, error)

	// Meta retrieves the combined control metadata for this work
	// unit.
	Meta() (WorkUnitMeta, error)

	// SetMeta updates the control metadata for this work unit.
	// Note that zero values for fields here are meaningful, and
	// so callers will generally want to call Meta() and then
	// update fields.
	SetMeta(WorkUnitMeta) error

	// Priority gets a priority score for this work unit.  Higher
	// priority executes sooner.
	//
	// TODO(dmaze): this call is redundant with Meta() and will
	// be removed in a future version of coordinate, not before
	// coordinate 0.4.0.
	Priority() (float64, error)

	// SetPriority changes the priority score for this work unit.
	// Higher priority executes sooner.
	//
	// TODO(dmaze): this call is redundant with SetMeta() and will
	// be removed in a future version of coordinate, not before
	// coordinate 0.4.0.
	SetPriority(float64) error

	// ActiveAttempt returns the current Attempt for this work
	// unit, if any.  If the work unit is completed, either
	// successfully or unsuccessfully, this is the Attempt that
	// finished it.  This may be an expired Attempt if no other
	// worker has started it yet.  If no Worker is currently
	// working on this work unit, returns nil.
	ActiveAttempt() (Attempt, error)

	// ClearActiveAttempt removes the current active attempt.
	// If there is an active attempt and it is Pending, this does
	// not attempt to proactively kill the attempt and does not
	// remove the attempt from the worker's active attempts list.
	ClearActiveAttempt() error

	// Attempts returns all current and past Attempts for this
	// work unit, if any.  This includes the attempt reported by
	// ActiveAttempt().
	Attempts() ([]Attempt, error)

	// NumAttempts returns the number of times this work unit has
	// been attempted.
	NumAttempts() (int, error)
}

// AttemptRequest describes parameters to Worker.RequestAttempts().
// Its zero value provides reasonable defaults, returning a single
// work unit from any work spec ignoring resource constraints if
// possible.
type AttemptRequest struct {
	// AvailableGb is the amount of memory that can be dedicated
	// to the returned work unit.  If zero, ignore this
	// constraint.  This is compared with the "min_gb" field in
	// the work spec.
	AvailableGb float64 `json:"available_gb"`

	// Lifetime is the minimum requested time to perform this
	// attempt; it must be completed or renewed by this deadline.
	// If zero, use a system-provided default, generally 15
	// minutes.
	Lifetime time.Duration `json:"lifetime"`

	// NumberOfWorkUnits is the number of work units requested.
	// If zero, actually use one.  All of the returned attempts
	// will be for work units in the same work spec.  Fewer work
	// units, maybe as few as zero, can be returned if they are
	// not available.
	NumberOfWorkUnits int `json:"number_of_work_units"`

	// WorkSpecs limits this request to only consider specific
	// work spec(s).  If this is nil or an empty slice, any work
	// spec is acceptable; otherwise only work units from the
	// named work specs will be returned.  It is not an error if
	// these work specs do not exist.  This could cause no work
	// units to be returned if none of the named work specs have
	// available work units, even though other work specs do.
	WorkSpecs []string `json:"work_specs"`

	// Runtimes limits this request to only allow specific
	// language runtimes.  If this is nil or an empty slice, any
	// runtime is acceptable; otherwise only work units from work
	// specs whose WorkSpecMeta.Runtime exactly matches one of
	// these strings will be returned.  This could cause no work
	// units to be returned if none of the work specs with any of
	// these runtimes have work, even though other work specs that
	// use other runtimes do.
	Runtimes []string `json:"runtimes"`
}

// A Worker is a process that is doing work.  Workers may be
// hierarchical, for instance with a parent Worker that does not do
// work itself but supervises its children.  A Worker chooses its own
// name, often a UUID.  It may be performing some number of Attempts;
// typically none if it is only a parent, exactly one if it runs work
// units serially, or multiple if it requests multiple work units in one
// shot or can actively run work units in parallel.
//
// In addition to their name, Worker objects track a worker-provided
// "environment" dictionary.  They may be considered active or inactive,
// and have an expiration time after which they will become inactive.
// Workers have a worker-provided mode, with no semantics assigned to it.
type Worker interface {
	// Name returns the worker-chosen name of the worker.
	Name() string

	// Parent returns the parent of this worker, if any.  If this
	// worker does not have a parent, nil is returned; this is not
	// an error.
	Parent() (Worker, error)

	// SetParent changes (or assigns) the parent of this worker.
	// If a nil worker is passed, clear the parent.  This similarly
	// changes whether this worker is returned from the old and new
	// parents' Children() calls.
	SetParent(Worker) error

	// Children returns the children of this worker, if any.
	Children() ([]Worker, error)

	// Active determines whether or not this worker is currently
	// active.
	Active() (bool, error)

	// Deactivate immediately sets this worker to inactive.
	Deactivate() error

	// Mode gets the mode reported in the last call to Update().
	Mode() (string, error)

	// Data gets the data dictionary passed to the last call to
	// Update().
	Data() (map[string]interface{}, error)

	// Expiration gets the expiration time from the last call to
	// Update().
	Expiration() (time.Time, error)

	// LastUpdate gets the current time at the last call to Update().
	LastUpdate() (time.Time, error)

	// Update refreshes this worker's data.  The worker is set to
	// active.  The data, current and expiration times, and modes
	// are recorded for future calls to Data(), LastUpdate(),
	// Expiration(), and Mode(), respectively.
	Update(data map[string]interface{}, now, expiration time.Time, mode string) error

	// RequestAttempts tries to allocate new work to this worker.
	// With a zero-valued AttemptRequest, this will return at most
	// one new Attempt with a default expiration from any work
	// spec with no resource constraints.  This may return fewer
	// attempts than were requested, maybe even none, if work is
	// not available.
	//
	// Any Attempts returned from this method will also be
	// returned from AllAttempts(), and will be returned from
	// ActiveAttempts() until they are completed or expired.
	RequestAttempts(req AttemptRequest) ([]Attempt, error)

	// MakeAttempt creates an attempt for a specific work unit.
	// On success the new attempt is added to the current and
	// historic attempts for this worker, and becomes the active
	// attempt for the work unit.
	//
	// This method is principally intended for testing and
	// debugging.  It should not be used to resurrect an attempt
	// that has been preempted by another worker; the most likely
	// outcome of this is two workers fighting over the same unit
	// of work.  This will not check the state of the work unit,
	// and could restart a work unit that otherwise is in a
	// terminal state.
	MakeAttempt(WorkUnit, time.Duration) (Attempt, error)

	// ActiveAttempts returns all Attempts this worker is
	// currently performing, or an empty slice if this worker is
	// idle.
	ActiveAttempts() ([]Attempt, error)

	// AllAttempts returns all Attempts this worker has ever
	// performed, including those returned in ActiveAttempts().
	AllAttempts() ([]Attempt, error)

	// ChildAttempts returns any attempts this worker's
	// children are performing.  It is similar to calling
	// ActiveAttempt on each of Children, but is atomic.
	ChildAttempts() ([]Attempt, error)
}

// AttemptStatus is a brief representation of the current status of
// an Attempt.
type AttemptStatus int

const (
	// Pending attempts are not in any other state, and their
	// workers are still working on it.
	Pending AttemptStatus = iota

	// Expired attempts' expiration times have passed.  These
	// attempts should not be the active attempts for their work
	// units, but this constraint is not enforced anywhere.
	Expired

	// Finished attempts have been successfully completed by their
	// workers.
	Finished

	// Failed attempts have been unsuccessfully completed by their
	// workers.
	Failed

	// Retryable attempts have been unsuccessfully completed by
	// their workers, but the failures are identified as transient
	// such that later attempts at redoing the same work would
	// succeed.
	Retryable
)

// An Attempt is a persistent record that some worker is attempting to
// complete some specific work unit.  It has its own copy of the work
// unit data.
type Attempt interface {
	// WorkUnit returns the work unit that is being attempted.
	WorkUnit() WorkUnit

	// Worker returns the worker that is attempting the work.
	Worker() Worker

	// Status returns a high-level status of this Attempt.
	Status() (AttemptStatus, error)

	// Data returns the data map of this work unit, as seen
	// within this attempt.
	Data() (map[string]interface{}, error)

	// StartTime returns the time this attempt began.
	StartTime() (time.Time, error)

	// EndTime returns the time this attempt completed.  If
	// this attempt is not yet complete, returns a zero time.
	EndTime() (time.Time, error)

	// ExpirationTime returns the time by which the worker must
	// complete the work unit.  If this deadline passes, this
	// attempt may expire, and another worker can begin the work
	// unit.
	ExpirationTime() (time.Time, error)

	// Renew attempts to extend the time this worker has to
	// complete the attempt.  You must request a specific
	// duration, with time.Duration(15) * time.Minute being a
	// reasonable default.  Selecting 0 or a negative duration
	// will generally cause this Attempt's status to change to
	// Expired, though it is implementation-dependent whether that
	// change happens before or after this call actuall returns.
	// If data is non-nil, replaces the data stored in this
	// Attempt with a new map.
	//
	// This Attempt must be the active attempt for Renew() to have
	// any affect.  If it is not, the Attempt data will still be
	// updated, but Renew() will return ErrLostLease.
	//
	// The Status() of this Attempt must be Pending for Renew()
	// to have any affect.  If it is Expired but still is the
	// active Attempt, it can also be Renew()ed.  Otherwise, do
	// not update anything and return ErrNotPending.
	Renew(extendDuration time.Duration, data map[string]interface{}) error

	// Expire explicitly transitions an Attempt from Pending to
	// Expired status.  If data is non-nil, also updates the work
	// unit data.  If Status() is already Expired, has no effect.
	//
	// This method is intended to be called by a parent worker to
	// record the fact that it killed off a long-running work unit
	// that was about to expire.  As such it is possible that the
	// parent and child can both be trying to update the same
	// Attempt, resulting in conflicts in the data map.
	//
	// If the Status() of this Attempt is not Pending or Expired,
	// does nothing and returns ErrNotPending.
	Expire(data map[string]interface{}) error

	// Finish transitions an Attempt from Pending to Finished
	// status.  If data is non-nil, also updates the work unit
	// data.
	//
	// If the Status() of this attempt is not Pending, or if it
	// is not both Expired and the current active Attempt, returns
	// ErrNotPending and has no effect.
	Finish(data map[string]interface{}) error

	// Fail transitions an Attempt from Pending to Failed status.
	// If data is non-nil, also updates the work unit data.
	//
	// If the Status() of this attempt is not Pending, or if it
	// is not both Expired and the current active Attempt, returns
	// ErrNotPending and has no effect.
	Fail(data map[string]interface{}) error

	// Retry transitions an Attempt from Pending to Retryable
	// status.  If data is non-nil, also updates the work unit
	// data.  If delay is non-zero, sets the work unit to not
	// be allowed to restart until this time has passed.
	//
	// If the Status() of this attempt is not Pending, or if it
	// is not both Expired and the current active Attempt, returns
	// ErrNotPending and has no effect.
	Retry(data map[string]interface{}, delay time.Duration) error
}
