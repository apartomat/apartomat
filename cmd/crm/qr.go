package main

import (
	"image/color"
	"log/slog"
	"net/http"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/store/projectpages"
	qrsvg "github.com/wamuir/svg-qr-code"
)

func Qr(crm *crm.CRM) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			projectID = r.URL.Query().Get("id")
		)

		if projectID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		page, err := crm.ProjectPages.Get(r.Context(), projectpages.ProjectIDIn(projectID))
		if err != nil {
			slog.Error("can't get project page", slog.Any("err", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		qr, err := qrsvg.New(page.URL)
		if err != nil {
			slog.Error("can't create QR code", slog.Any("err", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var (
			c = color.RGBA{
				R: 53,
				G: 104,
				B: 47,
				A: 255,
			}

			back = color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 0,
			}
		)

		qr.Borderwidth = 0
		qr.Code.ForegroundColor = c
		qr.Code.BackgroundColor = back

		w.Header().Set("content-type", "image/svg+xml")

		if _, err := w.Write([]byte(qr.String())); err != nil {
			slog.Error("can't write image", slog.Any("err", err))
			return
		}
	}
}
