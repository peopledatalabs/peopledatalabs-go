package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/peopledatalabs/peopledatalabs-go/v6/model"

	"github.com/stretchr/testify/assert"
)

func TestJobPosting_Search_Query(t *testing.T) {
	// setup
	jp := JobPosting{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	numResults := 3

	// test
	params := model.JobPostingSearchParams{
		BaseParams: model.BaseParams{Size: numResults},
		Query: map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{"term": map[string]interface{}{"title_role": "engineering"}},
					},
				},
			},
		},
	}
	resp, err := jp.Search(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.LessOrEqual(t, len(resp.Data), numResults)
}

func TestJobPosting_Search_Params(t *testing.T) {
	// setup
	jp := JobPosting{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}
	numResults := 3

	// test
	params := model.JobPostingSearchParams{
		BaseParams:       model.BaseParams{Size: numResults},
		TitleRole:        "engineering",
		RemoteWorkPolicy: model.RemoteWorkPolicyRemote,
	}
	resp, err := jp.Search(context.Background(), params)

	// assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status)
	assert.LessOrEqual(t, len(resp.Data), numResults)
}

func TestJobPosting_Search_SizeOutOfRange(t *testing.T) {
	jp := JobPosting{Client: NewClient(os.Getenv("PDL_API_KEY"), "1.0.0")}

	_, err := jp.Search(context.Background(), model.JobPostingSearchParams{
		BaseParams: model.BaseParams{Size: 101},
		Title:      "engineer",
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "size must be between 1 and 100")
}

func TestJobPosting_Search_IsActiveOptIn(t *testing.T) {
	// IsActive is *bool so a nil (unset) pointer must not appear in the JSON
	// body, but an explicit false must round-trip.
	unset := model.JobPostingSearchParams{Title: "engineer"}
	body, err := json.Marshal(unset)
	assert.NoError(t, err)
	assert.False(t, strings.Contains(string(body), "is_active"))

	explicitFalse := false
	set := model.JobPostingSearchParams{Title: "engineer", IsActive: &explicitFalse}
	body, err = json.Marshal(set)
	assert.NoError(t, err)
	assert.Contains(t, string(body), `"is_active":false`)
}

func TestJobPosting_Search_ScrollTokenOpaque(t *testing.T) {
	// scroll_token must serialize verbatim — no client-side decoding.
	token := "eyJhIjogMX0="
	params := model.JobPostingSearchParams{ScrollToken: token}
	body, err := json.Marshal(params)
	assert.NoError(t, err)
	assert.Contains(t, string(body), `"scroll_token":"`+token+`"`)
}
