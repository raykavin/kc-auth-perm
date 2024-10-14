package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"keycloak-example/confs"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/goterm/term"
)

type (
	claims struct {
		GivenName   string      `json:"given_name"`
		RealmAccess realmAccess `json:"realm_access"`
	}

	realmAccess struct {
		Roles []string `json:"roles"`
	}

	authorizationFailedStruct struct {
		Status   string `json:"status"`
		HTTPCode int    `json:"httpCode"`
		Message  string `json:"message"`
	}
)

const (
	authHeader         = "Authorization"
	tokenPrefix        = "Bearer"
	invalidTokenMsg    = "invalid authorization token"
	providerErrMsg     = "authorization failed while getting the provider"
	verificationErrMsg = "authorization failed while verifying the token"
	claimsErrMsg       = "error parsing claims"
	roleErrMsg         = "user not allowed to access this resource"
)

// IsAuthorized is a middleware that checks if a user has the required role to access a resource.
// It retrieves and verifies the user's token and checks if the token contains the required role.
func IsAuthorized(role string, nextFunc http.HandlerFunc) http.HandlerFunc {
	var (
		realmUrl = confs.Default.Application.OIDC.ConfigurationUrl
		clientID = confs.Default.Application.OIDC.ClientID
		client   = &http.Client{Timeout: 30 * time.Second}

		oidcConfig = &oidc.Config{
			ClientID:                   clientID,
			SkipIssuerCheck:            false,
			SkipClientIDCheck:          false,
			InsecureSkipSignatureCheck: false,
			SkipExpiryCheck:            false,
		}
	)

	return func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken, err := extractBearerToken(r)
		if err != nil {
			authorizationFailed(err.Error(), w)
			return
		}

		// Get OIDC provider context using the client's HTTP context
		ctx := oidc.ClientContext(r.Context(), client)

		// Retrieve the OIDC provider using the realm URL
		provider, err := oidc.NewProvider(ctx, realmUrl)
		if err != nil {
			authorizationFailed(fmt.Sprintf("%s: %s", providerErrMsg, err.Error()), w)
			return
		}

		// Create a token verifier using the OIDC provider and configuration
		verifier := provider.Verifier(oidcConfig)

		// Verify the token
		idToken, err := verifier.Verify(ctx, rawAccessToken)
		if err != nil {
			authorizationFailed(fmt.Sprintf("%s: %s", verificationErrMsg, err.Error()), w)
			return
		}

		// Parse and validate the claims in the token
		idTokenClaims := new(claims)
		if err := idToken.Claims(&idTokenClaims); err != nil {
			authorizationFailed(fmt.Sprintf("%s: %s", claimsErrMsg, err.Error()), w)
			return
		}

		log.Println(term.Bluef("checking if user: %s contains a role permission: %s", idTokenClaims.GivenName, role))

		// Check if the user has the required role in their token
		if !slices.Contains(idTokenClaims.RealmAccess.Roles, role) {
			log.Println(term.Redf("user: %s does not contains a role permission: %s and access to the resource blocked", idTokenClaims.GivenName, role))
			authorizationFailed(roleErrMsg, w)
			return
		}

		log.Println(term.Greenf("access granted to %s to user: %s", r.URL.Path, idTokenClaims.GivenName))

		// If authorized, proceed to the next handler
		nextFunc(w, r)
	}
}

// extractBearerToken extracts the Bearer token from the Authorization header of the HTTP request.
// It validates the format of the token and returns an error if it is invalid.
func extractBearerToken(r *http.Request) (string, error) {
	authHeaderVal := r.Header.Get(authHeader)
	if authHeaderVal == "" {
		return "", errors.New(invalidTokenMsg)
	}

	// Split the authorization header value into two parts (e.g., "Bearer token")
	parts := strings.Split(authHeaderVal, " ")
	if len(parts) != 2 || parts[0] != tokenPrefix {
		return "", errors.New(invalidTokenMsg)
	}

	return parts[1], nil
}

// authorizationFailed sends an HTTP 401 Unauthorized response when authorization fails.
// It formats the failure response as a JSON object with a status, HTTP code, and message.
func authorizationFailed(message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	data := authorizationFailedStruct{
		Status:   "FAILED",
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	}
	res, _ := json.Marshal(data)
	w.Write(res)
}
