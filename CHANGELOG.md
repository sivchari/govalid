# Changelog

## [v1.5.0](https://github.com/sivchari/govalid/compare/v1.4.0...v1.5.0) - 2025-09-06
- Remove obsolete benchmarks by @BlackBuck in https://github.com/sivchari/govalid/pull/123
- arrange order by @sivchari in https://github.com/sivchari/govalid/pull/126
- use strings.ToLower to be consist between OSs by @sivchari in https://github.com/sivchari/govalid/pull/127
- rename files by @sivchari in https://github.com/sivchari/govalid/pull/128
- Collect multiple errors on validation run by @egor-denysenko in https://github.com/sivchari/govalid/pull/121
- Adding benchmark syncing to GHA by @BlackBuck in https://github.com/sivchari/govalid/pull/124
- add go report card badge in readme by @egor-denysenko in https://github.com/sivchari/govalid/pull/133
- always run sync-benchmark workflow by @sivchari in https://github.com/sivchari/govalid/pull/139
- Fix benchmark by @sivchari in https://github.com/sivchari/govalid/pull/141
- improve benchmark script by @sivchari in https://github.com/sivchari/govalid/pull/142
- Add pre-commit hooks by @t4kamura in https://github.com/sivchari/govalid/pull/137
- Handle multiple marker by @sivchari in https://github.com/sivchari/govalid/pull/125
- fix golden test by @sivchari in https://github.com/sivchari/govalid/pull/143
- chore: add govulncheck workflow by @shiiyan in https://github.com/sivchari/govalid/pull/146
- Generate validator interface for middleware by @shiiyan in https://github.com/sivchari/govalid/pull/136
- Run formatting file and sorting impport block by @sivchari in https://github.com/sivchari/govalid/pull/145
- chore: remove Hugo build artifacts from version control by @sivchari in https://github.com/sivchari/govalid/pull/148
- fix GHA user by @sivchari in https://github.com/sivchari/govalid/pull/150
- Add field path tracking for nested structures by @t4kamura in https://github.com/sivchari/govalid/pull/134
- docs: clarify error detection timing by @ras0q in https://github.com/sivchari/govalid/pull/138
- Add tagpr to control release version by @sivchari in https://github.com/sivchari/govalid/pull/152
- prepare release v1.5.0 by @sivchari in https://github.com/sivchari/govalid/pull/154
- update goreleaser by @sivchari in https://github.com/sivchari/govalid/pull/155

## [v1.4.0](https://github.com/sivchari/govalid/compare/v1.3.0...v1.4.0) - 2025-07-31
- Introduced registry-based system for complete automation of validator registration by @sivchari in https://github.com/sivchari/govalid/pull/90
- Restore length test by @sivchari in https://github.com/sivchari/govalid/pull/110
- add workflow to check generated code by @sivchari in https://github.com/sivchari/govalid/pull/115
- fix generation code by @sivchari in https://github.com/sivchari/govalid/pull/116
- Fix nested structures generate files even without marker comments by @t4kamura in https://github.com/sivchari/govalid/pull/97
- Adding alpha validator by @BlackBuck in https://github.com/sivchari/govalid/pull/79
- Remove reset stop timer from benchmark tests by @shiiyan in https://github.com/sivchari/govalid/pull/113
- implement numeric validation by @ferenc-zagon in https://github.com/sivchari/govalid/pull/112
- follow up benchmark syntax by @sivchari in https://github.com/sivchari/govalid/pull/122

## [v1.3.0](https://github.com/sivchari/govalid/compare/v1.2.0...v1.3.0) - 2025-07-24
- run fmt, then add format and mod diff check workflow by @sivchari in https://github.com/sivchari/govalid/pull/93
- enable gosec by @sivchari in https://github.com/sivchari/govalid/pull/95
- enable gci in golangci-lint by @ccoVeille in https://github.com/sivchari/govalid/pull/98
- enable thelper in golangci-lint by @ccoVeille in https://github.com/sivchari/govalid/pull/99
- make workflows pull_request_target by @sivchari in https://github.com/sivchari/govalid/pull/101
- Add nolintlint configuration by @sivchari in https://github.com/sivchari/govalid/pull/102
- fix release check workflow by @sivchari in https://github.com/sivchari/govalid/pull/107
- Feature/length marker by @taua-almeida in https://github.com/sivchari/govalid/pull/92
- fix fmt by @sivchari in https://github.com/sivchari/govalid/pull/108
- delete previous fuzzing result comments by @sivchari in https://github.com/sivchari/govalid/pull/109

## [v1.2.0](https://github.com/sivchari/govalid/compare/v1.0.0...v1.2.0) - 2025-07-16
- update docs by @sivchari in https://github.com/sivchari/govalid/pull/71
- update doc by @sivchari in https://github.com/sivchari/govalid/pull/73
- Contain structname by @sivchari in https://github.com/sivchari/govalid/pull/74

## [v1.1.0](https://github.com/sivchari/govalid/compare/v1.0.1...v1.1.0) - 2025-07-15
- run benchmark on GHA automatically by @sivchari in https://github.com/sivchari/govalid/pull/56
- fix benchmark.yaml by @sivchari in https://github.com/sivchari/govalid/pull/57
- Revert run bench automatically gha by @sivchari in https://github.com/sivchari/govalid/pull/60
- improve doc by @sivchari in https://github.com/sivchari/govalid/pull/61
- add CONTRIBUTING.md by @sivchari in https://github.com/sivchari/govalid/pull/63
- Feature/cel marker by @sivchari in https://github.com/sivchari/govalid/pull/58

## [v1.0.1](https://github.com/sivchari/govalid/commits/v1.0.1) - 2025-07-11
- Add markers analyzer by @sivchari in https://github.com/sivchari/govalid/pull/1
- add markers registry to manage some analyzers more easily by @sivchari in https://github.com/sivchari/govalid/pull/2
- update tools module path by @sivchari in https://github.com/sivchari/govalid/pull/3
- Implement required marker by @sivchari in https://github.com/sivchari/govalid/pull/4
- Setup documentation and add benchmark by @sivchari in https://github.com/sivchari/govalid/pull/13
- fix test by @sivchari in https://github.com/sivchari/govalid/pull/14
- Add min marker by @sivchari in https://github.com/sivchari/govalid/pull/19
- Add max marker by @sivchari in https://github.com/sivchari/govalid/pull/20
- rename min/max to lt/gt by @sivchari in https://github.com/sivchari/govalid/pull/21
- Use mit license by @sivchari in https://github.com/sivchari/govalid/pull/22
- Implement MaxLength marker with interface-based import system by @sivchari in https://github.com/sivchari/govalid/pull/23
- Implement MaxItems marker for slice/array validation by @sivchari in https://github.com/sivchari/govalid/pull/27
- Implement MinItems marker by @sivchari in https://github.com/sivchari/govalid/pull/24
- Implement MinLength marker by @sivchari in https://github.com/sivchari/govalid/pull/29
- Implement GTE markers by @sivchari in https://github.com/sivchari/govalid/pull/25
- Implement LTE marker by @sivchari in https://github.com/sivchari/govalid/pull/26
- Feature/enum marker by @sivchari in https://github.com/sivchari/govalid/pull/30
- Implement email marker with HTML5-compliant validation by @sivchari in https://github.com/sivchari/govalid/pull/34
- Implement UUID marker with RFC 4122 validation by @sivchari in https://github.com/sivchari/govalid/pull/36
- Implement URL marker with HTTP/HTTPS validation by @sivchari in https://github.com/sivchari/govalid/pull/35
- refactor by @sivchari in https://github.com/sivchari/govalid/pull/37
- Add fuzz test by @sivchari in https://github.com/sivchari/govalid/pull/38
- fix artifact actions by @sivchari in https://github.com/sivchari/govalid/pull/39
- add goreleaser by @sivchari in https://github.com/sivchari/govalid/pull/40
- Create govalid pages. by @sivchari in https://github.com/sivchari/govalid/pull/41
- update docs flow by @sivchari in https://github.com/sivchari/govalid/pull/42
- add token to push homebrew-tap by @sivchari in https://github.com/sivchari/govalid/pull/43
- Add token by @sivchari in https://github.com/sivchari/govalid/pull/44
- fix ENV to Env by @sivchari in https://github.com/sivchari/govalid/pull/45
- separate docs by @sivchari in https://github.com/sivchari/govalid/pull/46
- update docs by @sivchari in https://github.com/sivchari/govalid/pull/47
- fix test by @sivchari in https://github.com/sivchari/govalid/pull/48
- add benchmark to compete other major Go validation libraries by @sivchari in https://github.com/sivchari/govalid/pull/49

## [v1.0.0](https://github.com/sivchari/govalid/commits/v1.0.0) - 2025-07-10
