package controllers

import (
	"errors"

	"code.cloudfoundry.org/auctioneer"
	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/events"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/bbs/serviceclient"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/rep"
)

type ActualLRPLifecycleController struct {
	db               db.ActualLRPDB
	evacuationDB     db.EvacuationDB
	desiredLRPDB     db.DesiredLRPDB
	auctioneerClient auctioneer.Client
	serviceClient    serviceclient.ServiceClient
	repClientFactory rep.ClientFactory
	actualHub        events.Hub
}

func NewActualLRPLifecycleController(
	db db.ActualLRPDB,
	evacuationDB db.EvacuationDB,
	desiredLRPDB db.DesiredLRPDB,
	auctioneerClient auctioneer.Client,
	serviceClient serviceclient.ServiceClient,
	repClientFactory rep.ClientFactory,
	actualHub events.Hub,
) *ActualLRPLifecycleController {
	return &ActualLRPLifecycleController{
		db:               db,
		evacuationDB:     evacuationDB,
		desiredLRPDB:     desiredLRPDB,
		auctioneerClient: auctioneerClient,
		serviceClient:    serviceClient,
		repClientFactory: repClientFactory,
		actualHub:        actualHub,
	}
}

func (h *ActualLRPLifecycleController) ClaimActualLRP(logger lager.Logger, processGuid string, index int32, actualLRPInstanceKey *models.ActualLRPInstanceKey) error {
	before, after, err := h.db.ClaimActualLRP(logger, processGuid, index, actualLRPInstanceKey)
	if err != nil {
		return err
	}

	if !after.Equal(before) {
		go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
	}
	return nil
}

func (h *ActualLRPLifecycleController) StartActualLRP(logger lager.Logger, actualLRPKey *models.ActualLRPKey, actualLRPInstanceKey *models.ActualLRPInstanceKey, actualLRPNetInfo *models.ActualLRPNetInfo) error {
	lrpg, err := h.db.ActualLRPGroupByProcessGuidAndIndex(logger, actualLRPKey.ProcessGuid, actualLRPKey.Index)
	if err != nil {
		logger.Error("err-when-finding-suspect", err)
		if err == models.ErrResourceNotFound {
			logger.Info("create-missing-actual-lrp", lager.Data{"guid": actualLRPKey.ProcessGuid, "index": actualLRPKey.Index, "instance-guid": actualLRPInstanceKey.InstanceGuid})
			return h.db.CreateRunningActualLRP(logger, actualLRPKey, actualLRPInstanceKey, actualLRPNetInfo)
		}
		return err
	}

	// TODO: decide what happens if the suspect is already running
	if lrpg.Suspect != nil {
		logger = logger.Session("found-suspect", lager.Data{"guid": actualLRPKey.ProcessGuid, "index": actualLRPKey.Index})
		suspectLRP := lrpg.Suspect
		if suspectLRP.ActualLRPInstanceKey.InstanceGuid != actualLRPInstanceKey.InstanceGuid &&
			actualLRPInstanceKey.InstanceGuid == lrpg.Instance.ActualLRPInstanceKey.InstanceGuid {
			logger.Info("starting-shadow", lager.Data{"suspect-instance-guid": suspectLRP.ActualLRPInstanceKey.InstanceGuid, "shadow-instance-guid": actualLRPInstanceKey.InstanceGuid})
			h.db.RemoveActualLRP(logger, lrpg.Suspect.ProcessGuid, lrpg.Suspect.Index, &lrpg.Suspect.ActualLRPInstanceKey)
			h.actualHub.Emit(models.NewActualLRPRemovedEvent(&models.ActualLRPGroup{Suspect: suspectLRP}))
		} else {
			err := errors.New("rep-restarting-suspect")
			logger.Error("rep-cannot-restart-suspect", err, lager.Data{"suspect-instance-guid": suspectLRP.ActualLRPInstanceKey.InstanceGuid})
			return err
		}
	}

	before, after, err := h.db.StartActualLRP(logger, actualLRPKey, actualLRPInstanceKey, actualLRPNetInfo)
	if err != nil {
		return err
	}

	var lrpGroup *models.ActualLRPGroup
	lrpGroup, err = h.db.ActualLRPGroupByProcessGuidAndIndex(logger, actualLRPKey.ProcessGuid, actualLRPKey.Index)
	if err != nil {
		return err
	}

	if lrpGroup.Evacuating != nil {
		h.evacuationDB.RemoveEvacuatingActualLRP(logger, &lrpGroup.Evacuating.ActualLRPKey, &lrpGroup.Evacuating.ActualLRPInstanceKey)
	}

	go func() {
		if before == nil {
			h.actualHub.Emit(models.NewActualLRPCreatedEvent(after))
		} else if !before.Equal(after) {
			h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
		}
		if lrpGroup.Evacuating != nil {
			h.actualHub.Emit(models.NewActualLRPRemovedEvent(&models.ActualLRPGroup{Evacuating: lrpGroup.Evacuating}))
		}
	}()
	return nil
}

func (h *ActualLRPLifecycleController) CrashActualLRP(logger lager.Logger, actualLRPKey *models.ActualLRPKey, actualLRPInstanceKey *models.ActualLRPInstanceKey, errorMessage string) error {
	lrpg, err := h.db.ActualLRPGroupByProcessGuidAndIndex(logger, actualLRPKey.ProcessGuid, actualLRPKey.Index)
	if err != nil {
		logger.Error("err-when-finding-actual-lrp-group", err)
		return err
	}

	logger.Info("actual-lrpg-crash", lager.Data{"lrpg": lrpg})
	if lrpg.Suspect != nil {
		if lrpg.Suspect.ActualLRPInstanceKey.InstanceGuid == actualLRPInstanceKey.InstanceGuid {
			suspectLRP := lrpg.Suspect
			logger.Info("found-crashed-suspect", lager.Data{"guid": actualLRPKey.ProcessGuid, "index": actualLRPKey.Index, "instance-guid": actualLRPInstanceKey.InstanceGuid})
			err = h.db.RemoveActualLRP(logger, suspectLRP.ProcessGuid, suspectLRP.Index, &suspectLRP.ActualLRPInstanceKey)
			h.actualHub.Emit(models.NewActualLRPRemovedEvent(&models.ActualLRPGroup{Suspect: suspectLRP}))
			return err
		}

		if lrpg.Instance.ActualLRPInstanceKey.InstanceGuid == actualLRPInstanceKey.InstanceGuid {
			instanceLRP := lrpg.Instance
			logger.Info("found-crashed-instance", lager.Data{"guid": actualLRPKey.ProcessGuid, "index": actualLRPKey.Index, "instance-guid": actualLRPInstanceKey.InstanceGuid})
			err = h.db.RemoveActualLRP(logger, instanceLRP.ProcessGuid, instanceLRP.Index, &instanceLRP.ActualLRPInstanceKey)
			if err != nil {
				return err
			}
			_, err = h.db.CreateUnclaimedActualLRP(logger, actualLRPKey)
			if err != nil {
				return err
			}
		}
	}

	if lrpg.Suspect == nil {
		before, after, shouldRestart, err := h.db.CrashActualLRP(logger, actualLRPKey, actualLRPInstanceKey, errorMessage)
		if err != nil {
			return err
		}

		beforeActualLRP, _ := before.Resolve()
		afterActualLRP, _ := after.Resolve()
		go h.actualHub.Emit(models.NewActualLRPCrashedEvent(beforeActualLRP, afterActualLRP))
		go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
		if !shouldRestart {
			return nil
		}
	}

	desiredLRP, err := h.desiredLRPDB.DesiredLRPByProcessGuid(logger, actualLRPKey.ProcessGuid)
	if err != nil {
		logger.Error("failed-fetching-desired-lrp", err)
		return err
	}

	schedInfo := desiredLRP.DesiredLRPSchedulingInfo()
	startRequest := auctioneer.NewLRPStartRequestFromSchedulingInfo(&schedInfo, int(actualLRPKey.Index))
	logger.Info("start-lrp-auction-request", lager.Data{"app_guid": schedInfo.ProcessGuid, "index": int(actualLRPKey.Index)})
	err = h.auctioneerClient.RequestLRPAuctions(logger, []*auctioneer.LRPStartRequest{&startRequest})
	logger.Info("finished-lrp-auction-request", lager.Data{"app_guid": schedInfo.ProcessGuid, "index": int(actualLRPKey.Index)})
	if err != nil {
		logger.Error("failed-requesting-auction", err)
	}
	return nil
}

func (h *ActualLRPLifecycleController) FailActualLRP(logger lager.Logger, key *models.ActualLRPKey, errorMessage string) error {
	before, after, err := h.db.FailActualLRP(logger, key, errorMessage)
	if err != nil {
		return err
	}

	go h.actualHub.Emit(models.NewActualLRPChangedEvent(before, after))
	return nil
}

func (h *ActualLRPLifecycleController) RemoveActualLRP(logger lager.Logger, processGuid string, index int32, instanceKey *models.ActualLRPInstanceKey) error {
	beforeActualLRPGroup, err := h.db.ActualLRPGroupByProcessGuidAndIndex(logger, processGuid, index)
	if err != nil {
		return err
	}

	err = h.db.RemoveActualLRP(logger, processGuid, index, instanceKey)
	if err != nil {
		return err

	}

	// what event to emit when removing an ActualLRPGroup with both Shadow and Suspect
	go h.actualHub.Emit(models.NewActualLRPRemovedEvent(beforeActualLRPGroup))
	return nil
}

func (h *ActualLRPLifecycleController) RetireActualLRP(logger lager.Logger, key *models.ActualLRPKey) error {
	var err error
	var cell *models.CellPresence

	logger = logger.Session("retire-actual-lrp", lager.Data{"process_guid": key.ProcessGuid, "index": key.Index})

	for retryCount := 0; retryCount < models.RetireActualLRPRetryAttempts; retryCount++ {
		var lrpGroup *models.ActualLRPGroup
		lrpGroup, err = h.db.ActualLRPGroupByProcessGuidAndIndex(logger, key.ProcessGuid, key.Index)
		if err != nil {
			return err
		}

		// if there is a suspect actual lrp, just remove it
		if lrpGroup.Suspect != nil {
			suspectActualLRP := lrpGroup.Suspect
			logger = logger.Session("found-suspect", lager.Data{"guid": suspectActualLRP.ProcessGuid, "index": suspectActualLRP.Index, "instance-guid": suspectActualLRP.ActualLRPInstanceKey.InstanceGuid})
			err = h.db.RemoveActualLRP(logger, suspectActualLRP.ProcessGuid, suspectActualLRP.Index, &suspectActualLRP.ActualLRPInstanceKey)
			logger.Error("retrying-failed-retire-of-actual-lrp", err, lager.Data{"attempt": retryCount + 1})
			continue
		}

		lrp := lrpGroup.Instance
		if lrp == nil {
			return models.ErrResourceNotFound
		}

		switch lrp.State {
		case models.ActualLRPStateUnclaimed, models.ActualLRPStateCrashed:
			err = h.db.RemoveActualLRP(logger, lrp.ProcessGuid, lrp.Index, &lrp.ActualLRPInstanceKey)
			if err == nil {
				go h.actualHub.Emit(models.NewActualLRPRemovedEvent(lrpGroup))
			}
		case models.ActualLRPStateClaimed, models.ActualLRPStateRunning:
			cell, err = h.serviceClient.CellById(logger, lrp.CellId)
			if err != nil {
				bbsErr := models.ConvertError(err)
				if bbsErr.Type == models.Error_ResourceNotFound {
					err = h.db.RemoveActualLRP(logger, lrp.ProcessGuid, lrp.Index, &lrp.ActualLRPInstanceKey)
					if err == nil {
						go h.actualHub.Emit(models.NewActualLRPRemovedEvent(lrpGroup))
					}
				}
				return err
			}

			var client rep.Client
			client, err = h.repClientFactory.CreateClient(cell.RepAddress, cell.RepUrl)
			if err != nil {
				return err
			}
			err = client.StopLRPInstance(logger, lrp.ActualLRPKey, lrp.ActualLRPInstanceKey)
		}

		// if there is a suspect actual lrp, what do? remove it?
		// or just notice its existence and have some sort of a retry logic?
		if lrpGroup.Suspect != nil {
			// suspectActualLRP := lrpGroup.Suspect
			// logger = logger.Session("found-suspect", lager.Data{"guid": suspectActualLRP.ProcessGuid, "index": suspectActualLRP.Index, "instance-guid": suspectActualLRP.ActualLRPInstanceKey.InstanceGuid})
			// err = h.db.RemoveActualLRP(logger, suspectActualLRP.ProcessGuid, suspectActualLRP.Index, &suspectActualLRP.ActualLRPInstanceKey)
			logger.Info("suspect-actual-lrp-found")
		} else if err == nil {
			return nil
		}

		logger.Error("retrying-failed-retire-of-actual-lrp", err, lager.Data{"attempt": retryCount + 1})
	}

	return err
}
