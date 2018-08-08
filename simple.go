package main

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"fmt"
	"encoding/json"
)

func main()  {
	// Step 1: Create bundle
	bundle := &i18n.Bundle{DefaultLanguage: language.English}

	// Step 2: Create localizer for that bundle using one or more language tags
	loc := i18n.NewLocalizer(bundle, language.English.String())

	// Step 3: Define messages
	messages := &i18n.Message{
		ID: "Emails",
		Description: "The number of unread emails a user has",
		One: "{{.Name}} has {{.Count}} email.",
		Other: "{{.Name}} has {{.Count}} emails.",
	}

	// Step 3: Localize Messages
	messagesCount := 2
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: messages,
		TemplateData: map[string]interface{}{
			"Name": "Theo",
			"Count": messagesCount,
		},
		PluralCount: messagesCount,
	})

	fmt.Println(translation)

	// Define different delimiters
	messages = &i18n.Message{
		ID: "Notifications",
		Description: "The number of unread notifications a user has",
		One: "<<.Name>> has <<.Count>> notification.",
		Other: "<<.Name>> has <<.Count>> notifications.",
		LeftDelim:  "<<",
		RightDelim: ">>",
	}

	notificationsCount := 1
	translation = loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: messages,
		TemplateData: map[string]interface{}{
			"Name": "Nick",
			"Count": notificationsCount,
		},
		PluralCount: notificationsCount,
	})

	fmt.Println(translation)

	// Unmarshaling from files
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("en.json")
	bundle.MustLoadMessageFile("el.json")

	loc = i18n.NewLocalizer(bundle, "el")
	messagesCount = 10
	translation = loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "messages",
		TemplateData: map[string]interface{}{
			"Name": "Alex",
			"Count": messagesCount,
		},
		PluralCount: messagesCount,
	})

	fmt.Println(translation)
}
