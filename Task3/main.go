package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	models "task/Models"
)

func convertEvent(event models.Event) models.ConvertedEvent {
	converted := models.ConvertedEvent{
		Event:       event.Ev,
		EventType:   event.Et,
		AppID:       event.ID,
		UserID:      event.UID,
		MessageID:   event.MID,
		PageTitle:   event.T,
		PageURL:     event.P,
		BrowserLang: event.L,
		ScreenSize:  event.SC,
		Attributes: map[string]models.Attribute{
			event.Atk1: {Value: event.Atv1, Type: event.Atr1},
			event.Atk2: {Value: event.Atv2, Type: event.Atr2},
		},
		Traits: map[string]models.Trait{
			event.Uatk1: {Value: event.Uatv1, Type: event.Uatr1},
			event.Uatk2: {Value: event.Uatv2, Type: event.Uatr2},
			event.Uatk3: {Value: event.Uatv3, Type: event.Uatr3},
		},
	}
	return converted
}

// ... (previous code)

func sendToWebhook(convertedEvent models.ConvertedEvent) error {
	webhookURL := "https://gateway.webhooks.io/v1/hello/INcd7e636de46f4ad196cb11346c95b735"
	payload, err := json.Marshal(convertedEvent)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	fmt.Printf("resp: %v\n", resp)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Webhook request failed with status: %d", resp.StatusCode)
	}

	return nil
}

func handleRequest(ch chan models.Event) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var event models.Event
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&event)
		if err != nil {
			http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
			return
		}

		ch <- event
		w.WriteHeader(http.StatusOK)
	}
}

func worker(ch chan models.Event) {
	for {
		event := <-ch
		convertedEvent := convertEvent(event)
		err := sendToWebhook(convertedEvent)
		if err != nil {
			fmt.Println("Error sending to webhook:", err)
		}
		fmt.Println("Worker processed event:", event)
	}
}

func main() {
	eventChannel := make(chan models.Event)

	// Start the worker
	go worker(eventChannel)

	http.HandleFunc("/", handleRequest(eventChannel))

	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
