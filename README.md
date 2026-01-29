# Telnyx CLI

The official CLI for the Telnyx REST API.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/telnyx-cli/cmd/telnyx@latest'
```

### Running Locally

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
telnyx [resource] [command] [flags]
```

```sh
telnyx calls dial \
  --connection-id conn12345 \
  --from +15557654321 \
  --to +15551234567 \
  --answering-machine-detection detect \
  --answering-machine-detection-config '{after_greeting_silence_millis: 1000, between_words_silence_millis: 1000, greeting_duration_millis: 1000, greeting_silence_duration_millis: 2000, greeting_total_analysis_time_millis: 50000, initial_silence_millis: 1000, maximum_number_of_words: 1000, maximum_word_length_millis: 2000, silence_threshold: 512, total_analysis_time_millis: 5000}' \
  --audio-url http://www.example.com/sounds/greeting.wav \
  --billing-group-id f5586561-8ff0-4291-a0ac-84fe544797bd \
  --bridge-intent \
  --bridge-on-answer \
  --client-state aGF2ZSBhIG5pY2UgZGF5ID1d \
  --command-id 891510ac-f3e4-11e8-af5b-de00688a4901 \
  --conference-config '{id: 0ccc7b54-4df3-4bca-a65a-3da1ecc777f0, beep_enabled: on_exit, conference_name: telnyx-conference, early_media: false, end_conference_on_exit: true, hold: true, hold_audio_url: http://example.com/message.wav, hold_media_name: my_media_uploaded_to_media_storage_api, mute: true, soft_end_conference_on_exit: true, start_conference_on_create: false, start_conference_on_enter: true, supervisor_role: whisper, whisper_call_control_ids: [v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ, v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw]}' \
  --custom-header '{name: head_1, value: val_1}' \
  --custom-header '{name: head_2, value: val_2}' \
  --dialogflow-config '{analyze_sentiment: false, partial_automated_agent_reply: false}' \
  --enable-dialogflow \
  --from-display-name 'Company Name' \
  --link-to 'ilditnZK_eVysupV21KzmzN_sM29ygfauQojpm4BgFtfX5hXAcjotg==' \
  --media-encryption SRTP \
  --media-name my_media_uploaded_to_media_storage_api \
  --park-after-unbridge self \
  --preferred-codecs G722,PCMU,PCMA,G729,OPUS,VP8,H264 \
  --record record-from-answer \
  --record-channels single \
  --record-custom-file-name my_recording_file_name \
  --record-format wav \
  --record-max-length 1000 \
  --record-timeout-secs 100 \
  --record-track outbound \
  --record-trim trim-silence \
  --send-silence-when-idle \
  --sip-auth-password password \
  --sip-auth-username username \
  --sip-header "{name: User-to-User, value: '12345'}" \
  --sip-region Canada \
  --sip-transport-protocol TLS \
  --sound-modifications '{octaves: 0.1, pitch: 0.8, semitone: -2, track: both}' \
  --stream-bidirectional-codec G722 \
  --stream-bidirectional-mode rtp \
  --stream-bidirectional-sampling-rate 16000 \
  --stream-bidirectional-target-legs both \
  --stream-codec PCMA \
  --stream-establish-before-call-originate \
  --stream-track both_tracks \
  --stream-url wss://www.example.com/websocket \
  --supervise-call-control-id v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg \
  --supervisor-role barge \
  --time-limit-secs 60 \
  --timeout-secs 60 \
  --transcription \
  --transcription-config '{client_state: aGF2ZSBhIG5pY2UgZGF5ID1d, command_id: 891510ac-f3e4-11e8-af5b-de00688a4901, transcription_engine: Google, transcription_engine_config: {enable_speaker_diarization: true, hints: [string], interim_results: true, language: en, max_speaker_count: 4, min_speaker_count: 4, model: latest_long, profanity_filter: true, speech_context: [{boost: 1, phrases: [string]}], transcription_engine: Google, use_enhanced: true}, transcription_tracks: both}' \
  --webhook-url https://your-webhook.url/events \
  --webhook-url-method POST
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
