
# Eventbrite Sniper

Watches Eventbrite and obtains a ticket spot using firefox through Selenium.
Once a spot in the queue has been obtained, the link is sent by both SMS and email accordingly to the settings provided.




### Caveats

This relies on finding a "Register" button in the DOM.
It is very liable to break in the case of the CSS changing.




### Notifications

Twilio is used to send text messages.
Sendgrid is used to send emails.





## Running It


First run Selenium:

```bash
java -jar selenium-server-standalone-2.45.0.jar
```


Then run the sniper.

```bash
go build
source your_settings_file.sh
./eventbrite_sniper
```


### Settings

The following settings are needed as environment variables.
You can fill out the provided `secrets.sh` file with your credentials.

```
TWILIO_ACCOUNT    # acct number
TWILIO_TOKEN      # acct token
TWILIO_NUMBER     # your twilio number
SENDGRID_USER     # sendgrid username
SENDGRID_PASS     # sendgrid password
RECIPIENT_EMAIL   # email to send link to
RECIPIENT_NUMBER  # number to text
```


