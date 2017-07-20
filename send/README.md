---
sms send

- built with go
- command line or environment variable configuration
- requires account setup, this is assuming twilio api and pre
  configured phone number and account
- go get 
- make will build from the makefile

```
Usage of send:


Send an sms message from the command line or env variables.  Either
use the flag for text or env var. Override env or flag with command
line text. Args after the flags will be used as the text message. 

Skip flag/option arguments until last id marker "--" is seen 

  -account-sid string
    	usage: Account from twilio: https://www.twilio.com/console env var name(SMS_ACCOUNT_SID) : (string)
  -auth-token string
    	usage: Secret API tokey from twilio env var name(SMS_AUTH_TOKEN) : (string)
  -from-phone string
    	usage: Twilio allocated phone number, https://www.twilio.com/console/phone-numbers/getting-started env var name(SMS_FROM_PHONE) : (string)
  -text string
    	usage: Send text as SMS message env var name(SMS_TEXT) : (string)
  -to-phone string
    	usage: Send SMS message to this destination env var name(SMS_TO_PHONE) : (string)
```
