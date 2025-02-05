package findupes

type (
	Result struct {
		SourceID string
		MatchID  string
		Accuracy Accuracy
	}

	Results []Result
)

func NewResult(src, match string, acc Accuracy) Result {
	return Result{src, match, acc}
}

func (r Result) Export() []string {
	return []string{r.SourceID, r.MatchID, r.Accuracy.String()}
}

func (rs Results) Export() [][]string {
	res := make([][]string, len(rs))

	for i, r := range rs {
		res[i] = r.Export()
	}

	return res
}
