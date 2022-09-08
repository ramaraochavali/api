// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha3

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for DestinationRule
func (this *DestinationRule) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DestinationRule
func (this *DestinationRule) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicy
func (this *TrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicy
func (this *TrafficPolicy) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicy_PortTrafficPolicy
func (this *TrafficPolicy_PortTrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicy_PortTrafficPolicy
func (this *TrafficPolicy_PortTrafficPolicy) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicy_TunnelSettings
func (this *TrafficPolicy_TunnelSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicy_TunnelSettings
func (this *TrafficPolicy_TunnelSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Subset
func (this *Subset) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Subset
func (this *Subset) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings
func (this *LoadBalancerSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings
func (this *LoadBalancerSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings_ConsistentHashLB
func (this *LoadBalancerSettings_ConsistentHashLB) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings_ConsistentHashLB_RingHash
func (this *LoadBalancerSettings_ConsistentHashLB_RingHash) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings_ConsistentHashLB_RingHash
func (this *LoadBalancerSettings_ConsistentHashLB_RingHash) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings_ConsistentHashLB_MagLev
func (this *LoadBalancerSettings_ConsistentHashLB_MagLev) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings_ConsistentHashLB_MagLev
func (this *LoadBalancerSettings_ConsistentHashLB_MagLev) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings_ConsistentHashLB_HTTPCookie
func (this *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings_ConsistentHashLB_HTTPCookie
func (this *LoadBalancerSettings_ConsistentHashLB_HTTPCookie) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LoadBalancerSettings_PersistentSession
func (this *LoadBalancerSettings_PersistentSession) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LoadBalancerSettings_PersistentSession
func (this *LoadBalancerSettings_PersistentSession) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings
func (this *ConnectionPoolSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings
func (this *ConnectionPoolSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings_TCPSettings
func (this *ConnectionPoolSettings_TCPSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings_TCPSettings
func (this *ConnectionPoolSettings_TCPSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings_TCPSettings_TcpKeepalive
func (this *ConnectionPoolSettings_TCPSettings_TcpKeepalive) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings_TCPSettings_TcpKeepalive
func (this *ConnectionPoolSettings_TCPSettings_TcpKeepalive) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings_HTTPSettings
func (this *ConnectionPoolSettings_HTTPSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings_HTTPSettings
func (this *ConnectionPoolSettings_HTTPSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OutlierDetection
func (this *OutlierDetection) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OutlierDetection
func (this *OutlierDetection) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ClientTLSSettings
func (this *ClientTLSSettings) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ClientTLSSettings
func (this *ClientTLSSettings) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PersistentCookie
func (this *PersistentCookie) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PersistentCookie
func (this *PersistentCookie) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LocalityLoadBalancerSetting
func (this *LocalityLoadBalancerSetting) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LocalityLoadBalancerSetting
func (this *LocalityLoadBalancerSetting) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LocalityLoadBalancerSetting_Distribute
func (this *LocalityLoadBalancerSetting_Distribute) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LocalityLoadBalancerSetting_Distribute
func (this *LocalityLoadBalancerSetting_Distribute) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for LocalityLoadBalancerSetting_Failover
func (this *LocalityLoadBalancerSetting_Failover) MarshalJSON() ([]byte, error) {
	str, err := DestinationRuleMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for LocalityLoadBalancerSetting_Failover
func (this *LocalityLoadBalancerSetting_Failover) UnmarshalJSON(b []byte) error {
	return DestinationRuleUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	DestinationRuleMarshaler   = &jsonpb.Marshaler{}
	DestinationRuleUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
