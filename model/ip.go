package model

import (
	"errors"
)

type IPBaseParams struct {
	IP                string `json:"ip,omitempty" url:"ip,omitempty"`                                   // IP that is used as the seed for enrichment
	ReturnIPLocation  bool   `json:"return_ip_location,omitempty" url:"return_ip_location,omitempty"`   // If true, the response will include the location of the IP
	ReturnIPMetadata  bool   `json:"return_ip_metadata,omitempty" url:"return_ip_metadata,omitempty"`   // If true, the response will include the metadata of the IP
	ReturnPerson      bool   `json:"return_person,omitempty" url:"return_person,omitempty"`             // If true, the response will include the person fields
	ReturnIfUnmatched bool   `json:"return_if_unmatched,omitempty" url:"return_if_unmatched,omitempty"` // If true, the response will return metadata/location even if no company is found
}

type IPParams struct {
	BaseParams
	IPBaseParams
}

func (params IPParams) Validate() error {
	if params.IP == "" {
		return errors.New("ip: 'IP' para must not be empty")
	}
	return nil
}

type IPResponse struct {
	Status int      `json:"status"`
	Data   IPResult `json:"data"`
}

type IPResult struct {
	IP      IPInfo    `json:"ip"`
	Company IPCompany `json:"company"`
	Person  IPPerson  `json:"person"`
}

type IPInfo struct {
	Address  string     `json:"address"`
	Metadata IPMetadata `json:"metadata"`
	Location IPLocation `json:"location"`
}

type IPMetadata struct {
	Version int  `json:"version"`
	Mobile  bool `json:"mobile"`
	Hosting bool `json:"hosting"`
	Proxy   bool `json:"proxy"`
	Tor     bool `json:"tor"`
	VPN     bool `json:"vpn"`
	Relay   bool `json:"relay"`
	Service bool `json:"service"`
}

type IPLocation struct {
	Name       string `json:"name"`
	Locality   string `json:"locality"`
	Region     string `json:"region"`
	Metro      string `json:"metro"`
	Country    string `json:"country"`
	Continent  string `json:"continent"`
	PostalCode string `json:"postal_code"`
	Geo        string `json:"geo"`
	Timezone   string `json:"timezone"`
}

type IPCompany struct {
	Confidence      string            `json:"confidence"`
	ID              string            `json:"id"`
	Domain          string            `json:"domain"`
	Name            string            `json:"name"`
	Location        IPCompanyLocation `json:"location"`
	Size            string            `json:"size"`
	Industry        string            `json:"industry"`
	InferredRevenue string            `json:"inferred_revenue"`
	EmployeeCount   int               `json:"employee_count"`
	Tags            []string          `json:"tags"`
}

type IPCompanyLocation struct {
	Name          string `json:"name"`
	Locality      string `json:"locality"`
	Region        string `json:"region"`
	Metro         string `json:"metro"`
	Country       string `json:"country"`
	Continent     string `json:"continent"`
	StreetAddress string `json:"street_address"`
	AddressLine2  string `json:"address_line_2"`
	PostalCode    string `json:"postal_code"`
	Geo           string `json:"geo"`
}

type IPPerson struct {
	Confidence      string   `json:"confidence"`
	JobTitleSubrole string   `json:"job_title_subrole"`
	JobTitleRole    string   `json:"job_title_role"`
	JobTitleLevels  []string `json:"job_title_levels"`
}
