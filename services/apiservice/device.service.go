package apiservice

import (
	"fmt"
	"net/http"

	"github.com/codebulletin/AQMFluxAPI/db"
	"github.com/codebulletin/AQMFluxAPI/logger"
	"github.com/codebulletin/AQMFluxAPI/repo"
	"github.com/codebulletin/AQMFluxAPI/types"
	"github.com/codebulletin/AQMFluxAPI/utils"
)

type DeviceService struct {
	db db.DB
	logger logger.Logger
}

func NewDeviceService(db db.DB) *DeviceService {
	logger := logger.GetLogger()
	return &DeviceService{
		db: db,
		logger: logger,
	}
}

func (h *DeviceService) CreateDevice(w http.ResponseWriter, r *http.Request) {
	tx, err := h.db.Begin()
	if err != nil {
		h.logger.Error("Error starting transaction: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error starting transaction: %v", err))
		return
	}

	defer tx.Rollback()

	query := repo.New(h.db)
	qtx := query.WithTx(tx)
	defer query.Close()

	var device types.NewDevice
	err = utils.ParseRequest(r, &device)

	if err != nil {
		h.logger.Error("Error parsing request: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	err = qtx.CreateDevice(r.Context(), repo.CreateDeviceParams{
		DeviceName: device.Name,
		LocationID: device.Location,
		DeviceDesc: device.Desc,
		DeviceID: device.Id,
	})

	if err != nil {
		h.logger.Error("Error creating device: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating device: %v", err))
		return
	}

	err = qtx.CreateDeviceAddr(r.Context(), repo.CreateDeviceAddrParams{
		DeviceID: device.Id,
		IpAddr: device.IPAddress,
		MacAddr: device.MACAddress,
		Port: device.Port,
	})

	if err != nil {
		h.logger.Error("Error creating device address: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating device address: %v", err))
		return
	}

	for _, sensor := range device.Sensors {
		err = qtx.AddSensorToDevice(r.Context(), repo.AddSensorToDeviceParams{
			SensorID: sensor,
			DeviceID: device.Id,
		})

		if err != nil {
			h.logger.Error("Error creating device sensor: %v", err)
			utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating device sensor: %v", err))
			return
		}
	}

	err = tx.Commit()

	if err != nil {
		h.logger.Error("Error committing transaction: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error committing transaction: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, types.OkCreatedJsonMessage{
		Message: "Device created successfully",
		Data: device,
	})
}

func (h *DeviceService) UpdateDevice(w http.ResponseWriter, r *http.Request) {
}

func (h *DeviceService) GetDevice(w http.ResponseWriter, r *http.Request) {
	query := repo.New(h.db)
	defer query.Close()

	devices, err := query.GetAllDeviceInformation(r.Context())

	if err != nil {
		h.logger.Error("Error getting device information: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error getting device information: %v", err))
		return
	}

	var deviceList []types.NewDevice

	for _, device := range devices {

		sensors, err := query.GetDeviceSensors(r.Context(), device.DeviceID)

		if err != nil {
			h.logger.Error("Error getting device sensors: %v", err)
			utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error getting device sensors: %v", err))
			return
		}

		var sensorList []int32
		for _, sensor := range sensors {
			sensorList = append(sensorList, sensor.SensorID)
		}

		deviceList = append(deviceList, types.NewDevice{
			Name: device.DeviceName,
			Id: device.DeviceID,
			Desc: device.DeviceDesc,
			IPAddress: device.IpAddr,
			MACAddress: device.MacAddr,
			Port: device.Port,
			Location: device.LocationID,
			Sensors: sensorList,
		})
	}

	err = utils.WriteJSON(w, http.StatusOK, deviceList)

	if err != nil {
		h.logger.Error("Error writing device information: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error writing device information: %v", err))
		return
	}
}

func (h *DeviceService) GetDeviceList (w http.ResponseWriter, r *http.Request) {
	query := repo.New(h.db)
	defer query.Close()

	data, err := query.GetDevicesList(r.Context())

	if err != nil {
		h.logger.Error("Error getting device list: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error getting device list: %v", err))
		return
	}

	var devices = make([]types.List, len(data))

	for i, device := range data {
		devices[i] = types.List{
			Id: device.DeviceID,
			Name: device.DeviceName,
		}
	}

	err = utils.WriteJSON(w, http.StatusOK, devices)

	if err != nil {
		h.logger.Error("Error writing device list: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error writing device list: %v", err))
		return
	}
}