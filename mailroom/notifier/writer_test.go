// Copyright 2024 SeatGeek, Inc.
//
// Licensed under the terms of the Apache-2.0 license. See LICENSE file in project root for terms.

package notifier_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/seatgeek/mailroom/mailroom/identifier"
	"github.com/seatgeek/mailroom/mailroom/notification"
	"github.com/seatgeek/mailroom/mailroom/notifier"
	"github.com/stretchr/testify/require"
)

func TestWriterNotifier_Push(t *testing.T) {
	t.Parallel()

	buffer := &bytes.Buffer{}
	notifier := notifier.NewWriterNotifier("buffer", buffer)

	err := notifier.Push(context.Background(), notification.NewBuilder("a1c11a53-c4be-488f-89b6-f83bf2d48dab", "com.example.test").
		WithRecipientIdentifiers(identifier.New(identifier.GenericUsername, "codell")).
		WithDefaultMessage("Hello, world!").
		Build(),
	)

	require.NoError(t, err)
	require.Equal(t, "notification: id=a1c11a53-c4be-488f-89b6-f83bf2d48dab type=com.example.test, to=[username:codell], message=Hello, world!\n", buffer.String())
}
