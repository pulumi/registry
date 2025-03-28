---
title: Top-level resource versions.
meta_desc: Table to help migrate between top-level resources for Azure Native v1 and v2
---

## Top-Level Resource Versions

|Service|Resource|REST version in v1|REST version in v2|
|---|---|---|---|
|Aad|DomainService|2021-03-01|2022-12-01|
|Aad|OuContainer|2021-03-01|2022-12-01|
|AadIam|azureADMetric|2020-07-01-preview|2020-07-01-preview|
|AadIam|DiagnosticSetting|2017-04-01|2017-04-01|
|AadIam|PrivateEndpointConnection|2020-03-01|2020-03-01|
|AadIam|privateLinkForAzureAd|2020-03-01|2020-03-01|
|Addons|listSupportPlanTypeInfo|2018-03-01|2018-03-01|
|Addons|SupportPlanType|2018-03-01|2018-03-01|
|Advisor|Suppression|2020-01-01|2023-01-01|
|AgFoodPlatform|DataConnector|not present|2023-06-01-preview|
|AgFoodPlatform|DataManagerForAgricultureResource|not present|2023-06-01-preview|
|AgFoodPlatform|Extension|2020-05-12-preview|2023-06-01-preview|
|AgFoodPlatform|FarmBeatsModel|2020-05-12-preview|Removed from default due to deprecation|
|AgFoodPlatform|PrivateEndpointConnection|not present|2023-06-01-preview|
|AgFoodPlatform|Solution|not present|2023-06-01-preview|
|AlertsManagement|ActionRuleByName|2019-05-05-preview|2019-05-05-preview|
|AlertsManagement|AlertProcessingRuleByName|not present|2023-05-01-preview|
|AlertsManagement|PrometheusRuleGroup|not present|2023-03-01|
|AlertsManagement|SmartDetectorAlertRule|2019-06-01|2021-04-01|
|AnalysisServices|listServerGatewayStatus|2017-08-01|2017-08-01|
|AnalysisServices|ServerDetails|2017-08-01|2017-08-01|
|ApiCenter|Service|not present|2023-07-01-preview|
|ApiManagement|Api|2020-12-01|2022-08-01|
|ApiManagement|ApiDiagnostic|2020-12-01|2022-08-01|
|ApiManagement|ApiDiagnosticLogger|2018-01-01|Removed from more recent versions|
|ApiManagement|ApiIssue|2020-12-01|2022-08-01|
|ApiManagement|ApiIssueAttachment|2020-12-01|2022-08-01|
|ApiManagement|ApiIssueComment|2020-12-01|2022-08-01|
|ApiManagement|ApiManagementService|2020-12-01|2022-08-01|
|ApiManagement|ApiOperation|2020-12-01|2022-08-01|
|ApiManagement|ApiOperationPolicy|2020-12-01|2022-08-01|
|ApiManagement|ApiOperationsPolicy|2016-10-10|Renamed to ApiOperationPolicy|
|ApiManagement|ApiPolicy|2020-12-01|2022-08-01|
|ApiManagement|ApiRelease|2020-12-01|2022-08-01|
|ApiManagement|ApiSchema|2020-12-01|2022-08-01|
|ApiManagement|ApiTagDescription|2020-12-01|2022-08-01|
|ApiManagement|ApiVersionSet|2020-12-01|2022-08-01|
|ApiManagement|ApiWiki|not present|2022-08-01|
|ApiManagement|Authorization|not present|2022-08-01|
|ApiManagement|AuthorizationAccessPolicy|not present|2022-08-01|
|ApiManagement|AuthorizationProvider|not present|2022-08-01|
|ApiManagement|AuthorizationServer|2020-12-01|2022-08-01|
|ApiManagement|Backend|2020-12-01|2022-08-01|
|ApiManagement|Cache|2020-12-01|2022-08-01|
|ApiManagement|Certificate|2020-12-01|2022-08-01|
|ApiManagement|ContentItem|2020-12-01|2022-08-01|
|ApiManagement|ContentType|2020-12-01|2022-08-01|
|ApiManagement|Diagnostic|2020-12-01|2022-08-01|
|ApiManagement|DiagnosticLogger|2018-01-01|Removed from more recent versions|
|ApiManagement|Documentation|not present|2022-08-01|
|ApiManagement|EmailTemplate|2020-12-01|2022-08-01|
|ApiManagement|Gateway|2020-12-01|2022-08-01|
|ApiManagement|GatewayApiEntityTag|2020-12-01|2022-08-01|
|ApiManagement|GatewayCertificateAuthority|2020-12-01|2022-08-01|
|ApiManagement|GatewayHostnameConfiguration|2020-12-01|2022-08-01|
|ApiManagement|getApiManagementServiceDomainOwnershipIdentifier|2020-12-01|2022-08-01|
|ApiManagement|getApiManagementServiceSsoToken|2020-12-01|2022-08-01|
|ApiManagement|getAuthorizationLoginLinkPost|not present|2022-08-01|
|ApiManagement|getUserSharedAccessToken|2020-12-01|2022-08-01|
|ApiManagement|GlobalSchema|not present|2022-08-01|
|ApiManagement|GraphQLApiResolver|not present|2022-08-01|
|ApiManagement|GraphQLApiResolverPolicy|not present|2022-08-01|
|ApiManagement|Group|2020-12-01|2022-08-01|
|ApiManagement|GroupUser|2020-12-01|2022-08-01|
|ApiManagement|IdentityProvider|2020-12-01|2022-08-01|
|ApiManagement|listAuthorizationServerSecrets|2020-12-01|2022-08-01|
|ApiManagement|listDelegationSettingSecrets|2020-12-01|2021-08-01|
|ApiManagement|listGatewayDebugCredentials|not present|2023-03-01-preview|
|ApiManagement|listGatewayKeys|2020-12-01|2022-08-01|
|ApiManagement|listGatewayTrace|not present|2023-03-01-preview|
|ApiManagement|listIdentityProviderSecrets|2020-12-01|2022-08-01|
|ApiManagement|listNamedValue|2020-12-01|2022-08-01|
|ApiManagement|listOpenIdConnectProviderSecrets|2020-12-01|2022-08-01|
|ApiManagement|listPolicyFragmentReferences|2021-12-01-preview|2022-08-01|
|ApiManagement|listSubscriptionSecrets|2020-12-01|2022-08-01|
|ApiManagement|listTenantAccessGitSecrets|2019-12-01|Replaced by listTenantAccessSecrets|
|ApiManagement|listTenantAccessSecrets|2020-12-01|2022-08-01|
|ApiManagement|listWorkspaceNamedValue|not present|2022-09-01-preview|
|ApiManagement|listWorkspacePolicyFragmentReferences|not present|2022-09-01-preview|
|ApiManagement|listWorkspaceSubscriptionSecrets|not present|2022-09-01-preview|
|ApiManagement|Logger|2020-12-01|2022-08-01|
|ApiManagement|NamedValue|2020-12-01|2022-08-01|
|ApiManagement|NotificationRecipientEmail|2020-12-01|2022-08-01|
|ApiManagement|NotificationRecipientUser|2020-12-01|2022-08-01|
|ApiManagement|OpenIdConnectProvider|2020-12-01|2022-08-01|
|ApiManagement|Policy|2020-12-01|2022-08-01|
|ApiManagement|PolicyFragment|2021-12-01-preview|2022-08-01|
|ApiManagement|PrivateEndpointConnectionByName|2021-04-01-preview|2022-08-01|
|ApiManagement|Product|2020-12-01|2022-08-01|
|ApiManagement|ProductApi|2020-12-01|2022-08-01|
|ApiManagement|ProductApiLink|not present|2022-09-01-preview|
|ApiManagement|ProductGroup|2020-12-01|2022-08-01|
|ApiManagement|ProductGroupLink|not present|2022-09-01-preview|
|ApiManagement|ProductPolicy|2020-12-01|2022-08-01|
|ApiManagement|ProductWiki|not present|2022-08-01|
|ApiManagement|Property|2019-01-01|Replaced by NamedValue|
|ApiManagement|Schema|2021-04-01-preview|2021-04-01-preview|
|ApiManagement|Subscription|2020-12-01|2022-08-01|
|ApiManagement|Tag|2020-12-01|2022-08-01|
|ApiManagement|TagApiLink|not present|2022-09-01-preview|
|ApiManagement|TagByApi|2020-12-01|2022-08-01|
|ApiManagement|TagByOperation|2020-12-01|2022-08-01|
|ApiManagement|TagByProduct|2020-12-01|2022-08-01|
|ApiManagement|TagOperationLink|not present|2022-09-01-preview|
|ApiManagement|TagProductLink|not present|2022-09-01-preview|
|ApiManagement|User|2020-12-01|2022-08-01|
|ApiManagement|Workspace|not present|2022-09-01-preview|
|ApiManagement|WorkspaceApi|not present|2022-09-01-preview|
|ApiManagement|WorkspaceApiOperation|not present|2022-09-01-preview|
|ApiManagement|WorkspaceApiOperationPolicy|not present|2022-09-01-preview|
|ApiManagement|WorkspaceApiPolicy|not present|2022-09-01-preview|
|ApiManagement|WorkspaceApiRelease|not present|2022-09-01-preview|
|ApiManagement|WorkspaceApiSchema|not present|2022-09-01-preview|
|ApiManagement|WorkspaceApiVersionSet|not present|2022-09-01-preview|
|ApiManagement|WorkspaceGlobalSchema|not present|2022-09-01-preview|
|ApiManagement|WorkspaceGroup|not present|2022-09-01-preview|
|ApiManagement|WorkspaceGroupUser|not present|2022-09-01-preview|
|ApiManagement|WorkspaceNamedValue|not present|2022-09-01-preview|
|ApiManagement|WorkspaceNotificationRecipientEmail|not present|2022-09-01-preview|
|ApiManagement|WorkspaceNotificationRecipientUser|not present|2022-09-01-preview|
|ApiManagement|WorkspacePolicy|not present|2022-09-01-preview|
|ApiManagement|WorkspacePolicyFragment|not present|2022-09-01-preview|
|ApiManagement|WorkspaceProduct|not present|2022-09-01-preview|
|ApiManagement|WorkspaceProductApiLink|not present|2022-09-01-preview|
|ApiManagement|WorkspaceProductGroupLink|not present|2022-09-01-preview|
|ApiManagement|WorkspaceProductPolicy|not present|2022-09-01-preview|
|ApiManagement|WorkspaceSubscription|not present|2022-09-01-preview|
|ApiManagement|WorkspaceTag|not present|2022-09-01-preview|
|ApiManagement|WorkspaceTagApiLink|not present|2022-09-01-preview|
|ApiManagement|WorkspaceTagOperationLink|not present|2022-09-01-preview|
|ApiManagement|WorkspaceTagProductLink|not present|2022-09-01-preview|
|App|Certificate|2022-03-01|2022-10-01|
|App|ConnectedEnvironment|not present|2022-10-01|
|App|ConnectedEnvironmentsCertificate|not present|2022-10-01|
|App|ConnectedEnvironmentsDaprComponent|not present|2022-10-01|
|App|ConnectedEnvironmentsStorage|not present|2022-10-01|
|App|ContainerApp|2022-03-01|2022-10-01|
|App|ContainerAppsAuthConfig|2022-03-01|2022-10-01|
|App|ContainerAppsSourceControl|2022-03-01|2022-10-01|
|App|DaprComponent|2022-03-01|2022-10-01|
|App|getContainerAppAuthToken|not present|2022-10-01|
|App|getManagedEnvironmentAuthToken|not present|2022-10-01|
|App|Job|not present|2023-04-01-preview|
|App|listConnectedEnvironmentsDaprComponentSecrets|not present|2022-10-01|
|App|listContainerAppCustomHostNameAnalysis|2022-03-01|2022-10-01|
|App|listContainerAppSecrets|2022-03-01|2022-10-01|
|App|listDaprComponentSecrets|2022-03-01|2022-10-01|
|App|listJobSecrets|not present|2023-04-01-preview|
|App|ManagedCertificate|not present|2023-04-01-preview|
|App|ManagedEnvironment|2022-03-01|2022-10-01|
|App|ManagedEnvironmentsStorage|2022-03-01|2022-10-01|
|AppComplianceAutomation|Report|2022-11-16-preview|2022-11-16-preview|
|AppConfiguration|ConfigurationStore|2020-06-01|2023-03-01|
|AppConfiguration|KeyValue|2020-07-01-preview|2023-03-01|
|AppConfiguration|listConfigurationStoreKeys|2020-06-01|2023-03-01|
|AppConfiguration|listConfigurationStoreKeyValue|2020-06-01|Renamed to listConfigurationStoreKeys|
|AppConfiguration|PrivateEndpointConnection|2020-06-01|2023-03-01|
|AppConfiguration|Replica|not present|2023-03-01|
|AppPlatform|ApiPortal|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|ApiPortalCustomDomain|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|Apm|not present|2023-05-01-preview|
|AppPlatform|App|2020-07-01|2023-05-01-preview|
|AppPlatform|ApplicationAccelerator|not present|2023-05-01-preview|
|AppPlatform|ApplicationLiveView|not present|2023-05-01-preview|
|AppPlatform|Binding|2020-07-01|2023-05-01-preview|
|AppPlatform|BuildpackBinding|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|BuildServiceAgentPool|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|BuildServiceBuild|not present|2023-05-01-preview|
|AppPlatform|BuildServiceBuilder|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|Certificate|2020-07-01|2023-05-01-preview|
|AppPlatform|ConfigServer|2020-07-01|2023-05-01-preview|
|AppPlatform|ConfigurationService|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|ContainerRegistry|not present|2023-05-01-preview|
|AppPlatform|CustomDomain|2020-07-01|2023-05-01-preview|
|AppPlatform|CustomizedAccelerator|not present|2023-05-01-preview|
|AppPlatform|Deployment|2020-07-01|2023-05-01-preview|
|AppPlatform|DevToolPortal|not present|2023-05-01-preview|
|AppPlatform|Gateway|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|GatewayCustomDomain|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|GatewayRouteConfig|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|getAppResourceUploadUrl|2020-07-01|2023-05-01-preview|
|AppPlatform|getBuildServiceBuildResultLog|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|getBuildServiceResourceUploadUrl|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|getDeploymentLogFileUrl|2020-07-01|2023-05-01-preview|
|AppPlatform|getDeploymentRemoteDebuggingConfig|not present|2023-05-01-preview|
|AppPlatform|listApmSecretKeys|not present|2023-05-01-preview|
|AppPlatform|listBuildServiceBuilderDeployments|not present|2023-05-01-preview|
|AppPlatform|listGatewayEnvSecrets|not present|2023-05-01-preview|
|AppPlatform|listServiceGloballyEnabledApms|not present|2023-05-01-preview|
|AppPlatform|listServiceTestKeys|2020-07-01|2023-05-01-preview|
|AppPlatform|MonitoringSetting|2020-07-01|2023-05-01-preview|
|AppPlatform|Service|2020-07-01|2023-05-01-preview|
|AppPlatform|ServiceRegistry|2022-01-01-preview|2023-05-01-preview|
|AppPlatform|Storage|2021-09-01-preview|2023-05-01-preview|
|Attestation|AttestationProvider|2020-10-01|2021-06-01|
|Attestation|PrivateEndpointConnection|2020-10-01|2021-06-01|
|Authorization|AccessReviewHistoryDefinitionById|2021-11-16-preview|2021-12-01-preview|
|Authorization|AccessReviewScheduleDefinitionById|2021-03-01-preview|2021-12-01-preview|
|Authorization|ManagementLockAtResourceGroupLevel|2017-04-01|2020-05-01|
|Authorization|ManagementLockAtResourceLevel|2017-04-01|2020-05-01|
|Authorization|ManagementLockAtSubscriptionLevel|2017-04-01|2020-05-01|
|Authorization|ManagementLockByScope|2017-04-01|2020-05-01|
|Authorization|PolicyAssignment|2020-09-01|2022-06-01|
|Authorization|PolicyDefinition|2020-09-01|2021-06-01|
|Authorization|PolicyDefinitionAtManagementGroup|2020-09-01|2021-06-01|
|Authorization|PolicyExemption|2020-07-01-preview|2022-07-01-preview|
|Authorization|PolicySetDefinition|2020-09-01|2021-06-01|
|Authorization|PolicySetDefinitionAtManagementGroup|2020-09-01|2021-06-01|
|Authorization|PrivateLinkAssociation|2020-05-01|2020-05-01|
|Authorization|ResourceManagementPrivateLink|2020-05-01|2020-05-01|
|Authorization|RoleAssignment|2020-10-01-preview|2022-04-01|
|Authorization|RoleDefinition|2018-01-01-preview|2022-05-01-preview|
|Authorization|RoleManagementPolicyAssignment|2020-10-01|2020-10-01|
|Authorization|ScopeAccessReviewHistoryDefinitionById|2021-12-01-preview|2021-12-01-preview|
|Authorization|ScopeAccessReviewScheduleDefinitionById|2021-12-01-preview|2021-12-01-preview|
|Authorization|Variable|not present|2022-08-01-preview|
|Authorization|VariableAtManagementGroup|not present|2022-08-01-preview|
|Authorization|VariableValue|not present|2022-08-01-preview|
|Authorization|VariableValueAtManagementGroup|not present|2022-08-01-preview|
|Automanage|Account|2020-06-30-preview|2020-06-30-preview|
|Automanage|ConfigurationProfile|not present|2022-05-04|
|Automanage|ConfigurationProfileAssignment|2020-06-30-preview|2022-05-04|
|Automanage|ConfigurationProfileHCIAssignment|not present|2022-05-04|
|Automanage|ConfigurationProfileHCRPAssignment|not present|2022-05-04|
|Automanage|ConfigurationProfilePreference|2020-06-30-preview|2020-06-30-preview|
|Automanage|ConfigurationProfilesVersion|not present|2022-05-04|
|Automation|AutomationAccount|2021-06-22|2022-08-08|
|Automation|Certificate|2019-06-01|2022-08-08|
|Automation|Connection|2019-06-01|2022-08-08|
|Automation|ConnectionType|2019-06-01|2022-08-08|
|Automation|Credential|2019-06-01|2022-08-08|
|Automation|DscConfiguration|2019-06-01|2022-08-08|
|Automation|DscNodeConfiguration|2019-06-01|2022-08-08|
|Automation|HybridRunbookWorker|2021-06-22|2022-08-08|
|Automation|HybridRunbookWorkerGroup|2021-06-22|2022-08-08|
|Automation|JobSchedule|2019-06-01|2022-08-08|
|Automation|listKeyByAutomationAccount|2021-06-22|2022-08-08|
|Automation|Module|2019-06-01|2022-08-08|
|Automation|PrivateEndpointConnection|2020-01-13-preview|2020-01-13-preview|
|Automation|Python2Package|2019-06-01|2022-08-08|
|Automation|Python3Package|not present|2022-08-08|
|Automation|Runbook|2019-06-01|2022-08-08|
|Automation|Schedule|2019-06-01|2022-08-08|
|Automation|SoftwareUpdateConfigurationByName|2019-06-01|2019-06-01|
|Automation|SourceControl|2019-06-01|2022-08-08|
|Automation|Variable|2019-06-01|2022-08-08|
|Automation|Watcher|2019-06-01|2020-01-13-preview|
|Automation|Webhook|2015-10-31|2015-10-31|
|AutonomousDevelopmentPlatform|Account|2021-02-01-preview|2021-11-01-preview|
|AutonomousDevelopmentPlatform|DataPool|2021-02-01-preview|2021-11-01-preview|
|AVS|Addon|2020-07-17-preview|2022-05-01|
|AVS|Authorization|2020-03-20|2022-05-01|
|AVS|CloudLink|2021-06-01|2022-05-01|
|AVS|Cluster|2020-03-20|2022-05-01|
|AVS|Datastore|2021-01-01-preview|2022-05-01|
|AVS|getScriptExecutionLogs|2021-06-01|2022-05-01|
|AVS|GlobalReachConnection|2020-07-17-preview|2022-05-01|
|AVS|HcxEnterpriseSite|2020-03-20|2022-05-01|
|AVS|listClusterZones|not present|2022-05-01|
|AVS|listPrivateCloudAdminCredentials|2020-03-20|2022-05-01|
|AVS|PlacementPolicy|2021-12-01|2022-05-01|
|AVS|PrivateCloud|2020-03-20|2022-05-01|
|AVS|ScriptExecution|2021-06-01|2022-05-01|
|AVS|WorkloadNetworkDhcp|2020-07-17-preview|2022-05-01|
|AVS|WorkloadNetworkDnsService|2020-07-17-preview|2022-05-01|
|AVS|WorkloadNetworkDnsZone|2020-07-17-preview|2022-05-01|
|AVS|WorkloadNetworkPortMirroring|2020-07-17-preview|2022-05-01|
|AVS|WorkloadNetworkPublicIP|2021-06-01|2022-05-01|
|AVS|WorkloadNetworkSegment|2020-07-17-preview|2022-05-01|
|AVS|WorkloadNetworkVMGroup|2020-07-17-preview|2022-05-01|
|AzureActiveDirectory|B2CTenant|2019-01-01-preview|2021-04-01|
|AzureActiveDirectory|GuestUsage|2020-05-01-preview|2021-04-01|
|AzureArcData|ActiveDirectoryConnector|2022-03-01-preview|2023-01-15-preview|
|AzureArcData|DataController|2021-06-01-preview|2023-01-15-preview|
|AzureArcData|FailoverGroup|not present|2023-01-15-preview|
|AzureArcData|PostgresInstance|2021-06-01-preview|2023-01-15-preview|
|AzureArcData|SqlManagedInstance|2021-06-01-preview|2023-01-15-preview|
|AzureArcData|SqlServerDatabase|not present|2023-01-15-preview|
|AzureArcData|SqlServerInstance|2021-06-01-preview|2023-01-15-preview|
|AzureData|SqlServer|2019-07-24-preview|2019-07-24-preview|
|AzureData|SqlServerRegistration|2019-07-24-preview|2019-07-24-preview|
|AzureSphere|Catalog|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|Deployment|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|Device|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|DeviceGroup|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|Image|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|listCatalogDeployments|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|listCatalogDeviceGroups|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|listCatalogDeviceInsights|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|listCatalogDevices|2022-09-01-preview|2022-09-01-preview|
|AzureSphere|Product|2022-09-01-preview|2022-09-01-preview|
|AzureStack|CustomerSubscription|2017-06-01|2022-06-01|
|AzureStack|getProduct|2017-06-01|2022-06-01|
|AzureStack|getProducts|2017-06-01|2022-06-01|
|AzureStack|getRegistrationActivationKey|2017-06-01|2022-06-01|
|AzureStack|LinkedSubscription|2020-06-01-preview|2020-06-01-preview|
|AzureStack|listProductDetails|2017-06-01|2022-06-01|
|AzureStack|listProducts|2017-06-01|2022-06-01|
|AzureStack|Registration|2017-06-01|2022-06-01|
|AzureStackHCI|ArcSetting|2021-01-01-preview|2023-03-01|
|AzureStackHCI|Cluster|2020-10-01|2023-03-01|
|AzureStackHCI|Extension|2021-01-01-preview|2023-03-01|
|AzureStackHCI|GalleryImage|not present|2022-12-15-preview|
|AzureStackHCI|GuestAgent|not present|2022-12-15-preview|
|AzureStackHCI|HybridIdentityMetadatum|not present|2022-12-15-preview|
|AzureStackHCI|MachineExtension|not present|2022-12-15-preview|
|AzureStackHCI|MarketplaceGalleryImage|not present|2022-12-15-preview|
|AzureStackHCI|NetworkInterface|not present|2022-12-15-preview|
|AzureStackHCI|StorageContainer|not present|2022-12-15-preview|
|AzureStackHCI|Update|not present|2023-03-01|
|AzureStackHCI|UpdateRun|not present|2023-03-01|
|AzureStackHCI|UpdateSummary|not present|2023-03-01|
|AzureStackHCI|VirtualHardDisk|not present|2022-12-15-preview|
|AzureStackHCI|VirtualMachine|not present|2022-12-15-preview|
|AzureStackHCI|VirtualNetwork|not present|2022-12-15-preview|
|BareMetalInfrastructure|AzureBareMetalStorageInstance|not present|2023-04-06|
|Batch|Application|2021-01-01|2023-05-01|
|Batch|ApplicationPackage|2021-01-01|2023-05-01|
|Batch|BatchAccount|2021-01-01|2023-05-01|
|Batch|Certificate|2021-01-01|Deprecated and removed by Azure|
|Batch|listBatchAccountKeys|2021-01-01|2023-05-01|
|Batch|Pool|2021-01-01|2023-05-01|
|Billing|BillingRoleAssignmentByBillingAccount|2019-10-01-preview|2019-10-01-preview|
|Billing|BillingRoleAssignmentByDepartment|2019-10-01-preview|2019-10-01-preview|
|Billing|BillingRoleAssignmentByEnrollmentAccount|2019-10-01-preview|2019-10-01-preview|
|Billing|listBillingAccountInvoiceSectionsByCreateSubscriptionPermission|2020-05-01|2020-05-01|
|Blockchain|BlockchainMember|2018-06-01-preview|2018-06-01-preview|
|Blockchain|listBlockchainMemberApiKeys|2018-06-01-preview|2018-06-01-preview|
|Blockchain|listLocationConsortiums|2018-06-01-preview|2018-06-01-preview|
|Blockchain|listTransactionNodeApiKeys|2018-06-01-preview|2018-06-01-preview|
|Blockchain|TransactionNode|2018-06-01-preview|2018-06-01-preview|
|Blueprint|Artifact|2018-11-01-preview|2018-11-01-preview|
|Blueprint|Assignment|2018-11-01-preview|2018-11-01-preview|
|Blueprint|Blueprint|2018-11-01-preview|2018-11-01-preview|
|Blueprint|PublishedBlueprint|2018-11-01-preview|2018-11-01-preview|
|BotService|Bot|2021-03-01|2022-09-15|
|BotService|BotConnection|2021-03-01|2022-09-15|
|BotService|Channel|2021-03-01|2022-09-15|
|BotService|EnterpriseChannel|2018-07-12|[Deprecated by Azure](https:||learn.microsoft.com|en-us|dotnet|api|microsoft.bot.connector.channels.enterprisechannel?view=botbuilder-dotnet-stable)|
|BotService|listBotConnectionServiceProviders|2021-03-01|2022-09-15|
|BotService|listBotConnectionWithSecrets|2021-03-01|2022-09-15|
|BotService|listChannelWithKeys|2021-03-01|2022-09-15|
|BotService|listQnAMakerEndpointKey|not present|2022-09-15|
|BotService|PrivateEndpointConnection|2021-05-01-preview|2022-09-15|
|Cache|AccessPolicy|not present|2023-05-01-preview|
|Cache|AccessPolicyAssignment|not present|2023-05-01-preview|
|Cache|Database|2021-03-01|2023-03-01-preview|
|Cache|EnterprisePrivateEndpointConnection|not present|2023-03-01-preview|
|Cache|FirewallRule|2020-06-01|2023-04-01|
|Cache|LinkedServer|2020-06-01|2023-04-01|
|Cache|listDatabaseKeys|2021-03-01|2023-03-01-preview|
|Cache|listRedisKeys|2020-06-01|2023-04-01|
|Cache|PatchSchedule|2020-06-01|2023-04-01|
|Cache|PrivateEndpointConnection|2021-03-01|2023-04-01|
|Cache|Redis|2020-06-01|2023-04-01|
|Cache|RedisEnterprise|2021-03-01|2023-03-01-preview|
|Cdn|AFDCustomDomain|2020-09-01|2023-05-01|
|Cdn|AFDEndpoint|2020-09-01|2023-05-01|
|Cdn|AFDOrigin|2020-09-01|2023-05-01|
|Cdn|AFDOriginGroup|2020-09-01|2023-05-01|
|Cdn|CustomDomain|2020-09-01|2023-05-01|
|Cdn|Endpoint|2020-09-01|2023-05-01|
|Cdn|getProfileSupportedOptimizationTypes|2020-09-01|2023-05-01|
|Cdn|Origin|2020-09-01|2023-05-01|
|Cdn|OriginGroup|2020-09-01|2023-05-01|
|Cdn|Policy|2020-09-01|2023-05-01|
|Cdn|Profile|2020-09-01|2023-05-01|
|Cdn|Route|2020-09-01|2023-05-01|
|Cdn|Rule|2020-09-01|2023-05-01|
|Cdn|RuleSet|2020-09-01|2023-05-01|
|Cdn|Secret|2020-09-01|2023-05-01|
|Cdn|SecurityPolicy|2020-09-01|2023-05-01|
|CertificateRegistration|AppServiceCertificateOrder|2020-10-01|2022-09-01|
|CertificateRegistration|AppServiceCertificateOrderCertificate|2020-10-01|2022-09-01|
|ChangeAnalysis|ConfigurationProfile|2020-04-01-preview|2020-04-01-preview|
|Chaos|Capability|2021-09-15-preview|2023-04-15-preview|
|Chaos|Experiment|2021-09-15-preview|2023-04-15-preview|
|Chaos|Target|2021-09-15-preview|2023-04-15-preview|
|CognitiveServices|Account|2017-04-18|2023-05-01|
|CognitiveServices|CommitmentPlan|2021-10-01|2023-05-01|
|CognitiveServices|CommitmentPlanAssociation|not present|2023-05-01|
|CognitiveServices|Deployment|2021-10-01|2023-05-01|
|CognitiveServices|listAccountKeys|2017-04-18|2023-05-01|
|CognitiveServices|PrivateEndpointConnection|2017-04-18|2023-05-01|
|CognitiveServices|SharedCommitmentPlan|not present|2023-05-01|
|Communication|CommunicationService|2020-08-20|2023-03-31|
|Communication|Domain|2021-10-01-preview|2023-03-31|
|Communication|EmailService|2021-10-01-preview|2023-03-31|
|Communication|listCommunicationServiceKeys|2020-08-20|2023-03-31|
|Communication|listEmailServiceVerifiedExchangeOnlineDomains|2021-10-01-preview|2023-03-31|
|Communication|SenderUsername|not present|2023-03-31|
|Compute|AvailabilitySet|2020-12-01|2023-03-01|
|Compute|CapacityReservation|2021-04-01|2023-03-01|
|Compute|CapacityReservationGroup|2021-04-01|2023-03-01|
|Compute|CloudService|2021-03-01|2022-09-04|
|Compute|DedicatedHost|2020-12-01|2023-03-01|
|Compute|DedicatedHostGroup|2020-12-01|2023-03-01|
|Compute|Disk|2020-12-01|2022-07-02|
|Compute|DiskAccess|2020-12-01|2022-07-02|
|Compute|DiskAccessAPrivateEndpointConnection|2020-12-01|2022-07-02|
|Compute|DiskEncryptionSet|2020-12-01|2022-07-02|
|Compute|Gallery|2020-09-30|2022-03-03|
|Compute|GalleryApplication|2020-09-30|2022-03-03|
|Compute|GalleryApplicationVersion|2020-09-30|2022-03-03|
|Compute|GalleryImage|2020-09-30|2022-03-03|
|Compute|GalleryImageVersion|2020-09-30|2022-03-03|
|Compute|getLogAnalyticExportRequestRateByInterval|2020-12-01|2023-03-01|
|Compute|getLogAnalyticExportThrottledRequests|2020-12-01|2023-03-01|
|Compute|Image|2020-12-01|2023-03-01|
|Compute|ProximityPlacementGroup|2020-12-01|2023-03-01|
|Compute|RestorePoint|2021-03-01|2023-03-01|
|Compute|RestorePointCollection|2021-03-01|2023-03-01|
|Compute|Snapshot|2020-12-01|2022-07-02|
|Compute|SshPublicKey|2020-12-01|2023-03-01|
|Compute|VirtualMachine|2021-03-01|2023-03-01|
|Compute|VirtualMachineExtension|2021-03-01|2023-03-01|
|Compute|VirtualMachineRunCommandByVirtualMachine|2021-03-01|2023-03-01|
|Compute|VirtualMachineScaleSet|2021-03-01|2023-03-01|
|Compute|VirtualMachineScaleSetExtension|2021-03-01|2023-03-01|
|Compute|VirtualMachineScaleSetVM|2021-03-01|2023-03-01|
|Compute|VirtualMachineScaleSetVMExtension|2021-03-01|2023-03-01|
|Compute|VirtualMachineScaleSetVMRunCommand|2021-03-01|2023-03-01|
|ConfidentialLedger|Ledger|2020-12-01-preview|2022-05-13|
|ConfidentialLedger|ManagedCCF|not present|2023-01-26-preview|
|Confluent|Organization|2020-03-01|2021-12-01|
|ConnectedVMwarevSphere|Cluster|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|Datastore|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|GuestAgent|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|Host|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|HybridIdentityMetadatum|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|InventoryItem|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|MachineExtension|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|ResourcePool|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|VCenter|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|VirtualMachine|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|VirtualMachineTemplate|2020-10-01-preview|2022-07-15-preview|
|ConnectedVMwarevSphere|VirtualNetwork|2020-10-01-preview|2022-07-15-preview|
|Consumption|Budget|2019-10-01|2023-05-01|
|ContainerInstance|ContainerGroup|2021-03-01|2023-05-01|
|ContainerRegistry|AgentPool|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|BuildStep|not present|2018-02-01-preview|
|ContainerRegistry|CacheRule|not present|2023-01-01-preview|
|ContainerRegistry|ConnectedRegistry|2020-11-01-preview|2023-01-01-preview|
|ContainerRegistry|CredentialSet|not present|2023-01-01-preview|
|ContainerRegistry|ExportPipeline|2020-11-01-preview|2023-01-01-preview|
|ContainerRegistry|getBuildLogLink|2018-02-01-preview|2018-02-01-preview|
|ContainerRegistry|getRegistryBuildSourceUploadUrl|2018-02-01-preview|2018-02-01-preview|
|ContainerRegistry|getRegistryCredentials|2016-06-27-preview|2016-06-27-preview|
|ContainerRegistry|getWebhookCallbackConfig|2019-05-01|2022-12-01|
|ContainerRegistry|ImportPipeline|2020-11-01-preview|2023-01-01-preview|
|ContainerRegistry|listAgentPoolQueueStatus|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|listBuildStepBuildArguments|2018-02-01-preview|2018-02-01-preview|
|ContainerRegistry|listBuildTaskSourceRepositoryProperties|2018-02-01-preview|2018-02-01-preview|
|ContainerRegistry|listRegistryBuildSourceUploadUrl|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|listRegistryCredentials|2019-05-01|2022-12-01|
|ContainerRegistry|listRunLogSasUrl|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|listTaskDetails|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|listTaskRunDetails|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|listWebhookEvents|2019-05-01|2022-12-01|
|ContainerRegistry|PipelineRun|2020-11-01-preview|2023-01-01-preview|
|ContainerRegistry|PrivateEndpointConnection|2020-11-01-preview|2022-12-01|
|ContainerRegistry|Registry|2019-05-01|2022-12-01|
|ContainerRegistry|Replication|2019-05-01|2022-12-01|
|ContainerRegistry|ScopeMap|2020-11-01-preview|2022-12-01|
|ContainerRegistry|Task|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|TaskRun|2019-06-01-preview|2019-06-01-preview|
|ContainerRegistry|Token|2020-11-01-preview|2022-12-01|
|ContainerRegistry|Webhook|2019-05-01|2022-12-01|
|ContainerService|AgentPool|2021-03-01|2023-04-01|
|ContainerService|Fleet|not present|2023-03-15-preview|
|ContainerService|FleetMember|not present|2023-03-15-preview|
|ContainerService|listFleetCredentials|not present|2023-03-15-preview|
|ContainerService|listManagedClusterAccessProfile|2020-03-01|2020-03-01|
|ContainerService|listManagedClusterAdminCredentials|2021-03-01|2023-04-01|
|ContainerService|listManagedClusterMonitoringUserCredentials|2021-03-01|2023-04-01|
|ContainerService|listManagedClusterUserCredentials|2021-03-01|2023-04-01|
|ContainerService|MaintenanceConfiguration|2021-03-01|2023-04-01|
|ContainerService|ManagedCluster|2021-03-01|2023-04-01|
|ContainerService|ManagedClusterSnapshot|2022-02-02-preview|2023-05-02-preview|
|ContainerService|OpenShiftManagedCluster|2019-04-30|2019-10-27-preview|
|ContainerService|PrivateEndpointConnection|2021-03-01|2023-04-01|
|ContainerService|Snapshot|2021-08-01|2023-04-01|
|ContainerService|TrustedAccessRoleBinding|2022-04-02-preview|2023-05-02-preview|
|ContainerService|UpdateRun|not present|2023-03-15-preview|
|ContainerStorage|Pool|2023-03-01-preview|2023-03-01-preview|
|ContainerStorage|Volume|2023-03-01-preview|2023-03-01-preview|
|ContainerStorage|VolumeSnapshot|2023-03-01-preview|2023-03-01-preview|
|CostManagement|Budget|not present|2023-04-01-preview|
|CostManagement|CloudConnector|2019-03-01-preview|2019-03-01-preview|
|CostManagement|Connector|not present|2018-08-01-preview|
|CostManagement|CostAllocationRule|2020-03-01-preview|2020-03-01-preview|
|CostManagement|Export|2020-06-01|2023-03-01|
|CostManagement|MarkupRule|not present|2022-10-05-preview|
|CostManagement|Report|2018-08-01-preview|2018-08-01-preview|
|CostManagement|ReportByBillingAccount|2018-08-01-preview|2018-08-01-preview|
|CostManagement|ReportByDepartment|2018-08-01-preview|2018-08-01-preview|
|CostManagement|ReportByResourceGroupName|2018-08-01-preview|2018-08-01-preview|
|CostManagement|ScheduledAction|2022-04-01-preview|2023-03-01|
|CostManagement|ScheduledActionByScope|2022-04-01-preview|2023-03-01|
|CostManagement|Setting|2019-11-01|2019-11-01|
|CostManagement|SettingByScope|not present|2022-10-05-preview|
|CostManagement|View|2019-11-01|2023-03-01|
|CostManagement|ViewByScope|2019-11-01|2023-03-01|
|CustomerInsights|Connector|2017-04-26|2017-04-26|
|CustomerInsights|ConnectorMapping|2017-04-26|2017-04-26|
|CustomerInsights|getImageUploadUrlForData|2017-04-26|2017-04-26|
|CustomerInsights|getImageUploadUrlForEntityType|2017-04-26|2017-04-26|
|CustomerInsights|getPredictionModelStatus|2017-04-26|2017-04-26|
|CustomerInsights|getPredictionTrainingResults|2017-04-26|2017-04-26|
|CustomerInsights|getProfileEnrichingKpis|2017-04-26|2017-04-26|
|CustomerInsights|Hub|2017-04-26|2017-04-26|
|CustomerInsights|Kpi|2017-04-26|2017-04-26|
|CustomerInsights|Link|2017-04-26|2017-04-26|
|CustomerInsights|Prediction|2017-04-26|2017-04-26|
|CustomerInsights|Profile|2017-04-26|2017-04-26|
|CustomerInsights|Relationship|2017-04-26|2017-04-26|
|CustomerInsights|RelationshipLink|2017-04-26|2017-04-26|
|CustomerInsights|RoleAssignment|2017-04-26|2017-04-26|
|CustomerInsights|View|2017-04-26|2017-04-26|
|CustomProviders|Association|2018-09-01-preview|2018-09-01-preview|
|CustomProviders|CustomResourceProvider|2018-09-01-preview|2018-09-01-preview|
|Dashboard|Grafana|2022-05-01-preview|2022-08-01|
|Dashboard|PrivateEndpointConnection|2022-05-01-preview|2022-08-01|
|DataBox|Job|2020-11-01|2022-12-01|
|DataBox|listJobCredentials|2020-11-01|2022-12-01|
|DataBoxEdge|Addon|2020-12-01|2022-03-01|
|DataBoxEdge|BandwidthSchedule|2020-12-01|2022-03-01|
|DataBoxEdge|Container|2020-12-01|2022-03-01|
|DataBoxEdge|Device|2020-12-01|2022-03-01|
|DataBoxEdge|getDeviceExtendedInformation|2020-12-01|2022-03-01|
|DataBoxEdge|listOrderDCAccessCode|2020-12-01|2022-03-01|
|DataBoxEdge|MonitoringConfig|2020-12-01|2022-03-01|
|DataBoxEdge|Order|2020-12-01|2022-03-01|
|DataBoxEdge|Role|2020-12-01|2022-03-01|
|DataBoxEdge|Share|2020-12-01|2022-03-01|
|DataBoxEdge|StorageAccount|2020-12-01|2022-03-01|
|DataBoxEdge|StorageAccountCredential|2020-12-01|2022-03-01|
|DataBoxEdge|Trigger|2020-12-01|2022-03-01|
|DataBoxEdge|User|2020-12-01|2022-03-01|
|Databricks|AccessConnector|2022-04-01-preview|2023-05-01|
|Databricks|PrivateEndpointConnection|2022-04-01-preview|2023-02-01|
|Databricks|vNetPeering|2018-04-01|2023-02-01|
|Databricks|Workspace|2018-04-01|2023-02-01|
|DataCatalog|ADCCatalog|2016-03-30|2016-03-30|
|Datadog|getMonitorDefaultKey|2021-03-01|2022-06-01|
|Datadog|listMonitorApiKeys|2021-03-01|2022-06-01|
|Datadog|listMonitorHosts|2021-03-01|2022-06-01|
|Datadog|listMonitorLinkedResources|2021-03-01|2022-06-01|
|Datadog|listMonitorMonitoredResources|2021-03-01|2022-06-01|
|Datadog|Monitor|2021-03-01|2022-06-01|
|DataFactory|ChangeDataCapture|not present|2018-06-01|
|DataFactory|CredentialOperation|2018-06-01|2018-06-01|
|DataFactory|DataFlow|2018-06-01|2018-06-01|
|DataFactory|Dataset|2018-06-01|2018-06-01|
|DataFactory|Factory|2018-06-01|2018-06-01|
|DataFactory|getExposureControlFeatureValue|2018-06-01|2018-06-01|
|DataFactory|getExposureControlFeatureValueByFactory|2018-06-01|2018-06-01|
|DataFactory|getFactoryDataPlaneAccess|2018-06-01|2018-06-01|
|DataFactory|getFactoryGitHubAccessToken|2018-06-01|2018-06-01|
|DataFactory|getIntegrationRuntimeConnectionInfo|2018-06-01|2018-06-01|
|DataFactory|getIntegrationRuntimeObjectMetadatum|2018-06-01|2018-06-01|
|DataFactory|getIntegrationRuntimeStatus|2018-06-01|2018-06-01|
|DataFactory|getTriggerEventSubscriptionStatus|2018-06-01|2018-06-01|
|DataFactory|GlobalParameter|2018-06-01|2018-06-01|
|DataFactory|IntegrationRuntime|2018-06-01|2018-06-01|
|DataFactory|LinkedService|2018-06-01|2018-06-01|
|DataFactory|listIntegrationRuntimeAuthKeys|2018-06-01|2018-06-01|
|DataFactory|ManagedPrivateEndpoint|2018-06-01|2018-06-01|
|DataFactory|Pipeline|2018-06-01|2018-06-01|
|DataFactory|PrivateEndpointConnection|2018-06-01|2018-06-01|
|DataFactory|Trigger|2018-06-01|2018-06-01|
|DataLakeAnalytics|Account|2016-11-01|2019-11-01-preview|
|DataLakeAnalytics|ComputePolicy|2016-11-01|2019-11-01-preview|
|DataLakeAnalytics|DataLakeStoreAccount|2016-11-01|2019-11-01-preview|
|DataLakeAnalytics|FirewallRule|2016-11-01|2019-11-01-preview|
|DataLakeAnalytics|listStorageAccountSasTokens|2016-11-01|2019-11-01-preview|
|DataLakeAnalytics|StorageAccount|2016-11-01|2019-11-01-preview|
|DataLakeStore|Account|2016-11-01|2016-11-01|
|DataLakeStore|FirewallRule|2016-11-01|2016-11-01|
|DataLakeStore|TrustedIdProvider|2016-11-01|2016-11-01|
|DataLakeStore|VirtualNetworkRule|2016-11-01|2016-11-01|
|DataMigration|DatabaseMigrationsSqlDb|2022-03-30-preview|2022-03-30-preview|
|DataMigration|File|2018-07-15-preview|2021-06-30|
|DataMigration|listSqlMigrationServiceAuthKeys|2021-10-30-preview|2022-03-30-preview|
|DataMigration|listSqlMigrationServiceMonitoringData|2021-10-30-preview|2022-03-30-preview|
|DataMigration|Project|2018-04-19|2021-06-30|
|DataMigration|Service|2018-04-19|2021-06-30|
|DataMigration|ServiceTask|not present|2021-06-30|
|DataMigration|SqlMigrationService|2021-10-30-preview|2022-03-30-preview|
|DataMigration|Task|2018-04-19|2021-06-30|
|DataProtection|BackupInstance|2021-01-01|2023-01-01|
|DataProtection|BackupPolicy|2021-01-01|2023-01-01|
|DataProtection|BackupVault|2021-01-01|2023-01-01|
|DataProtection|DppResourceGuardProxy|not present|2023-01-01|
|DataProtection|ResourceGuard|2021-10-01-preview|2023-01-01|
|DataShare|Account|2020-09-01|2021-08-01|
|DataShare|DataSet|2020-09-01|2021-08-01|
|DataShare|DataSetMapping|2020-09-01|2021-08-01|
|DataShare|Invitation|2020-09-01|2021-08-01|
|DataShare|listShareSubscriptionSourceShareSynchronizationSettings|2020-09-01|2021-08-01|
|DataShare|listShareSubscriptionSynchronizationDetails|2020-09-01|2021-08-01|
|DataShare|listShareSubscriptionSynchronizations|2020-09-01|2021-08-01|
|DataShare|listShareSynchronizationDetails|2020-09-01|2021-08-01|
|DataShare|listShareSynchronizations|2020-09-01|2021-08-01|
|DataShare|Share|2020-09-01|2021-08-01|
|DataShare|ShareSubscription|2020-09-01|2021-08-01|
|DataShare|SynchronizationSetting|2020-09-01|2021-08-01|
|DataShare|Trigger|2020-09-01|2021-08-01|
|DBforMariaDB|Configuration|2018-06-01|2018-06-01|
|DBforMariaDB|Database|2018-06-01|2018-06-01|
|DBforMariaDB|FirewallRule|2018-06-01|2018-06-01|
|DBforMariaDB|PrivateEndpointConnection|2018-06-01|2018-06-01|
|DBforMariaDB|Server|2018-06-01|2018-06-01|
|DBforMariaDB|VirtualNetworkRule|2018-06-01|2018-06-01|
|DBforMySQL|AzureADAdministrator|not present|2022-01-01|
|DBforMySQL|Configuration|2017-12-01|2022-01-01|
|DBforMySQL|Database|2017-12-01|2022-01-01|
|DBforMySQL|FirewallRule|2017-12-01|2022-01-01|
|DBforMySQL|getGetPrivateDnsZoneSuffixExecute|2021-05-01-preview|2022-01-01|
|DBforMySQL|PrivateEndpointConnection|2018-06-01|2022-09-30-preview|
|DBforMySQL|Server|2017-12-01|2022-01-01|
|DBforMySQL|ServerAdministrator|2017-12-01|[This was used for Single Server only which has been replaced with Flexible Server](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server)|
|DBforMySQL|ServerKey|2020-01-01|[This was used for Single Server only which has been replaced with Flexible Server](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server)|
|DBforMySQL|VirtualNetworkRule|2017-12-01|[This was used for Single Server only which has been replaced with Flexible Server](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server)|
|DBforPostgreSQL|Administrator|not present|2022-12-01|
|DBforPostgreSQL|Cluster|not present|2022-11-08|
|DBforPostgreSQL|Configuration|2017-12-01|2022-12-01|
|DBforPostgreSQL|Database|2017-12-01|2022-12-01|
|DBforPostgreSQL|FirewallRule|2017-12-01|2022-12-01|
|DBforPostgreSQL|getGetPrivateDnsZoneSuffixExecute|2022-01-20-preview|2022-12-01|
|DBforPostgreSQL|Migration|not present|2023-03-01-preview|
|DBforPostgreSQL|PrivateEndpointConnection|2018-06-01|2022-11-08|
|DBforPostgreSQL|Role|not present|2022-11-08|
|DBforPostgreSQL|Server|2017-12-01|2022-12-01|
|DBforPostgreSQL|ServerAdministrator|2017-12-01|[This was used for Single Server only which has been replaced with Flexible Server](https://learn.microsoft.com/en-us/azure/postgresql/single-server/whats-happening-to-postgresql-single-server)|
|DBforPostgreSQL|ServerKey|2020-01-01|[This was used for Single Server only which has been replaced with Flexible Server](https://learn.microsoft.com/en-us/azure/postgresql/single-server/whats-happening-to-postgresql-single-server)|
|DBforPostgreSQL|ServerSecurityAlertPolicy|2017-12-01|[This was used for Single Server only which has been replaced with Flexible Server](https://learn.microsoft.com/en-us/azure/postgresql/single-server/whats-happening-to-postgresql-single-server)|
|DBforPostgreSQL|VirtualNetworkRule|2017-12-01|[This was used for Single Server only which has been replaced with Flexible Server](https://learn.microsoft.com/en-us/azure/postgresql/single-server/whats-happening-to-postgresql-single-server)|
|DelegatedNetwork|ControllerDetails|2021-03-15|2021-03-15|
|DelegatedNetwork|DelegatedSubnetServiceDetails|2021-03-15|2021-03-15|
|DelegatedNetwork|OrchestratorInstanceServiceDetails|2021-03-15|2021-03-15|
|DeploymentManager|ArtifactSource|2019-11-01-preview|2019-11-01-preview|
|DeploymentManager|Rollout|2019-11-01-preview|2019-11-01-preview|
|DeploymentManager|Service|2019-11-01-preview|2019-11-01-preview|
|DeploymentManager|ServiceTopology|2019-11-01-preview|2019-11-01-preview|
|DeploymentManager|ServiceUnit|2019-11-01-preview|2019-11-01-preview|
|DeploymentManager|Step|2019-11-01-preview|2019-11-01-preview|
|DesktopVirtualization|Application|2021-02-01-preview|2022-09-09|
|DesktopVirtualization|ApplicationGroup|2021-02-01-preview|2022-09-09|
|DesktopVirtualization|HostPool|2021-02-01-preview|2022-09-09|
|DesktopVirtualization|MSIXPackage|2021-02-01-preview|2022-09-09|
|DesktopVirtualization|PrivateEndpointConnectionByHostPool|2021-04-01-preview|2022-10-14-preview|
|DesktopVirtualization|PrivateEndpointConnectionByWorkspace|2021-04-01-preview|2022-10-14-preview|
|DesktopVirtualization|ScalingPlan|2021-02-01-preview|2022-09-09|
|DesktopVirtualization|ScalingPlanPooledSchedule|2022-04-01-preview|2022-09-09|
|DesktopVirtualization|Workspace|2021-02-01-preview|2022-09-09|
|DevCenter|AttachedNetworkByDevCenter|2022-09-01-preview|2023-04-01|
|DevCenter|Catalog|2022-09-01-preview|2023-04-01|
|DevCenter|DevBoxDefinition|2022-09-01-preview|2023-04-01|
|DevCenter|DevCenter|2022-09-01-preview|2023-04-01|
|DevCenter|EnvironmentType|2022-09-01-preview|2023-04-01|
|DevCenter|Gallery|2022-09-01-preview|2023-04-01|
|DevCenter|NetworkConnection|2022-09-01-preview|2023-04-01|
|DevCenter|Pool|2022-09-01-preview|2023-04-01|
|DevCenter|Project|2022-09-01-preview|2023-04-01|
|DevCenter|ProjectEnvironmentType|2022-09-01-preview|2023-04-01|
|DevCenter|Schedule|2022-09-01-preview|2023-04-01|
|DevHub|getGitHubOAuth|2022-04-01-preview|2022-10-11-preview|
|DevHub|Workflow|2022-04-01-preview|2022-10-11-preview|
|Devices|Certificate|2020-08-31|2022-11-15-preview|
|Devices|DpsCertificate|2020-03-01|2022-12-12|
|Devices|IotDpsResource|2020-03-01|2022-12-12|
|Devices|IotDpsResourcePrivateEndpointConnection|2020-03-01|2022-12-12|
|Devices|IotHubResource|2020-08-31|2022-11-15-preview|
|Devices|IotHubResourceEventHubConsumerGroup|2020-08-31|2022-11-15-preview|
|Devices|listIotDpsResourceKeys|2020-03-01|2022-12-12|
|Devices|listIotDpsResourceKeysForKeyName|2020-03-01|2022-12-12|
|Devices|listIotHubResourceKeys|2020-08-31|2022-11-15-preview|
|Devices|listIotHubResourceKeysForKeyName|2020-08-31|2022-11-15-preview|
|Devices|PrivateEndpointConnection|2020-08-31|2022-11-15-preview|
|DeviceUpdate|Account|2020-03-01-preview|2023-07-01|
|DeviceUpdate|Instance|2020-03-01-preview|2023-07-01|
|DeviceUpdate|PrivateEndpointConnection|2020-03-01-preview|2023-07-01|
|DeviceUpdate|PrivateEndpointConnectionProxy|2020-03-01-preview|2023-07-01|
|DevSpaces|Controller|2019-04-01|2019-04-01|
|DevSpaces|listControllerConnectionDetails|2019-04-01|2019-04-01|
|DevTestLab|ArtifactSource|2018-09-15|2018-09-15|
|DevTestLab|CustomImage|2018-09-15|2018-09-15|
|DevTestLab|Disk|2018-09-15|2018-09-15|
|DevTestLab|Environment|2018-09-15|2018-09-15|
|DevTestLab|Formula|2018-09-15|2018-09-15|
|DevTestLab|getVirtualMachineRdpFileContents|2018-09-15|2018-09-15|
|DevTestLab|GlobalSchedule|2018-09-15|2018-09-15|
|DevTestLab|Lab|2018-09-15|2018-09-15|
|DevTestLab|listLabVhds|2018-09-15|2018-09-15|
|DevTestLab|listScheduleApplicable|2018-09-15|2018-09-15|
|DevTestLab|listServiceFabricApplicableSchedules|2018-09-15|2018-09-15|
|DevTestLab|listVirtualMachineApplicableSchedules|2018-09-15|2018-09-15|
|DevTestLab|NotificationChannel|2018-09-15|2018-09-15|
|DevTestLab|Policy|2018-09-15|2018-09-15|
|DevTestLab|Schedule|2018-09-15|2018-09-15|
|DevTestLab|Secret|2018-09-15|2018-09-15|
|DevTestLab|ServiceFabric|2018-09-15|2018-09-15|
|DevTestLab|ServiceFabricSchedule|2018-09-15|2018-09-15|
|DevTestLab|ServiceRunner|2018-09-15|2018-09-15|
|DevTestLab|User|2018-09-15|2018-09-15|
|DevTestLab|VirtualMachine|2018-09-15|2018-09-15|
|DevTestLab|VirtualMachineSchedule|2018-09-15|2018-09-15|
|DevTestLab|VirtualNetwork|2018-09-15|2018-09-15|
|DigitalTwins|DigitalTwin|2020-12-01|2023-01-31|
|DigitalTwins|DigitalTwinsEndpoint|2020-12-01|2023-01-31|
|DigitalTwins|PrivateEndpointConnection|2020-12-01|2023-01-31|
|DigitalTwins|TimeSeriesDatabaseConnection|2021-06-30-preview|2023-01-31|
|DocumentDB|CassandraCluster|2021-03-01-preview|2023-04-15|
|DocumentDB|CassandraDataCenter|2021-03-01-preview|2023-04-15|
|DocumentDB|CassandraResourceCassandraKeyspace|2021-03-15|2023-04-15|
|DocumentDB|CassandraResourceCassandraTable|2021-03-15|2023-04-15|
|DocumentDB|CassandraResourceCassandraView|2021-07-01-preview|2023-03-15-preview|
|DocumentDB|DatabaseAccount|2021-03-15|2023-04-15|
|DocumentDB|GraphResourceGraph|2021-07-01-preview|2023-03-15-preview|
|DocumentDB|GremlinResourceGremlinDatabase|2021-03-15|2023-04-15|
|DocumentDB|GremlinResourceGremlinGraph|2021-03-15|2023-04-15|
|DocumentDB|listDatabaseAccountConnectionStrings|2021-03-15|2023-04-15|
|DocumentDB|listDatabaseAccountKeys|2021-03-15|2023-04-15|
|DocumentDB|listMongoClusterConnectionStrings|not present|2023-03-15-preview|
|DocumentDB|listNotebookWorkspaceConnectionInfo|2021-03-15|2023-04-15|
|DocumentDB|MongoCluster|not present|2023-03-15-preview|
|DocumentDB|MongoClusterFirewallRule|not present|2023-03-15-preview|
|DocumentDB|MongoDBResourceMongoDBCollection|2021-03-15|2023-04-15|
|DocumentDB|MongoDBResourceMongoDBDatabase|2021-03-15|2023-04-15|
|DocumentDB|MongoDBResourceMongoRoleDefinition|2021-10-15-preview|2023-04-15|
|DocumentDB|MongoDBResourceMongoUserDefinition|2021-10-15-preview|2023-04-15|
|DocumentDB|NotebookWorkspace|2021-03-15|2023-04-15|
|DocumentDB|PrivateEndpointConnection|2021-03-15|2023-04-15|
|DocumentDB|Service|2021-04-01-preview|2023-04-15|
|DocumentDB|SqlResourceSqlContainer|2021-03-15|2023-04-15|
|DocumentDB|SqlResourceSqlDatabase|2021-03-15|2023-04-15|
|DocumentDB|SqlResourceSqlRoleAssignment|2021-03-01-preview|2023-04-15|
|DocumentDB|SqlResourceSqlRoleDefinition|2021-03-01-preview|2023-04-15|
|DocumentDB|SqlResourceSqlStoredProcedure|2021-03-15|2023-04-15|
|DocumentDB|SqlResourceSqlTrigger|2021-03-15|2023-04-15|
|DocumentDB|SqlResourceSqlUserDefinedFunction|2021-03-15|2023-04-15|
|DocumentDB|TableResourceTable|2021-03-15|2023-04-15|
|DomainRegistration|Domain|2020-10-01|2022-09-01|
|DomainRegistration|DomainOwnershipIdentifier|2020-10-01|2022-09-01|
|DomainRegistration|listDomainRecommendations|2020-10-01|2022-09-01|
|DomainRegistration|listTopLevelDomainAgreements|2020-10-01|2022-09-01|
|Dynamics365Fraudprotection|InstanceDetails|2021-02-01-preview|2021-02-01-preview|
|Easm|LabelByWorkspace|2022-04-01-preview|2023-04-01-preview|
|Easm|Workspace|2022-04-01-preview|2023-04-01-preview|
|EdgeOrder|Address|not present|2022-05-01-preview|
|EdgeOrder|AddressByName|2021-12-01|2021-12-01|
|EdgeOrder|listConfigurations|2021-12-01|2021-12-01|
|EdgeOrder|listProductFamilies|2021-12-01|2021-12-01|
|EdgeOrder|listProductsAndConfigurationProductFamilies|not present|2022-05-01-preview|
|EdgeOrder|listProductsAndConfigurations|not present|2022-05-01-preview|
|EdgeOrder|OrderItem|not present|2022-05-01-preview|
|EdgeOrder|OrderItemByName|2021-12-01|2021-12-01|
|Education|Lab|2021-12-01-preview|2021-12-01-preview|
|Education|Student|2021-12-01-preview|2021-12-01-preview|
|Elastic|getOrganizationApiKey|not present|2023-06-01|
|Elastic|getOrganizationElasticToAzureSubscriptionMapping|not present|2023-06-15-preview|
|Elastic|listAllTrafficFilter|not present|2023-06-01|
|Elastic|listDeploymentInfo|2020-07-01|2023-06-01|
|Elastic|listlistAssociatedTrafficFilter|not present|2023-06-01|
|Elastic|listMonitoredResource|2020-07-01|2023-06-01|
|Elastic|listUpgradableVersionDetails|2021-10-01-preview|2023-06-01|
|Elastic|listVMHost|2020-07-01|2023-06-01|
|Elastic|Monitor|2020-07-01|2023-06-01|
|Elastic|TagRule|2020-07-01|2023-06-01|
|ElasticSan|ElasticSan|2021-11-20-preview|2021-11-20-preview|
|ElasticSan|Volume|2021-11-20-preview|2021-11-20-preview|
|ElasticSan|VolumeGroup|2021-11-20-preview|2021-11-20-preview|
|EngagementFabric|Account|2018-09-01-preview|2018-09-01-preview|
|EngagementFabric|Channel|2018-09-01-preview|2018-09-01-preview|
|EngagementFabric|listAccountChannelTypes|2018-09-01-preview|2018-09-01-preview|
|EngagementFabric|listAccountKeys|2018-09-01-preview|2018-09-01-preview|
|EnterpriseKnowledgeGraph|EnterpriseKnowledgeGraph|2018-12-03|2018-12-03|
|EventGrid|CaCertificate|not present|2023-06-01-preview|
|EventGrid|Channel|2021-10-15-preview|2022-06-15|
|EventGrid|Client|not present|2023-06-01-preview|
|EventGrid|ClientGroup|not present|2023-06-01-preview|
|EventGrid|Domain|2020-06-01|2022-06-15|
|EventGrid|DomainEventSubscription|2021-10-15-preview|2022-06-15|
|EventGrid|DomainTopic|2020-06-01|2022-06-15|
|EventGrid|DomainTopicEventSubscription|2021-10-15-preview|2022-06-15|
|EventGrid|EventChannel|2021-06-01-preview|Removed from defaults due to deprecation|
|EventGrid|EventSubscription|2020-06-01|2022-06-15|
|EventGrid|getChannelFullUrl|2021-10-15-preview|2022-06-15|
|EventGrid|getDomainEventSubscriptionDeliveryAttributes|2021-10-15-preview|2022-06-15|
|EventGrid|getDomainEventSubscriptionFullUrl|2021-10-15-preview|2022-06-15|
|EventGrid|getDomainTopicEventSubscriptionDeliveryAttributes|2021-10-15-preview|2022-06-15|
|EventGrid|getDomainTopicEventSubscriptionFullUrl|2021-10-15-preview|2022-06-15|
|EventGrid|getEventSubscriptionDeliveryAttributes|2021-06-01-preview|2022-06-15|
|EventGrid|getEventSubscriptionFullUrl|2020-06-01|2022-06-15|
|EventGrid|getPartnerTopicEventSubscriptionDeliveryAttributes|2021-06-01-preview|2022-06-15|
|EventGrid|getPartnerTopicEventSubscriptionFullUrl|2021-06-01-preview|2022-06-15|
|EventGrid|getSystemTopicEventSubscriptionDeliveryAttributes|2021-06-01-preview|2022-06-15|
|EventGrid|getSystemTopicEventSubscriptionFullUrl|2021-06-01-preview|2022-06-15|
|EventGrid|getTopicEventSubscriptionDeliveryAttributes|2021-10-15-preview|2022-06-15|
|EventGrid|getTopicEventSubscriptionFullUrl|2021-10-15-preview|2022-06-15|
|EventGrid|listDomainSharedAccessKeys|2020-06-01|2022-06-15|
|EventGrid|listNamespaceSharedAccessKeys|not present|2023-06-01-preview|
|EventGrid|listNamespaceTopicSharedAccessKeys|not present|2023-06-01-preview|
|EventGrid|listPartnerNamespaceSharedAccessKeys|2021-06-01-preview|2022-06-15|
|EventGrid|listTopicSharedAccessKeys|2020-06-01|2022-06-15|
|EventGrid|Namespace|not present|2023-06-01-preview|
|EventGrid|NamespaceTopic|not present|2023-06-01-preview|
|EventGrid|NamespaceTopicEventSubscription|not present|2023-06-01-preview|
|EventGrid|PartnerConfiguration|2021-10-15-preview|2022-06-15|
|EventGrid|PartnerDestination|2021-10-15-preview|2023-06-01-preview|
|EventGrid|PartnerNamespace|2021-06-01-preview|2022-06-15|
|EventGrid|PartnerRegistration|2021-06-01-preview|2022-06-15|
|EventGrid|PartnerTopic|2021-10-15-preview|2022-06-15|
|EventGrid|PartnerTopicEventSubscription|2020-04-01-preview|2022-06-15|
|EventGrid|PermissionBinding|not present|2023-06-01-preview|
|EventGrid|PrivateEndpointConnection|2020-06-01|2022-06-15|
|EventGrid|SystemTopic|2021-06-01-preview|2022-06-15|
|EventGrid|SystemTopicEventSubscription|2020-04-01-preview|2022-06-15|
|EventGrid|Topic|2020-06-01|2022-06-15|
|EventGrid|TopicEventSubscription|2021-10-15-preview|2022-06-15|
|EventGrid|TopicSpace|not present|2023-06-01-preview|
|EventHub|ApplicationGroup|2022-01-01-preview|2022-10-01-preview|
|EventHub|Cluster|2018-01-01-preview|2022-10-01-preview|
|EventHub|ConsumerGroup|2017-04-01|2022-10-01-preview|
|EventHub|DisasterRecoveryConfig|2017-04-01|2022-10-01-preview|
|EventHub|EventHub|2017-04-01|2022-10-01-preview|
|EventHub|EventHubAuthorizationRule|2017-04-01|2022-10-01-preview|
|EventHub|listDisasterRecoveryConfigKeys|2017-04-01|2022-10-01-preview|
|EventHub|listEventHubKeys|2017-04-01|2022-10-01-preview|
|EventHub|listNamespaceKeys|2017-04-01|2022-10-01-preview|
|EventHub|Namespace|2017-04-01|2022-10-01-preview|
|EventHub|NamespaceAuthorizationRule|2017-04-01|2022-10-01-preview|
|EventHub|NamespaceIpFilterRule|2018-01-01-preview|2018-01-01-preview|
|EventHub|NamespaceNetworkRuleSet|2017-04-01|2022-10-01-preview|
|EventHub|NamespaceVirtualNetworkRule|2018-01-01-preview|2018-01-01-preview|
|EventHub|PrivateEndpointConnection|2018-01-01-preview|2022-10-01-preview|
|EventHub|SchemaRegistry|2022-01-01-preview|2022-10-01-preview|
|ExtendedLocation|CustomLocation|2021-03-15-preview|2021-08-15|
|ExtendedLocation|ResourceSyncRule|2021-08-31-preview|2021-08-31-preview|
|Features|SubscriptionFeatureRegistration|2021-07-01|2021-07-01|
|FluidRelay|FluidRelayServer|2021-03-12-preview|2022-06-01|
|FluidRelay|getFluidRelayServerKeys|2021-03-12-preview|Replaced with `listFluidRelayServerKeys`|
|FluidRelay|listFluidRelayServerKeys|2022-04-21|2022-06-01|
|GraphServices|Account|2022-09-22-preview|2023-04-13|
|GuestConfiguration|GuestConfigurationAssignment|2020-06-25|2022-01-25|
|GuestConfiguration|GuestConfigurationAssignmentsVMSS|not present|2022-01-25|
|GuestConfiguration|GuestConfigurationConnectedVMwarevSphereAssignment|2020-06-25|2022-01-25|
|GuestConfiguration|GuestConfigurationHCRPAssignment|2020-06-25|2022-01-25|
|HanaOnAzure|HanaInstance|2017-11-03-preview|This was replaced by ProviderInstance|
|HanaOnAzure|ProviderInstance|2020-02-07-preview|2020-02-07-preview|
|HanaOnAzure|SapMonitor|2020-02-07-preview|2020-02-07-preview|
|HardwareSecurityModules|CloudHsmCluster|not present|2022-08-31-preview|
|HardwareSecurityModules|CloudHsmClusterPrivateEndpointConnection|not present|2022-08-31-preview|
|HardwareSecurityModules|DedicatedHsm|2018-10-31-preview|2021-11-30|
|HDInsight|Application|2018-06-01-preview|2021-06-01|
|HDInsight|Cluster|2018-06-01-preview|2021-06-01|
|HDInsight|Extension|2018-06-01-preview|2021-06-01|
|HDInsight|ExtensionAzureMonitorStatus|2018-06-01-preview|2021-06-01|
|HDInsight|ExtensionMonitoringStatus|2018-06-01-preview|2021-06-01|
|HDInsight|getClusterGatewaySettings|2018-06-01-preview|2021-06-01|
|HDInsight|listVirtualMachineHosts|2018-06-01-preview|2021-06-01|
|HDInsight|PrivateEndpointConnection|2021-06-01|2021-06-01|
|HealthBot|Bot|2020-12-08|2023-05-01|
|HealthBot|listBotSecrets|not present|2023-05-01|
|HealthcareApis|AnalyticsConnector|not present|2022-10-01-preview|
|HealthcareApis|DicomService|2022-05-15|2023-02-28|
|HealthcareApis|FhirService|2022-05-15|2023-02-28|
|HealthcareApis|IotConnector|2022-05-15|2023-02-28|
|HealthcareApis|IotConnectorFhirDestination|2022-05-15|2023-02-28|
|HealthcareApis|PrivateEndpointConnection|2022-05-15|2023-02-28|
|HealthcareApis|Service|2022-05-15|2023-02-28|
|HealthcareApis|Workspace|2022-05-15|2023-02-28|
|HealthcareApis|WorkspacePrivateEndpointConnection|2022-05-15|2023-02-28|
|HybridCloud|CloudConnection|2023-01-01-preview|2023-01-01-preview|
|HybridCloud|CloudConnector|2023-01-01-preview|2023-01-01-preview|
|HybridCompute|Machine|2020-08-02|2022-12-27|
|HybridCompute|MachineExtension|2020-08-02|2022-12-27|
|HybridCompute|MachineRunCommand|not present|2023-04-25-preview|
|HybridCompute|PrivateEndpointConnection|2021-03-25-preview|2022-12-27|
|HybridCompute|PrivateLinkScope|2021-03-25-preview|2022-12-27|
|HybridCompute|PrivateLinkScopedResource|2020-08-15-preview|2020-08-15-preview|
|HybridConnectivity|Endpoint|2022-05-01-preview|2023-03-15|
|HybridConnectivity|listEndpointCredentials|2022-05-01-preview|2023-03-15|
|HybridConnectivity|listEndpointIngressGatewayCredentials|not present|2023-03-15|
|HybridConnectivity|listEndpointManagedProxyDetails|2022-05-01-preview|2023-03-15|
|HybridConnectivity|ServiceConfiguration|not present|2023-03-15|
|HybridContainerService|agentPool|2022-05-01-preview|2022-09-01-preview|
|HybridContainerService|HybridIdentityMetadatum|2022-05-01-preview|2022-09-01-preview|
|HybridContainerService|ProvisionedCluster|2022-05-01-preview|2022-09-01-preview|
|HybridContainerService|storageSpaceRetrieve|2022-05-01-preview|2022-09-01-preview|
|HybridContainerService|virtualNetworkRetrieve|2022-05-01-preview|2022-09-01-preview|
|HybridData|DataManager|2019-06-01|2019-06-01|
|HybridData|DataStore|2019-06-01|2019-06-01|
|HybridData|JobDefinition|2019-06-01|2019-06-01|
|HybridNetwork|Device|2020-01-01-preview|2022-01-01-preview|
|HybridNetwork|listDeviceRegistrationKey|2020-01-01-preview|2022-01-01-preview|
|HybridNetwork|listVendorSkusCredential|2022-01-01-preview|2022-01-01-preview|
|HybridNetwork|NetworkFunction|2020-01-01-preview|2022-01-01-preview|
|HybridNetwork|Vendor|2020-01-01-preview|2022-01-01-preview|
|HybridNetwork|VendorSkuPreview|2020-01-01-preview|2022-01-01-preview|
|HybridNetwork|VendorSkus|2020-01-01-preview|2022-01-01-preview|
|ImportExport|Job|2020-08-01|2021-01-01|
|ImportExport|listBitLockerKey|2020-08-01|2021-01-01|
|Insights|ActionGroup|2019-06-01|2023-01-01|
|Insights|ActivityLogAlert|2020-10-01|2023-01-01-preview|
|Insights|AlertRule|2016-03-01|2016-03-01|
|Insights|AnalyticsItem|2015-05-01|2015-05-01|
|Insights|AutoscaleSetting|2015-04-01|2022-10-01|
|Insights|Component|2015-05-01|2020-02-02|
|Insights|ComponentCurrentBillingFeature|2015-05-01|2015-05-01|
|Insights|ComponentLinkedStorageAccount|2020-03-01-preview|2020-03-01-preview|
|Insights|DataCollectionEndpoint|2021-09-01-preview|2022-06-01|
|Insights|DataCollectionRule|2019-11-01-preview|2022-06-01|
|Insights|DataCollectionRuleAssociation|2019-11-01-preview|2022-06-01|
|Insights|DiagnosticSetting|2017-05-01-preview|2021-05-01-preview|
|Insights|ExportConfiguration|2015-05-01|2015-05-01|
|Insights|Favorite|2015-05-01|2015-05-01|
|Insights|getDiagnosticServiceTokenReadOnly|2021-03-03-preview|2021-03-03-preview|
|Insights|getDiagnosticServiceTokenReadWrite|2021-03-03-preview|2021-03-03-preview|
|Insights|getLiveToken|2020-06-02-preview|2021-10-14|
|Insights|getTestResultFile|2020-02-10-preview|2020-02-10-preview|
|Insights|guestDiagnosticsSetting|2018-06-01-preview|2018-06-01-preview|
|Insights|GuestDiagnosticsSettingsAssociation|2018-06-01-preview|2018-06-01-preview|
|Insights|listEASubscriptionListMigrationDatePost|2017-10-01|2017-10-01|
|Insights|LogProfile|2016-03-01|2016-03-01|
|Insights|ManagementGroupDiagnosticSetting|2020-01-01-preview|2021-05-01-preview|
|Insights|MetricAlert|2018-03-01|2018-03-01|
|Insights|MyWorkbook|2020-10-20|2021-03-08|
|Insights|PrivateEndpointConnection|2019-10-17-preview|2021-07-01-preview|
|Insights|PrivateLinkScope|2019-10-17-preview|2021-07-01-preview|
|Insights|PrivateLinkScopedResource|2019-10-17-preview|2021-07-01-preview|
|Insights|ProactiveDetectionConfiguration|2015-05-01|2018-05-01-preview|
|Insights|ScheduledQueryRule|2018-04-16|2023-03-15-preview|
|Insights|SubscriptionDiagnosticSetting|2017-05-01-preview|2021-05-01-preview|
|Insights|TenantActionGroup|not present|2023-05-01-preview|
|Insights|WebTest|2015-05-01|2022-06-15|
|Insights|Workbook|2020-10-20|2022-04-01|
|Insights|WorkbookTemplate|2019-10-17-preview|2020-11-20|
|Intune|AndroidMAMPolicyByName|2015-01-14-preview|2015-01-14-preview|
|Intune|IoMAMPolicyByName|2015-01-14-preview|2015-01-14-preview|
|IoTCentral|App|2021-06-01|2021-06-01|
|IoTCentral|PrivateEndpointConnection|2021-11-01-preview|2021-11-01-preview|
|IoTFirmwareDefense|Firmware|not present|2023-02-08-preview|
|IoTFirmwareDefense|Workspace|not present|2023-02-08-preview|
|IoTSecurity|DefenderSetting|2021-02-01-preview|2021-02-01-preview|
|IoTSecurity|DeviceGroup|2021-02-01-preview|2021-02-01-preview|
|IoTSecurity|OnPremiseSensor|2021-02-01-preview|2021-02-01-preview|
|IoTSecurity|Sensor|2021-02-01-preview|2021-02-01-preview|
|IoTSecurity|Site|2021-02-01-preview|2021-02-01-preview|
|KeyVault|Key|2019-09-01|2023-02-01|
|KeyVault|ManagedHsm|2021-06-01-preview|2023-02-01|
|KeyVault|MHSMPrivateEndpointConnection|2021-06-01-preview|2023-02-01|
|KeyVault|PrivateEndpointConnection|2019-09-01|2023-02-01|
|KeyVault|Secret|2019-09-01|2023-02-01|
|KeyVault|Vault|2019-09-01|2023-02-01|
|Kubernetes|ConnectedCluster|2021-03-01|2022-05-01-preview|
|Kubernetes|listConnectedClusterUserCredential|2022-05-01-preview|2022-05-01-preview|
|Kubernetes|listConnectedClusterUserCredentials|2021-04-01-preview|2021-04-01-preview|
|KubernetesConfiguration|Extension|2020-07-01-preview|2023-05-01|
|KubernetesConfiguration|FluxConfiguration|2021-11-01-preview|2023-05-01|
|KubernetesConfiguration|PrivateEndpointConnection|2022-04-02-preview|2022-04-02-preview|
|KubernetesConfiguration|PrivateLinkScope|2022-04-02-preview|2022-04-02-preview|
|KubernetesConfiguration|SourceControlConfiguration|2021-03-01|2023-05-01|
|Kusto|AttachedDatabaseConfiguration|2021-01-01|2022-12-29|
|Kusto|Cluster|2021-01-01|2022-12-29|
|Kusto|ClusterPrincipalAssignment|2021-01-01|2022-12-29|
|Kusto|Database|2021-01-01|2022-12-29|
|Kusto|DatabasePrincipalAssignment|2021-01-01|2022-12-29|
|Kusto|DataConnection|2021-01-01|2022-12-29|
|Kusto|EventHubConnection|2018-09-07-preview|2018-09-07-preview|
|Kusto|listClusterFollowerDatabases|2021-01-01|2022-12-29|
|Kusto|listClusterLanguageExtensions|2021-01-01|2022-12-29|
|Kusto|listDatabasePrincipals|2021-01-01|2022-12-29|
|Kusto|ManagedPrivateEndpoint|2021-08-27|2022-12-29|
|Kusto|PrivateEndpointConnection|2021-08-27|2022-12-29|
|Kusto|Script|2021-01-01|2022-12-29|
|LabServices|Environment|2018-10-15|2018-10-15|
|LabServices|EnvironmentSetting|2018-10-15|2018-10-15|
|LabServices|GalleryImage|2018-10-15|2018-10-15|
|LabServices|getGlobalUserEnvironment|2018-10-15|2018-10-15|
|LabServices|getGlobalUserOperationBatchStatus|2018-10-15|2018-10-15|
|LabServices|getGlobalUserOperationStatus|2018-10-15|2018-10-15|
|LabServices|getGlobalUserPersonalPreferences|2018-10-15|2018-10-15|
|LabServices|getLabAccountRegionalAvailability|2018-10-15|2018-10-15|
|LabServices|Lab|2018-10-15|2022-08-01|
|LabServices|LabAccount|2018-10-15|2018-10-15|
|LabServices|LabPlan|2021-10-01-preview|2022-08-01|
|LabServices|listGlobalUserEnvironments|2018-10-15|2018-10-15|
|LabServices|listGlobalUserLabs|2018-10-15|2018-10-15|
|LabServices|Schedule|2021-10-01-preview|2022-08-01|
|LabServices|User|2018-10-15|2022-08-01|
|LoadTestService|LoadTest|2021-12-01-preview|2022-12-01|
|Logic|IntegrationAccount|2019-05-01|2019-05-01|
|Logic|IntegrationAccountAgreement|2019-05-01|2019-05-01|
|Logic|IntegrationAccountAssembly|2019-05-01|2019-05-01|
|Logic|IntegrationAccountBatchConfiguration|2019-05-01|2019-05-01|
|Logic|IntegrationAccountCertificate|2019-05-01|2019-05-01|
|Logic|IntegrationAccountMap|2019-05-01|2019-05-01|
|Logic|IntegrationAccountPartner|2019-05-01|2019-05-01|
|Logic|IntegrationAccountSchema|2019-05-01|2019-05-01|
|Logic|IntegrationAccountSession|2019-05-01|2019-05-01|
|Logic|IntegrationServiceEnvironment|2019-05-01|2019-05-01|
|Logic|IntegrationServiceEnvironmentManagedApi|2019-05-01|2019-05-01|
|Logic|listIntegrationAccountAgreementContentCallbackUrl|2019-05-01|2019-05-01|
|Logic|listIntegrationAccountAssemblyContentCallbackUrl|2019-05-01|2019-05-01|
|Logic|listIntegrationAccountCallbackUrl|2019-05-01|2019-05-01|
|Logic|listIntegrationAccountKeyVaultKeys|2019-05-01|2019-05-01|
|Logic|listIntegrationAccountMapContentCallbackUrl|2019-05-01|2019-05-01|
|Logic|listIntegrationAccountPartnerContentCallbackUrl|2019-05-01|2019-05-01|
|Logic|listIntegrationAccountSchemaContentCallbackUrl|2019-05-01|2019-05-01|
|Logic|listWorkflowAccessKeySecretKeys|2015-02-01-preview|2015-02-01-preview|
|Logic|listWorkflowCallbackUrl|2019-05-01|2019-05-01|
|Logic|listWorkflowRunActionExpressionTraces|2019-05-01|2019-05-01|
|Logic|listWorkflowRunActionRepetitionExpressionTraces|2019-05-01|2019-05-01|
|Logic|listWorkflowSwagger|2019-05-01|2019-05-01|
|Logic|listWorkflowTriggerCallbackUrl|2019-05-01|2019-05-01|
|Logic|listWorkflowVersionCallbackUrl|not present|2016-06-01|
|Logic|listWorkflowVersionTriggerCallbackUrl|2019-05-01|2019-05-01|
|Logic|RosettaNetProcessConfiguration|2016-06-01|2016-06-01|
|Logic|Workflow|2019-05-01|2019-05-01|
|Logic|WorkflowAccessKey|2015-02-01-preview|2015-02-01-preview|
|Logz|listMonitorMonitoredResources|2020-10-01|2022-01-01-preview|
|Logz|listMonitorUserRoles|2020-10-01|2022-01-01-preview|
|Logz|listMonitorVMHosts|2020-10-01|2022-01-01-preview|
|Logz|listSubAccountMonitoredResources|2020-10-01|2022-01-01-preview|
|Logz|listSubAccountVMHosts|2020-10-01|2022-01-01-preview|
|Logz|MetricsSource|2022-01-01-preview|2022-01-01-preview|
|Logz|MetricsSourceTagRule|2022-01-01-preview|2022-01-01-preview|
|Logz|Monitor|2020-10-01|2022-01-01-preview|
|Logz|SubAccount|2020-10-01|2022-01-01-preview|
|Logz|SubAccountTagRule|2020-10-01|2022-01-01-preview|
|Logz|TagRule|2020-10-01|2022-01-01-preview|
|M365SecurityAndCompliance|PrivateEndpointConnectionsAdtAPI|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|PrivateEndpointConnectionsComp|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|PrivateEndpointConnectionsForEDM|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|PrivateEndpointConnectionsForMIPPolicySync|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|PrivateEndpointConnectionsForSCCPowershell|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|PrivateEndpointConnectionsSec|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|privateLinkServicesForEDMUpload|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|privateLinkServicesForM365ComplianceCenter|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|privateLinkServicesForM365SecurityCenter|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|privateLinkServicesForMIPPolicySync|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|privateLinkServicesForO365ManagementActivityAPI|2021-03-25-preview|2021-03-25-preview|
|M365SecurityAndCompliance|privateLinkServicesForSCCPowershell|2021-03-25-preview|2021-03-25-preview|
|MachineLearning|CommitmentPlan|2016-05-01-preview|2016-05-01-preview|
|MachineLearning|listWorkspaceKeys|2016-04-01|2019-10-01|
|MachineLearning|WebService|2017-01-01|2017-01-01|
|MachineLearning|Workspace|2016-04-01|2019-10-01|
|MachineLearningCompute|listOperationalizationClusterKeys|2017-08-01-preview|2017-08-01-preview|
|MachineLearningCompute|OperationalizationCluster|2017-08-01-preview|2017-08-01-preview|
|MachineLearningExperimentation|Account|2017-05-01-preview|2017-05-01-preview|
|MachineLearningExperimentation|Project|2017-05-01-preview|2017-05-01-preview|
|MachineLearningExperimentation|Workspace|2017-05-01-preview|2017-05-01-preview|
|MachineLearningServices|BatchDeployment|2021-03-01-preview|2023-04-01|
|MachineLearningServices|BatchEndpoint|2021-03-01-preview|2023-04-01|
|MachineLearningServices|CodeContainer|2021-03-01-preview|2023-04-01|
|MachineLearningServices|CodeVersion|2021-03-01-preview|2023-04-01|
|MachineLearningServices|ComponentContainer|2022-02-01-preview|2023-04-01|
|MachineLearningServices|ComponentVersion|2022-02-01-preview|2023-04-01|
|MachineLearningServices|Compute|not present|2023-04-01|
|MachineLearningServices|DataContainer|2021-03-01-preview|2023-04-01|
|MachineLearningServices|Datastore|not present|2023-04-01|
|MachineLearningServices|DataVersion|2021-03-01-preview|2023-04-01|
|MachineLearningServices|EnvironmentContainer|2021-03-01-preview|2023-04-01|
|MachineLearningServices|EnvironmentSpecificationVersion|2021-03-01-preview|2021-03-01-preview|
|MachineLearningServices|EnvironmentVersion|not present|2023-04-01|
|MachineLearningServices|FeaturesetContainerEntity|not present|2023-04-01-preview|
|MachineLearningServices|FeaturesetVersion|not present|2023-04-01-preview|
|MachineLearningServices|FeaturestoreEntityContainerEntity|not present|2023-04-01-preview|
|MachineLearningServices|FeaturestoreEntityVersion|not present|2023-04-01-preview|
|MachineLearningServices|getFeaturesetVersionFeature|not present|2023-02-01-preview|
|MachineLearningServices|getOnlineDeploymentLogs|2021-03-01-preview|2023-04-01|
|MachineLearningServices|getOnlineEndpointToken|2021-03-01-preview|2023-04-01|
|MachineLearningServices|Job|2021-03-01-preview|2023-04-01|
|MachineLearningServices|LabelingJob|2020-09-01-preview|2023-04-01-preview|
|MachineLearningServices|LinkedService|2020-09-01-preview|2020-09-01-preview|
|MachineLearningServices|LinkedWorkspace|2020-03-01|2020-05-15-preview|
|MachineLearningServices|listBatchEndpointKeys|2021-03-01-preview|2023-04-01|
|MachineLearningServices|listComputeKeys|not present|2023-04-01|
|MachineLearningServices|listComputeNodes|not present|2023-04-01|
|MachineLearningServices|listDatastoreSecrets|2021-03-01-preview|2023-04-01|
|MachineLearningServices|listFeaturesetVersionFeatures|not present|2023-02-01-preview|
|MachineLearningServices|listFeaturesetVersionMaterializationJobs|not present|2023-04-01-preview|
|MachineLearningServices|listMachineLearningComputeKeys|2021-01-01|Renamed to listComputeKeys|
|MachineLearningServices|listMachineLearningComputeNodes|2021-01-01|Renamed to listComputeKeys|
|MachineLearningServices|listNotebookKeys|2021-01-01|Renamed to listWorkspaceNotebookKeys|
|MachineLearningServices|listOnlineEndpointKeys|2021-03-01-preview|2023-04-01|
|MachineLearningServices|listStorageAccountKeys|2021-01-01|Renamed to listWorkspaceStorageAccountKeys|
|MachineLearningServices|listWorkspaceConnectionSecrets|not present|2023-06-01-preview|
|MachineLearningServices|listWorkspaceKeys|2021-01-01|2023-04-01|
|MachineLearningServices|listWorkspaceNotebookAccessToken|2021-01-01|2023-04-01|
|MachineLearningServices|listWorkspaceNotebookKeys|not present|2023-04-01|
|MachineLearningServices|listWorkspaceStorageAccountKeys|not present|2023-04-01|
|MachineLearningServices|MachineLearningCompute|2021-01-01|Renamed to Compute|
|MachineLearningServices|MachineLearningDataset|2020-05-01-preview|2020-05-01-preview|
|MachineLearningServices|MachineLearningDatastore|2020-05-01-preview|2020-05-01-preview|
|MachineLearningServices|MachineLearningService|2021-01-01|No longer listed in [documentation](https:||learn.microsoft.com|en-us|rest|api|azureml|)|
|MachineLearningServices|ManagedNetworkSettingsRule|not present|2023-04-01-preview|
|MachineLearningServices|ModelContainer|2021-03-01-preview|2023-04-01|
|MachineLearningServices|ModelVersion|2021-03-01-preview|2023-04-01|
|MachineLearningServices|OnlineDeployment|2021-03-01-preview|2023-04-01|
|MachineLearningServices|OnlineEndpoint|2021-03-01-preview|2023-04-01|
|MachineLearningServices|PrivateEndpointConnection|2021-01-01|2023-04-01|
|MachineLearningServices|Registry|not present|2023-04-01|
|MachineLearningServices|RegistryCodeContainer|not present|2023-04-01|
|MachineLearningServices|RegistryCodeVersion|not present|2023-04-01|
|MachineLearningServices|RegistryComponentContainer|not present|2023-04-01|
|MachineLearningServices|RegistryComponentVersion|not present|2023-04-01|
|MachineLearningServices|RegistryDataContainer|not present|2023-04-01|
|MachineLearningServices|RegistryDataVersion|not present|2023-04-01|
|MachineLearningServices|RegistryEnvironmentContainer|not present|2023-04-01|
|MachineLearningServices|RegistryEnvironmentVersion|not present|2023-04-01|
|MachineLearningServices|RegistryModelContainer|not present|2023-04-01|
|MachineLearningServices|RegistryModelVersion|not present|2023-04-01|
|MachineLearningServices|Schedule|not present|2023-04-01|
|MachineLearningServices|Workspace|2021-01-01|2023-04-01|
|MachineLearningServices|WorkspaceConnection|2021-01-01|2023-04-01|
|Maintenance|ConfigurationAssignment|2021-04-01-preview|2022-11-01-preview|
|Maintenance|ConfigurationAssignmentParent|2021-04-01-preview|2022-11-01-preview|
|Maintenance|MaintenanceConfiguration|2020-04-01|2022-11-01-preview|
|ManagedIdentity|FederatedIdentityCredential|2022-01-31-preview|2023-01-31|
|ManagedIdentity|listUserAssignedIdentityAssociatedResources|2022-01-31-preview|2022-01-31-preview|
|ManagedIdentity|UserAssignedIdentity|2018-11-30|2023-01-31|
|ManagedNetwork|ManagedNetwork|2019-06-01-preview|2019-06-01-preview|
|ManagedNetwork|ManagedNetworkGroup|2019-06-01-preview|2019-06-01-preview|
|ManagedNetwork|ManagedNetworkPeeringPolicy|2019-06-01-preview|2019-06-01-preview|
|ManagedNetwork|ScopeAssignment|2019-06-01-preview|2019-06-01-preview|
|ManagedNetworkFabric|AccessControlList|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|ExternalNetwork|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|getL2IsolationDomainArpEntries|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|getNetworkDeviceDynamicInterfaceMaps|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|getNetworkDeviceStaticInterfaceMaps|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|getNetworkDeviceStatus|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|getNetworkInterfaceStatus|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|InternalNetwork|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|IpCommunity|not present|2023-02-01-preview|
|ManagedNetworkFabric|IpCommunityList|2023-02-01-preview|Renamed to IpCommunity|
|ManagedNetworkFabric|IpExtendedCommunity|not present|2023-02-01-preview|
|ManagedNetworkFabric|IpPrefix|not present|2023-02-01-preview|
|ManagedNetworkFabric|IpPrefixList|2023-02-01-preview|Renamed to IpPrefix|
|ManagedNetworkFabric|L2IsolationDomain|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|L3IsolationDomain|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|NetworkDevice|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|NetworkFabric|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|NetworkFabricController|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|NetworkInterface|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|NetworkRack|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|NetworkToNetworkInterconnect|2023-02-01-preview|2023-02-01-preview|
|ManagedNetworkFabric|RoutePolicy|2023-02-01-preview|2023-02-01-preview|
|ManagedServices|RegistrationAssignment|2019-09-01|2022-10-01|
|ManagedServices|RegistrationDefinition|2019-09-01|2022-10-01|
|Management|getEntity|2020-05-01|2021-04-01|
|Management|HierarchySetting|2020-05-01|2021-04-01|
|Management|ManagementGroup|2020-05-01|2021-04-01|
|Management|ManagementGroupSubscription|2020-05-01|2021-04-01|
|ManagementPartner|Partner|2018-02-01|2018-02-01|
|Maps|Account|2018-05-01|2021-02-01|
|Maps|Creator|2020-02-01-preview|2021-02-01|
|Maps|listAccountKeys|2018-05-01|2021-02-01|
|Maps|listAccountSas|2021-12-01-preview|2021-12-01-preview|
|Maps|PrivateAtlase|2020-02-01-preview|2020-02-01-preview|
|Marketplace|listPrivateStoreNewPlansNotifications|2021-12-01|2023-01-01|
|Marketplace|listPrivateStoreStopSellOffersPlansNotifications|2021-12-01|2023-01-01|
|Marketplace|listPrivateStoreSubscriptionsContext|2021-12-01|2023-01-01|
|Marketplace|PrivateStore|2021-12-01|2023-01-01|
|Marketplace|PrivateStoreCollection|2021-12-01|2023-01-01|
|Marketplace|PrivateStoreCollectionOffer|2021-12-01|2023-01-01|
|Marketplace|PrivateStoreOffer|2020-01-01|Renamed to PrivateStoreCollectionOffer|
|Media|AccountFilter|2020-05-01|2023-01-01|
|Media|Asset|2020-05-01|2023-01-01|
|Media|AssetFilter|2020-05-01|2023-01-01|
|Media|ContentKeyPolicy|2020-05-01|2023-01-01|
|Media|getAssetEncryptionKey|2020-05-01|2023-01-01|
|Media|getContentKeyPolicyPropertiesWithSecrets|2020-05-01|2023-01-01|
|Media|getLiveEventStatus|not present|2022-11-01|
|Media|getLiveEventStreamEvents|not present|2022-11-01|
|Media|getLiveEventTrackIngestHeartbeats|not present|2022-11-01|
|Media|Job|2020-05-01|2022-07-01|
|Media|listAssetContainerSas|2020-05-01|2023-01-01|
|Media|listAssetStreamingLocators|2020-05-01|2023-01-01|
|Media|listMediaServiceEdgePolicies|2020-05-01|2023-01-01|
|Media|listMediaServiceKeys|2015-10-01|2015-10-01|
|Media|listStreamingLocatorContentKeys|2020-05-01|2023-01-01|
|Media|listStreamingLocatorPaths|2020-05-01|2023-01-01|
|Media|LiveEvent|2020-05-01|2022-11-01|
|Media|LiveOutput|2020-05-01|2022-11-01|
|Media|MediaGraph|2020-02-01-preview|2020-02-01-preview|
|Media|MediaService|2020-05-01|2023-01-01|
|Media|PrivateEndpointConnection|2020-05-01|2023-01-01|
|Media|StreamingEndpoint|2020-05-01|2022-11-01|
|Media|StreamingLocator|2020-05-01|2023-01-01|
|Media|StreamingPolicy|2020-05-01|2023-01-01|
|Media|Track|2021-11-01|2023-01-01|
|Media|Transform|2020-05-01|2022-07-01|
|Migrate|Assessment|2019-10-01|2019-10-01|
|Migrate|getProjectKeys|2018-02-02|2018-02-02|
|Migrate|getSolutionConfig|2018-09-01-preview|2018-09-01-preview|
|Migrate|getWorkloadDeploymentSecretConfigurations|not present|2022-05-01-preview|
|Migrate|Group|2019-10-01|2019-10-01|
|Migrate|HyperVCollector|2019-10-01|2019-10-01|
|Migrate|ImportCollector|2019-10-01|2019-10-01|
|Migrate|MigrateAgent|not present|2022-05-01-preview|
|Migrate|MigrateProject|2018-09-01-preview|2018-09-01-preview|
|Migrate|MigrateProjectsControllerMigrateProject|not present|2020-05-01|
|Migrate|ModernizeProject|not present|2022-05-01-preview|
|Migrate|MoveCollection|2021-01-01|2022-08-01|
|Migrate|MoveResource|2021-01-01|2022-08-01|
|Migrate|PrivateEndpointConnection|2019-10-01|2019-10-01|
|Migrate|PrivateEndpointConnectionControllerPrivateEndpointConnection|not present|2020-05-01|
|Migrate|Project|2019-10-01|2019-10-01|
|Migrate|ServerCollector|2019-10-01|2019-10-01|
|Migrate|Solution|2018-09-01-preview|2018-09-01-preview|
|Migrate|VMwareCollector|2019-10-01|2019-10-01|
|Migrate|WorkloadDeployment|not present|2022-05-01-preview|
|Migrate|WorkloadInstance|not present|2022-05-01-preview|
|MixedReality|listObjectAnchorsAccountKeys|2021-03-01-preview|2021-03-01-preview|
|MixedReality|listRemoteRenderingAccountKeys|2021-01-01|2021-01-01|
|MixedReality|listSpatialAnchorsAccountKeys|2021-01-01|2021-01-01|
|MixedReality|ObjectAnchorsAccount|2021-03-01-preview|2021-03-01-preview|
|MixedReality|RemoteRenderingAccount|2021-01-01|2021-01-01|
|MixedReality|SpatialAnchorsAccount|2021-01-01|2021-01-01|
|MobileNetwork|AttachedDataNetwork|2022-04-01-preview|2023-06-01|
|MobileNetwork|DataNetwork|2022-04-01-preview|2023-06-01|
|MobileNetwork|DiagnosticsPackage|not present|2023-06-01|
|MobileNetwork|listMobileNetworkSimIds|2022-04-01-preview|2022-04-01-preview|
|MobileNetwork|MobileNetwork|2022-04-01-preview|2023-06-01|
|MobileNetwork|PacketCapture|not present|2023-06-01|
|MobileNetwork|PacketCoreControlPlane|2022-04-01-preview|2023-06-01|
|MobileNetwork|PacketCoreDataPlane|2022-04-01-preview|2023-06-01|
|MobileNetwork|Service|2022-04-01-preview|2023-06-01|
|MobileNetwork|Sim|2022-04-01-preview|2023-06-01|
|MobileNetwork|SimGroup|2022-04-01-preview|2023-06-01|
|MobileNetwork|SimPolicy|2022-04-01-preview|2023-06-01|
|MobileNetwork|Site|2022-04-01-preview|2023-06-01|
|MobileNetwork|Slice|2022-04-01-preview|2023-06-01|
|Monitor|AzureMonitorWorkspace|2021-06-03-preview|2023-04-03|
|NetApp|Account|2020-12-01|2022-11-01|
|NetApp|Backup|2020-12-01|2022-11-01|
|NetApp|BackupPolicy|2020-12-01|2022-11-01|
|NetApp|getSubvolumeMetadata|2021-10-01|2022-11-01|
|NetApp|getVolumeGroupIdForLdapUser|not present|2022-11-01|
|NetApp|listVolumeReplications|2022-01-01|2022-11-01|
|NetApp|Pool|2020-12-01|2022-11-01|
|NetApp|Snapshot|2020-12-01|2022-11-01|
|NetApp|SnapshotPolicy|2020-12-01|2022-11-01|
|NetApp|Subvolume|2021-10-01|2022-11-01|
|NetApp|Volume|2020-12-01|2022-11-01|
|NetApp|VolumeGroup|2021-10-01|2022-11-01|
|NetApp|VolumeQuotaRule|2022-01-01|2022-11-01|
|Network|AdminRule|2021-02-01-preview|2023-02-01|
|Network|AdminRuleCollection|2021-02-01-preview|2023-02-01|
|Network|ApplicationGateway|2020-11-01|2023-02-01|
|Network|ApplicationGatewayPrivateEndpointConnection|2020-11-01|2023-02-01|
|Network|ApplicationSecurityGroup|2020-11-01|2023-02-01|
|Network|AzureFirewall|2020-11-01|2023-02-01|
|Network|BastionHost|2020-11-01|2023-02-01|
|Network|ConfigurationPolicyGroup|2022-01-01|2023-02-01|
|Network|ConnectionMonitor|2020-11-01|2023-02-01|
|Network|ConnectivityConfiguration|2021-02-01-preview|2023-02-01|
|Network|CustomIPPrefix|2020-11-01|2023-02-01|
|Network|DdosCustomPolicy|2020-11-01|2023-02-01|
|Network|DdosProtectionPlan|2020-11-01|2023-02-01|
|Network|DnsForwardingRuleset|2020-04-01-preview|2022-07-01|
|Network|DnsResolver|2020-04-01-preview|2022-07-01|
|Network|DnssecConfig|not present|2023-07-01-preview|
|Network|DscpConfiguration|2020-11-01|2023-02-01|
|Network|Endpoint|2018-08-01|2022-04-01|
|Network|Experiment|2019-11-01|2019-11-01|
|Network|ExpressRouteCircuit|2020-11-01|2023-02-01|
|Network|ExpressRouteCircuitAuthorization|2020-11-01|2023-02-01|
|Network|ExpressRouteCircuitConnection|2020-11-01|2023-02-01|
|Network|ExpressRouteCircuitPeering|2020-11-01|2023-02-01|
|Network|ExpressRouteConnection|2020-11-01|2023-02-01|
|Network|ExpressRouteCrossConnectionPeering|2020-11-01|2023-02-01|
|Network|ExpressRouteGateway|2020-11-01|2023-02-01|
|Network|ExpressRoutePort|2020-11-01|2023-02-01|
|Network|ExpressRoutePortAuthorization|2022-01-01|2023-02-01|
|Network|FirewallPolicy|2020-11-01|2023-02-01|
|Network|FirewallPolicyRuleCollectionGroup|2020-11-01|2023-02-01|
|Network|FirewallPolicyRuleGroup|2020-04-01|2020-04-01|
|Network|FlowLog|2020-11-01|2023-02-01|
|Network|ForwardingRule|2020-04-01-preview|2022-07-01|
|Network|FrontDoor|2020-05-01|2021-06-01|
|Network|getActiveSessions|2020-11-01|2023-02-01|
|Network|getApplicationGatewayBackendHealthOnDemand|2020-11-01|2023-02-01|
|Network|getBastionShareableLink|2020-11-01|2023-02-01|
|Network|getDnsResourceReferenceByTarResources|2018-05-01|2023-07-01-preview|
|Network|getP2sVpnGatewayP2sVpnConnectionHealth|2020-11-01|2023-02-01|
|Network|getP2sVpnGatewayP2sVpnConnectionHealthDetailed|2020-11-01|2023-02-01|
|Network|getVirtualNetworkGatewayAdvertisedRoutes|2020-11-01|2023-02-01|
|Network|getVirtualNetworkGatewayBgpPeerStatus|2020-11-01|2023-02-01|
|Network|getVirtualNetworkGatewayConnectionIkeSas|2022-01-01|2023-02-01|
|Network|getVirtualNetworkGatewayLearnedRoutes|2020-11-01|2023-02-01|
|Network|getVirtualNetworkGatewayVpnclientConnectionHealth|2020-11-01|2023-02-01|
|Network|getVirtualNetworkGatewayVpnclientIpsecParameters|2020-11-01|2023-02-01|
|Network|getVirtualNetworkGatewayVpnProfilePackageUrl|2022-01-01|2023-02-01|
|Network|getVpnLinkConnectionIkeSas|2022-01-01|2023-02-01|
|Network|HubRouteTable|2020-11-01|2023-02-01|
|Network|HubVirtualNetworkConnection|2020-11-01|2023-02-01|
|Network|InboundEndpoint|2020-04-01-preview|2022-07-01|
|Network|InboundNatRule|2020-11-01|2023-02-01|
|Network|InterfaceEndpoint|not present|2019-02-01|
|Network|IpAllocation|2020-11-01|2023-02-01|
|Network|IpGroup|2020-11-01|2023-02-01|
|Network|listActiveConnectivityConfiguration|2021-02-01-preview|2021-02-01-preview|
|Network|listActiveConnectivityConfigurations|not present|2023-02-01|
|Network|listActiveSecurityAdminRule|2021-02-01-preview|2021-02-01-preview|
|Network|listActiveSecurityAdminRules|not present|2023-02-01|
|Network|listActiveSecurityUserRule|2021-02-01-preview|2021-02-01-preview|
|Network|listActiveSecurityUserRules|not present|2022-04-01-preview|
|Network|listDnsForwardingRulesetByVirtualNetwork|2020-04-01-preview|2022-07-01|
|Network|listDnsResolverByVirtualNetwork|2020-04-01-preview|2022-07-01|
|Network|listEffectiveConnectivityConfiguration|2021-02-01-preview|2021-02-01-preview|
|Network|listEffectiveVirtualNetworkByNetworkGroup|2021-02-01-preview|2021-02-01-preview|
|Network|listEffectiveVirtualNetworkByNetworkManager|2021-02-01-preview|2022-04-01-preview|
|Network|listFirewallPolicyIdpsSignature|2022-01-01|2023-02-01|
|Network|listFirewallPolicyIdpsSignaturesFilterValue|2022-01-01|2023-02-01|
|Network|listListEffectiveVirtualNetworkByNetworkGroup|not present|2022-04-01-preview|
|Network|listNetworkManagerDeploymentStatus|2021-02-01-preview|2023-02-01|
|Network|listNetworkManagerEffectiveConnectivityConfigurations|not present|2023-02-01|
|Network|listNetworkManagerEffectiveSecurityAdminRule|2021-02-01-preview|2021-02-01-preview|
|Network|listNetworkManagerEffectiveSecurityAdminRules|not present|2023-02-01|
|Network|LoadBalancer|2020-11-01|2023-02-01|
|Network|LoadBalancerBackendAddressPool|2020-11-01|2023-02-01|
|Network|LocalNetworkGateway|2020-11-01|2023-02-01|
|Network|ManagementGroupNetworkManagerConnection|2021-05-01-preview|2023-02-01|
|Network|NatGateway|2020-11-01|2023-02-01|
|Network|NatRule|2020-11-01|2023-02-01|
|Network|NetworkExperimentProfile|2019-11-01|2019-11-01|
|Network|NetworkGroup|2021-02-01-preview|2023-02-01|
|Network|NetworkInterface|2020-11-01|2023-02-01|
|Network|NetworkInterfaceTapConfiguration|2020-11-01|2023-02-01|
|Network|NetworkManager|2021-02-01-preview|2023-02-01|
|Network|NetworkProfile|2020-11-01|2023-02-01|
|Network|NetworkSecurityGroup|2020-11-01|2023-02-01|
|Network|NetworkSecurityPerimeter|2021-02-01-preview|2021-03-01-preview|
|Network|NetworkVirtualAppliance|2020-11-01|2023-02-01|
|Network|NetworkVirtualApplianceConnection|not present|2023-02-01|
|Network|NetworkWatcher|2020-11-01|2023-02-01|
|Network|NspAccessRule|2021-02-01-preview|2021-02-01-preview|
|Network|NspAssociation|2021-02-01-preview|2021-02-01-preview|
|Network|NspAssociationsProxy|2021-02-01-preview|No longer present|
|Network|NspLink|not present|2021-02-01-preview|
|Network|NspProfile|2021-02-01-preview|2021-02-01-preview|
|Network|OutboundEndpoint|2020-04-01-preview|2022-07-01|
|Network|P2sVpnGateway|2020-11-01|2023-02-01|
|Network|P2sVpnServerConfiguration|2019-07-01|2019-07-01|
|Network|PacketCapture|2020-11-01|2023-02-01|
|Network|Policy|2020-11-01|2022-05-01|
|Network|PrivateDnsZoneGroup|2020-11-01|2023-02-01|
|Network|PrivateEndpoint|2020-11-01|2023-02-01|
|Network|PrivateLinkService|2020-11-01|2023-02-01|
|Network|PrivateLinkServicePrivateEndpointConnection|2020-11-01|2023-02-01|
|Network|PrivateRecordSet|2020-06-01|2020-06-01|
|Network|PrivateResolverVirtualNetworkLink|not present|2022-07-01|
|Network|PrivateZone|2020-06-01|2020-06-01|
|Network|Profile|2018-08-01|2022-04-01|
|Network|PublicIPAddress|2020-11-01|2023-02-01|
|Network|PublicIPPrefix|2020-11-01|2023-02-01|
|Network|RecordSet|2018-05-01|2023-07-01-preview|
|Network|Route|2020-11-01|2023-02-01|
|Network|RouteFilter|2020-11-01|2023-02-01|
|Network|RouteFilterRule|2020-11-01|2023-02-01|
|Network|RouteMap|not present|2023-02-01|
|Network|RouteTable|2020-11-01|2023-02-01|
|Network|RoutingIntent|2022-01-01|2023-02-01|
|Network|RulesEngine|2020-05-01|2021-06-01|
|Network|ScopeConnection|2022-02-01-preview|2023-02-01|
|Network|SecurityAdminConfiguration|2021-02-01-preview|2023-02-01|
|Network|SecurityPartnerProvider|2020-11-01|2023-02-01|
|Network|SecurityRule|2020-11-01|2023-02-01|
|Network|SecurityUserConfiguration|2021-02-01-preview|2022-04-01-preview|
|Network|ServiceEndpointPolicy|2020-11-01|2023-02-01|
|Network|ServiceEndpointPolicyDefinition|2020-11-01|2023-02-01|
|Network|StaticMember|2022-02-01-preview|2023-02-01|
|Network|Subnet|2020-11-01|2023-02-01|
|Network|SubscriptionNetworkManagerConnection|2022-02-01-preview|2023-02-01|
|Network|TrafficManagerUserMetricsKey|2018-08-01|2022-04-01|
|Network|UserRule|2021-02-01-preview|2022-04-01-preview|
|Network|UserRuleCollection|2021-02-01-preview|2022-04-01-preview|
|Network|VirtualApplianceSite|2020-11-01|2023-02-01|
|Network|VirtualHub|2020-11-01|2023-02-01|
|Network|VirtualHubBgpConnection|2020-11-01|2023-02-01|
|Network|VirtualHubIpConfiguration|2020-11-01|2023-02-01|
|Network|VirtualHubRouteTableV2|2020-11-01|2023-02-01|
|Network|VirtualNetwork|2020-11-01|2023-02-01|
|Network|VirtualNetworkGateway|2020-11-01|2023-02-01|
|Network|VirtualNetworkGatewayConnection|2020-11-01|2023-02-01|
|Network|VirtualNetworkGatewayNatRule|2021-03-01|2023-02-01|
|Network|VirtualNetworkLink|2020-06-01|2020-06-01|
|Network|VirtualNetworkPeering|2020-11-01|2023-02-01|
|Network|VirtualNetworkTap|2020-11-01|2023-02-01|
|Network|VirtualRouter|2022-01-01|2023-02-01|
|Network|VirtualRouterPeering|2022-01-01|2023-02-01|
|Network|VirtualWan|2020-11-01|2023-02-01|
|Network|VpnConnection|2020-11-01|2023-02-01|
|Network|VpnGateway|2020-11-01|2023-02-01|
|Network|VpnServerConfiguration|2020-11-01|2023-02-01|
|Network|VpnSite|2020-11-01|2023-02-01|
|Network|WebApplicationFirewallPolicy|2020-11-01|2023-02-01|
|Network|Zone|2018-05-01|2023-07-01-preview|
|NetworkCloud|AgentPool|not present|2023-05-01-preview|
|NetworkCloud|BareMetalMachine|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|BareMetalMachineKeySet|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|BmcKeySet|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|CloudServicesNetwork|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|Cluster|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|ClusterManager|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|Console|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|DefaultCniNetwork|2022-12-12-preview|2022-12-12-preview|
|NetworkCloud|HybridAksCluster|2022-12-12-preview|2022-12-12-preview|
|NetworkCloud|KubernetesCluster|not present|2023-05-01-preview|
|NetworkCloud|L2Network|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|L3Network|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|MetricsConfiguration|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|Rack|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|StorageAppliance|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|TrunkedNetwork|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|VirtualMachine|2022-12-12-preview|2023-05-01-preview|
|NetworkCloud|Volume|2022-12-12-preview|2023-05-01-preview|
|NetworkFunction|AzureTrafficCollector|2022-05-01|2022-11-01|
|NetworkFunction|CollectorPolicy|2022-05-01|2022-11-01|
|NotificationHubs|getNamespacePnsCredentials|not present|2023-01-01-preview|
|NotificationHubs|getNotificationHubPnsCredentials|2017-04-01|2023-01-01-preview|
|NotificationHubs|listNamespaceKeys|2017-04-01|2023-01-01-preview|
|NotificationHubs|listNotificationHubKeys|2017-04-01|2023-01-01-preview|
|NotificationHubs|Namespace|2017-04-01|2023-01-01-preview|
|NotificationHubs|NamespaceAuthorizationRule|2017-04-01|2023-01-01-preview|
|NotificationHubs|NotificationHub|2017-04-01|2023-01-01-preview|
|NotificationHubs|NotificationHubAuthorizationRule|2017-04-01|2023-01-01-preview|
|NotificationHubs|PrivateEndpointConnection|not present|2023-01-01-preview|
|OffAzure|HyperVSite|2020-01-01|2020-07-07|
|OffAzure|MasterSite|2020-07-07|2020-07-07|
|OffAzure|PrivateEndpointConnection|2020-07-07|2020-07-07|
|OffAzure|Site|2020-01-01|2020-07-07|
|OffAzure|VCenter|2020-01-01|2020-07-07|
|OpenEnergyPlatform|EnergyService|2022-04-04-preview|2022-04-04-preview|
|OpenEnergyPlatform|listEnergyServicePartitions|2022-04-04-preview|2022-04-04-preview|
|OperationalInsights|Cluster|2020-10-01|2021-06-01|
|OperationalInsights|DataExport|2020-08-01|2020-08-01|
|OperationalInsights|DataSource|2020-08-01|2020-08-01|
|OperationalInsights|getSharedKeys|2020-08-01|2020-08-01|
|OperationalInsights|getWorkspaceSharedKeys|not present|2015-11-01-preview|
|OperationalInsights|LinkedService|2020-08-01|2020-08-01|
|OperationalInsights|LinkedStorageAccount|2020-08-01|2020-08-01|
|OperationalInsights|MachineGroup|2015-11-01-preview|2015-11-01-preview|
|OperationalInsights|Query|2019-09-01|2019-09-01|
|OperationalInsights|QueryPack|2019-09-01|2019-09-01|
|OperationalInsights|SavedSearch|2020-08-01|2020-08-01|
|OperationalInsights|StorageInsightConfig|2020-08-01|2020-08-01|
|OperationalInsights|Table|2021-12-01-preview|2022-10-01|
|OperationalInsights|Workspace|2020-10-01|2022-10-01|
|OperationsManagement|ManagementAssociation|2015-11-01-preview|2015-11-01-preview|
|OperationsManagement|ManagementConfiguration|2015-11-01-preview|2015-11-01-preview|
|OperationsManagement|Solution|2015-11-01-preview|2015-11-01-preview|
|Orbital|Contact|not present|2022-11-01|
|Orbital|ContactProfile|not present|2022-11-01|
|Orbital|listSpacecraftAvailableContacts|not present|2022-11-01|
|Orbital|Spacecraft|not present|2022-11-01|
|Peering|ConnectionMonitorTest|2021-06-01|2022-10-01|
|Peering|PeerAsn|2021-01-01|2022-10-01|
|Peering|Peering|2021-01-01|2022-10-01|
|Peering|PeeringService|2021-01-01|2022-10-01|
|Peering|Prefix|2021-01-01|2022-10-01|
|Peering|RegisteredAsn|2021-01-01|2022-10-01|
|Peering|RegisteredPrefix|2021-01-01|2022-10-01|
|PolicyInsights|AttestationAtResource|2021-01-01|2022-09-01|
|PolicyInsights|AttestationAtResourceGroup|2021-01-01|2022-09-01|
|PolicyInsights|AttestationAtSubscription|2021-01-01|2022-09-01|
|PolicyInsights|listRemediationDeploymentsAtManagementGroup|2019-07-01|2021-10-01|
|PolicyInsights|listRemediationDeploymentsAtResource|2019-07-01|2021-10-01|
|PolicyInsights|listRemediationDeploymentsAtResourceGroup|2019-07-01|2021-10-01|
|PolicyInsights|listRemediationDeploymentsAtSubscription|2019-07-01|2021-10-01|
|PolicyInsights|RemediationAtManagementGroup|2019-07-01|2021-10-01|
|PolicyInsights|RemediationAtResource|2019-07-01|2021-10-01|
|PolicyInsights|RemediationAtResourceGroup|2019-07-01|2021-10-01|
|PolicyInsights|RemediationAtSubscription|2019-07-01|2021-10-01|
|Portal|Console|2018-10-01|2018-10-01|
|Portal|ConsoleWithLocation|2018-10-01|2018-10-01|
|Portal|Dashboard|2020-09-01-preview|2020-09-01-preview|
|Portal|listListTenantConfigurationViolation|2020-09-01-preview|2020-09-01-preview|
|Portal|TenantConfiguration|2020-09-01-preview|2020-09-01-preview|
|Portal|UserSettings|2018-10-01|2018-10-01|
|Portal|UserSettingsWithLocation|2018-10-01|2018-10-01|
|PowerBI|listWorkspaceCollectionAccessKeys|2016-01-29|2016-01-29|
|PowerBI|PowerBIResource|2020-06-01|2020-06-01|
|PowerBI|PrivateEndpointConnection|2020-06-01|2020-06-01|
|PowerBI|WorkspaceCollection|2016-01-29|2016-01-29|
|PowerBIDedicated|AutoScaleVCore|2021-01-01|2021-01-01|
|PowerBIDedicated|CapacityDetails|2021-01-01|2021-01-01|
|PowerPlatform|Account|2020-10-30-preview|2020-10-30-preview|
|PowerPlatform|EnterprisePolicy|2020-10-30-preview|2020-10-30-preview|
|PowerPlatform|PrivateEndpointConnection|2020-10-30-preview|2020-10-30-preview|
|ProfessionalService|ProfessionalServiceSubscriptionLevel|2023-07-01-preview|2023-07-01-preview|
|ProviderHub|DefaultRollout|2020-11-20|2021-09-01-preview|
|ProviderHub|NotificationRegistration|2020-11-20|2021-09-01-preview|
|ProviderHub|OperationByProviderRegistration|2020-11-20|2021-09-01-preview|
|ProviderHub|ProviderRegistration|2020-11-20|2021-09-01-preview|
|ProviderHub|ResourceTypeRegistration|2020-11-20|2021-09-01-preview|
|ProviderHub|Skus|2020-11-20|2021-09-01-preview|
|ProviderHub|SkusNestedResourceTypeFirst|2020-11-20|2021-09-01-preview|
|ProviderHub|SkusNestedResourceTypeSecond|2020-11-20|2021-09-01-preview|
|ProviderHub|SkusNestedResourceTypeThird|2020-11-20|2021-09-01-preview|
|Purview|Account|2020-12-01-preview|2021-12-01|
|Purview|KafkaConfiguration|not present|2021-12-01|
|Purview|listAccountKeys|2020-12-01-preview|2021-12-01|
|Purview|listFeatureAccount|not present|2021-12-01|
|Purview|listFeatureSubscription|not present|2021-12-01|
|Purview|PrivateEndpointConnection|2020-12-01-preview|2021-12-01|
|Quantum|Workspace|2019-11-04-preview|2022-01-10-preview|
|RecommendationsService|Account|2022-02-01|2022-02-01|
|RecommendationsService|Modeling|2022-02-01|2022-02-01|
|RecommendationsService|ServiceEndpoint|2022-02-01|2022-02-01|
|RecoveryServices|getRecoveryPointAccessToken|2018-12-20|2023-01-15|
|RecoveryServices|PrivateEndpointConnection|2021-02-01|2023-04-01|
|RecoveryServices|ProtectedItem|2021-02-01|2023-04-01|
|RecoveryServices|ProtectionContainer|2021-02-01|2023-04-01|
|RecoveryServices|ProtectionIntent|2021-02-01|2023-04-01|
|RecoveryServices|ProtectionPolicy|2021-02-01|2023-04-01|
|RecoveryServices|ReplicationFabric|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationMigrationItem|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationNetworkMapping|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationPolicy|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationProtectedItem|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationProtectionContainerMapping|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationRecoveryPlan|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationRecoveryServicesProvider|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationStorageClassificationMapping|2018-07-10|2023-04-01|
|RecoveryServices|ReplicationvCenter|2018-07-10|2023-04-01|
|RecoveryServices|ResourceGuardProxy|2021-02-01-preview|2023-04-01|
|RecoveryServices|Vault|2021-01-01|2023-04-01|
|RedHatOpenShift|listOpenShiftClusterAdminCredentials|2021-09-01-preview|2022-09-04|
|RedHatOpenShift|listOpenShiftClusterCredentials|2020-04-30|2022-09-04|
|RedHatOpenShift|MachinePool|not present|2022-09-04|
|RedHatOpenShift|OpenShiftCluster|2020-04-30|2022-09-04|
|RedHatOpenShift|Secret|not present|2022-09-04|
|RedHatOpenShift|SyncIdentityProvider|not present|2022-09-04|
|RedHatOpenShift|SyncSet|not present|2022-09-04|
|Relay|HybridConnection|2017-04-01|2021-11-01|
|Relay|HybridConnectionAuthorizationRule|2017-04-01|2021-11-01|
|Relay|listHybridConnectionKeys|2017-04-01|2021-11-01|
|Relay|listNamespaceKeys|2017-04-01|2021-11-01|
|Relay|listWCFRelayKeys|2017-04-01|2021-11-01|
|Relay|Namespace|2017-04-01|2021-11-01|
|Relay|NamespaceAuthorizationRule|2017-04-01|2021-11-01|
|Relay|PrivateEndpointConnection|2018-01-01-preview|2021-11-01|
|Relay|WCFRelay|2017-04-01|2021-11-01|
|Relay|WCFRelayAuthorizationRule|2017-04-01|2021-11-01|
|ResourceConnector|Appliance|2021-10-31-preview|2022-10-27|
|ResourceConnector|listApplianceClusterCustomerUserCredential|2022-04-15-preview|2022-04-15-preview|
|ResourceConnector|listApplianceClusterUserCredential|2021-10-31-preview|2022-10-27|
|ResourceConnector|listApplianceKeys|not present|2022-10-27|
|ResourceGraph|GraphQuery|2018-09-01-preview|2020-04-01-preview|
|ResourceHealth|listSecurityAdvisoryImpactedResourceBySubscriptionIdAndEventId|2022-10-01-preview|2022-10-01|
|ResourceHealth|listSecurityAdvisoryImpactedResourceByTenantIdAndEventId|2022-10-01-preview|2022-10-01|
|Resources|Deployment|2021-01-01|2022-09-01|
|Resources|DeploymentAtManagementGroupScope|2021-01-01|2022-09-01|
|Resources|DeploymentAtScope|2021-01-01|2022-09-01|
|Resources|DeploymentAtSubscriptionScope|2021-01-01|2022-09-01|
|Resources|DeploymentAtTenantScope|2021-01-01|2022-09-01|
|Resources|DeploymentScript|2020-10-01|2020-10-01|
|Resources|DeploymentStackAtManagementGroup|not present|2022-08-01-preview|
|Resources|DeploymentStackAtResourceGroup|not present|2022-08-01-preview|
|Resources|DeploymentStackAtSubscription|not present|2022-08-01-preview|
|Resources|Resource|2019-05-01|2022-09-01|
|Resources|ResourceGroup|2019-05-01|2022-09-01|
|Resources|TagAtScope|2019-10-01|2022-09-01|
|Resources|TemplateSpec|2022-02-01|2022-02-01|
|Resources|TemplateSpecVersion|2022-02-01|2022-02-01|
|SaaS|listSaasResourceAccessToken|2018-03-01-beta|2018-03-01-beta|
|SaaS|listSaasSubscriptionLevelAccessToken|2018-03-01-beta|2018-03-01-beta|
|SaaS|SaasSubscriptionLevel|2018-03-01-beta|2018-03-01-beta|
|Scheduler|Job|2016-03-01|2016-03-01|
|Scheduler|JobCollection|2016-03-01|2016-03-01|
|ScVmm|AvailabilitySet|2020-06-05-preview|2022-05-21-preview|
|ScVmm|Cloud|2020-06-05-preview|2022-05-21-preview|
|ScVmm|GuestAgent|not present|2022-05-21-preview|
|ScVmm|HybridIdentityMetadata|not present|2022-05-21-preview|
|ScVmm|InventoryItem|2020-06-05-preview|2022-05-21-preview|
|ScVmm|MachineExtension|not present|2022-05-21-preview|
|ScVmm|VirtualMachine|2020-06-05-preview|2022-05-21-preview|
|ScVmm|VirtualMachineTemplate|2020-06-05-preview|2022-05-21-preview|
|ScVmm|VirtualNetwork|2020-06-05-preview|2022-05-21-preview|
|ScVmm|VmmServer|2020-06-05-preview|2022-05-21-preview|
|Search|listAdminKey|2020-08-01|2022-09-01|
|Search|listQueryKeyBySearchService|2020-08-01|2022-09-01|
|Search|PrivateEndpointConnection|2020-08-01|2022-09-01|
|Search|Service|2020-08-01|2022-09-01|
|Search|SharedPrivateLinkResource|2020-08-01|2022-09-01|
|Security|AdaptiveApplicationControl|2020-01-01|2020-01-01|
|Security|AdvancedThreatProtection|2019-01-01|2019-01-01|
|Security|AlertsSuppressionRule|2019-01-01-preview|2019-01-01-preview|
|Security|APICollection|not present|2022-11-20-preview|
|Security|Application|2022-07-01-preview|2022-07-01-preview|
|Security|Assessment|2020-01-01|2021-06-01|
|Security|AssessmentMetadataInSubscription|2020-01-01|2021-06-01|
|Security|AssessmentsMetadataSubscription|not present|2019-01-01-preview|
|Security|Assignment|2021-08-01-preview|2021-08-01-preview|
|Security|Automation|2019-01-01-preview|2019-01-01-preview|
|Security|Connector|2020-01-01-preview|2020-01-01-preview|
|Security|CustomAssessmentAutomation|2021-07-01-preview|2021-07-01-preview|
|Security|CustomEntityStoreAssignment|2021-07-01-preview|2021-07-01-preview|
|Security|DeviceSecurityGroup|2019-08-01|2019-08-01|
|Security|GovernanceAssignment|not present|2022-01-01-preview|
|Security|GovernanceRule|not present|2022-01-01-preview|
|Security|IngestionSetting|2021-01-15-preview|2021-01-15-preview|
|Security|IotSecuritySolution|2019-08-01|2019-08-01|
|Security|JitNetworkAccessPolicy|2020-01-01|2020-01-01|
|Security|listIngestionSettingConnectionStrings|2021-01-15-preview|2021-01-15-preview|
|Security|listIngestionSettingTokens|2021-01-15-preview|2021-01-15-preview|
|Security|SecurityConnector|2021-07-01-preview|2023-03-01-preview|
|Security|SecurityConnectorApplication|2022-07-01-preview|2022-07-01-preview|
|Security|SecurityContact|2020-01-01-preview|2020-01-01-preview|
|Security|SecurityOperator|not present|2023-01-01-preview|
|Security|ServerVulnerabilityAssessment|2020-01-01|2020-01-01|
|Security|ServerVulnerabilityAssessmentsSetting|not present|2023-05-01|
|Security|SqlVulnerabilityAssessmentBaselineRule|2020-07-01-preview|2023-02-01-preview|
|Security|Standard|2021-08-01-preview|2021-08-01-preview|
|Security|WorkspaceSetting|2017-08-01-preview|2017-08-01-preview|
|SecurityAndCompliance|PrivateEndpointConnectionsAdtAPI|2021-03-08|2021-03-08|
|SecurityAndCompliance|PrivateEndpointConnectionsComp|2021-03-08|2021-03-08|
|SecurityAndCompliance|PrivateEndpointConnectionsForEDM|2021-03-08|2021-03-08|
|SecurityAndCompliance|PrivateEndpointConnectionsForMIPPolicySync|2021-03-08|2021-03-08|
|SecurityAndCompliance|PrivateEndpointConnectionsForSCCPowershell|2021-03-08|2021-03-08|
|SecurityAndCompliance|PrivateEndpointConnectionsSec|2021-03-08|2021-03-08|
|SecurityAndCompliance|privateLinkServicesForEDMUpload|2021-03-08|2021-03-08|
|SecurityAndCompliance|privateLinkServicesForM365ComplianceCenter|2021-03-08|2021-03-08|
|SecurityAndCompliance|privateLinkServicesForM365SecurityCenter|2021-03-08|2021-03-08|
|SecurityAndCompliance|privateLinkServicesForMIPPolicySync|2021-03-08|2021-03-08|
|SecurityAndCompliance|privateLinkServicesForO365ManagementActivityAPI|2021-03-08|2021-03-08|
|SecurityAndCompliance|privateLinkServicesForSCCPowershell|2021-03-08|2021-03-08|
|SecurityDevOps|AzureDevOpsConnector|2022-09-01-preview|2022-09-01-preview|
|SecurityDevOps|GitHubConnector|2022-09-01-preview|2022-09-01-preview|
|SecurityInsights|Action|2020-01-01|2023-02-01|
|SecurityInsights|AlertRule|2020-01-01|2023-02-01|
|SecurityInsights|AutomationRule|2019-01-01-preview|2023-02-01|
|SecurityInsights|Bookmark|2020-01-01|2023-02-01|
|SecurityInsights|BookmarkRelation|2019-01-01-preview|2023-06-01-preview|
|SecurityInsights|ContentPackage|not present|2023-06-01-preview|
|SecurityInsights|ContentTemplate|not present|2023-06-01-preview|
|SecurityInsights|DataConnector|2020-01-01|2023-02-01|
|SecurityInsights|EntityQuery|2021-03-01-preview|2023-06-01-preview|
|SecurityInsights|FileImport|not present|2023-06-01-preview|
|SecurityInsights|getEntitiesGetTimeline|2019-01-01-preview|2023-06-01-preview|
|SecurityInsights|getEntityInsights|2019-01-01-preview|2023-06-01-preview|
|SecurityInsights|Hunt|not present|2023-06-01-preview|
|SecurityInsights|HuntComment|not present|2023-06-01-preview|
|SecurityInsights|HuntRelation|not present|2023-06-01-preview|
|SecurityInsights|Incident|2020-01-01|2023-02-01|
|SecurityInsights|IncidentComment|2021-03-01-preview|2023-02-01|
|SecurityInsights|IncidentRelation|2021-03-01-preview|2023-02-01|
|SecurityInsights|IncidentTask|not present|2023-06-01-preview|
|SecurityInsights|listSourceControlRepositories|2021-03-01-preview|2023-06-01-preview|
|SecurityInsights|Metadata|2021-03-01-preview|2023-02-01|
|SecurityInsights|ProductSetting|2021-03-01-preview|2023-06-01-preview|
|SecurityInsights|SecurityMLAnalyticsSetting|2022-05-01-preview|2023-02-01|
|SecurityInsights|SentinelOnboardingState|2021-03-01-preview|2023-02-01|
|SecurityInsights|SourceControl|2021-03-01-preview|2023-05-01-preview|
|SecurityInsights|ThreatIntelligenceIndicator|2019-01-01-preview|2023-02-01|
|SecurityInsights|Watchlist|2021-03-01-preview|2023-02-01|
|SecurityInsights|WatchlistItem|2021-03-01-preview|2023-02-01|
|SecurityInsights|WorkspaceManagerAssignment|not present|2023-06-01-preview|
|SecurityInsights|WorkspaceManagerConfiguration|not present|2023-06-01-preview|
|SecurityInsights|WorkspaceManagerGroup|not present|2023-06-01-preview|
|SecurityInsights|WorkspaceManagerMember|not present|2023-06-01-preview|
|SerialConsole|SerialPort|2018-05-01|2018-05-01|
|ServiceBus|DisasterRecoveryConfig|2017-04-01|2022-01-01-preview|
|ServiceBus|listDisasterRecoveryConfigKeys|2017-04-01|2022-01-01-preview|
|ServiceBus|listNamespaceKeys|2017-04-01|2022-01-01-preview|
|ServiceBus|listQueueKeys|2017-04-01|2022-01-01-preview|
|ServiceBus|listTopicKeys|2017-04-01|2022-01-01-preview|
|ServiceBus|MigrationConfig|2017-04-01|2022-01-01-preview|
|ServiceBus|Namespace|2017-04-01|2022-01-01-preview|
|ServiceBus|NamespaceAuthorizationRule|2017-04-01|2022-01-01-preview|
|ServiceBus|NamespaceIpFilterRule|2018-01-01-preview|2018-01-01-preview|
|ServiceBus|NamespaceNetworkRuleSet|2017-04-01|2022-01-01-preview|
|ServiceBus|NamespaceVirtualNetworkRule|2018-01-01-preview|2018-01-01-preview|
|ServiceBus|PrivateEndpointConnection|2018-01-01-preview|2022-01-01-preview|
|ServiceBus|Queue|2017-04-01|2022-01-01-preview|
|ServiceBus|QueueAuthorizationRule|2017-04-01|2022-01-01-preview|
|ServiceBus|Rule|2017-04-01|2022-01-01-preview|
|ServiceBus|Subscription|2017-04-01|2022-01-01-preview|
|ServiceBus|Topic|2017-04-01|2022-01-01-preview|
|ServiceBus|TopicAuthorizationRule|2017-04-01|2022-01-01-preview|
|ServiceFabric|Application|2020-03-01|This was deprecated [along with the non-managed cluster](https://learn.microsoft.com/en-us/azure/service-fabric/faq-managed-cluster)|
|ServiceFabric|ApplicationType|2020-03-01|This was deprecated [along with the non-managed cluster](https://learn.microsoft.com/en-us/azure/service-fabric/faq-managed-cluster)|
|ServiceFabric|ApplicationTypeVersion|2020-03-01|This was deprecated [along with the non-managed cluster](https://learn.microsoft.com/en-us/azure/service-fabric/faq-managed-cluster)|
|ServiceFabric|Cluster|2020-03-01|Replaced with the more recent ManagedCluster. This is still available via the version-specific module `v20210601`|
|ServiceFabric|getmanagedAzResiliencyStatus|2022-02-01-preview|2023-03-01-preview|
|ServiceFabric|listListUpgradableVersionPost|2020-12-01-preview|This was deprecated [along with the non-managed cluster](https://learn.microsoft.com/en-us/azure/service-fabric/faq-managed-cluster)|
|ServiceFabric|ManagedCluster|2020-01-01-preview|2023-03-01-preview|
|ServiceFabric|ManagedClusterApplication|not present|2023-03-01-preview|
|ServiceFabric|ManagedClusterApplicationType|not present|2023-03-01-preview|
|ServiceFabric|ManagedClusterApplicationTypeVersion|not present|2023-03-01-preview|
|ServiceFabric|ManagedClusterService|not present|2023-03-01-preview|
|ServiceFabric|NodeType|2020-01-01-preview|2023-03-01-preview|
|ServiceFabric|Service|2020-03-01|This was deprecated [along with the non-managed cluster](https://learn.microsoft.com/en-us/azure/service-fabric/faq-managed-cluster)|
|ServiceFabricMesh|Application|2018-09-01-preview|2018-09-01-preview|
|ServiceFabricMesh|Gateway|2018-09-01-preview|2018-09-01-preview|
|ServiceFabricMesh|listSecretValue|2018-09-01-preview|2018-09-01-preview|
|ServiceFabricMesh|Network|2018-09-01-preview|2018-09-01-preview|
|ServiceFabricMesh|Secret|2018-09-01-preview|2018-09-01-preview|
|ServiceFabricMesh|SecretValue|2018-09-01-preview|2018-09-01-preview|
|ServiceFabricMesh|Volume|2018-09-01-preview|2018-09-01-preview|
|ServiceLinker|Connector|not present|2022-11-01-preview|
|ServiceLinker|ConnectorDryrun|not present|2022-11-01-preview|
|ServiceLinker|Linker|2021-11-01-preview|2022-11-01-preview|
|ServiceLinker|LinkerDryrun|not present|2022-11-01-preview|
|ServiceLinker|listLinkerConfigurations|2021-11-01-preview|2022-11-01-preview|
|ServiceNetworking|AssociationsInterface|2022-10-01-preview|2023-05-01-preview|
|ServiceNetworking|FrontendsInterface|2022-10-01-preview|2023-05-01-preview|
|ServiceNetworking|TrafficControllerInterface|2022-10-01-preview|2023-05-01-preview|
|SignalRService|listSignalRKeys|2020-05-01|2023-02-01|
|SignalRService|SignalR|2020-05-01|2023-02-01|
|SignalRService|SignalRCustomCertificate|2022-02-01|2023-02-01|
|SignalRService|SignalRCustomDomain|2022-02-01|2023-02-01|
|SignalRService|SignalRPrivateEndpointConnection|2020-05-01|2023-02-01|
|SignalRService|SignalRReplica|not present|2023-03-01-preview|
|SignalRService|SignalRSharedPrivateLinkResource|2021-04-01-preview|2023-02-01|
|SoftwarePlan|HybridUseBenefit|2019-06-01-preview|2019-12-01|
|Solutions|Application|2019-07-01|2021-07-01|
|Solutions|ApplicationDefinition|2019-07-01|2021-07-01|
|Solutions|JitRequest|2019-07-01|2021-07-01|
|Solutions|listApplicationAllowedUpgradePlans|2021-07-01|2021-07-01|
|Solutions|listApplicationTokens|not present|2021-07-01|
|Sql|BackupLongTermRetentionPolicy|not present|2017-03-01-preview|
|Sql|BackupShortTermRetentionPolicy|2020-11-01-preview|2021-11-01|
|Sql|Database|2020-11-01-preview|2021-11-01|
|Sql|DatabaseAdvisor|2020-11-01-preview|2021-11-01|
|Sql|DatabaseBlobAuditingPolicy|2020-11-01-preview|2021-11-01|
|Sql|DatabaseSecurityAlertPolicy|2020-11-01-preview|2021-11-01|
|Sql|DatabaseSqlVulnerabilityAssessmentRuleBaseline|not present|2022-11-01-preview|
|Sql|DatabaseThreatDetectionPolicy|not present|2014-04-01|
|Sql|DatabaseVulnerabilityAssessment|2020-11-01-preview|2021-11-01|
|Sql|DatabaseVulnerabilityAssessmentRuleBaseline|2020-11-01-preview|2021-11-01|
|Sql|DataMaskingPolicy|2014-04-01|2021-11-01|
|Sql|DisasterRecoveryConfiguration|2014-04-01|2014-04-01|
|Sql|DistributedAvailabilityGroup|2021-05-01-preview|2021-11-01|
|Sql|ElasticPool|2020-11-01-preview|2021-11-01|
|Sql|EncryptionProtector|2020-11-01-preview|2021-11-01|
|Sql|ExtendedDatabaseBlobAuditingPolicy|2020-11-01-preview|2021-11-01|
|Sql|ExtendedServerBlobAuditingPolicy|2020-11-01-preview|2021-11-01|
|Sql|FailoverGroup|2020-11-01-preview|2021-11-01|
|Sql|FirewallRule|2020-11-01-preview|2021-11-01|
|Sql|GeoBackupPolicy|2014-04-01|2021-11-01|
|Sql|InstanceFailoverGroup|2020-11-01-preview|2021-11-01|
|Sql|InstancePool|2020-11-01-preview|2021-11-01|
|Sql|IPv6FirewallRule|2021-08-01-preview|2021-11-01|
|Sql|Job|2020-11-01-preview|2021-11-01|
|Sql|JobAgent|2020-11-01-preview|2021-11-01|
|Sql|JobCredential|2020-11-01-preview|2021-11-01|
|Sql|JobStep|2020-11-01-preview|2021-11-01|
|Sql|JobTargetGroup|2020-11-01-preview|2021-11-01|
|Sql|LongTermRetentionPolicy|2020-11-01-preview|2021-11-01|
|Sql|ManagedDatabase|2020-11-01-preview|2021-11-01|
|Sql|ManagedDatabaseSensitivityLabel|2020-11-01-preview|2021-11-01|
|Sql|ManagedDatabaseVulnerabilityAssessment|2020-11-01-preview|2021-11-01|
|Sql|ManagedDatabaseVulnerabilityAssessmentRuleBaseline|2020-11-01-preview|2021-11-01|
|Sql|ManagedInstance|2020-11-01-preview|2021-11-01|
|Sql|ManagedInstanceAdministrator|2020-11-01-preview|2021-11-01|
|Sql|ManagedInstanceAzureADOnlyAuthentication|2020-11-01-preview|2021-11-01|
|Sql|ManagedInstanceKey|2020-11-01-preview|2021-11-01|
|Sql|ManagedInstanceLongTermRetentionPolicy|not present|2022-11-01-preview|
|Sql|ManagedInstancePrivateEndpointConnection|2020-11-01-preview|2021-11-01|
|Sql|ManagedInstanceVulnerabilityAssessment|2020-11-01-preview|2021-11-01|
|Sql|ManagedServerDnsAlias|2021-11-01-preview|2021-11-01|
|Sql|OutboundFirewallRule|2021-02-01-preview|2021-11-01|
|Sql|PrivateEndpointConnection|2020-11-01-preview|2021-11-01|
|Sql|SensitivityLabel|2020-11-01-preview|2021-11-01|
|Sql|Server|2020-11-01-preview|2021-11-01|
|Sql|ServerAdvisor|2020-11-01-preview|2021-11-01|
|Sql|ServerAzureADAdministrator|2020-11-01-preview|2021-11-01|
|Sql|ServerAzureADOnlyAuthentication|2020-11-01-preview|2021-11-01|
|Sql|ServerBlobAuditingPolicy|2020-11-01-preview|2021-11-01|
|Sql|ServerCommunicationLink|2014-04-01|2014-04-01|
|Sql|ServerDnsAlias|2020-11-01-preview|2021-11-01|
|Sql|ServerKey|2020-11-01-preview|2021-11-01|
|Sql|ServerSecurityAlertPolicy|2020-11-01-preview|2021-11-01|
|Sql|ServerTrustCertificate|2021-05-01-preview|2021-11-01|
|Sql|ServerTrustGroup|2020-11-01-preview|2021-11-01|
|Sql|ServerVulnerabilityAssessment|2020-11-01-preview|2021-11-01|
|Sql|SqlVulnerabilityAssessmentRuleBaseline|not present|2022-11-01-preview|
|Sql|SqlVulnerabilityAssessmentsSetting|not present|2022-11-01-preview|
|Sql|StartStopManagedInstanceSchedule|not present|2022-11-01-preview|
|Sql|SyncAgent|2020-11-01-preview|2021-11-01|
|Sql|SyncGroup|2020-11-01-preview|2021-11-01|
|Sql|SyncMember|2020-11-01-preview|2021-11-01|
|Sql|TransparentDataEncryption|2014-04-01|2021-11-01|
|Sql|VirtualNetworkRule|2020-11-01-preview|2021-11-01|
|Sql|WorkloadClassifier|2020-11-01-preview|2021-11-01|
|Sql|WorkloadGroup|2020-11-01-preview|2021-11-01|
|SqlVirtualMachine|AvailabilityGroupListener|2017-03-01-preview|2022-02-01|
|SqlVirtualMachine|SqlVirtualMachine|2017-03-01-preview|2022-02-01|
|SqlVirtualMachine|SqlVirtualMachineGroup|2017-03-01-preview|2022-02-01|
|Storage|BlobContainer|2021-02-01|2022-09-01|
|Storage|BlobContainerImmutabilityPolicy|2021-02-01|2022-09-01|
|Storage|BlobInventoryPolicy|2021-02-01|2022-09-01|
|Storage|BlobServiceProperties|2021-02-01|2022-09-01|
|Storage|EncryptionScope|2021-02-01|2022-09-01|
|Storage|FileServiceProperties|2021-02-01|2022-09-01|
|Storage|FileShare|2021-02-01|2022-09-01|
|Storage|listLocalUserKeys|2021-08-01|2022-09-01|
|Storage|listStorageAccountKeys|2021-02-01|2022-09-01|
|Storage|listStorageAccountSAS|2021-02-01|2022-09-01|
|Storage|listStorageAccountServiceSAS|2021-02-01|2022-09-01|
|Storage|LocalUser|2021-08-01|2022-09-01|
|Storage|ManagementPolicy|2021-02-01|2022-09-01|
|Storage|ObjectReplicationPolicy|2021-02-01|2022-09-01|
|Storage|PrivateEndpointConnection|2021-02-01|2022-09-01|
|Storage|Queue|2021-02-01|2022-09-01|
|Storage|QueueServiceProperties|2021-02-01|2022-09-01|
|Storage|StorageAccount|2021-02-01|2022-09-01|
|Storage|Table|2021-02-01|2022-09-01|
|Storage|TableServiceProperties|2021-02-01|2022-09-01|
|StorageCache|amlFilesystem|not present|2023-05-01|
|StorageCache|Cache|2021-03-01|2023-05-01|
|StorageCache|getRequiredAmlFSSubnetsSize|not present|2023-05-01|
|StorageCache|StorageTarget|2021-03-01|2023-05-01|
|StorageMover|Agent|2022-07-01-preview|2023-03-01|
|StorageMover|Endpoint|2022-07-01-preview|2023-03-01|
|StorageMover|JobDefinition|2022-07-01-preview|2023-03-01|
|StorageMover|Project|2022-07-01-preview|2023-03-01|
|StorageMover|StorageMover|2022-07-01-preview|2023-03-01|
|StoragePool|DiskPool|2020-03-15-preview|2021-08-01|
|StoragePool|IscsiTarget|2020-03-15-preview|2021-08-01|
|StorageSync|CloudEndpoint|2020-03-01|2022-06-01|
|StorageSync|PrivateEndpointConnection|2020-03-01|2022-06-01|
|StorageSync|RegisteredServer|2020-03-01|2022-06-01|
|StorageSync|ServerEndpoint|2020-03-01|2022-06-01|
|StorageSync|StorageSyncService|2020-03-01|2022-06-01|
|StorageSync|SyncGroup|2020-03-01|2022-06-01|
|StorSimple|AccessControlRecord|2017-06-01|2017-06-01|
|StorSimple|BackupPolicy|2017-06-01|2017-06-01|
|StorSimple|BackupSchedule|2017-06-01|2017-06-01|
|StorSimple|BandwidthSetting|2017-06-01|2017-06-01|
|StorSimple|getManagerDevicePublicEncryptionKey|2017-06-01|2017-06-01|
|StorSimple|listDeviceFailoverSets|2017-06-01|2017-06-01|
|StorSimple|listDeviceFailoverTars|2017-06-01|2017-06-01|
|StorSimple|listManagerActivationKey|2017-06-01|2017-06-01|
|StorSimple|listManagerPublicEncryptionKey|2017-06-01|2017-06-01|
|StorSimple|Manager|2017-06-01|2017-06-01|
|StorSimple|ManagerExtendedInfo|2017-06-01|2017-06-01|
|StorSimple|StorageAccountCredential|2017-06-01|2017-06-01|
|StorSimple|Volume|2017-06-01|2017-06-01|
|StorSimple|VolumeContainer|2017-06-01|2017-06-01|
|StreamAnalytics|Cluster|2020-03-01-preview|2020-03-01|
|StreamAnalytics|Function|2016-03-01|2020-03-01|
|StreamAnalytics|Input|2016-03-01|2020-03-01|
|StreamAnalytics|listClusterStreamingJobs|2020-03-01-preview|2020-03-01|
|StreamAnalytics|Output|2016-03-01|2020-03-01|
|StreamAnalytics|PrivateEndpoint|2020-03-01-preview|2020-03-01|
|StreamAnalytics|StreamingJob|2016-03-01|2020-03-01|
|Subscription|Alias|2020-09-01|2021-10-01|
|Synapse|BigDataPool|2021-03-01|2021-06-01|
|Synapse|Database|2021-04-01-preview|2021-04-01-preview|
|Synapse|DatabasePrincipalAssignment|2021-04-01-preview|2021-04-01-preview|
|Synapse|DataConnection|2021-04-01-preview|2021-04-01-preview|
|Synapse|getIntegrationRuntimeConnectionInfo|2021-03-01|2021-06-01|
|Synapse|getIntegrationRuntimeObjectMetadatum|2021-03-01|2021-06-01|
|Synapse|getIntegrationRuntimeStatus|2021-03-01|2021-06-01|
|Synapse|IntegrationRuntime|2021-03-01|2021-06-01|
|Synapse|IpFirewallRule|2021-03-01|2021-06-01|
|Synapse|Key|2021-03-01|2021-06-01|
|Synapse|kustoPool|2021-04-01-preview|Renamed to KustoPool|
|Synapse|KustoPool|not present|2021-06-01-preview|
|Synapse|KustoPoolAttachedDatabaseConfiguration|2021-06-01-preview|2021-06-01-preview|
|Synapse|KustoPoolDatabase|not present|2021-06-01-preview|
|Synapse|KustoPoolDatabasePrincipalAssignment|not present|2021-06-01-preview|
|Synapse|KustoPoolDataConnection|not present|2021-06-01-preview|
|Synapse|KustoPoolPrincipalAssignment|2021-04-01-preview|2021-06-01-preview|
|Synapse|listIntegrationRuntimeAuthKey|2021-03-01|2021-06-01|
|Synapse|listKustoPoolFollowerDatabases|2021-06-01-preview|2021-06-01-preview|
|Synapse|listKustoPoolLanguageExtensions|2021-06-01-preview|2021-06-01-preview|
|Synapse|PrivateEndpointConnection|2021-03-01|2021-06-01|
|Synapse|PrivateLinkHub|2021-03-01|2021-06-01|
|Synapse|SqlPool|2021-03-01|2021-06-01|
|Synapse|SqlPoolSensitivityLabel|2021-03-01|2021-06-01|
|Synapse|SqlPoolTransparentDataEncryption|2021-03-01|2021-06-01|
|Synapse|SqlPoolVulnerabilityAssessment|2021-03-01|2021-06-01|
|Synapse|SqlPoolVulnerabilityAssessmentRuleBaseline|2021-03-01|2021-06-01|
|Synapse|SqlPoolWorkloadClassifier|2021-03-01|2021-06-01|
|Synapse|SqlPoolWorkloadGroup|2021-03-01|2021-06-01|
|Synapse|Workspace|2021-03-01|2021-06-01|
|Synapse|WorkspaceAadAdmin|2021-03-01|2021-06-01|
|Synapse|WorkspaceManagedSqlServerVulnerabilityAssessment|2021-03-01|2021-06-01|
|Synapse|WorkspaceSqlAadAdmin|2021-03-01|2021-06-01|
|Syntex|DocumentProcessor|2022-09-15-preview|2022-09-15-preview|
|TestBase|CustomerEvent|2022-04-01-preview|2022-04-01-preview|
|TestBase|FavoriteProcess|2022-04-01-preview|2022-04-01-preview|
|TestBase|getBillingHubServiceFreeHourBalance|2022-04-01-preview|2022-04-01-preview|
|TestBase|getBillingHubServiceUsage|2022-04-01-preview|2022-04-01-preview|
|TestBase|getPackageDownloadURL|2022-04-01-preview|2022-04-01-preview|
|TestBase|getTestBaseAccountFileUploadUrl|2022-04-01-preview|2022-04-01-preview|
|TestBase|getTestResultConsoleLogDownloadURL|2022-04-01-preview|2022-04-01-preview|
|TestBase|getTestResultDownloadURL|2022-04-01-preview|2022-04-01-preview|
|TestBase|getTestResultVideoDownloadURL|2022-04-01-preview|2022-04-01-preview|
|TestBase|Package|2022-04-01-preview|2022-04-01-preview|
|TestBase|TestBaseAccount|2022-04-01-preview|2022-04-01-preview|
|TimeSeriesInsights|AccessPolicy|2020-05-15|2020-05-15|
|TimeSeriesInsights|Environment|2020-05-15|2020-05-15|
|TimeSeriesInsights|EventSource|2020-05-15|2020-05-15|
|TimeSeriesInsights|PrivateEndpointConnection|2021-03-31-preview|2021-03-31-preview|
|TimeSeriesInsights|ReferenceDataSet|2020-05-15|2020-05-15|
|VideoAnalyzer|AccessPolicy|2021-05-01-preview|2021-11-01-preview|
|VideoAnalyzer|EdgeModule|2021-05-01-preview|2021-11-01-preview|
|VideoAnalyzer|listEdgeModuleProvisioningToken|2021-05-01-preview|2021-11-01-preview|
|VideoAnalyzer|listVideoContentToken|2021-11-01-preview|2021-11-01-preview|
|VideoAnalyzer|listVideoStreamingToken|2021-05-01-preview|Replaced with listVideoContentToken|
|VideoAnalyzer|LivePipeline|2021-11-01-preview|2021-11-01-preview|
|VideoAnalyzer|PipelineJob|2021-11-01-preview|2021-11-01-preview|
|VideoAnalyzer|PipelineTopology|2021-11-01-preview|2021-11-01-preview|
|VideoAnalyzer|PrivateEndpointConnection|2021-11-01-preview|2021-11-01-preview|
|VideoAnalyzer|Video|2021-05-01-preview|2021-11-01-preview|
|VideoAnalyzer|VideoAnalyzer|2021-05-01-preview|2021-11-01-preview|
|VideoIndexer|Account|2021-10-18-preview|2022-08-01|
|VirtualMachineImages|Trigger|not present|2022-07-01|
|VirtualMachineImages|VirtualMachineImageTemplate|2020-02-14|2022-07-01|
|VisualStudio|Account|2014-04-01-preview|2017-11-01-preview|
|VisualStudio|Extension|2014-04-01-preview|2017-11-01-preview|
|VMwareCloudSimple|DedicatedCloudNode|2019-04-01|2019-04-01|
|VMwareCloudSimple|DedicatedCloudService|2019-04-01|2019-04-01|
|VMwareCloudSimple|VirtualMachine|2019-04-01|2019-04-01|
|VoiceServices|CommunicationsGateway|2022-12-01-preview|2023-04-03|
|VoiceServices|Contact|2022-12-01-preview|2022-12-01-preview|
|VoiceServices|TestLine|2022-12-01-preview|2023-04-03|
|Web|AppServiceEnvironment|2020-12-01|2022-09-01|
|Web|AppServiceEnvironmentAseCustomDnsSuffixConfiguration|2022-03-01|2022-09-01|
|Web|AppServiceEnvironmentPrivateEndpointConnection|2020-12-01|2022-09-01|
|Web|AppServicePlan|2020-12-01|2022-09-01|
|Web|AppServicePlanRouteForVnet|2020-12-01|2022-09-01|
|Web|Certificate|2020-12-01|2022-09-01|
|Web|Connection|2016-06-01|2016-06-01|
|Web|ConnectionGateway|2016-06-01|2016-06-01|
|Web|CustomApi|2016-06-01|2016-06-01|
|Web|KubeEnvironment|2021-01-01|2022-09-01|
|Web|listAppServicePlanHybridConnectionKeys|2020-12-01|2022-09-01|
|Web|listConnectionConsentLinks|2016-06-01|2016-06-01|
|Web|listConnectionKeys|2015-08-01-preview|2015-08-01-preview|
|Web|listContainerAppSecrets|2021-03-01|This moved into the ContainerApp service|
|Web|listCustomApiWsdlInterfaces|2016-06-01|2016-06-01|
|Web|listSiteIdentifiersAssignedToHostName|2020-12-01|2022-09-01|
|Web|listStaticSiteAppSettings|2020-12-01|2022-09-01|
|Web|listStaticSiteBuildAppSettings|2020-12-01|2022-09-01|
|Web|listStaticSiteBuildFunctionAppSettings|2020-12-01|2022-09-01|
|Web|listStaticSiteConfiguredRoles|2020-12-01|2022-09-01|
|Web|listStaticSiteFunctionAppSettings|2020-12-01|2022-09-01|
|Web|listStaticSiteSecrets|2020-12-01|2022-09-01|
|Web|listStaticSiteUsers|2020-12-01|2022-09-01|
|Web|listWebAppApplicationSettings|2020-12-01|2022-09-01|
|Web|listWebAppApplicationSettingsSlot|2020-12-01|2022-09-01|
|Web|listWebAppAuthSettings|2020-12-01|2022-09-01|
|Web|listWebAppAuthSettingsSlot|2020-12-01|2022-09-01|
|Web|listWebAppAzureStorageAccounts|2020-12-01|2022-09-01|
|Web|listWebAppAzureStorageAccountsSlot|2020-12-01|2022-09-01|
|Web|listWebAppBackupConfiguration|2020-12-01|2022-09-01|
|Web|listWebAppBackupConfigurationSlot|2020-12-01|2022-09-01|
|Web|listWebAppBackupStatusSecrets|2020-12-01|2022-09-01|
|Web|listWebAppBackupStatusSecretsSlot|2020-12-01|2022-09-01|
|Web|listWebAppConnectionStrings|2020-12-01|2022-09-01|
|Web|listWebAppConnectionStringsSlot|2020-12-01|2022-09-01|
|Web|listWebAppFunctionKeys|2020-12-01|2022-09-01|
|Web|listWebAppFunctionKeysSlot|2020-12-01|2022-09-01|
|Web|listWebAppFunctionSecrets|2020-12-01|2022-09-01|
|Web|listWebAppFunctionSecretsSlot|2020-12-01|2022-09-01|
|Web|listWebAppHostKeys|2020-12-01|2022-09-01|
|Web|listWebAppHostKeysSlot|2020-12-01|2022-09-01|
|Web|listWebAppHybridConnectionKeys|2018-11-01|2018-11-01|
|Web|listWebAppHybridConnectionKeysSlot|2018-11-01|2018-11-01|
|Web|listWebAppMetadata|2020-12-01|2022-09-01|
|Web|listWebAppMetadataSlot|2020-12-01|2022-09-01|
|Web|listWebAppPublishingCredentials|2020-12-01|2022-09-01|
|Web|listWebAppPublishingCredentialsSlot|2020-12-01|2022-09-01|
|Web|listWebAppSiteBackups|2020-12-01|2022-09-01|
|Web|listWebAppSiteBackupsSlot|2020-12-01|2022-09-01|
|Web|listWebAppSitePushSettings|2020-12-01|2022-09-01|
|Web|listWebAppSitePushSettingsSlot|2020-12-01|2022-09-01|
|Web|listWebAppSyncFunctionTriggers|2020-12-01|2022-09-01|
|Web|listWebAppSyncFunctionTriggersSlot|2020-12-01|2022-09-01|
|Web|listWebAppSyncStatus|2020-12-01|2022-09-01|
|Web|listWebAppSyncStatusSlot|2020-12-01|2022-09-01|
|Web|listWebAppWorkflowsConnections|not present|2022-09-01|
|Web|listWebAppWorkflowsConnectionsSlot|not present|2022-09-01|
|Web|listWorkflowRunActionExpressionTraces|2022-03-01|2022-09-01|
|Web|listWorkflowRunActionRepetitionExpressionTraces|2022-03-01|2022-09-01|
|Web|listWorkflowTriggerCallbackUrl|2022-03-01|2022-09-01|
|Web|StaticSite|2020-12-01|2022-09-01|
|Web|StaticSiteBuildDatabaseConnection|not present|2022-09-01|
|Web|StaticSiteCustomDomain|2020-12-01|2022-09-01|
|Web|StaticSiteDatabaseConnection|not present|2022-09-01|
|Web|StaticSiteLinkedBackend|2022-03-01|2022-09-01|
|Web|StaticSiteLinkedBackendForBuild|2022-03-01|2022-09-01|
|Web|StaticSitePrivateEndpointConnection|2020-12-01|2022-09-01|
|Web|StaticSiteUserProvidedFunctionAppForStaticSite|2020-12-01|2022-09-01|
|Web|StaticSiteUserProvidedFunctionAppForStaticSiteBuild|2020-12-01|2022-09-01|
|Web|WebApp|2020-12-01|2022-09-01|
|Web|WebAppApplicationSettings|2020-12-01|2022-09-01|
|Web|WebAppApplicationSettingsSlot|2020-12-01|2022-09-01|
|Web|WebAppAuthSettings|2020-12-01|2022-09-01|
|Web|WebAppAuthSettingsSlot|2020-12-01|2022-09-01|
|Web|WebAppAuthSettingsV2|2020-12-01|2021-02-01|
|Web|WebAppAuthSettingsV2Slot|2020-12-01|2021-02-01|
|Web|WebAppAzureStorageAccounts|2020-12-01|2022-09-01|
|Web|WebAppAzureStorageAccountsSlot|2020-12-01|2022-09-01|
|Web|WebAppBackupConfiguration|2020-12-01|2022-09-01|
|Web|WebAppBackupConfigurationSlot|2020-12-01|2022-09-01|
|Web|WebAppConnectionStrings|2020-12-01|2022-09-01|
|Web|WebAppConnectionStringsSlot|2020-12-01|2022-09-01|
|Web|WebAppDeployment|2020-12-01|2022-09-01|
|Web|WebAppDeploymentSlot|2020-12-01|2022-09-01|
|Web|WebAppDiagnosticLogsConfiguration|2020-12-01|2022-09-01|
|Web|WebAppDomainOwnershipIdentifier|2020-12-01|2022-09-01|
|Web|WebAppDomainOwnershipIdentifierSlot|2020-12-01|2022-09-01|
|Web|WebAppFtpAllowed|not present|2022-09-01|
|Web|WebAppFunction|2020-12-01|2022-09-01|
|Web|WebAppHostNameBinding|2020-12-01|2022-09-01|
|Web|WebAppHostNameBindingSlot|2020-12-01|2022-09-01|
|Web|WebAppHybridConnection|2020-12-01|2022-09-01|
|Web|WebAppHybridConnectionSlot|2020-12-01|2022-09-01|
|Web|WebAppInstanceFunctionSlot|2020-12-01|2022-09-01|
|Web|WebAppMetadata|2020-12-01|2022-09-01|
|Web|WebAppMetadataSlot|2020-12-01|2022-09-01|
|Web|WebAppPremierAddOn|2020-12-01|2022-09-01|
|Web|WebAppPremierAddOnSlot|2020-12-01|2022-09-01|
|Web|WebAppPrivateEndpointConnection|2020-12-01|2022-09-01|
|Web|WebAppPrivateEndpointConnectionSlot|2020-12-01|2022-09-01|
|Web|WebAppPublicCertificate|2020-12-01|2022-09-01|
|Web|WebAppPublicCertificateSlot|2020-12-01|2022-09-01|
|Web|WebAppRelayServiceConnection|2020-12-01|2022-09-01|
|Web|WebAppRelayServiceConnectionSlot|2020-12-01|2022-09-01|
|Web|WebAppScmAllowed|not present|2022-09-01|
|Web|WebAppSiteExtension|2020-12-01|2022-09-01|
|Web|WebAppSiteExtensionSlot|2020-12-01|2022-09-01|
|Web|WebAppSitePushSettings|2020-12-01|2022-09-01|
|Web|WebAppSitePushSettingsSlot|2020-12-01|2022-09-01|
|Web|WebAppSlot|2020-12-01|2022-09-01|
|Web|WebAppSlotConfigurationNames|2020-12-01|2022-09-01|
|Web|WebAppSourceControl|2020-12-01|2022-09-01|
|Web|WebAppSourceControlSlot|2020-12-01|2022-09-01|
|Web|WebAppSwiftVirtualNetworkConnection|2020-10-01|2022-09-01|
|Web|WebAppSwiftVirtualNetworkConnectionSlot|2020-10-01|2022-09-01|
|Web|WebAppVnetConnection|2020-12-01|2022-09-01|
|Web|WebAppVnetConnectionSlot|2020-12-01|2022-09-01|
|WebPubSub|listWebPubSubKeys|2021-04-01-preview|2023-02-01|
|WebPubSub|WebPubSub|2021-04-01-preview|2023-02-01|
|WebPubSub|WebPubSubCustomCertificate|not present|2023-02-01|
|WebPubSub|WebPubSubCustomDomain|not present|2023-02-01|
|WebPubSub|WebPubSubHub|2021-10-01|2023-02-01|
|WebPubSub|WebPubSubPrivateEndpointConnection|2021-04-01-preview|2023-02-01|
|WebPubSub|WebPubSubReplica|not present|2023-03-01-preview|
|WebPubSub|WebPubSubSharedPrivateLinkResource|2021-04-01-preview|2023-02-01|
|WindowsESU|MultipleActivationKey|2019-09-16-preview|2019-09-16-preview|
|WindowsIoT|Service|2019-06-01|2019-06-01|
|Workloads|getSAPAvailabilityZoneDetails|2021-12-01-preview|2023-04-01|
|Workloads|getSAPDiskConfigurations|2021-12-01-preview|2023-04-01|
|Workloads|getSAPSizingRecommendations|2021-12-01-preview|2023-04-01|
|Workloads|getSAPSupportedSku|2021-12-01-preview|2023-04-01|
|Workloads|monitor|2021-12-01-preview|2023-04-01|
|Workloads|PhpWorkload|2021-12-01-preview|Deprecated by Azure|
|Workloads|ProviderInstance|2021-12-01-preview|2023-04-01|
|Workloads|SAPApplicationServerInstance|2021-12-01-preview|2023-04-01|
|Workloads|SAPCentralInstance|2021-12-01-preview|2023-04-01|
|Workloads|SAPDatabaseInstance|2021-12-01-preview|2023-04-01|
|Workloads|SapLandscapeMonitor|not present|2023-04-01|
|Workloads|SAPVirtualInstance|2021-12-01-preview|2023-04-01|
|Workloads|WordpressInstance|2021-12-01-preview|Deprecated by Azure|
