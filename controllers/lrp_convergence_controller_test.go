package controllers_test

import (
	"errors"

	"code.cloudfoundry.org/auctioneer"
	"code.cloudfoundry.org/auctioneer/auctioneerfakes"
	"code.cloudfoundry.org/bbs/controllers"
	"code.cloudfoundry.org/bbs/controllers/fakes"
	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/db/dbfakes"
	"code.cloudfoundry.org/bbs/events/eventfakes"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/bbs/models/test/model_helpers"
	"code.cloudfoundry.org/bbs/serviceclient/serviceclientfakes"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	"code.cloudfoundry.org/rep/repfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("LRP Convergence Controllers", func() {
	var (
		err                  error
		logger               *lagertest.TestLogger
		fakeLRPDB            *dbfakes.FakeLRPDB
		actualHub            *eventfakes.FakeHub
		retirer              *fakes.FakeRetirer
		fakeAuctioneerClient *auctioneerfakes.FakeClient

		keysToAuction        []*auctioneer.LRPStartRequest
		keysToRetire         []*models.ActualLRPKey
		keysWithMissingCells []*models.ActualLRPKeyWithSchedulingInfo

		retiringActualLRP1 *models.ActualLRP
		retiringActualLRP2 *models.ActualLRP

		desiredLRP1, desiredLRP2 models.DesiredLRPSchedulingInfo
		cellID                   string
		cellSet                  models.CellSet

		controller *controllers.LRPConvergenceController
	)

	BeforeEach(func() {
		fakeLRPDB = new(dbfakes.FakeLRPDB)
		fakeAuctioneerClient = new(auctioneerfakes.FakeClient)
		logger = lagertest.NewTestLogger("test")

		// request1 := auctioneer.NewLRPStartRequestFromModel(model_helpers.NewValidDesiredLRP("to-auction-1"), 1, 2)
		// request2 := auctioneer.NewLRPStartRequestFromModel(model_helpers.NewValidDesiredLRP("to-auction-2"), 0, 4)
		// keysToAuction = []*auctioneer.LRPStartRequest{&request1, &request2}

		retiringActualLRP1 = model_helpers.NewValidActualLRP("to-retire-1", 0)
		retiringActualLRP2 = model_helpers.NewValidActualLRP("to-retire-2", 1)
		keysToRetire = []*models.ActualLRPKey{&retiringActualLRP1.ActualLRPKey, &retiringActualLRP2.ActualLRPKey}

		desiredLRP1 = model_helpers.NewValidDesiredLRP("to-unclaim-1").DesiredLRPSchedulingInfo()
		desiredLRP2 = model_helpers.NewValidDesiredLRP("to-unclaim-2").DesiredLRPSchedulingInfo()

		cellID = "cell-id"
		instanceKey := models.NewActualLRPInstanceKey("instance-guid", cellID)

		retiringActualLRP1.CellId = cellID
		retiringActualLRP1.ActualLRPInstanceKey = instanceKey
		retiringActualLRP1.State = models.ActualLRPStateClaimed
		group1 := &models.ActualLRPGroup{Instance: retiringActualLRP1}

		retiringActualLRP2.CellId = cellID
		retiringActualLRP2.ActualLRPInstanceKey = instanceKey
		retiringActualLRP2.State = models.ActualLRPStateClaimed
		group2 := &models.ActualLRPGroup{Instance: retiringActualLRP2}

		fakeLRPDB.ActualLRPGroupByProcessGuidAndIndexStub = func(_ lager.Logger, processGuid string, _ int32) (*models.ActualLRPGroup, error) {
			if processGuid == retiringActualLRP1.ProcessGuid {
				return group1, nil
			}
			if processGuid == retiringActualLRP2.ProcessGuid {
				return group2, nil
			}

			return nil, models.ErrResourceNotFound
		}

		result := db.ConvergenceResult{
			KeysToRetire:         keysToRetire,
			KeysWithMissingCells: keysWithMissingCells,
			Events:               nil,
		}
		fakeLRPDB.ConvergeLRPsReturns(result)

		logger.RegisterSink(lager.NewWriterSink(GinkgoWriter, lager.DEBUG))

		fakeServiceClient = new(serviceclientfakes.FakeServiceClient)
		fakeRepClientFactory = new(repfakes.FakeClientFactory)
		fakeRepClient = new(repfakes.FakeClient)
		fakeRepClientFactory.CreateClientReturns(fakeRepClient, nil)
		fakeServiceClient.CellByIdReturns(nil, errors.New("hi"))

		cellPresence := models.NewCellPresence("cell-id", "1.1.1.1", "", "z1", models.CellCapacity{}, nil, nil, nil, nil)
		cellSet = models.CellSet{"cell-id": &cellPresence}
		fakeServiceClient.CellsReturns(cellSet, nil)

		actualHub = &eventfakes.FakeHub{}
		retirer = &fakes.FakeRetirer{}
		controller = controllers.NewLRPConvergenceController(logger, fakeLRPDB, actualHub, fakeAuctioneerClient, fakeServiceClient, retirer, 2)
	})

	JustBeforeEach(func() {
		err = controller.ConvergeLRPs(logger)
	})

	It("calls ConvergeLRPs", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(fakeLRPDB.ConvergeLRPsCallCount()).To(Equal(1))
		_, actualCellSet := fakeLRPDB.ConvergeLRPsArgsForCall(0)
		Expect(actualCellSet).To(BeEquivalentTo(cellSet))
	})

	Context("when there are unstarted ActualLRPs", func() {
		var (
			key            *models.ActualLRPKey
			schedulingInfo models.DesiredLRPSchedulingInfo
			before, after  *models.ActualLRPGroup
		)

		BeforeEach(func() {
			lrp := model_helpers.NewValidDesiredLRP("some-guid")
			schedulingInfo = lrp.DesiredLRPSchedulingInfo()
			key = &models.ActualLRPKey{
				ProcessGuid: lrp.ProcessGuid,
				Index:       0,
				Domain:      lrp.Domain,
			}
			lrpKeyWithSchedulingInfo := &models.ActualLRPKeyWithSchedulingInfo{
				Key:            key,
				SchedulingInfo: &schedulingInfo,
			}
			fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{
				UnstartedLRPKeys: []*models.ActualLRPKeyWithSchedulingInfo{lrpKeyWithSchedulingInfo},
			})

			crashedLRP := model_helpers.NewValidActualLRP("some-guid", 0)
			crashedLRP.State = models.ActualLRPStateCrashed
			before = &models.ActualLRPGroup{Instance: crashedLRP}
			unclaimedLRP := *crashedLRP
			unclaimedLRP.State = models.ActualLRPStateUnclaimed
			after = &models.ActualLRPGroup{Instance: &unclaimedLRP}
			fakeLRPDB.UnclaimActualLRPReturns(before, after, nil)
		})

		It("auctions off the returned keys", func() {
			Expect(fakeAuctioneerClient.RequestLRPAuctionsCallCount()).To(Equal(1))

			_, startAuctions := fakeAuctioneerClient.RequestLRPAuctionsArgsForCall(0)
			Expect(startAuctions).To(HaveLen(1))
			request := auctioneer.NewLRPStartRequestFromModel(model_helpers.NewValidDesiredLRP("some-guid"), 0)
			Expect(startAuctions).To(ContainElement(&request))
		})

		It("transition the LPR to UNCLAIMED state", func() {
			Eventually(fakeLRPDB.UnclaimActualLRPCallCount).Should(Equal(1))
			_, actualKey := fakeLRPDB.UnclaimActualLRPArgsForCall(0)
			Expect(actualKey).To(Equal(key))
		})

		It("emits an LRPChanged event", func() {
			Eventually(actualHub.EmitCallCount).Should(Equal(1))
			event := actualHub.EmitArgsForCall(0)
			Expect(event).To(Equal(models.NewActualLRPChangedEvent(before, after)))
		})
	})

	Context("when there are missing ActualLRPs", func() {
		var (
			key            *models.ActualLRPKey
			schedulingInfo models.DesiredLRPSchedulingInfo
			after          *models.ActualLRPGroup
		)

		BeforeEach(func() {
			lrp := model_helpers.NewValidDesiredLRP("some-guid")
			schedulingInfo = lrp.DesiredLRPSchedulingInfo()
			key = &models.ActualLRPKey{
				ProcessGuid: lrp.ProcessGuid,
				Index:       0,
				Domain:      lrp.Domain,
			}
			lrpKeyWithSchedulingInfo := &models.ActualLRPKeyWithSchedulingInfo{
				Key:            key,
				SchedulingInfo: &schedulingInfo,
			}
			fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{
				MissingLRPKeys: []*models.ActualLRPKeyWithSchedulingInfo{lrpKeyWithSchedulingInfo},
			})

			unclaimedLRP := model_helpers.NewValidActualLRP("some-guid", 0)
			unclaimedLRP.State = models.ActualLRPStateUnclaimed
			after = &models.ActualLRPGroup{Instance: unclaimedLRP}
			fakeLRPDB.CreateUnclaimedActualLRPReturns(after, nil)
		})

		It("auctions off the returned keys", func() {
			Expect(fakeAuctioneerClient.RequestLRPAuctionsCallCount()).To(Equal(1))

			_, startAuctions := fakeAuctioneerClient.RequestLRPAuctionsArgsForCall(0)
			Expect(startAuctions).To(HaveLen(1))
			request := auctioneer.NewLRPStartRequestFromModel(model_helpers.NewValidDesiredLRP("some-guid"), 0)
			Expect(startAuctions).To(ContainElement(&request))
		})

		It("creates the LPR record in the database", func() {
			Eventually(fakeLRPDB.CreateUnclaimedActualLRPCallCount).Should(Equal(1))
			_, actualKey := fakeLRPDB.CreateUnclaimedActualLRPArgsForCall(0)
			Expect(actualKey).To(Equal(key))
		})

		It("emits a LPRCreated event", func() {
			Eventually(actualHub.EmitCallCount).Should(Equal(1))
			event := actualHub.EmitArgsForCall(0)
			Expect(event).To(Equal(models.NewActualLRPCreatedEvent(after)))
		})
	})

	Context("when fetching the cells fails", func() {
		BeforeEach(func() {
			fakeServiceClient.CellsReturns(nil, errors.New("kaboom"))
		})

		It("does not return an error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("does not call ConvergeLRPs", func() {
			Expect(fakeLRPDB.ConvergeLRPsCallCount()).To(Equal(0))
		})

		It("logs the error", func() {
			Eventually(logger).Should(gbytes.Say("failed-listing-cells"))
		})
	})

	Context("when fetching the cells returns ErrResourceNotFound", func() {
		BeforeEach(func() {
			fakeServiceClient.CellsReturns(nil, models.ErrResourceNotFound)
		})

		It("calls ConvergeLRPs with an empty CellSet", func() {
			Expect(err).NotTo(HaveOccurred())
			Expect(fakeLRPDB.ConvergeLRPsCallCount()).To(Equal(1))
			_, actualCellSet := fakeLRPDB.ConvergeLRPsArgsForCall(0)
			Expect(actualCellSet).To(BeEquivalentTo(models.CellSet{}))
		})
	})

	Context("when no lrps to auction", func() {
		BeforeEach(func() {
			fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{})
		})

		It("doesn't start the auctions", func() {
			Consistently(fakeAuctioneerClient.RequestLRPAuctionsCallCount).Should(Equal(0))
		})
	})

	Context("when there are suspect LRPs with existing cells", func() {
		var (
			suspectActualLRP             *models.ActualLRP
			ordinaryActualLRP            *models.ActualLRP
			suspectKeysWithExistingCells []*models.ActualLRPKey
		)

		BeforeEach(func() {
			ordinaryActualLRP = model_helpers.NewValidActualLRP("suspect-1", 0)
			suspectActualLRP = model_helpers.NewValidActualLRP("suspect-1", 0)
			group := &models.ActualLRPGroup{Instance: suspectActualLRP}

			suspectKeysWithExistingCells = []*models.ActualLRPKey{&suspectActualLRP.ActualLRPKey}

			fakeLRPDB.ActualLRPGroupByProcessGuidAndIndexReturns(&models.ActualLRPGroup{Instance: ordinaryActualLRP}, nil)
			fakeLRPDB.ChangeActualLRPPresenceReturns(group, group, nil)
			fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{
				SuspectKeysWithExistingCells: suspectKeysWithExistingCells,
				KeysWithMissingCells:         keysWithMissingCells,
			})
		})

		It("remove the Ordinary LRP", func() {
			Eventually(fakeLRPDB.RemoveActualLRPCallCount).Should(Equal(1))

			_, guid, index, key := fakeLRPDB.RemoveActualLRPArgsForCall(0)

			Expect(guid).To(Equal(ordinaryActualLRP.ProcessGuid))
			Expect(index).To(Equal(ordinaryActualLRP.Index))
			Expect(key).To(BeNil())
		})

		It("changes the suspect LRP presence to Ordinary", func() {
			Eventually(fakeLRPDB.ChangeActualLRPPresenceCallCount).Should(Equal(1))
			_, lrpKey, presence := fakeLRPDB.ChangeActualLRPPresenceArgsForCall(0)

			Expect(lrpKey).To(Equal(&suspectActualLRP.ActualLRPKey))
			Expect(presence).To(Equal(models.ActualLRP_Ordinary))
		})

		It("does not emit any events", func() {
			Consistently(actualHub.EmitCallCount).Should(Equal(0))
		})

		Context("when the suspect lrp cannot be changed to ordinary lrp", func() {
			Context("and the error is unrecoverable", func() {
				BeforeEach(func() {
					fakeLRPDB.ChangeActualLRPPresenceReturns(nil, nil, models.NewUnrecoverableError(nil))
				})

				It("logs the error", func() {
					Eventually(logger).Should(gbytes.Say("unrecoverable-error"))
				})

				It("returns the error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err).Should(Equal(models.NewUnrecoverableError(nil)))
				})
			})
		})

		Context("when the ordinary lrp cannot be removed", func() {
			BeforeEach(func() {
				fakeLRPDB.RemoveActualLRPReturns(errors.New("booom!"))
			})

			It("does not change the suspect LRP presence", func() {
				Consistently(fakeLRPDB.ChangeActualLRPPresenceCallCount).Should(BeZero())
			})

			Context("and the error is unrecoverable", func() {
				BeforeEach(func() {
					fakeLRPDB.RemoveActualLRPReturns(models.NewUnrecoverableError(nil))
				})

				It("logs the error", func() {
					Eventually(logger).Should(gbytes.Say("unrecoverable-error"))
				})

				It("returns the error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err).Should(Equal(models.NewUnrecoverableError(nil)))
				})
			})
		})
	})

	Context("when there are LRPs with missing cells", func() {
		var (
			suspectActualLRP1 *models.ActualLRP
			suspectActualLRP2 *models.ActualLRP
		)

		BeforeEach(func() {
			suspectActualLRP1 = model_helpers.NewValidActualLRP("to-unclaim-1", 0)
			suspectActualLRP2 = model_helpers.NewValidActualLRP("to-unclaim-2", 1)

			fakeLRPDB.ChangeActualLRPPresenceStub = func(_ lager.Logger, key *models.ActualLRPKey, _ models.ActualLRP_Presence) (*models.ActualLRPGroup, *models.ActualLRPGroup, error) {
				if key.ProcessGuid == suspectActualLRP1.ProcessGuid {
					return &models.ActualLRPGroup{Instance: suspectActualLRP1},
						&models.ActualLRPGroup{Instance: suspectActualLRP1}, nil
				}
				if key.ProcessGuid == suspectActualLRP2.ProcessGuid {
					return &models.ActualLRPGroup{Instance: suspectActualLRP2},
						&models.ActualLRPGroup{Instance: suspectActualLRP2}, nil
				}
				return nil, nil, models.ErrResourceNotFound
			}

			keysWithMissingCells = []*models.ActualLRPKeyWithSchedulingInfo{
				{Key: &suspectActualLRP1.ActualLRPKey, SchedulingInfo: &desiredLRP1},
				{Key: &suspectActualLRP2.ActualLRPKey, SchedulingInfo: &desiredLRP2},
			}
			fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{
				KeysWithMissingCells: keysWithMissingCells,
			})
		})

		It("change the LRP presence to 'Suspect' and auctions a new actual lrp", func() {
			Eventually(fakeLRPDB.ChangeActualLRPPresenceCallCount).Should(Equal(2))

			unclaimedKeys := []*models.ActualLRPKey{}
			for i := 0; i < fakeLRPDB.ChangeActualLRPPresenceCallCount(); i++ {
				_, key, presence := fakeLRPDB.ChangeActualLRPPresenceArgsForCall(i)
				Expect(presence).To(Equal(models.ActualLRP_Suspect))
				unclaimedKeys = append(unclaimedKeys, key)
			}
			Expect(unclaimedKeys).To(ContainElement(&suspectActualLRP1.ActualLRPKey))
			Expect(unclaimedKeys).To(ContainElement(&suspectActualLRP2.ActualLRPKey))
		})

		It("auctions new lrps", func() {
			Expect(fakeAuctioneerClient.RequestLRPAuctionsCallCount()).To(Equal(1))

			unclaimedStartRequest1 := auctioneer.NewLRPStartRequestFromSchedulingInfo(&desiredLRP1, 0)
			unclaimedStartRequest2 := auctioneer.NewLRPStartRequestFromSchedulingInfo(&desiredLRP2, 1)

			keysToAuction := []*auctioneer.LRPStartRequest{&unclaimedStartRequest1, &unclaimedStartRequest2}

			_, startAuctions := fakeAuctioneerClient.RequestLRPAuctionsArgsForCall(0)
			Expect(startAuctions).To(HaveLen(2))
			Expect(startAuctions).To(ConsistOf(keysToAuction))
		})

		It("emits the proper events", func() {
			Consistently(actualHub.EmitCallCount).Should(Equal(0))
		})


		})

		Context("when the DB returns an unrecoverable error", func() {
			BeforeEach(func() {
				fakeLRPDB.ChangeActualLRPPresenceReturns(nil, nil, models.NewUnrecoverableError(nil))
			})

			It("logs the error", func() {
				Eventually(logger).Should(gbytes.Say("unrecoverable-error"))
			})

			It("returns the error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).Should(Equal(models.NewUnrecoverableError(nil)))
			})
		})

		Context("when changing the actual lrp presence fails", func() {
			BeforeEach(func() {
				fakeLRPDB.ChangeActualLRPPresenceReturns(nil, nil, errors.New("terrrible"))
				fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{
					KeysWithMissingCells: keysWithMissingCells,
				})
			})

			XIt("auctions off the returned keys", func() {
				// how can it auction off keys if it returns before adding anything to the startRequests?
				Expect(fakeAuctioneerClient.RequestLRPAuctionsCallCount()).To(Equal(1))

				_, startAuctions := fakeAuctioneerClient.RequestLRPAuctionsArgsForCall(0)
				Expect(startAuctions).To(HaveLen(2))
				Expect(startAuctions).To(ConsistOf(keysToAuction))
			})

			It("does not emit change events", func() {
				Eventually(fakeLRPDB.ChangeActualLRPPresenceCallCount).Should(Equal(2))
				Consistently(actualHub.EmitCallCount).Should(Equal(0))
			})
		})
	})

	Describe("stopping extra LRPs", func() {
		var (
			cellPresence models.CellPresence
		)

		Context("when the cell", func() {
			Context("is present", func() {
				Context("and the extra lrp is suspect", func() {
					var (
						key            *models.ActualLRPKey
						schedulingInfo models.DesiredLRPSchedulingInfo
						/*before,*/ after      *models.ActualLRPGroup
						runningLRP, suspectLRP *models.ActualLRP
					)
					BeforeEach(func() {
						lrp := model_helpers.NewValidDesiredLRP("some-guid")
						schedulingInfo = lrp.DesiredLRPSchedulingInfo()
						key = &models.ActualLRPKey{
							ProcessGuid: lrp.ProcessGuid,
							Index:       0,
							Domain:      lrp.Domain,
						}
						/*lrpKeyWithSchedulingInfo := &models.ActualLRPKeyWithSchedulingInfo{
							Key:            key,
							SchedulingInfo: &schedulingInfo,
						}*/
						// convergence should return the following
						fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{

							SuspectKeysWithPresentCells: []*models.ActualLRPKey{key},
						})
						///
						runningLRP = model_helpers.NewValidActualLRP("some-guid", 1)
						runningLRP.State = models.ActualLRPStateClaimed
						runningLRP.Presence = models.ActualLRP_Ordinary
						suspectLRP = model_helpers.NewValidActualLRP("some-guid", 0)
						suspectLRP.State = models.ActualLRPStateClaimed
						suspectLRP.Presence = models.ActualLRP_Suspect
						after = &models.ActualLRPGroup{Instance: runningLRP}
						fakeLRPDB.ActualLRPGroupByProcessGuidAndIndexReturns(after, nil)
					})
					It("removes the suspect LRP", func() {
						Eventually(fakeLRPDB.RemoveActualLRPCallCount).Should(Equal(1))
						_, removedProcessGUID, removedIndex, _ := fakeLRPDB.RemoveActualLRPArgsForCall(0)
						Expect(removedIndex).To(Equal(suspectLRP.ActualLRPKey.Index))
						Expect(removedProcessGUID).To(Equal(suspectLRP.ActualLRPKey.ProcessGuid))
						// we aren't testing the ActualLRPInstanceKey that is returned...
					})
					It("emits an ActualLRPRemovedEvent containing the suspect LRP", func() {
						Eventually(actualHub.EmitCallCount).Should(Equal(1))
						event := actualHub.EmitArgsForCall(0)
						Expect(event).To(Equal(models.NewActualLRPRemovedEvent(after)))
					})

				})
				BeforeEach(func() {
					cellPresence = models.NewCellPresence(
						cellID,
						"cell1.addr",
						"",
						"the-zone",
						models.NewCellCapacity(128, 1024, 6),
						[]string{},
						[]string{},
						[]string{},
						[]string{},
					)

					fakeServiceClient.CellByIdReturns(&cellPresence, nil)
				})

				It("stops the LRPs", func() {
					Eventually(retirer.RetireActualLRPCallCount()).Should(Equal(2))

					stoppedKeys := make([]*models.ActualLRPKey, 2)

					for i := 0; i < 2; i++ {
						_, key := retirer.RetireActualLRPArgsForCall(i)
						stoppedKeys[i] = key
					}

					Expect(stoppedKeys).To(ContainElement(&retiringActualLRP1.ActualLRPKey))
					Expect(stoppedKeys).To(ContainElement(&retiringActualLRP2.ActualLRPKey))
				})

				Context("when the retirer returns an error", func() {
					BeforeEach(func() {
						retirer.RetireActualLRPReturns(errors.New("BOOM!!!"))
					})

					It("should log the error", func() {
						Expect(logger.Buffer()).To(gbytes.Say("BOOM!!!"))
					})

					It("should return the error", func() {
						Expect(err).NotTo(HaveOccurred())
					})
				})
			})

			Context("is not present", func() {
				BeforeEach(func() {
					fakeServiceClient.CellByIdReturns(nil,
						&models.Error{
							Type:    models.Error_ResourceNotFound,
							Message: "cell not found",
						})
				})

				Context("removing the actualLRP succeeds", func() {
					It("removes the LRPs", func() {
						Eventually(retirer.RetireActualLRPCallCount()).Should(Equal(2))

						deletedKeys := make([]*models.ActualLRPKey, 2)

						for i := 0; i < 2; i++ {
							_, key := retirer.RetireActualLRPArgsForCall(i)
							deletedKeys[i] = key
						}

						Expect(deletedKeys).To(ContainElement(&retiringActualLRP1.ActualLRPKey))
						Expect(deletedKeys).To(ContainElement(&retiringActualLRP2.ActualLRPKey))
					})
				})
			})
		})
	})

	Context("when the db returns events", func() {
		var expectedRemovedEvent *models.ActualLRPRemovedEvent
		BeforeEach(func() {
			group1 := &models.ActualLRPGroup{Instance: model_helpers.NewValidActualLRP("evacuating-lrp", 0)}
			expectedRemovedEvent = models.NewActualLRPRemovedEvent(group1)
			events := []models.Event{expectedRemovedEvent}
			fakeLRPDB.ConvergeLRPsReturns(db.ConvergenceResult{
				Events: events,
			})
		})

		It("emits those events", func() {

			Eventually(actualHub.EmitCallCount).Should(Equal(1))
			event := actualHub.EmitArgsForCall(0)
			removedEvent, ok := event.(*models.ActualLRPRemovedEvent)
			Expect(ok).To(BeTrue())
			Expect(removedEvent).To(Equal(expectedRemovedEvent))
		})
	})

	Context("Given two containers i1 suspect container on cell1 and i2 on cell2", func() {
		Context("when cell1 is Present", func() {
	  	Context("and container i1 is in STATE RUNNING", func() {
		  	Context("and cell2 is present", func() {
			   Context("and container i2 is CREATED", func() {
			    It("removes i2 and i1 is no longer suspect", func() {})
			})
			Context("and container i2 is RUNNING", func() {
			It("removes the suspect container i1", func() {})
			It("emits an ActualLRPRemovedEvent for i1", func(){})
			})
			})
		})
		})
		Context("When cell1 is Missing", func() {
			Context("And container i1 is RUNNING", func() {
				Context("And cell2 is Present", func() {
					Context("And container i2 does not exist", func() {
						It("moves i1 to suspect", func() {})
						It("creates an unclaimed ActualLRP", func(){})
						It("Emits an ActualLRPCreated event", func() {})
					})
					Context("And container i2 is CREATED", func() {
						It("does nothing", func() {})
					})
					Context("and i2 is RUNNING", func() {
						It("removes suspect instance i1", func() {})
						It("emits an ActualLRPRemovedEvent for i1", func() {})

					})
				})
		})
		})
	})
})
