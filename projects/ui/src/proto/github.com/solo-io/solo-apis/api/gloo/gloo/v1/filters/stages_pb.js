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

var extproto_ext_pb = require('../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.filters.gloo.solo.io.FilterStage', null, global);
goog.exportSymbol('proto.filters.gloo.solo.io.FilterStage.Predicate', null, global);
goog.exportSymbol('proto.filters.gloo.solo.io.FilterStage.Stage', null, global);

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
proto.filters.gloo.solo.io.FilterStage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.filters.gloo.solo.io.FilterStage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.filters.gloo.solo.io.FilterStage.displayName = 'proto.filters.gloo.solo.io.FilterStage';
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
proto.filters.gloo.solo.io.FilterStage.prototype.toObject = function(opt_includeInstance) {
  return proto.filters.gloo.solo.io.FilterStage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.filters.gloo.solo.io.FilterStage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.filters.gloo.solo.io.FilterStage.toObject = function(includeInstance, msg) {
  var f, obj = {
    stage: jspb.Message.getFieldWithDefault(msg, 1, 0),
    predicate: jspb.Message.getFieldWithDefault(msg, 2, 0)
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
 * @return {!proto.filters.gloo.solo.io.FilterStage}
 */
proto.filters.gloo.solo.io.FilterStage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.filters.gloo.solo.io.FilterStage;
  return proto.filters.gloo.solo.io.FilterStage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.filters.gloo.solo.io.FilterStage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.filters.gloo.solo.io.FilterStage}
 */
proto.filters.gloo.solo.io.FilterStage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.filters.gloo.solo.io.FilterStage.Stage} */ (reader.readEnum());
      msg.setStage(value);
      break;
    case 2:
      var value = /** @type {!proto.filters.gloo.solo.io.FilterStage.Predicate} */ (reader.readEnum());
      msg.setPredicate(value);
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
proto.filters.gloo.solo.io.FilterStage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.filters.gloo.solo.io.FilterStage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.filters.gloo.solo.io.FilterStage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.filters.gloo.solo.io.FilterStage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getStage();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getPredicate();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.filters.gloo.solo.io.FilterStage.Stage = {
  FAULTSTAGE: 0,
  CORSSTAGE: 1,
  WAFSTAGE: 2,
  AUTHNSTAGE: 3,
  AUTHZSTAGE: 4,
  RATELIMITSTAGE: 5,
  ACCEPTEDSTAGE: 6,
  OUTAUTHSTAGE: 7,
  ROUTESTAGE: 8
};

/**
 * @enum {number}
 */
proto.filters.gloo.solo.io.FilterStage.Predicate = {
  DURING: 0,
  BEFORE: 1,
  AFTER: 2
};

/**
 * optional Stage stage = 1;
 * @return {!proto.filters.gloo.solo.io.FilterStage.Stage}
 */
proto.filters.gloo.solo.io.FilterStage.prototype.getStage = function() {
  return /** @type {!proto.filters.gloo.solo.io.FilterStage.Stage} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.filters.gloo.solo.io.FilterStage.Stage} value */
proto.filters.gloo.solo.io.FilterStage.prototype.setStage = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional Predicate predicate = 2;
 * @return {!proto.filters.gloo.solo.io.FilterStage.Predicate}
 */
proto.filters.gloo.solo.io.FilterStage.prototype.getPredicate = function() {
  return /** @type {!proto.filters.gloo.solo.io.FilterStage.Predicate} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {!proto.filters.gloo.solo.io.FilterStage.Predicate} value */
proto.filters.gloo.solo.io.FilterStage.prototype.setPredicate = function(value) {
  jspb.Message.setProto3EnumField(this, 2, value);
};


goog.object.extend(exports, proto.filters.gloo.solo.io);