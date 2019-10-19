/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A StoreApiController binds http requests to an api service and writes the service results to the http response
type StoreApiController struct {
    service StoreApiServicer
}

// NewStoreApiController creates a default api controller
func NewStoreApiController(s StoreApiServicer) StoreApiRouter {
    return &StoreApiController{ service: s }
}

// Routes returns all of the api route for the StoreApiController
func (c *StoreApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteOrder",
			strings.ToUpper("Delete"),
			"/v2/store/order/{orderId}",
			c.DeleteOrder,
		},
		{
			"GetInventory",
			strings.ToUpper("Get"),
			"/v2/store/inventory",
			c.GetInventory,
		},
		{
			"GetOrderById",
			strings.ToUpper("Get"),
			"/v2/store/order/{orderId}",
			c.GetOrderById,
		},
		{
			"PlaceOrder",
			strings.ToUpper("Post"),
			"/v2/store/order",
			c.PlaceOrder,
		},
	}
}

// DeleteOrder - Delete purchase order by ID
func (c *StoreApiController) DeleteOrder(w http.ResponseWriter, r *http.Request) { 
    params := mux.Vars(r)
	orderId := params["orderId"]
    result, err := c.service.DeleteOrder(orderId)
    if err != nil {
        w.WriteHeader(500)
        return
    }
	
    EncodeJSONResponse(result, nil,  w)
}

// GetInventory - Returns pet inventories by status
func (c *StoreApiController) GetInventory(w http.ResponseWriter, r *http.Request) { 
    result, err := c.service.GetInventory()
    if err != nil {
        w.WriteHeader(500)
        return
    }
	
    EncodeJSONResponse(result, nil,  w)
}

// GetOrderById - Find purchase order by ID
func (c *StoreApiController) GetOrderById(w http.ResponseWriter, r *http.Request) { 
    params := mux.Vars(r)
	orderId, err := parseIntParameter(params["orderId"])
	if err != nil {
        w.WriteHeader(500)
        return
    }
	
    result, err := c.service.GetOrderById(orderId)
    if err != nil {
        w.WriteHeader(500)
        return
    }
	
    EncodeJSONResponse(result, nil,  w)
}

// PlaceOrder - Place an order for a pet
func (c *StoreApiController) PlaceOrder(w http.ResponseWriter, r *http.Request) { 
    body := &Order{}
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        w.WriteHeader(500)
        return
    }
	
    result, err := c.service.PlaceOrder(*body)
    if err != nil {
        w.WriteHeader(500)
        return
    }
	
    EncodeJSONResponse(result, nil,  w)
}
