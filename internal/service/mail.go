package service

import (
	"context"
	ht "html/template"
	tt "text/template"

	"github.com/XIAOR1Ck/localsewa-backend/internal/config"
	"github.com/wneessen/go-mail"
)

type MailService struct {
	client *mail.Client
}

type MailFormat struct {
	MailTo, Subject, Name string
}

// Sernder Information

const (
	senderName    = "LocalSewa Co."
	senderAddress = "noreply@localsewa.com.np"
)

const (
	textBodyTemplate = `Hi {{.Name}},

we are writing your to Thank you for joining Localsewa.
I hope all your requirements can be fulfilled here.

Check out all our service providers here at https://localsewa.com.np

With Regards,
  LocalSewa Co.`

	htmlBodyTemplate = `<p>Hi {{.Firstname}},</p>
<p>we are writing your to Thank you for joining "<strong>Localsewa</strong>".</p>
<p>I hope all your requirements can be fulfilled here.</p>
<p>Check out all our service providers here at <a href="https://localsewa.com.np" target="_blank">https://localsewa.com.np</a></p>
<p>With Regards<br />
&nbsp;&nbsp;LocalSewa Co.</p>`
)

func NewMailService(client *mail.Client) *MailService {
	return &MailService{
		client: client,
	}
}

func NewMailClient(cfg *config.Config) (*mail.Client, error) {
	return mail.NewClient(
		"smtp.gmail.com",
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(cfg.SMTPUser),
		mail.WithPassword(cfg.SMTPPass),
	)
}

func (s *MailService) SendMail(ctx context.Context, data MailFormat) error {
	message := mail.NewMsg()
	textTpl, err := tt.New("texttpl").Parse(textBodyTemplate)
	if err != nil {
		return err
	}
	htmlTpl, err := ht.New("htmltpl").Parse(htmlBodyTemplate)
	if err != nil {
		return err
	}
	if err := message.EnvelopeFrom("noreply-1@localsewa.com.np"); err != nil {
		return err
	}
	if err := message.FromFormat(senderName, senderAddress); err != nil {
		return err
	}
	if err := message.AddToFormat(data.Name, data.MailTo); err != nil {
		return err
	}
	message.SetMessageID()
	message.SetDate()
	message.Subject(data.Subject)

	// message template
	if err := message.SetBodyTextTemplate(textTpl, data); err != nil {
		return err
	}

	if err := message.AddAlternativeHTMLTemplate(htmlTpl, data); err != nil {
		return nil
	}

	if err := s.client.DialAndSendWithContext(ctx, message); err != nil {
		return err
	}
	return nil

}
