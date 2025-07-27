# discord

# Usage

#### to use webhookURL

```go
if err := discord.SendMessage("message", webhookURL); err != nil {
    log.Printf("failed to send error message to Discord: %v", err)
    return
}
```

#### to use env

```go
os.Setenv("DISCORD_WEBHOOK", "https://example.com")
if err := discord.SendMessage("message"); err != nil {
    log.Printf("failed to send error message to Discord: %v", err)
    return
}
```

#### to use withMention

```go
if err := discord.SendMessageWithMention("message"); err != nil {
    log.Printf("failed to send error message to Discord: %v", err)
    return
}

// send message:
// @here
// message
```
