package executor

// Report interface defines a report of a Job. A report must define how the Job went using the
// Status() method and a the description about the report using the String().
type Report interface {
	Status() int
	String() string
}

// Job interface defines a job for the Executor. 'Execute' method is used to start the job which
// must return a Report. String() gives a small description about the Job.
type Job interface {
	Execute() Report
	String() string
}
