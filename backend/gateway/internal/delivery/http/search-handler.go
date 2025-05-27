package http

import (
	"context"
	"fmt"
	"net/http"

	"quickflow/gateway/internal/delivery/http/forms"
	"quickflow/gateway/internal/errors"
	http2 "quickflow/gateway/utils/http"
	"quickflow/shared/logger"
	"quickflow/shared/models"
)

type SearchUseCase interface {
	SearchSimilarUser(ctx context.Context, toSearch string, postsCount uint) ([]models.PublicUserInfo, error)
}

type SearchHandler struct {
	searchUseCase    SearchUseCase
	communityService CommunityService
	profileService   ProfileUseCase
}

func NewSearchHandler(searchUseCase SearchUseCase, communityService CommunityService, profileService ProfileUseCase) *SearchHandler {
	return &SearchHandler{
		searchUseCase:    searchUseCase,
		communityService: communityService,
		profileService:   profileService,
	}
}

func (s *SearchHandler) SearchSimilarUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var searchForm forms.SearchForm
	err := searchForm.Unpack(r.URL.Query())
	if err != nil {
		logger.Error(ctx, "Failed to decode request body for user search: "+err.Error())
		http2.WriteJSONError(w, errors2.New(errors2.BadRequestErrorCode, "Failed to decode request body", http.StatusBadRequest))
		return
	}

	users, err := s.searchUseCase.SearchSimilarUser(ctx, searchForm.ToSearch, searchForm.Count)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Failed to search similar users: %s", err.Error()))
		http2.WriteJSONError(w, err)
		return
	}

	var publicUsersInfoOut []forms.PublicUserInfoOut
	for _, user := range users {
		publicUsersInfoOut = append(publicUsersInfoOut, forms.PublicUserInfoToOut(user, ""))
	}

	out := forms.PayloadWrapper[[]forms.PublicUserInfoOut]{Payload: publicUsersInfoOut}

	w.Header().Set("Content-Type", "application/json")
	js, err := out.MarshalJSON()
	if err != nil {
		logger.Error(ctx, "Failed to marshal json payload", err)
		http2.WriteJSONError(w, err)
		return
	}
	if _, err = w.Write(js); err != nil {
		logger.Error(ctx, "Failed to encode feedback output", err)
		http2.WriteJSONError(w, errors2.New(errors2.InternalErrorCode, "Failed to encode feedback output", http.StatusInternalServerError))
	}
}

func (s *SearchHandler) SearchSimilarCommunities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var searchForm forms.SearchForm
	err := searchForm.Unpack(r.URL.Query())
	if err != nil {
		logger.Error(ctx, "Failed to decode request body for user search: "+err.Error())
		http2.WriteJSONError(w, errors2.New(errors2.BadRequestErrorCode, "Failed to decode request body", http.StatusBadRequest))
		return
	}

	communities, err := s.communityService.SearchSimilarCommunities(ctx, searchForm.ToSearch, int(searchForm.Count))
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Failed to search similar communities: %s", err.Error()))
		http2.WriteJSONError(w, err)
		return
	}

	communitiesOut := make([]forms.CommunityForm, len(communities))
	for i, community := range communities {
		info, err := s.profileService.GetPublicUserInfo(ctx, community.OwnerID)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Failed to get user info: %s", err.Error()))
			http2.WriteJSONError(w, err)
		}
		communitiesOut[i] = forms.ToCommunityForm(*community, info)
	}

	out := forms.PayloadWrapper[[]forms.CommunityForm]{Payload: communitiesOut}

	w.Header().Set("Content-Type", "application/json")
	js, err := out.MarshalJSON()
	if err != nil {
		logger.Error(ctx, "Failed to marshal json payload", err)
		http2.WriteJSONError(w, err)
		return
	}
	if _, err = w.Write(js); err != nil {
		logger.Error(ctx, "Failed to encode feedback output", err)
		http2.WriteJSONError(w, errors2.New(errors2.InternalErrorCode, "Failed to encode feedback output", http.StatusInternalServerError))
	}
}
