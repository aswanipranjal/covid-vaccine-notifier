Hi, I live in India at the time of writing this India is under it's second wave of covid and things are brutal.
Being a country of ~136.6 crores (1.4 B), resources are scares (beds, oxygen, proper medical care) and that means it'll be a while until this is over.

Because of our brilliant administration, me and lot of other techies are looking for solutions to get the vaccine and have started automating the notification/booking process as the last resort.

This is an attempt for me to get notified of slots available.

## Setup
This is supposed to be used as a cron job to get notifications on  whatsapp/sms via Twilio. Run it on a linux box in GCP (won't work on Digital ocean, havent tried AWS)  as a cron job (every 1-2 minutes since the solts get over very fast)

The script pings the cowin platform to check for hospitals whichhave vaccines available for people below 45 years of age.

You'll need to set up your accounton twillio and attach a verified phone number to get notifications.
[Follow this guide to set it up](https://www.twilio.com/docs/sms/send-messages#send-an-sms-with-twilios-api)

## Running the script
default flags:
```
s: start date (default today's date format: DD-MM-YYYY)
d: district id
i: number of days from start that you want run the script for
```

```
$ go run main.go -s 05-05-2021 -i 2 -d 294 
```
