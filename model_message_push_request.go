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

type MessagePushRequest struct {
	// action size range is 0,3
	Actions []Action `json:"actions,omitempty"`
	// message content
	Content string `json:"content"`
	// when value is false, will not send client push
	NeedPush bool `json:"needPush"`
	// message title
	Title string `json:"title"`
}
