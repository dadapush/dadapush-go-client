/*
 * DaDaPush Public API
 *
 * DaDaPush: Real-time Notifications App Send real-time notifications through our API without coding and maintaining your own app for iOS or Android devices.
 *
 * API version: v1
 * Contact: contacts@dadapush.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dadapushclient

type PageResponseOfMessageObject struct {
	Content       []MessageObject `json:"content,omitempty"`
	TotalElements int64           `json:"totalElements,omitempty"`
	TotalPages    int64           `json:"totalPages,omitempty"`
}
