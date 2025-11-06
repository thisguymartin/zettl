package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"time"

	domain "thisguymartin/zettl/internal/ui"
)

const (
	openAIBaseURL      = "https://api.openai.com/v1"
	embeddingModel     = "text-embedding-3-small"
	chatModel          = "gpt-4o-mini"
	embeddingDimension = 1536
)

// AIService handles AI operations including embeddings and chat
type AIService struct {
	apiKey     string
	httpClient *http.Client
}

// NewAIService creates a new AI service instance
func NewAIService(apiKey string) *AIService {
	return &AIService{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// EmbeddingRequest represents the OpenAI embedding API request
type EmbeddingRequest struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

// EmbeddingResponse represents the OpenAI embedding API response
type EmbeddingResponse struct {
	Data []struct {
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
}

// ChatRequest represents the OpenAI chat completion request
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

// ChatMessage represents a chat message
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatResponse represents the OpenAI chat completion response
type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

// GenerateEmbedding generates an embedding vector for the given text
func (s *AIService) GenerateEmbedding(text string) ([]float64, error) {
	reqBody := EmbeddingRequest{
		Input: text,
		Model: embeddingModel,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", openAIBaseURL+"/embeddings", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var embeddingResp EmbeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&embeddingResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(embeddingResp.Data) == 0 {
		return nil, fmt.Errorf("no embedding data in response")
	}

	return embeddingResp.Data[0].Embedding, nil
}

// GenerateNoteEmbedding creates an embedding for a note (title + content)
func (s *AIService) GenerateNoteEmbedding(note *domain.Note) (string, error) {
	text := fmt.Sprintf("%s\n\n%s", note.Title, note.Content)
	embedding, err := s.GenerateEmbedding(text)
	if err != nil {
		return "", err
	}

	// Convert to JSON string for storage
	jsonData, err := json.Marshal(embedding)
	if err != nil {
		return "", fmt.Errorf("failed to marshal embedding: %w", err)
	}

	return string(jsonData), nil
}

// CosineSimilarity calculates the cosine similarity between two vectors
func CosineSimilarity(a, b []float64) float64 {
	if len(a) != len(b) {
		return 0.0
	}

	var dotProduct, normA, normB float64
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}

	if normA == 0 || normB == 0 {
		return 0.0
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// SearchSimilarNotes finds notes similar to the query using vector similarity
func (s *AIService) SearchSimilarNotes(query string, notes []domain.Note, topK int) ([]domain.Note, error) {
	// Generate embedding for the query
	queryEmbedding, err := s.GenerateEmbedding(query)
	if err != nil {
		return nil, fmt.Errorf("failed to generate query embedding: %w", err)
	}

	// Calculate similarity scores
	type scoredNote struct {
		note  domain.Note
		score float64
	}

	var scored []scoredNote
	for _, note := range notes {
		if note.Embedding == "" {
			continue
		}

		// Parse note embedding
		var noteEmbedding []float64
		if err := json.Unmarshal([]byte(note.Embedding), &noteEmbedding); err != nil {
			continue
		}

		score := CosineSimilarity(queryEmbedding, noteEmbedding)
		scored = append(scored, scoredNote{note: note, score: score})
	}

	// Sort by score (descending)
	for i := 0; i < len(scored); i++ {
		for j := i + 1; j < len(scored); j++ {
			if scored[j].score > scored[i].score {
				scored[i], scored[j] = scored[j], scored[i]
			}
		}
	}

	// Return top K results
	if topK > len(scored) {
		topK = len(scored)
	}

	result := make([]domain.Note, topK)
	for i := 0; i < topK; i++ {
		result[i] = scored[i].note
	}

	return result, nil
}

// Chat sends a message to the AI and returns the response
func (s *AIService) Chat(userMessage string, history []domain.ChatMessage, context []domain.Note) (string, error) {
	// Build conversation with context
	var messages []ChatMessage

	// Add system message with context from notes
	if len(context) > 0 {
		var contextParts []string
		contextParts = append(contextParts, "You are a helpful AI assistant that helps users with their notes. Here are some relevant notes for context:")
		for i, note := range context {
			contextParts = append(contextParts, fmt.Sprintf("\n--- Note %d: %s ---\n%s", i+1, note.Title, note.Content))
		}
		contextParts = append(contextParts, "\nPlease use these notes to provide helpful and contextual responses.")

		messages = append(messages, ChatMessage{
			Role:    "system",
			Content: strings.Join(contextParts, "\n"),
		})
	} else {
		messages = append(messages, ChatMessage{
			Role:    "system",
			Content: "You are a helpful AI assistant that helps users with their notes and tasks.",
		})
	}

	// Add chat history
	for _, msg := range history {
		messages = append(messages, ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	// Add current user message
	messages = append(messages, ChatMessage{
		Role:    "user",
		Content: userMessage,
	})

	// Create request
	reqBody := ChatRequest{
		Model:    chatModel,
		Messages: messages,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", openAIBaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	return chatResp.Choices[0].Message.Content, nil
}
