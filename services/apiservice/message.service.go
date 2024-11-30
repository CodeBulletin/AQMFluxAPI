package apiservice

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/repo"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type MessageService struct {
	db db.DB
	logger logger.Logger
}

func NewMessageService(db db.DB) *MessageService {
	logger := logger.GetLogger()
	return &MessageService{db: db, logger: logger}
}

func (m *MessageService) CreateMessage(w http.ResponseWriter, r *http.Request) {
	query := repo.New(m.db)
	defer query.Close()

	// Parse the request body into a Message struct
	var msg types.NewMessage
	err := utils.ParseRequest(r, &msg)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Create the message in the database
	data, err := query.CreateMessage(r.Context(), repo.CreateMessageParams{
		Title: msg.Title,
		Topic: msg.Topic,
		Payload: msg.Payload,
		Tags: sql.NullString{String: msg.Tags, Valid: true},
		Messagepriority: int32(msg.Priority),
	})

	if err != nil {
		m.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Write the response
	err = utils.WriteJSON(w, http.StatusCreated, types.OkCreatedJsonMessage{
		Message: "Message created successfully",
		Data:    data,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (m *MessageService) GetMessages(w http.ResponseWriter, r *http.Request) {
	query := repo.New(m.db)
	defer query.Close()

	// Get all messages from the database
	data, err := query.GetMessageslist(r.Context())

	if err != nil {
		m.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var messages []types.Messages

	for _, msg := range data {
		messages = append(messages, types.Messages{
			Id:    msg.ID,
			Title: msg.Title,
		})
	}

	// Write the response
	err = utils.WriteJSON(w, http.StatusOK, messages)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (m *MessageService) GetMessage(w http.ResponseWriter, r *http.Request) {
	query := repo.New(m.db)
	defer query.Close()

	// Get the message ID from the URL
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Get the message from the database
	data, err := query.GetMessages(r.Context(), int32(id))

	if err != nil {
		m.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var message types.Message = types.Message{
		Id:    data.ID,
		Title: data.Title,
		Topic: data.Topic,
		Payload: data.Payload,
		Tags: data.Tags.String,
		Priority: int(data.Messagepriority),
	}

	// Write the response
	err = utils.WriteJSON(w, http.StatusOK, message)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (m *MessageService) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	query := repo.New(m.db)
	defer query.Close()

	// Get the message ID from the URL
	// Parse the request body into a Message struct
	var msg types.Message
	err := utils.ParseRequest(r, &msg)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Update the message in the database
	data, err := query.UpdateMessage(r.Context(), repo.UpdateMessageParams{
		ID: msg.Id,
		Title: msg.Title,
		Topic: msg.Topic,
		Payload: msg.Payload,
		Tags: sql.NullString{String: msg.Tags, Valid: true},
		Messagepriority: int32(msg.Priority),
	})

	if err != nil {
		m.logger.Error(err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Write the response
	err = utils.WriteJSON(w, http.StatusOK, types.OkCreatedJsonMessage{
		Message: "Message updated successfully",
		Data:    data,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (m *MessageService) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	query := repo.New(m.db)
	defer query.Close()

	// Get the message ID from the URL
	// Parse the request body into a Message struct
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Delete the message from the database
	data, err := query.DeleteMessage(r.Context(), int32(id))

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Write the response
	err = utils.WriteJSON(w, http.StatusOK, types.OkCreatedJsonMessage{
		Message: "Message deleted successfully",
		Data:   data,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}