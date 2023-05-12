/* eslint-disable */
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_rest_rest_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/rest/rest_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_grpc_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/grpc/grpc_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_graphql_graphql_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/graphql/graphql_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_json_grpc_json_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/grpc_json/grpc_json_pb.js');
var extproto_ext_pb = require('../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.options.gloo.solo.io.ServiceSpec', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.options.gloo.solo.io.ServiceSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.options.gloo.solo.io.ServiceSpec.oneofGroups_);
};
goog.inherits(proto.options.gloo.solo.io.ServiceSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.options.gloo.solo.io.ServiceSpec.displayName = 'proto.options.gloo.solo.io.ServiceSpec';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.options.gloo.solo.io.ServiceSpec.oneofGroups_ = [[1,2,3,4]];

/**
 * @enum {number}
 */
proto.options.gloo.solo.io.ServiceSpec.PluginTypeCase = {
  PLUGIN_TYPE_NOT_SET: 0,
  REST: 1,
  GRPC: 2,
  GRPC_JSON_TRANSCODER: 3,
  GRAPHQL: 4
};

/**
 * @return {proto.options.gloo.solo.io.ServiceSpec.PluginTypeCase}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.getPluginTypeCase = function() {
  return /** @type {proto.options.gloo.solo.io.ServiceSpec.PluginTypeCase} */(jspb.Message.computeOneofCase(this, proto.options.gloo.solo.io.ServiceSpec.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.options.gloo.solo.io.ServiceSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.options.gloo.solo.io.ServiceSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.options.gloo.solo.io.ServiceSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    rest: (f = msg.getRest()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_rest_rest_pb.ServiceSpec.toObject(includeInstance, f),
    grpc: (f = msg.getGrpc()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_grpc_pb.ServiceSpec.toObject(includeInstance, f),
    grpcJsonTranscoder: (f = msg.getGrpcJsonTranscoder()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_json_grpc_json_pb.GrpcJsonTranscoder.toObject(includeInstance, f),
    graphql: (f = msg.getGraphql()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_graphql_graphql_pb.ServiceSpec.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.options.gloo.solo.io.ServiceSpec}
 */
proto.options.gloo.solo.io.ServiceSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.options.gloo.solo.io.ServiceSpec;
  return proto.options.gloo.solo.io.ServiceSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.options.gloo.solo.io.ServiceSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.options.gloo.solo.io.ServiceSpec}
 */
proto.options.gloo.solo.io.ServiceSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_rest_rest_pb.ServiceSpec;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_rest_rest_pb.ServiceSpec.deserializeBinaryFromReader);
      msg.setRest(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_grpc_pb.ServiceSpec;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_grpc_pb.ServiceSpec.deserializeBinaryFromReader);
      msg.setGrpc(value);
      break;
    case 3:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_json_grpc_json_pb.GrpcJsonTranscoder;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_json_grpc_json_pb.GrpcJsonTranscoder.deserializeBinaryFromReader);
      msg.setGrpcJsonTranscoder(value);
      break;
    case 4:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_graphql_graphql_pb.ServiceSpec;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_graphql_graphql_pb.ServiceSpec.deserializeBinaryFromReader);
      msg.setGraphql(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.options.gloo.solo.io.ServiceSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.options.gloo.solo.io.ServiceSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.options.gloo.solo.io.ServiceSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRest();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_rest_rest_pb.ServiceSpec.serializeBinaryToWriter
    );
  }
  f = message.getGrpc();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_grpc_pb.ServiceSpec.serializeBinaryToWriter
    );
  }
  f = message.getGrpcJsonTranscoder();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_json_grpc_json_pb.GrpcJsonTranscoder.serializeBinaryToWriter
    );
  }
  f = message.getGraphql();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_graphql_graphql_pb.ServiceSpec.serializeBinaryToWriter
    );
  }
};


/**
 * optional rest.options.gloo.solo.io.ServiceSpec rest = 1;
 * @return {?proto.rest.options.gloo.solo.io.ServiceSpec}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.getRest = function() {
  return /** @type{?proto.rest.options.gloo.solo.io.ServiceSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_rest_rest_pb.ServiceSpec, 1));
};


/** @param {?proto.rest.options.gloo.solo.io.ServiceSpec|undefined} value */
proto.options.gloo.solo.io.ServiceSpec.prototype.setRest = function(value) {
  jspb.Message.setOneofWrapperField(this, 1, proto.options.gloo.solo.io.ServiceSpec.oneofGroups_[0], value);
};


proto.options.gloo.solo.io.ServiceSpec.prototype.clearRest = function() {
  this.setRest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.hasRest = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional grpc.options.gloo.solo.io.ServiceSpec grpc = 2;
 * @return {?proto.grpc.options.gloo.solo.io.ServiceSpec}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.getGrpc = function() {
  return /** @type{?proto.grpc.options.gloo.solo.io.ServiceSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_grpc_pb.ServiceSpec, 2));
};


/** @param {?proto.grpc.options.gloo.solo.io.ServiceSpec|undefined} value */
proto.options.gloo.solo.io.ServiceSpec.prototype.setGrpc = function(value) {
  jspb.Message.setOneofWrapperField(this, 2, proto.options.gloo.solo.io.ServiceSpec.oneofGroups_[0], value);
};


proto.options.gloo.solo.io.ServiceSpec.prototype.clearGrpc = function() {
  this.setGrpc(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.hasGrpc = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional grpc_json.options.gloo.solo.io.GrpcJsonTranscoder grpc_json_transcoder = 3;
 * @return {?proto.grpc_json.options.gloo.solo.io.GrpcJsonTranscoder}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.getGrpcJsonTranscoder = function() {
  return /** @type{?proto.grpc_json.options.gloo.solo.io.GrpcJsonTranscoder} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_grpc_json_grpc_json_pb.GrpcJsonTranscoder, 3));
};


/** @param {?proto.grpc_json.options.gloo.solo.io.GrpcJsonTranscoder|undefined} value */
proto.options.gloo.solo.io.ServiceSpec.prototype.setGrpcJsonTranscoder = function(value) {
  jspb.Message.setOneofWrapperField(this, 3, proto.options.gloo.solo.io.ServiceSpec.oneofGroups_[0], value);
};


proto.options.gloo.solo.io.ServiceSpec.prototype.clearGrpcJsonTranscoder = function() {
  this.setGrpcJsonTranscoder(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.hasGrpcJsonTranscoder = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional graphql.options.gloo.solo.io.ServiceSpec graphql = 4;
 * @return {?proto.graphql.options.gloo.solo.io.ServiceSpec}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.getGraphql = function() {
  return /** @type{?proto.graphql.options.gloo.solo.io.ServiceSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_options_graphql_graphql_pb.ServiceSpec, 4));
};


/** @param {?proto.graphql.options.gloo.solo.io.ServiceSpec|undefined} value */
proto.options.gloo.solo.io.ServiceSpec.prototype.setGraphql = function(value) {
  jspb.Message.setOneofWrapperField(this, 4, proto.options.gloo.solo.io.ServiceSpec.oneofGroups_[0], value);
};


proto.options.gloo.solo.io.ServiceSpec.prototype.clearGraphql = function() {
  this.setGraphql(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.options.gloo.solo.io.ServiceSpec.prototype.hasGraphql = function() {
  return jspb.Message.getField(this, 4) != null;
};


goog.object.extend(exports, proto.options.gloo.solo.io);