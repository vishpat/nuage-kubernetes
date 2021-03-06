/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IKEGatewayProfileIdentity represents the Identity of the object
var IKEGatewayProfileIdentity = bambou.Identity{
	Name:     "ikegatewayprofile",
	Category: "ikegatewayprofiles",
}

// IKEGatewayProfilesList represents a list of IKEGatewayProfiles
type IKEGatewayProfilesList []*IKEGatewayProfile

// IKEGatewayProfilesAncestor is the interface of an ancestor of a IKEGatewayProfile must implement.
type IKEGatewayProfilesAncestor interface {
	IKEGatewayProfiles(*bambou.FetchingInfo) (IKEGatewayProfilesList, *bambou.Error)
	CreateIKEGatewayProfiles(*IKEGatewayProfile) *bambou.Error
}

// IKEGatewayProfile represents the model of a ikegatewayprofile
type IKEGatewayProfile struct {
	ID                               string `json:"ID,omitempty"`
	ParentID                         string `json:"parentID,omitempty"`
	ParentType                       string `json:"parentType,omitempty"`
	Owner                            string `json:"owner,omitempty"`
	IKEGatewayIdentifier             string `json:"IKEGatewayIdentifier,omitempty"`
	IKEGatewayIdentifierType         string `json:"IKEGatewayIdentifierType,omitempty"`
	Name                             string `json:"name,omitempty"`
	LastUpdatedBy                    string `json:"lastUpdatedBy,omitempty"`
	ServiceClass                     string `json:"serviceClass,omitempty"`
	Description                      string `json:"description,omitempty"`
	AntiReplayCheck                  bool   `json:"antiReplayCheck"`
	EntityScope                      string `json:"entityScope,omitempty"`
	AssociatedEnterpriseID           string `json:"associatedEnterpriseID,omitempty"`
	AssociatedIKEAuthenticationID    string `json:"associatedIKEAuthenticationID,omitempty"`
	AssociatedIKEAuthenticationType  string `json:"associatedIKEAuthenticationType,omitempty"`
	AssociatedIKEEncryptionProfileID string `json:"associatedIKEEncryptionProfileID,omitempty"`
	AssociatedIKEGatewayID           string `json:"associatedIKEGatewayID,omitempty"`
	ExternalID                       string `json:"externalID,omitempty"`
}

// NewIKEGatewayProfile returns a new *IKEGatewayProfile
func NewIKEGatewayProfile() *IKEGatewayProfile {

	return &IKEGatewayProfile{}
}

// Identity returns the Identity of the object.
func (o *IKEGatewayProfile) Identity() bambou.Identity {

	return IKEGatewayProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKEGatewayProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKEGatewayProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKEGatewayProfile from the server
func (o *IKEGatewayProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKEGatewayProfile into the server
func (o *IKEGatewayProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKEGatewayProfile from the server
func (o *IKEGatewayProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKEGatewayProfile
func (o *IKEGatewayProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKEGatewayProfile
func (o *IKEGatewayProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKEGatewayProfile
func (o *IKEGatewayProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKEGatewayProfile
func (o *IKEGatewayProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
