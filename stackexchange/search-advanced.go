package stackexchange

import (
	"fmt"

	"github.com/aframevr/slackoverflow/std"
)

// SearchAdvanced - https://api.stackexchange.com/docs/advanced-search
type SearchAdvanced struct {
	Parameters Parameters
	Paging
	Result *QuestionsWrapperObj
}

// Init initializes Search Advanced module
func (sa *SearchAdvanced) Init() {
	sa.Parameters.Allow("site", "stackoverflow",
		"site where to check questions from")
	sa.Parameters.Allow("q", "",
		"a free form text parameter, will match all question properties based on an undocumented algorithm.")
	sa.Parameters.Allow("accepted", "",
		"true to return only questions with accepted answers, false to return only those without. Omit to elide constraint.")
	sa.Parameters.Allow("answers", "",
		"the minimum number of answers returned questions must have.")
	sa.Parameters.Allow("body", "",
		"text which must appear in returned questions' bodies.")
	sa.Parameters.Allow("closed", "",
		"true to return only closed questions, false to return only open ones. Omit to elide constraint.")
	sa.Parameters.Allow("migrated", "",
		"true to return only questions migrated away from a site, false to return only those not. Omit to elide constraint.")
	sa.Parameters.Allow("notice", "",
		"true to return only questions with post notices, false to return only those without. Omit to elide constraint.")
	sa.Parameters.Allow("nottagged", "",
		"a semicolon delimited list of tags, none of which will be present on returned questions.")
	sa.Parameters.Allow("tagged", "",
		"a semicolon delimited list of tags, of which at least one will be present on all returned questions.")
	sa.Parameters.Allow("title", "",
		"text which must appear in returned questions titles.")
	sa.Parameters.Allow("user", "",
		"the id of the user who must own the questions returned.")
	sa.Parameters.Allow("url", "",
		"a url which must be contained in a post, may include a wildcard.")
	sa.Parameters.Allow("views", "",
		"the minimum number of views returned questions must have.")
	sa.Parameters.Allow("wiki", "",
		"true to return only community wiki questions, false to return only non-community wiki ones. Omit to elide constraint.")
	sa.Parameters.Allow("sort", "creation",
		"The sorts accepted by this method operate on the follow fields of the question object: activity, creation, votes, relevance")
	sa.Parameters.Allow("order", "asc",
		"Order results is ascending or descending")
	sa.Parameters.Allow("fromdate", "",
		"From which date to search")
	sa.Parameters.Allow("todate", "",
		"Up to which date to search")
	sa.Parameters.Allow("filter", "!6hYwbNNZ(*eH3a3)XT0aZCOGTo-kwAtAoVF5vC378NPI6Y",
		"Defined Custom Filters https://api.stackexchange.com/docs/filters")
	sa.Parameters.Allow("pagesize", 100,
		"API. page starts at and defaults to 1, pagesize can be any value between 0 and 100")
	sa.Parameters.Allow("page", 1,
		"Current page to be fetched")
	sa.Parameters.Allow("key", stackexchange.GetAPIKey(),
		"Pass this as key when making requests against the Stack Exchange API to receive a higher request quota.")
}

// Get request
func (sa *SearchAdvanced) Get() (bool, error) {

	url, err := sa.GetURL()

	if err != nil {
		return false, err
	}

	err = std.HTTPGetByURL(url, &sa.Result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	sa.Paging.curentPage = sa.Result.Page
	sa.Paging.hasMore = sa.Result.HasMore
	stackexchange.SetQuotaMax(sa.Result.QuotaMax)
	stackexchange.SetQuotaRemaining(sa.Result.QuotaRemaining)

	return true, err
}

// GetURL composed from current parameters
func (sa *SearchAdvanced) GetURL() (string, error) {
	endpoint, err := stackexchange.GetEndpont("search/advanced")

	query := endpoint.Query()
	// Apply default parameters
	sa.Parameters.ApplyDefaults()

	// Apply defined parameters
	for param, value := range sa.Parameters.GetApplied() {
		query.Set(param, value.String())
	}

	endpoint.RawQuery = query.Encode()

	return endpoint.String(), err
}

// DrawQuery output table of search query
func (sa *SearchAdvanced) DrawQuery() {
	std.Hr()
	std.Body("Search Advanced query will be performed with following parameters")
	query := std.NewTable("parameter", "defined value", "default", "description")
	for key, allowed := range sa.Parameters.GetAllowed() {
		query.AddRow(key, sa.Parameters.ValueOf(key), allowed.String(), allowed.Decription())
	}
	query.Print()
	std.Hr()
	url, _ := sa.GetURL()
	std.Body("Resulting Url: %s", url)
	std.Hr()
}
