package mail

import (
	"context"
	"fmt"
	"log"
	"net/mail"

	"github.com/can3p/gogo/sender"
	"github.com/can3p/gogo-cli-example/pkg/links"
	"github.com/can3p/gogo-cli-example/pkg/model/core"
	"github.com/pkg/errors"
)

func ConfirmSignup(ctx context.Context, s sender.Sender, user *core.User) error {
	if user.EmailConfirmSeed.String == "" {
		return errors.Errorf("cannot send confirm email for user with empty confirmation seed, user id = %s", user.ID)
	}

	link := links.AbsLink("confirm_signup", user.EmailConfirmSeed.String)
	to := user.Email

	mail := &sender.Mail{
		From: mail.Address{
			Address: "dpetroff@gmail.com",
			Name:    "Your gogo-cli-example",
		},
		To: []mail.Address{
			{
				Address: to,
			},
		},
		Subject: "Welcome to gogo-cli-example",
		Text: fmt.Sprintf(`
	Hi!

	Thank you for your interest in gogo-cli-example! Please follow the link to confirm your email address

	%s`, link),
		Html: fmt.Sprintf(`
	<p>Hi!</p>

	<p>Thank you for your interest in gogo-cli-example! Please follow the link to confirm your email address</p>

	<a href="%s">%s</a>`, link, link),
	}

	err := s.Send(ctx, mail)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
