package models

type CodingSubmissionResults struct {
	CaseResults []CodingSubmissionCaseResults
}

type CodingSubmissionCaseResults struct {
	CaseName string
	Result   CodingSubmissionResult
}

type CodingSubmissionResult string

const (
	ResultAccepted          CodingSubmissionResult = "ACCEPTED"
	ResultWrongAnswer       CodingSubmissionResult = "WRONG_ANSWER"
	ResultTimeLimitExceeded CodingSubmissionResult = "TIME_LIMIT_EXCEEDED"
	ResultMemoryExhausted   CodingSubmissionResult = "MEMORY_EXHAUSTED"
	ResultRuntimeError      CodingSubmissionResult = "RUNTIME_ERROR"
	ResultPresentationError CodingSubmissionResult = "PRESENTATION_ERROR"
)
