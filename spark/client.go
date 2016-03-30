package main

import (
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
)

// Send message to the room with specified id
func sendMessageToRoom(ctx context.Context, roomID string, message string) {
	log.Debugf(ctx, "Would send message %s to room %s", message, roomID)
	log.Warningf(ctx, "Not implemented yet")
}


