# CHANGELOG

## v1.3.0 (2021-05-09)

### Others

- chore(readme): update badges (2021-05-09)

- chore: bump year and name in license (2021-05-09)

- chore: bump go in github action (2021-05-09)

- build: remove Makefile (2021-05-09)

- Skip decoding if Content-Length is zero (#63) (2020-10-29)

- Replace Travis CI with Github Actions (2020-10-11)

- Update .travis.yml with Go v1.13 and v1.14 (2020-07-11)

- Read body to completion before close (#59) (2020-05-13)

- Remove README links to missing projects (#58) (2019-11-23)

- Mention that Do/Receive won't decode for status code 204 (#56) (2019-09-01)

- Bump examples module's version of sling module (2019-07-07)

- Update CHANGES for v1.3.0 release (2019-07-07)

- Add a go.mod with go-querystring dependency (2019-07-07)

- Add Go v1.12 and remove Go v1.9 (2019-03-05)

- Replace test script with a make target (2018-11-25)

- Limit .travis.yml email notifications (2018-11-25)

- Add changelog entry for ResponseDecoder (2018-11-25)

- Add ResponseDecoder interface (#49) (2018-11-26)

- Update CHANGES for v1.2.0 (2018-11-17)

- Update .travis.yml for Go v1.11, remove v1.8 (2018-11-17)

- Switch to the golang.org/x/lint/golint import path (2018-10-13)

- Add Option, Trace, and Connect HTTP methods (2018-01-25)

- Fix gopher with slingshot image path (2018-06-03)

- Update .travis.yml for Go 1.9 and 1.10 (2018-05-19)

- Merge pull request #34 from ae6rt/tracing (2017-06-28)

- Merge pull request #31 from AaronO/fix/eof-on-204 (2017-03-21)

- Add Go 1.8 to .travis.yml (2017-02-19)

- Add CHANGES for v1.1.0 (2016-12-18)

- Merge pull request #26 from dghubble/json-decode-any-content-type (2016-10-11)

- Merge pull request #25 from dghubble/body-provider-candidate (2016-10-10)

- Add test script and Go 1.7 (2016-08-20)

- Merge branch 'add-doer-interface' (2016-07-22)

- Merge pull request #20 from rliu054/master (2016-07-12)

- Merge pull request #18 from JesusTinoco/patch-1 (2016-06-30)

- Merge branch 'AaronO-feature/basicauth' (2016-06-14)

- Add HN API to the examples list (2016-04-23)

- Add Go 1.6 to test matrix, remove 1.3 and 1.4 (2016-04-09)

- Show build status of master branch only (2016-02-17)

- Add swagger-codegen to notable users (2016-02-17)

- Fix a test printf format which is an error on Go tip (2016-02-03)

- Update godocs and gocover links from http to https (2016-02-03)

- Run tests on travis Go 1.5, stop testing Go 1.2 (2015-11-05)

- Improve README, godocs, and example API client (2015-11-04)

- Update CHANGES and doc for added Body setter (2015-09-14)

- Add Sling plain Body setter (2015-09-10)

- Add README link about Contributing (2015-08-30)

- Organize imports (2015-08-17)

- Add to list of open-source projects using Sling (2015-06-15)

- Fix some typos (2015-05-23)

- Merge pull request #7 from dghubble/dev (2015-05-23)

- Fix all non-breaking golint violations (2015-04-26)

- Fix typo in README example (2015-04-26)

- Update README with details about the new BodyStruct method (2015-04-25)

- Add 'go vet' to travis yml config (2015-04-21)

- Add missing arguments to Errorf in TestSlingNew (2015-04-17)

- Add BodyStruct setter which adds a url encoded struct form to Body (2015-04-21)

- Add gocover README badge (2015-04-14)

- Add README details about setting Headers, expand doc.go (2015-04-14)

- Set Content-Type application/json when jsonBody is non-nil (2015-04-14)

- Improve Github example to show unauthenticated and OAuth2 requests (2015-04-13)

- Add support for Add'ing and Set'ing Headers (2015-04-13)

- Add details to README usage intro (2015-04-13)

- Ensure Sling jsonBody value is copied to child Slings (2015-04-13)

- Add JsonBody description to README (2015-04-12)

- Add JsonBody setter for encoding JSON structs to Request.Body (2015-04-10)

- Add Gopher image to README (2015-04-05)

- Drop Go 1.0 from travis.yml, not supported (2015-04-05)

- Add .travis.yml file and ci/doc README badges (2015-04-05)

- Rename Request to New and HttpRequest to Request (2015-04-05)

- Rename Do to Receive and rename Fire to Do (2015-04-05)

- Expose Sling HttpClient and improve README (2015-04-05)

- Add QueryStruct setter to support encoding url query structs (2015-04-05)

- Added Client setter and made New() take no parameters (2015-04-04)

- Change Path to extend the Sling RawUrl (2015-04-04)

- Add README description and LICENSE (2015-04-04)

- Add Head, Put, Patch, and Delete setters and tests (2015-04-04)

- Handle nil v parameter to Sling Fire and Do methods (2015-04-04)

- Add Sling copier, setters, new http Request method, and Do (2015-04-02)

- Add tiny Github API example with example usage (2015-04-02)

- Add initial Sling which decodes Response Body JSON (2015-04-01)

- Doc improvements and fixes (2015-05-23)

- Fix default Request method to be "GET" instead of "" (2015-05-22)

- Improve documentation about Sling extension via New() (2015-05-22)

- Internalize Sling HTTPClient and Header fields (2015-05-22)

- Check Content-Type is application/json before JSON decoding responses (2015-05-22)

- Add CHANGES.md Changelog (2015-05-19)

- Breaking name changes for consistency and golint compliance (2015-05-19)

- Simplify oauth2 example and make example pass golint (2015-05-19)

- Add tests for Receive method with success and failure responses (2015-05-17)

- Add tests for Do method with success and failure responses (2015-05-17)

- Change Receive and Do to support JSON decoding error structs (2015-05-17)

- Add basic auth support: Sling.SetBasicAuth() (2016-06-09)

- Add go-smith to the examples apis using sling (2016-06-27)

- remove gocover.io link (2016-07-13)

- Add Doer interface and Doer setter (2016-07-17)

- Stop exposing body provider implementations (2016-10-10)

- Rework body handling to use providers (2016-08-21)

- Decode JSON success/failure regardless of Content-Type (2016-10-10)

- Fix typo (2017-03-12)

- Don't try to decode on 204s (no content) (2017-03-12)

- [issue/33] Add an example to the README that shows how to do Go http tracing. (2017-05-20)