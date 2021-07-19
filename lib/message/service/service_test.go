package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tylerlofgren/noiseaware/lib/message"
	"github.com/tylerlofgren/noiseaware/lib/message/repo/mock"
	"testing"
)

func TestMessageService_Get(t *testing.T) {
	expected := message.Message{
		Timestamp: 1,
		Symbol: "aaa",
		Volume: 10,
		Temperature: 10,
	}

	tests := []struct {
		desc string
		expected message.Message
		error error
	} {
		{"success", expected, nil},
		{"error", message.Message{}, errors.New("failure")},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mockRepo := mock.NewMockRepo(mockCtrl)
			mockRepo.EXPECT().Get().Return(tc.expected, tc.error).Times(1)

			ms := New(mockRepo)
			result, err := ms.Get()

			assert.Equal(t, tc.error, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
