/* eslint-disable */
// package: matchers.core.gloo.solo.io
// file: github.com/solo-io/solo-apis/api/gloo/gloo/v1/core/matchers/matchers.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as extproto_ext_pb from "../../../../../../../../../extproto/ext_pb";

export class Matcher extends jspb.Message {
  hasPrefix(): boolean;
  clearPrefix(): void;
  getPrefix(): string;
  setPrefix(value: string): void;

  hasExact(): boolean;
  clearExact(): void;
  getExact(): string;
  setExact(value: string): void;

  hasRegex(): boolean;
  clearRegex(): void;
  getRegex(): string;
  setRegex(value: string): void;

  hasConnectMatcher(): boolean;
  clearConnectMatcher(): void;
  getConnectMatcher(): Matcher.ConnectMatcher | undefined;
  setConnectMatcher(value?: Matcher.ConnectMatcher): void;

  hasCaseSensitive(): boolean;
  clearCaseSensitive(): void;
  getCaseSensitive(): google_protobuf_wrappers_pb.BoolValue | undefined;
  setCaseSensitive(value?: google_protobuf_wrappers_pb.BoolValue): void;

  clearHeadersList(): void;
  getHeadersList(): Array<HeaderMatcher>;
  setHeadersList(value: Array<HeaderMatcher>): void;
  addHeaders(value?: HeaderMatcher, index?: number): HeaderMatcher;

  clearQueryParametersList(): void;
  getQueryParametersList(): Array<QueryParameterMatcher>;
  setQueryParametersList(value: Array<QueryParameterMatcher>): void;
  addQueryParameters(value?: QueryParameterMatcher, index?: number): QueryParameterMatcher;

  clearMethodsList(): void;
  getMethodsList(): Array<string>;
  setMethodsList(value: Array<string>): void;
  addMethods(value: string, index?: number): string;

  getPathSpecifierCase(): Matcher.PathSpecifierCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Matcher.AsObject;
  static toObject(includeInstance: boolean, msg: Matcher): Matcher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Matcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Matcher;
  static deserializeBinaryFromReader(message: Matcher, reader: jspb.BinaryReader): Matcher;
}

export namespace Matcher {
  export type AsObject = {
    prefix: string,
    exact: string,
    regex: string,
    connectMatcher?: Matcher.ConnectMatcher.AsObject,
    caseSensitive?: google_protobuf_wrappers_pb.BoolValue.AsObject,
    headersList: Array<HeaderMatcher.AsObject>,
    queryParametersList: Array<QueryParameterMatcher.AsObject>,
    methodsList: Array<string>,
  }

  export class ConnectMatcher extends jspb.Message {
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ConnectMatcher.AsObject;
    static toObject(includeInstance: boolean, msg: ConnectMatcher): ConnectMatcher.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ConnectMatcher, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ConnectMatcher;
    static deserializeBinaryFromReader(message: ConnectMatcher, reader: jspb.BinaryReader): ConnectMatcher;
  }

  export namespace ConnectMatcher {
    export type AsObject = {
    }
  }

  export enum PathSpecifierCase {
    PATH_SPECIFIER_NOT_SET = 0,
    PREFIX = 1,
    EXACT = 2,
    REGEX = 3,
    CONNECT_MATCHER = 9,
  }
}

export class HeaderMatcher extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  getRegex(): boolean;
  setRegex(value: boolean): void;

  getInvertMatch(): boolean;
  setInvertMatch(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HeaderMatcher.AsObject;
  static toObject(includeInstance: boolean, msg: HeaderMatcher): HeaderMatcher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HeaderMatcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HeaderMatcher;
  static deserializeBinaryFromReader(message: HeaderMatcher, reader: jspb.BinaryReader): HeaderMatcher;
}

export namespace HeaderMatcher {
  export type AsObject = {
    name: string,
    value: string,
    regex: boolean,
    invertMatch: boolean,
  }
}

export class QueryParameterMatcher extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getValue(): string;
  setValue(value: string): void;

  getRegex(): boolean;
  setRegex(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryParameterMatcher.AsObject;
  static toObject(includeInstance: boolean, msg: QueryParameterMatcher): QueryParameterMatcher.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: QueryParameterMatcher, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryParameterMatcher;
  static deserializeBinaryFromReader(message: QueryParameterMatcher, reader: jspb.BinaryReader): QueryParameterMatcher;
}

export namespace QueryParameterMatcher {
  export type AsObject = {
    name: string,
    value: string,
    regex: boolean,
  }
}