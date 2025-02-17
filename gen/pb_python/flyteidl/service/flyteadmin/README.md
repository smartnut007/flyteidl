# flyteadmin
No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

This Python package is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: version not set
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.PythonClientCodegen

## Requirements.

Python 2.7 and 3.4+

## Installation & Usage
### pip install

If the python package is hosted on Github, you can install directly from Github

```sh
pip install git+https://github.com/GIT_USER_ID/GIT_REPO_ID.git
```
(you may need to run `pip` with root permission: `sudo pip install git+https://github.com/GIT_USER_ID/GIT_REPO_ID.git`)

Then import the package:
```python
import flyteadmin 
```

### Setuptools

Install via [Setuptools](http://pypi.python.org/pypi/setuptools).

```sh
python setup.py install --user
```
(or `sudo python setup.py install` to install the package for all users)

Then import the package:
```python
import flyteadmin
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```python
from __future__ import print_function
import time
import flyteadmin
from flyteadmin.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = flyteadmin.AdminServiceApi(flyteadmin.ApiClient(configuration))
body = flyteadmin.AdminExecutionCreateRequest() # AdminExecutionCreateRequest | 

try:
    api_response = api_instance.create_execution(body)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AdminServiceApi->create_execution: %s\n" % e)

```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminServiceApi* | [**create_execution**](docs/AdminServiceApi.md#create_execution) | **POST** /api/v1/executions | 
*AdminServiceApi* | [**create_launch_plan**](docs/AdminServiceApi.md#create_launch_plan) | **POST** /api/v1/launch_plans | 
*AdminServiceApi* | [**create_node_event**](docs/AdminServiceApi.md#create_node_event) | **POST** /api/v1/events/nodes | 
*AdminServiceApi* | [**create_task**](docs/AdminServiceApi.md#create_task) | **POST** /api/v1/tasks | 
*AdminServiceApi* | [**create_task_event**](docs/AdminServiceApi.md#create_task_event) | **POST** /api/v1/events/tasks | 
*AdminServiceApi* | [**create_workflow**](docs/AdminServiceApi.md#create_workflow) | **POST** /api/v1/workflows | 
*AdminServiceApi* | [**create_workflow_event**](docs/AdminServiceApi.md#create_workflow_event) | **POST** /api/v1/events/workflows | 
*AdminServiceApi* | [**delete_project_domain_attributes**](docs/AdminServiceApi.md#delete_project_domain_attributes) | **DELETE** /api/v1/project_domain_attributes/{project}/{domain} | 
*AdminServiceApi* | [**delete_workflow_attributes**](docs/AdminServiceApi.md#delete_workflow_attributes) | **DELETE** /api/v1/workflow_attributes/{project}/{domain}/{workflow} | 
*AdminServiceApi* | [**get_active_launch_plan**](docs/AdminServiceApi.md#get_active_launch_plan) | **GET** /api/v1/active_launch_plans/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**get_execution**](docs/AdminServiceApi.md#get_execution) | **GET** /api/v1/executions/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**get_execution_data**](docs/AdminServiceApi.md#get_execution_data) | **GET** /api/v1/data/executions/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**get_launch_plan**](docs/AdminServiceApi.md#get_launch_plan) | **GET** /api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version} | 
*AdminServiceApi* | [**get_named_entity**](docs/AdminServiceApi.md#get_named_entity) | **GET** /api/v1/named_entities/{resource_type}/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**get_node_execution**](docs/AdminServiceApi.md#get_node_execution) | **GET** /api/v1/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id} | 
*AdminServiceApi* | [**get_node_execution_data**](docs/AdminServiceApi.md#get_node_execution_data) | **GET** /api/v1/data/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id} | 
*AdminServiceApi* | [**get_project_domain_attributes**](docs/AdminServiceApi.md#get_project_domain_attributes) | **GET** /api/v1/project_domain_attributes/{project}/{domain} | 
*AdminServiceApi* | [**get_task**](docs/AdminServiceApi.md#get_task) | **GET** /api/v1/tasks/{id.project}/{id.domain}/{id.name}/{id.version} | 
*AdminServiceApi* | [**get_task_execution**](docs/AdminServiceApi.md#get_task_execution) | **GET** /api/v1/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt} | 
*AdminServiceApi* | [**get_task_execution_data**](docs/AdminServiceApi.md#get_task_execution_data) | **GET** /api/v1/data/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt} | 
*AdminServiceApi* | [**get_version**](docs/AdminServiceApi.md#get_version) | **GET** /api/v1/version | 
*AdminServiceApi* | [**get_workflow**](docs/AdminServiceApi.md#get_workflow) | **GET** /api/v1/workflows/{id.project}/{id.domain}/{id.name}/{id.version} | 
*AdminServiceApi* | [**get_workflow_attributes**](docs/AdminServiceApi.md#get_workflow_attributes) | **GET** /api/v1/workflow_attributes/{project}/{domain}/{workflow} | 
*AdminServiceApi* | [**list_active_launch_plans**](docs/AdminServiceApi.md#list_active_launch_plans) | **GET** /api/v1/active_launch_plans/{project}/{domain} | 
*AdminServiceApi* | [**list_executions**](docs/AdminServiceApi.md#list_executions) | **GET** /api/v1/executions/{id.project}/{id.domain} | 
*AdminServiceApi* | [**list_launch_plan_ids**](docs/AdminServiceApi.md#list_launch_plan_ids) | **GET** /api/v1/launch_plan_ids/{project}/{domain} | 
*AdminServiceApi* | [**list_launch_plans**](docs/AdminServiceApi.md#list_launch_plans) | **GET** /api/v1/launch_plans/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**list_launch_plans2**](docs/AdminServiceApi.md#list_launch_plans2) | **GET** /api/v1/launch_plans/{id.project}/{id.domain} | 
*AdminServiceApi* | [**list_matchable_attributes**](docs/AdminServiceApi.md#list_matchable_attributes) | **GET** /api/v1/matchable_attributes | 
*AdminServiceApi* | [**list_named_entities**](docs/AdminServiceApi.md#list_named_entities) | **GET** /api/v1/named_entities/{resource_type}/{project}/{domain} | 
*AdminServiceApi* | [**list_node_executions**](docs/AdminServiceApi.md#list_node_executions) | **GET** /api/v1/node_executions/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name} | 
*AdminServiceApi* | [**list_node_executions_for_task**](docs/AdminServiceApi.md#list_node_executions_for_task) | **GET** /api/v1/children/task_executions/{task_execution_id.node_execution_id.execution_id.project}/{task_execution_id.node_execution_id.execution_id.domain}/{task_execution_id.node_execution_id.execution_id.name}/{task_execution_id.node_execution_id.node_id}/{task_execution_id.task_id.project}/{task_execution_id.task_id.domain}/{task_execution_id.task_id.name}/{task_execution_id.task_id.version}/{task_execution_id.retry_attempt} | 
*AdminServiceApi* | [**list_projects**](docs/AdminServiceApi.md#list_projects) | **GET** /api/v1/projects | 
*AdminServiceApi* | [**list_task_executions**](docs/AdminServiceApi.md#list_task_executions) | **GET** /api/v1/task_executions/{node_execution_id.execution_id.project}/{node_execution_id.execution_id.domain}/{node_execution_id.execution_id.name}/{node_execution_id.node_id} | 
*AdminServiceApi* | [**list_task_ids**](docs/AdminServiceApi.md#list_task_ids) | **GET** /api/v1/task_ids/{project}/{domain} | 
*AdminServiceApi* | [**list_tasks**](docs/AdminServiceApi.md#list_tasks) | **GET** /api/v1/tasks/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**list_tasks2**](docs/AdminServiceApi.md#list_tasks2) | **GET** /api/v1/tasks/{id.project}/{id.domain} | 
*AdminServiceApi* | [**list_workflow_ids**](docs/AdminServiceApi.md#list_workflow_ids) | **GET** /api/v1/workflow_ids/{project}/{domain} | 
*AdminServiceApi* | [**list_workflows**](docs/AdminServiceApi.md#list_workflows) | **GET** /api/v1/workflows/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**list_workflows2**](docs/AdminServiceApi.md#list_workflows2) | **GET** /api/v1/workflows/{id.project}/{id.domain} | 
*AdminServiceApi* | [**register_project**](docs/AdminServiceApi.md#register_project) | **POST** /api/v1/projects | 
*AdminServiceApi* | [**relaunch_execution**](docs/AdminServiceApi.md#relaunch_execution) | **POST** /api/v1/executions/relaunch | 
*AdminServiceApi* | [**terminate_execution**](docs/AdminServiceApi.md#terminate_execution) | **DELETE** /api/v1/executions/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**update_launch_plan**](docs/AdminServiceApi.md#update_launch_plan) | **PUT** /api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version} | 
*AdminServiceApi* | [**update_named_entity**](docs/AdminServiceApi.md#update_named_entity) | **PUT** /api/v1/named_entities/{resource_type}/{id.project}/{id.domain}/{id.name} | 
*AdminServiceApi* | [**update_project**](docs/AdminServiceApi.md#update_project) | **PUT** /api/v1/projects/{id} | flyteidl.admin.Project should be passed but the domains property should be empty; it will be ignored in the handler as domains cannot be updated via this API.
*AdminServiceApi* | [**update_project_domain_attributes**](docs/AdminServiceApi.md#update_project_domain_attributes) | **PUT** /api/v1/project_domain_attributes/{attributes.project}/{attributes.domain} | 
*AdminServiceApi* | [**update_workflow_attributes**](docs/AdminServiceApi.md#update_workflow_attributes) | **PUT** /api/v1/workflow_attributes/{attributes.project}/{attributes.domain}/{attributes.workflow} | 


## Documentation For Models

 - [AdminAbortMetadata](docs/AdminAbortMetadata.md)
 - [AdminAnnotations](docs/AdminAnnotations.md)
 - [AdminAuth](docs/AdminAuth.md)
 - [AdminAuthRole](docs/AdminAuthRole.md)
 - [AdminClusterResourceAttributes](docs/AdminClusterResourceAttributes.md)
 - [AdminCronSchedule](docs/AdminCronSchedule.md)
 - [AdminDomain](docs/AdminDomain.md)
 - [AdminEmailNotification](docs/AdminEmailNotification.md)
 - [AdminExecution](docs/AdminExecution.md)
 - [AdminExecutionClosure](docs/AdminExecutionClosure.md)
 - [AdminExecutionClusterLabel](docs/AdminExecutionClusterLabel.md)
 - [AdminExecutionCreateRequest](docs/AdminExecutionCreateRequest.md)
 - [AdminExecutionCreateResponse](docs/AdminExecutionCreateResponse.md)
 - [AdminExecutionList](docs/AdminExecutionList.md)
 - [AdminExecutionMetadata](docs/AdminExecutionMetadata.md)
 - [AdminExecutionQueueAttributes](docs/AdminExecutionQueueAttributes.md)
 - [AdminExecutionRelaunchRequest](docs/AdminExecutionRelaunchRequest.md)
 - [AdminExecutionSpec](docs/AdminExecutionSpec.md)
 - [AdminExecutionTerminateRequest](docs/AdminExecutionTerminateRequest.md)
 - [AdminExecutionTerminateResponse](docs/AdminExecutionTerminateResponse.md)
 - [AdminFixedRate](docs/AdminFixedRate.md)
 - [AdminFixedRateUnit](docs/AdminFixedRateUnit.md)
 - [AdminGetVersionResponse](docs/AdminGetVersionResponse.md)
 - [AdminLabels](docs/AdminLabels.md)
 - [AdminLaunchPlan](docs/AdminLaunchPlan.md)
 - [AdminLaunchPlanClosure](docs/AdminLaunchPlanClosure.md)
 - [AdminLaunchPlanCreateRequest](docs/AdminLaunchPlanCreateRequest.md)
 - [AdminLaunchPlanCreateResponse](docs/AdminLaunchPlanCreateResponse.md)
 - [AdminLaunchPlanList](docs/AdminLaunchPlanList.md)
 - [AdminLaunchPlanMetadata](docs/AdminLaunchPlanMetadata.md)
 - [AdminLaunchPlanSpec](docs/AdminLaunchPlanSpec.md)
 - [AdminLaunchPlanState](docs/AdminLaunchPlanState.md)
 - [AdminLaunchPlanUpdateRequest](docs/AdminLaunchPlanUpdateRequest.md)
 - [AdminLaunchPlanUpdateResponse](docs/AdminLaunchPlanUpdateResponse.md)
 - [AdminListMatchableAttributesResponse](docs/AdminListMatchableAttributesResponse.md)
 - [AdminLiteralMapBlob](docs/AdminLiteralMapBlob.md)
 - [AdminMatchableAttributesConfiguration](docs/AdminMatchableAttributesConfiguration.md)
 - [AdminMatchableResource](docs/AdminMatchableResource.md)
 - [AdminMatchingAttributes](docs/AdminMatchingAttributes.md)
 - [AdminNamedEntity](docs/AdminNamedEntity.md)
 - [AdminNamedEntityIdentifier](docs/AdminNamedEntityIdentifier.md)
 - [AdminNamedEntityIdentifierList](docs/AdminNamedEntityIdentifierList.md)
 - [AdminNamedEntityList](docs/AdminNamedEntityList.md)
 - [AdminNamedEntityMetadata](docs/AdminNamedEntityMetadata.md)
 - [AdminNamedEntityState](docs/AdminNamedEntityState.md)
 - [AdminNamedEntityUpdateRequest](docs/AdminNamedEntityUpdateRequest.md)
 - [AdminNamedEntityUpdateResponse](docs/AdminNamedEntityUpdateResponse.md)
 - [AdminNodeExecutionClosure](docs/AdminNodeExecutionClosure.md)
 - [AdminNodeExecutionEventRequest](docs/AdminNodeExecutionEventRequest.md)
 - [AdminNodeExecutionEventResponse](docs/AdminNodeExecutionEventResponse.md)
 - [AdminNodeExecutionGetDataResponse](docs/AdminNodeExecutionGetDataResponse.md)
 - [AdminNodeExecutionList](docs/AdminNodeExecutionList.md)
 - [AdminNodeExecutionMetaData](docs/AdminNodeExecutionMetaData.md)
 - [AdminNotification](docs/AdminNotification.md)
 - [AdminNotificationList](docs/AdminNotificationList.md)
 - [AdminPagerDutyNotification](docs/AdminPagerDutyNotification.md)
 - [AdminPluginOverride](docs/AdminPluginOverride.md)
 - [AdminPluginOverrides](docs/AdminPluginOverrides.md)
 - [AdminProject](docs/AdminProject.md)
 - [AdminProjectDomainAttributes](docs/AdminProjectDomainAttributes.md)
 - [AdminProjectDomainAttributesDeleteRequest](docs/AdminProjectDomainAttributesDeleteRequest.md)
 - [AdminProjectDomainAttributesDeleteResponse](docs/AdminProjectDomainAttributesDeleteResponse.md)
 - [AdminProjectDomainAttributesGetResponse](docs/AdminProjectDomainAttributesGetResponse.md)
 - [AdminProjectDomainAttributesUpdateRequest](docs/AdminProjectDomainAttributesUpdateRequest.md)
 - [AdminProjectDomainAttributesUpdateResponse](docs/AdminProjectDomainAttributesUpdateResponse.md)
 - [AdminProjectRegisterRequest](docs/AdminProjectRegisterRequest.md)
 - [AdminProjectRegisterResponse](docs/AdminProjectRegisterResponse.md)
 - [AdminProjectUpdateResponse](docs/AdminProjectUpdateResponse.md)
 - [AdminProjects](docs/AdminProjects.md)
 - [AdminRawOutputDataConfig](docs/AdminRawOutputDataConfig.md)
 - [AdminSchedule](docs/AdminSchedule.md)
 - [AdminSlackNotification](docs/AdminSlackNotification.md)
 - [AdminSort](docs/AdminSort.md)
 - [AdminSystemMetadata](docs/AdminSystemMetadata.md)
 - [AdminTask](docs/AdminTask.md)
 - [AdminTaskClosure](docs/AdminTaskClosure.md)
 - [AdminTaskCreateRequest](docs/AdminTaskCreateRequest.md)
 - [AdminTaskCreateResponse](docs/AdminTaskCreateResponse.md)
 - [AdminTaskExecutionClosure](docs/AdminTaskExecutionClosure.md)
 - [AdminTaskExecutionEventRequest](docs/AdminTaskExecutionEventRequest.md)
 - [AdminTaskExecutionEventResponse](docs/AdminTaskExecutionEventResponse.md)
 - [AdminTaskExecutionGetDataResponse](docs/AdminTaskExecutionGetDataResponse.md)
 - [AdminTaskExecutionList](docs/AdminTaskExecutionList.md)
 - [AdminTaskList](docs/AdminTaskList.md)
 - [AdminTaskResourceAttributes](docs/AdminTaskResourceAttributes.md)
 - [AdminTaskResourceSpec](docs/AdminTaskResourceSpec.md)
 - [AdminTaskSpec](docs/AdminTaskSpec.md)
 - [AdminUrlBlob](docs/AdminUrlBlob.md)
 - [AdminVersion](docs/AdminVersion.md)
 - [AdminWorkflow](docs/AdminWorkflow.md)
 - [AdminWorkflowAttributes](docs/AdminWorkflowAttributes.md)
 - [AdminWorkflowAttributesDeleteRequest](docs/AdminWorkflowAttributesDeleteRequest.md)
 - [AdminWorkflowAttributesDeleteResponse](docs/AdminWorkflowAttributesDeleteResponse.md)
 - [AdminWorkflowAttributesGetResponse](docs/AdminWorkflowAttributesGetResponse.md)
 - [AdminWorkflowAttributesUpdateRequest](docs/AdminWorkflowAttributesUpdateRequest.md)
 - [AdminWorkflowAttributesUpdateResponse](docs/AdminWorkflowAttributesUpdateResponse.md)
 - [AdminWorkflowClosure](docs/AdminWorkflowClosure.md)
 - [AdminWorkflowCreateRequest](docs/AdminWorkflowCreateRequest.md)
 - [AdminWorkflowCreateResponse](docs/AdminWorkflowCreateResponse.md)
 - [AdminWorkflowExecutionEventRequest](docs/AdminWorkflowExecutionEventRequest.md)
 - [AdminWorkflowExecutionEventResponse](docs/AdminWorkflowExecutionEventResponse.md)
 - [AdminWorkflowExecutionGetDataResponse](docs/AdminWorkflowExecutionGetDataResponse.md)
 - [AdminWorkflowList](docs/AdminWorkflowList.md)
 - [AdminWorkflowSpec](docs/AdminWorkflowSpec.md)
 - [BlobTypeBlobDimensionality](docs/BlobTypeBlobDimensionality.md)
 - [ComparisonExpressionOperator](docs/ComparisonExpressionOperator.md)
 - [ConjunctionExpressionLogicalOperator](docs/ConjunctionExpressionLogicalOperator.md)
 - [ConnectionSetIdList](docs/ConnectionSetIdList.md)
 - [CoreAlias](docs/CoreAlias.md)
 - [CoreBinary](docs/CoreBinary.md)
 - [CoreBinding](docs/CoreBinding.md)
 - [CoreBindingData](docs/CoreBindingData.md)
 - [CoreBindingDataCollection](docs/CoreBindingDataCollection.md)
 - [CoreBindingDataMap](docs/CoreBindingDataMap.md)
 - [CoreBlob](docs/CoreBlob.md)
 - [CoreBlobMetadata](docs/CoreBlobMetadata.md)
 - [CoreBlobType](docs/CoreBlobType.md)
 - [CoreBooleanExpression](docs/CoreBooleanExpression.md)
 - [CoreBranchNode](docs/CoreBranchNode.md)
 - [CoreCatalogArtifactTag](docs/CoreCatalogArtifactTag.md)
 - [CoreCatalogCacheStatus](docs/CoreCatalogCacheStatus.md)
 - [CoreCatalogMetadata](docs/CoreCatalogMetadata.md)
 - [CoreComparisonExpression](docs/CoreComparisonExpression.md)
 - [CoreCompiledTask](docs/CoreCompiledTask.md)
 - [CoreCompiledWorkflow](docs/CoreCompiledWorkflow.md)
 - [CoreCompiledWorkflowClosure](docs/CoreCompiledWorkflowClosure.md)
 - [CoreConjunctionExpression](docs/CoreConjunctionExpression.md)
 - [CoreConnectionSet](docs/CoreConnectionSet.md)
 - [CoreContainer](docs/CoreContainer.md)
 - [CoreContainerPort](docs/CoreContainerPort.md)
 - [CoreDataLoadingConfig](docs/CoreDataLoadingConfig.md)
 - [CoreError](docs/CoreError.md)
 - [CoreExecutionError](docs/CoreExecutionError.md)
 - [CoreIOStrategy](docs/CoreIOStrategy.md)
 - [CoreIdentifier](docs/CoreIdentifier.md)
 - [CoreIdentity](docs/CoreIdentity.md)
 - [CoreIfBlock](docs/CoreIfBlock.md)
 - [CoreIfElseBlock](docs/CoreIfElseBlock.md)
 - [CoreKeyValuePair](docs/CoreKeyValuePair.md)
 - [CoreLiteral](docs/CoreLiteral.md)
 - [CoreLiteralCollection](docs/CoreLiteralCollection.md)
 - [CoreLiteralMap](docs/CoreLiteralMap.md)
 - [CoreLiteralType](docs/CoreLiteralType.md)
 - [CoreNode](docs/CoreNode.md)
 - [CoreNodeExecutionIdentifier](docs/CoreNodeExecutionIdentifier.md)
 - [CoreNodeExecutionPhase](docs/CoreNodeExecutionPhase.md)
 - [CoreNodeMetadata](docs/CoreNodeMetadata.md)
 - [CoreOAuth2Client](docs/CoreOAuth2Client.md)
 - [CoreOAuth2TokenRequest](docs/CoreOAuth2TokenRequest.md)
 - [CoreOAuth2TokenRequestType](docs/CoreOAuth2TokenRequestType.md)
 - [CoreOperand](docs/CoreOperand.md)
 - [CoreOutputReference](docs/CoreOutputReference.md)
 - [CoreParameter](docs/CoreParameter.md)
 - [CoreParameterMap](docs/CoreParameterMap.md)
 - [CorePrimitive](docs/CorePrimitive.md)
 - [CoreQualityOfService](docs/CoreQualityOfService.md)
 - [CoreQualityOfServiceSpec](docs/CoreQualityOfServiceSpec.md)
 - [CoreResourceType](docs/CoreResourceType.md)
 - [CoreResources](docs/CoreResources.md)
 - [CoreRetryStrategy](docs/CoreRetryStrategy.md)
 - [CoreRuntimeMetadata](docs/CoreRuntimeMetadata.md)
 - [CoreScalar](docs/CoreScalar.md)
 - [CoreSchemaType](docs/CoreSchemaType.md)
 - [CoreSecret](docs/CoreSecret.md)
 - [CoreSecurityContext](docs/CoreSecurityContext.md)
 - [CoreSimpleType](docs/CoreSimpleType.md)
 - [CoreTaskExecutionIdentifier](docs/CoreTaskExecutionIdentifier.md)
 - [CoreTaskExecutionPhase](docs/CoreTaskExecutionPhase.md)
 - [CoreTaskLog](docs/CoreTaskLog.md)
 - [CoreTaskMetadata](docs/CoreTaskMetadata.md)
 - [CoreTaskNode](docs/CoreTaskNode.md)
 - [CoreTaskTemplate](docs/CoreTaskTemplate.md)
 - [CoreTypedInterface](docs/CoreTypedInterface.md)
 - [CoreVariable](docs/CoreVariable.md)
 - [CoreVariableMap](docs/CoreVariableMap.md)
 - [CoreVoid](docs/CoreVoid.md)
 - [CoreWorkflowExecutionIdentifier](docs/CoreWorkflowExecutionIdentifier.md)
 - [CoreWorkflowExecutionPhase](docs/CoreWorkflowExecutionPhase.md)
 - [CoreWorkflowMetadata](docs/CoreWorkflowMetadata.md)
 - [CoreWorkflowMetadataDefaults](docs/CoreWorkflowMetadataDefaults.md)
 - [CoreWorkflowNode](docs/CoreWorkflowNode.md)
 - [CoreWorkflowTemplate](docs/CoreWorkflowTemplate.md)
 - [DataLoadingConfigLiteralMapFormat](docs/DataLoadingConfigLiteralMapFormat.md)
 - [EventExternalResourceInfo](docs/EventExternalResourceInfo.md)
 - [EventNodeExecutionEvent](docs/EventNodeExecutionEvent.md)
 - [EventParentNodeExecutionMetadata](docs/EventParentNodeExecutionMetadata.md)
 - [EventParentTaskExecutionMetadata](docs/EventParentTaskExecutionMetadata.md)
 - [EventResourcePoolInfo](docs/EventResourcePoolInfo.md)
 - [EventTaskExecutionEvent](docs/EventTaskExecutionEvent.md)
 - [EventTaskExecutionMetadata](docs/EventTaskExecutionMetadata.md)
 - [EventWorkflowExecutionEvent](docs/EventWorkflowExecutionEvent.md)
 - [ExecutionErrorErrorKind](docs/ExecutionErrorErrorKind.md)
 - [ExecutionMetadataExecutionMode](docs/ExecutionMetadataExecutionMode.md)
 - [FlyteidladminDynamicWorkflowNodeMetadata](docs/FlyteidladminDynamicWorkflowNodeMetadata.md)
 - [FlyteidladminNodeExecution](docs/FlyteidladminNodeExecution.md)
 - [FlyteidladminTaskExecution](docs/FlyteidladminTaskExecution.md)
 - [FlyteidladminTaskNodeMetadata](docs/FlyteidladminTaskNodeMetadata.md)
 - [FlyteidladminWorkflowNodeMetadata](docs/FlyteidladminWorkflowNodeMetadata.md)
 - [FlyteidlcoreSchema](docs/FlyteidlcoreSchema.md)
 - [FlyteidleventDynamicWorkflowNodeMetadata](docs/FlyteidleventDynamicWorkflowNodeMetadata.md)
 - [FlyteidleventTaskNodeMetadata](docs/FlyteidleventTaskNodeMetadata.md)
 - [FlyteidleventWorkflowNodeMetadata](docs/FlyteidleventWorkflowNodeMetadata.md)
 - [IOStrategyDownloadMode](docs/IOStrategyDownloadMode.md)
 - [IOStrategyUploadMode](docs/IOStrategyUploadMode.md)
 - [PluginOverrideMissingPluginBehavior](docs/PluginOverrideMissingPluginBehavior.md)
 - [ProjectProjectState](docs/ProjectProjectState.md)
 - [ProtobufListValue](docs/ProtobufListValue.md)
 - [ProtobufNullValue](docs/ProtobufNullValue.md)
 - [ProtobufStruct](docs/ProtobufStruct.md)
 - [ProtobufValue](docs/ProtobufValue.md)
 - [QualityOfServiceTier](docs/QualityOfServiceTier.md)
 - [ResourcesResourceEntry](docs/ResourcesResourceEntry.md)
 - [ResourcesResourceName](docs/ResourcesResourceName.md)
 - [RuntimeMetadataRuntimeType](docs/RuntimeMetadataRuntimeType.md)
 - [SchemaColumnSchemaColumnType](docs/SchemaColumnSchemaColumnType.md)
 - [SchemaTypeSchemaColumn](docs/SchemaTypeSchemaColumn.md)
 - [SecretMountType](docs/SecretMountType.md)
 - [SortDirection](docs/SortDirection.md)
 - [TaskExecutionMetadataInstanceClass](docs/TaskExecutionMetadataInstanceClass.md)
 - [TaskLogMessageFormat](docs/TaskLogMessageFormat.md)
 - [WorkflowMetadataOnFailurePolicy](docs/WorkflowMetadataOnFailurePolicy.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author



