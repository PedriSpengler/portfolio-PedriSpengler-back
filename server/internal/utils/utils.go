package utils

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"gopkg.in/gomail.v2"
    "os"
)

// Dados básicos para enviar um e-mail de boas-vindas
type EmailData struct {
    Email   string
    Subject string // assunto
    Message string
    Name    string
}

// SendEmail envia um e-mail usando o template fornecido
func SendEmail(recipient EmailData) error {
    from := os.Getenv("SMTP_EMAIL")
    smtpUser := os.Getenv("SMTP_EMAIL")
    smtpHost := "smtp.gmail.com"
    smtpPort := 587
    smtpPassword := os.Getenv("SMTP_PASSWORD")
    


    // Template simples
    tmpl, err := template.New("welcome").Parse(`
        <html>
        <body>
            <h3>Enviado pelo Portfólio Pedro Spengler</h3>
            <h1>{{.Name}} te enviou uma mensagem pelo email: {{.Email}}!</h1>
            <p> {{.Message}} </p>
        </body>
        </html>
    `)
    if err != nil {
        log.Printf("Erro ao criar template: %v", err)
        return err
    }

    // Preparando o corpo do e-mail
    var body bytes.Buffer
    if err := tmpl.Execute(&body, recipient); err != nil {
        log.Printf("Erro ao renderizar template: %v", err)
        return err
    }

    // Criando a mensagem de e-mail
    m := gomail.NewMessage()
    m.SetHeader("From", from)
    m.SetHeader("To", from)
    m.SetHeader("Subject", recipient.Subject)
    m.SetBody("text/html", body.String())

    // Configurando o envio via SMTP
    d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)
    if err := d.DialAndSend(m); err != nil {
        log.Printf("Erro ao enviar e-mail para %s: %v", recipient.Email, err)
        return err
    }

    log.Printf("E-mail enviado com sucesso para %s", recipient.Email)
    return nil
}

func CheckGoogleCaptcha(response string) bool {
	var googleCaptcha string = os.Getenv("RECAPTCHA_SECRET")
	req, _ := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", nil)
	q := req.URL.Query()
	q.Add("secret", googleCaptcha)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	var googleResponse map[string]interface{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &googleResponse)
	return googleResponse["success"].(bool)

}