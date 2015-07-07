package handlers

import (
	"net/http"

	"github.com/cloudfoundry-incubator/bbs"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/rata"
)

func New(db db.DB, logger lager.Logger) http.Handler {
	domainHandler := NewDomainHandler(db, logger)
	actualLRPHandler := NewActualLRPHandler(db, logger)

	actions := rata.Handlers{
		// Domains
		bbs.DomainsRoute:      route(domainHandler.GetAll),
		bbs.UpsertDomainRoute: route(domainHandler.Upsert),

		// Actual LRPs
		bbs.ActualLRPGroupsRoute:                     route(actualLRPHandler.ActualLRPGroups),
		bbs.ActualLRPGroupsByProcessGuidRoute:        route(actualLRPHandler.ActualLRPGroupsByProcessGuid),
		bbs.ActualLRPGroupByProcessGuidAndIndexRoute: route(actualLRPHandler.ActualLRPGroupByProcessGuidAndIndex),
	}

	handler, err := rata.NewRouter(bbs.Routes, actions)
	if err != nil {
		panic("unable to create router: " + err.Error())
	}

	return LogWrap(handler, logger)
}

func route(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(f)
}
