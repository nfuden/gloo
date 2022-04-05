export const petstoreYaml = `apiVersion: graphql.gloo.solo.io/v1beta1
kind: GraphQLApi
metadata:
  creationTimestamp: "2022-01-25T19:01:30Z"
  generation: 1
  managedFields:
  - apiVersion: graphql.gloo.solo.io/v1beta1
    fieldsType: FieldsV1
    fieldsV1:
      f:spec:
        .: {}
        f:executableSchema:
          .: {}
          f:executor:
            .: {}
            f:local:
              .: {}
              f:enableIntrospection: {}
              f:resolutions:
                .: {}
                f:default-petstore-8080_gloo-system|Mutation|addPet:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:body:
                        .: {}
                        f:category: {}
                        f:id: {}
                        f:name: {}
                        f:photoUrls: {}
                        f:status: {}
                        f:tags: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|createUser:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:body:
                        .: {}
                        f:email: {}
                        f:firstName: {}
                        f:id: {}
                        f:lastName: {}
                        f:password: {}
                        f:phone: {}
                        f:userStatus: {}
                        f:username: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|createUsersWithArrayInput:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:body: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|createUsersWithListInput:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:body: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|deleteOrder:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|deletePet:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                        f:api_key: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|deleteUser:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|placeOrder:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:body:
                        .: {}
                        f:complete: {}
                        f:id: {}
                        f:petId: {}
                        f:quantity: {}
                        f:shipDate: {}
                        f:status: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|updatePet:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:body:
                        .: {}
                        f:category: {}
                        f:id: {}
                        f:name: {}
                        f:photoUrls: {}
                        f:status: {}
                        f:tags: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|updatePetWithForm:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|updateUser:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:body:
                        .: {}
                        f:email: {}
                        f:firstName: {}
                        f:id: {}
                        f:lastName: {}
                        f:password: {}
                        f:phone: {}
                        f:userStatus: {}
                        f:username: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Mutation|uploadFile:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|findPetsByStatus:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                      f:queryParams:
                        .: {}
                        f:status: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|findPetsByTags:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                      f:queryParams:
                        .: {}
                        f:tags: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|getInventory:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|getOrderById:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|getPetById:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|getUserByName:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|loginUser:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                      f:queryParams:
                        .: {}
                        f:password: {}
                        f:username: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
                f:default-petstore-8080_gloo-system|Query|logoutUser:
                  .: {}
                  f:restResolver:
                    .: {}
                    f:request:
                      .: {}
                      f:headers:
                        .: {}
                        f::method: {}
                        f::path: {}
                    f:upstreamRef:
                      .: {}
                      f:name: {}
                      f:namespace: {}
          f:schemaDefinition: {}
    manager: discovery
    operation: Update
    time: "2022-01-25T19:01:30Z"
  name: default-petstore-8080
  namespace: gloo-system
  resourceVersion: "29317"
  uid: 97a4963c-8d73-4f5f-bf30-28aa4c076527
spec:
  executableSchema:
    executor:
      local:
        enableIntrospection: true
        resolutions:
          default-petstore-8080_gloo-system|Mutation|addPet:
            restResolver:
              request:
                body:
                  category: '{$args.petInput.category}'
                  id: '{$args.petInput.id}'
                  name: '{$args.petInput.name}'
                  photoUrls: '{$args.petInput.photoUrls}'
                  status: '{$args.petInput.status}'
                  tags: '{$args.petInput.tags}'
                headers:
                  :method: POST
                  :path: /v3/pet
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|createUser:
            restResolver:
              request:
                body:
                  email: '{$args.userInput.email}'
                  firstName: '{$args.userInput.firstName}'
                  id: '{$args.userInput.id}'
                  lastName: '{$args.userInput.lastName}'
                  password: '{$args.userInput.password}'
                  phone: '{$args.userInput.phone}'
                  userStatus: '{$args.userInput.userStatus}'
                  username: '{$args.userInput.username}'
                headers:
                  :method: POST
                  :path: /v3/user
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|createUsersWithArrayInput:
            restResolver:
              request:
                body: '{$args.userCreateWithArrayInput}'
                headers:
                  :method: POST
                  :path: /v3/user/createWithArray
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|createUsersWithListInput:
            restResolver:
              request:
                body: '{$args.userCreateWithListInput}'
                headers:
                  :method: POST
                  :path: /v3/user/createWithList
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|deleteOrder:
            restResolver:
              request:
                headers:
                  :method: DELETE
                  :path: /v3/store/order/{$args.orderId}
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|deletePet:
            restResolver:
              request:
                headers:
                  :method: DELETE
                  :path: /v3/pet/{$args.apiKey}
                  api_key: '{$args.apiKey}'
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|deleteUser:
            restResolver:
              request:
                headers:
                  :method: DELETE
                  :path: /v3/user/{$args.username}
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|placeOrder:
            restResolver:
              request:
                body:
                  complete: '{$args.orderInput.complete}'
                  id: '{$args.orderInput.id}'
                  petId: '{$args.orderInput.petId}'
                  quantity: '{$args.orderInput.quantity}'
                  shipDate: '{$args.orderInput.shipDate}'
                  status: '{$args.orderInput.status}'
                headers:
                  :method: POST
                  :path: /v3/store/order
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|updatePet:
            restResolver:
              request:
                body:
                  category: '{$args.petInput.category}'
                  id: '{$args.petInput.id}'
                  name: '{$args.petInput.name}'
                  photoUrls: '{$args.petInput.photoUrls}'
                  status: '{$args.petInput.status}'
                  tags: '{$args.petInput.tags}'
                headers:
                  :method: PUT
                  :path: /v3/pet
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|updatePetWithForm:
            restResolver:
              request:
                headers:
                  :method: POST
                  :path: /v3/pet/{$args.petId}
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|updateUser:
            restResolver:
              request:
                body:
                  email: '{$args.userInput.email}'
                  firstName: '{$args.userInput.firstName}'
                  id: '{$args.userInput.id}'
                  lastName: '{$args.userInput.lastName}'
                  password: '{$args.userInput.password}'
                  phone: '{$args.userInput.phone}'
                  userStatus: '{$args.userInput.userStatus}'
                  username: '{$args.userInput.username}'
                headers:
                  :method: PUT
                  :path: /v3/user/{$args.username}
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Mutation|uploadFile:
            restResolver:
              request:
                headers:
                  :method: POST
                  :path: /v3/pet/{$args.petId}/uploadImage
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|findPetsByStatus:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/pet/findByStatus
                queryParams:
                  status: '{$args.status}'
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|findPetsByTags:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/pet/findByTags
                queryParams:
                  tags: '{$args.tags}'
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|getInventory:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/store/inventory
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|getOrderById:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/store/order/{$args.orderId}
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|getPetById:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/pet/{$args.petId}
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|getUserByName:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/user/{$args.username}
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|loginUser:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/user/login
                queryParams:
                  password: '{$args.username}'
                  username: '{$args.username}'
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
          default-petstore-8080_gloo-system|Query|logoutUser:
            restResolver:
              request:
                headers:
                  :method: GET
                  :path: /v3/user/logout
              upstreamRef:
                name: default-petstore-8080
                namespace: gloo-system
    schemaDefinition: "\"\"\"Marks an element of a GraphQL schema as no longer supported.\"\"\"\ndirective
      @deprecated(\n  \n  \"\"\"Explains why this element was deprecated, usually
      also including a suggestion for how to access supported similar data. Formattedin
      [Markdown](https://daringfireball.net/projects/markdown/).\"\"\"\n  reason:
      String = \"No longer supported\"\n) on FIELD_DEFINITION | ENUM_VALUE\n\ntype
      Mutation {\n  \n  \"\"\"\n  Place an order for a pet\n  \n  Equivalent to OpenApiSpec
      'OpenAPI Petstore' POST /store/order\n  \"\"\"\n  placeOrder(\n    \n    \"\"\"An
      order for a pets from the pet store\"\"\"\n    orderInput: OrderInput!\n  ):
      Order @resolve(name: \"default-petstore-8080_gloo-system|Mutation|placeOrder\")\n
      \ \n  \"\"\"\n  Deletes a pet\n  \n  Equivalent to OpenApiSpec 'OpenAPI Petstore'
      DELETE /pet/{petId}\n  \"\"\"\n  deletePet(\n    apiKey: String\n    \n    \"\"\"Pet
      id to delete\"\"\"\n    petId: Int!\n  ): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|deletePet\")\n
      \ \n  \"\"\"\n  For valid response try integer IDs with value < 1000. Anything
      above 1000 or nonintegers will generate API errors\n  \n  Equivalent to OpenApiSpec
      'OpenAPI Petstore' DELETE /store/order/{orderId}\n  \"\"\"\n  deleteOrder(\n
      \   \n    \"\"\"ID of the order that needs to be deleted\"\"\"\n    orderId:
      String!\n  ): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|deleteOrder\")\n
      \ \n  \"\"\"\n  Creates list of users with given input array\n  \n  Equivalent
      to OpenApiSpec 'OpenAPI Petstore' POST /user/createWithArray\n  \"\"\"\n  createUsersWithArrayInput(userCreateWithArrayInput:
      [UserInput]!): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|createUsersWithArrayInput\")\n
      \ \n  \"\"\"\n  Creates list of users with given input array\n  \n  Equivalent
      to OpenApiSpec 'OpenAPI Petstore' POST /user/createWithList\n  \"\"\"\n  createUsersWithListInput(userCreateWithListInput:
      [UserInput]!): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|createUsersWithListInput\")\n
      \ \n  \"\"\"\n  This can only be done by the logged in user.\n  \n  Equivalent
      to OpenApiSpec 'OpenAPI Petstore' PUT /user/{username}\n  \"\"\"\n  updateUser(\n
      \   \n    \"\"\"name that need to be deleted\"\"\"\n    username: String!\n
      \   \n    \"\"\"A User who is purchasing from the pet store\"\"\"\n    userInput:
      UserInput!\n  ): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|updateUser\")\n
      \ \n  \"\"\"\n  uploads an image\n  \n  Equivalent to OpenApiSpec 'OpenAPI Petstore'
      POST /pet/{petId}/uploadImage\n  \"\"\"\n  uploadFile(\n    \n    \"\"\"ID of
      pet to update\"\"\"\n    petId: Int!\n    \n    \"\"\"String represents payload
      of content type 'multipart/form-data'\"\"\"\n    multipartFormDataInput: String\n
      \ ): ApiResponse @resolve(name: \"default-petstore-8080_gloo-system|Mutation|uploadFile\")\n
      \ \n  \"\"\"\n  This can only be done by the logged in user.\n  \n  Equivalent
      to OpenApiSpec 'OpenAPI Petstore' DELETE /user/{username}\n  \"\"\"\n  deleteUser(\n
      \   \n    \"\"\"The name that needs to be deleted\"\"\"\n    username: String!\n
      \ ): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|deleteUser\")\n
      \ \n  \"\"\"\n  Updates a pet in the store with form data\n  \n  Equivalent
      to OpenApiSpec 'OpenAPI Petstore' POST /pet/{petId}\n  \"\"\"\n  updatePetWithForm(\n
      \   \n    \"\"\"ID of pet that needs to be updated\"\"\"\n    petId: Int!\n
      \   \n    \"\"\"String represents payload of content type 'application/x-www-form-urlencoded'\"\"\"\n
      \   applicationXWwwFormUrlencodedInput: String\n  ): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|updatePetWithForm\")\n
      \ \n  \"\"\"\n  Update an existing pet\n  \n  Equivalent to OpenApiSpec 'OpenAPI
      Petstore' PUT /pet\n  \"\"\"\n  updatePet(\n    \n    \"\"\"A pet for sale in
      the pet store\"\"\"\n    petInput: PetInput!\n  ): JSON @resolve(name: \"default-petstore-8080_gloo-system|Mutation|updatePet\")\n
      \ \n  \"\"\"\n  Add a new pet to the store\n  \n  Equivalent to OpenApiSpec
      'OpenAPI Petstore' POST /pet\n  \"\"\"\n  addPet(\n    \n    \"\"\"A pet for
      sale in the pet store\"\"\"\n    petInput: PetInput!\n  ): JSON @resolve(name:
      \"default-petstore-8080_gloo-system|Mutation|addPet\")\n  \n  \"\"\"\n  This
      can only be done by the logged in user.\n  \n  Equivalent to OpenApiSpec 'OpenAPI
      Petstore' POST /user\n  \"\"\"\n  createUser(\n    \n    \"\"\"A User who is
      purchasing from the pet store\"\"\"\n    userInput: UserInput!\n  ): JSON @resolve(name:
      \"default-petstore-8080_gloo-system|Mutation|createUser\")\n}\n\nenum Status3ListItem
      {\n  sold\n  available\n  pending\n}\n\ntype ApiResponse {\n  code: Int\n  message:
      String\n  type: String\n}\n\ntype User {\n  password: String\n  phone: String\n
      \ \n  \"\"\"User Status\"\"\"\n  userStatus: Int\n  username: String\n  email:
      String\n  firstName: String\n  id: Int\n  lastName: String\n}\n\nenum Status
      {\n  available\n  pending\n  sold\n}\n\n\"\"\"The JSON scalar type represents
      JSON values as specified by [ECMA-404](http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf).\"\"\"\nscalar
      JSON\n\ntype Order {\n  petId: Int\n  quantity: Int\n  shipDate: String\n  \n
      \ \"\"\"Order Status\"\"\"\n  status: Status2\n  complete: Boolean\n  id: Int\n}\n\n\"\"\"A
      pet for sale in the pet store\"\"\"\ninput PetInput {\n  \n  \"\"\"pet status
      in the store\"\"\"\n  status: Status\n  tags: [TagInput]\n  \n  \"\"\"A category
      for a pet\"\"\"\n  category: CategoryInput\n  id: Int\n  name: String!\n  photoUrls:
      [String]!\n}\n\ntype Query {\n  \n  \"\"\"\n  Returns a map of status codes
      to quantities\n  \n  Equivalent to OpenApiSpec 'OpenAPI Petstore' GET /store/inventory\n
      \ \"\"\"\n  getInventory: JSON @resolve(name: \"default-petstore-8080_gloo-system|Query|getInventory\")\n
      \ \n  \"\"\"\n  Multiple tags can be provided with comma separated strings.
      Use tag1, tag2, tag3 for testing.\n  \n  Equivalent to OpenApiSpec 'OpenAPI
      Petstore' GET /pet/findByTags\n  \"\"\"\n  findPetsByTags(\n    \n    \"\"\"Tags
      to filter by\"\"\"\n    tags: [String]!\n  ): [Pet] @resolve(name: \"default-petstore-8080_gloo-system|Query|findPetsByTags\")\n
      \ \n  \"\"\"\n  Multiple status values can be provided with comma separated
      strings\n  \n  Equivalent to OpenApiSpec 'OpenAPI Petstore' GET /pet/findByStatus\n
      \ \"\"\"\n  findPetsByStatus(\n    \n    \"\"\"Status values that need to be
      considered for filter\"\"\"\n    status: [Status3ListItem]!\n  ): [Pet] @resolve(name:
      \"default-petstore-8080_gloo-system|Query|findPetsByStatus\")\n  \n  \"\"\"\n
      \ For valid response try integer IDs with value <= 5 or > 10. Other values will
      generated exceptions\n  \n  Equivalent to OpenApiSpec 'OpenAPI Petstore' GET
      /store/order/{orderId}\n  \"\"\"\n  getOrderById(\n    \n    \"\"\"ID of pet
      that needs to be fetched\"\"\"\n    orderId: Int!\n  ): Order @resolve(name:
      \"default-petstore-8080_gloo-system|Query|getOrderById\")\n  \n  \"\"\"\n  Logs
      out current logged in user session\n  \n  Equivalent to OpenApiSpec 'OpenAPI
      Petstore' GET /user/logout\n  \"\"\"\n  logoutUser: JSON @resolve(name: \"default-petstore-8080_gloo-system|Query|logoutUser\")\n
      \ \n  \"\"\"\n  Returns a single pet\n  \n  Equivalent to OpenApiSpec 'OpenAPI
      Petstore' GET /pet/{petId}\n  \"\"\"\n  getPetById(\n    \n    \"\"\"ID of pet
      to return\"\"\"\n    petId: Int!\n  ): Pet @resolve(name: \"default-petstore-8080_gloo-system|Query|getPetById\")\n
      \ \n  \"\"\"\n  Get user by user name\n  \n  Equivalent to OpenApiSpec 'OpenAPI
      Petstore' GET /user/{username}\n  \"\"\"\n  getUserByName(\n    \n    \"\"\"The
      name that needs to be fetched. Use user1 for testing.\"\"\"\n    username: String!\n
      \ ): User @resolve(name: \"default-petstore-8080_gloo-system|Query|getUserByName\")\n
      \ \n  \"\"\"\n  Logs user into the system\n  \n  Equivalent to OpenApiSpec 'OpenAPI
      Petstore' GET /user/login\n  \"\"\"\n  loginUser(\n    \n    \"\"\"The user
      name for login\"\"\"\n    username: String!\n    \n    \"\"\"The password for
      login in clear text\"\"\"\n    password: String!\n  ): String @resolve(name:
      \"default-petstore-8080_gloo-system|Query|loginUser\")\n}\n\n\"\"\"A tag for
      a pet\"\"\"\ninput TagInput {\n  name: String\n  id: Int\n}\n\ntype Category
      {\n  id: Int\n  name: String\n}\n\ntype Pet {\n  name: String!\n  photoUrls:
      [String]!\n  \n  \"\"\"pet status in the store\"\"\"\n  status: Status\n  tags:
      [Tag]\n  \n  \"\"\"A category for a pet\"\"\"\n  category: Category\n  id: Int\n}\n\nenum
      Status2 {\n  placed\n  approved\n  delivered\n}\n\ntype Tag {\n  id: Int\n  name:
      String\n}\n\n\"\"\"A category for a pet\"\"\"\ninput CategoryInput {\n  id:
      Int\n  name: String\n}\n\n\"\"\"An order for a pets from the pet store\"\"\"\ninput
      OrderInput {\n  complete: Boolean\n  id: Int\n  petId: Int\n  quantity: Int\n
      \ shipDate: String\n  \n  \"\"\"Order Status\"\"\"\n  status: Status2\n}\n\n\"\"\"A
      User who is purchasing from the pet store\"\"\"\ninput UserInput {\n  id: Int\n
      \ lastName: String\n  password: String\n  phone: String\n  \n  \"\"\"User Status\"\"\"\n
      \ userStatus: Int\n  username: String\n  email: String\n  firstName: String\n}\n"
`;
