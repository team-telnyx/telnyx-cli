# Changelog

## 0.10.0 (2026-03-27)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.9.0...v0.10.0)

### Features

* add default description for enum CLI flags without an explicit description ([3984c4e](https://github.com/team-telnyx/telnyx-cli/commit/3984c4e4c8046eb4d7d9ec5c778342d9e3672495))
* Add Minimax provider support to Voice Designs and Voice Clones API spec ([541089f](https://github.com/team-telnyx/telnyx-cli/commit/541089f8f551a6848a0884f742918b8f19b9c540))
* **api:** manual updates ([ce17e31](https://github.com/team-telnyx/telnyx-cli/commit/ce17e31cf722f807f8e9448e351b4347281ea6e3))
* **api:** manual updates ([96cf66a](https://github.com/team-telnyx/telnyx-cli/commit/96cf66af4a01a5ee2a11da36f033a46fa0bfd5a2))
* **api:** manual updates ([cbe7a45](https://github.com/team-telnyx/telnyx-cli/commit/cbe7a458eef881414a47693bb3ec9f96ad989c28))
* **api:** manual updates ([333d745](https://github.com/team-telnyx/telnyx-cli/commit/333d745c5e9db537609369073a3aa0a2d9ab84a4))
* **api:** Merge pull request [#30](https://github.com/team-telnyx/telnyx-cli/issues/30) from stainless-sdks/fix-schemaUnionDiscriminatorMissing ([473ae31](https://github.com/team-telnyx/telnyx-cli/commit/473ae316481cd10bf367b5f5d5a8dc910495c1cd))
* New tools api ([53bbc57](https://github.com/team-telnyx/telnyx-cli/commit/53bbc57b833b0ebe491d95006feeb8a630e1d79f))
* set CLI flag constant values automatically where `x-stainless-const` is set ([bdffc96](https://github.com/team-telnyx/telnyx-cli/commit/bdffc96a0e197d00b951c269f5a57affbd4c991b))
* TELAPPS-5685: Add store_fields_as_variables to WebhookToolParams ([5486544](https://github.com/team-telnyx/telnyx-cli/commit/548654484ae8321c06bd823ca2e69697d97aacd0))
* **wireless:** add traffic policy profiles endpoints to OpenAPI spec ([72a9cc3](https://github.com/team-telnyx/telnyx-cli/commit/72a9cc3a8ea4e66ca445ee870f1b9fd4881f9606))


### Bug Fixes

* cli no longer hangs when stdin is attached to a pipe with empty input ([df986e3](https://github.com/team-telnyx/telnyx-cli/commit/df986e320b6369da40abf118a94657362176518d))
* **cli:** update go_sdk_version to v4.55.0 ([c171aa5](https://github.com/team-telnyx/telnyx-cli/commit/c171aa511fdecb16d36876708fb1ce1d5815fbf4))


### Chores

* **ci:** skip lint on metadata-only changes ([d1cdb62](https://github.com/team-telnyx/telnyx-cli/commit/d1cdb629356dddb5142ebe84f1dfdc2a9191ffd4))
* **internal:** update gitignore ([6fe564c](https://github.com/team-telnyx/telnyx-cli/commit/6fe564c2080e8987788fb81e6d540c4fe0296d6a))
* omit full usage information when missing required CLI parameters ([83bd447](https://github.com/team-telnyx/telnyx-cli/commit/83bd447d7ef30e59fddfa89bbf660ae3352fcf88))


### Documentation

* **branded-calling:** add Number Reputation API specs ([18ce61e](https://github.com/team-telnyx/telnyx-cli/commit/18ce61e40942615778c80aae22bc0d543a94ff9e))
* fix voice settings available voices link ([2be640f](https://github.com/team-telnyx/telnyx-cli/commit/2be640f4c61a53798fa194c820b9e27d5dfb6319))
* WhatsApp template components schema ([b92f5ed](https://github.com/team-telnyx/telnyx-cli/commit/b92f5ed6f05917352af8bc80eff59980a3aec528))

## 0.9.0 (2026-03-20)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.8.0...v0.9.0)

### Features

* TELAPPS-5668: Add call.cost webhook event documentation ([da5c801](https://github.com/team-telnyx/telnyx-cli/commit/da5c8013374f59bb7d7f0b324baf40f5f43cc4bc))


### Documentation

* **tts:** Add Telnyx.Ultra model documentation ([0b07c86](https://github.com/team-telnyx/telnyx-cli/commit/0b07c8641d8cc4a3abbcb80e31cd16e3f1e5caeb))

## 0.8.0 (2026-03-18)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.7.0...v0.8.0)

### Features

* add ai_assistant_join call control command OpenAPI spec ([58b7b54](https://github.com/team-telnyx/telnyx-cli/commit/58b7b548e6211dc4d3706672236238ec34f1764f))
* add message_history, send_message_history_updates, participants to AIAssistantStartRequest ([abed3f5](https://github.com/team-telnyx/telnyx-cli/commit/abed3f5a2d8a395b7de4813b3a2ca36296254206))
* add public x402 payment endpoints to external specs ([ac28064](https://github.com/team-telnyx/telnyx-cli/commit/ac2806411a686b191ab02d7b583cfd346e24cf3b))
* AI-2131: Add expressive_mode boolean to VoiceSettings ([4e4708d](https://github.com/team-telnyx/telnyx-cli/commit/4e4708d9bb0da6f724df514e10cd8b14865c1254))
* **api:** manual updates ([7618cf5](https://github.com/team-telnyx/telnyx-cli/commit/7618cf5fcd3a15bc95759f332abe50034677701e))
* **api:** manual updates ([2deb7c4](https://github.com/team-telnyx/telnyx-cli/commit/2deb7c4960b9d370a01ac56289ac4bfea583cabe))
* port-4690: fix LOA configuration preview path (singular → plural) ([2ee8e0f](https://github.com/team-telnyx/telnyx-cli/commit/2ee8e0fcdfb2b123d8b5102804d69014ed550eb5))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([fdf002a](https://github.com/team-telnyx/telnyx-cli/commit/fdf002a1884a8128f4252b69e33ddc5ac3647a8f))
* better support passing client args in any position ([6ab5f27](https://github.com/team-telnyx/telnyx-cli/commit/6ab5f277f2b681cdebeac00a174b72e2bee00c29))
* **call-recordings:** align OpenAPI spec with implementation ([f12b073](https://github.com/team-telnyx/telnyx-cli/commit/f12b0732903f59f48b559d4009722f91420de294))
* **cli:** remove need to manually pin version of Go SDK ([4c94b4e](https://github.com/team-telnyx/telnyx-cli/commit/4c94b4eb34023101ad046ac13c54815b43d24e7f))
* **codegen:** correct type mismatches and method names ([e8abe2d](https://github.com/team-telnyx/telnyx-cli/commit/e8abe2d31d321cb1235f038bb0f67d5eac3ffeaf))
* improve linking behavior when developing on a branch not in the Go SDK ([e01f881](https://github.com/team-telnyx/telnyx-cli/commit/e01f881d5664c8c5b4ddeb15f8f8af6e64fbec4c))
* improved workflow for developing on branches ([20ed1bc](https://github.com/team-telnyx/telnyx-cli/commit/20ed1bcb504fe7b022b29176cebf6769faf1716c))
* no longer require an API key when building on production repos ([2d78dd2](https://github.com/team-telnyx/telnyx-cli/commit/2d78dd27d017f8d7276960dd85b13dad95e03856))
* only set client options when the corresponding CLI flag or env var is explicitly set ([12701b6](https://github.com/team-telnyx/telnyx-cli/commit/12701b68896497ff10cee0dc9b458dc46f0d1a01))


### Chores

* **deps:** bump telnyx-go to v4.46.0 ([96e7bd3](https://github.com/team-telnyx/telnyx-cli/commit/96e7bd398a9661bf144b15e5e4951bb306e8650a))
* **internal:** codegen related update ([b7f6323](https://github.com/team-telnyx/telnyx-cli/commit/b7f6323114f0116b18b061079fbd4d695c31c77f))
* **internal:** codegen related update ([58d0f2b](https://github.com/team-telnyx/telnyx-cli/commit/58d0f2b282831466d622d69f943dd7fd71d7c6ef))
* **internal:** tweak CI branches ([ea3aad5](https://github.com/team-telnyx/telnyx-cli/commit/ea3aad5b69f4f89e6567fca687736bdc40b921cb))

## 0.7.0 (2026-03-12)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.6.0...v0.7.0)

### Features

* add `--max-items` flag for paginated/streaming endpoints ([bd3994b](https://github.com/team-telnyx/telnyx-cli/commit/bd3994be31287298c33cb6a873ce0c70d4a6542c))
* Add enable_thinking parameter to chat completions API ([021b800](https://github.com/team-telnyx/telnyx-cli/commit/021b80038e113da5fe7b3624df3d642cd43865d9))
* Add Voice Designs and Voice Clones API specification ([f896bb5](https://github.com/team-telnyx/telnyx-cli/commit/f896bb5749bdb5b9c222f1015430595e3b7a378f))
* **api:** manual updates ([58fb3be](https://github.com/team-telnyx/telnyx-cli/commit/58fb3beeb7a44cdeeb9cc6544bbe927ed8d406d2))
* **api:** manual updates ([ea5e255](https://github.com/team-telnyx/telnyx-cli/commit/ea5e25534f5b6e1b5067340253011b01e1748f2b))
* **api:** manual updates ([786dcca](https://github.com/team-telnyx/telnyx-cli/commit/786dccaa28cc9b1543c48d8274cd63b2a388718a))
* **api:** manual updates ([0572d09](https://github.com/team-telnyx/telnyx-cli/commit/0572d09fcb828e004ffd81275b040f61c49debc7))
* **api:** Merge pull request [#27](https://github.com/team-telnyx/telnyx-cli/issues/27) from stainless-sdks/fix/whatsapp-message-templates-conflict ([89030b7](https://github.com/team-telnyx/telnyx-cli/commit/89030b7ca216936df0e176ffbf6ad909e0cede38))
* Assistant tags ([b0d6cbd](https://github.com/team-telnyx/telnyx-cli/commit/b0d6cbde68a5b86236cf406df8025ecc4676ea54))
* CW-2881 publish wireless VoLTE docs to prod ([63b959a](https://github.com/team-telnyx/telnyx-cli/commit/63b959a9ce479774d71ab342878fd1ebb9656e7c))
* **stt:** add WebSocket event schemas for Stainless SDK generation ([a9eeede](https://github.com/team-telnyx/telnyx-cli/commit/a9eeede2238287bd16de3ace3241043430f076d6))
* support passing required body params through pipes ([882ac2e](https://github.com/team-telnyx/telnyx-cli/commit/882ac2e8e347d31fd44fbe6e64712c8340ab47e9))
* TELAPPS-ENGDESK-49737 Add prevent_double_bridge param to dial ([5d442cb](https://github.com/team-telnyx/telnyx-cli/commit/5d442cb78662343e9d3532bee48731354f6f61d7))


### Bug Fixes

* add missing client parameter flags to test cases ([f9b4d80](https://github.com/team-telnyx/telnyx-cli/commit/f9b4d80accac21463b8fe3e7e7001508d7d4324a))
* add missing example parameters for test cases ([e99ea97](https://github.com/team-telnyx/telnyx-cli/commit/e99ea971ad64c60500134c12f69d9eaf0934866e))
* add missing vertical enum values for 10DLC brand creation (ENGDESK-49040) ([84ae9a0](https://github.com/team-telnyx/telnyx-cli/commit/84ae9a0511193999938289c0c90eb8265512748a))
* fix for encoding arrays with `any` type items ([95fa3bd](https://github.com/team-telnyx/telnyx-cli/commit/95fa3bdfb8e1ec5aed1e467122820f69fbbb3124))
* fix for test cases with newlines in YAML and better error reporting ([8ca922a](https://github.com/team-telnyx/telnyx-cli/commit/8ca922a94ac0ba0f13c321f6d468b40c49a9fe5b))
* rename whatsapp.message_templates to whatsapp.templates to avoid conflict ([89030b7](https://github.com/team-telnyx/telnyx-cli/commit/89030b7ca216936df0e176ffbf6ad909e0cede38))
* update wait_seconds example to 0.5 ([5364833](https://github.com/team-telnyx/telnyx-cli/commit/53648337ba3c65fb1a2f34eb8bd1bd546c6fa1bb))


### Reverts

* restore stainless.yml changes removed in 1de6067 ([a0bc2eb](https://github.com/team-telnyx/telnyx-cli/commit/a0bc2ebcc09e2541668c8e3ee7f0d6f26c0bc3d7))


### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([db57625](https://github.com/team-telnyx/telnyx-cli/commit/db5762538c2dac26c7dabff0145b7812a32a8a1c))
* **deps:** bump telnyx-go to v4.45.0 ([f1d102a](https://github.com/team-telnyx/telnyx-cli/commit/f1d102a91308a15168aa54b8f873e53afe1f9bb2))
* update example date in usage-reports ([68b47ec](https://github.com/team-telnyx/telnyx-cli/commit/68b47ec73184456dffc28a04911f98fbea1d6ac7))

## 0.6.0 (2026-03-04)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.5.0...v0.6.0)

### Features

* [TDA-6425] Add Session Analysis API spec to public docs ([ba91a39](https://github.com/team-telnyx/telnyx-cli/commit/ba91a3993cd01974d1ddbc36e41a3ce729eab1aa))
* **api:** manual updates ([a45318c](https://github.com/team-telnyx/telnyx-cli/commit/a45318cd42315b4a1fa863b518ab41d07b05fe86))
* **api:** manual updates ([ba4a6a5](https://github.com/team-telnyx/telnyx-cli/commit/ba4a6a51b47f2dc980d52f69a040ccfe909f6378))

## 0.5.0 (2026-03-03)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.4.0...v0.5.0)

### Features

* add support for file downloads from binary response endpoints ([9bf74c0](https://github.com/team-telnyx/telnyx-cli/commit/9bf74c0f81e2ef4896e21c2dd53cde1703822f49))
* **api:** manual updates ([a04a5fa](https://github.com/team-telnyx/telnyx-cli/commit/a04a5fa9c65a748f94aa3e5583cca365f9fd182e))
* **api:** manual updates ([58fe50e](https://github.com/team-telnyx/telnyx-cli/commit/58fe50eb2c39102a968412326de82cceaca557b2))
* **api:** manual updates ([7842380](https://github.com/team-telnyx/telnyx-cli/commit/78423802757d9956522ad5aa32c25abaa3f2f931))
* **api:** manual updates ([b51de06](https://github.com/team-telnyx/telnyx-cli/commit/b51de06bac644ed850418240a65e68b8199138fc))
* improved documentation and flags for client options ([4c81556](https://github.com/team-telnyx/telnyx-cli/commit/4c815564351e65c96476b537e6f4e59dfb02e09c))
* Merge TTS file-based spec into text-to-speech.json ([35f28ce](https://github.com/team-telnyx/telnyx-cli/commit/35f28ce1edbc7f82af9e3d045dabcd6593247198))


### Bug Fixes

* avoid printing usage errors twice ([09be78b](https://github.com/team-telnyx/telnyx-cli/commit/09be78b0a27b111da194288eacfa1f69d93a6a5d))


### Chores

* **ci:** shorten artifact link ([abec876](https://github.com/team-telnyx/telnyx-cli/commit/abec8765c8b48bf661b5b0601bfb11c2a733034b))

## 0.4.0 (2026-02-27)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.3.0...v0.4.0)

### Features

* Add TTS file-based endpoint spec ([f8add93](https://github.com/team-telnyx/telnyx-cli/commit/f8add93c7a8c7af1c58b3ff2b3ebc7065704cb1a))

## 0.3.0 (2026-02-26)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.2.1...v0.3.0)

### Features

* Add missing TTS voice settings schemas and update voice descriptions ([f3f285e](https://github.com/team-telnyx/telnyx-cli/commit/f3f285eaca192811a6af90702d9c042b767e49dd))
* Add text-to-speech WebSocket streaming OpenAPI spec ([8eda057](https://github.com/team-telnyx/telnyx-cli/commit/8eda05748827bfdf4265f6dd038cd475079238aa))
* **api:** manual updates ([8cdcc53](https://github.com/team-telnyx/telnyx-cli/commit/8cdcc5311cbba81d6cc0404170414d4dd4985089))
* **api:** manual updates ([792e45d](https://github.com/team-telnyx/telnyx-cli/commit/792e45dc48a15106fec9eca81a135ce8bd0fd59d))
* PORTAL-5923: Add stored_payment_transactions endpoint to OpenAPI docs ([a0244ae](https://github.com/team-telnyx/telnyx-cli/commit/a0244ae16a737ed136ed4bccaf7e182b4c840703))
* TELAPPS Add prevent_double_bridge to bridge command ([e9cd493](https://github.com/team-telnyx/telnyx-cli/commit/e9cd493c1bcb67fe4f7856179bb6a192c3d8a66a))
* TELAPPS-ENGDESK-48951 add channels to conf record start ([9977d90](https://github.com/team-telnyx/telnyx-cli/commit/9977d901462b8c127ba0bb1bc348acb6e4e67315))


### Bug Fixes

* more gracefully handle empty stdin input ([b31dd84](https://github.com/team-telnyx/telnyx-cli/commit/b31dd84d5d98cf50a60c68c49e0861428a14abc8))


### Chores

* bring back other changes ([3b85ea7](https://github.com/team-telnyx/telnyx-cli/commit/3b85ea7a8084f03f3d8cc71b035087d2bc6343dd))
* **cli:** bump go sdk version ([8d04de4](https://github.com/team-telnyx/telnyx-cli/commit/8d04de4df7c7be888ab3d4a9bb03f4d41738594c))
* zip READMEs as part of build artifact ([bfea13f](https://github.com/team-telnyx/telnyx-cli/commit/bfea13fc0b3e5f3ff8a51b8f18b0240c3a845f13))


### Documentation

* **call-control:** Add missing params to hangup, bridge, answer ([0ad8643](https://github.com/team-telnyx/telnyx-cli/commit/0ad8643b35eece56c8bcfa8279300485976c20de))
* **call-control:** Add queue CRUD endpoints ([598e40f](https://github.com/team-telnyx/telnyx-cli/commit/598e40fcab223d1c366ffbd82b8a483561c8e950))

## 0.2.1 (2026-02-24)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/team-telnyx/telnyx-cli/compare/v0.2.0...v0.2.1)

### Bug Fixes

* pin formatting for headers to always use repeat/dot formats ([540d172](https://github.com/team-telnyx/telnyx-cli/commit/540d172963da808ed4770f833454c64add9b1e68))

## 0.2.0 (2026-02-24)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.1.0...v0.2.0)

### Features

* Add dynamic_variables field to scheduled event schemas ([f36ef47](https://github.com/team-telnyx/telnyx-cli/commit/f36ef47b93715b4c424b39ce816839a4c3525ba1))
* Add Label parameter to Dial Conference Participant endpoint ([1f9b454](https://github.com/team-telnyx/telnyx-cli/commit/1f9b454ff2d11f635fd93c6fd11e4dc3ea8d73ec))
* Add Minimax and Resemble voice providers for speak commands ([a8f32cd](https://github.com/team-telnyx/telnyx-cli/commit/a8f32cda4fe7ec72c49e2a2f42767601dae3a616))
* Add OpenAI-compatible embeddings API spec ([12c7ec8](https://github.com/team-telnyx/telnyx-cli/commit/12c7ec895a81ee8a68f4fdf6fac1f12ead123f23))
* Add smart encoding fields to messaging API spec ([38e99db](https://github.com/team-telnyx/telnyx-cli/commit/38e99dbf3aef18bc9238dadd7551e972866a8ab5))
* AI-2078 Feature: API endpoint to disable AI assistant mid-conversation ([25e7e1d](https://github.com/team-telnyx/telnyx-cli/commit/25e7e1d698d9fa2dc0c7943325abf881afcf0d33))
* AI-2086: Add AI Missions endpoints to inference spec ([f237c1c](https://github.com/team-telnyx/telnyx-cli/commit/f237c1ca18c586988d236cba8b33adb7bf736510))
* AI-2093: Add language_boost to MiniMax voice settings ([30a20a2](https://github.com/team-telnyx/telnyx-cli/commit/30a20a25ff75f6c959d5d33edda93a8b1d2265ce))
* **api:** manual updates ([7e81f77](https://github.com/team-telnyx/telnyx-cli/commit/7e81f7752d891dd9ecf4b818932d7da20a541708))
* **api:** manual updates ([28cf507](https://github.com/team-telnyx/telnyx-cli/commit/28cf507080c906f1d1d5530efbd12fe840d28ba8))
* **api:** manual updates ([a9d3a59](https://github.com/team-telnyx/telnyx-cli/commit/a9d3a5992c3cfb81c1c78affc09b1fb6faacae0e))
* **api:** manual updates ([700b54f](https://github.com/team-telnyx/telnyx-cli/commit/700b54fab2c69a7471a496e7f689962a4d7a746d))
* **api:** manual updates ([83d4f70](https://github.com/team-telnyx/telnyx-cli/commit/83d4f703af63e1d84a2429d0134e3be5fee311b1))
* **api:** manual updates ([0eb3575](https://github.com/team-telnyx/telnyx-cli/commit/0eb3575db43072836fcc27dae4efde35cf30d272))
* **api:** manual updates ([5b60f77](https://github.com/team-telnyx/telnyx-cli/commit/5b60f779f626fa1c0d9075835ffc4d6320fcb7f1))
* **api:** manual updates ([7d6e087](https://github.com/team-telnyx/telnyx-cli/commit/7d6e087e4edc927da0d818dd7a05dc51c0ef2d7a))
* **api:** manual updates ([bb18e78](https://github.com/team-telnyx/telnyx-cli/commit/bb18e7879d3ff4d9503fb4a6f16f61ff7d5377fd))
* **api:** manual updates ([f8faa2f](https://github.com/team-telnyx/telnyx-cli/commit/f8faa2f8a764f71e54368ccae73a06bfa39c5f42))
* **api:** manual updates ([e4c5d3d](https://github.com/team-telnyx/telnyx-cli/commit/e4c5d3d9612cdc5de7318ea8a1f93b8ca0764f43))
* ENGDESK-49554: Add quail_voice_focus to noise suppression engine enums ([414b057](https://github.com/team-telnyx/telnyx-cli/commit/414b057ad5d787fa4bf8edaae094de1d20708b70))
* fix-stainless-sdk-timeout ([a5a03a8](https://github.com/team-telnyx/telnyx-cli/commit/a5a03a8ebcf1f214ed5c7c39f5de5583c91b6a9b))
* improved support for passing files for `any`-typed arguments ([eae5e29](https://github.com/team-telnyx/telnyx-cli/commit/eae5e295423c0af62b29a2e00d9902223158b8a7))
* Revert "fix emergency settings -schema" ([c19077a](https://github.com/team-telnyx/telnyx-cli/commit/c19077a34da37bc0680b244c8615c06ca39b6061))
* TELAPPS Add ApplicationSid param ([de31a0a](https://github.com/team-telnyx/telnyx-cli/commit/de31a0a5a1a7ceb01d0d10fef2aeda322e5aeac9))
* TELAPPS Add interim_results to deepgram config ([2a71ce8](https://github.com/team-telnyx/telnyx-cli/commit/2a71ce87a8a8fc96d65a0e2bd6569385e0062253))


### Bug Fixes

* fix for generated tests for some array flags ([ec0818d](https://github.com/team-telnyx/telnyx-cli/commit/ec0818d9713495ce282fe25969a37ca4bb51f47d))
* fix for when terminal width is not available ([ef7727f](https://github.com/team-telnyx/telnyx-cli/commit/ef7727fbe4e2625e42fd7d78b8b397f88b5f6971))
* move unsupported string formats to x-format ([1fffb82](https://github.com/team-telnyx/telnyx-cli/commit/1fffb82d2112a8403f58b7a5dc56ff97e6637338))
* OAS drift — 10dlc.json (messaging-campaign-registry) ([04c3211](https://github.com/team-telnyx/telnyx-cli/commit/04c3211b7f03e9b8267944454d6f201c75d87015))
* OAS drift — messaging.json (messaging-settings + messaging-outbound) ([745863b](https://github.com/team-telnyx/telnyx-cli/commit/745863b25c8e999a622232c1abdbba7cb3232320))
* OAS drift — toll-free-verification.json (messaging-tf-verify) ([9c33f82](https://github.com/team-telnyx/telnyx-cli/commit/9c33f825b002436ee6f82cad2a2de73bfbbb508b))
* OAS drift — verify.json (messaging-2fa) ([6f1da41](https://github.com/team-telnyx/telnyx-cli/commit/6f1da41f9263b3d8a177bbb35d84f9a17a09a2d6))
* preserve filename in content-disposition for file uploads ([99dc47f](https://github.com/team-telnyx/telnyx-cli/commit/99dc47f839983745e72f870fb1d1c0cf3e09841c))
* prevent tests from hanging on streaming/paginated endpoints ([6d971e0](https://github.com/team-telnyx/telnyx-cli/commit/6d971e0f824240a8ca40e7dc0abddf9a36f54f94))
* remove invalid discriminators from string enum schemas ([c8ae448](https://github.com/team-telnyx/telnyx-cli/commit/c8ae4486d459ab10790726a197f8e1029bd5f4d6))
* remove unused imports ([ccac068](https://github.com/team-telnyx/telnyx-cli/commit/ccac06829545a25376a4b28abd99aefbb4af2559))
* StringFormatNotSupported ([9d9799e](https://github.com/team-telnyx/telnyx-cli/commit/9d9799e797cbbc4c489a0580480e64b95f350d34))


### Chores

* **internal:** remove mock server code ([45c7d7f](https://github.com/team-telnyx/telnyx-cli/commit/45c7d7f98248534859165fa80fd662600a8bb843))
* update mock server docs ([bdba363](https://github.com/team-telnyx/telnyx-cli/commit/bdba363d15e54b16a955c87cd97e5c993c48eecd))


### Documentation

* add service_provider_login_url to authentication provider settings ([9c603ac](https://github.com/team-telnyx/telnyx-cli/commit/9c603ac57cf797a5bd0cd3ed4b34ff0dccc87fc0))
* **call-control:** Add missing conference endpoints ([c7d2377](https://github.com/team-telnyx/telnyx-cli/commit/c7d23777d60cad89b73e3dab9836fef45da308c7))
* **call-control:** Add missing parameters to call control endpoints ([a2c2c2e](https://github.com/team-telnyx/telnyx-cli/commit/a2c2c2e2d237f967d6078ae8786b281bdabd74ab))
* **call-scripting:** add Timeout and TimeLimit to InitiateTexmlCall ([054196d](https://github.com/team-telnyx/telnyx-cli/commit/054196d438a3dc65b26c05aa81953cd388bde603))

## 0.1.0 (2026-02-05)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/team-telnyx/telnyx-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([67928c9](https://github.com/team-telnyx/telnyx-cli/commit/67928c9f50f1ca0947f534386de839c971b1e678))


### Chores

* **config:** add new target ([cd918f1](https://github.com/team-telnyx/telnyx-cli/commit/cd918f14d14b1722ed3dfdeac85e778c7c4f4035))
* **internal:** codegen related update ([63de3f5](https://github.com/team-telnyx/telnyx-cli/commit/63de3f56f61d3fe01531d2886aad0e265530f9c3))
* **internal:** ignore stainless-internal artifacts ([77ccad9](https://github.com/team-telnyx/telnyx-cli/commit/77ccad95dd7a6450ddf3e064218f1c7a680891f6))
* prepare for npm publish ([#4](https://github.com/team-telnyx/telnyx-cli/issues/4)) ([8a6c36c](https://github.com/team-telnyx/telnyx-cli/commit/8a6c36c5b3ced0e4cda56dd75cd8ec7e00619ff8))
* sync repo ([814e54a](https://github.com/team-telnyx/telnyx-cli/commit/814e54ac71657d4b4ca7d0817d579a14e4ceeac2))
