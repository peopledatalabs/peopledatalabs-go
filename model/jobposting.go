package model

import "errors"

const (
	RemoteWorkPolicyRemote = "remote"
	RemoteWorkPolicyOnsite = "onsite"

	SalaryPeriodYear  = "year"
	SalaryPeriodMonth = "month"
	SalaryPeriodWeek  = "week"
	SalaryPeriodDay   = "day"
	SalaryPeriodHour  = "hour"
)

// JobPostingSearchParams covers both the Elasticsearch query form (set Query)
// and the field-based parameter form (leave Query nil and set any of the
// filter fields). The PDL API decides which form to honor based on the body
// it receives.
//
// IsActive is *bool so that "unset" (nil) is omitted from the request while
// an explicit false is still sent. ScrollToken is the opaque base64 token
// returned by the API on the previous page; pass it back unchanged.
type JobPostingSearchParams struct {
	BaseParams
	Query       interface{} `json:"query,omitempty"`
	ScrollToken string      `json:"scroll_token,omitempty"`

	ID                 string `json:"id,omitempty"`
	FirstSeenMin       string `json:"first_seen_min,omitempty"`
	FirstSeenMax       string `json:"first_seen_max,omitempty"`
	DeactivatedDateMin string `json:"deactivated_date_min,omitempty"`
	DeactivatedDateMax string `json:"deactivated_date_max,omitempty"`

	Title        string `json:"title,omitempty"`
	TitleClass   string `json:"title_class,omitempty"`
	TitleRole    string `json:"title_role,omitempty"`
	TitleSubRole string `json:"title_sub_role,omitempty"`
	TitleLevels  string `json:"title_levels,omitempty"`

	CompanyID         string `json:"company_id,omitempty"`
	CompanyName       string `json:"company_name,omitempty"`
	CompanyIndustry   string `json:"company_industry,omitempty"`
	CompanyIndustryV2 string `json:"company_industry_v2,omitempty"`
	CompanyWebsite    string `json:"company_website,omitempty"`
	CompanyProfile    string `json:"company_profile,omitempty"`

	Location    string `json:"location,omitempty"`
	Description string `json:"description,omitempty"`

	SalaryRangeMin int    `json:"salary_range_min,omitempty"`
	SalaryRangeMax int    `json:"salary_range_max,omitempty"`
	SalaryCurrency string `json:"salary_currency,omitempty"`
	SalaryPeriod   string `json:"salary_period,omitempty"`

	RemoteWorkPolicy string `json:"remote_work_policy,omitempty"`
	InferredSkills   string `json:"inferred_skills,omitempty"`
	LastVerifiedMin  string `json:"last_verified_min,omitempty"`
	LastVerifiedMax  string `json:"last_verified_max,omitempty"`

	IsActive *bool `json:"is_active,omitempty"`
}

func (params JobPostingSearchParams) Validate() error {
	if params.Size < 0 || params.Size > 100 {
		return errors.New("job_posting search: size must be between 1 and 100")
	}
	return nil
}

type JobPosting struct {
	ID                 string   `json:"id"`
	Title              string   `json:"title"`
	TitleClass         string   `json:"title_class"`
	TitleRole          string   `json:"title_role"`
	TitleSubRole       string   `json:"title_sub_role"`
	TitleLevels        []string `json:"title_levels"`
	CompanyID          string   `json:"company_id"`
	CompanyName        string   `json:"company_name"`
	CompanyIndustry    string   `json:"company_industry"`
	CompanyIndustryV2  string   `json:"company_industry_v2"`
	CompanyWebsite     string   `json:"company_website"`
	CompanyProfile     string   `json:"company_profile"`
	Location           string   `json:"location"`
	Description        string   `json:"description"`
	URLs               []string `json:"urls"`
	FirstSeen          string   `json:"first_seen"`
	LastVerified       string   `json:"last_verified"`
	DeactivatedDate    string   `json:"deactivated_date"`
	IsActive           bool     `json:"is_active"`
	RemoteWorkPolicy   string   `json:"remote_work_policy"`
	SalaryRangeMin     int      `json:"salary_range_min"`
	SalaryRangeMax     int      `json:"salary_range_max"`
	SalaryCurrency     string   `json:"salary_currency"`
	SalaryPeriod       string   `json:"salary_period"`
	InferredSkills     []string `json:"inferred_skills"`
}

type SearchJobPostingResponse struct {
	Status int `json:"status"`
	Error  struct {
		Type    []string `json:"type"`
		Message string   `json:"message"`
	} `json:"error"`
	Data        []JobPosting `json:"data"`
	Total       int          `json:"total"`
	ScrollToken string       `json:"scroll_token"`
}
