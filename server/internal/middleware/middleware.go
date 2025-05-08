package middleware

import (
	"net"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/time/rate"

	"execute/internal/handlers/auth"
)

// 60 requests for 30 minutes
var limiterStore = NewIPRateLimiter(60.0/1800, 60)

// IPRateLimiter implements a simple per-IP rate limiter
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.Mutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter creates a new IPRateLimiter with the given parameters
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.Mutex{},
		r:   r,
		b:   b,
	}
}

// getLimiter returns the rate limiter for the given IP, creating one if necessary
func (i *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
	}
	return limiter
}

// getIP extracts the client's real IP address from the request
// It first checks the X-Forwarded-For header and falls back to RemoteAddr
func getIP(r *http.Request) string {
	// If behind a proxy, the real IP might be in the X-Forwarded-For header
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		parts := strings.Split(xff, ",")
		ip := strings.TrimSpace(parts[0])
		if net.ParseIP(ip) != nil {
			return ip
		}
	}
	// If no X-Forwarded-For header, use RemoteAddr (strip the port if present)
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

// rateLimitMiddleware is a middleware that checks if the request
// from an IP is allowed to proceed
func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := getIP(r)
		limiter := limiterStore.getLimiter(ip)
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// SetContentTypeMiddleware is a middleware that sets the Content-Type header to application/json
func setContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Helper function to apply multiple middlewares to a handler
func ApplyMiddlewares(handler http.Handler) http.Handler {
	return rateLimitMiddleware(setContentTypeMiddleware(handler))
}

// authMiddleware is a middleware that checks for a valid session cookie.
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		username, exists := auth.GetSessionUsername(cookie.Value)
		if !exists || strings.TrimSpace(username) == "" {
			http.NotFound(w, r)
			return
		}
		// You can set the username in the request context here if needed.
		next.ServeHTTP(w, r)
	})
}

// Helper function to apply multiple middlewares to a handler
func ApplyAuthMiddlewares(handler http.Handler) http.Handler {
	return rateLimitMiddleware(authMiddleware(setContentTypeMiddleware(handler)))
}

// corsMiddleware sets CORS headers to allow corss-origin requests
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
