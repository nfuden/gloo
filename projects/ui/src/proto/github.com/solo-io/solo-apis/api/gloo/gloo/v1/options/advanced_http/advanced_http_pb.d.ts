/* eslint-disable */
// package: advancedhttp.options.gloo.solo.io
// file: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/advanced_http/advanced_http.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as validate_validate_pb from "../../../../../../../../../validate/validate_pb";

export class ResponseAssertions extends jspb.Message {
  clearResponseMatchersList(): void;
  getResponseMatchersList(): Array<ResponseMatcher>;
  setResponseMatchersList(value: Array<ResponseMatcher>): void;
  addResponseMatchers(value?: ResponseMatcher, index?: number): ResponseMatcher;

  getNoMatchHealth(): HealthCheckResultMap[keyof HealthCheckResultMap];
  setNoMatchHealth(value: HealthCheckResultMap[keyof HealthCheckResultMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResponseAssertions.AsObject;
  static toObject(includeInstance: boolean, msg: ResponseAssertions): ResponseAssertions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResponseAssertions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResponseAssertions;
  static deserializeBinaryFromReader(message: ResponseAssertions, reader: jspb.BinaryReader): ResponseAssertions;
}

export namespace ResponseAssertions {
  export type AsObject = {
    responseMatchersList: Array<ResponseMatcher.AsObject>,
    noMatchHealth: HealthCheckResultMap[keyof HealthCheckResultMap],
  }
}

export class ResponseMatcher extends jspb.Message {
  hasResponseMatch(): boolean;
  clearResponseMatch(): void;
  getResponseMatch(): ResponseMatch | undefined;
  setResponseMatch(value?: ResponseMatch): void;

  getMatchHealth(): HealthCheckResultMap[keyof HealthCheckResultMap];
  setMatchHealth(value: HealthCheckResultMap[keyof HealthCheckResultMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResponseMatcher.AsObject;
  static toObject(includeInstance: boolean, msg: ResponseMatcher): ResponseMatcher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResponseMatcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResponseMatcher;
  static deserializeBinaryFromReader(message: ResponseMatcher, reader: jspb.BinaryReader): ResponseMatcher;
}

export namespace ResponseMatcher {
  export type AsObject = {
    responseMatch?: ResponseMatch.AsObject,
    matchHealth: HealthCheckResultMap[keyof HealthCheckResultMap],
  }
}

export class ResponseMatch extends jspb.Message {
  hasJsonKey(): boolean;
  clearJsonKey(): void;
  getJsonKey(): JsonKey | undefined;
  setJsonKey(value?: JsonKey): void;

  getIgnoreErrorOnParse(): boolean;
  setIgnoreErrorOnParse(value: boolean): void;

  hasHeader(): boolean;
  clearHeader(): void;
  getHeader(): string;
  setHeader(value: string): void;

  hasBody(): boolean;
  clearBody(): void;
  getBody(): google_protobuf_empty_pb.Empty | undefined;
  setBody(value?: google_protobuf_empty_pb.Empty): void;

  getRegex(): string;
  setRegex(value: string): void;

  getSourceCase(): ResponseMatch.SourceCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ResponseMatch.AsObject;
  static toObject(includeInstance: boolean, msg: ResponseMatch): ResponseMatch.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ResponseMatch, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ResponseMatch;
  static deserializeBinaryFromReader(message: ResponseMatch, reader: jspb.BinaryReader): ResponseMatch;
}

export namespace ResponseMatch {
  export type AsObject = {
    jsonKey?: JsonKey.AsObject,
    ignoreErrorOnParse: boolean,
    header: string,
    body?: google_protobuf_empty_pb.Empty.AsObject,
    regex: string,
  }

  export enum SourceCase {
    SOURCE_NOT_SET = 0,
    HEADER = 3,
    BODY = 4,
  }
}

export class JsonKey extends jspb.Message {
  clearPathList(): void;
  getPathList(): Array<JsonKey.PathSegment>;
  setPathList(value: Array<JsonKey.PathSegment>): void;
  addPath(value?: JsonKey.PathSegment, index?: number): JsonKey.PathSegment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JsonKey.AsObject;
  static toObject(includeInstance: boolean, msg: JsonKey): JsonKey.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JsonKey, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JsonKey;
  static deserializeBinaryFromReader(message: JsonKey, reader: jspb.BinaryReader): JsonKey;
}

export namespace JsonKey {
  export type AsObject = {
    pathList: Array<JsonKey.PathSegment.AsObject>,
  }

  export class PathSegment extends jspb.Message {
    hasKey(): boolean;
    clearKey(): void;
    getKey(): string;
    setKey(value: string): void;

    getSegmentCase(): PathSegment.SegmentCase;
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PathSegment.AsObject;
    static toObject(includeInstance: boolean, msg: PathSegment): PathSegment.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PathSegment, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PathSegment;
    static deserializeBinaryFromReader(message: PathSegment, reader: jspb.BinaryReader): PathSegment;
  }

  export namespace PathSegment {
    export type AsObject = {
      key: string,
    }

    export enum SegmentCase {
      SEGMENT_NOT_SET = 0,
      KEY = 1,
    }
  }
}

export interface HealthCheckResultMap {
  DEFAULT: 0;
  HEALTHY: 1;
  DEGRADED: 2;
  UNHEALTHY: 3;
}

export const HealthCheckResult: HealthCheckResultMap;