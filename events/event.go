package events

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type EventHandler struct {
	ws            *websocket.Conn
	eventHandlers eventHandlers
}

// Connect to the server for processing events
func Connect(host string, clientToken string, accessToken string) (*EventHandler, error) {
	origin := "https://" + host + "/wss/connector"
	url := "wss://" + host + "/ws/connector?ClientToken=" + clientToken + "&AccessToken=" + accessToken
	fmt.Println(origin)
	ws, err := websocket.Dial(url, "wss", origin)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = ws.Close()
		}
	}()

	var h = &EventHandler{
		ws: ws,
	}

	return h, nil
}

// Start handling events from the server
func (h *EventHandler) Handling() {
	var (
		handlers *MewsEvents
		err      error
	)

	decoder := json.NewDecoder(h.ws)
	for {
		if err = decoder.Decode(&handlers); err != nil {
			log.Println(err)
			break
		}
		for i := 0; i < len(handlers.Events); i++ {
			h.eventHandlers.Send(&handlers.Events[i])
		}
	}
}

// Handler function definition.
// If the function returns error then procession will be stopped.
type Handler = func(event *MewsEvent) error

// Add handler in events processing list
func (h *EventHandler) AddHandler(handler Handler, events ...string) {
	if len(events) == 0 {
		return // no events for processing
	}
	// make communication channel and make it available for all events names
	go func() {
		var eventChan = make(chan *MewsEvent, 1)
		defer close(eventChan) // close channel on exit
		h.eventHandlers.Store(eventChan, events...)
		defer h.eventHandlers.Delete(eventChan) // remove handler on exit

		for {
			select {
			case resp := <-eventChan: // got event from the server
				// empty event on close
				if resp == nil {
					return
				}
				// handle event and check for error
				switch err := handler(resp); err {
				case nil:
					continue // wait for next event
				case Stop:
					return
				default:
					return
				}
			}
		}
	}()
}

// HandleWait call handler function for processing all events from the server.
// timeout define maximum time to wait response from the server
// Function return ErrTimeout on no response from the server.
// Function infinite wait for response in case timeout is equal 0 or below.
func (h *EventHandler) HandleWait(handler Handler, timeout time.Duration,
	events ...string) (err error) {
	if len(events) == 0 {
		return nil // no events to processing
	}
	// make communication channel and make it available for all events names
	var eventChan = make(chan *MewsEvent, 1)
	defer close(eventChan)                      // close channel on exit
	h.eventHandlers.Store(eventChan, events...) // registering events for handler
	defer h.eventHandlers.Delete(eventChan)     // remove handlers on exit

	// set up timer for the operation
	var timeoutTimer = time.NewTimer(timeout)
	if timeout <= 0 {
		<-timeoutTimer.C // reset timer
	}
	for {
		select {
		case resp := <-eventChan: // got event from the server
			// empty event on close
			if resp == nil {
				timeoutTimer.Stop()
				return nil
			}
			// run event handler and process return value
			switch err = handler(resp); err {
			case nil:
				if timeout > 0 { // move timer
					timeoutTimer.Reset(timeout)
				}
				continue // wait for next event for processing
			case Stop:
				return nil
			default:
				return err
			}
		case <-timeoutTimer.C:
			return ErrTimeout // timeout error
		}
	}
}

// Handle simple call HandleWait with 0 timeout.
func (h *EventHandler) Handle(handler Handler, events ...string) error {
	return h.HandleWait(handler, 0, events...)
}

func (h *EventHandler) Close() error {
	h.eventHandlers.Close()
	if err := h.ws.Close(); err != nil {
		h.ws = nil
		return err
	}
	h.ws = nil
	return nil
}

func (h *EventHandler) IsRunning() bool {
	if h.ws == nil {
		return false
	}
	return true
}
