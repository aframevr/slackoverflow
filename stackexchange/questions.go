package stackexchange

import (
	"fmt"

	"github.com/aframevr/slackoverflow/std"
)

// Questions - https://api.stackexchange.com/docs/questions-by-ids
type Questions struct {
	Parameters Parameters
	Paging
	Result *QuestionsWrapperObj
}

// Init initializes Questions module
func (q *Questions) Init() {
	q.Parameters.Allow("site", "stackoverflow",
		"site where to check questions from")

	q.Parameters.Allow("sort", "activity",
		"The sorts accepted by this method operate on the follow fields of the question object: activity, creation, votes")
	q.Parameters.Allow("order", "asc",
		"Order results is ascending or descending")
	q.Parameters.Allow("fromdate", "",
		"From which date to search")
	q.Parameters.Allow("todate", "",
		"Up to which date to search")

	q.Parameters.Allow("filter", "!6hYwbNNZ(*eH3a3)XT0aZCOGTo-kwAtAoVF5vC378NPI6Y",
		"Defined Custom Filters https://api.stackexchange.com/docs/filters")
	q.Parameters.Allow("pagesize", 100,
		"API. page starts at and defaults to 1, pagesize can be any value between 0 and 100")
	q.Parameters.Allow("page", 1,
		"Current page to be fetched")
	q.Parameters.Allow("min", "",
		"min and max specify the range of a field must fall in (that field being specified by sort)")
	q.Parameters.Allow("max", "",
		"min and max specify the range of a field must fall in (that field being specified by sort)")
	q.Parameters.Allow("key", stackexchange.GetAPIKey(),
		"Pass this as key when making requests against the Stack Exchange API to receive a higher request quota.")
}

// DrawQuery output table of questions query
func (q *Questions) DrawQuery(ids string) {
	std.Hr()
	std.Body("Questions query will be performed with following parameters")
	query := std.NewTable("parameter", "defined value", "default", "description")
	for key, allowed := range q.Parameters.GetAllowed() {
		query.AddRow(key, q.Parameters.ValueOf(key), allowed.String(), allowed.Decription())
	}
	var idss string
	if len(ids) > 10 {
		idss = ids[0:10]
	} else {
		idss = ids
	}
	query.AddRow("ids", idss+"...", "",
		"{ids} can contain up to 100 semicolon delimited ids, to find ids programatically look for question_id ")
	query.Print()
	std.Hr()
	url, _ := q.GetURL(ids)
	std.Body("Resulting Url: %s", url)
	std.Hr()
}

// Get request
func (q *Questions) Get(ids string) (bool, error) {

	url, err := q.GetURL(ids)

	if err != nil {
		return false, err
	}

	err = std.HTTPGetByURL(url, &q.Result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	q.Paging.curentPage = q.Result.Page
	q.Paging.hasMore = q.Result.HasMore
	stackexchange.SetQuotaMax(q.Result.QuotaMax)
	stackexchange.SetQuotaRemaining(q.Result.QuotaRemaining)

	return true, err
}

// GetURL composed from current parameters
func (q *Questions) GetURL(ids string) (string, error) {
	// Apply default parameters
	q.Parameters.ApplyDefaults()

	// Make sure to set failing id if no ids are supplied
	if len(ids) == 0 {
		ids = "100"
	}

	endpoint, err := stackexchange.GetEndpont("questions/" + ids)
	query := endpoint.Query()

	// Remove IDS since that is actually set in path
	q.Parameters.Delete("ids")

	// Apply defined parameters
	for param, value := range q.Parameters.GetApplied() {
		query.Set(param, value.String())
	}

	endpoint.RawQuery = query.Encode()

	return endpoint.String(), err
}
