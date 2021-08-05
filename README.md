This is a demo using Google Cloud Text-to-Speech to synthesize strings read from a Google Sheets spreadsheet and write the resulting audio to MP3 files in Google Drive. It has rudimentary change tracking using checksums. See the basic usage by running `go run ./tts.go --help`. 

The associated Level Up episode, with hands-on walkthrough of usage, can be viewed on the [Google Cloud APAC YouTube Channel](https://youtu.be/aMRqPGD0_To)

You will need to create your own credentials.json and service-account.json files, as shown in the video above. If you have issues with authentication, delete any existing `token.json` file and try again.

**Warning**: This tutorial code is provided as-is, with no warranty or support. It can overwrite columns in sheets without confirmation, so be sure to try it out and make sure you know how it works before running it on data you care about! 

## License

Apache 2.0
