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

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type DaDaPushMessageApiService service

/*
DaDaPushMessageApiService push Message to a Channel
&lt;h2&gt;Rate Limit:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;1 request per 1s&lt;/li&gt;&lt;li&gt;30 request per 1m&lt;/li&gt;&lt;li&gt;500 request per 1h&lt;/li&gt;&lt;/ul&gt;&lt;h2&gt;Result code/errmsg List:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;0: ok&lt;/li&gt;&lt;li&gt;1: server error&lt;/li&gt;&lt;li&gt;101: channel is exists&lt;/li&gt;&lt;li&gt;102: channel is not exists&lt;/li&gt;&lt;li&gt;103: channel token error&lt;/li&gt;&lt;li&gt;104: channel is not exists&lt;/li&gt;&lt;li&gt;105: message is not exists&lt;/li&gt;&lt;li&gt;204: bad request&lt;/li&gt;&lt;li&gt;205: permission deny&lt;/li&gt;&lt;li&gt;206: too many request, please after 5 minutes to try!&lt;/li&gt;&lt;li&gt;301: duplicate username/email&lt;/li&gt;&lt;li&gt;302: user is not exists&lt;/li&gt;&lt;li&gt;303: user password is error&lt;/li&gt;&lt;li&gt;304: client push token is error&lt;/li&gt;&lt;li&gt;305: user is disabled&lt;/li&gt;&lt;li&gt;306: your subscription is expired&lt;/li&gt;&lt;li&gt;307: user not subscribe channel&lt;/li&gt;&lt;/ul&gt;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body body
 * @param optional nil or *CreateMessageOpts - Optional Parameters:
 * @param "ChannelToken" -  see: https://www.dadapush.com/channel/list
@return ResultOfMessagePushResponse
*/

func (a *DaDaPushMessageApiService) CreateMessage(ctx context.Context, body MessagePushRequest, channelToken string) (ResultOfMessagePushResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ResultOfMessagePushResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/message"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["x-channel-token"] = parameterToString(channelToken, "")
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	_ = localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResultOfMessagePushResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DaDaPushMessageApiService delete a Channel Message
&lt;h2&gt;Rate Limit:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;10 request per 1s&lt;/li&gt;&lt;li&gt;100 request per 1m&lt;/li&gt;&lt;li&gt;1000 request per 1h&lt;/li&gt;&lt;/ul&gt;&lt;h2&gt;Result code/errmsg List:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;0: ok&lt;/li&gt;&lt;li&gt;1: server error&lt;/li&gt;&lt;li&gt;101: channel is exists&lt;/li&gt;&lt;li&gt;102: channel is not exists&lt;/li&gt;&lt;li&gt;103: channel token error&lt;/li&gt;&lt;li&gt;104: channel is not exists&lt;/li&gt;&lt;li&gt;105: message is not exists&lt;/li&gt;&lt;li&gt;204: bad request&lt;/li&gt;&lt;li&gt;205: permission deny&lt;/li&gt;&lt;li&gt;206: too many request, please after 5 minutes to try!&lt;/li&gt;&lt;li&gt;301: duplicate username/email&lt;/li&gt;&lt;li&gt;302: user is not exists&lt;/li&gt;&lt;li&gt;303: user password is error&lt;/li&gt;&lt;li&gt;304: client push token is error&lt;/li&gt;&lt;li&gt;305: user is disabled&lt;/li&gt;&lt;li&gt;306: your subscription is expired&lt;/li&gt;&lt;li&gt;307: user not subscribe channel&lt;/li&gt;&lt;/ul&gt;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param messageId messageId
 * @param optional nil or *DeleteMessageOpts - Optional Parameters:
 * @param "ChannelToken"  -  see: https://www.dadapush.com/channel/list
@return Result
*/

func (a *DaDaPushMessageApiService) DeleteMessage(ctx context.Context, messageId int64, channelToken string) (Result, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Result
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/message/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", fmt.Sprintf("%v", messageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	var localVarHttpContentTypes []string

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["x-channel-token"] = parameterToString(channelToken, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	_ = localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Result
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DaDaPushMessageApiService get a Channel Message
&lt;h2&gt;Rate Limit:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;10 request per 1s&lt;/li&gt;&lt;li&gt;100 request per 1m&lt;/li&gt;&lt;li&gt;1000 request per 1h&lt;/li&gt;&lt;/ul&gt;&lt;h2&gt;Result code/errmsg List:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;0: ok&lt;/li&gt;&lt;li&gt;1: server error&lt;/li&gt;&lt;li&gt;101: channel is exists&lt;/li&gt;&lt;li&gt;102: channel is not exists&lt;/li&gt;&lt;li&gt;103: channel token error&lt;/li&gt;&lt;li&gt;104: channel is not exists&lt;/li&gt;&lt;li&gt;105: message is not exists&lt;/li&gt;&lt;li&gt;204: bad request&lt;/li&gt;&lt;li&gt;205: permission deny&lt;/li&gt;&lt;li&gt;206: too many request, please after 5 minutes to try!&lt;/li&gt;&lt;li&gt;301: duplicate username/email&lt;/li&gt;&lt;li&gt;302: user is not exists&lt;/li&gt;&lt;li&gt;303: user password is error&lt;/li&gt;&lt;li&gt;304: client push token is error&lt;/li&gt;&lt;li&gt;305: user is disabled&lt;/li&gt;&lt;li&gt;306: your subscription is expired&lt;/li&gt;&lt;li&gt;307: user not subscribe channel&lt;/li&gt;&lt;/ul&gt;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param messageId messageId
 * @param optional nil or *GetMessageOpts - Optional Parameters:
 * @param "ChannelToken" -  see: https://www.dadapush.com/channel/list
@return ResultOfMessageObject
*/

func (a *DaDaPushMessageApiService) GetMessage(ctx context.Context, messageId int64, channelToken string) (ResultOfMessageObject, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ResultOfMessageObject
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/message/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", fmt.Sprintf("%v", messageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	var localVarHttpContentTypes []string

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["x-channel-token"] = parameterToString(channelToken, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	_ = localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResultOfMessageObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DaDaPushMessageApiService get Message List
&lt;h2&gt;Rate Limit:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;1 request per 1s&lt;/li&gt;&lt;li&gt;45 request per 1m&lt;/li&gt;&lt;/ul&gt;&lt;h2&gt;Result code/errmsg List:&lt;/h2&gt;&lt;ul&gt;&lt;li&gt;0: ok&lt;/li&gt;&lt;li&gt;1: server error&lt;/li&gt;&lt;li&gt;101: channel is exists&lt;/li&gt;&lt;li&gt;102: channel is not exists&lt;/li&gt;&lt;li&gt;103: channel token error&lt;/li&gt;&lt;li&gt;104: channel is not exists&lt;/li&gt;&lt;li&gt;105: message is not exists&lt;/li&gt;&lt;li&gt;204: bad request&lt;/li&gt;&lt;li&gt;205: permission deny&lt;/li&gt;&lt;li&gt;206: too many request, please after 5 minutes to try!&lt;/li&gt;&lt;li&gt;301: duplicate username/email&lt;/li&gt;&lt;li&gt;302: user is not exists&lt;/li&gt;&lt;li&gt;303: user password is error&lt;/li&gt;&lt;li&gt;304: client push token is error&lt;/li&gt;&lt;li&gt;305: user is disabled&lt;/li&gt;&lt;li&gt;306: your subscription is expired&lt;/li&gt;&lt;li&gt;307: user not subscribe channel&lt;/li&gt;&lt;/ul&gt;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param page greater than 1
 * @param pageSize range is 1,50
 * @param optional nil or *GetMessagesOpts - Optional Parameters:
 * @param "ChannelToken" -  see: https://www.dadapush.com/channel/list
@return ResultOfPageResponseOfMessageObject
*/

func (a *DaDaPushMessageApiService) GetMessages(ctx context.Context, page int32, pageSize int32, channelToken string) (ResultOfPageResponseOfMessageObject, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ResultOfPageResponseOfMessageObject
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("page", parameterToString(page, ""))
	localVarQueryParams.Add("pageSize", parameterToString(pageSize, ""))
	// to determine the Content-Type header
	var localVarHttpContentTypes []string

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	localVarHeaderParams["x-channel-token"] = parameterToString(channelToken, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	_ = localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ResultOfPageResponseOfMessageObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
