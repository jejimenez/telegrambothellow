package main

import (
    tb "gopkg.in/tucnak/telebot.v2"
    "log"
    "os"
)

func main() {
    f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    defer f.Close()

    log.SetOutput(f)
    log.Println("This is a test log entry")
    var (
        port      = os.Getenv("PORT")
        publicURL = os.Getenv("PUBLIC_URL")
        token     = os.Getenv("TOKEN")
    )

    webhook := &tb.Webhook{
        Listen:   ":" + port,
        Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
    }

    pref := tb.Settings{
        Token:  token,
        Poller: webhook,
    }

    b, err := tb.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }

    b.Handle("/hello", func(m *tb.Message) {
        b.Send(m.Sender, "Hi!")
    })

    b.Start()
}