package middleware

import (
    "context"
    "net/http"
    "strings"
    firebase "firebase.google.com/go/v4"
)

// AuthMiddleware maneja la autenticación usando Firebase Auth
type AuthMiddleware struct {
    app *firebase.App
}

// NewAuthMiddleware crea una nueva instancia del middleware de autenticación
func NewAuthMiddleware(app *firebase.App) *AuthMiddleware {
    return &AuthMiddleware{
        app: app,
    }
}

// Authenticate verifica el token JWT de Firebase
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        idToken := strings.TrimPrefix(authHeader, "Bearer ")
        auth, err := m.app.Auth(r.Context())
        if err != nil {
            http.Error(w, "Internal error", http.StatusInternalServerError)
            return
        }

        token, err := auth.VerifyIDToken(r.Context(), idToken)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Añadir información del usuario al contexto
        ctx := context.WithValue(r.Context(), "user_id", token.UID)
        ctx = context.WithValue(ctx, "user_email", token.Claims["email"])
        ctx = context.WithValue(ctx, "user_role", token.Claims["role"])
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// GetUserFromContext obtiene la información del usuario del contexto
func GetUserFromContext(ctx context.Context) (userId string, userEmail string, userRole string) {
    userId, _ = ctx.Value("user_id").(string)
    userEmail, _ = ctx.Value("user_email").(string)
    userRole, _ = ctx.Value("user_role").(string)
    return
}
