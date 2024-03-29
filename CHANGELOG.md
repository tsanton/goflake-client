# Changelog

## [0.6.0](https://github.com/Tsanton/goflake-client/compare/v0.5.0...v0.6.0) (2023-04-27)


### ⚠ BREAKING CHANGES

* role inheritance with ISnowflakePrincipal

### Features

* principal direct ascendants ([9bb4ea8](https://github.com/Tsanton/goflake-client/commit/9bb4ea863e61ff23c41b0b7dda45f9a1deceff9b))


### Code Refactoring

* role inheritance with ISnowflakePrincipal ([c1b82ba](https://github.com/Tsanton/goflake-client/commit/c1b82ba004cc3278b8a830ff1335395a8c79b811))

## [0.5.0](https://github.com/Tsanton/goflake-client/compare/v0.4.0...v0.5.0) (2023-04-27)


### ⚠ BREAKING CHANGES

* remove assets dependency on types, opening up for external package implementation of interface

### Code Refactoring

* remove assets dependency on types, opening up for external package implementation of interface ([7dc662d](https://github.com/Tsanton/goflake-client/commit/7dc662d80219ce09167ea824d981a99d9bc21798))

## [0.4.0](https://github.com/Tsanton/goflake-client/compare/v0.3.1...v0.4.0) (2023-04-03)


### ⚠ BREAKING CHANGES

* utilize tf-provider-snowflake auth

### Code Refactoring

* utilize tf-provider-snowflake auth ([c9a8715](https://github.com/Tsanton/goflake-client/commit/c9a87151fbea62428d7af0abca3721dd3505e248))

## [0.3.1](https://github.com/Tsanton/goflake-client/compare/v0.3.0...v0.3.1) (2023-04-03)


### Bug Fixes

* performance enchancing describes for grants ([05dbe35](https://github.com/Tsanton/goflake-client/commit/05dbe3535a5f4151ba3481b024ac60e11ee090ee))

## [0.3.0](https://github.com/Tsanton/goflake-client/compare/v0.2.0...v0.3.0) (2023-04-03)


### ⚠ BREAKING CHANGES

* lifting principal up into grant_asset

### Code Refactoring

* lifting principal up into grant_asset ([ff5c15a](https://github.com/Tsanton/goflake-client/commit/ff5c15aff1ce592039942178e7ae7958c88e1f84))

## [0.2.0](https://github.com/Tsanton/goflake-client/compare/v0.1.2...v0.2.0) (2023-03-25)


### ⚠ BREAKING CHANGES

* added tag resource

### Features

* added boolean column option to table asset ([daff0f9](https://github.com/Tsanton/goflake-client/commit/daff0f9271697f2e6c642578c7d32a12d26fc7ab))
* added date column option to table asset ([8d22da1](https://github.com/Tsanton/goflake-client/commit/8d22da184f19e2aeb515cc6d095489c8ea7e2032))
* added number column option to table asset ([ee765d5](https://github.com/Tsanton/goflake-client/commit/ee765d567ddce3bf5ca962fe9fa11d3764e58a4f))
* added table asset, describable and entity ([01b1070](https://github.com/Tsanton/goflake-client/commit/01b107026586475226db0d36e8b35d0a083770af))
* added tag resource ([d406d90](https://github.com/Tsanton/goflake-client/commit/d406d907c7e8aa26fe029d369675c61e6bdcb760))
* added time column option to table asset ([cce7cef](https://github.com/Tsanton/goflake-client/commit/cce7cef30308a6bbb5a4564b2fd7decfc2523826))
* added timestamp column option to table asset ([ae693bc](https://github.com/Tsanton/goflake-client/commit/ae693bc147b7a118632f21278c7471872bb3a458))
* added variant column option to table asset ([c8793a3](https://github.com/Tsanton/goflake-client/commit/c8793a375ef6f23d4627a2f645b4fb374993dec7))
* grant_action and account_grant ([327fe8b](https://github.com/Tsanton/goflake-client/commit/327fe8ba577c59b958b233e7d55dd6881083d36f))
* init commit database roles ([e72d8d0](https://github.com/Tsanton/goflake-client/commit/e72d8d0beeaa2dfaedd6190b164b1ebb6afcae6b))
* ISnowflakePrincipal as database owner ([2f909c7](https://github.com/Tsanton/goflake-client/commit/2f909c71c58a1b84e9820bc2485a80443258e4e5))
* role_hierarchy replaces with principal ascendants and descendants ([30481bb](https://github.com/Tsanton/goflake-client/commit/30481bbb8fe7a733354e5967827957eb5d3dc633))


### Bug Fixes

* RoleFutureGrant json attr name error ([aaa8d32](https://github.com/Tsanton/goflake-client/commit/aaa8d32c2f957785e387fc662f50acf838dd2760))

## [0.1.2](https://github.com/Tsanton/goflake-client/compare/v0.1.1...v0.1.2) (2023-01-24)


### Bug Fixes

* added warehouse ([a7425d8](https://github.com/Tsanton/goflake-client/commit/a7425d85d2f1a5a1ab5e07fbd10a3cfbb1608fb4))

## [0.1.1](https://github.com/Tsanton/goflake-client/compare/v0.1.0...v0.1.1) (2023-01-24)


### Features

* added account & database grants ([3f11e88](https://github.com/Tsanton/goflake-client/commit/3f11e88a9744ed83aa0e926bde141460ae0156b4))
* schema object future role grant ([d5b885d](https://github.com/Tsanton/goflake-client/commit/d5b885d0136c45ded8b73a04606ef22972ef4301))
* schema role privileges ([aa43398](https://github.com/Tsanton/goflake-client/commit/aa43398d156a725933d6d3cb6ae3aa03a6b58b3d))

## [0.1.0](https://github.com/Tsanton/goflake-client/compare/v0.1.0...v0.1.0) (2023-01-19)


### Features

* init commit ([627f652](https://github.com/Tsanton/goflake-client/commit/627f652386803bb23ac3c390ee8d31e9d5eefb51))


### Miscellaneous Chores

* release 0.1.0 ([e0bfcd1](https://github.com/Tsanton/goflake-client/commit/e0bfcd1838fa856c53b2aa6bbf1c63135ba2c73c))
* release 0.1.0 ([e840b04](https://github.com/Tsanton/goflake-client/commit/e840b04f425092594ba8465c2f7b9a3804064fe5))
* release 0.1.0 ([6fb13b1](https://github.com/Tsanton/goflake-client/commit/6fb13b15822f16dec20880b6116a9cc374d69386))
