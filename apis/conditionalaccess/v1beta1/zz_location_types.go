/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CountryObservation struct {

	// List of countries and/or regions in two-letter format specified by ISO 3166-2.
	CountriesAndRegions []*string `json:"countriesAndRegions,omitempty" tf:"countries_and_regions,omitempty"`

	// Whether IP addresses that don't map to a country or region should be included in the named location. Defaults to false.
	IncludeUnknownCountriesAndRegions *bool `json:"includeUnknownCountriesAndRegions,omitempty" tf:"include_unknown_countries_and_regions,omitempty"`
}

type CountryParameters struct {

	// List of countries and/or regions in two-letter format specified by ISO 3166-2.
	// +kubebuilder:validation:Required
	CountriesAndRegions []*string `json:"countriesAndRegions" tf:"countries_and_regions,omitempty"`

	// Whether IP addresses that don't map to a country or region should be included in the named location. Defaults to false.
	// +kubebuilder:validation:Optional
	IncludeUnknownCountriesAndRegions *bool `json:"includeUnknownCountriesAndRegions,omitempty" tf:"include_unknown_countries_and_regions,omitempty"`
}

type IPObservation struct {

	// List of IP address ranges in IPv4 CIDR format (e.g. 1.2.3.4/32) or any allowable IPv6 format from IETF RFC596.
	IPRanges []*string `json:"ipRanges,omitempty" tf:"ip_ranges,omitempty"`

	// Whether the named location is trusted. Defaults to false.
	Trusted *bool `json:"trusted,omitempty" tf:"trusted,omitempty"`
}

type IPParameters struct {

	// List of IP address ranges in IPv4 CIDR format (e.g. 1.2.3.4/32) or any allowable IPv6 format from IETF RFC596.
	// +kubebuilder:validation:Required
	IPRanges []*string `json:"ipRanges" tf:"ip_ranges,omitempty"`

	// Whether the named location is trusted. Defaults to false.
	// +kubebuilder:validation:Optional
	Trusted *bool `json:"trusted,omitempty" tf:"trusted,omitempty"`
}

type LocationObservation struct {

	// A country block as documented below, which configures a country-based named location.
	Country []CountryObservation `json:"country,omitempty" tf:"country,omitempty"`

	// The friendly name for this named location.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the named location.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An ip block as documented below, which configures an IP-based named location.
	IP []IPObservation `json:"ip,omitempty" tf:"ip,omitempty"`
}

type LocationParameters struct {

	// A country block as documented below, which configures a country-based named location.
	// +kubebuilder:validation:Optional
	Country []CountryParameters `json:"country,omitempty" tf:"country,omitempty"`

	// The friendly name for this named location.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// An ip block as documented below, which configures an IP-based named location.
	// +kubebuilder:validation:Optional
	IP []IPParameters `json:"ip,omitempty" tf:"ip,omitempty"`
}

// LocationSpec defines the desired state of Location
type LocationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationParameters `json:"forProvider"`
}

// LocationStatus defines the observed state of Location.
type LocationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Location is the Schema for the Locations API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Location struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   LocationSpec   `json:"spec"`
	Status LocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationList contains a list of Locations
type LocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Location `json:"items"`
}

// Repository type metadata.
var (
	Location_Kind             = "Location"
	Location_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Location_Kind}.String()
	Location_KindAPIVersion   = Location_Kind + "." + CRDGroupVersion.String()
	Location_GroupVersionKind = CRDGroupVersion.WithKind(Location_Kind)
)

func init() {
	SchemeBuilder.Register(&Location{}, &LocationList{})
}
