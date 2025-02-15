package installer

var generatedMigrationCRDV1 = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.1
  creationTimestamp: null
  name: migrations.schemas.schemahero.io
spec:
  group: schemas.schemahero.io
  names:
    kind: Migration
    listKind: MigrationList
    plural: migrations
    singular: migration
  scope: Namespaced
  versions:
  - name: v1alpha4
    schema:
      openAPIV3Schema:
        description: Migration is the Schema for the migrations API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: MigrationSpec defines the desired state of Migration
            properties:
              editedDDL:
                type: string
              generatedDDL:
                type: string
              tableName:
                type: string
              tableNamespace:
                type: string
            required:
            - tableName
            - tableNamespace
            type: object
          status:
            description: MigrationStatus defines the observed state of Migration
            properties:
              approvedAt:
                format: int64
                type: integer
              executedAt:
                format: int64
                type: integer
              invalidatedAt:
                description: InvalidatedAt is the unix nano timestamp when this plan was determined to be invalid or outdated
                format: int64
                type: integer
              plannedAt:
                description: PlannedAt is the unix nano timestamp when the plan was generated
                format: int64
                type: integer
              rejectedAt:
                format: int64
                type: integer
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`
