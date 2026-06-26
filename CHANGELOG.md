# Changelog

## 0.18.0 (2026-06-26)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.17.0...v0.18.0)

### Features

* Add sip region to scheduled events ([c3d55fd](https://github.com/team-telnyx/telnyx-cli/commit/c3d55fd16b5c8d3888e30b398374edf57b8c2b7c))
* add transcriptions-search API spec (dev/external) ([146b2e5](https://github.com/team-telnyx/telnyx-cli/commit/146b2e5126f2dec4f4dfaf92b54861d2372d9a33))
* branded-calling: drop 'simplified'/'no enterprise_id' framing from DIR endpoints ([c8705d1](https://github.com/team-telnyx/telnyx-cli/commit/c8705d1852e4b6cc2a301a467a36acc008bc2f8b))
* branded-calling: sync PATCH /dir/{dir_id} spec with deployed service behavior ([dda24e9](https://github.com/team-telnyx/telnyx-cli/commit/dda24e93ea89f3b643af9a7e6b86c5c335c187c7))
* **branded-calling:** document LOA (DIR) + reputation remediation endpoints ([764ff5f](https://github.com/team-telnyx/telnyx-cli/commit/764ff5fdda0cbbe604676981aaa64692ffda9ad7))
* ENGDESK-51290: Add Aicoustic NS engines to portal backend ([42efb8b](https://github.com/team-telnyx/telnyx-cli/commit/42efb8bdff068f13aba049f45ee9bfb3a8e40056))
* Fix cross-product leakage: de-dup shared enterprise/ToS endpoints + neutralize descriptions ([729d827](https://github.com/team-telnyx/telnyx-cli/commit/729d827ebfae5199ec32c14e72c16296542d3167))
* NUM-6470: Add branded-calling-v2 OpenAPI specs (external + internal) ([1e72096](https://github.com/team-telnyx/telnyx-cli/commit/1e720962d7fb5f9a25885a89be3f060ce38e69b1))
* **spec:** backfill parameter descriptions to 100% coverage ([447a2c9](https://github.com/team-telnyx/telnyx-cli/commit/447a2c9b83361e263029b915d54c4b0d129b0e0c))
* specs: add GET /terms_of_service/info (branded-calling + number-reputation) ([55dd1d4](https://github.com/team-telnyx/telnyx-cli/commit/55dd1d4b5dcf9c45848d32246a6fa8b546ec7a93))
* specs: drop legacy plain filter params (keep filter[] canonical) on list endpoints ([0e03c64](https://github.com/team-telnyx/telnyx-cli/commit/0e03c6487279a716f21fe849e5840c3c9c7ab19a))
* specs: remove em dashes from branded-calling + number-reputation descriptions ([51fe586](https://github.com/team-telnyx/telnyx-cli/commit/51fe58672c7d7434e425b0f472f1c2b56137f944))
* TELAPPS-719: add Inworld inworld-tts-2 model + delivery_mode to call-control voice settings ([4c0a9bc](https://github.com/team-telnyx/telnyx-cli/commit/4c0a9bc79e2ba2b42f5952e65e3622843a2480ed))
* **transcriptions-search:** hide record_type, remove document_id, scrub internal references ([0196578](https://github.com/team-telnyx/telnyx-cli/commit/0196578ba02913154cfc569a5ceb16b4bc061121))


### Bug Fixes

* **number-reputation:** remediation contact_email is optional ([f6cc43c](https://github.com/team-telnyx/telnyx-cli/commit/f6cc43cbaf24132f3c414dc1c4187d289dbe6f5c))
* restore stainless.yml — only remove production_repo for Python ([2c2e64d](https://github.com/team-telnyx/telnyx-cli/commit/2c2e64d9c0dbf07a44ed5b0a43c1c2212329457b))


### Chores

* add release-please workflow + fix config for STLC cutover ([#29](https://github.com/team-telnyx/telnyx-cli/issues/29)) ([a50a4e2](https://github.com/team-telnyx/telnyx-cli/commit/a50a4e27bfb7a5b7d3fc23082820aeb608c742d1))
* **internal:** codegen related update ([593ad8e](https://github.com/team-telnyx/telnyx-cli/commit/593ad8ec4fa0ec035026625dc431804a71a9ecf1))
* **internal:** codegen related update ([07781f4](https://github.com/team-telnyx/telnyx-cli/commit/07781f4407346d6d7b28a9bf2c50294bde7ef324))
* promote from staging 5c6693b ([0c930fa](https://github.com/team-telnyx/telnyx-cli/commit/0c930facd8c950ac1f0f87c66c87832c3d5e0e95))
* promote from staging 6e95420 ([bc888d1](https://github.com/team-telnyx/telnyx-cli/commit/bc888d1701663f61a21c4b0a3e9fbeb5ad5fe0cd))
* promote from staging b63e558 ([cb0e871](https://github.com/team-telnyx/telnyx-cli/commit/cb0e87194af593e4e588f5dfe220aaed0f30f282))
* promote from staging eeb1233 ([53be537](https://github.com/team-telnyx/telnyx-cli/commit/53be537997bfbbf7311bc8323f6fc07c3c8834a5))
* update example date in audit-logs for pipeline sync ([ec95d54](https://github.com/team-telnyx/telnyx-cli/commit/ec95d541c92f29764ccd51a571654568f5095b3f))
* update example date in audit-logs for SaaS rebuild ([766f3b3](https://github.com/team-telnyx/telnyx-cli/commit/766f3b37eaa951d8525ddceb4ddf1ff3899685ec))

## [0.23.3](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.23.2...v0.23.3) (2026-06-26)


### Chores

* remove staging release-please.yml ([229cf3a](https://github.com/team-telnyx/telnyx-cli-staging/commit/229cf3a2732a7809219fddf45d593782bfe97436))

## [0.23.2](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.23.1...v0.23.2) (2026-06-26)


### Chores

* add promote-to-prod workflow for STLC cutover ([eeb1233](https://github.com/team-telnyx/telnyx-cli-staging/commit/eeb123339ba78e449e0f6593e89e246f99407d0c))
* add promote-to-prod workflow for STLC cutover ([b8c2c9c](https://github.com/team-telnyx/telnyx-cli-staging/commit/b8c2c9cd5edc486ff8a26bf86606c2c3219c8c3c))

## [0.23.1](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.23.0...v0.23.1) (2026-06-17)


### Bug Fixes

* **ci:** use github.head_ref for branch resolution on PRs ([968d5c0](https://github.com/team-telnyx/telnyx-cli-staging/commit/968d5c08369984f7bb6aec69c9467e7860478e78))
* **ci:** use github.head_ref for branch resolution on PRs ([b07fa52](https://github.com/team-telnyx/telnyx-cli-staging/commit/b07fa5213bbf1e1ce5e73fdc961d8eb108930feb))

## [0.23.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.22.3...v0.23.0) (2026-06-16)


### Bug Fixes

* rename DIR method names to match published SDKs ([32adb02](https://github.com/team-telnyx/telnyx-cli-staging/commit/32adb02928b2d0d18ad15531c7fdbb5884612cff))
* restore go.mod replace directive for staging SDK ([b6eb4ba](https://github.com/team-telnyx/telnyx-cli-staging/commit/b6eb4baf27f191b8cbd4ca27bb28d30ba573bca7))


### Chores

* preserve repo-owned files not part of SDK generation ([2ae6bab](https://github.com/team-telnyx/telnyx-cli-staging/commit/2ae6bab855bea1f42fbf15f3c0a81f454d41f704))
* release cli 0.23.0 ([295b557](https://github.com/team-telnyx/telnyx-cli-staging/commit/295b557a62a3ee0ba10c55457b44d6171601ffb7))
* sync OpenAPI spec from e44bbbc ([e547311](https://github.com/team-telnyx/telnyx-cli-staging/commit/e54731189aa3418a65e581922affae1ff36f4e9d))

## [0.22.3](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.22.2...v0.22.3) (2026-06-10)


### Bug Fixes

* add missing go.sum entries for telnyx-go custom lib dependencies ([88fcb00](https://github.com/team-telnyx/telnyx-cli-staging/commit/88fcb00c71931f4f02a67f8801b73cfaa8f37787))
* add missing go.sum entries for telnyx-go custom lib dependencies ([fc05967](https://github.com/team-telnyx/telnyx-cli-staging/commit/fc059677d31d10e6f5ad4218cdfd1bb1fba1d660))

## [0.22.2](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.22.1...v0.22.2) (2026-06-09)


### Chores

* sync OpenAPI spec from c319cee ([25d4013](https://github.com/team-telnyx/telnyx-cli-staging/commit/25d401398cfa01251eb54472aa2529f1b669cce7))

## [0.22.1](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.22.0...v0.22.1) (2026-06-09)


### Bug Fixes

* **ci:** fix bash syntax error in release-please auto-merge step ([c8801fd](https://github.com/team-telnyx/telnyx-cli-staging/commit/c8801fd0786e7dff9fd37de73f01b73bb2fced8b))


### Refactors

* **ci:** align auto-merge step with canonical form in sibling staging repos ([b5a7548](https://github.com/team-telnyx/telnyx-cli-staging/commit/b5a75480e34fae22d8ed534c08ce51bb08662305))

## [0.22.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.21.1...v0.22.0) (2026-06-08)


### Chores

* preserve repo-owned files not part of SDK generation ([01441dd](https://github.com/team-telnyx/telnyx-cli-staging/commit/01441ddc20abdc81fe388ffffa6b0c36f8cc7360))
* preserve repo-owned files not part of SDK generation ([5e65e50](https://github.com/team-telnyx/telnyx-cli-staging/commit/5e65e50e14200aee54a9e9d4b318a0dde0f9cd98))
* release cli 0.22.0 ([104fbe9](https://github.com/team-telnyx/telnyx-cli-staging/commit/104fbe9ddd12fe6191f0cede5b023ba04179809b))
* release cli 0.22.0 ([9e0b640](https://github.com/team-telnyx/telnyx-cli-staging/commit/9e0b6406abf458853f31a774232489242e77cf91))
* sync OpenAPI spec from 6eae6a5 ([2d0f166](https://github.com/team-telnyx/telnyx-cli-staging/commit/2d0f1660be78e8e5d28f6ebc540e6c90b8395862))
* sync OpenAPI spec from a13d4b1 ([2c70cc2](https://github.com/team-telnyx/telnyx-cli-staging/commit/2c70cc28f1022ad41a15dfa4584717f50bbac54f))
* sync OpenAPI spec from aae7c19 ([ca80e41](https://github.com/team-telnyx/telnyx-cli-staging/commit/ca80e4160db40b912c4624f4f86d9310ffc94bf0))
* sync OpenAPI spec from aae7c19 ([af5777c](https://github.com/team-telnyx/telnyx-cli-staging/commit/af5777c7703b51ab29c92afbda39f6c3dd0b2c53))

## [0.21.1](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.21.0...v0.21.1) (2026-06-07)


### Bug Fixes

* add Go SDK auth + clone to publish-release.yml ([a8a9778](https://github.com/team-telnyx/telnyx-cli-staging/commit/a8a9778c1cba2f69764f633dc4d7a2be0d10d4b3))
* add Go SDK auth + clone to publish-release.yml ([a30435a](https://github.com/team-telnyx/telnyx-cli-staging/commit/a30435a0e98b9fd7d254b38c0cb07cc8cf6b2c5d))
* extract PR number from JSON output for auto-merge ([ad2ea47](https://github.com/team-telnyx/telnyx-cli-staging/commit/ad2ea470465c6d41684dbc8aff13a2a8eddc0930))
* extract PR number from JSON output for auto-merge ([02f5e3b](https://github.com/team-telnyx/telnyx-cli-staging/commit/02f5e3bcaadfb98f148fe2cca34d5d7f4bbaadd6))
* use release-please output directly for auto-merge (avoid race condition) ([1b676e4](https://github.com/team-telnyx/telnyx-cli-staging/commit/1b676e40c2e770341fe93a659e2f6ea2a68562b6))
* use release-please output directly for auto-merge (avoid race condition) ([1a6b29b](https://github.com/team-telnyx/telnyx-cli-staging/commit/1a6b29ba1c282558c8013ef8fd538ccad0faf0ca))

## [0.21.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.20.0...v0.21.0) (2026-06-07)


### Chores

* preserve repo-owned files not part of SDK generation ([22f502f](https://github.com/team-telnyx/telnyx-cli-staging/commit/22f502fa8c4d6dbe30fce4ae1f64f92498bdbf65))
* preserve repo-owned files not part of SDK generation ([b347d4a](https://github.com/team-telnyx/telnyx-cli-staging/commit/b347d4abc34161ca34147bad497b068942ca0bef))
* release cli 0.21.0 ([5917ba9](https://github.com/team-telnyx/telnyx-cli-staging/commit/5917ba976c45bb92d0a382ff360d8e2299387b93))
* release cli 0.21.0 ([4d177db](https://github.com/team-telnyx/telnyx-cli-staging/commit/4d177db027836a6c98cad8547faf5647ad8336dd))
* sync OpenAPI spec from 0193002 ([f09871e](https://github.com/team-telnyx/telnyx-cli-staging/commit/f09871e725e6c735521e8c552e2cd48665afdfe0))
* sync OpenAPI spec from 0193002 ([c8639d2](https://github.com/team-telnyx/telnyx-cli-staging/commit/c8639d2eca0b93ec8fe8a84b198615d0f76e65e4))
* sync OpenAPI spec from 8faa4be ([9e9115a](https://github.com/team-telnyx/telnyx-cli-staging/commit/9e9115a1b71199bd98f12136ed1b3c7b62bb9e1a))
* sync OpenAPI spec from 8faa4be ([4181610](https://github.com/team-telnyx/telnyx-cli-staging/commit/41816106e17dd210817afd6342f3ac2a8b589d46))

## [0.20.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.19.0...v0.20.0) (2026-06-07)


### Features

* enable GitHub auto-merge on release PRs ([230caa1](https://github.com/team-telnyx/telnyx-cli-staging/commit/230caa1ed406daaa9bce769fed911b2e479b18fb))
* enable GitHub auto-merge on release PRs ([da7701e](https://github.com/team-telnyx/telnyx-cli-staging/commit/da7701ef8c0a07dc98df11ef49bf707047189a9f))


### Bug Fixes

* correct auto-merge output name and PR search pattern ([d6ebb1f](https://github.com/team-telnyx/telnyx-cli-staging/commit/d6ebb1f88ddf4273ea90842b9cd893b7140ff661))
* correct auto-merge output name and PR search pattern ([d2deabf](https://github.com/team-telnyx/telnyx-cli-staging/commit/d2deabfd2a3330e7735dd5652cd2c9af5b37b5a6))
* restore ${{ }} expressions in release-please workflow ([2fc911f](https://github.com/team-telnyx/telnyx-cli-staging/commit/2fc911f31cc35d36f6942bd18c9689f78d7903e2))
* restore ${{ }} expressions in release-please workflow ([7c25175](https://github.com/team-telnyx/telnyx-cli-staging/commit/7c25175801469d22a4eeb4fb13e946936271974a))

## [0.19.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.18.0...v0.19.0) (2026-06-07)


### Features

* enable automerge for release PRs ([188aab2](https://github.com/team-telnyx/telnyx-cli-staging/commit/188aab2b855ce4a8de28b44d1bdf15e7e69cbd23))
* enable automerge for release PRs ([94bb69f](https://github.com/team-telnyx/telnyx-cli-staging/commit/94bb69fe6ac2165a79d2f86a3d6d6814434889b1))

## [0.18.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.17.1...v0.18.0) (2026-06-07)


### Bug Fixes

* correct YAML quoting in CI if condition ([c26fe5a](https://github.com/team-telnyx/telnyx-cli-staging/commit/c26fe5aa11c84f70939079bd8b4eb411a710ed87))
* remove text_to_speech retrieve_speech mapping ([e5e09a7](https://github.com/team-telnyx/telnyx-cli-staging/commit/e5e09a7255d1c0f96885d7972dfddb5ad347b83b))
* remove text_to_speech retrieve_speech mapping ([5b9ed8e](https://github.com/team-telnyx/telnyx-cli-staging/commit/5b9ed8ed534c6f317805aa784994ddede815e4c7))
* use PAT for release-please to trigger CI ([13330b5](https://github.com/team-telnyx/telnyx-cli-staging/commit/13330b5a1fb0b71b636fc4cb7973a36beee85238))
* use PAT for release-please to trigger CI ([e9c9f67](https://github.com/team-telnyx/telnyx-cli-staging/commit/e9c9f67e9e64472b5fafc72f62586bbfbb832ee4))


### Chores

* preserve repo-owned files not part of SDK generation ([371916a](https://github.com/team-telnyx/telnyx-cli-staging/commit/371916adf6c0ff9b3c686c434788f73cebb4c14a))
* preserve repo-owned files not part of SDK generation ([8fc3915](https://github.com/team-telnyx/telnyx-cli-staging/commit/8fc391554e8968e0b796fed29750af9ba91a7108))
* release cli 0.18.0 ([440b0d7](https://github.com/team-telnyx/telnyx-cli-staging/commit/440b0d7b5750d3a8084da2e05adb0078f6f6b351))
* release cli 0.18.0 ([467447f](https://github.com/team-telnyx/telnyx-cli-staging/commit/467447f0b28cc6c202461cb0e26e053138b62dd8))

## [0.17.1](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.17.0...v0.17.1) (2026-06-06)


### Bug Fixes

* run CI build/lint on internal PRs not just forks ([c523abc](https://github.com/team-telnyx/telnyx-cli-staging/commit/c523abcc3d8baa398d00a8420e094bfe29031c5c))

## [0.17.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.16.0...v0.17.0) (2026-06-06)


### Bug Fixes

* **cli:** reconcile stainless.yml method surface to canonical telnyx-go ([11ad49a](https://github.com/team-telnyx/telnyx-cli-staging/commit/11ad49ab2044d31dc8cbad4d11e80f9f2757826a))
* **go:** use model mapping for DirPhoneNumberStatus dedup ([8101742](https://github.com/team-telnyx/telnyx-cli-staging/commit/81017422d4180c9af071b5c348df47fe89f2267a))
* update custom code to use published Go SDK v4.76.0 API names ([d532315](https://github.com/team-telnyx/telnyx-cli-staging/commit/d532315a5f09201ca3b597fb0e8df53b31b5423a))
* update custom code to use published Go SDK v4.76.0 API names ([37017e4](https://github.com/team-telnyx/telnyx-cli-staging/commit/37017e4cc0d291686a4f4a39c2e173abed543dbf))


### Chores

* preserve repo-owned files not part of SDK generation ([ac3f4d4](https://github.com/team-telnyx/telnyx-cli-staging/commit/ac3f4d4293f2a0293c3f8e4798ba699821178eac))
* release cli 0.17.0 ([02aab60](https://github.com/team-telnyx/telnyx-cli-staging/commit/02aab60462e00a7cd70bc140ec413d6c104ad321))

## [0.16.0](https://github.com/team-telnyx/telnyx-cli-staging/compare/v0.15.0...v0.16.0) (2026-06-04)


### Chores

* release cli 0.16.0 ([c10ada3](https://github.com/team-telnyx/telnyx-cli-staging/commit/c10ada3adafefcf5c5b40121f42fb1cf05eaf87c))
* release cli 0.16.0 ([e4c9580](https://github.com/team-telnyx/telnyx-cli-staging/commit/e4c9580d502424a0a56b06ae8dd2ff39bfae6fd6))
