<p align="center">
<img src="https://www.peopledatalabs.com/images/logos/company-logo.png" style="background-color: white; padding: 5px 10px;" width="250" alt="People Data Labs Logo">
</p>
<h1 align="center">People Data Labs Go Client</h1>
<p align="center">Official Go client for the People Data Labs API.</p>

<p align="center">
  <a href="https://github.com/peopledatalabs/peopledatalabs-go">
    <img src="https://img.shields.io/badge/repo%20status-Active-limegreen" alt="Repo Status">
  </a>&nbsp;
  <a href="https://pkg.go.dev/github.com/peopledatalabs/peopledatalabs-go">
    <img src="https://img.shields.io/github/go-mod/go-version/peopledatalabs/peopledatalabs-go" alt="Go 5.0.1" />
  </a>&nbsp;
  <a href="https://github.com/peopledatalabs/peopledatalabs-go/actions/workflows/test.yaml">
    <img src="https://github.com/peopledatalabs/peopledatalabs-go/actions/workflows/test.yaml/badge.svg" alt="Tests Status" />
  </a>
</p>

## Table of Contents

- [🔧 Installation](#installation)
- [🚀 Usage](#usage)
- [🏝 Sandbox Usage](#sandbox)
- [🌐 Endpoints](#endpoints)
- [📘 Documentation](#documentation)

## 🔧 Installation <a name="installation"></a>

1. To use _peopledatalabs-go_ SDK in your project initialize go modules then run:
    ```bash
    go get github.com/peopledatalabs/peopledatalabs-go/v6
    ```
2. Sign up for a [free PDL API key](https://www.peopledatalabs.com/signup).
3. Set your API key as a environment variable.

## 🚀 Usage <a name="usage"></a>

First, create the PeopleDataLabs client:

```go
package main

import (
    pdl "github.com/peopledatalabs/peopledatalabs-go/v6"
    pdlmodel "github.com/peopledatalabs/peopledatalabs-go/v6/model"
)


func main() {
    apiKey := "YOUR_API_KEY_HERE"
    // Set API KEY as env variable
    // apiKey := os.Getenv("API_KEY")

    client := pdl.New(apiKey)
}
```
Then, send requests to any PDL API Endpoint.

### Person Data

#### Enrichment

```go
params := pdlmodel.EnrichPersonParams{
    PersonParams: pdlmodel.PersonParams{
        Phone: []string{"4155688415"},
    },
}

result, err := client.Person.Enrich(ctx, params)

if err == nil {
    fmt.Printf("Status: %d, FullName: %s\n", result.Status, result.Data.FullName)
}
```

#### Bulk Enrichment

```go
params := pdlmodel.BulkEnrichPersonParams{
    BaseParams:       model.BaseParams{Pretty: true},
    AdditionalParams: model.AdditionalParams{MinLikelihood: 6, IncludeIfMatched: true, Required: "full_name"},
    Requests: []pdlmodel.BulkEnrichSinglePersonParams{
        {
            Params: pdlmodel.PersonParams{
                Profile:  []string{"linkedin.com/in/seanthorne"},
                Location: []string{"SF Bay Area"},
                Name:     []string{"Sean F. Thorne"},
            },
        },
        {
            Params: pdlmodel.PersonParams{
                Profile:   []string{"https://www.linkedin.com/in/haydenconrad/"},
                FirstName: []string{"Hayden"},
                LastName:  []string{"Conrad"},
            },
        },
    },
}

result, err := client.Person.BulkEnrich(ctx, params)
```

#### Search (Elasticsearch)

```go
elasticSearchQuery := map[string]interface{}{
    "query": map[string]interface{}{
        "bool": map[string]interface{}{
            "must": []map[string]interface{}{
                {"term": map[string]interface{}{"location_country": "mexico"}},
                {"term": map[string]interface{}{"job_title_role": "health"}},
            },
        },
    },
}

params := pdlmodel.SearchParams{
    BaseParams: pdlmodel.BaseParams{
        Size: 10,
    },
    SearchBaseParams: pdlmodel.SearchBaseParams{
        Query:   elasticSearchQuery,
        Dataset: "phone,mobile_phone",
    },
}
result, err := client.Person.Search(ctx, params)
```

#### Search (SQL)

```go
sqlQuery := "SELECT * FROM person" +
    " WHERE location_country='mexico'" +
    " AND job_title_role='health'" +
    " AND phone_numbers IS NOT NULL;"

params := pdlmodel.SearchParams{
    BaseParams: pdlmodel.BaseParams{
        Size: 10,
    },
    SearchBaseParams: pdlmodel.SearchBaseParams{
        SQL:     sqlQuery,
        Dataset: "phone,mobile_phone",
    },
}
result, err := client.Person.Search(ctx, params)
```

#### `PDL_ID` (Retrieve API)

```go
params := pdlmodel.RetrievePersonParams{PersonID: "qEnOZ5Oh0poWnQ1luFBfVw_0000"}

result, err := client.Person.Retrieve(ctx, params)
```

#### Bulk Retrieve API

```go
params := pdlmodel.BulkRetrievePersonParams{
    Requests: []pdlmodel.BulkRetrieveSinglePersonParams{
        {ID: "qEnOZ5Oh0poWnQ1luFBfVw_0000"},
        {ID: "PzFD15NINdBWNULBBkwlig_0000"},
    }
}

result, err := client.Person.BulkRetrieve(ctx, params)
```

#### Fuzzy Enrichment (Identify API)

```go
params := pdlmodel.IdentifyPersonParams{PersonParams: pdlmodel.PersonParams{Name: []string{"sean thorne"}}}

result, err := client.Person.Identify(ctx, params)
```

### Person Changelog

```go
params := pdlmodel.ChangelogPersonParams{
    CurrentVersion:  "31.0",
    OriginVersion:   "30.2",
    Type:           "updated",
}

resp, err := client.person.Changelog(ctx, params)
```

### Company Data

#### Enrichment

```go
params := pdlmodel.EnrichCompanyParams{
    CompanyParams: pdlmodel.CompanyParams{Website: "peopledatalabs.com"},
}

result, err := client.Company.Enrich(ctx, params)
```

#### Bulk Enrichment

```go
params := pdlmodel.BulkEnrichCompanyParams{
    BaseParams:       model.BaseParams{Pretty: true},
    Requests: []pdlmodel.BulkEnrichSingleCompanyParams{
        {
            Params: pdlmodel.CompanyParams{
                Profile:  "linkedin.com/company/peopledatalabs",
            },
        },
        {
            Params: pdlmodel.CompanyParams{
                Profile:   "https://www.linkedin.com/company/apple/",
            },
        },
    },
}

result, err := client.Company.BulkEnrich(ctx, params)
```

#### Search (Elasticsearch)

```go
elasticSearchQuery := map[string]interface{}{
    "query": map[string]interface{}{
        "bool": map[string]interface{}{
            "must": []map[string]interface{}{
                {"term": map[string]interface{}{"tags": "bigdata"}},
                {"term": map[string]interface{}{"industry": "financial services"}},
                {"term": map[string]interface{}{"location.country": "united states"}},
            },
        },
    },
}

params := pdlmodel.SearchParams{
    BaseParams:       pdlmodel.BaseParams{Size: 10},
    SearchBaseParams: pdlmodel.SearchBaseParams{Query: elasticSearchQuery},
}

result, err := client.Company.Search(ctx, params)
```

#### Search (SQL)

```go
sqlQuery := "SELECT * FROM company" +
    " WHERE tags='big data'" +
    " AND industry='financial services'" +
    " AND location.country='united states';"

params := pdlmodel.SearchParams{
    BaseParams:       pdlmodel.BaseParams{Size: 10},
    SearchBaseParams: pdlmodel.SearchBaseParams{SQL: sqlQuery},
}

result, err := client.Company.Search(ctx, params)

```

### Supporting APIs

#### Get Autocomplete Suggestions

```go
params := pdlmodel.AutocompleteParams{
    BaseParams:             pdlmodel.BaseParams{Size: 10},
    AutocompleteBaseParams: pdlmodel.AutocompleteBaseParams{Field: "title", Text: "full"},
}

result, err := client.Autocomplete(ctx, params)
```

#### Clean Raw Company Strings

```go
params := pdlmodel.CleanCompanyParams{Name: "peOple DaTa LabS"}

result, err := client.Company.Clean(ctx, params)
```

#### Clean Raw Location Strings

```go
params := pdlmodel.CleanLocationParams{
    LocationParams: pdlmodel.LocationParams{
        Location: "455 Market Street, San Francisco, California 94105, US",
    },
}

result, err := client.Location.Clean(ctx, params)
```

#### Clean Raw School Strings

```go
params := pdlmodel.CleanSchoolParams{
    SchoolParams: pdlmodel.SchoolParams{Name: "university of oregon"},
}

result, err := client.School.Clean(ctx, params)
```

#### Enrich Job Title

```go
params := model.JobTitleParams{
    BaseParams:             model.BaseParams{Pretty: true},
    JobTitleBaseParams:     model.JobTitleBaseParams{JobTitle: "data scientist"},
}

result, err := client.JobTitle(ctx, params)
```

#### Enrich IP

```go
params := model.IPParams{
    BaseParams:             model.BaseParams{Pretty: true},
    IPBaseParams:           model.IPBaseParams{IP: "72.212.42.169"},
}

result, err := client.IP(ctx, params)
```

## 🏝 Sandbox Usage <a name="sandbox"></a>
```go
# To enable sandbox usage, use the following

import (
    pdl "github.com/peopledatalabs/peopledatalabs-go/v6"
    "github.com/peopledatalabs/peopledatalabs-go/v6/api"
    pdlmodel "github.com/peopledatalabs/peopledatalabs-go/v6/model"
)

client := pdl.New(apiKey, api.ClientOptions(func(c *api.Client) {
    c.Sandbox = true
}))
```

## 🌐 Endpoints <a name="endpoints"></a>

**Person Endpoints**

| API Endpoint                                                                           | SDK Function                         |
|----------------------------------------------------------------------------------------|--------------------------------------|
| [Person Enrichment API](https://docs.peopledatalabs.com/docs/enrichment-api)           | `client.Person.Enrich(params)`       |
| [Person Bulk Enrichment API](https://docs.peopledatalabs.com/docs/bulk-enrichment-api) | `client.Person.BulkEnrich(params)`   |
| [Person Search API](https://docs.peopledatalabs.com/docs/search-api)                   | `client.Person.Search(params)`       |
| [Person Retrieve API](https://docs.peopledatalabs.com/docs/person-retrieve-api)        | `client.Person.Retrieve(params)`     |
| [Person Bulk Retrieve API](https://docs.peopledatalabs.com/docs/bulk-person-retrieve)  | `client.Person.BulkRetrieve(params)` |
| [Person Identify API](https://docs.peopledatalabs.com/docs/identify-api)               | `client.Person.Identify(params)`     |
| [Person Changelog API](https://docs.peopledatalabs.com/docs/person-changelog-api)      | `client.Person.Changelog(params)`    |

**Company Endpoints**

| API Endpoint                                                                                    | SDK Function                    |
| ----------------------------------------------------------------------------------------------- |---------------------------------|
| [Company Enrichment API](https://docs.peopledatalabs.com/docs/company-enrichment-api)           | `client.Company.Enrich(params)` |
| [Company Bulk Enrichment API](https://docs.peopledatalabs.com/docs/bulk-company-enrichment-api) | `client.Company.BulkEnrich(params)`   |
| [Company Search API](https://docs.peopledatalabs.com/docs/company-search-api)                   | `client.Company.Search(params)` |

**Supporting Endpoints**

| API Endpoint                                                                            | SDK Function                    |
| --------------------------------------------------------------------------------------- |---------------------------------|
| [Autocomplete API](https://docs.peopledatalabs.com/docs/autocomplete-api)               | `client.Autocomplete(params)`   |
| [Company Cleaner API](https://docs.peopledatalabs.com/docs/cleaner-apis#companyclean)   | `client.Company.Clean(params)`  |
| [Location Cleaner API](https://docs.peopledatalabs.com/docs/cleaner-apis#locationclean) | `client.Location.Clean(params)` |
| [School Cleaner API](https://docs.peopledatalabs.com/docs/cleaner-apis#schoolclean)     | `client.School.Clean(params)`   |
| [Job Title Enrichment API](https://docs.peopledatalabs.com/docs/job-title-enrichment-api) | `client.JobTitle(params)` |
| [IP Enrichment API](https://docs.peopledatalabs.com/docs/ip-enrichment-api) | `client.IP(params)` |

## 📘 Documentation <a name="documentation"></a>

All of our API endpoints are documented at: https://docs.peopledatalabs.com/

These docs describe the supported input parameters, output responses and also provide additional technical context.

As illustrated in the [Endpoints](#endpoints) section above, each of our API endpoints is mapped to a specific method in the API Client. For each of these class methods, **all function inputs are mapped as input parameters to the respective API endpoint**, meaning that you can use the API documentation linked above to determine the input parameters for each endpoint.
