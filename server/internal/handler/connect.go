package handler

import (
	"encoding/json"
	"envia-email/internal/utils"
	"fmt"
	"net/http"
)

type EmailRequest struct {
    Email   string `json:"email"`
    Subject string `json:"subject"`
    Name    string `json:"name"`
    Message string `json:"message"`
    Captcha string `json:"captcha"`
}

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var req EmailRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Erro ao processar os dados: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Validar os campos obrigatórios
    if req.Email == "" || req.Name == "" || req.Subject == "" || req.Message == "" || req.Captcha == "" {
        http.Error(w, "Campos email, assunto, nome e mensagem são obrigatórios", http.StatusBadRequest)
        return
    }
    fmt.Println(req.Captcha)
    valid := utils.CheckGoogleCaptcha(req.Captcha)
	if !valid {
		http.Error(w, "Captcha está inválido!", http.StatusBadRequest)
        return
	}

    // Criar os dados para envio do e-mail
    emailData := utils.EmailData{
        Email:   req.Email,
        Subject: req.Subject,
        Message: req.Message,
        Name:    req.Name,  
    }

    // Enviar o e-mail usando a função do pacote utils
    if err := utils.SendEmail(emailData); err != nil {
        http.Error(w, "Erro ao enviar e-mail: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Resposta de sucesso
    response := "E-mail enviado com sucesso!"
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}