# CHANGELOG

## v0.1.3 on 2022-10-04
- Added capability to serve in-memory files when used as library (under `/inmemory/`). If no such files are passed via `InMemoryFile` Setting, the `/inmemory` path can be used for static content.

## v0.1.2 on 2022-10-04
- Moved the webserver into a package. It can now be imported in other projects.
- Added an intermediate layer for the settings (like directory, port, ...). These can be used when the webserver is used in other projects.
- Updated Go version and all dependencies.

## v0.1.1 on 2021-08-30
- Changed default port back to 3000.
- Updated `README.md` for usage and flags.

## v0.1.0 on 2021-08-30
- Complete rework; now uses cobra as cli-framework.
- Now supports flags:
  --verbose for logging each called path
  --port for customizing port if needed (default is 80)
  --ipaddress for customizing ip to bind to (default is 0.0.0.0)
  --directory for serving another directoy as './'
- 'Release' and 'Prerelease' actions now have proper names

## v0.0.3 on 2021-08-28
- added variables for ip and port

## v0.0.2 on 2021-05-09
- added windows-support for cross-platform development

## v0.0.1 on 2021-04-30
- initial release
- added github actions release workflow
