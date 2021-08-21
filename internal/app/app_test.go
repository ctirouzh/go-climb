package app

import (
	"github.com/pkritiotis/go-climb/internal/app/crag/commands"
	"github.com/pkritiotis/go-climb/internal/app/crag/queries"
	"github.com/pkritiotis/go-climb/internal/app/notification"
	"github.com/pkritiotis/go-climb/internal/domain/crag"
	"github.com/pkritiotis/go-climb/internal/pkg/time"
	"github.com/pkritiotis/go-climb/internal/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	mockRepo := crag.MockRepository{}
	UUIDProvider := &uuid.MockProvider{}
	timeProvider := &time.MockProvider{}
	notificationService := notification.MockNotificationService{}

	type args struct {
		up                  uuid.Provider
		tp                  time.Provider
		cragRepo            crag.Repository
		notificationService notification.Service
	}
	tests := []struct {
		name string
		args args
		want App
	}{
		{
			name: "should initialize application layer",
			args: args{
				cragRepo:            mockRepo,
				notificationService: notificationService,
				up:                  UUIDProvider,
				tp:                  timeProvider,
			},
			want: App{
				Queries: Queries{
					GetAllCragsHandler: queries.NewGetAllCragsRequestHandler(mockRepo),
					GetCragHandler:     queries.NewGetCragRequestHandler(mockRepo),
				},
				Commands: Commands{
					AddCragHandler:    commands.NewAddCragRequestHandler(UUIDProvider, timeProvider, mockRepo, notificationService),
					UpdateCragHandler: commands.NewUpdateCragRequestHandler(mockRepo),
					DeleteCragHandler: commands.NewDeleteCragRequestHandler(mockRepo),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApp(tt.args.cragRepo, tt.args.notificationService, tt.args.up, tt.args.tp)
			assert.Equal(t, tt.want, got)
		})
	}
}
