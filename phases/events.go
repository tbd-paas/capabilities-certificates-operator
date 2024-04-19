package phases

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Event int

const (
	Unknown Event = iota
	Created
	Updated
	Deleted
)

const (
	UnknownString = "Unknown"
	CreatedString = "Created"
	UpdatedString = "Updated"
	DeletedString = "Deleted"
)

// String returns the string value of an event.
func (event Event) String() string {
	return map[Event]string{
		Unknown: UnknownString,
		Created: CreatedString,
		Updated: UpdatedString,
		Deleted: DeletedString,
	}[event]
}

// Type returns the type of event.
func (event Event) Type() string {
	return map[Event]string{
		Unknown: UnknownString,
		Created: corev1.EventTypeNormal,
		Updated: corev1.EventTypeNormal,
		Deleted: corev1.EventTypeNormal,
	}[event]
}

// RegisterAction registers an event.
func (event Event) RegisterAction(recorder record.EventRecorder, child, parent client.Object) {
	recorder.Event(
		child,
		event.Type(),
		fmt.Sprintf("%s%s", child.GetObjectKind().GroupVersionKind(), event.String()),
		fmt.Sprintf(
			"%s child resource '%s' managed by parent resource '%s'",
			event.String(),
			getMessageString(child),
			getMessageString(parent),
		),
	)
}

// getMessageString gets the message string for an object.  The message string is the message that is
// displayed when a resource is actioned upon.
func getMessageString(object client.Object) string {
	return fmt.Sprintf("%s/%s", object.GetObjectKind().GroupVersionKind(), object.GetName())
}