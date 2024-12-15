package handler

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"gcommons/otel"
)

func RollDice(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "RollDice")
	defer span.End()

	_, span2 := otel.Tracer.Start(ctx, "random number")
	roll := 1 + rand.Intn(6)
	defer span2.End()

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(w, resp); err != nil {
		log.Printf("Write failed: %v\n", err)
	}
}
