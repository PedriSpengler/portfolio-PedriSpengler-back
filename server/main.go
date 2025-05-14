package main

import (
    "envia-email/internal/handler"
    "log"
    "net/http"
    "os"
    "github.com/joho/godotenv"
    "github.com/rs/cors"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Printf("Aviso: arquivo .env não carregado: %v", err)
    }
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // fallback para rodar localmente
    }

    // Configuração do CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"https://www.pedrispengler.com.br"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization", "accessToken", "Refresh-Token", "Access-Control-Allow-Origin", "Cache-Control", "X-Requested-With"},
        ExposedHeaders:   []string{"Content-Type", "Authorization"},
    })

    // Envolve seu handler com CORS
    handlerWithCORS := c.Handler(http.HandlerFunc(handler.SendEmailHandler))

    http.Handle("/api/send-email", handlerWithCORS)

    log.Printf("Servidor escutando na porta %s...", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}