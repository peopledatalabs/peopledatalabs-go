# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [1.0.0] - 2022-06-11

- 🎉 First release!

## [1.0.1] - 2022-06-19

- Incremented default timeout from 5 seconds to 10 seconds
- Added `AdditionalParams` to Bulk Person Retrieve params

## [1.1.0] - 2022-06-25

- Removed methods that allow no context to be provided
- Removed commented parameters on Company Cleaner endpoint

## [1.2.0] - 2022-08-13

- Add Support for Sandbox, Job Title Enrich and Skill Enrich Endpoints

## [1.2.1] - 2022-08-13

- Add Documentation and tests for Sandbox

## [1.2.2] - 2022-08-30

- Move API config public

## [1.2.3] - 2022-08-30

- Update Docs

## [1.2.4] - 2022-04-28

- Add Support for pdl_id

## [1.2.5] - 2022-05-09

- Dependency Updates

## [1.3.0] - 2023-08-01

- Add Support for IP Enrichment

## [1.3.1] - 2023-08-13

- Add return_if_unmatched to IP enrichment

## [1.3.2] - 2023-08-26

- Fix Bulk Person Enrichment 415 Error

## [1.3.3] - 2023-08-28

- Fix Bulk Person Enrichment to Support Base and Additional Params

## [1.3.4] - 2023-08-28

- Add Id to Company Enrichment

## [1.4.0] - 2023-11-20

- Add Support for Beta AutoComplete Endpoint

## [1.4.1] - 2023-11-20

- Update Models for Company and Autocomplete Endpoints

## [1.4.2] - 2024-01-09

- Add Support for Bulk Company Enrichment
- Add LinkedIn Slug to Company Response
- Add Size to Company Enrichment Params

## [1.4.3] - 2024-03-04

- Update Dependencies

## [1.4.4] - 2024-03-04

- Remove beta from AutoComplete Endpoint

## [2.0.0] - 2024-04-02

- Change gender to sex in Person Schema
- Add job_last_changed to Person Schema
- Add job_last_verified to Person Schema

## [2.0.1] - 2024-05-13

- Add DatasetVersion to Company Schema
- Add DatasetVersion to IP Schema

## [3.0.0] - 2024-07-14

- Remove job_last_updated from Person Schema
- Add linkedin_follower_count to Company Schema

## [3.0.1] - 2024-08-05

- Fix Education Schema to include summary

## [3.0.2] - 2024-08-08

- Add Score to company cleaner response

## [3.1.0] - 2024-08-14

- Add UltimateParentTicker to Company Schema
- Add UltimateParentMicExchange to Company Schema
- Add MicExchange to Company Schema

## [3.2.0] - 2024-09-09

- Add Support for Updated Title Roles
- Update Person Schema to include job_title_roles and title.class
- Add Support for Min Confidence in IP Enrichment

## [3.3.0] - 2024-11-28

- Add Headline to Person Schema
- Update Readme

## [3.4.0] - 2024-11-28

- Add Support for Class and Updated Title Roles in Autocomplete

## [4.0.0] - 2025-02-21

- Remove Support for Updated Title Roles

## [5.0.0] - 2025-04-28

- Remove Support for Skill Enrichment API
