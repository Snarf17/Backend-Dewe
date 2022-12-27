package handlers

import (
	dto "dewetour/dto/result"
	transactiondto "dewetour/dto/transaction"
	"dewetour/models"
	"dewetour/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type Transactionhand struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *Transactionhand {
	return &Transactionhand{TransactionRepository}
}

func (h *Transactionhand) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *Transactionhand) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	// trip.Image = path_file + trip.Image
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

// var file = "http://localhost:9000/uploads/"

func (h *Transactionhand) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	UserID := int(UserInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// request := new(tripdto.CreateTripRequest)

	qty, _ := strconv.Atoi(r.FormValue("qty"))
	total, _ := strconv.Atoi(r.FormValue("total"))
	TripID, _ := strconv.Atoi(r.FormValue("trip_id"))
	// UserID, _ := strconv.Atoi(r.FormValue("user_id"))
	request := transactiondto.CreateTransaction{
		TripID:     TripID,
		CounterQty: qty,
		Total:      total,
		Status:     r.FormValue("status"),
		Attachment: filename,
		UserID:     UserID,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.Transaction{
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		TripID:     request.TripID,
		Attachment: filename,
		UserID:     UserID,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	get, err := h.TransactionRepository.GetTransaction(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	transaction.Attachment = path_file + transaction.Attachment
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: get}
	json.NewEncoder(w).Encode(response)
}

func (h *Transactionhand) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	UserInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(UserInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	req := transactiondto.UpdateTransaction{
		Attachment: filename,
		UserID:     userID,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	getData, err := h.TransactionRepository.GetTransaction(int(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	qty, _ := strconv.Atoi(r.FormValue("qty"))
	if qty != 0 {
		getData.CounterQty = qty
	}
	total, _ := strconv.Atoi(r.FormValue("total"))
	if total != 0 {
		getData.Total = total
	}
	if r.FormValue("status") != "" {
		getData.Status = r.FormValue("status")
	}
	tripID, _ := strconv.Atoi(r.FormValue("trip_id"))
	if tripID != 0 {
		getData.TripID = tripID
	}
	if req.Attachment != "" {
		getData.Attachment = req.Attachment
	}

	// update data
	data, err := h.TransactionRepository.UpdateTransaction(getData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// get data
	GetID, err := h.TransactionRepository.GetTransaction(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}
	// success
	getData.Attachment = path_file + getData.Attachment
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: GetID}
	json.NewEncoder(w).Encode(response)
}
