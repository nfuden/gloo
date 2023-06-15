/* eslint-disable */
// package: gloo.solo.io
// file: github.com/solo-io/solo-apis/api/gloo/gloo/v1/load_balancer.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as extproto_ext_pb from "../../../../../../../extproto/ext_pb";

export class LoadBalancerConfig extends jspb.Message {
  hasHealthyPanicThreshold(): boolean;
  clearHealthyPanicThreshold(): void;
  getHealthyPanicThreshold(): google_protobuf_wrappers_pb.DoubleValue | undefined;
  setHealthyPanicThreshold(value?: google_protobuf_wrappers_pb.DoubleValue): void;

  hasUpdateMergeWindow(): boolean;
  clearUpdateMergeWindow(): void;
  getUpdateMergeWindow(): google_protobuf_duration_pb.Duration | undefined;
  setUpdateMergeWindow(value?: google_protobuf_duration_pb.Duration): void;

  hasRoundRobin(): boolean;
  clearRoundRobin(): void;
  getRoundRobin(): LoadBalancerConfig.RoundRobin | undefined;
  setRoundRobin(value?: LoadBalancerConfig.RoundRobin): void;

  hasLeastRequest(): boolean;
  clearLeastRequest(): void;
  getLeastRequest(): LoadBalancerConfig.LeastRequest | undefined;
  setLeastRequest(value?: LoadBalancerConfig.LeastRequest): void;

  hasRandom(): boolean;
  clearRandom(): void;
  getRandom(): LoadBalancerConfig.Random | undefined;
  setRandom(value?: LoadBalancerConfig.Random): void;

  hasRingHash(): boolean;
  clearRingHash(): void;
  getRingHash(): LoadBalancerConfig.RingHash | undefined;
  setRingHash(value?: LoadBalancerConfig.RingHash): void;

  hasMaglev(): boolean;
  clearMaglev(): void;
  getMaglev(): LoadBalancerConfig.Maglev | undefined;
  setMaglev(value?: LoadBalancerConfig.Maglev): void;

  hasLocalityWeightedLbConfig(): boolean;
  clearLocalityWeightedLbConfig(): void;
  getLocalityWeightedLbConfig(): google_protobuf_empty_pb.Empty | undefined;
  setLocalityWeightedLbConfig(value?: google_protobuf_empty_pb.Empty): void;

  getTypeCase(): LoadBalancerConfig.TypeCase;
  getLocalityConfigCase(): LoadBalancerConfig.LocalityConfigCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoadBalancerConfig.AsObject;
  static toObject(includeInstance: boolean, msg: LoadBalancerConfig): LoadBalancerConfig.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoadBalancerConfig, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoadBalancerConfig;
  static deserializeBinaryFromReader(message: LoadBalancerConfig, reader: jspb.BinaryReader): LoadBalancerConfig;
}

export namespace LoadBalancerConfig {
  export type AsObject = {
    healthyPanicThreshold?: google_protobuf_wrappers_pb.DoubleValue.AsObject,
    updateMergeWindow?: google_protobuf_duration_pb.Duration.AsObject,
    roundRobin?: LoadBalancerConfig.RoundRobin.AsObject,
    leastRequest?: LoadBalancerConfig.LeastRequest.AsObject,
    random?: LoadBalancerConfig.Random.AsObject,
    ringHash?: LoadBalancerConfig.RingHash.AsObject,
    maglev?: LoadBalancerConfig.Maglev.AsObject,
    localityWeightedLbConfig?: google_protobuf_empty_pb.Empty.AsObject,
  }

  export class RoundRobin extends jspb.Message {
    hasSlowStartConfig(): boolean;
    clearSlowStartConfig(): void;
    getSlowStartConfig(): LoadBalancerConfig.SlowStartConfig | undefined;
    setSlowStartConfig(value?: LoadBalancerConfig.SlowStartConfig): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RoundRobin.AsObject;
    static toObject(includeInstance: boolean, msg: RoundRobin): RoundRobin.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RoundRobin, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RoundRobin;
    static deserializeBinaryFromReader(message: RoundRobin, reader: jspb.BinaryReader): RoundRobin;
  }

  export namespace RoundRobin {
    export type AsObject = {
      slowStartConfig?: LoadBalancerConfig.SlowStartConfig.AsObject,
    }
  }

  export class LeastRequest extends jspb.Message {
    getChoiceCount(): number;
    setChoiceCount(value: number): void;

    hasSlowStartConfig(): boolean;
    clearSlowStartConfig(): void;
    getSlowStartConfig(): LoadBalancerConfig.SlowStartConfig | undefined;
    setSlowStartConfig(value?: LoadBalancerConfig.SlowStartConfig): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LeastRequest.AsObject;
    static toObject(includeInstance: boolean, msg: LeastRequest): LeastRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LeastRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LeastRequest;
    static deserializeBinaryFromReader(message: LeastRequest, reader: jspb.BinaryReader): LeastRequest;
  }

  export namespace LeastRequest {
    export type AsObject = {
      choiceCount: number,
      slowStartConfig?: LoadBalancerConfig.SlowStartConfig.AsObject,
    }
  }

  export class Random extends jspb.Message {
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Random.AsObject;
    static toObject(includeInstance: boolean, msg: Random): Random.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Random, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Random;
    static deserializeBinaryFromReader(message: Random, reader: jspb.BinaryReader): Random;
  }

  export namespace Random {
    export type AsObject = {
    }
  }

  export class RingHashConfig extends jspb.Message {
    getMinimumRingSize(): number;
    setMinimumRingSize(value: number): void;

    getMaximumRingSize(): number;
    setMaximumRingSize(value: number): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RingHashConfig.AsObject;
    static toObject(includeInstance: boolean, msg: RingHashConfig): RingHashConfig.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RingHashConfig, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RingHashConfig;
    static deserializeBinaryFromReader(message: RingHashConfig, reader: jspb.BinaryReader): RingHashConfig;
  }

  export namespace RingHashConfig {
    export type AsObject = {
      minimumRingSize: number,
      maximumRingSize: number,
    }
  }

  export class RingHash extends jspb.Message {
    hasRingHashConfig(): boolean;
    clearRingHashConfig(): void;
    getRingHashConfig(): LoadBalancerConfig.RingHashConfig | undefined;
    setRingHashConfig(value?: LoadBalancerConfig.RingHashConfig): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RingHash.AsObject;
    static toObject(includeInstance: boolean, msg: RingHash): RingHash.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RingHash, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RingHash;
    static deserializeBinaryFromReader(message: RingHash, reader: jspb.BinaryReader): RingHash;
  }

  export namespace RingHash {
    export type AsObject = {
      ringHashConfig?: LoadBalancerConfig.RingHashConfig.AsObject,
    }
  }

  export class Maglev extends jspb.Message {
    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Maglev.AsObject;
    static toObject(includeInstance: boolean, msg: Maglev): Maglev.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Maglev, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Maglev;
    static deserializeBinaryFromReader(message: Maglev, reader: jspb.BinaryReader): Maglev;
  }

  export namespace Maglev {
    export type AsObject = {
    }
  }

  export class SlowStartConfig extends jspb.Message {
    hasSlowStartWindow(): boolean;
    clearSlowStartWindow(): void;
    getSlowStartWindow(): google_protobuf_duration_pb.Duration | undefined;
    setSlowStartWindow(value?: google_protobuf_duration_pb.Duration): void;

    hasAggression(): boolean;
    clearAggression(): void;
    getAggression(): google_protobuf_wrappers_pb.DoubleValue | undefined;
    setAggression(value?: google_protobuf_wrappers_pb.DoubleValue): void;

    hasMinWeightPercent(): boolean;
    clearMinWeightPercent(): void;
    getMinWeightPercent(): google_protobuf_wrappers_pb.DoubleValue | undefined;
    setMinWeightPercent(value?: google_protobuf_wrappers_pb.DoubleValue): void;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SlowStartConfig.AsObject;
    static toObject(includeInstance: boolean, msg: SlowStartConfig): SlowStartConfig.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SlowStartConfig, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SlowStartConfig;
    static deserializeBinaryFromReader(message: SlowStartConfig, reader: jspb.BinaryReader): SlowStartConfig;
  }

  export namespace SlowStartConfig {
    export type AsObject = {
      slowStartWindow?: google_protobuf_duration_pb.Duration.AsObject,
      aggression?: google_protobuf_wrappers_pb.DoubleValue.AsObject,
      minWeightPercent?: google_protobuf_wrappers_pb.DoubleValue.AsObject,
    }
  }

  export enum TypeCase {
    TYPE_NOT_SET = 0,
    ROUND_ROBIN = 3,
    LEAST_REQUEST = 4,
    RANDOM = 5,
    RING_HASH = 6,
    MAGLEV = 7,
  }

  export enum LocalityConfigCase {
    LOCALITY_CONFIG_NOT_SET = 0,
    LOCALITY_WEIGHTED_LB_CONFIG = 8,
  }
}