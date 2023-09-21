package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"io"
	"log"
	"net/http"

	"strings"

	"github.com/Tito-74/Savannah-e-shop/database"
	"github.com/Tito-74/Savannah-e-shop/models"
	"github.com/joho/godotenv"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	database.DatabaseConnect()
	configURL := os.Getenv("KEYCLOAK_CONFIG_URL")
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		panic(err)
	}

	clientID := os.Getenv("KEYCLOAK_CLIENTID")
	clientSecret := os.Getenv("KEYCLOAK_CLIENT_SECRET")

	fmt.Println("secret", clientSecret)

	redirectURL := os.Getenv("KEYCLOAK_REDIRECT_URL")
	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state := os.Getenv("KEYCLOAK_STATE")

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get("Authorization")
		if rawAccessToken == "" {
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			fmt.Print("parts: ", len(parts))
			w.WriteHeader(400)
			return
		}
		_, err := verifier.Verify(ctx, parts[1])

		if err != nil {
			fmt.Print("error: ", err)
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}
	})

	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
			return
		}
		database := database.Database.Db
		rawAccessToken := r.Header.Get("Authorization")

		if rawAccessToken == "" {
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			fmt.Print("parts: ", len(parts))
			w.WriteHeader(400)
			return
		}
		_, err := verifier.Verify(ctx, parts[1])

		if err != nil {
			fmt.Print("error: ", err)
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}

		var customer models.Customer
		json.Unmarshal(body, &customer)

		// add to the customers table
		err = database.Debug().Create(&customer).Error
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
				return
			}
			// database := database.Database.Db
			rawAccessToken := r.Header.Get("Authorization")

			if rawAccessToken == "" {
				http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
				return
			}

			parts := strings.Split(rawAccessToken, " ")
			if len(parts) != 2 {
				fmt.Print("parts: ", len(parts))
				w.WriteHeader(400)
				return
			}
			_, err := verifier.Verify(ctx, parts[1])

			if err != nil {
				fmt.Print("error: ", err)
				http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
				return
			}

			body, err := io.ReadAll(r.Body)

			if err != nil {
				log.Fatalln(err)
			}

			var order models.Orders
			json.Unmarshal(body, &order)

			// Append to the order table
			if result := database.Create(&order); result.Error != nil {
				fmt.Println(result.Error)
			}
			// send message
			var customer models.Customer
			database.Find(&customer, "id =?", order.CustomerId)
			phone := customer.Phone
			name := customer.Name
			res := SendMessage(name, phone, &order)

			log.Println("response", res)

			// Send a 201 created response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Created successfully")

		})

		// Send a 201 created response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Created successfully")
	})

	http.HandleFunc("/demo/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	log.Fatal(http.ListenAndServe("localhost:8181", nil))
}
