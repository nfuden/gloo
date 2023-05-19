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

var udpa_annotations_status_pb = require('../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/status_pb.js');
var udpa_annotations_versioning_pb = require('../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/versioning_pb.js');
var extproto_ext_pb = require('../../../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.solo.io.envoy.type.v3.SemanticVersion', null, global);

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
proto.solo.io.envoy.type.v3.SemanticVersion = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.solo.io.envoy.type.v3.SemanticVersion, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.solo.io.envoy.type.v3.SemanticVersion.displayName = 'proto.solo.io.envoy.type.v3.SemanticVersion';
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
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.toObject = function(opt_includeInstance) {
  return proto.solo.io.envoy.type.v3.SemanticVersion.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.solo.io.envoy.type.v3.SemanticVersion} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.type.v3.SemanticVersion.toObject = function(includeInstance, msg) {
  var f, obj = {
    majorNumber: jspb.Message.getFieldWithDefault(msg, 1, 0),
    minorNumber: jspb.Message.getFieldWithDefault(msg, 2, 0),
    patch: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.solo.io.envoy.type.v3.SemanticVersion}
 */
proto.solo.io.envoy.type.v3.SemanticVersion.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.solo.io.envoy.type.v3.SemanticVersion;
  return proto.solo.io.envoy.type.v3.SemanticVersion.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.solo.io.envoy.type.v3.SemanticVersion} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.solo.io.envoy.type.v3.SemanticVersion}
 */
proto.solo.io.envoy.type.v3.SemanticVersion.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setMajorNumber(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setMinorNumber(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setPatch(value);
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
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.solo.io.envoy.type.v3.SemanticVersion.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.solo.io.envoy.type.v3.SemanticVersion} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.type.v3.SemanticVersion.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMajorNumber();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getMinorNumber();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
  f = message.getPatch();
  if (f !== 0) {
    writer.writeUint32(
      3,
      f
    );
  }
};


/**
 * optional uint32 major_number = 1;
 * @return {number}
 */
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.getMajorNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.setMajorNumber = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional uint32 minor_number = 2;
 * @return {number}
 */
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.getMinorNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.setMinorNumber = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional uint32 patch = 3;
 * @return {number}
 */
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.getPatch = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.solo.io.envoy.type.v3.SemanticVersion.prototype.setPatch = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


goog.object.extend(exports, proto.solo.io.envoy.type.v3);