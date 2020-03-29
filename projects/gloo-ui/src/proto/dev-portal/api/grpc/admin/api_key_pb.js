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

var gogoproto_gogo_pb = require('../../../../gogoproto/gogo_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var dev$portal_api_dev$portal_v1_common_pb = require('../../../../dev-portal/api/dev-portal/v1/common_pb.js');
var dev$portal_api_grpc_common_common_pb = require('../../../../dev-portal/api/grpc/common/common_pb.js');
goog.exportSymbol('proto.admin.devportal.solo.io.ApiKey', null, global);
goog.exportSymbol('proto.admin.devportal.solo.io.ApiKeyList', null, global);

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
proto.admin.devportal.solo.io.ApiKey = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.admin.devportal.solo.io.ApiKey, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.admin.devportal.solo.io.ApiKey.displayName = 'proto.admin.devportal.solo.io.ApiKey';
}


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
proto.admin.devportal.solo.io.ApiKey.prototype.toObject = function(opt_includeInstance) {
  return proto.admin.devportal.solo.io.ApiKey.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.admin.devportal.solo.io.ApiKey} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.admin.devportal.solo.io.ApiKey.toObject = function(includeInstance, msg) {
  var f, obj = {
    metadata: (f = msg.getMetadata()) && dev$portal_api_grpc_common_common_pb.ObjectMeta.toObject(includeInstance, f),
    value: jspb.Message.getFieldWithDefault(msg, 2, ""),
    user: (f = msg.getUser()) && dev$portal_api_dev$portal_v1_common_pb.ObjectRef.toObject(includeInstance, f),
    keyScope: (f = msg.getKeyScope()) && dev$portal_api_dev$portal_v1_common_pb.ObjectRef.toObject(includeInstance, f)
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
 * @return {!proto.admin.devportal.solo.io.ApiKey}
 */
proto.admin.devportal.solo.io.ApiKey.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.admin.devportal.solo.io.ApiKey;
  return proto.admin.devportal.solo.io.ApiKey.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.admin.devportal.solo.io.ApiKey} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.admin.devportal.solo.io.ApiKey}
 */
proto.admin.devportal.solo.io.ApiKey.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new dev$portal_api_grpc_common_common_pb.ObjectMeta;
      reader.readMessage(value,dev$portal_api_grpc_common_common_pb.ObjectMeta.deserializeBinaryFromReader);
      msg.setMetadata(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setValue(value);
      break;
    case 3:
      var value = new dev$portal_api_dev$portal_v1_common_pb.ObjectRef;
      reader.readMessage(value,dev$portal_api_dev$portal_v1_common_pb.ObjectRef.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    case 4:
      var value = new dev$portal_api_dev$portal_v1_common_pb.ObjectRef;
      reader.readMessage(value,dev$portal_api_dev$portal_v1_common_pb.ObjectRef.deserializeBinaryFromReader);
      msg.setKeyScope(value);
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
proto.admin.devportal.solo.io.ApiKey.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.admin.devportal.solo.io.ApiKey.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.admin.devportal.solo.io.ApiKey} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.admin.devportal.solo.io.ApiKey.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMetadata();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      dev$portal_api_grpc_common_common_pb.ObjectMeta.serializeBinaryToWriter
    );
  }
  f = message.getValue();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      dev$portal_api_dev$portal_v1_common_pb.ObjectRef.serializeBinaryToWriter
    );
  }
  f = message.getKeyScope();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      dev$portal_api_dev$portal_v1_common_pb.ObjectRef.serializeBinaryToWriter
    );
  }
};


/**
 * optional common.devportal.solo.io.ObjectMeta metadata = 1;
 * @return {?proto.common.devportal.solo.io.ObjectMeta}
 */
proto.admin.devportal.solo.io.ApiKey.prototype.getMetadata = function() {
  return /** @type{?proto.common.devportal.solo.io.ObjectMeta} */ (
    jspb.Message.getWrapperField(this, dev$portal_api_grpc_common_common_pb.ObjectMeta, 1));
};


/** @param {?proto.common.devportal.solo.io.ObjectMeta|undefined} value */
proto.admin.devportal.solo.io.ApiKey.prototype.setMetadata = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.admin.devportal.solo.io.ApiKey.prototype.clearMetadata = function() {
  this.setMetadata(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.admin.devportal.solo.io.ApiKey.prototype.hasMetadata = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string value = 2;
 * @return {string}
 */
proto.admin.devportal.solo.io.ApiKey.prototype.getValue = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.admin.devportal.solo.io.ApiKey.prototype.setValue = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional devportal.solo.io.ObjectRef user = 3;
 * @return {?proto.devportal.solo.io.ObjectRef}
 */
proto.admin.devportal.solo.io.ApiKey.prototype.getUser = function() {
  return /** @type{?proto.devportal.solo.io.ObjectRef} */ (
    jspb.Message.getWrapperField(this, dev$portal_api_dev$portal_v1_common_pb.ObjectRef, 3));
};


/** @param {?proto.devportal.solo.io.ObjectRef|undefined} value */
proto.admin.devportal.solo.io.ApiKey.prototype.setUser = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.admin.devportal.solo.io.ApiKey.prototype.clearUser = function() {
  this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.admin.devportal.solo.io.ApiKey.prototype.hasUser = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional devportal.solo.io.ObjectRef key_scope = 4;
 * @return {?proto.devportal.solo.io.ObjectRef}
 */
proto.admin.devportal.solo.io.ApiKey.prototype.getKeyScope = function() {
  return /** @type{?proto.devportal.solo.io.ObjectRef} */ (
    jspb.Message.getWrapperField(this, dev$portal_api_dev$portal_v1_common_pb.ObjectRef, 4));
};


/** @param {?proto.devportal.solo.io.ObjectRef|undefined} value */
proto.admin.devportal.solo.io.ApiKey.prototype.setKeyScope = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.admin.devportal.solo.io.ApiKey.prototype.clearKeyScope = function() {
  this.setKeyScope(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.admin.devportal.solo.io.ApiKey.prototype.hasKeyScope = function() {
  return jspb.Message.getField(this, 4) != null;
};



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
proto.admin.devportal.solo.io.ApiKeyList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.admin.devportal.solo.io.ApiKeyList.repeatedFields_, null);
};
goog.inherits(proto.admin.devportal.solo.io.ApiKeyList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.admin.devportal.solo.io.ApiKeyList.displayName = 'proto.admin.devportal.solo.io.ApiKeyList';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.admin.devportal.solo.io.ApiKeyList.repeatedFields_ = [1];



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
proto.admin.devportal.solo.io.ApiKeyList.prototype.toObject = function(opt_includeInstance) {
  return proto.admin.devportal.solo.io.ApiKeyList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.admin.devportal.solo.io.ApiKeyList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.admin.devportal.solo.io.ApiKeyList.toObject = function(includeInstance, msg) {
  var f, obj = {
    apiKeysList: jspb.Message.toObjectList(msg.getApiKeysList(),
    proto.admin.devportal.solo.io.ApiKey.toObject, includeInstance)
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
 * @return {!proto.admin.devportal.solo.io.ApiKeyList}
 */
proto.admin.devportal.solo.io.ApiKeyList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.admin.devportal.solo.io.ApiKeyList;
  return proto.admin.devportal.solo.io.ApiKeyList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.admin.devportal.solo.io.ApiKeyList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.admin.devportal.solo.io.ApiKeyList}
 */
proto.admin.devportal.solo.io.ApiKeyList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.admin.devportal.solo.io.ApiKey;
      reader.readMessage(value,proto.admin.devportal.solo.io.ApiKey.deserializeBinaryFromReader);
      msg.addApiKeys(value);
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
proto.admin.devportal.solo.io.ApiKeyList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.admin.devportal.solo.io.ApiKeyList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.admin.devportal.solo.io.ApiKeyList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.admin.devportal.solo.io.ApiKeyList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApiKeysList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.admin.devportal.solo.io.ApiKey.serializeBinaryToWriter
    );
  }
};


/**
 * repeated ApiKey api_keys = 1;
 * @return {!Array<!proto.admin.devportal.solo.io.ApiKey>}
 */
proto.admin.devportal.solo.io.ApiKeyList.prototype.getApiKeysList = function() {
  return /** @type{!Array<!proto.admin.devportal.solo.io.ApiKey>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.admin.devportal.solo.io.ApiKey, 1));
};


/** @param {!Array<!proto.admin.devportal.solo.io.ApiKey>} value */
proto.admin.devportal.solo.io.ApiKeyList.prototype.setApiKeysList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.admin.devportal.solo.io.ApiKey=} opt_value
 * @param {number=} opt_index
 * @return {!proto.admin.devportal.solo.io.ApiKey}
 */
proto.admin.devportal.solo.io.ApiKeyList.prototype.addApiKeys = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.admin.devportal.solo.io.ApiKey, opt_index);
};


proto.admin.devportal.solo.io.ApiKeyList.prototype.clearApiKeysList = function() {
  this.setApiKeysList([]);
};


goog.object.extend(exports, proto.admin.devportal.solo.io);
