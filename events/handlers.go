package events

import (
	"sync"
)

type MewsEvent struct {
	State           string `json:"State"`
	Type            string `json:"Type"`
	Id              string `json:"Id"`
	StartUtc        string `json:"StartUtc,omitempty"`
	EndUtc          string `json:"EndUtc,omitempty"`
	AssignedSpaceId string `json:"AssignedSpaceId,omitempty"`
}

type MewsEvents struct {
	Events []MewsEvent `json:"Events"`
}

// Stop event to stop process by Handler.
var Stop error = new(emptyError)

// ErrTimeout возвращается когда ответ от сервера на команду не получен за время
// ReadTimeout.
var ErrTimeout error = new(timeoutError)

type timeoutError struct{}

func (timeoutError) Error() string   { return "response timeout" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }

type emptyError struct{}

func (emptyError) Error() string { return "" }


// eventHandlers list of handlers grouped by event name
type eventHandlers struct {
	handlers sync.Map
}

// Store add new handler to process the new event.
func (ehs *eventHandlers) Store(ch chan<- *MewsEvent, events ...string) {
	for _, event := range events {
		list, ok := ehs.handlers.Load(event)
		if !ok {
			list = mapOfHandlerChan{ch: struct{}{}}
		} else {
			list.(mapOfHandlerChan)[ch] = struct{}{}
		}
		ehs.handlers.Store(event, list) // сохраняем обновленный список
	}
}

// Delete remove registered handler for thr event.
func (ehs *eventHandlers) Delete(ch chan<- *MewsEvent) {
	ehs.handlers.Range(func(event, list interface{}) bool {
		var handlers = list.(mapOfHandlerChan)
		if _, ok := handlers[ch]; ok {
			if len(handlers) < 2 {
				ehs.handlers.Delete(event)
			} else {
				delete(handlers, ch)
				ehs.handlers.Store(event, handlers)
			}
		}
		return true
	})
}

// Send send event to all registered handlers
func (ehs *eventHandlers) Send(resp *MewsEvent) {
	if list, ok := ehs.handlers.Load(resp.Type); ok {
		for handler := range list.(mapOfHandlerChan) {
			handler <- resp
		}
	}
}

// Close send empty event to all handlers to stop processing new events.
func (ehs *eventHandlers) Close() {
	// build collection of unique handlers
	var handlers = make(mapOfHandlerChan)
	ehs.handlers.Range(func(event, list interface{}) bool {
		for handler := range list.(mapOfHandlerChan) {
			handlers[handler] = struct{}{}
		}
		ehs.handlers.Delete(event) // this delete must be fastens
		return true
	})
	// send empty responses on close
	for handler := range handlers {
		handler <- nil
	}
}

type (
	// mapOfHandlerChan used as alias to define list of events handlers.
	mapOfHandlerChan = map[chan<- *MewsEvent]struct{}
	// responseChan define channel to process event.
	responseChan = chan *MewsEvent
)
