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

var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var github_com_solo$io_solo$kit_api_external_envoy_api_v2_core_base_pb = require('../../../../../../../../../github.com/solo-io/solo-kit/api/external/envoy/api/v2/core/base_pb.js');
var extproto_ext_pb = require('../../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.tcp.options.gloo.solo.io.HeaderValue', null, global);
goog.exportSymbol('proto.tcp.options.gloo.solo.io.HeaderValueOption', null, global);
goog.exportSymbol('proto.tcp.options.gloo.solo.io.TcpProxySettings', null, global);
goog.exportSymbol('proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig', null, global);

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
proto.tcp.options.gloo.solo.io.TcpProxySettings = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tcp.options.gloo.solo.io.TcpProxySettings, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.tcp.options.gloo.solo.io.TcpProxySettings.displayName = 'proto.tcp.options.gloo.solo.io.TcpProxySettings';
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
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.toObject = function(opt_includeInstance) {
  return proto.tcp.options.gloo.solo.io.TcpProxySettings.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tcp.options.gloo.solo.io.TcpProxySettings} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.toObject = function(includeInstance, msg) {
  var f, obj = {
    maxConnectAttempts: (f = msg.getMaxConnectAttempts()) && google_protobuf_wrappers_pb.UInt32Value.toObject(includeInstance, f),
    idleTimeout: (f = msg.getIdleTimeout()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    tunnelingConfig: (f = msg.getTunnelingConfig()) && proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.toObject(includeInstance, f)
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
 * @return {!proto.tcp.options.gloo.solo.io.TcpProxySettings}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tcp.options.gloo.solo.io.TcpProxySettings;
  return proto.tcp.options.gloo.solo.io.TcpProxySettings.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tcp.options.gloo.solo.io.TcpProxySettings} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tcp.options.gloo.solo.io.TcpProxySettings}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.UInt32Value;
      reader.readMessage(value,google_protobuf_wrappers_pb.UInt32Value.deserializeBinaryFromReader);
      msg.setMaxConnectAttempts(value);
      break;
    case 2:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setIdleTimeout(value);
      break;
    case 12:
      var value = new proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig;
      reader.readMessage(value,proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.deserializeBinaryFromReader);
      msg.setTunnelingConfig(value);
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
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tcp.options.gloo.solo.io.TcpProxySettings.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tcp.options.gloo.solo.io.TcpProxySettings} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMaxConnectAttempts();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.UInt32Value.serializeBinaryToWriter
    );
  }
  f = message.getIdleTimeout();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getTunnelingConfig();
  if (f != null) {
    writer.writeMessage(
      12,
      f,
      proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.serializeBinaryToWriter
    );
  }
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
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.repeatedFields_, null);
};
goog.inherits(proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.displayName = 'proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.repeatedFields_ = [13];



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
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    hostname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    headersToAddList: jspb.Message.toObjectList(msg.getHeadersToAddList(),
    proto.tcp.options.gloo.solo.io.HeaderValueOption.toObject, includeInstance)
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
 * @return {!proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig;
  return proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setHostname(value);
      break;
    case 13:
      var value = new proto.tcp.options.gloo.solo.io.HeaderValueOption;
      reader.readMessage(value,proto.tcp.options.gloo.solo.io.HeaderValueOption.deserializeBinaryFromReader);
      msg.addHeadersToAdd(value);
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
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHostname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getHeadersToAddList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      13,
      f,
      proto.tcp.options.gloo.solo.io.HeaderValueOption.serializeBinaryToWriter
    );
  }
};


/**
 * optional string hostname = 1;
 * @return {string}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.getHostname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.setHostname = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated HeaderValueOption headers_to_add = 13;
 * @return {!Array<!proto.tcp.options.gloo.solo.io.HeaderValueOption>}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.getHeadersToAddList = function() {
  return /** @type{!Array<!proto.tcp.options.gloo.solo.io.HeaderValueOption>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.tcp.options.gloo.solo.io.HeaderValueOption, 13));
};


/** @param {!Array<!proto.tcp.options.gloo.solo.io.HeaderValueOption>} value */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.setHeadersToAddList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 13, value);
};


/**
 * @param {!proto.tcp.options.gloo.solo.io.HeaderValueOption=} opt_value
 * @param {number=} opt_index
 * @return {!proto.tcp.options.gloo.solo.io.HeaderValueOption}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.addHeadersToAdd = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 13, opt_value, proto.tcp.options.gloo.solo.io.HeaderValueOption, opt_index);
};


proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig.prototype.clearHeadersToAddList = function() {
  this.setHeadersToAddList([]);
};


/**
 * optional google.protobuf.UInt32Value max_connect_attempts = 1;
 * @return {?proto.google.protobuf.UInt32Value}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.getMaxConnectAttempts = function() {
  return /** @type{?proto.google.protobuf.UInt32Value} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.UInt32Value, 1));
};


/** @param {?proto.google.protobuf.UInt32Value|undefined} value */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.setMaxConnectAttempts = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.clearMaxConnectAttempts = function() {
  this.setMaxConnectAttempts(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.hasMaxConnectAttempts = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.Duration idle_timeout = 2;
 * @return {?proto.google.protobuf.Duration}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.getIdleTimeout = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 2));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.setIdleTimeout = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.clearIdleTimeout = function() {
  this.setIdleTimeout(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.hasIdleTimeout = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional TunnelingConfig tunneling_config = 12;
 * @return {?proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.getTunnelingConfig = function() {
  return /** @type{?proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig} */ (
    jspb.Message.getWrapperField(this, proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig, 12));
};


/** @param {?proto.tcp.options.gloo.solo.io.TcpProxySettings.TunnelingConfig|undefined} value */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.setTunnelingConfig = function(value) {
  jspb.Message.setWrapperField(this, 12, value);
};


proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.clearTunnelingConfig = function() {
  this.setTunnelingConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.tcp.options.gloo.solo.io.TcpProxySettings.prototype.hasTunnelingConfig = function() {
  return jspb.Message.getField(this, 12) != null;
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
proto.tcp.options.gloo.solo.io.HeaderValueOption = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tcp.options.gloo.solo.io.HeaderValueOption, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.tcp.options.gloo.solo.io.HeaderValueOption.displayName = 'proto.tcp.options.gloo.solo.io.HeaderValueOption';
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
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.toObject = function(opt_includeInstance) {
  return proto.tcp.options.gloo.solo.io.HeaderValueOption.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tcp.options.gloo.solo.io.HeaderValueOption} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.toObject = function(includeInstance, msg) {
  var f, obj = {
    header: (f = msg.getHeader()) && proto.tcp.options.gloo.solo.io.HeaderValue.toObject(includeInstance, f),
    append: (f = msg.getAppend()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f)
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
 * @return {!proto.tcp.options.gloo.solo.io.HeaderValueOption}
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tcp.options.gloo.solo.io.HeaderValueOption;
  return proto.tcp.options.gloo.solo.io.HeaderValueOption.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tcp.options.gloo.solo.io.HeaderValueOption} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tcp.options.gloo.solo.io.HeaderValueOption}
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.tcp.options.gloo.solo.io.HeaderValue;
      reader.readMessage(value,proto.tcp.options.gloo.solo.io.HeaderValue.deserializeBinaryFromReader);
      msg.setHeader(value);
      break;
    case 2:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setAppend(value);
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
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tcp.options.gloo.solo.io.HeaderValueOption.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tcp.options.gloo.solo.io.HeaderValueOption} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHeader();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.tcp.options.gloo.solo.io.HeaderValue.serializeBinaryToWriter
    );
  }
  f = message.getAppend();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
};


/**
 * optional HeaderValue header = 1;
 * @return {?proto.tcp.options.gloo.solo.io.HeaderValue}
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.getHeader = function() {
  return /** @type{?proto.tcp.options.gloo.solo.io.HeaderValue} */ (
    jspb.Message.getWrapperField(this, proto.tcp.options.gloo.solo.io.HeaderValue, 1));
};


/** @param {?proto.tcp.options.gloo.solo.io.HeaderValue|undefined} value */
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.setHeader = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.clearHeader = function() {
  this.setHeader(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.hasHeader = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.BoolValue append = 2;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.getAppend = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 2));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.setAppend = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.clearAppend = function() {
  this.setAppend(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.tcp.options.gloo.solo.io.HeaderValueOption.prototype.hasAppend = function() {
  return jspb.Message.getField(this, 2) != null;
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
proto.tcp.options.gloo.solo.io.HeaderValue = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tcp.options.gloo.solo.io.HeaderValue, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.tcp.options.gloo.solo.io.HeaderValue.displayName = 'proto.tcp.options.gloo.solo.io.HeaderValue';
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
proto.tcp.options.gloo.solo.io.HeaderValue.prototype.toObject = function(opt_includeInstance) {
  return proto.tcp.options.gloo.solo.io.HeaderValue.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tcp.options.gloo.solo.io.HeaderValue} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.HeaderValue.toObject = function(includeInstance, msg) {
  var f, obj = {
    key: jspb.Message.getFieldWithDefault(msg, 1, ""),
    value: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.tcp.options.gloo.solo.io.HeaderValue}
 */
proto.tcp.options.gloo.solo.io.HeaderValue.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tcp.options.gloo.solo.io.HeaderValue;
  return proto.tcp.options.gloo.solo.io.HeaderValue.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tcp.options.gloo.solo.io.HeaderValue} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tcp.options.gloo.solo.io.HeaderValue}
 */
proto.tcp.options.gloo.solo.io.HeaderValue.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setKey(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setValue(value);
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
proto.tcp.options.gloo.solo.io.HeaderValue.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tcp.options.gloo.solo.io.HeaderValue.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tcp.options.gloo.solo.io.HeaderValue} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tcp.options.gloo.solo.io.HeaderValue.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKey();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getValue();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string key = 1;
 * @return {string}
 */
proto.tcp.options.gloo.solo.io.HeaderValue.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.tcp.options.gloo.solo.io.HeaderValue.prototype.setKey = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string value = 2;
 * @return {string}
 */
proto.tcp.options.gloo.solo.io.HeaderValue.prototype.getValue = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.tcp.options.gloo.solo.io.HeaderValue.prototype.setValue = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


goog.object.extend(exports, proto.tcp.options.gloo.solo.io);