package model

type Person struct {
	Id                              string   `json:"id"`                                    // PDL persistent ID
	FullName                        string   `json:"full_name"`                             // The first and the last name fields appended with a space
	FirstName                       string   `json:"first_name"`                            // A person's first name
	MiddleInitial                   string   `json:"middle_initial"`                        // A person's middle initial
	MiddleName                      string   `json:"middle_name"`                           // A person's middle name
	LastInitial                     string   `json:"last_initial"`                          // The first letter of a person's last name
	LastName                        string   `json:"last_name"`                             // A person's last name
	Gender                          string   `json:"gender"`                                // The person's gender
	BirthYear                       int      `json:"birth_year"`                            // Approximated birth date associated with this person profile. If a profile has a birth_date, the birth_year field will match
	BirthDate                       string   `json:"birth_date"`                            // Date string in YYYY-mm-dd format. Can be accurate to the day (YYYY-MM-DD), month (YYYY-MM) or year (YYYY). If this exists, birth_year will agree with this
	LinkedinUrl                     string   `json:"linkedin_url"`                          // Main linkedin profile for this record based on source agreement
	LinkedinUsername                string   `json:"linkedin_username"`                     // Main linkedin username for this record based on source agreement
	LinkedinId                      string   `json:"linkedin_id"`                           // Main linkedin profile id for this record based on source agreement
	FacebookUrl                     string   `json:"facebook_url"`                          // facebook profile
	FacebookUsername                string   `json:"facebook_username"`                     // facebook username
	FacebookId                      string   `json:"facebook_id"`                           // persistent facebook id associated with a person's facebook profile
	TwitterUrl                      string   `json:"twitter_url"`                           // Twitter URL
	TwitterUsername                 string   `json:"twitter_username"`                      // Twitter Username
	GithubUrl                       string   `json:"github_url"`                            // Main github profile for this record based on source agreement
	GithubUsername                  string   `json:"github_username"`                       // Main github profile username for this record based on source agreement
	WorkEmail                       string   `json:"work_email"`                            // Current Professional email
	PersonalEmails                  []string `json:"personal_emails"`                       // List of all emails tagged as type = personal
	RecommendedPersonalEmail        string   `json:"recommended_personal_email"`            // Highly confident personal email associated with this person
	MobilePhone                     string   `json:"mobile_phone"`                          // Highly confident direct dial mobile phone associated with this person
	Industry                        string   `json:"industry"`                              // The most relevant industry for this record based primarily on their tagged personal industries and secondarily on the industries of the companies that they have worked for
	JobTitle                        string   `json:"job_title"`                             // A person's current job title
	JobTitleRole                    string   `json:"job_title_role"`                        // A person's current job title derived role
	JobTitleSubRole                 string   `json:"job_title_sub_role"`                    // A person's job title derived subrole. Each subrole maps to a role
	JobTitleLevels                  []string `json:"job_title_levels"`                      // A person's current job title derived levels
	JobCompanyId                    string   `json:"job_company_id"`                        // A person's current company's PDL ID
	JobCompanyName                  string   `json:"job_company_name"`                      // A person's current company's name
	JobCompanyWebsite               string   `json:"job_company_website"`                   // A person's current company's website
	JobCompanySize                  string   `json:"job_company_size"`                      // A person's current company's size range
	JobCompanyFounded               int      `json:"job_company_founded"`                   // A person's current company's founded date
	JobCompanyIndustry              string   `json:"job_company_industry"`                  // A person's current company's industry
	JobCompanyLinkedinUrl           string   `json:"job_company_linkedin_url"`              // A person's current company's linkedin url
	JobCompanyLinkedinId            string   `json:"job_company_linkedin_id"`               // A person's current company's linkedin id
	JobCompanyFacebookUrl           string   `json:"job_company_facebook_url"`              // A person's current company's facebook url
	JobCompanyTwitterUrl            string   `json:"job_company_twitter_url"`               // A person's current company's twitter url
	JobCompanyLocationName          string   `json:"job_company_location_name"`             // A person's current company's HQ canonical location
	JobCompanyLocationLocality      string   `json:"job_company_location_locality"`         // A person's current company's HQ locality
	JobCompanyLocationMetro         string   `json:"job_company_location_metro"`            // A person's current company's HQ metro area
	JobCompanyLocationRegion        string   `json:"job_company_location_region"`           // A person's current company's HQ region
	JobCompanyLocationGeo           string   `json:"job_company_location_geo"`              // A person's current company's HQ geo
	JobCompanyLocationStreetAddress string   `json:"job_company_location_street_address"`   // A person's current company's HQ street_address
	JobCompanyLocationAddressLine2  string   `json:"job_company_location_address_line_2"`   // A person's current company's HQ address line 2
	JobCompanyLocationPostalCode    string   `json:"job_company_location_postal_code"`      // A person's current company's HQ postal code
	JobCompanyLocationCountry       string   `json:"job_company_location_country"`          // A person's current company's HQ country
	JobCompanyLocationContinent     string   `json:"job_company_location_continent"`        // A person's current company's HQ continent
	JobCompanyEmployeeCount         int      `json:"job_company_employee_count"`            // A person's current company's employee count
	JobCompanyInferredRevenue       string   `json:"job_company_inferred_revenue"`          // A person's current company's inferred revenue
	JobCompanyMonthGrowthRate       float64  `json:"job_company_12mo_employee_growth_rate"` // A person's current company's 12 month employee growth rate
	JobCompanyTotalFundingRaised    float64  `json:"job_company_total_funding_raised"`      // A person's current company's total funding raised
	JobLastUpdated                  string   `json:"job_last_updated"`                      // YYYY-MM-DD Indicates the timestamp of the most recent source that agrees with this information
	JobStartDate                    string   `json:"job_start_date"`                        // YYYY-MM-DD Indicates the start period of the object. Can be accurate to the day (YYYY-MM-DD), month (YYYY-MM) or year (YYYY)
	LocationName                    string   `json:"location_name"`                         // the current canonical location of the person
	LocationLocality                string   `json:"location_locality"`                     // the current locality of the person
	LocationMetro                   string   `json:"location_metro"`                        // the current MSA of the person
	LocationRegion                  string   `json:"location_region"`                       // the current region of the person
	LocationCountry                 string   `json:"location_country"`                      // the current country of the person
	LocationContinent               string   `json:"location_continent"`                    // the current continent of the person
	LocationStreetAddress           string   `json:"location_street_address"`               // the current street address of the person
	LocationAddressLine2            string   `json:"location_address_line_2"`               // the current address line 2 of the person
	LocationPostalCode              string   `json:"location_postal_code"`                  // the current postal code of the person
	LocationGeo                     string   `json:"location_geo"`                          // the current geo of the person
	LocationLastUpdated             string   `json:"location_last_updated"`                 // YYYY-MM-DD Indicates the timestamp of the most recent source that agrees with this information
	PhoneNumbers                    []string `json:"phone_numbers"`                         // Phone numbers associated with this person profile in E164 format
	Emails                          []struct {
		Address    string  `json:"address"`     // The full parsed email
		Type       string  `json:"type"`        // The type of email either current_professional, professional, personal or null
		FirstSeen  *string `json:"first_seen"`  // The date when this email was first associated to this record
		LastSeen   *string `json:"last_seen"`   // The date when this email was last associated to this record
		NumSources *int    `json:"num_sources"` // The number of sources that have contributed to the association of this email to this record
	} `json:"emails"` // Emails associated with this person profile
	Interests       []string  `json:"interests"`        // Interests associated with the profile
	Skills          []string  `json:"skills"`           // Skills associated with the profile
	LocationNames   []string  `json:"location_names"`   // List of all canonical location names associated with the person
	Regions         []string  `json:"regions"`          // List of regions associated with the person
	Countries       []string  `json:"countries"`        // List of countries associated with a person
	StreetAddresses []Address `json:"street_addresses"` // List of full parsed addresses associated with the person
	Experience      []struct {
		Company struct {
			Name        string   `json:"name"`         // The name associated with the company
			Size        string   `json:"size"`         // The size range of the company
			Id          string   `json:"id"`           // Our current NOT PERSISTENT ids that tie company data to the canonical data
			Founded     int      `json:"founded"`      // YYYY-MM-DD The year that the company was founded
			Industry    string   `json:"industry"`     // The industry associated with the company
			Location    Location `json:"location"`     // Information for the associated company location
			LinkedinUrl string   `json:"linkedin_url"` // The linkedin url associated with the company
			LinkedinId  string   `json:"linkedin_id"`  // The linkedin id associated with the company
			FacebookUrl string   `json:"facebook_url"` // The facebook url associated with the company
			TwitterUrl  string   `json:"twitter_url"`  // The twitter associated with the company
			Website     string   `json:"website"`      // The website associated with the company
			Raw         []string `json:"raw"`          // Raw company names
			Ticker      *string  `json:"ticker"`       // Company Ticker, (only for public companies)
			Type        *string  `json:"type"`         // Company Type
		} `json:"company"` // Information for the associated company
		LocationNames []string `json:"location_names"` // Canonical locations associated with this particular job/experience object (where the person is working, which may or may not be where the company is headquartered.)
		EndDate       string   `json:"end_date"`       // YYYY-MM-DD Indicates the end period of the object
		StartDate     string   `json:"start_date"`     // YYYY-MM-DD Indicates the start period of the object. Can be accurate to the day (YYYY-MM-DD), month (YYYY-MM) or year (YYYY)
		Title         struct {
			Name    string   `json:"name"`     // The inputted title from our data sources with some basic cleaning and mapping in order to help with merging
			Role    string   `json:"role"`     // A person's job title derived role
			SubRole string   `json:"sub_role"` // A person's job title derived subrole. Each subrole maps to a role
			Levels  []string `json:"levels"`   // Levels associated with a title
			Raw     []string `json:"raw"`      // Raw titles
		} `json:"title"` // A dictionary object that provides a canonicalized title, role and level
		IsPrimary bool    `json:"is_primary"` // Indicates if the experience is the primary experience object in our dataset. This experience object will exist in the job_XXX fields
		FirstSeen *string `json:"first_seen"` // The date when this professional position was first associated to this record in our data
		LastSeen  *string `json:"last_seen"`  // The date when this professional position was last associated to this record in our data
		Summary   *string `json:"summary"`    // User-inputted summary of experience
	} `json:"experience"` // Experience objects associated with this person profile
	Education []struct {
		School    School   `json:"school"`     // Information for the associated school
		StartDate string   `json:"start_date"` // Indicates the start period of the object
		EndDate   string   `json:"end_date"`   // Indicates the end period of the object
		Gpa       float64  `json:"gpa"`        // The gpa associated with the given degree
		Degrees   []string `json:"degrees"`    // A list of canonical degrees associated with this education object
		Majors    []string `json:"majors"`     // A list of majors associated with this education object
		Minors    []string `json:"minors"`     // A list of minors associated with this education object
		Raw       []string `json:"raw"`        // Raw education input information. Parsed into the degrees/majors/minors fields
	} `json:"education"` // Education objects associated with this person profile
	Profiles            []Profile `json:"profiles"`             // Social media profiles associated with this person profile
	Phones              []Phone   `json:"phones"`               // A list of the phone numbers known to be associated with this record
	LinkedinConnections *int      `json:"linkedin_connections"` // The number of linkedin connections the person has
	FacebookFriends     *int      `json:"facebook_friends"`     // The number of facebook friends the person has
	NameAliases         []string  `json:"name_aliases"`         // Any additional associated names or aliases beyond the primary one currently displayed in the name field
	PossibleEmails      []struct {
	} `json:"possible_emails"` // Any additional associated emails to this person record with a lower level of confidence than the currently-displayed ones in the emails array
	PossiblePhones          []Phone   `json:"possible_phones"`           // Any additional associated phones to the person record with a lower level of confidence to the currently-displayed ones in the phone_numbers array
	PossibleProfiles        []Profile `json:"possible_profiles"`         // Any additional associated profiles to the person record w/ a lower level of confidence to the currently-displayed ones in the profiles
	PossibleStreetAddresses []Address `json:"possible_street_addresses"` // Any additional associated addresses to the person record with a lower level of confidence to the currently-displayed ones in the street_addresses array
	PossibleBirthDates      []string  `json:"possible_birth_dates"`      // Any additional associated birth dates to the person record with a lower level of confidence to the currently-displayed one in the birth_date field
	PossibleLocationNames   []string  `json:"possible_location_names"`   // Inferred potential locations the person has lived in based on phone area codes, university location, other associations
	JobHistory              []struct {
		FirstSeen string `json:"first_seen"` // The date when this professional position was first associated to this record in our data
		LastSeen  string `json:"last_seen"`  // The date when this professional position was last associated to this record in our data
	} `json:"job_history"` // Any additional professional positions associated to this person record beyond the ones we currently display in the experience array. Usually these are positions that have been removed or changed on resumes
	NumRecords     *int    `json:"num_records"`     // The number of unique raw records contributing to this specific PDL profile
	NumSources     *int    `json:"num_sources"`     // The number of unique sources contributing to this specific PDL profile
	FirstSeen      *string `json:"first_seen"`      // The date when this record was first created in our data
	DatasetVersion string  `json:"dataset_version"` // Explains the current major or minor release number.
	Certifications []struct {
		EndDate      string `json:"end_date"`     // YYYY-MM-DD Expiration of certification
		Name         string `json:"name"`         // Inputted Name of the certification
		Organization string `json:"organization"` // Inputted Organization awarding certification
		StartDate    string `json:"start_date"`   // Inputted Organization awarding certification
	} `json:"certifications"` // Certification objects associated with this person profile
	InferredSalary            *string `json:"inferred_salary"`              // inferred salary range
	InferredYearsExperience   *int    `json:"inferred_years_experience"`    // Inferred years work experience
	JobCompanyTicker          *string `json:"job_company_ticker"`           // Current Company Ticker
	JobCompanyType            *string `json:"job_company_type"`             // Current Company Type
	JobOnetCode               *string `json:"job_onet_code"`                // The 8 digit O*NET code for a person’s current job title, following the 2018 SOC guidelines.
	JobOnetMajorGroup         *string `json:"job_onet_major_group"`         // The O*NET Major Group associated with a person’s current job title.
	JobOnetMinorGroup         *string `json:"job_onet_minor_group"`         // The O*NET Minor Group associated with a person’s current job title.
	JobOnetBroadOccupation    *string `json:"job_onet_broad_occupation"`    // The O*NET Broad Occupation category associated with a person’s current job title.
	JobOnetSpecificOccupation *string `json:"job_onet_specific_occupation"` // The O*NET Detailed Occupation category associated with a person’s current job title.
	JobOnetTitle              *string `json:"job_onet_title"`               // An O*NET Alternative Title associated with a person’s current job title.
	JobSummary                *string `json:"job_summary"`                  // User-inputted summary of experience
	Languages                 []struct {
		Name        string `json:"name"`        // Name of the canonical language the person inputted
		Proficiency int    `json:"proficiency"` // Self-identified proficiency score 1 (limited) - 5 (fluent)
	} `json:"languages"` // Self-identified languages spoken
	Summary *string `json:"summary"` // Self-written summaries tied to person profile (often linkedin summaries)
}

type Company struct {
	DisplayName           string   `json:"display_name"`            // The company's main display name
	Name                  string   `json:"name"`                    // The company's main common name
	Size                  string   `json:"size"`                    // A range representing the number of people working at the company
	EmployeeCount         int      `json:"employee_count"`          // The current number of employees working at the company based on number of PDL profiles.
	LinkedInEmployeeCount int      `json:"linkedin_employee_count"` // The current number of employees working at the company based on number of LinkedIn profiles.
	Id                    string   `json:"id"`                      // PDL company ID. This is currently non-persistent and generated from the company's primary linkedin username
	Founded               int      `json:"founded"`                 // The founded year of the company
	Industry              string   `json:"industry"`                // Self reported industry
	LinkedinId            string   `json:"linkedin_id"`             // Main LinkedIn profile ID for the company
	LinkedinUrl           string   `json:"linkedin_url"`            // Main LinkedIn profile URL for the company
	FacebookUrl           string   `json:"facebook_url"`            // Main Facebook profile URL for the company
	TwitterUrl            string   `json:"twitter_url"`             // Main Twitter profile URL for the company
	Profiles              []string `json:"profiles"`                // List of all known social profile URLs for the company
	Website               string   `json:"website"`                 // Primary company website
	Ticker                string   `json:"ticker"`                  // The company ticker (only for public companies).
	Type                  string   `json:"type"`                    // The company type.
	Summary               string   `json:"summary"`                 // A description of the company.
	Tags                  []string `json:"tags"`                    // Tags associated with the company.
	Headline              string   `json:"headline"`                // The company’s headline summary.
	DisplayNameHistory    []string `json:"display_name_history"`    // A list of all known display names for the company.
	AlternativeNames      []string `json:"alternative_names"`       // A list of names associated with this company.
	AlternativeDomains    []string `json:"alternative_domains"`     // A list of alternate domains associated with this company.
	AffiliatedProfiles    []string `json:"affiliated_profiles"`     // Company IDs that are affiliated with the queried company (parents & subsidiaries).
	Location              Location `json:"location"`                // Location of the company’s current HQ.
	NAICS                 []struct {
		NaicsCode        string `json:"naics_code"`        // The NAICS code associated with a company’s industry classification.
		Sector           string `json:"sector"`            // The industry classification according to the first 2 digits in the NAICS code.
		SubSector        string `json:"sub_sector"`        // The industry classification according to the first 3 digits in the NAICS code.
		IndustryGroup    string `json:"industry_group"`    // The industry classification according to the first 4 digits in the NAICS code.
		NaicsIndustry    string `json:"naics_industry"`    // The industry classification according to the first 5 digits in the NAICS code.
		NationalIndustry string `json:"national_industry"` // The industry classification according to all 6 digits in the NAICS code.
	} `json:"naics"` // Industry classifications for a company according to the North American Industry Classification System (NAICS). A company can (and frequently does) have multiple NAICS codes.
	SIC []struct {
		SicCode        string `json:"sic_code"`        // The SIC code associated with a company’s industry classification.
		MajorGroup     string `json:"major_group"`     // The industry classification according to the first 2 digits in the SIC code.
		IndustryGroup  string `json:"industry_group"`  // The industry classification according to the first 3 digits in the SIC code.
		IndustrySector string `json:"industry_sector"` // The industry classification according to all 4 digits in the SIC code.
	} `json:"sic"` //Industry classifications for a company according to the Standard Industrial Classification (SIC) system. A company can (and frequently does) have multiple SIC codes.
	EmployeeGrowthRate     map[string]float64 `json:"employee_growth_rate"`      // The percentage increase in total headcount from N months prior.
	EmployeeChurnRate      map[string]float64 `json:"employee_churn_rate"`       // The rate of change in employee headcount from N months prior.
	AverageEmployeeTenure  float64            `json:"average_employee_tenure"`   // Average years of experience at the company.
	AverageTenureByRole    map[string]float64 `json:"average_tenure_by_role"`    // Average years of experience at the company by job role.
	AverageTenureByLevel   map[string]float64 `json:"average_tenure_by_level"`   // Average years of experience at the company by job level.
	EmployeeCountByCountry map[string]int     `json:"employee_count_by_country"` // The number of current employees broken out by country.
	TopUsEmployeeMetros    map[string]struct {
		CurrentHeadcount int     `json:"current_headcount"`    // Number of employees in the metro
		MonthGrowthRate  float64 `json:"12_month_growth_rate"` // Growth rate in the metro over the last 12 months, precise to 4th decimal place
	} `json:"top_us_employee_metros"` // The top 10 US metros where employees are based.
	EmployeeCountByMonth        map[string]int            `json:"employee_count_by_month"`          // The number of employees at the end of each month.
	GrossAdditionsByMonth       map[string]int            `json:"gross_additions_by_month"`         // The total number of profiles that joined the company each month.
	GrossDeparturesByMonth      map[string]int            `json:"gross_departures_by_month"`        // The total number of profiles that left the company each month.
	EmployeeCountByMonthByRole  map[string]map[string]int `json:"employee_count_by_month_by_role"`  // The number of employees at the end of each month, broken down by job role.
	EmployeeCountByMonthByLevel map[string]map[string]int `json:"employee_count_by_month_by_level"` // The number of employees at the end of each month, broken down by job level.
	RecentExecHires             []struct {
		JoinedDate                     string   `json:"joined_date"`                         // The month the Exec joined the company
		PdlId                          string   `json:"pdl_id"`                              // ID of the Exec in our Person dataset
		JobTitle                       string   `json:"job_title"`                           // Exec's current job title at the company
		JobTitleRole                   string   `json:"job_title_role"`                      // Exec's current job role at the company. Will be one of the Canonical Job Roles.
		JobTitleSubRole                string   `json:"job_title_sub_role"`                  // Exec's current job subrole at the company. Will be one of the Canonical Job Subroles.
		JobTitleLevels                 []string `json:"job_title_levels"`                    // Exec's current job level at the company. Will be in the Canonical Job Levels.
		PreviousCompanyId              string   `json:"previous_company_id"`                 // ID of company the Exec left
		PreviousCompanyJobTitle        string   `json:"previous_company_job_title"`          // Exec's previous job title at the old company
		PreviousCompanyJobTitleRole    string   `json:"previous_company_job_title_role"`     // Exec's previous job role at the old company. Will be one of the Canonical Job Roles.
		PreviousCompanyJobTitleSubRole string   `json:"previous_company_job_title_sub_role"` // Exec's previous job subrole at the old company. Will be one of the Canonical Job Subroles.
		PreviousCompanyJobTitleLevels  []string `json:"previous_company_job_title_levels"`   // Exec's previous job levels at the old company. Will be in the Canonical Job Levels.
	} `json:"recent_exec_hires"` // The profiles of all of CXOs, owners and VPs that have joined the company in the last 3 months.
	RecentExecDepartures []struct {
		DepartedDate              string   `json:"departed_date"`                  // (Date: YYYY-MM) The month the Exec left the company
		PdlId                     string   `json:"pdl_id"`                         // ID of the Exec in our Person dataset
		JobTitle                  string   `json:"job_title"`                      // Exec's previous job title at the company
		JobTitleRole              string   `json:"job_title_role"`                 // Exec's previous job role at the company. Will be one of the Canonical Job Roles.
		JobTitleSubRole           string   `json:"job_title_sub_role"`             // Exec's previous job subrole at the company. Will be one of the Canonical Job Subroles.
		JobTitleLevels            []string `json:"job_title_levels"`               // Exec's previous job levels at the company. Will be in the Canonical Job Levels.
		NewCompanyId              string   `json:"new_company_id"`                 // ID of company the Exec joined
		NewCompanyJobTitle        string   `json:"new_company_job_title"`          // Exec's current job title at the new company
		NewCompanyJobTitleRole    string   `json:"new_company_job_title_role"`     // Exec's current job role at the new company. Will be one of the Canonical Job Roles.
		NewCompanyJobTitleSubRole string   `json:"new_company_job_title_sub_role"` // Exec's current job subrole at the new company. Will be one of the Canonical Job Subroles.
		NewCompanyJobTitleLevels  []string `json:"new_company_job_title_levels"`   // Exec's current job levels at the new company. Will be in the Canonical Job Levels.
	} `json:"recent_exec_departures"` // The profiles of all of CXOs, owners and VPs that have left the company in the last 3 months.
	TopPreviousEmployersByRole map[string]map[string]int `json:"top_previous_employers_by_role"` // The top 10 previous companies employees worked for and how many current employees were previously employed by them.
	TopNextEmployersByRole     map[string]map[string]int `json:"top_next_employers_by_role"`     // The top 10 companies employees moved to and how many employees moved there.
	TotalFundingRaised         float64                   `json:"total_funding_raised"`           // The total amount of funding raised by the company
	LatestFundingStage         string                    `json:"latest_funding_stage"`           // The latest funding stage of the company
	LastFundingDate            string                    `json:"last_funding_date"`              // The date of the latest funding round
	NumberFundingRounds        int                       `json:"number_funding_rounds"`          // The number of funding rounds the company has raised
	FundingStages              []string                  `json:"funding_stages"`                 // The funding stages the company has raised
	FundingDetails             []struct {
		FundingRoundDate     string   `json:"funding_round_date"`    // The date of the funding round
		FundingRaised        int      `json:"funding_raised"`        // The amount of funding raised in the funding round
		FundingCurrency      string   `json:"funding_currency"`      // The currency of the funding round
		FundingType          string   `json:"funding_type"`          // The type of funding round
		InvestingCompanies   []string `json:"investing_companies"`   // The companies that invested in the funding round
		InvestingIndividuals []string `json:"investing_individuals"` // The individuals that invested in the funding round
	} `json:"funding_details"` // The details of the funding rounds the company has raised
}

type School struct {
	Name        string   `json:"name"`         // The name associated with the school
	Type        string   `json:"type"`         // The type of school
	Id          string   `json:"id"`           // Our current NOT PERSISTENT ids that tie company data to the canonical data
	Location    Location `json:"location"`     // The location associated with the school
	LinkedinUrl string   `json:"linkedin_url"` // The linkedin url associated with the school
	LinkedinId  string   `json:"linkedin_id"`  // The linkedin ID associated with the school
	FacebookUrl string   `json:"facebook_url"` // The facebook url associated with the school
	TwitterUrl  string   `json:"twitter_url"`  // The twitter url associated with the school
	Website     string   `json:"website"`      // The website associated with the school, could include subdomains
	Domain      string   `json:"domain"`       // The website associated with the school
	Raw         []string `json:"raw"`          // Raw school names
	Summary     *string  `json:"summary"`      // User-inputted summary of education
}

type Location struct {
	Name          string `json:"name"`           // The canonical location name associated with the company HQ
	Locality      string `json:"locality"`       // Company locality
	Region        string `json:"region"`         // Company region
	Metro         string `json:"metro"`          // Company metro area
	Country       string `json:"country"`        // Company country
	Continent     string `json:"continent"`      // The continent associated with the company HQ
	StreetAddress string `json:"street_address"` // Company HQ address
	AddressLine2  string `json:"address_line_2"` // The adress line 2 associated with the company HQ
	PostalCode    string `json:"postal_code"`    // The postal code associated with the company HQ
	Geo           string `json:"geo"`            // The geo code associated with the company HQ
}

type Phone struct {
	Number     string `json:"number"`      // An individual phone number associated with this record
	FirstSeen  string `json:"first_seen"`  // The date when this phone number was first associated to this record
	LastSeen   string `json:"last_seen"`   // The date when this phone number was last associated to this record
	NumSources int    `json:"num_sources"` // The number of sources that have contributed to the association of this phone number to this record
}

type Profile struct {
	Network    string  `json:"network"`     // The network the profile exists on
	Id         string  `json:"id"`          // The persistent id related to this social profile (varies by social network)
	Url        string  `json:"url"`         // The url of the social profile
	Username   string  `json:"username"`    // The username associated with the profile
	FirstSeen  *string `json:"first_seen"`  // The date when this profile was first associated to this record
	LastSeen   *string `json:"last_seen"`   // The date when this profile was last associated to this record
	NumSources *int    `json:"num_sources"` // The number of sources that have contributed to the association of this profile to this record
}

type Address struct {
	StreetAddress string  `json:"street_address"` // The street address associated with the location object
	AddressLine2  string  `json:"address_line_2"` // The secondary street address associated with the location object
	Name          string  `json:"name"`           // A string that appends location fields together to create a standard location field
	Locality      string  `json:"locality"`       // The administrative locality associated with the location object
	Metro         string  `json:"metro"`          // The metro area associated with the location object
	Region        string  `json:"region"`         // The administrative region associated with the location object
	PostalCode    string  `json:"postal_code"`    // The postal code associated with the location object
	Country       string  `json:"country"`        // The country associated with the location object
	Geo           string  `json:"geo"`            // The geolocation associated with the location object in latitude, longitude format
	Continent     string  `json:"continent"`      // The continent associated with the country in the location object
	FirstSeen     *string `json:"first_seen"`     // The date when this street address was first associated to this record
	LastSeen      *string `json:"last_seen"`      // The date when this street address was last associated to this record
	NumSources    *int    `json:"num_sources"`    // The number of sources that have contributed to the association of this street address to this record
}
