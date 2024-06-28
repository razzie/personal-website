package main

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w *GzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *GzipResponseWriter) WriteHeader(statusCode int) {
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(statusCode)
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz := gzip.NewWriter(w)
		defer gz.Close()

		gw := &GzipResponseWriter{ResponseWriter: w, Writer: gz}
		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gw, r)
	})
}
