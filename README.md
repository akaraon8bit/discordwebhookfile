# Discord Webhook
Borrowed from https://github.com/gtuk/discordwebhook
This package provides a super simple interface to send discord messages and file upload through webhooks in golang.

### Installation
```
go get github.com/akaraon8bit/discordwebhook
```

### Example
Below is the most basic example on how to send a message.
For a more advanced message structure see the structs in types.go and https://birdie0.github.io/discord-webhooks-guide/discord_webhook.html

```
package main

import "github.com/akaraon8bit/discordwebhookfile"

func main() {
   var username = "BotUser"
   var content = "This is a test message"
   var url = "https://discord.com/api/webhooks/..."
   var file = []string{"data.txt"}

   message :=  discordwebhookfile.MessageFiles{
       Username: &username,
       Content: &content,
       Files: &file,
   }

   err :=  discordwebhookfile.SendMessage(url, message)
   if err != nil {
       log.Fatal(err)
   }
}
```

### TODO
* Tests
* Documentation
