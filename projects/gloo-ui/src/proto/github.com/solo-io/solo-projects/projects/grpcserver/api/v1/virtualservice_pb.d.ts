// package: glooeeapi.solo.io
// file: github.com/solo-io/solo-projects/projects/grpcserver/api/v1/virtualservice.proto

import * as jspb from "google-protobuf";
import * as gogoproto_gogo_pb from "../../../../../../../gogoproto/gogo_pb";
import * as github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb from "../../../../../../../github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service_pb";
import * as github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb from "../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/proxy_pb";
import * as github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb from "../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/ssl_pb";
import * as github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb from "../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/extauth/extauth_pb";
import * as github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb from "../../../../../../../github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/ratelimit/ratelimit_pb";
import * as github_com_solo_io_solo_projects_projects_grpcserver_api_v1_types_pb from "../../../../../../../github.com/solo-io/solo-projects/projects/grpcserver/api/v1/types_pb";
import * as github_com_solo_io_solo_kit_api_v1_ref_pb from "../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";

export class VirtualServiceDetails extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasPlugins(): boolean;
  clearPlugins(): void;
  getPlugins(): Plugins | undefined;
  setPlugins(value?: Plugins): void;

  hasRaw(): boolean;
  clearRaw(): void;
  getRaw(): github_com_solo_io_solo_projects_projects_grpcserver_api_v1_types_pb.Raw | undefined;
  setRaw(value?: github_com_solo_io_solo_projects_projects_grpcserver_api_v1_types_pb.Raw): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VirtualServiceDetails.AsObject;
  static toObject(includeInstance: boolean, msg: VirtualServiceDetails): VirtualServiceDetails.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VirtualServiceDetails, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VirtualServiceDetails;
  static deserializeBinaryFromReader(message: VirtualServiceDetails, reader: jspb.BinaryReader): VirtualServiceDetails;
}

export namespace VirtualServiceDetails {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    plugins?: Plugins.AsObject,
    raw?: github_com_solo_io_solo_projects_projects_grpcserver_api_v1_types_pb.Raw.AsObject,
  }
}

export class Plugins extends jspb.Message {
  hasExtAuth(): boolean;
  clearExtAuth(): void;
  getExtAuth(): ExtAuthPlugin | undefined;
  setExtAuth(value?: ExtAuthPlugin): void;

  hasRateLimit(): boolean;
  clearRateLimit(): void;
  getRateLimit(): RateLimitPlugin | undefined;
  setRateLimit(value?: RateLimitPlugin): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Plugins.AsObject;
  static toObject(includeInstance: boolean, msg: Plugins): Plugins.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Plugins, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Plugins;
  static deserializeBinaryFromReader(message: Plugins, reader: jspb.BinaryReader): Plugins;
}

export namespace Plugins {
  export type AsObject = {
    extAuth?: ExtAuthPlugin.AsObject,
    rateLimit?: RateLimitPlugin.AsObject,
  }
}

export class ExtAuthPlugin extends jspb.Message {
  hasValue(): boolean;
  clearValue(): void;
  getValue(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.ExtAuthConfig | undefined;
  setValue(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.ExtAuthConfig): void;

  getError(): string;
  setError(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtAuthPlugin.AsObject;
  static toObject(includeInstance: boolean, msg: ExtAuthPlugin): ExtAuthPlugin.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExtAuthPlugin, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtAuthPlugin;
  static deserializeBinaryFromReader(message: ExtAuthPlugin, reader: jspb.BinaryReader): ExtAuthPlugin;
}

export namespace ExtAuthPlugin {
  export type AsObject = {
    value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.ExtAuthConfig.AsObject,
    error: string,
  }
}

export class RateLimitPlugin extends jspb.Message {
  hasValue(): boolean;
  clearValue(): void;
  getValue(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit | undefined;
  setValue(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit): void;

  getError(): string;
  setError(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RateLimitPlugin.AsObject;
  static toObject(includeInstance: boolean, msg: RateLimitPlugin): RateLimitPlugin.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RateLimitPlugin, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RateLimitPlugin;
  static deserializeBinaryFromReader(message: RateLimitPlugin, reader: jspb.BinaryReader): RateLimitPlugin;
}

export namespace RateLimitPlugin {
  export type AsObject = {
    value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit.AsObject,
    error: string,
  }
}

export class RepeatedStrings extends jspb.Message {
  clearValuesList(): void;
  getValuesList(): Array<string>;
  setValuesList(value: Array<string>): void;
  addValues(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RepeatedStrings.AsObject;
  static toObject(includeInstance: boolean, msg: RepeatedStrings): RepeatedStrings.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RepeatedStrings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RepeatedStrings;
  static deserializeBinaryFromReader(message: RepeatedStrings, reader: jspb.BinaryReader): RepeatedStrings;
}

export namespace RepeatedStrings {
  export type AsObject = {
    valuesList: Array<string>,
  }
}

export class RepeatedRoutes extends jspb.Message {
  clearValuesList(): void;
  getValuesList(): Array<github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route>;
  setValuesList(value: Array<github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route>): void;
  addValues(value?: github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route, index?: number): github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RepeatedRoutes.AsObject;
  static toObject(includeInstance: boolean, msg: RepeatedRoutes): RepeatedRoutes.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RepeatedRoutes, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RepeatedRoutes;
  static deserializeBinaryFromReader(message: RepeatedRoutes, reader: jspb.BinaryReader): RepeatedRoutes;
}

export namespace RepeatedRoutes {
  export type AsObject = {
    valuesList: Array<github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route.AsObject>,
  }
}

export class SslConfigValue extends jspb.Message {
  hasValue(): boolean;
  clearValue(): void;
  getValue(): github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig | undefined;
  setValue(value?: github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SslConfigValue.AsObject;
  static toObject(includeInstance: boolean, msg: SslConfigValue): SslConfigValue.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SslConfigValue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SslConfigValue;
  static deserializeBinaryFromReader(message: SslConfigValue, reader: jspb.BinaryReader): SslConfigValue;
}

export namespace SslConfigValue {
  export type AsObject = {
    value?: github_com_solo_io_gloo_projects_gloo_api_v1_ssl_pb.SslConfig.AsObject,
  }
}

export class IngressRateLimitValue extends jspb.Message {
  hasValue(): boolean;
  clearValue(): void;
  getValue(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit | undefined;
  setValue(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IngressRateLimitValue.AsObject;
  static toObject(includeInstance: boolean, msg: IngressRateLimitValue): IngressRateLimitValue.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IngressRateLimitValue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IngressRateLimitValue;
  static deserializeBinaryFromReader(message: IngressRateLimitValue, reader: jspb.BinaryReader): IngressRateLimitValue;
}

export namespace IngressRateLimitValue {
  export type AsObject = {
    value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit.AsObject,
  }
}

export class ExtAuthInput extends jspb.Message {
  hasConfig(): boolean;
  clearConfig(): void;
  getConfig(): ExtAuthInput.Config | undefined;
  setConfig(value?: ExtAuthInput.Config): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtAuthInput.AsObject;
  static toObject(includeInstance: boolean, msg: ExtAuthInput): ExtAuthInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExtAuthInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtAuthInput;
  static deserializeBinaryFromReader(message: ExtAuthInput, reader: jspb.BinaryReader): ExtAuthInput;
}

export namespace ExtAuthInput {
  export type AsObject = {
    config?: ExtAuthInput.Config.AsObject,
  }

  export class Config extends jspb.Message {
    hasOauth(): boolean;
    clearOauth(): void;
    getOauth(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.OAuth | undefined;
    setOauth(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.OAuth): void;

    hasCustomAuth(): boolean;
    clearCustomAuth(): void;
    getCustomAuth(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.CustomAuth | undefined;
    setCustomAuth(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.CustomAuth): void;

    getValueCase(): Config.ValueCase;
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Config.AsObject;
    static toObject(includeInstance: boolean, msg: Config): Config.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Config, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Config;
    static deserializeBinaryFromReader(message: Config, reader: jspb.BinaryReader): Config;
  }

  export namespace Config {
    export type AsObject = {
      oauth?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.OAuth.AsObject,
      customAuth?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.CustomAuth.AsObject,
    }

    export enum ValueCase {
      VALUE_NOT_SET = 0,
      OAUTH = 1,
      CUSTOM_AUTH = 2,
    }
  }
}

export class GetVirtualServiceRequest extends jspb.Message {
  hasRef(): boolean;
  clearRef(): void;
  getRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVirtualServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetVirtualServiceRequest): GetVirtualServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetVirtualServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVirtualServiceRequest;
  static deserializeBinaryFromReader(message: GetVirtualServiceRequest, reader: jspb.BinaryReader): GetVirtualServiceRequest;
}

export namespace GetVirtualServiceRequest {
  export type AsObject = {
    ref?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
  }
}

export class GetVirtualServiceResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVirtualServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetVirtualServiceResponse): GetVirtualServiceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetVirtualServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVirtualServiceResponse;
  static deserializeBinaryFromReader(message: GetVirtualServiceResponse, reader: jspb.BinaryReader): GetVirtualServiceResponse;
}

export namespace GetVirtualServiceResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

export class ListVirtualServicesRequest extends jspb.Message {
  clearNamespacesList(): void;
  getNamespacesList(): Array<string>;
  setNamespacesList(value: Array<string>): void;
  addNamespaces(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVirtualServicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListVirtualServicesRequest): ListVirtualServicesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListVirtualServicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVirtualServicesRequest;
  static deserializeBinaryFromReader(message: ListVirtualServicesRequest, reader: jspb.BinaryReader): ListVirtualServicesRequest;
}

export namespace ListVirtualServicesRequest {
  export type AsObject = {
    namespacesList: Array<string>,
  }
}

export class ListVirtualServicesResponse extends jspb.Message {
  clearVirtualServicesList(): void;
  getVirtualServicesList(): Array<github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService>;
  setVirtualServicesList(value: Array<github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService>): void;
  addVirtualServices(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService, index?: number): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService;

  clearVirtualServiceDetailsList(): void;
  getVirtualServiceDetailsList(): Array<VirtualServiceDetails>;
  setVirtualServiceDetailsList(value: Array<VirtualServiceDetails>): void;
  addVirtualServiceDetails(value?: VirtualServiceDetails, index?: number): VirtualServiceDetails;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVirtualServicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListVirtualServicesResponse): ListVirtualServicesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListVirtualServicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVirtualServicesResponse;
  static deserializeBinaryFromReader(message: ListVirtualServicesResponse, reader: jspb.BinaryReader): ListVirtualServicesResponse;
}

export namespace ListVirtualServicesResponse {
  export type AsObject = {
    virtualServicesList: Array<github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject>,
    virtualServiceDetailsList: Array<VirtualServiceDetails.AsObject>,
  }
}

export class VirtualServiceInput extends jspb.Message {
  hasRef(): boolean;
  clearRef(): void;
  getRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  getDisplayName(): string;
  setDisplayName(value: string): void;

  clearDomainsList(): void;
  getDomainsList(): Array<string>;
  setDomainsList(value: Array<string>): void;
  addDomains(value: string, index?: number): string;

  clearRoutesList(): void;
  getRoutesList(): Array<github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route>;
  setRoutesList(value: Array<github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route>): void;
  addRoutes(value?: github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route, index?: number): github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route;

  hasSecretRef(): boolean;
  clearSecretRef(): void;
  getSecretRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setSecretRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  hasRateLimitConfig(): boolean;
  clearRateLimitConfig(): void;
  getRateLimitConfig(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit | undefined;
  setRateLimitConfig(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit): void;

  hasBasicAuth(): boolean;
  clearBasicAuth(): void;
  getBasicAuth(): VirtualServiceInput.BasicAuthInput | undefined;
  setBasicAuth(value?: VirtualServiceInput.BasicAuthInput): void;

  hasOauth(): boolean;
  clearOauth(): void;
  getOauth(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.OAuth | undefined;
  setOauth(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.OAuth): void;

  hasCustomAuth(): boolean;
  clearCustomAuth(): void;
  getCustomAuth(): github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.CustomAuth | undefined;
  setCustomAuth(value?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.CustomAuth): void;

  getExtAuthConfigCase(): VirtualServiceInput.ExtAuthConfigCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VirtualServiceInput.AsObject;
  static toObject(includeInstance: boolean, msg: VirtualServiceInput): VirtualServiceInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VirtualServiceInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VirtualServiceInput;
  static deserializeBinaryFromReader(message: VirtualServiceInput, reader: jspb.BinaryReader): VirtualServiceInput;
}

export namespace VirtualServiceInput {
  export type AsObject = {
    ref?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    displayName: string,
    domainsList: Array<string>,
    routesList: Array<github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route.AsObject>,
    secretRef?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    rateLimitConfig?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_ratelimit_ratelimit_pb.IngressRateLimit.AsObject,
    basicAuth?: VirtualServiceInput.BasicAuthInput.AsObject,
    oauth?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.OAuth.AsObject,
    customAuth?: github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_plugins_extauth_extauth_pb.CustomAuth.AsObject,
  }

  export class BasicAuthInput extends jspb.Message {
    getRealm(): string;
    setRealm(value: string): void;

    getSpecCsv(): string;
    setSpecCsv(value: string): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BasicAuthInput.AsObject;
    static toObject(includeInstance: boolean, msg: BasicAuthInput): BasicAuthInput.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BasicAuthInput, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BasicAuthInput;
    static deserializeBinaryFromReader(message: BasicAuthInput, reader: jspb.BinaryReader): BasicAuthInput;
  }

  export namespace BasicAuthInput {
    export type AsObject = {
      realm: string,
      specCsv: string,
    }
  }

  export enum ExtAuthConfigCase {
    EXT_AUTH_CONFIG_NOT_SET = 0,
    BASIC_AUTH = 7,
    OAUTH = 8,
    CUSTOM_AUTH = 9,
  }
}

export class VirtualServiceInputV2 extends jspb.Message {
  hasRef(): boolean;
  clearRef(): void;
  getRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  hasDisplayName(): boolean;
  clearDisplayName(): void;
  getDisplayName(): google_protobuf_wrappers_pb.StringValue | undefined;
  setDisplayName(value?: google_protobuf_wrappers_pb.StringValue): void;

  hasDomains(): boolean;
  clearDomains(): void;
  getDomains(): RepeatedStrings | undefined;
  setDomains(value?: RepeatedStrings): void;

  hasRoutes(): boolean;
  clearRoutes(): void;
  getRoutes(): RepeatedRoutes | undefined;
  setRoutes(value?: RepeatedRoutes): void;

  hasSslConfig(): boolean;
  clearSslConfig(): void;
  getSslConfig(): SslConfigValue | undefined;
  setSslConfig(value?: SslConfigValue): void;

  hasRateLimitConfig(): boolean;
  clearRateLimitConfig(): void;
  getRateLimitConfig(): IngressRateLimitValue | undefined;
  setRateLimitConfig(value?: IngressRateLimitValue): void;

  hasExtAuthConfig(): boolean;
  clearExtAuthConfig(): void;
  getExtAuthConfig(): ExtAuthInput | undefined;
  setExtAuthConfig(value?: ExtAuthInput): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VirtualServiceInputV2.AsObject;
  static toObject(includeInstance: boolean, msg: VirtualServiceInputV2): VirtualServiceInputV2.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: VirtualServiceInputV2, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VirtualServiceInputV2;
  static deserializeBinaryFromReader(message: VirtualServiceInputV2, reader: jspb.BinaryReader): VirtualServiceInputV2;
}

export namespace VirtualServiceInputV2 {
  export type AsObject = {
    ref?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    displayName?: google_protobuf_wrappers_pb.StringValue.AsObject,
    domains?: RepeatedStrings.AsObject,
    routes?: RepeatedRoutes.AsObject,
    sslConfig?: SslConfigValue.AsObject,
    rateLimitConfig?: IngressRateLimitValue.AsObject,
    extAuthConfig?: ExtAuthInput.AsObject,
  }
}

export class CreateVirtualServiceRequest extends jspb.Message {
  hasInput(): boolean;
  clearInput(): void;
  getInput(): VirtualServiceInput | undefined;
  setInput(value?: VirtualServiceInput): void;

  hasInputV2(): boolean;
  clearInputV2(): void;
  getInputV2(): VirtualServiceInputV2 | undefined;
  setInputV2(value?: VirtualServiceInputV2): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateVirtualServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateVirtualServiceRequest): CreateVirtualServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateVirtualServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateVirtualServiceRequest;
  static deserializeBinaryFromReader(message: CreateVirtualServiceRequest, reader: jspb.BinaryReader): CreateVirtualServiceRequest;
}

export namespace CreateVirtualServiceRequest {
  export type AsObject = {
    input?: VirtualServiceInput.AsObject,
    inputV2?: VirtualServiceInputV2.AsObject,
  }
}

export class CreateVirtualServiceResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateVirtualServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateVirtualServiceResponse): CreateVirtualServiceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateVirtualServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateVirtualServiceResponse;
  static deserializeBinaryFromReader(message: CreateVirtualServiceResponse, reader: jspb.BinaryReader): CreateVirtualServiceResponse;
}

export namespace CreateVirtualServiceResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

export class UpdateVirtualServiceRequest extends jspb.Message {
  hasInput(): boolean;
  clearInput(): void;
  getInput(): VirtualServiceInput | undefined;
  setInput(value?: VirtualServiceInput): void;

  hasInputV2(): boolean;
  clearInputV2(): void;
  getInputV2(): VirtualServiceInputV2 | undefined;
  setInputV2(value?: VirtualServiceInputV2): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateVirtualServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateVirtualServiceRequest): UpdateVirtualServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateVirtualServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateVirtualServiceRequest;
  static deserializeBinaryFromReader(message: UpdateVirtualServiceRequest, reader: jspb.BinaryReader): UpdateVirtualServiceRequest;
}

export namespace UpdateVirtualServiceRequest {
  export type AsObject = {
    input?: VirtualServiceInput.AsObject,
    inputV2?: VirtualServiceInputV2.AsObject,
  }
}

export class UpdateVirtualServiceYamlRequest extends jspb.Message {
  hasEditedYamlData(): boolean;
  clearEditedYamlData(): void;
  getEditedYamlData(): github_com_solo_io_solo_projects_projects_grpcserver_api_v1_types_pb.EditedResourceYaml | undefined;
  setEditedYamlData(value?: github_com_solo_io_solo_projects_projects_grpcserver_api_v1_types_pb.EditedResourceYaml): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateVirtualServiceYamlRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateVirtualServiceYamlRequest): UpdateVirtualServiceYamlRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateVirtualServiceYamlRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateVirtualServiceYamlRequest;
  static deserializeBinaryFromReader(message: UpdateVirtualServiceYamlRequest, reader: jspb.BinaryReader): UpdateVirtualServiceYamlRequest;
}

export namespace UpdateVirtualServiceYamlRequest {
  export type AsObject = {
    editedYamlData?: github_com_solo_io_solo_projects_projects_grpcserver_api_v1_types_pb.EditedResourceYaml.AsObject,
  }
}

export class UpdateVirtualServiceResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateVirtualServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateVirtualServiceResponse): UpdateVirtualServiceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateVirtualServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateVirtualServiceResponse;
  static deserializeBinaryFromReader(message: UpdateVirtualServiceResponse, reader: jspb.BinaryReader): UpdateVirtualServiceResponse;
}

export namespace UpdateVirtualServiceResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

export class DeleteVirtualServiceRequest extends jspb.Message {
  hasRef(): boolean;
  clearRef(): void;
  getRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteVirtualServiceRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteVirtualServiceRequest): DeleteVirtualServiceRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteVirtualServiceRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteVirtualServiceRequest;
  static deserializeBinaryFromReader(message: DeleteVirtualServiceRequest, reader: jspb.BinaryReader): DeleteVirtualServiceRequest;
}

export namespace DeleteVirtualServiceRequest {
  export type AsObject = {
    ref?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
  }
}

export class DeleteVirtualServiceResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteVirtualServiceResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteVirtualServiceResponse): DeleteVirtualServiceResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteVirtualServiceResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteVirtualServiceResponse;
  static deserializeBinaryFromReader(message: DeleteVirtualServiceResponse, reader: jspb.BinaryReader): DeleteVirtualServiceResponse;
}

export namespace DeleteVirtualServiceResponse {
  export type AsObject = {
  }
}

export class RouteInput extends jspb.Message {
  hasVirtualServiceRef(): boolean;
  clearVirtualServiceRef(): void;
  getVirtualServiceRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setVirtualServiceRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  getIndex(): number;
  setIndex(value: number): void;

  hasRoute(): boolean;
  clearRoute(): void;
  getRoute(): github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route | undefined;
  setRoute(value?: github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RouteInput.AsObject;
  static toObject(includeInstance: boolean, msg: RouteInput): RouteInput.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RouteInput, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RouteInput;
  static deserializeBinaryFromReader(message: RouteInput, reader: jspb.BinaryReader): RouteInput;
}

export namespace RouteInput {
  export type AsObject = {
    virtualServiceRef?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    index: number,
    route?: github_com_solo_io_gloo_projects_gloo_api_v1_proxy_pb.Route.AsObject,
  }
}

export class CreateRouteRequest extends jspb.Message {
  hasInput(): boolean;
  clearInput(): void;
  getInput(): RouteInput | undefined;
  setInput(value?: RouteInput): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRouteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRouteRequest): CreateRouteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRouteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRouteRequest;
  static deserializeBinaryFromReader(message: CreateRouteRequest, reader: jspb.BinaryReader): CreateRouteRequest;
}

export namespace CreateRouteRequest {
  export type AsObject = {
    input?: RouteInput.AsObject,
  }
}

export class CreateRouteResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRouteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRouteResponse): CreateRouteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRouteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRouteResponse;
  static deserializeBinaryFromReader(message: CreateRouteResponse, reader: jspb.BinaryReader): CreateRouteResponse;
}

export namespace CreateRouteResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

export class UpdateRouteRequest extends jspb.Message {
  hasInput(): boolean;
  clearInput(): void;
  getInput(): RouteInput | undefined;
  setInput(value?: RouteInput): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRouteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRouteRequest): UpdateRouteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRouteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRouteRequest;
  static deserializeBinaryFromReader(message: UpdateRouteRequest, reader: jspb.BinaryReader): UpdateRouteRequest;
}

export namespace UpdateRouteRequest {
  export type AsObject = {
    input?: RouteInput.AsObject,
  }
}

export class UpdateRouteResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRouteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRouteResponse): UpdateRouteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateRouteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRouteResponse;
  static deserializeBinaryFromReader(message: UpdateRouteResponse, reader: jspb.BinaryReader): UpdateRouteResponse;
}

export namespace UpdateRouteResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

export class DeleteRouteRequest extends jspb.Message {
  hasVirtualServiceRef(): boolean;
  clearVirtualServiceRef(): void;
  getVirtualServiceRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setVirtualServiceRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  getIndex(): number;
  setIndex(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRouteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRouteRequest): DeleteRouteRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteRouteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRouteRequest;
  static deserializeBinaryFromReader(message: DeleteRouteRequest, reader: jspb.BinaryReader): DeleteRouteRequest;
}

export namespace DeleteRouteRequest {
  export type AsObject = {
    virtualServiceRef?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    index: number,
  }
}

export class DeleteRouteResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRouteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRouteResponse): DeleteRouteResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DeleteRouteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRouteResponse;
  static deserializeBinaryFromReader(message: DeleteRouteResponse, reader: jspb.BinaryReader): DeleteRouteResponse;
}

export namespace DeleteRouteResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

export class SwapRoutesRequest extends jspb.Message {
  hasVirtualServiceRef(): boolean;
  clearVirtualServiceRef(): void;
  getVirtualServiceRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setVirtualServiceRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  getIndex1(): number;
  setIndex1(value: number): void;

  getIndex2(): number;
  setIndex2(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SwapRoutesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SwapRoutesRequest): SwapRoutesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SwapRoutesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SwapRoutesRequest;
  static deserializeBinaryFromReader(message: SwapRoutesRequest, reader: jspb.BinaryReader): SwapRoutesRequest;
}

export namespace SwapRoutesRequest {
  export type AsObject = {
    virtualServiceRef?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    index1: number,
    index2: number,
  }
}

export class SwapRoutesResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SwapRoutesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SwapRoutesResponse): SwapRoutesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SwapRoutesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SwapRoutesResponse;
  static deserializeBinaryFromReader(message: SwapRoutesResponse, reader: jspb.BinaryReader): SwapRoutesResponse;
}

export namespace SwapRoutesResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

export class ShiftRoutesRequest extends jspb.Message {
  hasVirtualServiceRef(): boolean;
  clearVirtualServiceRef(): void;
  getVirtualServiceRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setVirtualServiceRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  getFromIndex(): number;
  setFromIndex(value: number): void;

  getToIndex(): number;
  setToIndex(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ShiftRoutesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ShiftRoutesRequest): ShiftRoutesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ShiftRoutesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ShiftRoutesRequest;
  static deserializeBinaryFromReader(message: ShiftRoutesRequest, reader: jspb.BinaryReader): ShiftRoutesRequest;
}

export namespace ShiftRoutesRequest {
  export type AsObject = {
    virtualServiceRef?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    fromIndex: number,
    toIndex: number,
  }
}

export class ShiftRoutesResponse extends jspb.Message {
  hasVirtualService(): boolean;
  clearVirtualService(): void;
  getVirtualService(): github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService | undefined;
  setVirtualService(value?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService): void;

  hasVirtualServiceDetails(): boolean;
  clearVirtualServiceDetails(): void;
  getVirtualServiceDetails(): VirtualServiceDetails | undefined;
  setVirtualServiceDetails(value?: VirtualServiceDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ShiftRoutesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ShiftRoutesResponse): ShiftRoutesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ShiftRoutesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ShiftRoutesResponse;
  static deserializeBinaryFromReader(message: ShiftRoutesResponse, reader: jspb.BinaryReader): ShiftRoutesResponse;
}

export namespace ShiftRoutesResponse {
  export type AsObject = {
    virtualService?: github_com_solo_io_gloo_projects_gateway_api_v1_virtual_service_pb.VirtualService.AsObject,
    virtualServiceDetails?: VirtualServiceDetails.AsObject,
  }
}

