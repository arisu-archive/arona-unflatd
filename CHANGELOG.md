# Changelog

## [0.0.3](https://github.com/arisu-archive/arona-unflatd/compare/v0.0.2...v0.0.3) (2026-08-23)


### Bug Fixes

* add missing namespace declaration ([c834b2a](https://github.com/arisu-archive/arona-unflatd/commit/c834b2ab7bf441e05426023863dca74b4ce4a346))
* cannot handle external type ([65855ca](https://github.com/arisu-archive/arona-unflatd/commit/65855cae94fb74c121dfb04f998eccbd0800409d))
* **codegen:** make schema recovery reliable ([#36](https://github.com/arisu-archive/arona-unflatd/issues/36)) ([8bd06ef](https://github.com/arisu-archive/arona-unflatd/commit/8bd06ef44f02ec2804eae82d50e046a3ca29167e))
* **deps:** update Go version to 1.23.x in CI and documentation ([a9ab577](https://github.com/arisu-archive/arona-unflatd/commit/a9ab57783bbfd5a23faa16d9a267c4f114c00187))
* **deps:** update module github.com/bmatcuk/doublestar/v4 to v4.10.0 ([#32](https://github.com/arisu-archive/arona-unflatd/issues/32)) ([360de47](https://github.com/arisu-archive/arona-unflatd/commit/360de47437d3ca458efb565a27539bcec59aaeee))
* **deps:** update module github.com/bmatcuk/doublestar/v4 to v4.9.1 ([#17](https://github.com/arisu-archive/arona-unflatd/issues/17)) ([29cf602](https://github.com/arisu-archive/arona-unflatd/commit/29cf60243e359246dee2cd0eb98d88f46dd0010a))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.26.0 ([#14](https://github.com/arisu-archive/arona-unflatd/issues/14)) ([18606ac](https://github.com/arisu-archive/arona-unflatd/commit/18606ac45c1c33cb8131e92bb46a22fa5d86dda6))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.27.1 ([#21](https://github.com/arisu-archive/arona-unflatd/issues/21)) ([e37efe8](https://github.com/arisu-archive/arona-unflatd/commit/e37efe8b16bf85dfb055f0a12d9214fa29d2c196))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.27.2 ([#22](https://github.com/arisu-archive/arona-unflatd/issues/22)) ([7d4dea2](https://github.com/arisu-archive/arona-unflatd/commit/7d4dea2558df5bfcec4ca250e6ffb80f7751aa17))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.27.3 ([#27](https://github.com/arisu-archive/arona-unflatd/issues/27)) ([e42ab3e](https://github.com/arisu-archive/arona-unflatd/commit/e42ab3e0acb68924b618df95e518ce5a7fd8ec97))
* **deps:** update module github.com/onsi/ginkgo/v2 to v2.28.3 ([#30](https://github.com/arisu-archive/arona-unflatd/issues/30)) ([ed53edc](https://github.com/arisu-archive/arona-unflatd/commit/ed53edc66ddf4380c25de78c54c0daa987b0a332))
* **deps:** update module github.com/onsi/gomega to v1.38.3 ([#28](https://github.com/arisu-archive/arona-unflatd/issues/28)) ([d958150](https://github.com/arisu-archive/arona-unflatd/commit/d958150f60789ccb924ca044117f9414adcec44c))
* **deps:** update module github.com/spf13/cobra to v1.10.1 ([#19](https://github.com/arisu-archive/arona-unflatd/issues/19)) ([015c54d](https://github.com/arisu-archive/arona-unflatd/commit/015c54d155edcd9dce3362def0716e5454cd8f14))
* **deps:** update module github.com/spf13/cobra to v1.10.2 ([#26](https://github.com/arisu-archive/arona-unflatd/issues/26)) ([05fa72f](https://github.com/arisu-archive/arona-unflatd/commit/05fa72f045fb51933a35c1ea93f1d04079e5fa11))
* ensure schemas are assigned the correct namespace before processing ([bc4844c](https://github.com/arisu-archive/arona-unflatd/commit/bc4844cee44a015fd3dd60e81f6d51a8050eab77))
* handle Nullable types with suffix "?" in ConvertFieldType ([608184f](https://github.com/arisu-archive/arona-unflatd/commit/608184f6367c92252449d23acaa4203e19c4044e))
* ignore non flatdata struct ([c101993](https://github.com/arisu-archive/arona-unflatd/commit/c10199375e449203cb96845c9b05aa6b785dd199))
* incorrectly dump mismatched namespace ([ecb2f2e](https://github.com/arisu-archive/arona-unflatd/commit/ecb2f2e2d8b79b326b4c7306f441f1fe369fefb6))
* missing private field ([bd4fa40](https://github.com/arisu-archive/arona-unflatd/commit/bd4fa40b19f75f15277c7436316b8b472ae34c60))
* missing property parsing query ([1065f96](https://github.com/arisu-archive/arona-unflatd/commit/1065f96560019fccdf2764499e76572dc286e27e))
* skip schemas that match the command's namespace during processing ([28d4d96](https://github.com/arisu-archive/arona-unflatd/commit/28d4d96e7e7aa3b976d8d32bb94490608b55a476))
* update field handling to use pointer types in schema and visitor ([bdf4ad0](https://github.com/arisu-archive/arona-unflatd/commit/bdf4ad0438d06908444f75adb1f2b17fed4e4f59))
* use parameter order as fbs definition field order ([f21f91c](https://github.com/arisu-archive/arona-unflatd/commit/f21f91ce8b901adc644310bedbf0785d7b6815cd))
* use path-schema pairs instead of map to prevent filename collisions ([7c9fa1a](https://github.com/arisu-archive/arona-unflatd/commit/7c9fa1afe5093332f317fc2f0ef62a593e7934b6))
