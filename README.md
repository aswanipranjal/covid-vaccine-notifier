# So that I can get vaccinated

Hi, I live in India at the time of writing this and covid is around me.
This is an attempt for me to get notified of slots available.

## Setup
This is supposed to be used as a cron job to get notifications on 
whatsapp/sms via Twilio. 

The script pings the cowin platform to check for hospitals which
have vaccines available for people below 45 years of age.

You'll need to set up your account and attach a verified phone 
number to get notifications.
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