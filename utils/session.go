package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
)

// Takes the JWT token from the request header and verifies it using keyfunc
func VerifySession(next http.Handler) http.Handler {
    jwksURL := os.Getenv("JWKS_URL")

    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        // Retrieve the JWT token from the request header
        token := strings.Replace( r.Header.Get("Authorization"), "Bearer ", "", 1)
        if token == "" {
            render.Render(w, r, GetHttpErrorFromError(ErrNoToken, "resource"))
            return
        }
        
        // Create a new keyfunc instance
        k, keyFuncErr := keyfunc.NewDefaultCtx(r.Context(), []string{jwksURL})
        if keyFuncErr != nil {
            log.Println("Failed to create keyfunc: ", keyFuncErr)
            render.Render(w, r, GetHttpErrorFromError(ErrBadToken, "resource"))
            return
        }

        log.Println("Token: ", token)
        // Parse the JWT
        parsed, err := jwt.Parse(token, k.Keyfunc);

        if err != nil {
            log.Println("Failed to parse token: ", err)
            
            render.Render(w, r, GetHttpErrorFromError(ErrBadToken, "resource"))
            return
        }

        uid:= parsed.Claims.(jwt.MapClaims)["sub"]
        if(uid == nil) {
            render.Render(w, r, GetHttpErrorFromError(ErrBadToken, "resource"))
            return
        }
        log.Println(uid)
        // Add the uid to the request context
        ctx := context.WithValue(r.Context(), "uid", uid)

        // Call the handler function
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func GetUidFromContext(ctx context.Context) string {
    uid := ctx.Value("uid")
    if uid == nil {
        return ""
    }
    return uid.(string)
}