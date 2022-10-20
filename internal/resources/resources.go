package resources

var TestAntiAffinityRule = newYamlFile("testData/testAntiAffinityRule.yaml")
var PluginsInitContainer = newYamlFile("testData/pluginsInitContainer.yaml")
var JsonLogging = newYamlFile("testData/jsonLogging.yaml")
var AcceptLicenseAgreementBoolYes = newYamlFile("testData/acceptLicenseAgreementBoolYes.yaml")
var AcceptLicenseAgreementBoolTrue = newYamlFile("testData/acceptLicenseAgreementBoolTrue.yaml")
var AcceptLicenseAgreement = newYamlFile("testData/acceptLicenseAgreement.yaml")
var ApocCorePlugin = newYamlFile("testData/apocCorePlugin.yaml")
var CsvMetrics = newYamlFile("testData/csvMetrics.yaml")
var DefaultStorageClass = newYamlFile("testData/defaultStorageClass.yaml")
var JvmAdditionalSettings = newYamlFile("testData/jvmAdditionalSettings.yaml")
var BoolsInConfig = newYamlFile("testData/boolsInConfig.yaml")
var IntsInConfig = newYamlFile("testData/intsInConfig.yaml")
var ChmodInitContainer = newYamlFile("testData/chmodInitContainer.yaml")
var ChmodInitContainerAndCustomInitContainer = newYamlFile("testData/chmodInitContainerAndCustomInitContainer.yaml")
var ReadReplicaUpstreamStrategy = newYamlFile("testData/read_replica_upstream_selection_strategy.yaml")
var ExcludeLoadBalancer = newYamlFile("testData/excludeLoadBalancer.yaml")
var EmptyImageCredentials = newYamlFile("testData/imagePullSecret/emptyImageCreds.yaml")
var DuplicateImageCredentials = newYamlFile("testData/imagePullSecret/duplicateImageCreds.yaml")
var MissingImageCredentials = newYamlFile("testData/imagePullSecret/missingImageCreds.yaml")
var EmptyImagePullSecrets = newYamlFile("testData/imagePullSecret/emptyImagePullSecrets.yaml")
var InvalidNodeSelectorLabels = newYamlFile("testData/nodeselector.yaml")
var ApocConfig = newYamlFile("testData/apocConfig.yaml")
var ApocClusterTestConfig = newYamlFile("testData/apocClusterTest.yaml")
var PodSpecAnnotations = newYamlFile("testData/podSpecAnnotations.yaml")
var StatefulSetAnnotations = newYamlFile("testData/statefulSetAnnotations.yaml")
var PriorityClassName = newYamlFile("testData/priorityClassName.yaml")
var AdditionalVolumesAndMounts = newYamlFile("testData/additionalVolumes.yaml")
var Tolerations = newYamlFile("testData/tolerations.yaml")
var NodeAffinity = newYamlFile("testData/nodeAffinity.yaml")
