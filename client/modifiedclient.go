// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"public": dara.String("alidns.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = GetEndpoint(client, dara.String("alidns"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func GetEndpoint(client *Client, productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a custom line for a domain name.
//
// Description:
//
// The end IP address of an IP address segment must be greater than or equal to its start IP address.
//
// The IP address ranges of segments cannot overlap across any custom lines for the domain name.
//
// @param request - AddCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomLineResponse
func AddCustomLineWithOptions(client *Client, request *AddCustomLineRequest, runtime *dara.RuntimeOptions) (_result *AddCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.IpSegment) {
		query["IpSegment"] = request.IpSegment
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineName) {
		query["LineName"] = request.LineName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomLine"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom line for a domain name.
//
// Description:
//
// The end IP address of an IP address segment must be greater than or equal to its start IP address.
//
// The IP address ranges of segments cannot overlap across any custom lines for the domain name.
//
// @param request - AddCustomLineRequest
//
// @return AddCustomLineResponse
func AddCustomLine(client *Client, request *AddCustomLineRequest) (_result *AddCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCustomLineResponse{}
	_body, _err := AddCustomLineWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a domain name to the DNS authoritative proxy service.
//
// @param request - AddDnsCacheDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsCacheDomainResponse
func AddDnsCacheDomainWithOptions(client *Client, request *AddDnsCacheDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDnsCacheDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheTtlMax) {
		query["CacheTtlMax"] = request.CacheTtlMax
	}

	if !dara.IsNil(request.CacheTtlMin) {
		query["CacheTtlMin"] = request.CacheTtlMin
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SourceDnsServer) {
		query["SourceDnsServer"] = request.SourceDnsServer
	}

	if !dara.IsNil(request.SourceEdns) {
		query["SourceEdns"] = request.SourceEdns
	}

	if !dara.IsNil(request.SourceProtocol) {
		query["SourceProtocol"] = request.SourceProtocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsCacheDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsCacheDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to the DNS authoritative proxy service.
//
// @param request - AddDnsCacheDomainRequest
//
// @return AddDnsCacheDomainResponse
func AddDnsCacheDomain(client *Client, request *AddDnsCacheDomainRequest) (_result *AddDnsCacheDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsCacheDomainResponse{}
	_body, _err := AddDnsCacheDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an access strategy.
//
// @param request - AddDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsGtmAccessStrategyResponse
func AddDnsGtmAccessStrategyWithOptions(client *Client, request *AddDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *AddDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultAddrPool) {
		query["DefaultAddrPool"] = request.DefaultAddrPool
	}

	if !dara.IsNil(request.DefaultAddrPoolType) {
		query["DefaultAddrPoolType"] = request.DefaultAddrPoolType
	}

	if !dara.IsNil(request.DefaultLatencyOptimization) {
		query["DefaultLatencyOptimization"] = request.DefaultLatencyOptimization
	}

	if !dara.IsNil(request.DefaultLbaStrategy) {
		query["DefaultLbaStrategy"] = request.DefaultLbaStrategy
	}

	if !dara.IsNil(request.DefaultMaxReturnAddrNum) {
		query["DefaultMaxReturnAddrNum"] = request.DefaultMaxReturnAddrNum
	}

	if !dara.IsNil(request.DefaultMinAvailableAddrNum) {
		query["DefaultMinAvailableAddrNum"] = request.DefaultMinAvailableAddrNum
	}

	if !dara.IsNil(request.FailoverAddrPool) {
		query["FailoverAddrPool"] = request.FailoverAddrPool
	}

	if !dara.IsNil(request.FailoverAddrPoolType) {
		query["FailoverAddrPoolType"] = request.FailoverAddrPoolType
	}

	if !dara.IsNil(request.FailoverLatencyOptimization) {
		query["FailoverLatencyOptimization"] = request.FailoverLatencyOptimization
	}

	if !dara.IsNil(request.FailoverLbaStrategy) {
		query["FailoverLbaStrategy"] = request.FailoverLbaStrategy
	}

	if !dara.IsNil(request.FailoverMaxReturnAddrNum) {
		query["FailoverMaxReturnAddrNum"] = request.FailoverMaxReturnAddrNum
	}

	if !dara.IsNil(request.FailoverMinAvailableAddrNum) {
		query["FailoverMinAvailableAddrNum"] = request.FailoverMinAvailableAddrNum
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lines) {
		query["Lines"] = request.Lines
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an access strategy.
//
// @param request - AddDnsGtmAccessStrategyRequest
//
// @return AddDnsGtmAccessStrategyResponse
func AddDnsGtmAccessStrategy(client *Client, request *AddDnsGtmAccessStrategyRequest) (_result *AddDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsGtmAccessStrategyResponse{}
	_body, _err := AddDnsGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddDnsGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsGtmAddressPoolResponse
func AddDnsGtmAddressPoolWithOptions(client *Client, request *AddDnsGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *AddDnsGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LbaStrategy) {
		query["LbaStrategy"] = request.LbaStrategy
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.MonitorStatus) {
		query["MonitorStatus"] = request.MonitorStatus
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddDnsGtmAddressPoolRequest
//
// @return AddDnsGtmAddressPoolResponse
func AddDnsGtmAddressPool(client *Client, request *AddDnsGtmAddressPoolRequest) (_result *AddDnsGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsGtmAddressPoolResponse{}
	_body, _err := AddDnsGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a health check.
//
// Description:
//
// **
//
// @param request - AddDnsGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsGtmMonitorResponse
func AddDnsGtmMonitorWithOptions(client *Client, request *AddDnsGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *AddDnsGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a health check.
//
// Description:
//
// **
//
// @param request - AddDnsGtmMonitorRequest
//
// @return AddDnsGtmMonitorResponse
func AddDnsGtmMonitor(client *Client, request *AddDnsGtmMonitorRequest) (_result *AddDnsGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsGtmMonitorResponse{}
	_body, _err := AddDnsGtmMonitorWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a domain name.
//
// Description:
//
// For more information, see <props="china">[Domain name validity](https://help.aliyun.com/document_detail/67788.html)<props="intl">[Domain name validity](https://www.alibabacloud.com/help/zh/doc-detail/67788.htm).
//
// @param request - AddDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainResponse
func AddDomainWithOptions(client *Client, request *AddDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name.
//
// Description:
//
// For more information, see <props="china">[Domain name validity](https://help.aliyun.com/document_detail/67788.html)<props="intl">[Domain name validity](https://www.alibabacloud.com/help/zh/doc-detail/67788.htm).
//
// @param request - AddDomainRequest
//
// @return AddDomainResponse
func AddDomain(client *Client, request *AddDomainRequest) (_result *AddDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainResponse{}
	_body, _err := AddDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backup for a domain based on the specified domain name and backup cycle.
//
// @param request - AddDomainBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainBackupResponse
func AddDomainBackupWithOptions(client *Client, request *AddDomainBackupRequest, runtime *dara.RuntimeOptions) (_result *AddDomainBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomainBackup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup for a domain based on the specified domain name and backup cycle.
//
// @param request - AddDomainBackupRequest
//
// @return AddDomainBackupResponse
func AddDomainBackup(client *Client, request *AddDomainBackupRequest) (_result *AddDomainBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainBackupResponse{}
	_body, _err := AddDomainBackupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a domain name group.
//
// @param request - AddDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainGroupResponse
func AddDomainGroupWithOptions(client *Client, request *AddDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *AddDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a domain name group.
//
// @param request - AddDomainGroupRequest
//
// @return AddDomainGroupResponse
func AddDomainGroup(client *Client, request *AddDomainGroupRequest) (_result *AddDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainGroupResponse{}
	_body, _err := AddDomainGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a DNS record.
//
// @param request - AddDomainRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainRecordResponse
func AddDomainRecordWithOptions(client *Client, request *AddDomainRecordRequest, runtime *dara.RuntimeOptions) (_result *AddDomainRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RR) {
		query["RR"] = request.RR
	}

	if !dara.IsNil(request.TTL) {
		query["TTL"] = request.TTL
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomainRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a DNS record.
//
// @param request - AddDomainRecordRequest
//
// @return AddDomainRecordResponse
func AddDomainRecord(client *Client, request *AddDomainRecordRequest) (_result *AddDomainRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainRecordResponse{}
	_body, _err := AddDomainRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access strategy based on the specified parameters.
//
// @param request - AddGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmAccessStrategyResponse
func AddGtmAccessStrategyWithOptions(client *Client, request *AddGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *AddGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessLines) {
		query["AccessLines"] = request.AccessLines
	}

	if !dara.IsNil(request.DefaultAddrPoolId) {
		query["DefaultAddrPoolId"] = request.DefaultAddrPoolId
	}

	if !dara.IsNil(request.FailoverAddrPoolId) {
		query["FailoverAddrPoolId"] = request.FailoverAddrPoolId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access strategy based on the specified parameters.
//
// @param request - AddGtmAccessStrategyRequest
//
// @return AddGtmAccessStrategyResponse
func AddGtmAccessStrategy(client *Client, request *AddGtmAccessStrategyRequest) (_result *AddGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmAccessStrategyResponse{}
	_body, _err := AddGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmAddressPoolResponse
func AddGtmAddressPoolWithOptions(client *Client, request *AddGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *AddGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MinAvailableAddrNum) {
		query["MinAvailableAddrNum"] = request.MinAvailableAddrNum
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.MonitorStatus) {
		query["MonitorStatus"] = request.MonitorStatus
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddGtmAddressPoolRequest
//
// @return AddGtmAddressPoolResponse
func AddGtmAddressPool(client *Client, request *AddGtmAddressPoolRequest) (_result *AddGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmAddressPoolResponse{}
	_body, _err := AddGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a health check.
//
// @param request - AddGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmMonitorResponse
func AddGtmMonitorWithOptions(client *Client, request *AddGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *AddGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a health check.
//
// @param request - AddGtmMonitorRequest
//
// @return AddGtmMonitorResponse
func AddGtmMonitor(client *Client, request *AddGtmMonitorRequest) (_result *AddGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmMonitorResponse{}
	_body, _err := AddGtmMonitorWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a disaster recovery plan.
//
// @param request - AddGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmRecoveryPlanResponse
func AddGtmRecoveryPlanWithOptions(client *Client, request *AddGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *AddGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaultAddrPool) {
		query["FaultAddrPool"] = request.FaultAddrPool
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a disaster recovery plan.
//
// @param request - AddGtmRecoveryPlanRequest
//
// @return AddGtmRecoveryPlanResponse
func AddGtmRecoveryPlan(client *Client, request *AddGtmRecoveryPlanRequest) (_result *AddGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmRecoveryPlanResponse{}
	_body, _err := AddGtmRecoveryPlanWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an authoritative record for recursive resolution.
//
// Description:
//
// - You can specify a domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve the list of DNS records for that domain name.
//
// - To find DNS records that contain a specific keyword, you can specify the keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord).
//
// - By default, the list of DNS records is sorted from newest to oldest.
//
// - You can specify a domain group ID (GroupId). The \\`All Domains\\` group includes all domain names. The \\`Default Group\\` includes domain names that are not assigned to a group.
//
// @param request - AddRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecursionRecordResponse
func AddRecursionRecordWithOptions(client *Client, request *AddRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *AddRecursionRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an authoritative record for recursive resolution.
//
// Description:
//
// - You can specify a domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve the list of DNS records for that domain name.
//
// - To find DNS records that contain a specific keyword, you can specify the keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord).
//
// - By default, the list of DNS records is sorted from newest to oldest.
//
// - You can specify a domain group ID (GroupId). The \\`All Domains\\` group includes all domain names. The \\`Default Group\\` includes domain names that are not assigned to a group.
//
// @param request - AddRecursionRecordRequest
//
// @return AddRecursionRecordResponse
func AddRecursionRecord(client *Client, request *AddRecursionRecordRequest) (_result *AddRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRecursionRecordResponse{}
	_body, _err := AddRecursionRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a built-in authoritative domain name zone for recursive resolution.
//
// Description:
//
// The end IP address of each IP range must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP ranges in all custom lines for a domain name cannot overlap.
//
// @param request - AddRecursionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecursionZoneResponse
func AddRecursionZoneWithOptions(client *Client, request *AddRecursionZoneRequest, runtime *dara.RuntimeOptions) (_result *AddRecursionZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProxyPattern) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecursionZone"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecursionZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a built-in authoritative domain name zone for recursive resolution.
//
// Description:
//
// The end IP address of each IP range must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP ranges in all custom lines for a domain name cannot overlap.
//
// @param request - AddRecursionZoneRequest
//
// @return AddRecursionZoneResponse
func AddRecursionZone(client *Client, request *AddRecursionZoneRequest) (_result *AddRecursionZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRecursionZoneResponse{}
	_body, _err := AddRecursionZoneWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the serverHold status to a specified domain name.
//
// Description:
//
// ## Request description
//
// - This API adds the serverHold property to a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRspDomainServerHoldStatusForGatewayResponse
func AddRspDomainServerHoldStatusForGatewayWithOptions(client *Client, request *AddRspDomainServerHoldStatusForGatewayRequest, runtime *dara.RuntimeOptions) (_result *AddRspDomainServerHoldStatusForGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRspDomainServerHoldStatusForGateway"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the serverHold status to a specified domain name.
//
// Description:
//
// ## Request description
//
// - This API adds the serverHold property to a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayRequest
//
// @return AddRspDomainServerHoldStatusForGatewayResponse
func AddRspDomainServerHoldStatusForGateway(client *Client, request *AddRspDomainServerHoldStatusForGatewayRequest) (_result *AddRspDomainServerHoldStatusForGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := AddRspDomainServerHoldStatusForGatewayWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the serverHold status for a specified domain name.
//
// Description:
//
// ## Request description
//
// - Adds the serverHold status for a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayOteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRspDomainServerHoldStatusForGatewayOteResponse
func AddRspDomainServerHoldStatusForGatewayOteWithOptions(client *Client, request *AddRspDomainServerHoldStatusForGatewayOteRequest, runtime *dara.RuntimeOptions) (_result *AddRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRspDomainServerHoldStatusForGatewayOte"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the serverHold status for a specified domain name.
//
// Description:
//
// ## Request description
//
// - Adds the serverHold status for a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayOteRequest
//
// @return AddRspDomainServerHoldStatusForGatewayOteResponse
func AddRspDomainServerHoldStatusForGatewayOte(client *Client, request *AddRspDomainServerHoldStatusForGatewayOteRequest) (_result *AddRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := AddRspDomainServerHoldStatusForGatewayOteWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds paid domain names in Alibaba Cloud DNS to an instance ID.
//
// Description:
//
// An instance is considered a new instance if its ID starts with \\"dns-\\". New instances support multiple domain names. You can call this operation to bind domain names directly to the instance. An error occurs if the number of domain names exceeds the instance\\"s limit.
//
// An instance is considered a legacy instance if its ID does not start with \\"dns-\\". Legacy instances support only one domain name. If you call this operation on a legacy instance that already has a domain name, the existing domain name is replaced.
//
// @param request - BindInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindInstanceDomainsResponse
func BindInstanceDomainsWithOptions(client *Client, request *BindInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *BindInstanceDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindInstanceDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindInstanceDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds paid domain names in Alibaba Cloud DNS to an instance ID.
//
// Description:
//
// An instance is considered a new instance if its ID starts with \\"dns-\\". New instances support multiple domain names. You can call this operation to bind domain names directly to the instance. An error occurs if the number of domain names exceeds the instance\\"s limit.
//
// An instance is considered a legacy instance if its ID does not start with \\"dns-\\". Legacy instances support only one domain name. If you call this operation on a legacy instance that already has a domain name, the existing domain name is replaced.
//
// @param request - BindInstanceDomainsRequest
//
// @return BindInstanceDomainsResponse
func BindInstanceDomains(client *Client, request *BindInstanceDomainsRequest) (_result *BindInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindInstanceDomainsResponse{}
	_body, _err := BindInstanceDomainsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a domain name to a new group.
//
// Description:
//
// You can specify the ID of a domain name group (GroupId). The All Domains group contains all domain names, while the Default group contains domain names that are not assigned to any group.
//
// @param request - ChangeDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDomainGroupResponse
func ChangeDomainGroupWithOptions(client *Client, request *ChangeDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a domain name to a new group.
//
// Description:
//
// You can specify the ID of a domain name group (GroupId). The All Domains group contains all domain names, while the Default group contains domain names that are not assigned to any group.
//
// @param request - ChangeDomainGroupRequest
//
// @return ChangeDomainGroupResponse
func ChangeDomainGroup(client *Client, request *ChangeDomainGroupRequest) (_result *ChangeDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeDomainGroupResponse{}
	_body, _err := ChangeDomainGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the domain name that is attached to a Cloud DNS product.
//
// Description:
//
// - **You can call this operation to change the domain name that is attached to a Cloud DNS product. To detach a domain name, call this operation and leave the NewDomain parameter empty.**
//
// - **This operation applies to instances of earlier versions. If you use a new edition, such as Personal Edition, Enterprise Standard Edition, or Enterprise Ultimate Edition, call the BindInstanceDomains operation instead.**
//
// @param request - ChangeDomainOfDnsProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDomainOfDnsProductResponse
func ChangeDomainOfDnsProductWithOptions(client *Client, request *ChangeDomainOfDnsProductRequest, runtime *dara.RuntimeOptions) (_result *ChangeDomainOfDnsProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewDomain) {
		query["NewDomain"] = request.NewDomain
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDomainOfDnsProduct"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDomainOfDnsProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the domain name that is attached to a Cloud DNS product.
//
// Description:
//
// - **You can call this operation to change the domain name that is attached to a Cloud DNS product. To detach a domain name, call this operation and leave the NewDomain parameter empty.**
//
// - **This operation applies to instances of earlier versions. If you use a new edition, such as Personal Edition, Enterprise Standard Edition, or Enterprise Ultimate Edition, call the BindInstanceDomains operation instead.**
//
// @param request - ChangeDomainOfDnsProductRequest
//
// @return ChangeDomainOfDnsProductResponse
func ChangeDomainOfDnsProduct(client *Client, request *ChangeDomainOfDnsProductRequest) (_result *ChangeDomainOfDnsProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeDomainOfDnsProductResponse{}
	_body, _err := ChangeDomainOfDnsProductWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Copies a Global Traffic Manager (GTM) configuration.
//
// @param request - CopyGtmConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyGtmConfigResponse
func CopyGtmConfigWithOptions(client *Client, request *CopyGtmConfigRequest, runtime *dara.RuntimeOptions) (_result *CopyGtmConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CopyType) {
		query["CopyType"] = request.CopyType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyGtmConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyGtmConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a Global Traffic Manager (GTM) configuration.
//
// @param request - CopyGtmConfigRequest
//
// @return CopyGtmConfigResponse
func CopyGtmConfig(client *Client, request *CopyGtmConfigRequest) (_result *CopyGtmConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyGtmConfigResponse{}
	_body, _err := CopyGtmConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers an agent.
//
// @param tmpReq - CreateAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAtiAgentRegisterInfoResponse
func CreateAtiAgentRegisterInfoWithOptions(client *Client, tmpReq *CreateAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAtiAgentRegisterInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Endpoints) {
		request.EndpointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoints, dara.String("Endpoints"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentDescription) {
		query["AgentDescription"] = request.AgentDescription
	}

	if !dara.IsNil(request.AgentDisplayName) {
		query["AgentDisplayName"] = request.AgentDisplayName
	}

	if !dara.IsNil(request.AgentHost) {
		query["AgentHost"] = request.AgentHost
	}

	if !dara.IsNil(request.AgentVersion) {
		query["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointsShrink) {
		query["Endpoints"] = request.EndpointsShrink
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers an agent.
//
// @param request - CreateAtiAgentRegisterInfoRequest
//
// @return CreateAtiAgentRegisterInfoResponse
func CreateAtiAgentRegisterInfo(client *Client, request *CreateAtiAgentRegisterInfoRequest) (_result *CreateAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAtiAgentRegisterInfoResponse{}
	_body, _err := CreateAtiAgentRegisterInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers an agent - Step 2: Generates a DNS record for domain ownership verification.
//
// @param request - CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse
func CreateAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(client *Client, request *CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAtiAgentRegisterInfoAcmeChallengeRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers an agent - Step 2: Generates a DNS record for domain ownership verification.
//
// @param request - CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @return CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse
func CreateAtiAgentRegisterInfoAcmeChallengeRecord(client *Client, request *CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest) (_result *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := CreateAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a real-name registrant.
//
// @param request - CreateAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAtiRegistrantResponse
func CreateAtiRegistrantWithOptions(client *Client, request *CreateAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *CreateAtiRegistrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cc) {
		query["Cc"] = request.Cc
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DocumentCode) {
		query["DocumentCode"] = request.DocumentCode
	}

	if !dara.IsNil(request.DocumentImage) {
		query["DocumentImage"] = request.DocumentImage
	}

	if !dara.IsNil(request.DocumentType) {
		query["DocumentType"] = request.DocumentType
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Street) {
		query["Street"] = request.Street
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a real-name registrant.
//
// @param request - CreateAtiRegistrantRequest
//
// @return CreateAtiRegistrantResponse
func CreateAtiRegistrant(client *Client, request *CreateAtiRegistrantRequest) (_result *CreateAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAtiRegistrantResponse{}
	_body, _err := CreateAtiRegistrantWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an address in Global Traffic Manager (GTM) 3.0.
//
// @param tmpReq - CreateCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmAddressResponse
func CreateCloudGtmAddressWithOptions(client *Client, tmpReq *CreateCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCloudGtmAddressShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HealthTasks) {
		request.HealthTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HealthTasks, dara.String("HealthTasks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AttributeInfo) {
		query["AttributeInfo"] = request.AttributeInfo
	}

	if !dara.IsNil(request.AvailableMode) {
		query["AvailableMode"] = request.AvailableMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	if !dara.IsNil(request.HealthTasksShrink) {
		query["HealthTasks"] = request.HealthTasksShrink
	}

	if !dara.IsNil(request.ManualAvailableStatus) {
		query["ManualAvailableStatus"] = request.ManualAvailableStatus
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an address in Global Traffic Manager (GTM) 3.0.
//
// @param request - CreateCloudGtmAddressRequest
//
// @return CreateCloudGtmAddressResponse
func CreateCloudGtmAddress(client *Client, request *CreateCloudGtmAddressRequest) (_result *CreateCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmAddressResponse{}
	_body, _err := CreateCloudGtmAddressWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an address pool based on the specified parameters.
//
// @param request - CreateCloudGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmAddressPoolResponse
func CreateCloudGtmAddressPoolWithOptions(client *Client, request *CreateCloudGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.AddressPoolType) {
		query["AddressPoolType"] = request.AddressPoolType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an address pool based on the specified parameters.
//
// @param request - CreateCloudGtmAddressPoolRequest
//
// @return CreateCloudGtmAddressPoolResponse
func CreateCloudGtmAddressPool(client *Client, request *CreateCloudGtmAddressPoolRequest) (_result *CreateCloudGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmAddressPoolResponse{}
	_body, _err := CreateCloudGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a configuration for a Global Traffic Manager (GTM) instance.
//
// @param request - CreateCloudGtmInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmInstanceConfigResponse
func CreateCloudGtmInstanceConfigWithOptions(client *Client, request *CreateCloudGtmInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ScheduleHostname) {
		query["ScheduleHostname"] = request.ScheduleHostname
	}

	if !dara.IsNil(request.ScheduleRrType) {
		query["ScheduleRrType"] = request.ScheduleRrType
	}

	if !dara.IsNil(request.ScheduleZoneMode) {
		query["ScheduleZoneMode"] = request.ScheduleZoneMode
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmInstanceConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration for a Global Traffic Manager (GTM) instance.
//
// @param request - CreateCloudGtmInstanceConfigRequest
//
// @return CreateCloudGtmInstanceConfigResponse
func CreateCloudGtmInstanceConfig(client *Client, request *CreateCloudGtmInstanceConfigRequest) (_result *CreateCloudGtmInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmInstanceConfigResponse{}
	_body, _err := CreateCloudGtmInstanceConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a health check template in Global Traffic Manager (GTM) 3.0.
//
// @param tmpReq - CreateCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmMonitorTemplateResponse
func CreateCloudGtmMonitorTemplateWithOptions(client *Client, tmpReq *CreateCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCloudGtmMonitorTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IspCityNodes) {
		request.IspCityNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IspCityNodes, dara.String("IspCityNodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.ExtendInfo) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.FailureRate) {
		query["FailureRate"] = request.FailureRate
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.IspCityNodesShrink) {
		query["IspCityNodes"] = request.IspCityNodesShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a health check template in Global Traffic Manager (GTM) 3.0.
//
// @param request - CreateCloudGtmMonitorTemplateRequest
//
// @return CreateCloudGtmMonitorTemplateResponse
func CreateCloudGtmMonitorTemplate(client *Client, request *CreateCloudGtmMonitorTemplateRequest) (_result *CreateCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmMonitorTemplateResponse{}
	_body, _err := CreateCloudGtmMonitorTemplateWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public DNS AccessKey.
//
// @param request - CreatePdnsAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdnsAppKeyResponse
func CreatePdnsAppKeyWithOptions(client *Client, request *CreatePdnsAppKeyRequest, runtime *dara.RuntimeOptions) (_result *CreatePdnsAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdnsAppKey"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdnsAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public DNS AccessKey.
//
// @param request - CreatePdnsAppKeyRequest
//
// @return CreatePdnsAppKeyResponse
func CreatePdnsAppKey(client *Client, request *CreatePdnsAppKeyRequest) (_result *CreatePdnsAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePdnsAppKeyResponse{}
	_body, _err := CreatePdnsAppKeyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Public DNS UDP IP Address Segment
//
// @param request - CreatePdnsUdpIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdnsUdpIpSegmentResponse
func CreatePdnsUdpIpSegmentWithOptions(client *Client, request *CreatePdnsUdpIpSegmentRequest, runtime *dara.RuntimeOptions) (_result *CreatePdnsUdpIpSegmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.IpToken) {
		query["IpToken"] = request.IpToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdnsUdpIpSegment"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdnsUdpIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Public DNS UDP IP Address Segment
//
// @param request - CreatePdnsUdpIpSegmentRequest
//
// @return CreatePdnsUdpIpSegmentResponse
func CreatePdnsUdpIpSegment(client *Client, request *CreatePdnsUdpIpSegmentRequest) (_result *CreatePdnsUdpIpSegmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePdnsUdpIpSegmentResponse{}
	_body, _err := CreatePdnsUdpIpSegmentWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes Agent registration information.
//
// @param request - DeleteAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAtiAgentRegisterInfoResponse
func DeleteAtiAgentRegisterInfoWithOptions(client *Client, request *DeleteAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *DeleteAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes Agent registration information.
//
// @param request - DeleteAtiAgentRegisterInfoRequest
//
// @return DeleteAtiAgentRegisterInfoResponse
func DeleteAtiAgentRegisterInfo(client *Client, request *DeleteAtiAgentRegisterInfoRequest) (_result *DeleteAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAtiAgentRegisterInfoResponse{}
	_body, _err := DeleteAtiAgentRegisterInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a verified registrant.
//
// @param request - DeleteAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAtiRegistrantResponse
func DeleteAtiRegistrantWithOptions(client *Client, request *DeleteAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *DeleteAtiRegistrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a verified registrant.
//
// @param request - DeleteAtiRegistrantRequest
//
// @return DeleteAtiRegistrantResponse
func DeleteAtiRegistrant(client *Client, request *DeleteAtiRegistrantRequest) (_result *DeleteAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAtiRegistrantResponse{}
	_body, _err := DeleteAtiRegistrantWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address.
//
// @param request - DeleteCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmAddressResponse
func DeleteCloudGtmAddressWithOptions(client *Client, request *DeleteCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address.
//
// @param request - DeleteCloudGtmAddressRequest
//
// @return DeleteCloudGtmAddressResponse
func DeleteCloudGtmAddress(client *Client, request *DeleteCloudGtmAddressRequest) (_result *DeleteCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmAddressResponse{}
	_body, _err := DeleteCloudGtmAddressWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteCloudGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmAddressPoolResponse
func DeleteCloudGtmAddressPoolWithOptions(client *Client, request *DeleteCloudGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteCloudGtmAddressPoolRequest
//
// @return DeleteCloudGtmAddressPoolResponse
func DeleteCloudGtmAddressPool(client *Client, request *DeleteCloudGtmAddressPoolRequest) (_result *DeleteCloudGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmAddressPoolResponse{}
	_body, _err := DeleteCloudGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access domain name from the configuration of a Global Traffic Manager (GTM) 3.0 instance.
//
// @param request - DeleteCloudGtmInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmInstanceConfigResponse
func DeleteCloudGtmInstanceConfigWithOptions(client *Client, request *DeleteCloudGtmInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmInstanceConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access domain name from the configuration of a Global Traffic Manager (GTM) 3.0 instance.
//
// @param request - DeleteCloudGtmInstanceConfigRequest
//
// @return DeleteCloudGtmInstanceConfigResponse
func DeleteCloudGtmInstanceConfig(client *Client, request *DeleteCloudGtmInstanceConfigRequest) (_result *DeleteCloudGtmInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmInstanceConfigResponse{}
	_body, _err := DeleteCloudGtmInstanceConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a health check template.
//
// @param request - DeleteCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmMonitorTemplateResponse
func DeleteCloudGtmMonitorTemplateWithOptions(client *Client, request *DeleteCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a health check template.
//
// @param request - DeleteCloudGtmMonitorTemplateRequest
//
// @return DeleteCloudGtmMonitorTemplateResponse
func DeleteCloudGtmMonitorTemplate(client *Client, request *DeleteCloudGtmMonitorTemplateRequest) (_result *DeleteCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmMonitorTemplateResponse{}
	_body, _err := DeleteCloudGtmMonitorTemplateWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a batch of custom lines by specifying their unique IDs.
//
// @param request - DeleteCustomLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomLinesResponse
func DeleteCustomLinesWithOptions(client *Client, request *DeleteCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomLinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineIds) {
		query["LineIds"] = request.LineIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a batch of custom lines by specifying their unique IDs.
//
// @param request - DeleteCustomLinesRequest
//
// @return DeleteCustomLinesResponse
func DeleteCustomLines(client *Client, request *DeleteCustomLinesRequest) (_result *DeleteCustomLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomLinesResponse{}
	_body, _err := DeleteCustomLinesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified domain name from the authoritative DNS proxy.
//
// @param request - DeleteDnsCacheDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsCacheDomainResponse
func DeleteDnsCacheDomainWithOptions(client *Client, request *DeleteDnsCacheDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsCacheDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnsCacheDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnsCacheDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified domain name from the authoritative DNS proxy.
//
// @param request - DeleteDnsCacheDomainRequest
//
// @return DeleteDnsCacheDomainResponse
func DeleteDnsCacheDomain(client *Client, request *DeleteDnsCacheDomainRequest) (_result *DeleteDnsCacheDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDnsCacheDomainResponse{}
	_body, _err := DeleteDnsCacheDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access policy by its ID.
//
// @param request - DeleteDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsGtmAccessStrategyResponse
func DeleteDnsGtmAccessStrategyWithOptions(client *Client, request *DeleteDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access policy by its ID.
//
// @param request - DeleteDnsGtmAccessStrategyRequest
//
// @return DeleteDnsGtmAccessStrategyResponse
func DeleteDnsGtmAccessStrategy(client *Client, request *DeleteDnsGtmAccessStrategyRequest) (_result *DeleteDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDnsGtmAccessStrategyResponse{}
	_body, _err := DeleteDnsGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address pool by its ID.
//
// @param request - DeleteDnsGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsGtmAddressPoolResponse
func DeleteDnsGtmAddressPoolWithOptions(client *Client, request *DeleteDnsGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnsGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnsGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address pool by its ID.
//
// @param request - DeleteDnsGtmAddressPoolRequest
//
// @return DeleteDnsGtmAddressPoolResponse
func DeleteDnsGtmAddressPool(client *Client, request *DeleteDnsGtmAddressPoolRequest) (_result *DeleteDnsGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDnsGtmAddressPoolResponse{}
	_body, _err := DeleteDnsGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified domain name.
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func DeleteDomainWithOptions(client *Client, request *DeleteDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified domain name.
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func DeleteDomain(client *Client, request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := DeleteDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a domain name group and moves its domain names to the default group.
//
// Description:
//
// > The default group cannot be deleted.
//
// @param request - DeleteDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainGroupResponse
func DeleteDomainGroupWithOptions(client *Client, request *DeleteDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a domain name group and moves its domain names to the default group.
//
// Description:
//
// > The default group cannot be deleted.
//
// @param request - DeleteDomainGroupRequest
//
// @return DeleteDomainGroupResponse
func DeleteDomainGroup(client *Client, request *DeleteDomainGroupRequest) (_result *DeleteDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := DeleteDomainGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a DNS record based on the specified request parameters.
//
// @param request - DeleteDomainRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainRecordResponse
func DeleteDomainRecordWithOptions(client *Client, request *DeleteDomainRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DNS record based on the specified request parameters.
//
// @param request - DeleteDomainRecordRequest
//
// @return DeleteDomainRecordResponse
func DeleteDomainRecord(client *Client, request *DeleteDomainRecordRequest) (_result *DeleteDomainRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainRecordResponse{}
	_body, _err := DeleteDomainRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access policy.
//
// @param request - DeleteGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGtmAccessStrategyResponse
func DeleteGtmAccessStrategyWithOptions(client *Client, request *DeleteGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access policy.
//
// @param request - DeleteGtmAccessStrategyRequest
//
// @return DeleteGtmAccessStrategyResponse
func DeleteGtmAccessStrategy(client *Client, request *DeleteGtmAccessStrategyRequest) (_result *DeleteGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGtmAccessStrategyResponse{}
	_body, _err := DeleteGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGtmAddressPoolResponse
func DeleteGtmAddressPoolWithOptions(client *Client, request *DeleteGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteGtmAddressPoolRequest
//
// @return DeleteGtmAddressPoolResponse
func DeleteGtmAddressPool(client *Client, request *DeleteGtmAddressPoolRequest) (_result *DeleteGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGtmAddressPoolResponse{}
	_body, _err := DeleteGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a disaster recovery plan.
//
// @param request - DeleteGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGtmRecoveryPlanResponse
func DeleteGtmRecoveryPlanWithOptions(client *Client, request *DeleteGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a disaster recovery plan.
//
// @param request - DeleteGtmRecoveryPlanRequest
//
// @return DeleteGtmRecoveryPlanResponse
func DeleteGtmRecoveryPlan(client *Client, request *DeleteGtmRecoveryPlanRequest) (_result *DeleteGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGtmRecoveryPlanResponse{}
	_body, _err := DeleteGtmRecoveryPlanWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DeleteRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecursionRecordResponse
func DeleteRecursionRecordWithOptions(client *Client, request *DeleteRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecursionRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DeleteRecursionRecordRequest
//
// @return DeleteRecursionRecordResponse
func DeleteRecursionRecord(client *Client, request *DeleteRecursionRecordRequest) (_result *DeleteRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecursionRecordResponse{}
	_body, _err := DeleteRecursionRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative zone used for recursive resolution.
//
// Description:
//
// If a zone contains locked DNS records, this operation does not delete them.
//
// @param request - DeleteRecursionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecursionZoneResponse
func DeleteRecursionZoneWithOptions(client *Client, request *DeleteRecursionZoneRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecursionZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecursionZone"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecursionZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative zone used for recursive resolution.
//
// Description:
//
// If a zone contains locked DNS records, this operation does not delete them.
//
// @param request - DeleteRecursionZoneRequest
//
// @return DeleteRecursionZoneResponse
func DeleteRecursionZone(client *Client, request *DeleteRecursionZoneRequest) (_result *DeleteRecursionZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecursionZoneResponse{}
	_body, _err := DeleteRecursionZoneWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the DNS records for a specified host record.
//
// Description:
//
// Locked DNS records will not be deleted.
//
// @param request - DeleteSubDomainRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubDomainRecordsResponse
func DeleteSubDomainRecordsWithOptions(client *Client, request *DeleteSubDomainRecordsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubDomainRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RR) {
		query["RR"] = request.RR
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubDomainRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubDomainRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the DNS records for a specified host record.
//
// Description:
//
// Locked DNS records will not be deleted.
//
// @param request - DeleteSubDomainRecordsRequest
//
// @return DeleteSubDomainRecordsResponse
func DeleteSubDomainRecords(client *Client, request *DeleteSubDomainRecordsRequest) (_result *DeleteSubDomainRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSubDomainRecordsResponse{}
	_body, _err := DeleteSubDomainRecordsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an Agent registration.
//
// @param request - DescribeAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiAgentRegisterInfoResponse
func DescribeAtiAgentRegisterInfoWithOptions(client *Client, request *DescribeAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Agent registration.
//
// @param request - DescribeAtiAgentRegisterInfoRequest
//
// @return DescribeAtiAgentRegisterInfoResponse
func DescribeAtiAgentRegisterInfo(client *Client, request *DescribeAtiAgentRegisterInfoRequest) (_result *DescribeAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiAgentRegisterInfoResponse{}
	_body, _err := DescribeAtiAgentRegisterInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为sdk查询agent信息功能
//
// @param request - DescribeAtiAgentRegisterInfoMarketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiAgentRegisterInfoMarketResponse
func DescribeAtiAgentRegisterInfoMarketWithOptions(client *Client, request *DescribeAtiAgentRegisterInfoMarketRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiAgentRegisterInfoMarketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentHost) {
		query["AgentHost"] = request.AgentHost
	}

	if !dara.IsNil(request.AgentVersion) {
		query["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiAgentRegisterInfoMarket"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiAgentRegisterInfoMarketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为sdk查询agent信息功能
//
// @param request - DescribeAtiAgentRegisterInfoMarketRequest
//
// @return DescribeAtiAgentRegisterInfoMarketResponse
func DescribeAtiAgentRegisterInfoMarket(client *Client, request *DescribeAtiAgentRegisterInfoMarketRequest) (_result *DescribeAtiAgentRegisterInfoMarketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiAgentRegisterInfoMarketResponse{}
	_body, _err := DescribeAtiAgentRegisterInfoMarketWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alert settings.
//
// @param request - DescribeAtiAlertSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiAlertSettingsResponse
func DescribeAtiAlertSettingsWithOptions(client *Client, request *DescribeAtiAlertSettingsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiAlertSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiAlertSettings"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiAlertSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert settings.
//
// @param request - DescribeAtiAlertSettingsRequest
//
// @return DescribeAtiAlertSettingsResponse
func DescribeAtiAlertSettings(client *Client, request *DescribeAtiAlertSettingsRequest) (_result *DescribeAtiAlertSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiAlertSettingsResponse{}
	_body, _err := DescribeAtiAlertSettingsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a certificate.
//
// @param request - DescribeAtiCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiCertificateResponse
func DescribeAtiCertificateWithOptions(client *Client, request *DescribeAtiCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentCertificateId) {
		query["AgentCertificateId"] = request.AgentCertificateId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiCertificate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a certificate.
//
// @param request - DescribeAtiCertificateRequest
//
// @return DescribeAtiCertificateResponse
func DescribeAtiCertificate(client *Client, request *DescribeAtiCertificateRequest) (_result *DescribeAtiCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiCertificateResponse{}
	_body, _err := DescribeAtiCertificateWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a real-name verified registrant.
//
// @param request - DescribeAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiRegistrantResponse
func DescribeAtiRegistrantWithOptions(client *Client, request *DescribeAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiRegistrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a real-name verified registrant.
//
// @param request - DescribeAtiRegistrantRequest
//
// @return DescribeAtiRegistrantResponse
func DescribeAtiRegistrant(client *Client, request *DescribeAtiRegistrantRequest) (_result *DescribeAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiRegistrantResponse{}
	_body, _err := DescribeAtiRegistrantWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution results of a batch operation task using a task ID. If you do not specify a task ID, the results of the most recent batch task are returned.
//
// @param request - DescribeBatchResultCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBatchResultCountResponse
func DescribeBatchResultCountWithOptions(client *Client, request *DescribeBatchResultCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeBatchResultCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchType) {
		query["BatchType"] = request.BatchType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBatchResultCount"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBatchResultCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution results of a batch operation task using a task ID. If you do not specify a task ID, the results of the most recent batch task are returned.
//
// @param request - DescribeBatchResultCountRequest
//
// @return DescribeBatchResultCountResponse
func DescribeBatchResultCount(client *Client, request *DescribeBatchResultCountRequest) (_result *DescribeBatchResultCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBatchResultCountResponse{}
	_body, _err := DescribeBatchResultCountWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a batch processing result.
//
// Description:
//
// *Prerequisite: You can call this operation after the batch task is complete.**
//
// @param request - DescribeBatchResultDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBatchResultDetailResponse
func DescribeBatchResultDetailWithOptions(client *Client, request *DescribeBatchResultDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeBatchResultDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchType) {
		query["BatchType"] = request.BatchType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBatchResultDetail"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBatchResultDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a batch processing result.
//
// Description:
//
// *Prerequisite: You can call this operation after the batch task is complete.**
//
// @param request - DescribeBatchResultDetailRequest
//
// @return DescribeBatchResultDetailResponse
func DescribeBatchResultDetail(client *Client, request *DescribeBatchResultDetailRequest) (_result *DescribeBatchResultDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBatchResultDetailResponse{}
	_body, _err := DescribeBatchResultDetailWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of an address based on the specified input parameters.
//
// @param request - DescribeCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressResponse
func DescribeCloudGtmAddressWithOptions(client *Client, request *DescribeCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of an address based on the specified input parameters.
//
// @param request - DescribeCloudGtmAddressRequest
//
// @return DescribeCloudGtmAddressResponse
func DescribeCloudGtmAddress(client *Client, request *DescribeCloudGtmAddressRequest) (_result *DescribeCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressResponse{}
	_body, _err := DescribeCloudGtmAddressWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified address pool.
//
// @param request - DescribeCloudGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressPoolResponse
func DescribeCloudGtmAddressPoolWithOptions(client *Client, request *DescribeCloudGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified address pool.
//
// @param request - DescribeCloudGtmAddressPoolRequest
//
// @return DescribeCloudGtmAddressPoolResponse
func DescribeCloudGtmAddressPool(client *Client, request *DescribeCloudGtmAddressPoolRequest) (_result *DescribeCloudGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressPoolResponse{}
	_body, _err := DescribeCloudGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about the instances that reference an address pool.
//
// @param request - DescribeCloudGtmAddressPoolReferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressPoolReferenceResponse
func DescribeCloudGtmAddressPoolReferenceWithOptions(client *Client, request *DescribeCloudGtmAddressPoolReferenceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressPoolReferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddressPoolReference"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressPoolReferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about the instances that reference an address pool.
//
// @param request - DescribeCloudGtmAddressPoolReferenceRequest
//
// @return DescribeCloudGtmAddressPoolReferenceResponse
func DescribeCloudGtmAddressPoolReference(client *Client, request *DescribeCloudGtmAddressPoolReferenceRequest) (_result *DescribeCloudGtmAddressPoolReferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressPoolReferenceResponse{}
	_body, _err := DescribeCloudGtmAddressPoolReferenceWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the address pools and Global Traffic Manager (GTM) 3.0 instances that reference a specified address.
//
// @param request - DescribeCloudGtmAddressReferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressReferenceResponse
func DescribeCloudGtmAddressReferenceWithOptions(client *Client, request *DescribeCloudGtmAddressReferenceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressReferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddressReference"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressReferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the address pools and Global Traffic Manager (GTM) 3.0 instances that reference a specified address.
//
// @param request - DescribeCloudGtmAddressReferenceRequest
//
// @return DescribeCloudGtmAddressReferenceResponse
func DescribeCloudGtmAddressReference(client *Client, request *DescribeCloudGtmAddressReferenceRequest) (_result *DescribeCloudGtmAddressReferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressReferenceResponse{}
	_body, _err := DescribeCloudGtmAddressReferenceWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global alert configuration for Global Traffic Manager (GTM).
//
// @param request - DescribeCloudGtmGlobalAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmGlobalAlertResponse
func DescribeCloudGtmGlobalAlertWithOptions(client *Client, request *DescribeCloudGtmGlobalAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmGlobalAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmGlobalAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmGlobalAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global alert configuration for Global Traffic Manager (GTM).
//
// @param request - DescribeCloudGtmGlobalAlertRequest
//
// @return DescribeCloudGtmGlobalAlertResponse
func DescribeCloudGtmGlobalAlert(client *Client, request *DescribeCloudGtmGlobalAlertRequest) (_result *DescribeCloudGtmGlobalAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmGlobalAlertResponse{}
	_body, _err := DescribeCloudGtmGlobalAlertWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert configuration for an instance.
//
// @param request - DescribeCloudGtmInstanceConfigAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmInstanceConfigAlertResponse
func DescribeCloudGtmInstanceConfigAlertWithOptions(client *Client, request *DescribeCloudGtmInstanceConfigAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmInstanceConfigAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmInstanceConfigAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmInstanceConfigAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert configuration for an instance.
//
// @param request - DescribeCloudGtmInstanceConfigAlertRequest
//
// @return DescribeCloudGtmInstanceConfigAlertResponse
func DescribeCloudGtmInstanceConfigAlert(client *Client, request *DescribeCloudGtmInstanceConfigAlertRequest) (_result *DescribeCloudGtmInstanceConfigAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmInstanceConfigAlertResponse{}
	_body, _err := DescribeCloudGtmInstanceConfigAlertWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the full configuration of a GTM 3.0 access domain name, including alert settings, address pools, and address details.
//
// @param request - DescribeCloudGtmInstanceConfigFullInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmInstanceConfigFullInfoResponse
func DescribeCloudGtmInstanceConfigFullInfoWithOptions(client *Client, request *DescribeCloudGtmInstanceConfigFullInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmInstanceConfigFullInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmInstanceConfigFullInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmInstanceConfigFullInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the full configuration of a GTM 3.0 access domain name, including alert settings, address pools, and address details.
//
// @param request - DescribeCloudGtmInstanceConfigFullInfoRequest
//
// @return DescribeCloudGtmInstanceConfigFullInfoResponse
func DescribeCloudGtmInstanceConfigFullInfo(client *Client, request *DescribeCloudGtmInstanceConfigFullInfoRequest) (_result *DescribeCloudGtmInstanceConfigFullInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmInstanceConfigFullInfoResponse{}
	_body, _err := DescribeCloudGtmInstanceConfigFullInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified health check template.
//
// @param request - DescribeCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmMonitorTemplateResponse
func DescribeCloudGtmMonitorTemplateWithOptions(client *Client, request *DescribeCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified health check template.
//
// @param request - DescribeCloudGtmMonitorTemplateRequest
//
// @return DescribeCloudGtmMonitorTemplateResponse
func DescribeCloudGtmMonitorTemplate(client *Client, request *DescribeCloudGtmMonitorTemplateRequest) (_result *DescribeCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmMonitorTemplateResponse{}
	_body, _err := DescribeCloudGtmMonitorTemplateWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCloudGtmSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmSummaryResponse
func DescribeCloudGtmSummaryWithOptions(client *Client, request *DescribeCloudGtmSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCloudGtmSummaryRequest
//
// @return DescribeCloudGtmSummaryResponse
func DescribeCloudGtmSummary(client *Client, request *DescribeCloudGtmSummaryRequest) (_result *DescribeCloudGtmSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmSummaryResponse{}
	_body, _err := DescribeCloudGtmSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the system lines supported by Global Traffic Manager (GTM).
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmSystemLinesResponse
func DescribeCloudGtmSystemLinesWithOptions(client *Client, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmSystemLinesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmSystemLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmSystemLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the system lines supported by Global Traffic Manager (GTM).
//
// @return DescribeCloudGtmSystemLinesResponse
func DescribeCloudGtmSystemLines(client *Client) (_result *DescribeCloudGtmSystemLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmSystemLinesResponse{}
	_body, _err := DescribeCloudGtmSystemLinesWithOptions(client, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a custom line by its unique ID.
//
// @param request - DescribeCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLineResponse
func DescribeCustomLineWithOptions(client *Client, request *DescribeCustomLineRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineId) {
		query["LineId"] = request.LineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomLine"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a custom line by its unique ID.
//
// @param request - DescribeCustomLineRequest
//
// @return DescribeCustomLineResponse
func DescribeCustomLine(client *Client, request *DescribeCustomLineRequest) (_result *DescribeCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomLineResponse{}
	_body, _err := DescribeCustomLineWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom lines for a domain name.
//
// @param request - DescribeCustomLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLinesResponse
func DescribeCustomLinesWithOptions(client *Client, request *DescribeCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom lines for a domain name.
//
// @param request - DescribeCustomLinesRequest
//
// @return DescribeCustomLinesResponse
func DescribeCustomLines(client *Client, request *DescribeCustomLinesRequest) (_result *DescribeCustomLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := DescribeCustomLinesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of subdomains that have weight configurations based on the specified parameters.
//
// @param request - DescribeDNSSLBSubDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDNSSLBSubDomainsResponse
func DescribeDNSSLBSubDomainsWithOptions(client *Client, request *DescribeDNSSLBSubDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDNSSLBSubDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDNSSLBSubDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDNSSLBSubDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of subdomains that have weight configurations based on the specified parameters.
//
// @param request - DescribeDNSSLBSubDomainsRequest
//
// @return DescribeDNSSLBSubDomainsResponse
func DescribeDNSSLBSubDomains(client *Client, request *DescribeDNSSLBSubDomainsRequest) (_result *DescribeDNSSLBSubDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDNSSLBSubDomainsResponse{}
	_body, _err := DescribeDNSSLBSubDomainsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries authoritative proxy domain names based on the specified parameters.
//
// @param request - DescribeDnsCacheDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsCacheDomainsResponse
func DescribeDnsCacheDomainsWithOptions(client *Client, request *DescribeDnsCacheDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsCacheDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsCacheDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsCacheDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries authoritative proxy domain names based on the specified parameters.
//
// @param request - DescribeDnsCacheDomainsRequest
//
// @return DescribeDnsCacheDomainsResponse
func DescribeDnsCacheDomains(client *Client, request *DescribeDnsCacheDomainsRequest) (_result *DescribeDnsCacheDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsCacheDomainsResponse{}
	_body, _err := DescribeDnsCacheDomainsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access strategies for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAccessStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAccessStrategiesResponse
func DescribeDnsGtmAccessStrategiesWithOptions(client *Client, request *DescribeDnsGtmAccessStrategiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAccessStrategies"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAccessStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access strategies for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAccessStrategiesRequest
//
// @return DescribeDnsGtmAccessStrategiesResponse
func DescribeDnsGtmAccessStrategies(client *Client, request *DescribeDnsGtmAccessStrategiesRequest) (_result *DescribeDnsGtmAccessStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategiesResponse{}
	_body, _err := DescribeDnsGtmAccessStrategiesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified access strategy.
//
// @param request - DescribeDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAccessStrategyResponse
func DescribeDnsGtmAccessStrategyWithOptions(client *Client, request *DescribeDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified access strategy.
//
// @param request - DescribeDnsGtmAccessStrategyRequest
//
// @return DescribeDnsGtmAccessStrategyResponse
func DescribeDnsGtmAccessStrategy(client *Client, request *DescribeDnsGtmAccessStrategyRequest) (_result *DescribeDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategyResponse{}
	_body, _err := DescribeDnsGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the available configurations for an access policy based on an instance ID.
//
// @param request - DescribeDnsGtmAccessStrategyAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAccessStrategyAvailableConfigResponse
func DescribeDnsGtmAccessStrategyAvailableConfigWithOptions(client *Client, request *DescribeDnsGtmAccessStrategyAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategyAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAccessStrategyAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the available configurations for an access policy based on an instance ID.
//
// @param request - DescribeDnsGtmAccessStrategyAvailableConfigRequest
//
// @return DescribeDnsGtmAccessStrategyAvailableConfigResponse
func DescribeDnsGtmAccessStrategyAvailableConfig(client *Client, request *DescribeDnsGtmAccessStrategyAvailableConfigRequest) (_result *DescribeDnsGtmAccessStrategyAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := DescribeDnsGtmAccessStrategyAvailableConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the region where an address is located.
//
// @param request - DescribeDnsGtmAddrAttributeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAddrAttributeInfoResponse
func DescribeDnsGtmAddrAttributeInfoWithOptions(client *Client, request *DescribeDnsGtmAddrAttributeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAddrAttributeInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addrs) {
		query["Addrs"] = request.Addrs
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAddrAttributeInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAddrAttributeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the region where an address is located.
//
// @param request - DescribeDnsGtmAddrAttributeInfoRequest
//
// @return DescribeDnsGtmAddrAttributeInfoResponse
func DescribeDnsGtmAddrAttributeInfo(client *Client, request *DescribeDnsGtmAddrAttributeInfoRequest) (_result *DescribeDnsGtmAddrAttributeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAddrAttributeInfoResponse{}
	_body, _err := DescribeDnsGtmAddrAttributeInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an address pool in a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAddressPoolAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAddressPoolAvailableConfigResponse
func DescribeDnsGtmAddressPoolAvailableConfigWithOptions(client *Client, request *DescribeDnsGtmAddressPoolAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAddressPoolAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAddressPoolAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAddressPoolAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an address pool in a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAddressPoolAvailableConfigRequest
//
// @return DescribeDnsGtmAddressPoolAvailableConfigResponse
func DescribeDnsGtmAddressPoolAvailableConfig(client *Client, request *DescribeDnsGtmAddressPoolAvailableConfigRequest) (_result *DescribeDnsGtmAddressPoolAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAddressPoolAvailableConfigResponse{}
	_body, _err := DescribeDnsGtmAddressPoolAvailableConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available alert contact groups.
//
// @param request - DescribeDnsGtmAvailableAlertGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAvailableAlertGroupResponse
func DescribeDnsGtmAvailableAlertGroupWithOptions(client *Client, request *DescribeDnsGtmAvailableAlertGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAvailableAlertGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAvailableAlertGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAvailableAlertGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available alert contact groups.
//
// @param request - DescribeDnsGtmAvailableAlertGroupRequest
//
// @return DescribeDnsGtmAvailableAlertGroupResponse
func DescribeDnsGtmAvailableAlertGroup(client *Client, request *DescribeDnsGtmAvailableAlertGroupRequest) (_result *DescribeDnsGtmAvailableAlertGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAvailableAlertGroupResponse{}
	_body, _err := DescribeDnsGtmAvailableAlertGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an instance based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceResponse
func DescribeDnsGtmInstanceWithOptions(client *Client, request *DescribeDnsGtmInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstance"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an instance based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceRequest
//
// @return DescribeDnsGtmInstanceResponse
func DescribeDnsGtmInstance(client *Client, request *DescribeDnsGtmInstanceRequest) (_result *DescribeDnsGtmInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceResponse{}
	_body, _err := DescribeDnsGtmInstanceWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an address pool by its ID.
//
// @param request - DescribeDnsGtmInstanceAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceAddressPoolResponse
func DescribeDnsGtmInstanceAddressPoolWithOptions(client *Client, request *DescribeDnsGtmInstanceAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstanceAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an address pool by its ID.
//
// @param request - DescribeDnsGtmInstanceAddressPoolRequest
//
// @return DescribeDnsGtmInstanceAddressPoolResponse
func DescribeDnsGtmInstanceAddressPool(client *Client, request *DescribeDnsGtmInstanceAddressPoolRequest) (_result *DescribeDnsGtmInstanceAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceAddressPoolResponse{}
	_body, _err := DescribeDnsGtmInstanceAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmInstanceAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceAddressPoolsResponse
func DescribeDnsGtmInstanceAddressPoolsWithOptions(client *Client, request *DescribeDnsGtmInstanceAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceAddressPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstanceAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmInstanceAddressPoolsRequest
//
// @return DescribeDnsGtmInstanceAddressPoolsResponse
func DescribeDnsGtmInstanceAddressPools(client *Client, request *DescribeDnsGtmInstanceAddressPoolsRequest) (_result *DescribeDnsGtmInstanceAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceAddressPoolsResponse{}
	_body, _err := DescribeDnsGtmInstanceAddressPoolsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the status of an instance based on its ID.
//
// @param request - DescribeDnsGtmInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceStatusResponse
func DescribeDnsGtmInstanceStatusWithOptions(client *Client, request *DescribeDnsGtmInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstanceStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the status of an instance based on its ID.
//
// @param request - DescribeDnsGtmInstanceStatusRequest
//
// @return DescribeDnsGtmInstanceStatusResponse
func DescribeDnsGtmInstanceStatus(client *Client, request *DescribeDnsGtmInstanceStatusRequest) (_result *DescribeDnsGtmInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceStatusResponse{}
	_body, _err := DescribeDnsGtmInstanceStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceSystemCnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceSystemCnameResponse
func DescribeDnsGtmInstanceSystemCnameWithOptions(client *Client, request *DescribeDnsGtmInstanceSystemCnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceSystemCnameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstanceSystemCname"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceSystemCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceSystemCnameRequest
//
// @return DescribeDnsGtmInstanceSystemCnameResponse
func DescribeDnsGtmInstanceSystemCname(client *Client, request *DescribeDnsGtmInstanceSystemCnameRequest) (_result *DescribeDnsGtmInstanceSystemCnameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceSystemCnameResponse{}
	_body, _err := DescribeDnsGtmInstanceSystemCnameWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - DescribeDnsGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstancesResponse
func DescribeDnsGtmInstancesWithOptions(client *Client, request *DescribeDnsGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - DescribeDnsGtmInstancesRequest
//
// @return DescribeDnsGtmInstancesResponse
func DescribeDnsGtmInstances(client *Client, request *DescribeDnsGtmInstancesRequest) (_result *DescribeDnsGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstancesResponse{}
	_body, _err := DescribeDnsGtmInstancesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs for an instance.
//
// @param request - DescribeDnsGtmLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmLogsResponse
func DescribeDnsGtmLogsWithOptions(client *Client, request *DescribeDnsGtmLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs for an instance.
//
// @param request - DescribeDnsGtmLogsRequest
//
// @return DescribeDnsGtmLogsResponse
func DescribeDnsGtmLogs(client *Client, request *DescribeDnsGtmLogsRequest) (_result *DescribeDnsGtmLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmLogsResponse{}
	_body, _err := DescribeDnsGtmLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available configurations for DNS health checks.
//
// @param request - DescribeDnsGtmMonitorAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmMonitorAvailableConfigResponse
func DescribeDnsGtmMonitorAvailableConfigWithOptions(client *Client, request *DescribeDnsGtmMonitorAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmMonitorAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmMonitorAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmMonitorAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available configurations for DNS health checks.
//
// @param request - DescribeDnsGtmMonitorAvailableConfigRequest
//
// @return DescribeDnsGtmMonitorAvailableConfigResponse
func DescribeDnsGtmMonitorAvailableConfig(client *Client, request *DescribeDnsGtmMonitorAvailableConfigRequest) (_result *DescribeDnsGtmMonitorAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmMonitorAvailableConfigResponse{}
	_body, _err := DescribeDnsGtmMonitorAvailableConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health check configuration for an address pool.
//
// @param request - DescribeDnsGtmMonitorConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmMonitorConfigResponse
func DescribeDnsGtmMonitorConfigWithOptions(client *Client, request *DescribeDnsGtmMonitorConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmMonitorConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmMonitorConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health check configuration for an address pool.
//
// @param request - DescribeDnsGtmMonitorConfigRequest
//
// @return DescribeDnsGtmMonitorConfigResponse
func DescribeDnsGtmMonitorConfig(client *Client, request *DescribeDnsGtmMonitorConfigRequest) (_result *DescribeDnsGtmMonitorConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmMonitorConfigResponse{}
	_body, _err := DescribeDnsGtmMonitorConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a paid Alibaba Cloud DNS instance by its instance ID.
//
// @param request - DescribeDnsProductInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsProductInstanceResponse
func DescribeDnsProductInstanceWithOptions(client *Client, request *DescribeDnsProductInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsProductInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsProductInstance"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsProductInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a paid Alibaba Cloud DNS instance by its instance ID.
//
// @param request - DescribeDnsProductInstanceRequest
//
// @return DescribeDnsProductInstanceResponse
func DescribeDnsProductInstance(client *Client, request *DescribeDnsProductInstanceRequest) (_result *DescribeDnsProductInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsProductInstanceResponse{}
	_body, _err := DescribeDnsProductInstanceWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of paid DNS product instances that match the specified parameters.
//
// Description:
//
// > **If the response does not contain a domain name, the Alibaba Cloud DNS instance is not associated with any domain names.**
//
// @param request - DescribeDnsProductInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsProductInstancesResponse
func DescribeDnsProductInstancesWithOptions(client *Client, request *DescribeDnsProductInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsProductInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsProductInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsProductInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of paid DNS product instances that match the specified parameters.
//
// Description:
//
// > **If the response does not contain a domain name, the Alibaba Cloud DNS instance is not associated with any domain names.**
//
// @param request - DescribeDnsProductInstancesRequest
//
// @return DescribeDnsProductInstancesResponse
func DescribeDnsProductInstances(client *Client, request *DescribeDnsProductInstancesRequest) (_result *DescribeDnsProductInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsProductInstancesResponse{}
	_body, _err := DescribeDnsProductInstancesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an overview of request statistics for a DNS over HTTPS (DoH) account.
//
// @param request - DescribeDohAccountStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohAccountStatisticsResponse
func DescribeDohAccountStatisticsWithOptions(client *Client, request *DescribeDohAccountStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohAccountStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohAccountStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohAccountStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an overview of request statistics for a DNS over HTTPS (DoH) account.
//
// @param request - DescribeDohAccountStatisticsRequest
//
// @return DescribeDohAccountStatisticsResponse
func DescribeDohAccountStatistics(client *Client, request *DescribeDohAccountStatisticsRequest) (_result *DescribeDohAccountStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohAccountStatisticsResponse{}
	_body, _err := DescribeDohAccountStatisticsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an overview of statistics for DNS over HTTPS (DoH) requests for a domain name.
//
// @param request - DescribeDohDomainStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohDomainStatisticsResponse
func DescribeDohDomainStatisticsWithOptions(client *Client, request *DescribeDohDomainStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohDomainStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohDomainStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohDomainStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an overview of statistics for DNS over HTTPS (DoH) requests for a domain name.
//
// @param request - DescribeDohDomainStatisticsRequest
//
// @return DescribeDohDomainStatisticsResponse
func DescribeDohDomainStatistics(client *Client, request *DescribeDohDomainStatisticsRequest) (_result *DescribeDohDomainStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohDomainStatisticsResponse{}
	_body, _err := DescribeDohDomainStatisticsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries request statistics for DNS over HTTPS (DoH) domain names.
//
// @param request - DescribeDohDomainStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohDomainStatisticsSummaryResponse
func DescribeDohDomainStatisticsSummaryWithOptions(client *Client, request *DescribeDohDomainStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohDomainStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohDomainStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohDomainStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries request statistics for DNS over HTTPS (DoH) domain names.
//
// @param request - DescribeDohDomainStatisticsSummaryRequest
//
// @return DescribeDohDomainStatisticsSummaryResponse
func DescribeDohDomainStatisticsSummary(client *Client, request *DescribeDohDomainStatisticsSummaryRequest) (_result *DescribeDohDomainStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohDomainStatisticsSummaryResponse{}
	_body, _err := DescribeDohDomainStatisticsSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries for statistics on DNS over HTTPS (DoH) requests for a subdomain.
//
// @param request - DescribeDohSubDomainStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohSubDomainStatisticsResponse
func DescribeDohSubDomainStatisticsWithOptions(client *Client, request *DescribeDohSubDomainStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohSubDomainStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohSubDomainStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohSubDomainStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries for statistics on DNS over HTTPS (DoH) requests for a subdomain.
//
// @param request - DescribeDohSubDomainStatisticsRequest
//
// @return DescribeDohSubDomainStatisticsResponse
func DescribeDohSubDomainStatistics(client *Client, request *DescribeDohSubDomainStatisticsRequest) (_result *DescribeDohSubDomainStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohSubDomainStatisticsResponse{}
	_body, _err := DescribeDohSubDomainStatisticsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a summary of request statistics for subdomains using DNS over HTTPS (DoH).
//
// @param request - DescribeDohSubDomainStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohSubDomainStatisticsSummaryResponse
func DescribeDohSubDomainStatisticsSummaryWithOptions(client *Client, request *DescribeDohSubDomainStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohSubDomainStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohSubDomainStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohSubDomainStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a summary of request statistics for subdomains using DNS over HTTPS (DoH).
//
// @param request - DescribeDohSubDomainStatisticsSummaryRequest
//
// @return DescribeDohSubDomainStatisticsSummaryResponse
func DescribeDohSubDomainStatisticsSummary(client *Client, request *DescribeDohSubDomainStatisticsSummaryRequest) (_result *DescribeDohSubDomainStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohSubDomainStatisticsSummaryResponse{}
	_body, _err := DescribeDohSubDomainStatisticsSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of a DNS over HTTPS (DoH) user.
//
// @param request - DescribeDohUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohUserInfoResponse
func DescribeDohUserInfoWithOptions(client *Client, request *DescribeDohUserInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohUserInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of a DNS over HTTPS (DoH) user.
//
// @param request - DescribeDohUserInfoRequest
//
// @return DescribeDohUserInfoResponse
func DescribeDohUserInfo(client *Client, request *DescribeDohUserInfoRequest) (_result *DescribeDohUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohUserInfoResponse{}
	_body, _err := DescribeDohUserInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Domain Name System Security Extensions (DNSSEC) information for a specified domain name.
//
// @param request - DescribeDomainDnssecInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainDnssecInfoResponse
func DescribeDomainDnssecInfoWithOptions(client *Client, request *DescribeDomainDnssecInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainDnssecInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainDnssecInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainDnssecInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Domain Name System Security Extensions (DNSSEC) information for a specified domain name.
//
// @param request - DescribeDomainDnssecInfoRequest
//
// @return DescribeDomainDnssecInfoResponse
func DescribeDomainDnssecInfo(client *Client, request *DescribeDomainDnssecInfoRequest) (_result *DescribeDomainDnssecInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainDnssecInfoResponse{}
	_body, _err := DescribeDomainDnssecInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain name groups.
//
// @param request - DescribeDomainGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainGroupsResponse
func DescribeDomainGroupsWithOptions(client *Client, request *DescribeDomainGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainGroups"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain name groups.
//
// @param request - DescribeDomainGroupsRequest
//
// @return DescribeDomainGroupsResponse
func DescribeDomainGroups(client *Client, request *DescribeDomainGroupsRequest) (_result *DescribeDomainGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainGroupsResponse{}
	_body, _err := DescribeDomainGroupsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specified domain name.
//
// Description:
//
// In this example, the domain name is bound to an instance of Alibaba Cloud DNS Ultimate Edition. For more information about line enumeration, see the RecordLines response parameter.
//
// @param request - DescribeDomainInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainInfoResponse
func DescribeDomainInfoWithOptions(client *Client, request *DescribeDomainInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NeedDetailAttributes) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specified domain name.
//
// Description:
//
// In this example, the domain name is bound to an instance of Alibaba Cloud DNS Ultimate Edition. For more information about line enumeration, see the RecordLines response parameter.
//
// @param request - DescribeDomainInfoRequest
//
// @return DescribeDomainInfoResponse
func DescribeDomainInfo(client *Client, request *DescribeDomainInfoRequest) (_result *DescribeDomainInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainInfoResponse{}
	_body, _err := DescribeDomainInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs for a domain name based on the specified parameters.
//
// @param request - DescribeDomainLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainLogsResponse
func DescribeDomainLogsWithOptions(client *Client, request *DescribeDomainLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs for a domain name based on the specified parameters.
//
// @param request - DescribeDomainLogsRequest
//
// @return DescribeDomainLogsResponse
func DescribeDomainLogs(client *Client, request *DescribeDomainLogsRequest) (_result *DescribeDomainLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainLogsResponse{}
	_body, _err := DescribeDomainLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current list of name servers for a domain name and determines whether the servers are managed by Alibaba Cloud DNS.
//
// Description:
//
// > This operation directly queries the authoritative server of the domain name registry to retrieve the DNS server names for the domain name. An error may be returned if the domain name is inactive. For example, if the domain name has a serverHold or clientHold status, or has not passed identity verification.
//
// @param request - DescribeDomainNsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainNsResponse
func DescribeDomainNsWithOptions(client *Client, request *DescribeDomainNsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainNsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainNs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainNsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current list of name servers for a domain name and determines whether the servers are managed by Alibaba Cloud DNS.
//
// Description:
//
// > This operation directly queries the authoritative server of the domain name registry to retrieve the DNS server names for the domain name. An error may be returned if the domain name is inactive. For example, if the domain name has a serverHold or clientHold status, or has not passed identity verification.
//
// @param request - DescribeDomainNsRequest
//
// @return DescribeDomainNsResponse
func DescribeDomainNs(client *Client, request *DescribeDomainNsRequest) (_result *DescribeDomainNsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainNsResponse{}
	_body, _err := DescribeDomainNsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a DNS record by its ID.
//
// @param request - DescribeDomainRecordInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRecordInfoResponse
func DescribeDomainRecordInfoWithOptions(client *Client, request *DescribeDomainRecordInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRecordInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRecordInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRecordInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a DNS record by its ID.
//
// @param request - DescribeDomainRecordInfoRequest
//
// @return DescribeDomainRecordInfoResponse
func DescribeDomainRecordInfo(client *Client, request *DescribeDomainRecordInfoRequest) (_result *DescribeDomainRecordInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRecordInfoResponse{}
	_body, _err := DescribeDomainRecordInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the DNS records for a specified root domain based on the input parameters.
//
// Description:
//
// - You can specify the domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve a list of DNS records.
//
// - You can specify a keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord) to query DNS records that contain the keyword.
//
// - By default, DNS records are sorted in descending order by the time they were added.
//
// - You can specify a domain group ID (GroupId) to query the DNS records in a specific group.
//
// @param request - DescribeDomainRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRecordsResponse
func DescribeDomainRecordsWithOptions(client *Client, request *DescribeDomainRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RRKeyWord) {
		query["RRKeyWord"] = request.RRKeyWord
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.TypeKeyWord) {
		query["TypeKeyWord"] = request.TypeKeyWord
	}

	if !dara.IsNil(request.ValueKeyWord) {
		query["ValueKeyWord"] = request.ValueKeyWord
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the DNS records for a specified root domain based on the input parameters.
//
// Description:
//
// - You can specify the domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve a list of DNS records.
//
// - You can specify a keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord) to query DNS records that contain the keyword.
//
// - By default, DNS records are sorted in descending order by the time they were added.
//
// - You can specify a domain group ID (GroupId) to query the DNS records in a specific group.
//
// @param request - DescribeDomainRecordsRequest
//
// @return DescribeDomainRecordsResponse
func DescribeDomainRecords(client *Client, request *DescribeDomainRecordsRequest) (_result *DescribeDomainRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRecordsResponse{}
	_body, _err := DescribeDomainRecordsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request volumes for all paid domain names in your account.
//
// @param request - DescribeDomainResolveStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResolveStatisticsSummaryResponse
func DescribeDomainResolveStatisticsSummaryWithOptions(client *Client, request *DescribeDomainResolveStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainResolveStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainResolveStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainResolveStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request volumes for all paid domain names in your account.
//
// @param request - DescribeDomainResolveStatisticsSummaryRequest
//
// @return DescribeDomainResolveStatisticsSummaryResponse
func DescribeDomainResolveStatisticsSummary(client *Client, request *DescribeDomainResolveStatisticsSummaryRequest) (_result *DescribeDomainResolveStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainResolveStatisticsSummaryResponse{}
	_body, _err := DescribeDomainResolveStatisticsSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of real-time requests for a specified primary domain name.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeDomainStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainStatisticsResponse
func DescribeDomainStatisticsWithOptions(client *Client, request *DescribeDomainStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of real-time requests for a specified primary domain name.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeDomainStatisticsRequest
//
// @return DescribeDomainStatisticsResponse
func DescribeDomainStatistics(client *Client, request *DescribeDomainStatisticsRequest) (_result *DescribeDomainStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainStatisticsResponse{}
	_body, _err := DescribeDomainStatisticsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of requests for all paid domain names in your account.
//
// @param request - DescribeDomainStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainStatisticsSummaryResponse
func DescribeDomainStatisticsSummaryWithOptions(client *Client, request *DescribeDomainStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of requests for all paid domain names in your account.
//
// @param request - DescribeDomainStatisticsSummaryRequest
//
// @return DescribeDomainStatisticsSummaryResponse
func DescribeDomainStatisticsSummary(client *Client, request *DescribeDomainStatisticsSummaryRequest) (_result *DescribeDomainStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainStatisticsSummaryResponse{}
	_body, _err := DescribeDomainStatisticsSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of domain names based on specified parameters.
//
// Description:
//
// 1. You can specify a page number (PageNumber) and page size (PageSize) to retrieve a paginated list of domain names.
//
// 2. You can specify a keyword (KeyWord) to query for domain names that contain the specified keyword.
//
// 3. By default, domain names are sorted in descending order of their creation time.
//
// 4. You can specify a domain name group ID (GroupId) to query for domain names in a specific group. This lets you retrieve all domain names or only the domain names that are not assigned to a group.
//
// @param request - DescribeDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainsResponse
func DescribeDomainsWithOptions(client *Client, request *DescribeDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.Starmark) {
		query["Starmark"] = request.Starmark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain names based on specified parameters.
//
// Description:
//
// 1. You can specify a page number (PageNumber) and page size (PageSize) to retrieve a paginated list of domain names.
//
// 2. You can specify a keyword (KeyWord) to query for domain names that contain the specified keyword.
//
// 3. By default, domain names are sorted in descending order of their creation time.
//
// 4. You can specify a domain name group ID (GroupId) to query for domain names in a specific group. This lets you retrieve all domain names or only the domain names that are not assigned to a group.
//
// @param request - DescribeDomainsRequest
//
// @return DescribeDomainsResponse
func DescribeDomains(client *Client, request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := DescribeDomainsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access policies for an instance.
//
// @param request - DescribeGtmAccessStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAccessStrategiesResponse
func DescribeGtmAccessStrategiesWithOptions(client *Client, request *DescribeGtmAccessStrategiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAccessStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmAccessStrategies"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAccessStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access policies for an instance.
//
// @param request - DescribeGtmAccessStrategiesRequest
//
// @return DescribeGtmAccessStrategiesResponse
func DescribeGtmAccessStrategies(client *Client, request *DescribeGtmAccessStrategiesRequest) (_result *DescribeGtmAccessStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategiesResponse{}
	_body, _err := DescribeGtmAccessStrategiesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an access policy based on the policy ID.
//
// @param request - DescribeGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAccessStrategyResponse
func DescribeGtmAccessStrategyWithOptions(client *Client, request *DescribeGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access policy based on the policy ID.
//
// @param request - DescribeGtmAccessStrategyRequest
//
// @return DescribeGtmAccessStrategyResponse
func DescribeGtmAccessStrategy(client *Client, request *DescribeGtmAccessStrategyRequest) (_result *DescribeGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategyResponse{}
	_body, _err := DescribeGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an access policy.
//
// @param request - DescribeGtmAccessStrategyAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAccessStrategyAvailableConfigResponse
func DescribeGtmAccessStrategyAvailableConfigWithOptions(client *Client, request *DescribeGtmAccessStrategyAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAccessStrategyAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmAccessStrategyAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an access policy.
//
// @param request - DescribeGtmAccessStrategyAvailableConfigRequest
//
// @return DescribeGtmAccessStrategyAvailableConfigResponse
func DescribeGtmAccessStrategyAvailableConfig(client *Client, request *DescribeGtmAccessStrategyAvailableConfigRequest) (_result *DescribeGtmAccessStrategyAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := DescribeGtmAccessStrategyAvailableConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of available alert contact groups for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmAvailableAlertGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAvailableAlertGroupResponse
func DescribeGtmAvailableAlertGroupWithOptions(client *Client, request *DescribeGtmAvailableAlertGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAvailableAlertGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmAvailableAlertGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAvailableAlertGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of available alert contact groups for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmAvailableAlertGroupRequest
//
// @return DescribeGtmAvailableAlertGroupResponse
func DescribeGtmAvailableAlertGroup(client *Client, request *DescribeGtmAvailableAlertGroupRequest) (_result *DescribeGtmAvailableAlertGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAvailableAlertGroupResponse{}
	_body, _err := DescribeGtmAvailableAlertGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceResponse
func DescribeGtmInstanceWithOptions(client *Client, request *DescribeGtmInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NeedDetailAttributes) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstance"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceRequest
//
// @return DescribeGtmInstanceResponse
func DescribeGtmInstance(client *Client, request *DescribeGtmInstanceRequest) (_result *DescribeGtmInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceResponse{}
	_body, _err := DescribeGtmInstanceWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the details of an address pool.
//
// @param request - DescribeGtmInstanceAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceAddressPoolResponse
func DescribeGtmInstanceAddressPoolWithOptions(client *Client, request *DescribeGtmInstanceAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstanceAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the details of an address pool.
//
// @param request - DescribeGtmInstanceAddressPoolRequest
//
// @return DescribeGtmInstanceAddressPoolResponse
func DescribeGtmInstanceAddressPool(client *Client, request *DescribeGtmInstanceAddressPoolRequest) (_result *DescribeGtmInstanceAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceAddressPoolResponse{}
	_body, _err := DescribeGtmInstanceAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceAddressPoolsResponse
func DescribeGtmInstanceAddressPoolsWithOptions(client *Client, request *DescribeGtmInstanceAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceAddressPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstanceAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceAddressPoolsRequest
//
// @return DescribeGtmInstanceAddressPoolsResponse
func DescribeGtmInstanceAddressPools(client *Client, request *DescribeGtmInstanceAddressPoolsRequest) (_result *DescribeGtmInstanceAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceAddressPoolsResponse{}
	_body, _err := DescribeGtmInstanceAddressPoolsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current status of an instance.
//
// @param request - DescribeGtmInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceStatusResponse
func DescribeGtmInstanceStatusWithOptions(client *Client, request *DescribeGtmInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstanceStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current status of an instance.
//
// @param request - DescribeGtmInstanceStatusRequest
//
// @return DescribeGtmInstanceStatusResponse
func DescribeGtmInstanceStatus(client *Client, request *DescribeGtmInstanceStatusRequest) (_result *DescribeGtmInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceStatusResponse{}
	_body, _err := DescribeGtmInstanceStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME domain name.
//
// @param request - DescribeGtmInstanceSystemCnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceSystemCnameResponse
func DescribeGtmInstanceSystemCnameWithOptions(client *Client, request *DescribeGtmInstanceSystemCnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceSystemCnameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstanceSystemCname"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceSystemCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME domain name.
//
// @param request - DescribeGtmInstanceSystemCnameRequest
//
// @return DescribeGtmInstanceSystemCnameResponse
func DescribeGtmInstanceSystemCname(client *Client, request *DescribeGtmInstanceSystemCnameRequest) (_result *DescribeGtmInstanceSystemCnameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceSystemCnameResponse{}
	_body, _err := DescribeGtmInstanceSystemCnameWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Global Traffic Manager (GTM) instances.
//
// @param request - DescribeGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstancesResponse
func DescribeGtmInstancesWithOptions(client *Client, request *DescribeGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NeedDetailAttributes) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Global Traffic Manager (GTM) instances.
//
// @param request - DescribeGtmInstancesRequest
//
// @return DescribeGtmInstancesResponse
func DescribeGtmInstances(client *Client, request *DescribeGtmInstancesRequest) (_result *DescribeGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstancesResponse{}
	_body, _err := DescribeGtmInstancesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of logs.
//
// @param request - DescribeGtmLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmLogsResponse
func DescribeGtmLogsWithOptions(client *Client, request *DescribeGtmLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of logs.
//
// @param request - DescribeGtmLogsRequest
//
// @return DescribeGtmLogsResponse
func DescribeGtmLogs(client *Client, request *DescribeGtmLogsRequest) (_result *DescribeGtmLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmLogsResponse{}
	_body, _err := DescribeGtmLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the available configurations for health checks.
//
// @param request - DescribeGtmMonitorAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmMonitorAvailableConfigResponse
func DescribeGtmMonitorAvailableConfigWithOptions(client *Client, request *DescribeGtmMonitorAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmMonitorAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmMonitorAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmMonitorAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the available configurations for health checks.
//
// @param request - DescribeGtmMonitorAvailableConfigRequest
//
// @return DescribeGtmMonitorAvailableConfigResponse
func DescribeGtmMonitorAvailableConfig(client *Client, request *DescribeGtmMonitorAvailableConfigRequest) (_result *DescribeGtmMonitorAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmMonitorAvailableConfigResponse{}
	_body, _err := DescribeGtmMonitorAvailableConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the health check configuration for an address pool.
//
// @param request - DescribeGtmMonitorConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmMonitorConfigResponse
func DescribeGtmMonitorConfigWithOptions(client *Client, request *DescribeGtmMonitorConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmMonitorConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmMonitorConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the health check configuration for an address pool.
//
// @param request - DescribeGtmMonitorConfigRequest
//
// @return DescribeGtmMonitorConfigResponse
func DescribeGtmMonitorConfig(client *Client, request *DescribeGtmMonitorConfigRequest) (_result *DescribeGtmMonitorConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmMonitorConfigResponse{}
	_body, _err := DescribeGtmMonitorConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmRecoveryPlanResponse
func DescribeGtmRecoveryPlanWithOptions(client *Client, request *DescribeGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanRequest
//
// @return DescribeGtmRecoveryPlanResponse
func DescribeGtmRecoveryPlan(client *Client, request *DescribeGtmRecoveryPlanRequest) (_result *DescribeGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlanResponse{}
	_body, _err := DescribeGtmRecoveryPlanWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the available configurations for a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmRecoveryPlanAvailableConfigResponse
func DescribeGtmRecoveryPlanAvailableConfigWithOptions(client *Client, request *DescribeGtmRecoveryPlanAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmRecoveryPlanAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmRecoveryPlanAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmRecoveryPlanAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the available configurations for a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanAvailableConfigRequest
//
// @return DescribeGtmRecoveryPlanAvailableConfigResponse
func DescribeGtmRecoveryPlanAvailableConfig(client *Client, request *DescribeGtmRecoveryPlanAvailableConfigRequest) (_result *DescribeGtmRecoveryPlanAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlanAvailableConfigResponse{}
	_body, _err := DescribeGtmRecoveryPlanAvailableConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of disaster recovery plans.
//
// @param request - DescribeGtmRecoveryPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmRecoveryPlansResponse
func DescribeGtmRecoveryPlansWithOptions(client *Client, request *DescribeGtmRecoveryPlansRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmRecoveryPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmRecoveryPlans"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmRecoveryPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of disaster recovery plans.
//
// @param request - DescribeGtmRecoveryPlansRequest
//
// @return DescribeGtmRecoveryPlansResponse
func DescribeGtmRecoveryPlans(client *Client, request *DescribeGtmRecoveryPlansRequest) (_result *DescribeGtmRecoveryPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlansResponse{}
	_body, _err := DescribeGtmRecoveryPlansWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of domain names that are attached to an instance.
//
// @param request - DescribeInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDomainsResponse
func DescribeInstanceDomainsWithOptions(client *Client, request *DescribeInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainKeywords) {
		query["DomainKeywords"] = request.DomainKeywords
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of domain names that are attached to an instance.
//
// @param request - DescribeInstanceDomainsRequest
//
// @return DescribeInstanceDomainsResponse
func DescribeInstanceDomains(client *Client, request *DescribeInstanceDomainsRequest) (_result *DescribeInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceDomainsResponse{}
	_body, _err := DescribeInstanceDomainsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a global overview of statistics for public authoritative DNS.
//
// Description:
//
// Real-time data is aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsGlobalOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsGlobalOverviewResponse
func DescribeInterAuthStatisticsGlobalOverviewWithOptions(client *Client, request *DescribeInterAuthStatisticsGlobalOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsGlobalOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OverviewPeriod) {
		query["OverviewPeriod"] = request.OverviewPeriod
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsGlobalOverview"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsGlobalOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a global overview of statistics for public authoritative DNS.
//
// Description:
//
// Real-time data is aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsGlobalOverviewRequest
//
// @return DescribeInterAuthStatisticsGlobalOverviewResponse
func DescribeInterAuthStatisticsGlobalOverview(client *Client, request *DescribeInterAuthStatisticsGlobalOverviewRequest) (_result *DescribeInterAuthStatisticsGlobalOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsGlobalOverviewResponse{}
	_body, _err := DescribeInterAuthStatisticsGlobalOverviewWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Historical statistics for public authoritative DNS resolution
//
// Description:
//
// Real-time data statistics are aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsHistoryResponse
func DescribeInterAuthStatisticsHistoryWithOptions(client *Client, request *DescribeInterAuthStatisticsHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.Rcode) {
		query["Rcode"] = request.Rcode
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.StatisticalType) {
		query["StatisticalType"] = request.StatisticalType
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsHistory"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Historical statistics for public authoritative DNS resolution
//
// Description:
//
// Real-time data statistics are aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsHistoryRequest
//
// @return DescribeInterAuthStatisticsHistoryResponse
func DescribeInterAuthStatisticsHistory(client *Client, request *DescribeInterAuthStatisticsHistoryRequest) (_result *DescribeInterAuthStatisticsHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsHistoryResponse{}
	_body, _err := DescribeInterAuthStatisticsHistoryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the summary list of public authoritative parse statistics.
//
// Description:
//
// Real-time data is aggregated by hour.
//
// @param request - DescribeInterAuthStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsSummaryResponse
func DescribeInterAuthStatisticsSummaryWithOptions(client *Client, request *DescribeInterAuthStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.GrowType) {
		query["GrowType"] = request.GrowType
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Rcode) {
		query["Rcode"] = request.Rcode
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.StatisticalType) {
		query["StatisticalType"] = request.StatisticalType
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the summary list of public authoritative parse statistics.
//
// Description:
//
// Real-time data is aggregated by hour.
//
// @param request - DescribeInterAuthStatisticsSummaryRequest
//
// @return DescribeInterAuthStatisticsSummaryResponse
func DescribeInterAuthStatisticsSummary(client *Client, request *DescribeInterAuthStatisticsSummaryRequest) (_result *DescribeInterAuthStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsSummaryResponse{}
	_body, _err := DescribeInterAuthStatisticsSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a zone-level overview of statistics for public authoritative DNS.
//
// @param request - DescribeInterAuthStatisticsZoneOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsZoneOverviewResponse
func DescribeInterAuthStatisticsZoneOverviewWithOptions(client *Client, request *DescribeInterAuthStatisticsZoneOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsZoneOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OverviewPeriod) {
		query["OverviewPeriod"] = request.OverviewPeriod
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsZoneOverview"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsZoneOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a zone-level overview of statistics for public authoritative DNS.
//
// @param request - DescribeInterAuthStatisticsZoneOverviewRequest
//
// @return DescribeInterAuthStatisticsZoneOverviewResponse
func DescribeInterAuthStatisticsZoneOverview(client *Client, request *DescribeInterAuthStatisticsZoneOverviewRequest) (_result *DescribeInterAuthStatisticsZoneOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsZoneOverviewResponse{}
	_body, _err := DescribeInterAuthStatisticsZoneOverviewWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query parsing logs
//
// @param request - DescribeInternetDnsLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetDnsLogsResponse
func DescribeInternetDnsLogsWithOptions(client *Client, request *DescribeInternetDnsLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetDnsLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Module) {
		query["Module"] = request.Module
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PreciseSort) {
		query["PreciseSort"] = request.PreciseSort
	}

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.RecursionProtocolType) {
		query["RecursionProtocolType"] = request.RecursionProtocolType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetDnsLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetDnsLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query parsing logs
//
// @param request - DescribeInternetDnsLogsRequest
//
// @return DescribeInternetDnsLogsResponse
func DescribeInternetDnsLogs(client *Client, request *DescribeInternetDnsLogsRequest) (_result *DescribeInternetDnsLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInternetDnsLogsResponse{}
	_body, _err := DescribeInternetDnsLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh instances.
//
// @param request - DescribeIspFlushCacheInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheInstancesResponse
func DescribeIspFlushCacheInstancesWithOptions(client *Client, request *DescribeIspFlushCacheInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIspFlushCacheInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh instances.
//
// @param request - DescribeIspFlushCacheInstancesRequest
//
// @return DescribeIspFlushCacheInstancesResponse
func DescribeIspFlushCacheInstances(client *Client, request *DescribeIspFlushCacheInstancesRequest) (_result *DescribeIspFlushCacheInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheInstancesResponse{}
	_body, _err := DescribeIspFlushCacheInstancesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the remaining number of cache refresh operations available.
//
// @param request - DescribeIspFlushCacheRemainQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheRemainQuotaResponse
func DescribeIspFlushCacheRemainQuotaWithOptions(client *Client, request *DescribeIspFlushCacheRemainQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheRemainQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIspFlushCacheRemainQuota"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheRemainQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the remaining number of cache refresh operations available.
//
// @param request - DescribeIspFlushCacheRemainQuotaRequest
//
// @return DescribeIspFlushCacheRemainQuotaResponse
func DescribeIspFlushCacheRemainQuota(client *Client, request *DescribeIspFlushCacheRemainQuotaRequest) (_result *DescribeIspFlushCacheRemainQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheRemainQuotaResponse{}
	_body, _err := DescribeIspFlushCacheRemainQuotaWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a cache flush task.
//
// @param request - DescribeIspFlushCacheTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheTaskResponse
func DescribeIspFlushCacheTaskWithOptions(client *Client, request *DescribeIspFlushCacheTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIspFlushCacheTask"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cache flush task.
//
// @param request - DescribeIspFlushCacheTaskRequest
//
// @return DescribeIspFlushCacheTaskResponse
func DescribeIspFlushCacheTask(client *Client, request *DescribeIspFlushCacheTaskRequest) (_result *DescribeIspFlushCacheTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheTaskResponse{}
	_body, _err := DescribeIspFlushCacheTaskWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh tasks.
//
// @param request - DescribeIspFlushCacheTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheTasksResponse
func DescribeIspFlushCacheTasksWithOptions(client *Client, request *DescribeIspFlushCacheTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIspFlushCacheTasks"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh tasks.
//
// @param request - DescribeIspFlushCacheTasksRequest
//
// @return DescribeIspFlushCacheTasksResponse
func DescribeIspFlushCacheTasks(client *Client, request *DescribeIspFlushCacheTasksRequest) (_result *DescribeIspFlushCacheTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheTasksResponse{}
	_body, _err := DescribeIspFlushCacheTasksWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the overview of Public DNS user data.
//
// @param request - DescribePdnsAccountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsAccountSummaryResponse
func DescribePdnsAccountSummaryWithOptions(client *Client, request *DescribePdnsAccountSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsAccountSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsAccountSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsAccountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the overview of Public DNS user data.
//
// @param request - DescribePdnsAccountSummaryRequest
//
// @return DescribePdnsAccountSummaryResponse
func DescribePdnsAccountSummary(client *Client, request *DescribePdnsAccountSummaryRequest) (_result *DescribePdnsAccountSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsAccountSummaryResponse{}
	_body, _err := DescribePdnsAccountSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the details of a public DNS AppKey
//
// @param request - DescribePdnsAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsAppKeyResponse
func DescribePdnsAppKeyWithOptions(client *Client, request *DescribePdnsAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKeyId) {
		query["AppKeyId"] = request.AppKeyId
	}

	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsAppKey"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the details of a public DNS AppKey
//
// @param request - DescribePdnsAppKeyRequest
//
// @return DescribePdnsAppKeyResponse
func DescribePdnsAppKey(client *Client, request *DescribePdnsAppKeyRequest) (_result *DescribePdnsAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsAppKeyResponse{}
	_body, _err := DescribePdnsAppKeyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the list of AppKeys of Public DNS
//
// @param request - DescribePdnsAppKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsAppKeysResponse
func DescribePdnsAppKeysWithOptions(client *Client, request *DescribePdnsAppKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsAppKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsAppKeys"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsAppKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the list of AppKeys of Public DNS
//
// @param request - DescribePdnsAppKeysRequest
//
// @return DescribePdnsAppKeysResponse
func DescribePdnsAppKeys(client *Client, request *DescribePdnsAppKeysRequest) (_result *DescribePdnsAppKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsAppKeysResponse{}
	_body, _err := DescribePdnsAppKeysWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation log list of Public DNS.
//
// @param request - DescribePdnsOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsOperateLogsResponse
func DescribePdnsOperateLogsWithOptions(client *Client, request *DescribePdnsOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsOperateLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsOperateLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsOperateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation log list of Public DNS.
//
// @param request - DescribePdnsOperateLogsRequest
//
// @return DescribePdnsOperateLogsResponse
func DescribePdnsOperateLogs(client *Client, request *DescribePdnsOperateLogsRequest) (_result *DescribePdnsOperateLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsOperateLogsResponse{}
	_body, _err := DescribePdnsOperateLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request statistics for Public DNS.
//
// @param request - DescribePdnsRequestStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsRequestStatisticResponse
func DescribePdnsRequestStatisticWithOptions(client *Client, request *DescribePdnsRequestStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsRequestStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsRequestStatistic"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsRequestStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request statistics for Public DNS.
//
// @param request - DescribePdnsRequestStatisticRequest
//
// @return DescribePdnsRequestStatisticResponse
func DescribePdnsRequestStatistic(client *Client, request *DescribePdnsRequestStatisticRequest) (_result *DescribePdnsRequestStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsRequestStatisticResponse{}
	_body, _err := DescribePdnsRequestStatisticWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request statistics for a specified subdomain.
//
// @param request - DescribePdnsRequestStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsRequestStatisticsResponse
func DescribePdnsRequestStatisticsWithOptions(client *Client, request *DescribePdnsRequestStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsRequestStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsRequestStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsRequestStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request statistics for a specified subdomain.
//
// @param request - DescribePdnsRequestStatisticsRequest
//
// @return DescribePdnsRequestStatisticsResponse
func DescribePdnsRequestStatistics(client *Client, request *DescribePdnsRequestStatisticsRequest) (_result *DescribePdnsRequestStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsRequestStatisticsResponse{}
	_body, _err := DescribePdnsRequestStatisticsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Public DNS threat logs.
//
// @param request - DescribePdnsThreatLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsThreatLogsResponse
func DescribePdnsThreatLogsWithOptions(client *Client, request *DescribePdnsThreatLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsThreatLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.ThreatLevel) {
		query["ThreatLevel"] = request.ThreatLevel
	}

	if !dara.IsNil(request.ThreatSourceIp) {
		query["ThreatSourceIp"] = request.ThreatSourceIp
	}

	if !dara.IsNil(request.ThreatType) {
		query["ThreatType"] = request.ThreatType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsThreatLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsThreatLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Public DNS threat logs.
//
// @param request - DescribePdnsThreatLogsRequest
//
// @return DescribePdnsThreatLogsResponse
func DescribePdnsThreatLogs(client *Client, request *DescribePdnsThreatLogsRequest) (_result *DescribePdnsThreatLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsThreatLogsResponse{}
	_body, _err := DescribePdnsThreatLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves Public DNS threat statistics.
//
// @param request - DescribePdnsThreatStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsThreatStatisticResponse
func DescribePdnsThreatStatisticWithOptions(client *Client, request *DescribePdnsThreatStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsThreatStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.ThreatSourceIp) {
		query["ThreatSourceIp"] = request.ThreatSourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsThreatStatistic"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsThreatStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves Public DNS threat statistics.
//
// @param request - DescribePdnsThreatStatisticRequest
//
// @return DescribePdnsThreatStatisticResponse
func DescribePdnsThreatStatistic(client *Client, request *DescribePdnsThreatStatisticRequest) (_result *DescribePdnsThreatStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsThreatStatisticResponse{}
	_body, _err := DescribePdnsThreatStatisticWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the public DNS threat statistics list.
//
// @param request - DescribePdnsThreatStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsThreatStatisticsResponse
func DescribePdnsThreatStatisticsWithOptions(client *Client, request *DescribePdnsThreatStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsThreatStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.ThreatLevel) {
		query["ThreatLevel"] = request.ThreatLevel
	}

	if !dara.IsNil(request.ThreatSourceIp) {
		query["ThreatSourceIp"] = request.ThreatSourceIp
	}

	if !dara.IsNil(request.ThreatType) {
		query["ThreatType"] = request.ThreatType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsThreatStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsThreatStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the public DNS threat statistics list.
//
// @param request - DescribePdnsThreatStatisticsRequest
//
// @return DescribePdnsThreatStatisticsResponse
func DescribePdnsThreatStatistics(client *Client, request *DescribePdnsThreatStatisticsRequest) (_result *DescribePdnsThreatStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsThreatStatisticsResponse{}
	_body, _err := DescribePdnsThreatStatisticsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of Public DNS UDP IP address ranges.
//
// @param request - DescribePdnsUdpIpSegmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsUdpIpSegmentsResponse
func DescribePdnsUdpIpSegmentsWithOptions(client *Client, request *DescribePdnsUdpIpSegmentsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsUdpIpSegmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsUdpIpSegments"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsUdpIpSegmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of Public DNS UDP IP address ranges.
//
// @param request - DescribePdnsUdpIpSegmentsRequest
//
// @return DescribePdnsUdpIpSegmentsResponse
func DescribePdnsUdpIpSegments(client *Client, request *DescribePdnsUdpIpSegmentsRequest) (_result *DescribePdnsUdpIpSegmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsUdpIpSegmentsResponse{}
	_body, _err := DescribePdnsUdpIpSegmentsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves user information for Public DNS.
//
// @param request - DescribePdnsUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsUserInfoResponse
func DescribePdnsUserInfoWithOptions(client *Client, request *DescribePdnsUserInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsUserInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves user information for Public DNS.
//
// @param request - DescribePdnsUserInfoRequest
//
// @return DescribePdnsUserInfoResponse
func DescribePdnsUserInfo(client *Client, request *DescribePdnsUserInfoRequest) (_result *DescribePdnsUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsUserInfoResponse{}
	_body, _err := DescribePdnsUserInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs for the DNS records of a domain name.
//
// @param request - DescribeRecordLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordLogsResponse
func DescribeRecordLogsWithOptions(client *Client, request *DescribeRecordLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs for the DNS records of a domain name.
//
// @param request - DescribeRecordLogsRequest
//
// @return DescribeRecordLogsResponse
func DescribeRecordLogs(client *Client, request *DescribeRecordLogsRequest) (_result *DescribeRecordLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordLogsResponse{}
	_body, _err := DescribeRecordLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordResolveStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordResolveStatisticsSummaryResponse
func DescribeRecordResolveStatisticsSummaryWithOptions(client *Client, request *DescribeRecordResolveStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordResolveStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordResolveStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordResolveStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordResolveStatisticsSummaryRequest
//
// @return DescribeRecordResolveStatisticsSummaryResponse
func DescribeRecordResolveStatisticsSummary(client *Client, request *DescribeRecordResolveStatisticsSummaryRequest) (_result *DescribeRecordResolveStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordResolveStatisticsSummaryResponse{}
	_body, _err := DescribeRecordResolveStatisticsSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries real-time data about DNS requests for a specified subdomain.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeRecordStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordStatisticsResponse
func DescribeRecordStatisticsWithOptions(client *Client, request *DescribeRecordStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries real-time data about DNS requests for a specified subdomain.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeRecordStatisticsRequest
//
// @return DescribeRecordStatisticsResponse
func DescribeRecordStatistics(client *Client, request *DescribeRecordStatisticsRequest) (_result *DescribeRecordStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordStatisticsResponse{}
	_body, _err := DescribeRecordStatisticsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordStatisticsSummaryResponse
func DescribeRecordStatisticsSummaryWithOptions(client *Client, request *DescribeRecordStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordStatisticsSummaryRequest
//
// @return DescribeRecordStatisticsSummaryResponse
func DescribeRecordStatisticsSummary(client *Client, request *DescribeRecordStatisticsSummaryRequest) (_result *DescribeRecordStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordStatisticsSummaryResponse{}
	_body, _err := DescribeRecordStatisticsSummaryWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DescribeRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecursionRecordResponse
func DescribeRecursionRecordWithOptions(client *Client, request *DescribeRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecursionRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DescribeRecursionRecordRequest
//
// @return DescribeRecursionRecordResponse
func DescribeRecursionRecord(client *Client, request *DescribeRecursionRecordRequest) (_result *DescribeRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecursionRecordResponse{}
	_body, _err := DescribeRecursionRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an authoritative zone for recursive resolution.
//
// Description:
//
// Real-time data is measured hourly.
//
// @param request - DescribeRecursionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecursionZoneResponse
func DescribeRecursionZoneWithOptions(client *Client, request *DescribeRecursionZoneRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecursionZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecursionZone"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecursionZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an authoritative zone for recursive resolution.
//
// Description:
//
// Real-time data is measured hourly.
//
// @param request - DescribeRecursionZoneRequest
//
// @return DescribeRecursionZoneResponse
func DescribeRecursionZone(client *Client, request *DescribeRecursionZoneRequest) (_result *DescribeRecursionZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecursionZoneResponse{}
	_body, _err := DescribeRecursionZoneWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all DNS records for a specific subdomain based on the specified parameters.
//
// @param request - DescribeSubDomainRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubDomainRecordsResponse
func DescribeSubDomainRecordsWithOptions(client *Client, request *DescribeSubDomainRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubDomainRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubDomainRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubDomainRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all DNS records for a specific subdomain based on the specified parameters.
//
// @param request - DescribeSubDomainRecordsRequest
//
// @return DescribeSubDomainRecordsResponse
func DescribeSubDomainRecords(client *Client, request *DescribeSubDomainRecordsRequest) (_result *DescribeSubDomainRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubDomainRecordsResponse{}
	_body, _err := DescribeSubDomainRecordsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all lines supported by Cloud DNS.
//
// @param request - DescribeSupportLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportLinesResponse
func DescribeSupportLinesWithOptions(client *Client, request *DescribeSupportLinesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupportLinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupportLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupportLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all lines supported by Cloud DNS.
//
// @param request - DescribeSupportLinesRequest
//
// @return DescribeSupportLinesResponse
func DescribeSupportLines(client *Client, request *DescribeSupportLinesRequest) (_result *DescribeSupportLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSupportLinesResponse{}
	_body, _err := DescribeSupportLinesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries existing tags.
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func DescribeTagsWithOptions(client *Client, request *DescribeTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries existing tags.
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
func DescribeTags(client *Client, request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := DescribeTagsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names transferred between accounts based on the specified input parameters.
//
// @param request - DescribeTransferDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransferDomainsResponse
func DescribeTransferDomainsWithOptions(client *Client, request *DescribeTransferDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransferDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FromUserId) {
		query["FromUserId"] = request.FromUserId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TransferType) {
		query["TransferType"] = request.TransferType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTransferDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTransferDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names transferred between accounts based on the specified input parameters.
//
// @param request - DescribeTransferDomainsRequest
//
// @return DescribeTransferDomainsResponse
func DescribeTransferDomains(client *Client, request *DescribeTransferDomainsRequest) (_result *DescribeTransferDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTransferDomainsResponse{}
	_body, _err := DescribeTransferDomainsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a disaster recovery plan.
//
// @param request - ExecuteGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteGtmRecoveryPlanResponse
func ExecuteGtmRecoveryPlanWithOptions(client *Client, request *ExecuteGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *ExecuteGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a disaster recovery plan.
//
// @param request - ExecuteGtmRecoveryPlanRequest
//
// @return ExecuteGtmRecoveryPlanResponse
func ExecuteGtmRecoveryPlan(client *Client, request *ExecuteGtmRecoveryPlanRequest) (_result *ExecuteGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteGtmRecoveryPlanResponse{}
	_body, _err := ExecuteGtmRecoveryPlanWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the primary domain name from a specified domain name.
//
// Description:
//
// # For more information about primary and subdomain name levels, see
//
// <props="china">[Domain name levels](https://help.aliyun.com/document_detail/39803.html?spm=a2c4g.2357293.0.0.211f41ffUR1cPb). For example, if you enter `www.abc.com`, the output is abc.com.
//
// <props="intl">[Domain name levels](https://www.alibabacloud.com/help/zh/faq-detail/39803.htm). For example, if you enter `www.abc.com`, the output is abc.com.
//
// @param request - GetMainDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMainDomainNameResponse
func GetMainDomainNameWithOptions(client *Client, request *GetMainDomainNameRequest, runtime *dara.RuntimeOptions) (_result *GetMainDomainNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputString) {
		query["InputString"] = request.InputString
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMainDomainName"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMainDomainNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the primary domain name from a specified domain name.
//
// Description:
//
// # For more information about primary and subdomain name levels, see
//
// <props="china">[Domain name levels](https://help.aliyun.com/document_detail/39803.html?spm=a2c4g.2357293.0.0.211f41ffUR1cPb). For example, if you enter `www.abc.com`, the output is abc.com.
//
// <props="intl">[Domain name levels](https://www.alibabacloud.com/help/zh/faq-detail/39803.htm). For example, if you enter `www.abc.com`, the output is abc.com.
//
// @param request - GetMainDomainNameRequest
//
// @return GetMainDomainNameResponse
func GetMainDomainName(client *Client, request *GetMainDomainNameRequest) (_result *GetMainDomainNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMainDomainNameResponse{}
	_body, _err := GetMainDomainNameWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a TXT record for domain and subdomain verification. This operation supports batch retrieval.
//
// @param request - GetTxtRecordForVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTxtRecordForVerifyResponse
func GetTxtRecordForVerifyWithOptions(client *Client, request *GetTxtRecordForVerifyRequest, runtime *dara.RuntimeOptions) (_result *GetTxtRecordForVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTxtRecordForVerify"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTxtRecordForVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a TXT record for domain and subdomain verification. This operation supports batch retrieval.
//
// @param request - GetTxtRecordForVerifyRequest
//
// @return GetTxtRecordForVerifyResponse
func GetTxtRecordForVerify(client *Client, request *GetTxtRecordForVerifyRequest) (_result *GetTxtRecordForVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTxtRecordForVerifyResponse{}
	_body, _err := GetTxtRecordForVerifyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Agent registration information.
//
// @param request - ListAtiAgentRegisterInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAtiAgentRegisterInfosResponse
func ListAtiAgentRegisterInfosWithOptions(client *Client, request *ListAtiAgentRegisterInfosRequest, runtime *dara.RuntimeOptions) (_result *ListAtiAgentRegisterInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentDisplayName) {
		query["AgentDisplayName"] = request.AgentDisplayName
	}

	if !dara.IsNil(request.AgentHost) {
		query["AgentHost"] = request.AgentHost
	}

	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		query["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAtiAgentRegisterInfos"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAtiAgentRegisterInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Agent registration information.
//
// @param request - ListAtiAgentRegisterInfosRequest
//
// @return ListAtiAgentRegisterInfosResponse
func ListAtiAgentRegisterInfos(client *Client, request *ListAtiAgentRegisterInfosRequest) (_result *ListAtiAgentRegisterInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAtiAgentRegisterInfosResponse{}
	_body, _err := ListAtiAgentRegisterInfosWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of real-name registered contacts.
//
// @param request - ListAtiChangeLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAtiChangeLogsResponse
func ListAtiChangeLogsWithOptions(client *Client, request *ListAtiChangeLogsRequest, runtime *dara.RuntimeOptions) (_result *ListAtiChangeLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OperatorAccount) {
		query["OperatorAccount"] = request.OperatorAccount
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAtiChangeLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAtiChangeLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of real-name registered contacts.
//
// @param request - ListAtiChangeLogsRequest
//
// @return ListAtiChangeLogsResponse
func ListAtiChangeLogs(client *Client, request *ListAtiChangeLogsRequest) (_result *ListAtiChangeLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAtiChangeLogsResponse{}
	_body, _err := ListAtiChangeLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实名注册人列表
//
// @param request - ListAtiRegistrantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAtiRegistrantsResponse
func ListAtiRegistrantsWithOptions(client *Client, request *ListAtiRegistrantsRequest, runtime *dara.RuntimeOptions) (_result *ListAtiRegistrantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAtiRegistrants"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAtiRegistrantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实名注册人列表
//
// @param request - ListAtiRegistrantsRequest
//
// @return ListAtiRegistrantsResponse
func ListAtiRegistrants(client *Client, request *ListAtiRegistrantsRequest) (_result *ListAtiRegistrantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAtiRegistrantsResponse{}
	_body, _err := ListAtiRegistrantsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of address pools in Global Traffic Manager (GTM) 3.0.
//
// @param request - ListCloudGtmAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAddressPoolsResponse
func ListCloudGtmAddressPoolsWithOptions(client *Client, request *ListCloudGtmAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAddressPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.AddressPoolType) {
		query["AddressPoolType"] = request.AddressPoolType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of address pools in Global Traffic Manager (GTM) 3.0.
//
// @param request - ListCloudGtmAddressPoolsRequest
//
// @return ListCloudGtmAddressPoolsResponse
func ListCloudGtmAddressPools(client *Client, request *ListCloudGtmAddressPoolsRequest) (_result *ListCloudGtmAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAddressPoolsResponse{}
	_body, _err := ListCloudGtmAddressPoolsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of addresses based on the specified parameters.
//
// @param request - ListCloudGtmAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAddressesResponse
func ListCloudGtmAddressesWithOptions(client *Client, request *ListCloudGtmAddressesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.MonitorTemplateId) {
		query["MonitorTemplateId"] = request.MonitorTemplateId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAddresses"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAddressesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of addresses based on the specified parameters.
//
// @param request - ListCloudGtmAddressesRequest
//
// @return ListCloudGtmAddressesResponse
func ListCloudGtmAddresses(client *Client, request *ListCloudGtmAddressesRequest) (_result *ListCloudGtmAddressesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAddressesResponse{}
	_body, _err := ListCloudGtmAddressesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of alert logs.
//
// @param request - ListCloudGtmAlertLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAlertLogsResponse
func ListCloudGtmAlertLogsWithOptions(client *Client, request *ListCloudGtmAlertLogsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAlertLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAlertLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAlertLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of alert logs.
//
// @param request - ListCloudGtmAlertLogsRequest
//
// @return ListCloudGtmAlertLogsResponse
func ListCloudGtmAlertLogs(client *Client, request *ListCloudGtmAlertLogsRequest) (_result *ListCloudGtmAlertLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAlertLogsResponse{}
	_body, _err := ListCloudGtmAlertLogsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListCloudGtmAvailableAlertGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAvailableAlertGroupsResponse
func ListCloudGtmAvailableAlertGroupsWithOptions(client *Client, request *ListCloudGtmAvailableAlertGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAvailableAlertGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAvailableAlertGroups"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAvailableAlertGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCloudGtmAvailableAlertGroupsRequest
//
// @return ListCloudGtmAvailableAlertGroupsResponse
func ListCloudGtmAvailableAlertGroups(client *Client, request *ListCloudGtmAvailableAlertGroupsRequest) (_result *ListCloudGtmAvailableAlertGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAvailableAlertGroupsResponse{}
	_body, _err := ListCloudGtmAvailableAlertGroupsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of instance configurations that contain access domain names and address pool information.
//
// @param request - ListCloudGtmInstanceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmInstanceConfigsResponse
func ListCloudGtmInstanceConfigsWithOptions(client *Client, request *ListCloudGtmInstanceConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmInstanceConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ScheduleDomainName) {
		query["ScheduleDomainName"] = request.ScheduleDomainName
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmInstanceConfigs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmInstanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of instance configurations that contain access domain names and address pool information.
//
// @param request - ListCloudGtmInstanceConfigsRequest
//
// @return ListCloudGtmInstanceConfigsResponse
func ListCloudGtmInstanceConfigs(client *Client, request *ListCloudGtmInstanceConfigsRequest) (_result *ListCloudGtmInstanceConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmInstanceConfigsResponse{}
	_body, _err := ListCloudGtmInstanceConfigsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of GTM 3.0 instances based on the specified parameters.
//
// @param request - ListCloudGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmInstancesResponse
func ListCloudGtmInstancesWithOptions(client *Client, request *ListCloudGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of GTM 3.0 instances based on the specified parameters.
//
// @param request - ListCloudGtmInstancesRequest
//
// @return ListCloudGtmInstancesResponse
func ListCloudGtmInstances(client *Client, request *ListCloudGtmInstancesRequest) (_result *ListCloudGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmInstancesResponse{}
	_body, _err := ListCloudGtmInstancesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the health check monitoring nodes based on the specified input parameters.
//
// @param request - ListCloudGtmMonitorNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmMonitorNodesResponse
func ListCloudGtmMonitorNodesWithOptions(client *Client, request *ListCloudGtmMonitorNodesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmMonitorNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ShowDisabledNodes) {
		query["ShowDisabledNodes"] = request.ShowDisabledNodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmMonitorNodes"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmMonitorNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the health check monitoring nodes based on the specified input parameters.
//
// @param request - ListCloudGtmMonitorNodesRequest
//
// @return ListCloudGtmMonitorNodesResponse
func ListCloudGtmMonitorNodes(client *Client, request *ListCloudGtmMonitorNodesRequest) (_result *ListCloudGtmMonitorNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmMonitorNodesResponse{}
	_body, _err := ListCloudGtmMonitorNodesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries health check templates based on the specified parameters.
//
// @param request - ListCloudGtmMonitorTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmMonitorTemplatesResponse
func ListCloudGtmMonitorTemplatesWithOptions(client *Client, request *ListCloudGtmMonitorTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmMonitorTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmMonitorTemplates"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmMonitorTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries health check templates based on the specified parameters.
//
// @param request - ListCloudGtmMonitorTemplatesRequest
//
// @return ListCloudGtmMonitorTemplatesResponse
func ListCloudGtmMonitorTemplates(client *Client, request *ListCloudGtmMonitorTemplatesRequest) (_result *ListCloudGtmMonitorTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmMonitorTemplatesResponse{}
	_body, _err := ListCloudGtmMonitorTemplatesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the built-in authoritative records for recursive resolution.
//
// @param request - ListRecursionRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecursionRecordsResponse
func ListRecursionRecordsWithOptions(client *Client, request *ListRecursionRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListRecursionRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecursionRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecursionRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the built-in authoritative records for recursive resolution.
//
// @param request - ListRecursionRecordsRequest
//
// @return ListRecursionRecordsResponse
func ListRecursionRecords(client *Client, request *ListRecursionRecordsRequest) (_result *ListRecursionRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecursionRecordsResponse{}
	_body, _err := ListRecursionRecordsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the zones that are used for recursive resolution of built-in authoritative domain names.
//
// Description:
//
// - Specify at least ResourceId.N or Tag.N (Tag.N.Key and Tag.N.Value) in your request to identify the resources to retrieve.
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with the tag key are returned. An error is returned if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, the resources that match all of them are returned.
//
// @param request - ListRecursionZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecursionZonesResponse
func ListRecursionZonesWithOptions(client *Client, request *ListRecursionZonesRequest, runtime *dara.RuntimeOptions) (_result *ListRecursionZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecursionZones"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecursionZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones that are used for recursive resolution of built-in authoritative domain names.
//
// Description:
//
// - Specify at least ResourceId.N or Tag.N (Tag.N.Key and Tag.N.Value) in your request to identify the resources to retrieve.
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with the tag key are returned. An error is returned if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, the resources that match all of them are returned.
//
// @param request - ListRecursionZonesRequest
//
// @return ListRecursionZonesResponse
func ListRecursionZones(client *Client, request *ListRecursionZonesRequest) (_result *ListRecursionZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecursionZonesResponse{}
	_body, _err := ListRecursionZonesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resources by tag.
//
// Description:
//
// - To specify the objects to retrieve, your request must contain at least one of the following parameters: `ResourceId.N` or `Tag.N` (which consists of Tag.N.Key and Tag.N.Value).
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with that tag key are returned. An error occurs if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified tag key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, only the resources that have all the specified key-value pairs are returned.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func ListTagResourcesWithOptions(client *Client, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resources by tag.
//
// Description:
//
// - To specify the objects to retrieve, your request must contain at least one of the following parameters: `ResourceId.N` or `Tag.N` (which consists of Tag.N.Key and Tag.N.Value).
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with that tag key are returned. An error occurs if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified tag key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, only the resources that have all the specified key-value pairs are returned.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func ListTagResources(client *Client, request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := ListTagResourcesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the DNS servers for a domain name from a third-party provider to Alibaba Cloud DNS.
//
// Description:
//
// After the operation is successful, the DNS servers are changed to Alibaba Cloud DNS servers. The names of these new servers end with hichina.com.
//
// > **Prerequisite: This operation applies to domain names that are registered with Alibaba Cloud and currently use third-party DNS servers.**
//
// @param request - ModifyHichinaDomainDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHichinaDomainDNSResponse
func ModifyHichinaDomainDNSWithOptions(client *Client, request *ModifyHichinaDomainDNSRequest, runtime *dara.RuntimeOptions) (_result *ModifyHichinaDomainDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHichinaDomainDNS"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHichinaDomainDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the DNS servers for a domain name from a third-party provider to Alibaba Cloud DNS.
//
// Description:
//
// After the operation is successful, the DNS servers are changed to Alibaba Cloud DNS servers. The names of these new servers end with hichina.com.
//
// > **Prerequisite: This operation applies to domain names that are registered with Alibaba Cloud and currently use third-party DNS servers.**
//
// @param request - ModifyHichinaDomainDNSRequest
//
// @return ModifyHichinaDomainDNSResponse
func ModifyHichinaDomainDNS(client *Client, request *ModifyHichinaDomainDNSRequest) (_result *ModifyHichinaDomainDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHichinaDomainDNSResponse{}
	_body, _err := ModifyHichinaDomainDNSWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a domain name to a different resource group.
//
// @param request - MoveDomainResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveDomainResourceGroupResponse
func MoveDomainResourceGroupWithOptions(client *Client, request *MoveDomainResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveDomainResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveDomainResourceGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveDomainResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a domain name to a different resource group.
//
// @param request - MoveDomainResourceGroupRequest
//
// @return MoveDomainResourceGroupResponse
func MoveDomainResourceGroup(client *Client, request *MoveDomainResourceGroupRequest) (_result *MoveDomainResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveDomainResourceGroupResponse{}
	_body, _err := MoveDomainResourceGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a Global Traffic Manager (GTM) instance to a new resource group.
//
// @param request - MoveGtmResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveGtmResourceGroupResponse
func MoveGtmResourceGroupWithOptions(client *Client, request *MoveGtmResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveGtmResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveGtmResourceGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveGtmResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a Global Traffic Manager (GTM) instance to a new resource group.
//
// @param request - MoveGtmResourceGroupRequest
//
// @return MoveGtmResourceGroupResponse
func MoveGtmResourceGroup(client *Client, request *MoveGtmResourceGroupRequest) (_result *MoveGtmResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveGtmResourceGroupResponse{}
	_body, _err := MoveGtmResourceGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or deletes domain names and DNS records in batches.
//
// Description:
//
// Use this operation for batch DNS tasks that do not require immediate execution.
//
// @param request - OperateBatchDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateBatchDomainResponse
func OperateBatchDomainWithOptions(client *Client, request *OperateBatchDomainRequest, runtime *dara.RuntimeOptions) (_result *OperateBatchDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainRecordInfo) {
		query["DomainRecordInfo"] = request.DomainRecordInfo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateBatchDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateBatchDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or deletes domain names and DNS records in batches.
//
// Description:
//
// Use this operation for batch DNS tasks that do not require immediate execution.
//
// @param request - OperateBatchDomainRequest
//
// @return OperateBatchDomainResponse
func OperateBatchDomain(client *Client, request *OperateBatchDomainRequest) (_result *OperateBatchDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateBatchDomainResponse{}
	_body, _err := OperateBatchDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Pause Public DNS Service
//
// @param request - PausePdnsServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePdnsServiceResponse
func PausePdnsServiceWithOptions(client *Client, request *PausePdnsServiceRequest, runtime *dara.RuntimeOptions) (_result *PausePdnsServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PausePdnsService"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PausePdnsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Pause Public DNS Service
//
// @param request - PausePdnsServiceRequest
//
// @return PausePdnsServiceResponse
func PausePdnsService(client *Client, request *PausePdnsServiceRequest) (_result *PausePdnsServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PausePdnsServiceResponse{}
	_body, _err := PausePdnsServiceWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Previews a disaster recovery plan.
//
// @param request - PreviewGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewGtmRecoveryPlanResponse
func PreviewGtmRecoveryPlanWithOptions(client *Client, request *PreviewGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *PreviewGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Previews a disaster recovery plan.
//
// @param request - PreviewGtmRecoveryPlanRequest
//
// @return PreviewGtmRecoveryPlanResponse
func PreviewGtmRecoveryPlan(client *Client, request *PreviewGtmRecoveryPlanRequest) (_result *PreviewGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreviewGtmRecoveryPlanResponse{}
	_body, _err := PreviewGtmRecoveryPlanWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Public DNS AppKey
//
// @param request - RemovePdnsAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePdnsAppKeyResponse
func RemovePdnsAppKeyWithOptions(client *Client, request *RemovePdnsAppKeyRequest, runtime *dara.RuntimeOptions) (_result *RemovePdnsAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKeyId) {
		query["AppKeyId"] = request.AppKeyId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePdnsAppKey"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePdnsAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Public DNS AppKey
//
// @param request - RemovePdnsAppKeyRequest
//
// @return RemovePdnsAppKeyResponse
func RemovePdnsAppKey(client *Client, request *RemovePdnsAppKeyRequest) (_result *RemovePdnsAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePdnsAppKeyResponse{}
	_body, _err := RemovePdnsAppKeyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Remove Public DNS UDP IP Segment
//
// @param request - RemovePdnsUdpIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePdnsUdpIpSegmentResponse
func RemovePdnsUdpIpSegmentWithOptions(client *Client, request *RemovePdnsUdpIpSegmentRequest, runtime *dara.RuntimeOptions) (_result *RemovePdnsUdpIpSegmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePdnsUdpIpSegment"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePdnsUdpIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Remove Public DNS UDP IP Segment
//
// @param request - RemovePdnsUdpIpSegmentRequest
//
// @return RemovePdnsUdpIpSegmentResponse
func RemovePdnsUdpIpSegment(client *Client, request *RemovePdnsUdpIpSegmentRequest) (_result *RemovePdnsUdpIpSegmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePdnsUdpIpSegmentResponse{}
	_body, _err := RemovePdnsUdpIpSegmentWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the serverHold status of a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status of a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRspDomainServerHoldStatusForGatewayResponse
func RemoveRspDomainServerHoldStatusForGatewayWithOptions(client *Client, request *RemoveRspDomainServerHoldStatusForGatewayRequest, runtime *dara.RuntimeOptions) (_result *RemoveRspDomainServerHoldStatusForGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveRspDomainServerHoldStatusForGateway"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the serverHold status of a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status of a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayRequest
//
// @return RemoveRspDomainServerHoldStatusForGatewayResponse
func RemoveRspDomainServerHoldStatusForGateway(client *Client, request *RemoveRspDomainServerHoldStatusForGatewayRequest) (_result *RemoveRspDomainServerHoldStatusForGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := RemoveRspDomainServerHoldStatusForGatewayWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the serverHold status for a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status for a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayOteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRspDomainServerHoldStatusForGatewayOteResponse
func RemoveRspDomainServerHoldStatusForGatewayOteWithOptions(client *Client, request *RemoveRspDomainServerHoldStatusForGatewayOteRequest, runtime *dara.RuntimeOptions) (_result *RemoveRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveRspDomainServerHoldStatusForGatewayOte"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the serverHold status for a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status for a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayOteRequest
//
// @return RemoveRspDomainServerHoldStatusForGatewayOteResponse
func RemoveRspDomainServerHoldStatusForGatewayOte(client *Client, request *RemoveRspDomainServerHoldStatusForGatewayOteRequest) (_result *RemoveRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := RemoveRspDomainServerHoldStatusForGatewayOteWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces the addresses in a specified address pool.
//
// @param tmpReq - ReplaceCloudGtmAddressPoolAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceCloudGtmAddressPoolAddressResponse
func ReplaceCloudGtmAddressPoolAddressWithOptions(client *Client, tmpReq *ReplaceCloudGtmAddressPoolAddressRequest, runtime *dara.RuntimeOptions) (_result *ReplaceCloudGtmAddressPoolAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplaceCloudGtmAddressPoolAddressShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Addresses) {
		request.AddressesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addresses, dara.String("Addresses"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.AddressesShrink) {
		query["Addresses"] = request.AddressesShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceCloudGtmAddressPoolAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceCloudGtmAddressPoolAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the addresses in a specified address pool.
//
// @param request - ReplaceCloudGtmAddressPoolAddressRequest
//
// @return ReplaceCloudGtmAddressPoolAddressResponse
func ReplaceCloudGtmAddressPoolAddress(client *Client, request *ReplaceCloudGtmAddressPoolAddressRequest) (_result *ReplaceCloudGtmAddressPoolAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplaceCloudGtmAddressPoolAddressResponse{}
	_body, _err := ReplaceCloudGtmAddressPoolAddressWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces the address pools associated with an instance.
//
// @param tmpReq - ReplaceCloudGtmInstanceConfigAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceCloudGtmInstanceConfigAddressPoolResponse
func ReplaceCloudGtmInstanceConfigAddressPoolWithOptions(client *Client, tmpReq *ReplaceCloudGtmInstanceConfigAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *ReplaceCloudGtmInstanceConfigAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddressPools) {
		request.AddressPoolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddressPools, dara.String("AddressPools"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolsShrink) {
		query["AddressPools"] = request.AddressPoolsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceCloudGtmInstanceConfigAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceCloudGtmInstanceConfigAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the address pools associated with an instance.
//
// @param request - ReplaceCloudGtmInstanceConfigAddressPoolRequest
//
// @return ReplaceCloudGtmInstanceConfigAddressPoolResponse
func ReplaceCloudGtmInstanceConfigAddressPool(client *Client, request *ReplaceCloudGtmInstanceConfigAddressPoolRequest) (_result *ReplaceCloudGtmInstanceConfigAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplaceCloudGtmInstanceConfigAddressPoolResponse{}
	_body, _err := ReplaceCloudGtmInstanceConfigAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Resume Public DNS Service
//
// @param request - ResumePdnsServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePdnsServiceResponse
func ResumePdnsServiceWithOptions(client *Client, request *ResumePdnsServiceRequest, runtime *dara.RuntimeOptions) (_result *ResumePdnsServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumePdnsService"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumePdnsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Resume Public DNS Service
//
// @param request - ResumePdnsServiceRequest
//
// @return ResumePdnsServiceResponse
func ResumePdnsService(client *Client, request *ResumePdnsServiceRequest) (_result *ResumePdnsServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumePdnsServiceResponse{}
	_body, _err := ResumePdnsServiceWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a domain name.
//
// Description:
//
// Before you can retrieve a domain name, you must verify it using a TXT record. Use this operation in conjunction with the <props="china">[Generate a TXT record](https://help.aliyun.com/document_detail/145533.html) <props="intl">[Generate a TXT record](https://www.alibabacloud.com/help/zh/alibaba-cloud-dns/latest/generating-a-txt-record) operation.
//
// @param request - RetrieveDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveDomainResponse
func RetrieveDomainWithOptions(client *Client, request *RetrieveDomainRequest, runtime *dara.RuntimeOptions) (_result *RetrieveDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetrieveDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a domain name.
//
// Description:
//
// Before you can retrieve a domain name, you must verify it using a TXT record. Use this operation in conjunction with the <props="china">[Generate a TXT record](https://help.aliyun.com/document_detail/145533.html) <props="intl">[Generate a TXT record](https://www.alibabacloud.com/help/zh/alibaba-cloud-dns/latest/generating-a-txt-record) operation.
//
// @param request - RetrieveDomainRequest
//
// @return RetrieveDomainResponse
func RetrieveDomain(client *Client, request *RetrieveDomainRequest) (_result *RetrieveDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetrieveDomainResponse{}
	_body, _err := RetrieveDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the registration information of an Agent.
//
// @param request - RevokeAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAtiAgentRegisterInfoResponse
func RevokeAtiAgentRegisterInfoWithOptions(client *Client, request *RevokeAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *RevokeAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.ReasonCode) {
		query["ReasonCode"] = request.ReasonCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the registration information of an Agent.
//
// @param request - RevokeAtiAgentRegisterInfoRequest
//
// @return RevokeAtiAgentRegisterInfoResponse
func RevokeAtiAgentRegisterInfo(client *Client, request *RevokeAtiAgentRegisterInfoRequest) (_result *RevokeAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeAtiAgentRegisterInfoResponse{}
	_body, _err := RevokeAtiAgentRegisterInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a real-name registrant.
//
// @param request - RevokeAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAtiRegistrantResponse
func RevokeAtiRegistrantWithOptions(client *Client, request *RevokeAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *RevokeAtiRegistrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a real-name registrant.
//
// @param request - RevokeAtiRegistrantRequest
//
// @return RevokeAtiRegistrantResponse
func RevokeAtiRegistrant(client *Client, request *RevokeAtiRegistrantRequest) (_result *RevokeAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeAtiRegistrantResponse{}
	_body, _err := RevokeAtiRegistrantWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back a disaster recovery plan.
//
// @param request - RollbackGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackGtmRecoveryPlanResponse
func RollbackGtmRecoveryPlanWithOptions(client *Client, request *RollbackGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *RollbackGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a disaster recovery plan.
//
// @param request - RollbackGtmRecoveryPlanRequest
//
// @return RollbackGtmRecoveryPlanResponse
func RollbackGtmRecoveryPlan(client *Client, request *RollbackGtmRecoveryPlanRequest) (_result *RollbackGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackGtmRecoveryPlanResponse{}
	_body, _err := RollbackGtmRecoveryPlanWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for agents in the Agent Marketplace.
//
// @param request - SearchAtiAgentRegisterInfoMarketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAtiAgentRegisterInfoMarketResponse
func SearchAtiAgentRegisterInfoMarketWithOptions(client *Client, request *SearchAtiAgentRegisterInfoMarketRequest, runtime *dara.RuntimeOptions) (_result *SearchAtiAgentRegisterInfoMarketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TrustLevel) {
		query["TrustLevel"] = request.TrustLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchAtiAgentRegisterInfoMarket"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAtiAgentRegisterInfoMarketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for agents in the Agent Marketplace.
//
// @param request - SearchAtiAgentRegisterInfoMarketRequest
//
// @return SearchAtiAgentRegisterInfoMarketResponse
func SearchAtiAgentRegisterInfoMarket(client *Client, request *SearchAtiAgentRegisterInfoMarketRequest) (_result *SearchAtiAgentRegisterInfoMarketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchAtiAgentRegisterInfoMarketResponse{}
	_body, _err := SearchAtiAgentRegisterInfoMarketWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for address pools by name, remarks, or other criteria.
//
// @param request - SearchCloudGtmAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmAddressPoolsResponse
func SearchCloudGtmAddressPoolsWithOptions(client *Client, request *SearchCloudGtmAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmAddressPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.AddressPoolType) {
		query["AddressPoolType"] = request.AddressPoolType
	}

	if !dara.IsNil(request.AvailableStatus) {
		query["AvailableStatus"] = request.AvailableStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for address pools by name, remarks, or other criteria.
//
// @param request - SearchCloudGtmAddressPoolsRequest
//
// @return SearchCloudGtmAddressPoolsResponse
func SearchCloudGtmAddressPools(client *Client, request *SearchCloudGtmAddressPoolsRequest) (_result *SearchCloudGtmAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmAddressPoolsResponse{}
	_body, _err := SearchCloudGtmAddressPoolsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for addresses based on criteria such as address name, remarks, referenced health check template, or address ID.
//
// @param request - SearchCloudGtmAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmAddressesResponse
func SearchCloudGtmAddressesWithOptions(client *Client, request *SearchCloudGtmAddressesRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.AvailableStatus) {
		query["AvailableStatus"] = request.AvailableStatus
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.MonitorTemplateName) {
		query["MonitorTemplateName"] = request.MonitorTemplateName
	}

	if !dara.IsNil(request.NameSearchCondition) {
		query["NameSearchCondition"] = request.NameSearchCondition
	}

	if !dara.IsNil(request.Names) {
		query["Names"] = request.Names
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RemarkSearchCondition) {
		query["RemarkSearchCondition"] = request.RemarkSearchCondition
	}

	if !dara.IsNil(request.Remarks) {
		query["Remarks"] = request.Remarks
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmAddresses"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmAddressesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for addresses based on criteria such as address name, remarks, referenced health check template, or address ID.
//
// @param request - SearchCloudGtmAddressesRequest
//
// @return SearchCloudGtmAddressesResponse
func SearchCloudGtmAddresses(client *Client, request *SearchCloudGtmAddressesRequest) (_result *SearchCloudGtmAddressesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmAddressesResponse{}
	_body, _err := SearchCloudGtmAddressesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves instance configurations that match the specified parameters.
//
// @param request - SearchCloudGtmInstanceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmInstanceConfigsResponse
func SearchCloudGtmInstanceConfigsWithOptions(client *Client, request *SearchCloudGtmInstanceConfigsRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmInstanceConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AvailableStatus) {
		query["AvailableStatus"] = request.AvailableStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ScheduleDomainName) {
		query["ScheduleDomainName"] = request.ScheduleDomainName
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmInstanceConfigs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmInstanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves instance configurations that match the specified parameters.
//
// @param request - SearchCloudGtmInstanceConfigsRequest
//
// @return SearchCloudGtmInstanceConfigsResponse
func SearchCloudGtmInstanceConfigs(client *Client, request *SearchCloudGtmInstanceConfigsRequest) (_result *SearchCloudGtmInstanceConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmInstanceConfigsResponse{}
	_body, _err := SearchCloudGtmInstanceConfigsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation searches for instances based on specified parameters.
//
// @param request - SearchCloudGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmInstancesResponse
func SearchCloudGtmInstancesWithOptions(client *Client, request *SearchCloudGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation searches for instances based on specified parameters.
//
// @param request - SearchCloudGtmInstancesRequest
//
// @return SearchCloudGtmInstancesResponse
func SearchCloudGtmInstances(client *Client, request *SearchCloudGtmInstancesRequest) (_result *SearchCloudGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmInstancesResponse{}
	_body, _err := SearchCloudGtmInstancesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for health check templates.
//
// @param request - SearchCloudGtmMonitorTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmMonitorTemplatesResponse
func SearchCloudGtmMonitorTemplatesWithOptions(client *Client, request *SearchCloudGtmMonitorTemplatesRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmMonitorTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmMonitorTemplates"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmMonitorTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for health check templates.
//
// @param request - SearchCloudGtmMonitorTemplatesRequest
//
// @return SearchCloudGtmMonitorTemplatesResponse
func SearchCloudGtmMonitorTemplates(client *Client, request *SearchCloudGtmMonitorTemplatesRequest) (_result *SearchCloudGtmMonitorTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmMonitorTemplatesResponse{}
	_body, _err := SearchCloudGtmMonitorTemplatesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for built-in authoritative DNS records used for recursive resolution.
//
// @param request - SearchRecursionRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchRecursionRecordsResponse
func SearchRecursionRecordsWithOptions(client *Client, request *SearchRecursionRecordsRequest, runtime *dara.RuntimeOptions) (_result *SearchRecursionRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchRecursionRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchRecursionRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for built-in authoritative DNS records used for recursive resolution.
//
// @param request - SearchRecursionRecordsRequest
//
// @return SearchRecursionRecordsResponse
func SearchRecursionRecords(client *Client, request *SearchRecursionRecordsRequest) (_result *SearchRecursionRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchRecursionRecordsResponse{}
	_body, _err := SearchRecursionRecordsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for zones of built-in authoritative domain names used for recursive resolution.
//
// Description:
//
// - To retrieve resources, you must specify at least `ResourceId.N` or `Tag.N` (`Tag.N.Key` and `Tag.N.Value`) in the request.
//
// - `Tag.N` is a resource tag that consists of a key-value pair. If you specify only `Tag.N.Key`, all tag values associated with that key are returned. If you specify only `Tag.N.Value`, an error is returned.
//
// - If you specify both `Tag.N` and `ResourceId.N`, the operation returns only the resources that are identified by `ResourceId.N` and match all the specified tag key-value pairs.
//
// - If you specify multiple tag key-value pairs, only resources that match all of them are returned.
//
// @param tmpReq - SearchRecursionZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchRecursionZonesResponse
func SearchRecursionZonesWithOptions(client *Client, tmpReq *SearchRecursionZonesRequest, runtime *dara.RuntimeOptions) (_result *SearchRecursionZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchRecursionZonesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EffectiveScopes) {
		request.EffectiveScopesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EffectiveScopes, dara.String("EffectiveScopes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EffectiveScopesShrink) {
		query["EffectiveScopes"] = request.EffectiveScopesShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchRecursionZones"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchRecursionZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for zones of built-in authoritative domain names used for recursive resolution.
//
// Description:
//
// - To retrieve resources, you must specify at least `ResourceId.N` or `Tag.N` (`Tag.N.Key` and `Tag.N.Value`) in the request.
//
// - `Tag.N` is a resource tag that consists of a key-value pair. If you specify only `Tag.N.Key`, all tag values associated with that key are returned. If you specify only `Tag.N.Value`, an error is returned.
//
// - If you specify both `Tag.N` and `ResourceId.N`, the operation returns only the resources that are identified by `ResourceId.N` and match all the specified tag key-value pairs.
//
// - If you specify multiple tag key-value pairs, only resources that match all of them are returned.
//
// @param request - SearchRecursionZonesRequest
//
// @return SearchRecursionZonesResponse
func SearchRecursionZones(client *Client, request *SearchRecursionZonesRequest) (_result *SearchRecursionZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchRecursionZonesResponse{}
	_body, _err := SearchRecursionZonesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the network traffic analysis feature for a Global Traffic Manager (GTM) instance. After this feature is enabled, you can view resolution logs and statistical reports for the domain name. You can also use the intelligent alerting feature based on abnormal metrics, such as resolution success rate and sudden changes in queries per second (QPS). This improves the observability and operations and maintenance (O&M) efficiency of the GTM instance.
//
// @param request - SetCloudGtmInstanceConfigLogSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCloudGtmInstanceConfigLogSwitchResponse
func SetCloudGtmInstanceConfigLogSwitchWithOptions(client *Client, request *SetCloudGtmInstanceConfigLogSwitchRequest, runtime *dara.RuntimeOptions) (_result *SetCloudGtmInstanceConfigLogSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCloudGtmInstanceConfigLogSwitch"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCloudGtmInstanceConfigLogSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the network traffic analysis feature for a Global Traffic Manager (GTM) instance. After this feature is enabled, you can view resolution logs and statistical reports for the domain name. You can also use the intelligent alerting feature based on abnormal metrics, such as resolution success rate and sudden changes in queries per second (QPS). This improves the observability and operations and maintenance (O&M) efficiency of the GTM instance.
//
// @param request - SetCloudGtmInstanceConfigLogSwitchRequest
//
// @return SetCloudGtmInstanceConfigLogSwitchResponse
func SetCloudGtmInstanceConfigLogSwitch(client *Client, request *SetCloudGtmInstanceConfigLogSwitchRequest) (_result *SetCloudGtmInstanceConfigLogSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCloudGtmInstanceConfigLogSwitchResponse{}
	_body, _err := SetCloudGtmInstanceConfigLogSwitchWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the weight configuration.
//
// @param request - SetDNSSLBStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDNSSLBStatusResponse
func SetDNSSLBStatusWithOptions(client *Client, request *SetDNSSLBStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDNSSLBStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Open) {
		query["Open"] = request.Open
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDNSSLBStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDNSSLBStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the weight configuration.
//
// @param request - SetDNSSLBStatusRequest
//
// @return SetDNSSLBStatusResponse
func SetDNSSLBStatus(client *Client, request *SetDNSSLBStatusRequest) (_result *SetDNSSLBStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDNSSLBStatusResponse{}
	_body, _err := SetDNSSLBStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the access mode for an access policy.
//
// Description:
//
// **
//
// @param request - SetDnsGtmAccessModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDnsGtmAccessModeResponse
func SetDnsGtmAccessModeWithOptions(client *Client, request *SetDnsGtmAccessModeRequest, runtime *dara.RuntimeOptions) (_result *SetDnsGtmAccessModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDnsGtmAccessMode"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDnsGtmAccessModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the access mode for an access policy.
//
// Description:
//
// **
//
// @param request - SetDnsGtmAccessModeRequest
//
// @return SetDnsGtmAccessModeResponse
func SetDnsGtmAccessMode(client *Client, request *SetDnsGtmAccessModeRequest) (_result *SetDnsGtmAccessModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDnsGtmAccessModeResponse{}
	_body, _err := SetDnsGtmAccessModeWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the health check status for an address pool.
//
// @param request - SetDnsGtmMonitorStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDnsGtmMonitorStatusResponse
func SetDnsGtmMonitorStatusWithOptions(client *Client, request *SetDnsGtmMonitorStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDnsGtmMonitorStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDnsGtmMonitorStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDnsGtmMonitorStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the health check status for an address pool.
//
// @param request - SetDnsGtmMonitorStatusRequest
//
// @return SetDnsGtmMonitorStatusResponse
func SetDnsGtmMonitorStatus(client *Client, request *SetDnsGtmMonitorStatusRequest) (_result *SetDnsGtmMonitorStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDnsGtmMonitorStatusResponse{}
	_body, _err := SetDnsGtmMonitorStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables Domain Name System Security Extensions (DNSSEC) for a domain name. This feature is available only to users of paid Alibaba Cloud DNS.
//
// @param request - SetDomainDnssecStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainDnssecStatusResponse
func SetDomainDnssecStatusWithOptions(client *Client, request *SetDomainDnssecStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDomainDnssecStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainDnssecStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainDnssecStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables Domain Name System Security Extensions (DNSSEC) for a domain name. This feature is available only to users of paid Alibaba Cloud DNS.
//
// @param request - SetDomainDnssecStatusRequest
//
// @return SetDomainDnssecStatusResponse
func SetDomainDnssecStatus(client *Client, request *SetDomainDnssecStatusRequest) (_result *SetDomainDnssecStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainDnssecStatusResponse{}
	_body, _err := SetDomainDnssecStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the status of a DNS record.
//
// @param request - SetDomainRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainRecordStatusResponse
func SetDomainRecordStatusWithOptions(client *Client, request *SetDomainRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDomainRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainRecordStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the status of a DNS record.
//
// @param request - SetDomainRecordStatusRequest
//
// @return SetDomainRecordStatusResponse
func SetDomainRecordStatus(client *Client, request *SetDomainRecordStatusRequest) (_result *SetDomainRecordStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainRecordStatusResponse{}
	_body, _err := SetDomainRecordStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the address pool switchover policy based on the request parameters.
//
// @param request - SetGtmAccessModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetGtmAccessModeResponse
func SetGtmAccessModeWithOptions(client *Client, request *SetGtmAccessModeRequest, runtime *dara.RuntimeOptions) (_result *SetGtmAccessModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetGtmAccessMode"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetGtmAccessModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the address pool switchover policy based on the request parameters.
//
// @param request - SetGtmAccessModeRequest
//
// @return SetGtmAccessModeResponse
func SetGtmAccessMode(client *Client, request *SetGtmAccessModeRequest) (_result *SetGtmAccessModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetGtmAccessModeResponse{}
	_body, _err := SetGtmAccessModeWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the health check status of an address pool.
//
// @param request - SetGtmMonitorStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetGtmMonitorStatusResponse
func SetGtmMonitorStatusWithOptions(client *Client, request *SetGtmMonitorStatusRequest, runtime *dara.RuntimeOptions) (_result *SetGtmMonitorStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetGtmMonitorStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetGtmMonitorStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the health check status of an address pool.
//
// @param request - SetGtmMonitorStatusRequest
//
// @return SetGtmMonitorStatusResponse
func SetGtmMonitorStatus(client *Client, request *SetGtmMonitorStatusRequest) (_result *SetGtmMonitorStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetGtmMonitorStatusResponse{}
	_body, _err := SetGtmMonitorStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits Agent registration information.
//
// @param request - SubmitAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAtiAgentRegisterInfoResponse
func SubmitAtiAgentRegisterInfoWithOptions(client *Client, request *SubmitAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *SubmitAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IdentityCsr) {
		query["IdentityCsr"] = request.IdentityCsr
	}

	if !dara.IsNil(request.ServerCertPem) {
		query["ServerCertPem"] = request.ServerCertPem
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits Agent registration information.
//
// @param request - SubmitAtiAgentRegisterInfoRequest
//
// @return SubmitAtiAgentRegisterInfoResponse
func SubmitAtiAgentRegisterInfo(client *Client, request *SubmitAtiAgentRegisterInfoRequest) (_result *SubmitAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAtiAgentRegisterInfoResponse{}
	_body, _err := SubmitAtiAgentRegisterInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a cache refresh task.
//
// @param request - SubmitIspFlushCacheTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIspFlushCacheTaskResponse
func SubmitIspFlushCacheTaskWithOptions(client *Client, request *SubmitIspFlushCacheTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitIspFlushCacheTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIspFlushCacheTask"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIspFlushCacheTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a cache refresh task.
//
// @param request - SubmitIspFlushCacheTaskRequest
//
// @return SubmitIspFlushCacheTaskResponse
func SubmitIspFlushCacheTask(client *Client, request *SubmitIspFlushCacheTaskRequest) (_result *SubmitIspFlushCacheTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitIspFlushCacheTaskResponse{}
	_body, _err := SubmitIspFlushCacheTaskWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches the access strategy mode of an instance.
//
// @param request - SwitchDnsGtmInstanceStrategyModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDnsGtmInstanceStrategyModeResponse
func SwitchDnsGtmInstanceStrategyModeWithOptions(client *Client, request *SwitchDnsGtmInstanceStrategyModeRequest, runtime *dara.RuntimeOptions) (_result *SwitchDnsGtmInstanceStrategyModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchDnsGtmInstanceStrategyMode"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchDnsGtmInstanceStrategyModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the access strategy mode of an instance.
//
// @param request - SwitchDnsGtmInstanceStrategyModeRequest
//
// @return SwitchDnsGtmInstanceStrategyModeResponse
func SwitchDnsGtmInstanceStrategyMode(client *Client, request *SwitchDnsGtmInstanceStrategyModeRequest) (_result *SwitchDnsGtmInstanceStrategyModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchDnsGtmInstanceStrategyModeResponse{}
	_body, _err := SwitchDnsGtmInstanceStrategyModeWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies tags for resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func TagResourcesWithOptions(client *Client, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies tags for resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func TagResources(client *Client, request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := TagResourcesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch transfers DNS permissions for multiple domain names to a specified execution account.
//
// @param request - TransferDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferDomainResponse
func TransferDomainWithOptions(client *Client, request *TransferDomainRequest, runtime *dara.RuntimeOptions) (_result *TransferDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch transfers DNS permissions for multiple domain names to a specified execution account.
//
// @param request - TransferDomainRequest
//
// @return TransferDomainResponse
func TransferDomain(client *Client, request *TransferDomainRequest) (_result *TransferDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferDomainResponse{}
	_body, _err := TransferDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches domain names from a paid Alibaba Cloud DNS instance.
//
// Description:
//
// An instance with an ID that starts with \\`dns-\\` is a new version instance. New version instances support attaching multiple domain names. You can call an API operation to attach domain names directly to an instance. An error is returned if the number of domain names exceeds the instance limit.
//
// An instance with an ID that does not start with \\`dns-\\` is a legacy instance. Legacy instances support only one domain name. Therefore, if you call this operation for an instance that already has a domain name attached, the domain name is replaced.
//
// @param request - UnbindInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindInstanceDomainsResponse
func UnbindInstanceDomainsWithOptions(client *Client, request *UnbindInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *UnbindInstanceDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindInstanceDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindInstanceDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches domain names from a paid Alibaba Cloud DNS instance.
//
// Description:
//
// An instance with an ID that starts with \\`dns-\\` is a new version instance. New version instances support attaching multiple domain names. You can call an API operation to attach domain names directly to an instance. An error is returned if the number of domain names exceeds the instance limit.
//
// An instance with an ID that does not start with \\`dns-\\` is a legacy instance. Legacy instances support only one domain name. Therefore, if you call this operation for an instance that already has a domain name attached, the domain name is replaced.
//
// @param request - UnbindInstanceDomainsRequest
//
// @return UnbindInstanceDomainsResponse
func UnbindInstanceDomains(client *Client, request *UnbindInstanceDomainsRequest) (_result *UnbindInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindInstanceDomainsResponse{}
	_body, _err := UnbindInstanceDomainsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func UntagResourcesWithOptions(client *Client, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func UntagResources(client *Client, request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := UntagResourcesWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify AppKey State
//
// @param request - UpdateAppKeyStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppKeyStateResponse
func UpdateAppKeyStateWithOptions(client *Client, request *UpdateAppKeyStateRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppKeyStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKeyId) {
		query["AppKeyId"] = request.AppKeyId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppKeyState"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppKeyStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify AppKey State
//
// @param request - UpdateAppKeyStateRequest
//
// @return UpdateAppKeyStateResponse
func UpdateAppKeyState(client *Client, request *UpdateAppKeyStateRequest) (_result *UpdateAppKeyStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppKeyStateResponse{}
	_body, _err := UpdateAppKeyStateWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the registration information of an Agent.
//
// @param tmpReq - UpdateAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAtiAgentRegisterInfoResponse
func UpdateAtiAgentRegisterInfoWithOptions(client *Client, tmpReq *UpdateAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAtiAgentRegisterInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Endpoints) {
		request.EndpointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoints, dara.String("Endpoints"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentDescription) {
		query["AgentDescription"] = request.AgentDescription
	}

	if !dara.IsNil(request.AgentDisplayName) {
		query["AgentDisplayName"] = request.AgentDisplayName
	}

	if !dara.IsNil(request.AgentHost) {
		query["AgentHost"] = request.AgentHost
	}

	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.AgentVersion) {
		query["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointsShrink) {
		query["Endpoints"] = request.EndpointsShrink
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the registration information of an Agent.
//
// @param request - UpdateAtiAgentRegisterInfoRequest
//
// @return UpdateAtiAgentRegisterInfoResponse
func UpdateAtiAgentRegisterInfo(client *Client, request *UpdateAtiAgentRegisterInfoRequest) (_result *UpdateAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAtiAgentRegisterInfoResponse{}
	_body, _err := UpdateAtiAgentRegisterInfoWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新告警设置
//
// @param request - UpdateAtiAlertSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAtiAlertSettingsResponse
func UpdateAtiAlertSettingsWithOptions(client *Client, request *UpdateAtiAlertSettingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateAtiAlertSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.AlertGroup) {
		query["AlertGroup"] = request.AlertGroup
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAtiAlertSettings"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAtiAlertSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新告警设置
//
// @param request - UpdateAtiAlertSettingsRequest
//
// @return UpdateAtiAlertSettingsResponse
func UpdateAtiAlertSettings(client *Client, request *UpdateAtiAlertSettingsRequest) (_result *UpdateAtiAlertSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAtiAlertSettingsResponse{}
	_body, _err := UpdateAtiAlertSettingsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a registrant profile.
//
// @param request - UpdateAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAtiRegistrantResponse
func UpdateAtiRegistrantWithOptions(client *Client, request *UpdateAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *UpdateAtiRegistrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cc) {
		query["Cc"] = request.Cc
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DocumentCode) {
		query["DocumentCode"] = request.DocumentCode
	}

	if !dara.IsNil(request.DocumentImage) {
		query["DocumentImage"] = request.DocumentImage
	}

	if !dara.IsNil(request.DocumentType) {
		query["DocumentType"] = request.DocumentType
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Street) {
		query["Street"] = request.Street
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a registrant profile.
//
// @param request - UpdateAtiRegistrantRequest
//
// @return UpdateAtiRegistrantResponse
func UpdateAtiRegistrant(client *Client, request *UpdateAtiRegistrantRequest) (_result *UpdateAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAtiRegistrantResponse{}
	_body, _err := UpdateAtiRegistrantWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of a specified address, such as the address name, type, and value.
//
// @param tmpReq - UpdateCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressResponse
func UpdateCloudGtmAddressWithOptions(client *Client, tmpReq *UpdateCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmAddressShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HealthTasks) {
		request.HealthTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HealthTasks, dara.String("HealthTasks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.AttributeInfo) {
		query["AttributeInfo"] = request.AttributeInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	if !dara.IsNil(request.HealthTasksShrink) {
		query["HealthTasks"] = request.HealthTasksShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of a specified address, such as the address name, type, and value.
//
// @param request - UpdateCloudGtmAddressRequest
//
// @return UpdateCloudGtmAddressResponse
func UpdateCloudGtmAddress(client *Client, request *UpdateCloudGtmAddressRequest) (_result *UpdateCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressResponse{}
	_body, _err := UpdateCloudGtmAddressWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address.
//
// Description:
//
// - The service status of an address is **active*	- if the address is **enabled*	- and its health check status is **Normal**.
//
// - The service status of an address is **unavailable*	- if the address is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressEnableStatusResponse
func UpdateCloudGtmAddressEnableStatusWithOptions(client *Client, request *UpdateCloudGtmAddressEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address.
//
// Description:
//
// - The service status of an address is **active*	- if the address is **enabled*	- and its health check status is **Normal**.
//
// - The service status of an address is **unavailable*	- if the address is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressEnableStatusRequest
//
// @return UpdateCloudGtmAddressEnableStatusResponse
func UpdateCloudGtmAddressEnableStatus(client *Client, request *UpdateCloudGtmAddressEnableStatusRequest) (_result *UpdateCloudGtmAddressEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressEnableStatusResponse{}
	_body, _err := UpdateCloudGtmAddressEnableStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the failover method for an address based on the specified parameters.
//
// @param request - UpdateCloudGtmAddressManualAvailableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressManualAvailableStatusResponse
func UpdateCloudGtmAddressManualAvailableStatusWithOptions(client *Client, request *UpdateCloudGtmAddressManualAvailableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressManualAvailableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.AvailableMode) {
		query["AvailableMode"] = request.AvailableMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ManualAvailableStatus) {
		query["ManualAvailableStatus"] = request.ManualAvailableStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressManualAvailableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressManualAvailableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the failover method for an address based on the specified parameters.
//
// @param request - UpdateCloudGtmAddressManualAvailableStatusRequest
//
// @return UpdateCloudGtmAddressManualAvailableStatusResponse
func UpdateCloudGtmAddressManualAvailableStatus(client *Client, request *UpdateCloudGtmAddressManualAvailableStatusRequest) (_result *UpdateCloudGtmAddressManualAvailableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressManualAvailableStatusResponse{}
	_body, _err := UpdateCloudGtmAddressManualAvailableStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolBasicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolBasicConfigResponse
func UpdateCloudGtmAddressPoolBasicConfigWithOptions(client *Client, request *UpdateCloudGtmAddressPoolBasicConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolBasicConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolBasicConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolBasicConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolBasicConfigRequest
//
// @return UpdateCloudGtmAddressPoolBasicConfigResponse
func UpdateCloudGtmAddressPoolBasicConfig(client *Client, request *UpdateCloudGtmAddressPoolBasicConfigRequest) (_result *UpdateCloudGtmAddressPoolBasicConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolBasicConfigResponse{}
	_body, _err := UpdateCloudGtmAddressPoolBasicConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address pool.
//
// Description:
//
// - An address pool is considered **active*	- if it is **enabled*	- and its health check status is **Normal**.
//
// - An address pool is considered **unavailable*	- if it is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressPoolEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolEnableStatusResponse
func UpdateCloudGtmAddressPoolEnableStatusWithOptions(client *Client, request *UpdateCloudGtmAddressPoolEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address pool.
//
// Description:
//
// - An address pool is considered **active*	- if it is **enabled*	- and its health check status is **Normal**.
//
// - An address pool is considered **unavailable*	- if it is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressPoolEnableStatusRequest
//
// @return UpdateCloudGtmAddressPoolEnableStatusResponse
func UpdateCloudGtmAddressPoolEnableStatus(client *Client, request *UpdateCloudGtmAddressPoolEnableStatusRequest) (_result *UpdateCloudGtmAddressPoolEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolEnableStatusResponse{}
	_body, _err := UpdateCloudGtmAddressPoolEnableStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolLbStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolLbStrategyResponse
func UpdateCloudGtmAddressPoolLbStrategyWithOptions(client *Client, request *UpdateCloudGtmAddressPoolLbStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolLbStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressLbStrategy) {
		query["AddressLbStrategy"] = request.AddressLbStrategy
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.SequenceLbStrategyMode) {
		query["SequenceLbStrategyMode"] = request.SequenceLbStrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolLbStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolLbStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolLbStrategyRequest
//
// @return UpdateCloudGtmAddressPoolLbStrategyResponse
func UpdateCloudGtmAddressPoolLbStrategy(client *Client, request *UpdateCloudGtmAddressPoolLbStrategyRequest) (_result *UpdateCloudGtmAddressPoolLbStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolLbStrategyResponse{}
	_body, _err := UpdateCloudGtmAddressPoolLbStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolRemarkResponse
func UpdateCloudGtmAddressPoolRemarkWithOptions(client *Client, request *UpdateCloudGtmAddressPoolRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolRemarkRequest
//
// @return UpdateCloudGtmAddressPoolRemarkResponse
func UpdateCloudGtmAddressPoolRemark(client *Client, request *UpdateCloudGtmAddressPoolRemarkRequest) (_result *UpdateCloudGtmAddressPoolRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolRemarkResponse{}
	_body, _err := UpdateCloudGtmAddressPoolRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for an address in Global Traffic Manager (GTM) 3.0.
//
// @param request - UpdateCloudGtmAddressRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressRemarkResponse
func UpdateCloudGtmAddressRemarkWithOptions(client *Client, request *UpdateCloudGtmAddressRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for an address in Global Traffic Manager (GTM) 3.0.
//
// @param request - UpdateCloudGtmAddressRemarkRequest
//
// @return UpdateCloudGtmAddressRemarkResponse
func UpdateCloudGtmAddressRemark(client *Client, request *UpdateCloudGtmAddressRemarkRequest) (_result *UpdateCloudGtmAddressRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressRemarkResponse{}
	_body, _err := UpdateCloudGtmAddressRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新全局流量管理告警设置
//
// @param tmpReq - UpdateCloudGtmGlobalAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmGlobalAlertResponse
func UpdateCloudGtmGlobalAlertWithOptions(client *Client, tmpReq *UpdateCloudGtmGlobalAlertRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmGlobalAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmGlobalAlertShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlertConfig) {
		request.AlertConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertConfig, dara.String("AlertConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AlertGroup) {
		request.AlertGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertGroup, dara.String("AlertGroup"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AlertConfigShrink) {
		query["AlertConfig"] = request.AlertConfigShrink
	}

	if !dara.IsNil(request.AlertGroupShrink) {
		query["AlertGroup"] = request.AlertGroupShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmGlobalAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmGlobalAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新全局流量管理告警设置
//
// @param request - UpdateCloudGtmGlobalAlertRequest
//
// @return UpdateCloudGtmGlobalAlertResponse
func UpdateCloudGtmGlobalAlert(client *Client, request *UpdateCloudGtmGlobalAlertRequest) (_result *UpdateCloudGtmGlobalAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmGlobalAlertResponse{}
	_body, _err := UpdateCloudGtmGlobalAlertWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateCloudGtmInstanceConfigAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigAlertResponse
func UpdateCloudGtmInstanceConfigAlertWithOptions(client *Client, tmpReq *UpdateCloudGtmInstanceConfigAlertRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmInstanceConfigAlertShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlertConfig) {
		request.AlertConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertConfig, dara.String("AlertConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AlertGroup) {
		request.AlertGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertGroup, dara.String("AlertGroup"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AlertConfigShrink) {
		query["AlertConfig"] = request.AlertConfigShrink
	}

	if !dara.IsNil(request.AlertGroupShrink) {
		query["AlertGroup"] = request.AlertGroupShrink
	}

	if !dara.IsNil(request.AlertMode) {
		query["AlertMode"] = request.AlertMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCloudGtmInstanceConfigAlertRequest
//
// @return UpdateCloudGtmInstanceConfigAlertResponse
func UpdateCloudGtmInstanceConfigAlert(client *Client, request *UpdateCloudGtmInstanceConfigAlertRequest) (_result *UpdateCloudGtmInstanceConfigAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigAlertResponse{}
	_body, _err := UpdateCloudGtmInstanceConfigAlertWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the global TTL configuration of a Global Traffic Manager (GTM) 3.0 instance based on the specified parameters.
//
// @param request - UpdateCloudGtmInstanceConfigBasicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigBasicResponse
func UpdateCloudGtmInstanceConfigBasicWithOptions(client *Client, request *UpdateCloudGtmInstanceConfigBasicRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigBasicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScheduleHostname) {
		query["ScheduleHostname"] = request.ScheduleHostname
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigBasic"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigBasicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the global TTL configuration of a Global Traffic Manager (GTM) 3.0 instance based on the specified parameters.
//
// @param request - UpdateCloudGtmInstanceConfigBasicRequest
//
// @return UpdateCloudGtmInstanceConfigBasicResponse
func UpdateCloudGtmInstanceConfigBasic(client *Client, request *UpdateCloudGtmInstanceConfigBasicRequest) (_result *UpdateCloudGtmInstanceConfigBasicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigBasicResponse{}
	_body, _err := UpdateCloudGtmInstanceConfigBasicWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enablement status of an instance configuration based on the input parameters.
//
// Description:
//
// - If a domain name is **enabled*	- and its health status is **Normal**, the service associated with the access domain name is **active**.
//
// - If a domain name is **disabled*	- or its health status is **abnormal**, the service associated with the access domain name is **unavailable**.
//
// @param request - UpdateCloudGtmInstanceConfigEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigEnableStatusResponse
func UpdateCloudGtmInstanceConfigEnableStatusWithOptions(client *Client, request *UpdateCloudGtmInstanceConfigEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enablement status of an instance configuration based on the input parameters.
//
// Description:
//
// - If a domain name is **enabled*	- and its health status is **Normal**, the service associated with the access domain name is **active**.
//
// - If a domain name is **disabled*	- or its health status is **abnormal**, the service associated with the access domain name is **unavailable**.
//
// @param request - UpdateCloudGtmInstanceConfigEnableStatusRequest
//
// @return UpdateCloudGtmInstanceConfigEnableStatusResponse
func UpdateCloudGtmInstanceConfigEnableStatus(client *Client, request *UpdateCloudGtmInstanceConfigEnableStatusRequest) (_result *UpdateCloudGtmInstanceConfigEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigEnableStatusResponse{}
	_body, _err := UpdateCloudGtmInstanceConfigEnableStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigLbStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigLbStrategyResponse
func UpdateCloudGtmInstanceConfigLbStrategyWithOptions(client *Client, request *UpdateCloudGtmInstanceConfigLbStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigLbStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolLbStrategy) {
		query["AddressPoolLbStrategy"] = request.AddressPoolLbStrategy
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SequenceLbStrategyMode) {
		query["SequenceLbStrategyMode"] = request.SequenceLbStrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigLbStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigLbStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigLbStrategyRequest
//
// @return UpdateCloudGtmInstanceConfigLbStrategyResponse
func UpdateCloudGtmInstanceConfigLbStrategy(client *Client, request *UpdateCloudGtmInstanceConfigLbStrategyRequest) (_result *UpdateCloudGtmInstanceConfigLbStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigLbStrategyResponse{}
	_body, _err := UpdateCloudGtmInstanceConfigLbStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigRemarkResponse
func UpdateCloudGtmInstanceConfigRemarkWithOptions(client *Client, request *UpdateCloudGtmInstanceConfigRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigRemarkRequest
//
// @return UpdateCloudGtmInstanceConfigRemarkResponse
func UpdateCloudGtmInstanceConfigRemark(client *Client, request *UpdateCloudGtmInstanceConfigRemarkRequest) (_result *UpdateCloudGtmInstanceConfigRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigRemarkResponse{}
	_body, _err := UpdateCloudGtmInstanceConfigRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCloudGtmInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceNameResponse
func UpdateCloudGtmInstanceNameWithOptions(client *Client, request *UpdateCloudGtmInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceName"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCloudGtmInstanceNameRequest
//
// @return UpdateCloudGtmInstanceNameResponse
func UpdateCloudGtmInstanceName(client *Client, request *UpdateCloudGtmInstanceNameRequest) (_result *UpdateCloudGtmInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceNameResponse{}
	_body, _err := UpdateCloudGtmInstanceNameWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a health check template.
//
// @param tmpReq - UpdateCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmMonitorTemplateResponse
func UpdateCloudGtmMonitorTemplateWithOptions(client *Client, tmpReq *UpdateCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmMonitorTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IspCityNodes) {
		request.IspCityNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IspCityNodes, dara.String("IspCityNodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.ExtendInfo) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.FailureRate) {
		query["FailureRate"] = request.FailureRate
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNodesShrink) {
		query["IspCityNodes"] = request.IspCityNodesShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a health check template.
//
// @param request - UpdateCloudGtmMonitorTemplateRequest
//
// @return UpdateCloudGtmMonitorTemplateResponse
func UpdateCloudGtmMonitorTemplate(client *Client, request *UpdateCloudGtmMonitorTemplateRequest) (_result *UpdateCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmMonitorTemplateResponse{}
	_body, _err := UpdateCloudGtmMonitorTemplateWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCloudGtmMonitorTemplateRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmMonitorTemplateRemarkResponse
func UpdateCloudGtmMonitorTemplateRemarkWithOptions(client *Client, request *UpdateCloudGtmMonitorTemplateRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmMonitorTemplateRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmMonitorTemplateRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmMonitorTemplateRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCloudGtmMonitorTemplateRemarkRequest
//
// @return UpdateCloudGtmMonitorTemplateRemarkResponse
func UpdateCloudGtmMonitorTemplateRemark(client *Client, request *UpdateCloudGtmMonitorTemplateRemarkRequest) (_result *UpdateCloudGtmMonitorTemplateRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmMonitorTemplateRemarkResponse{}
	_body, _err := UpdateCloudGtmMonitorTemplateRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom line specified by its unique ID.
//
// Description:
//
// For each IP range, the value of EndIp must be greater than or equal to the value of StartIp.
//
// The IP ranges of IP ranges cannot overlap across all custom lines for a domain name.
//
// @param request - UpdateCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomLineResponse
func UpdateCustomLineWithOptions(client *Client, request *UpdateCustomLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpSegment) {
		query["IpSegment"] = request.IpSegment
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineId) {
		query["LineId"] = request.LineId
	}

	if !dara.IsNil(request.LineName) {
		query["LineName"] = request.LineName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomLine"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom line specified by its unique ID.
//
// Description:
//
// For each IP range, the value of EndIp must be greater than or equal to the value of StartIp.
//
// The IP ranges of IP ranges cannot overlap across all custom lines for a domain name.
//
// @param request - UpdateCustomLineRequest
//
// @return UpdateCustomLineResponse
func UpdateCustomLine(client *Client, request *UpdateCustomLineRequest) (_result *UpdateCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomLineResponse{}
	_body, _err := UpdateCustomLineWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the weight of a DNS record based on the specified parameters.
//
// @param request - UpdateDNSSLBWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDNSSLBWeightResponse
func UpdateDNSSLBWeightWithOptions(client *Client, request *UpdateDNSSLBWeightRequest, runtime *dara.RuntimeOptions) (_result *UpdateDNSSLBWeightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDNSSLBWeight"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDNSSLBWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the weight of a DNS record based on the specified parameters.
//
// @param request - UpdateDNSSLBWeightRequest
//
// @return UpdateDNSSLBWeightResponse
func UpdateDNSSLBWeight(client *Client, request *UpdateDNSSLBWeightRequest) (_result *UpdateDNSSLBWeightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDNSSLBWeightResponse{}
	_body, _err := UpdateDNSSLBWeightWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a DNS authoritative proxy domain.
//
// @param request - UpdateDnsCacheDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsCacheDomainResponse
func UpdateDnsCacheDomainWithOptions(client *Client, request *UpdateDnsCacheDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsCacheDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheTtlMax) {
		query["CacheTtlMax"] = request.CacheTtlMax
	}

	if !dara.IsNil(request.CacheTtlMin) {
		query["CacheTtlMin"] = request.CacheTtlMin
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceDnsServer) {
		query["SourceDnsServer"] = request.SourceDnsServer
	}

	if !dara.IsNil(request.SourceEdns) {
		query["SourceEdns"] = request.SourceEdns
	}

	if !dara.IsNil(request.SourceProtocol) {
		query["SourceProtocol"] = request.SourceProtocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsCacheDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsCacheDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a DNS authoritative proxy domain.
//
// @param request - UpdateDnsCacheDomainRequest
//
// @return UpdateDnsCacheDomainResponse
func UpdateDnsCacheDomain(client *Client, request *UpdateDnsCacheDomainRequest) (_result *UpdateDnsCacheDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsCacheDomainResponse{}
	_body, _err := UpdateDnsCacheDomainWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remark for a domain name in the DNS cache.
//
// @param request - UpdateDnsCacheDomainRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsCacheDomainRemarkResponse
func UpdateDnsCacheDomainRemarkWithOptions(client *Client, request *UpdateDnsCacheDomainRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsCacheDomainRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsCacheDomainRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsCacheDomainRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remark for a domain name in the DNS cache.
//
// @param request - UpdateDnsCacheDomainRemarkRequest
//
// @return UpdateDnsCacheDomainRemarkResponse
func UpdateDnsCacheDomainRemark(client *Client, request *UpdateDnsCacheDomainRemarkRequest) (_result *UpdateDnsCacheDomainRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsCacheDomainRemarkResponse{}
	_body, _err := UpdateDnsCacheDomainRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an existing access policy.
//
// @param request - UpdateDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmAccessStrategyResponse
func UpdateDnsGtmAccessStrategyWithOptions(client *Client, request *UpdateDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.DefaultAddrPool) {
		query["DefaultAddrPool"] = request.DefaultAddrPool
	}

	if !dara.IsNil(request.DefaultAddrPoolType) {
		query["DefaultAddrPoolType"] = request.DefaultAddrPoolType
	}

	if !dara.IsNil(request.DefaultLatencyOptimization) {
		query["DefaultLatencyOptimization"] = request.DefaultLatencyOptimization
	}

	if !dara.IsNil(request.DefaultLbaStrategy) {
		query["DefaultLbaStrategy"] = request.DefaultLbaStrategy
	}

	if !dara.IsNil(request.DefaultMaxReturnAddrNum) {
		query["DefaultMaxReturnAddrNum"] = request.DefaultMaxReturnAddrNum
	}

	if !dara.IsNil(request.DefaultMinAvailableAddrNum) {
		query["DefaultMinAvailableAddrNum"] = request.DefaultMinAvailableAddrNum
	}

	if !dara.IsNil(request.FailoverAddrPool) {
		query["FailoverAddrPool"] = request.FailoverAddrPool
	}

	if !dara.IsNil(request.FailoverAddrPoolType) {
		query["FailoverAddrPoolType"] = request.FailoverAddrPoolType
	}

	if !dara.IsNil(request.FailoverLatencyOptimization) {
		query["FailoverLatencyOptimization"] = request.FailoverLatencyOptimization
	}

	if !dara.IsNil(request.FailoverLbaStrategy) {
		query["FailoverLbaStrategy"] = request.FailoverLbaStrategy
	}

	if !dara.IsNil(request.FailoverMaxReturnAddrNum) {
		query["FailoverMaxReturnAddrNum"] = request.FailoverMaxReturnAddrNum
	}

	if !dara.IsNil(request.FailoverMinAvailableAddrNum) {
		query["FailoverMinAvailableAddrNum"] = request.FailoverMinAvailableAddrNum
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lines) {
		query["Lines"] = request.Lines
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an existing access policy.
//
// @param request - UpdateDnsGtmAccessStrategyRequest
//
// @return UpdateDnsGtmAccessStrategyResponse
func UpdateDnsGtmAccessStrategy(client *Client, request *UpdateDnsGtmAccessStrategyRequest) (_result *UpdateDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmAccessStrategyResponse{}
	_body, _err := UpdateDnsGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of an address pool.
//
// @param request - UpdateDnsGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmAddressPoolResponse
func UpdateDnsGtmAddressPoolWithOptions(client *Client, request *UpdateDnsGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LbaStrategy) {
		query["LbaStrategy"] = request.LbaStrategy
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an address pool.
//
// @param request - UpdateDnsGtmAddressPoolRequest
//
// @return UpdateDnsGtmAddressPoolResponse
func UpdateDnsGtmAddressPool(client *Client, request *UpdateDnsGtmAddressPoolRequest) (_result *UpdateDnsGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmAddressPoolResponse{}
	_body, _err := UpdateDnsGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateDnsGtmInstanceGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmInstanceGlobalConfigResponse
func UpdateDnsGtmInstanceGlobalConfigWithOptions(client *Client, request *UpdateDnsGtmInstanceGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmInstanceGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.AlertGroup) {
		query["AlertGroup"] = request.AlertGroup
	}

	if !dara.IsNil(request.CnameType) {
		query["CnameType"] = request.CnameType
	}

	if !dara.IsNil(request.ForceUpdate) {
		query["ForceUpdate"] = request.ForceUpdate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PublicCnameMode) {
		query["PublicCnameMode"] = request.PublicCnameMode
	}

	if !dara.IsNil(request.PublicRr) {
		query["PublicRr"] = request.PublicRr
	}

	if !dara.IsNil(request.PublicUserDomainName) {
		query["PublicUserDomainName"] = request.PublicUserDomainName
	}

	if !dara.IsNil(request.PublicZoneName) {
		query["PublicZoneName"] = request.PublicZoneName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmInstanceGlobalConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmInstanceGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateDnsGtmInstanceGlobalConfigRequest
//
// @return UpdateDnsGtmInstanceGlobalConfigResponse
func UpdateDnsGtmInstanceGlobalConfig(client *Client, request *UpdateDnsGtmInstanceGlobalConfigRequest) (_result *UpdateDnsGtmInstanceGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmInstanceGlobalConfigResponse{}
	_body, _err := UpdateDnsGtmInstanceGlobalConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a health check configuration.
//
// @param request - UpdateDnsGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmMonitorResponse
func UpdateDnsGtmMonitorWithOptions(client *Client, request *UpdateDnsGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a health check configuration.
//
// @param request - UpdateDnsGtmMonitorRequest
//
// @return UpdateDnsGtmMonitorResponse
func UpdateDnsGtmMonitor(client *Client, request *UpdateDnsGtmMonitorRequest) (_result *UpdateDnsGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmMonitorResponse{}
	_body, _err := UpdateDnsGtmMonitorWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a domain name group.
//
// Description:
//
// This operation modifies the name of an existing domain name group.
//
// @param request - UpdateDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainGroupResponse
func UpdateDomainGroupWithOptions(client *Client, request *UpdateDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a domain name group.
//
// Description:
//
// This operation modifies the name of an existing domain name group.
//
// @param request - UpdateDomainGroupRequest
//
// @return UpdateDomainGroupResponse
func UpdateDomainGroup(client *Client, request *UpdateDomainGroupRequest) (_result *UpdateDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainGroupResponse{}
	_body, _err := UpdateDomainGroupWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a DNS record based on the specified parameters.
//
// @param request - UpdateDomainRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainRecordResponse
func UpdateDomainRecordWithOptions(client *Client, request *UpdateDomainRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RR) {
		query["RR"] = request.RR
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.TTL) {
		query["TTL"] = request.TTL
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a DNS record based on the specified parameters.
//
// @param request - UpdateDomainRecordRequest
//
// @return UpdateDomainRecordResponse
func UpdateDomainRecord(client *Client, request *UpdateDomainRecordRequest) (_result *UpdateDomainRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainRecordResponse{}
	_body, _err := UpdateDomainRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks of a DNS record.
//
// @param request - UpdateDomainRecordRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainRecordRemarkResponse
func UpdateDomainRecordRemarkWithOptions(client *Client, request *UpdateDomainRecordRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainRecordRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainRecordRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainRecordRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks of a DNS record.
//
// @param request - UpdateDomainRecordRemarkRequest
//
// @return UpdateDomainRecordRemarkResponse
func UpdateDomainRecordRemark(client *Client, request *UpdateDomainRecordRemarkRequest) (_result *UpdateDomainRecordRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainRecordRemarkResponse{}
	_body, _err := UpdateDomainRecordRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for a domain name.
//
// @param request - UpdateDomainRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainRemarkResponse
func UpdateDomainRemarkWithOptions(client *Client, request *UpdateDomainRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for a domain name.
//
// @param request - UpdateDomainRemarkRequest
//
// @return UpdateDomainRemarkResponse
func UpdateDomainRemark(client *Client, request *UpdateDomainRemarkRequest) (_result *UpdateDomainRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainRemarkResponse{}
	_body, _err := UpdateDomainRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an access policy.
//
// @param request - UpdateGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmAccessStrategyResponse
func UpdateGtmAccessStrategyWithOptions(client *Client, request *UpdateGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessLines) {
		query["AccessLines"] = request.AccessLines
	}

	if !dara.IsNil(request.DefaultAddrPoolId) {
		query["DefaultAddrPoolId"] = request.DefaultAddrPoolId
	}

	if !dara.IsNil(request.FailoverAddrPoolId) {
		query["FailoverAddrPoolId"] = request.FailoverAddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an access policy.
//
// @param request - UpdateGtmAccessStrategyRequest
//
// @return UpdateGtmAccessStrategyResponse
func UpdateGtmAccessStrategy(client *Client, request *UpdateGtmAccessStrategyRequest) (_result *UpdateGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmAccessStrategyResponse{}
	_body, _err := UpdateGtmAccessStrategyWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an address pool.
//
// @param request - UpdateGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmAddressPoolResponse
func UpdateGtmAddressPoolWithOptions(client *Client, request *UpdateGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MinAvailableAddrNum) {
		query["MinAvailableAddrNum"] = request.MinAvailableAddrNum
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an address pool.
//
// @param request - UpdateGtmAddressPoolRequest
//
// @return UpdateGtmAddressPoolResponse
func UpdateGtmAddressPool(client *Client, request *UpdateGtmAddressPoolRequest) (_result *UpdateGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmAddressPoolResponse{}
	_body, _err := UpdateGtmAddressPoolWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the global configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateGtmInstanceGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmInstanceGlobalConfigResponse
func UpdateGtmInstanceGlobalConfigWithOptions(client *Client, request *UpdateGtmInstanceGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmInstanceGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertGroup) {
		query["AlertGroup"] = request.AlertGroup
	}

	if !dara.IsNil(request.CnameCustomDomainName) {
		query["CnameCustomDomainName"] = request.CnameCustomDomainName
	}

	if !dara.IsNil(request.CnameMode) {
		query["CnameMode"] = request.CnameMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LbaStrategy) {
		query["LbaStrategy"] = request.LbaStrategy
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.UserDomainName) {
		query["UserDomainName"] = request.UserDomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmInstanceGlobalConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmInstanceGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the global configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateGtmInstanceGlobalConfigRequest
//
// @return UpdateGtmInstanceGlobalConfigResponse
func UpdateGtmInstanceGlobalConfig(client *Client, request *UpdateGtmInstanceGlobalConfigRequest) (_result *UpdateGtmInstanceGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmInstanceGlobalConfigResponse{}
	_body, _err := UpdateGtmInstanceGlobalConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a health check configuration.
//
// @param request - UpdateGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmMonitorResponse
func UpdateGtmMonitorWithOptions(client *Client, request *UpdateGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a health check configuration.
//
// @param request - UpdateGtmMonitorRequest
//
// @return UpdateGtmMonitorResponse
func UpdateGtmMonitor(client *Client, request *UpdateGtmMonitorRequest) (_result *UpdateGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmMonitorResponse{}
	_body, _err := UpdateGtmMonitorWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a disaster recovery plan.
//
// @param request - UpdateGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmRecoveryPlanResponse
func UpdateGtmRecoveryPlanWithOptions(client *Client, request *UpdateGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaultAddrPool) {
		query["FaultAddrPool"] = request.FaultAddrPool
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a disaster recovery plan.
//
// @param request - UpdateGtmRecoveryPlanRequest
//
// @return UpdateGtmRecoveryPlanResponse
func UpdateGtmRecoveryPlan(client *Client, request *UpdateGtmRecoveryPlanRequest) (_result *UpdateGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmRecoveryPlanResponse{}
	_body, _err := UpdateGtmRecoveryPlanWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cache refresh plan.
//
// @param request - UpdateIspFlushCacheInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIspFlushCacheInstanceConfigResponse
func UpdateIspFlushCacheInstanceConfigWithOptions(client *Client, request *UpdateIspFlushCacheInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateIspFlushCacheInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIspFlushCacheInstanceConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIspFlushCacheInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cache refresh plan.
//
// @param request - UpdateIspFlushCacheInstanceConfigRequest
//
// @return UpdateIspFlushCacheInstanceConfigResponse
func UpdateIspFlushCacheInstanceConfig(client *Client, request *UpdateIspFlushCacheInstanceConfigRequest) (_result *UpdateIspFlushCacheInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIspFlushCacheInstanceConfigResponse{}
	_body, _err := UpdateIspFlushCacheInstanceConfigWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a built-in authoritative record for recursive resolution.
//
// Description:
//
// If a DNS record is locked, it cannot be deleted.
//
// @param request - UpdateRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordResponse
func UpdateRecursionRecordWithOptions(client *Client, request *UpdateRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a built-in authoritative record for recursive resolution.
//
// Description:
//
// If a DNS record is locked, it cannot be deleted.
//
// @param request - UpdateRecursionRecordRequest
//
// @return UpdateRecursionRecordResponse
func UpdateRecursionRecord(client *Client, request *UpdateRecursionRecordRequest) (_result *UpdateRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordResponse{}
	_body, _err := UpdateRecursionRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enable status of a recursion record.
//
// @param request - UpdateRecursionRecordEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordEnableStatusResponse
func UpdateRecursionRecordEnableStatusWithOptions(client *Client, request *UpdateRecursionRecordEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enable status of a recursion record.
//
// @param request - UpdateRecursionRecordEnableStatusRequest
//
// @return UpdateRecursionRecordEnableStatusResponse
func UpdateRecursionRecordEnableStatus(client *Client, request *UpdateRecursionRecordEnableStatusRequest) (_result *UpdateRecursionRecordEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordEnableStatusResponse{}
	_body, _err := UpdateRecursionRecordEnableStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for a built-in authoritative record in HTTPDNS.
//
// @param request - UpdateRecursionRecordRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordRemarkResponse
func UpdateRecursionRecordRemarkWithOptions(client *Client, request *UpdateRecursionRecordRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for a built-in authoritative record in HTTPDNS.
//
// @param request - UpdateRecursionRecordRemarkRequest
//
// @return UpdateRecursionRecordRemarkResponse
func UpdateRecursionRecordRemark(client *Client, request *UpdateRecursionRecordRemarkRequest) (_result *UpdateRecursionRecordRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordRemarkResponse{}
	_body, _err := UpdateRecursionRecordRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the weight of a DNS record for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordWeightResponse
func UpdateRecursionRecordWeightWithOptions(client *Client, request *UpdateRecursionRecordWeightRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordWeightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordWeight"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the weight of a DNS record for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightRequest
//
// @return UpdateRecursionRecordWeightResponse
func UpdateRecursionRecordWeight(client *Client, request *UpdateRecursionRecordWeightRequest) (_result *UpdateRecursionRecordWeightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordWeightResponse{}
	_body, _err := UpdateRecursionRecordWeightWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enabled status of the weight algorithm for a DNS record used for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordWeightEnableStatusResponse
func UpdateRecursionRecordWeightEnableStatusWithOptions(client *Client, request *UpdateRecursionRecordWeightEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordWeightEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordWeightEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordWeightEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enabled status of the weight algorithm for a DNS record used for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightEnableStatusRequest
//
// @return UpdateRecursionRecordWeightEnableStatusResponse
func UpdateRecursionRecordWeightEnableStatus(client *Client, request *UpdateRecursionRecordWeightEnableStatusRequest) (_result *UpdateRecursionRecordWeightEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordWeightEnableStatusResponse{}
	_body, _err := UpdateRecursionRecordWeightEnableStatusWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the effective scope of a built-in authoritative domain name zone in HTTPDNS.
//
// @param tmpReq - UpdateRecursionZoneEffectiveScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionZoneEffectiveScopeResponse
func UpdateRecursionZoneEffectiveScopeWithOptions(client *Client, tmpReq *UpdateRecursionZoneEffectiveScopeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionZoneEffectiveScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRecursionZoneEffectiveScopeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EffectiveScopes) {
		request.EffectiveScopesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EffectiveScopes, dara.String("EffectiveScopes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EffectiveScopesShrink) {
		query["EffectiveScopes"] = request.EffectiveScopesShrink
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionZoneEffectiveScope"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionZoneEffectiveScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the effective scope of a built-in authoritative domain name zone in HTTPDNS.
//
// @param request - UpdateRecursionZoneEffectiveScopeRequest
//
// @return UpdateRecursionZoneEffectiveScopeResponse
func UpdateRecursionZoneEffectiveScope(client *Client, request *UpdateRecursionZoneEffectiveScopeRequest) (_result *UpdateRecursionZoneEffectiveScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionZoneEffectiveScopeResponse{}
	_body, _err := UpdateRecursionZoneEffectiveScopeWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the recursive proxy mode for a zone.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP segments for the domain name cannot overlap across custom lines.
//
// @param request - UpdateRecursionZoneProxyPatternRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionZoneProxyPatternResponse
func UpdateRecursionZoneProxyPatternWithOptions(client *Client, request *UpdateRecursionZoneProxyPatternRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionZoneProxyPatternResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProxyPattern) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionZoneProxyPattern"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionZoneProxyPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the recursive proxy mode for a zone.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP segments for the domain name cannot overlap across custom lines.
//
// @param request - UpdateRecursionZoneProxyPatternRequest
//
// @return UpdateRecursionZoneProxyPatternResponse
func UpdateRecursionZoneProxyPattern(client *Client, request *UpdateRecursionZoneProxyPatternRequest) (_result *UpdateRecursionZoneProxyPatternResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionZoneProxyPatternResponse{}
	_body, _err := UpdateRecursionZoneProxyPatternWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the remarks for a built-in authoritative domain name zone used for recursive resolution.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// For a domain name, the IP address ranges of all IP segments in all custom lines cannot overlap.
//
// @param request - UpdateRecursionZoneRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionZoneRemarkResponse
func UpdateRecursionZoneRemarkWithOptions(client *Client, request *UpdateRecursionZoneRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionZoneRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionZoneRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionZoneRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the remarks for a built-in authoritative domain name zone used for recursive resolution.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// For a domain name, the IP address ranges of all IP segments in all custom lines cannot overlap.
//
// @param request - UpdateRecursionZoneRemarkRequest
//
// @return UpdateRecursionZoneRemarkResponse
func UpdateRecursionZoneRemark(client *Client, request *UpdateRecursionZoneRemarkRequest) (_result *UpdateRecursionZoneRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionZoneRemarkResponse{}
	_body, _err := UpdateRecursionZoneRemarkWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// ## Request description
//
// - This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRspDomainServerProhibitStatusForGatewayResponse
func UpdateRspDomainServerProhibitStatusForGatewayWithOptions(client *Client, request *UpdateRspDomainServerProhibitStatusForGatewayRequest, runtime *dara.RuntimeOptions) (_result *UpdateRspDomainServerProhibitStatusForGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddStatusList) {
		query["AddStatusList"] = request.AddStatusList
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteStatusList) {
		query["DeleteStatusList"] = request.DeleteStatusList
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRspDomainServerProhibitStatusForGateway"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// ## Request description
//
// - This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayRequest
//
// @return UpdateRspDomainServerProhibitStatusForGatewayResponse
func UpdateRspDomainServerProhibitStatusForGateway(client *Client, request *UpdateRspDomainServerProhibitStatusForGatewayRequest) (_result *UpdateRspDomainServerProhibitStatusForGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayResponse{}
	_body, _err := UpdateRspDomainServerProhibitStatusForGatewayWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayOteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRspDomainServerProhibitStatusForGatewayOteResponse
func UpdateRspDomainServerProhibitStatusForGatewayOteWithOptions(client *Client, request *UpdateRspDomainServerProhibitStatusForGatewayOteRequest, runtime *dara.RuntimeOptions) (_result *UpdateRspDomainServerProhibitStatusForGatewayOteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddStatusList) {
		query["AddStatusList"] = request.AddStatusList
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteStatusList) {
		query["DeleteStatusList"] = request.DeleteStatusList
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRspDomainServerProhibitStatusForGatewayOte"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayOteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayOteRequest
//
// @return UpdateRspDomainServerProhibitStatusForGatewayOteResponse
func UpdateRspDomainServerProhibitStatusForGatewayOte(client *Client, request *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) (_result *UpdateRspDomainServerProhibitStatusForGatewayOteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayOteResponse{}
	_body, _err := UpdateRspDomainServerProhibitStatusForGatewayOteWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether an instance hostname can be added.
//
// @param request - ValidateDnsGtmCnameRrCanUseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateDnsGtmCnameRrCanUseResponse
func ValidateDnsGtmCnameRrCanUseWithOptions(client *Client, request *ValidateDnsGtmCnameRrCanUseRequest, runtime *dara.RuntimeOptions) (_result *ValidateDnsGtmCnameRrCanUseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CnameMode) {
		query["CnameMode"] = request.CnameMode
	}

	if !dara.IsNil(request.CnameRr) {
		query["CnameRr"] = request.CnameRr
	}

	if !dara.IsNil(request.CnameType) {
		query["CnameType"] = request.CnameType
	}

	if !dara.IsNil(request.CnameZone) {
		query["CnameZone"] = request.CnameZone
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateDnsGtmCnameRrCanUse"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateDnsGtmCnameRrCanUseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether an instance hostname can be added.
//
// @param request - ValidateDnsGtmCnameRrCanUseRequest
//
// @return ValidateDnsGtmCnameRrCanUseResponse
func ValidateDnsGtmCnameRrCanUse(client *Client, request *ValidateDnsGtmCnameRrCanUseRequest) (_result *ValidateDnsGtmCnameRrCanUseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateDnsGtmCnameRrCanUseResponse{}
	_body, _err := ValidateDnsGtmCnameRrCanUseWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates a UDP IP address segment for Public DNS.
//
// @param request - ValidatePdnsUdpIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidatePdnsUdpIpSegmentResponse
func ValidatePdnsUdpIpSegmentWithOptions(client *Client, request *ValidatePdnsUdpIpSegmentRequest, runtime *dara.RuntimeOptions) (_result *ValidatePdnsUdpIpSegmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.IpToken) {
		query["IpToken"] = request.IpToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidatePdnsUdpIpSegment"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidatePdnsUdpIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates a UDP IP address segment for Public DNS.
//
// @param request - ValidatePdnsUdpIpSegmentRequest
//
// @return ValidatePdnsUdpIpSegmentResponse
func ValidatePdnsUdpIpSegment(client *Client, request *ValidatePdnsUdpIpSegmentRequest) (_result *ValidatePdnsUdpIpSegmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidatePdnsUdpIpSegmentResponse{}
	_body, _err := ValidatePdnsUdpIpSegmentWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies DNS records.
//
// @param request - VerifyAtiAgentDnsRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAtiAgentDnsRecordsResponse
func VerifyAtiAgentDnsRecordsWithOptions(client *Client, request *VerifyAtiAgentDnsRecordsRequest, runtime *dara.RuntimeOptions) (_result *VerifyAtiAgentDnsRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyAtiAgentDnsRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyAtiAgentDnsRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies DNS records.
//
// @param request - VerifyAtiAgentDnsRecordsRequest
//
// @return VerifyAtiAgentDnsRecordsResponse
func VerifyAtiAgentDnsRecords(client *Client, request *VerifyAtiAgentDnsRecordsRequest) (_result *VerifyAtiAgentDnsRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyAtiAgentDnsRecordsResponse{}
	_body, _err := VerifyAtiAgentDnsRecordsWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 触发 ACME 预检
//
// @param request - VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse
func VerifyAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(client *Client, request *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest, runtime *dara.RuntimeOptions) (_result *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyAtiAgentRegisterInfoAcmeChallengeRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发 ACME 预检
//
// @param request - VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @return VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse
func VerifyAtiAgentRegisterInfoAcmeChallengeRecord(client *Client, request *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) (_result *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := VerifyAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(client, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
