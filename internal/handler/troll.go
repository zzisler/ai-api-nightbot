package handler

import (
	"ai-nightbot/internal/ai"
	"ai-nightbot/internal/config"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	AiClient *ai.Client
	Cfg      *config.Config
}

func (h *Handler) Troll(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("token")
	if token != h.Cfg.SecretToken {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "что-то пошло не так")
		return
	}

	user := r.URL.Query().Get("user")
	if user == "" {
		user = "Зритель"
	}

	text := r.URL.Query().Get("text")
	text = strings.TrimSpace(text)
	if text == "" {
		fmt.Fprintf(w, "%s, добавь вопрос", user)
		return
	}

	log.Printf("[%s] question: %s", user, text)

	content := buildUserMessage(user, text, h.Cfg.PromptTemplate)
	reply, err := h.AiClient.AiReq([]ai.ReqMessage{{Role: "user", Content: content}})
	if err != nil {
		log.Printf("[%s] AI error: %v", user, err)
		fmt.Fprintf(w, "%s, ИИ перегружен.", user)
		return
	}

	if reply.Content == "" {
		log.Printf("[%s] empty AI response", user)
		fmt.Fprintf(w, "%s, ИИ перегружен.", user)
		return
	}

	log.Printf("[%s] answer: %s", user, reply.Content)

	fmt.Fprintf(w, "%s", reply.Content)

}

func buildUserMessage(user, text, template string) string {
	return fmt.Sprintf(template, user, text)
}
