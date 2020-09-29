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

var envoy_config_core_v3_base_pb = require('../../../../envoy/config/core/v3/base_pb.js');
var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js');
var validate_validate_pb = require('../../../../validate/validate_pb.js');
var gogoproto_gogo_pb = require('../../../../gogoproto/gogo_pb.js');
goog.exportSymbol('proto.envoy.extensions.wasm.v3.PluginConfig', null, global);
goog.exportSymbol('proto.envoy.extensions.wasm.v3.VmConfig', null, global);
goog.exportSymbol('proto.envoy.extensions.wasm.v3.WasmService', null, global);

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
proto.envoy.extensions.wasm.v3.VmConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.extensions.wasm.v3.VmConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.extensions.wasm.v3.VmConfig.displayName = 'proto.envoy.extensions.wasm.v3.VmConfig';
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
proto.envoy.extensions.wasm.v3.VmConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.extensions.wasm.v3.VmConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.extensions.wasm.v3.VmConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.extensions.wasm.v3.VmConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    vmId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    runtime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    code: (f = msg.getCode()) && envoy_config_core_v3_base_pb.AsyncDataSource.toObject(includeInstance, f),
    configuration: (f = msg.getConfiguration()) && google_protobuf_any_pb.Any.toObject(includeInstance, f),
    allowPrecompiled: jspb.Message.getFieldWithDefault(msg, 5, false),
    nackOnCodeCacheMiss: jspb.Message.getFieldWithDefault(msg, 6, false)
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
 * @return {!proto.envoy.extensions.wasm.v3.VmConfig}
 */
proto.envoy.extensions.wasm.v3.VmConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.extensions.wasm.v3.VmConfig;
  return proto.envoy.extensions.wasm.v3.VmConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.extensions.wasm.v3.VmConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.extensions.wasm.v3.VmConfig}
 */
proto.envoy.extensions.wasm.v3.VmConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setVmId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRuntime(value);
      break;
    case 3:
      var value = new envoy_config_core_v3_base_pb.AsyncDataSource;
      reader.readMessage(value,envoy_config_core_v3_base_pb.AsyncDataSource.deserializeBinaryFromReader);
      msg.setCode(value);
      break;
    case 4:
      var value = new google_protobuf_any_pb.Any;
      reader.readMessage(value,google_protobuf_any_pb.Any.deserializeBinaryFromReader);
      msg.setConfiguration(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAllowPrecompiled(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setNackOnCodeCacheMiss(value);
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
proto.envoy.extensions.wasm.v3.VmConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.extensions.wasm.v3.VmConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.extensions.wasm.v3.VmConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.extensions.wasm.v3.VmConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVmId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRuntime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCode();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      envoy_config_core_v3_base_pb.AsyncDataSource.serializeBinaryToWriter
    );
  }
  f = message.getConfiguration();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_any_pb.Any.serializeBinaryToWriter
    );
  }
  f = message.getAllowPrecompiled();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getNackOnCodeCacheMiss();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
};


/**
 * optional string vm_id = 1;
 * @return {string}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.getVmId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.setVmId = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string runtime = 2;
 * @return {string}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.getRuntime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.setRuntime = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional envoy.config.core.v3.AsyncDataSource code = 3;
 * @return {?proto.envoy.config.core.v3.AsyncDataSource}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.getCode = function() {
  return /** @type{?proto.envoy.config.core.v3.AsyncDataSource} */ (
    jspb.Message.getWrapperField(this, envoy_config_core_v3_base_pb.AsyncDataSource, 3));
};


/** @param {?proto.envoy.config.core.v3.AsyncDataSource|undefined} value */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.setCode = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.envoy.extensions.wasm.v3.VmConfig.prototype.clearCode = function() {
  this.setCode(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.hasCode = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Any configuration = 4;
 * @return {?proto.google.protobuf.Any}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.getConfiguration = function() {
  return /** @type{?proto.google.protobuf.Any} */ (
    jspb.Message.getWrapperField(this, google_protobuf_any_pb.Any, 4));
};


/** @param {?proto.google.protobuf.Any|undefined} value */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.setConfiguration = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.envoy.extensions.wasm.v3.VmConfig.prototype.clearConfiguration = function() {
  this.setConfiguration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.hasConfiguration = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional bool allow_precompiled = 5;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.getAllowPrecompiled = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 5, false));
};


/** @param {boolean} value */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.setAllowPrecompiled = function(value) {
  jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional bool nack_on_code_cache_miss = 6;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.getNackOnCodeCacheMiss = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 6, false));
};


/** @param {boolean} value */
proto.envoy.extensions.wasm.v3.VmConfig.prototype.setNackOnCodeCacheMiss = function(value) {
  jspb.Message.setProto3BooleanField(this, 6, value);
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
proto.envoy.extensions.wasm.v3.PluginConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.envoy.extensions.wasm.v3.PluginConfig.oneofGroups_);
};
goog.inherits(proto.envoy.extensions.wasm.v3.PluginConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.extensions.wasm.v3.PluginConfig.displayName = 'proto.envoy.extensions.wasm.v3.PluginConfig';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.envoy.extensions.wasm.v3.PluginConfig.oneofGroups_ = [[3]];

/**
 * @enum {number}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.VmCase = {
  VM_NOT_SET: 0,
  VM_CONFIG: 3
};

/**
 * @return {proto.envoy.extensions.wasm.v3.PluginConfig.VmCase}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.getVmCase = function() {
  return /** @type {proto.envoy.extensions.wasm.v3.PluginConfig.VmCase} */(jspb.Message.computeOneofCase(this, proto.envoy.extensions.wasm.v3.PluginConfig.oneofGroups_[0]));
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
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.extensions.wasm.v3.PluginConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.extensions.wasm.v3.PluginConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.extensions.wasm.v3.PluginConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    rootId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    vmConfig: (f = msg.getVmConfig()) && proto.envoy.extensions.wasm.v3.VmConfig.toObject(includeInstance, f),
    configuration: (f = msg.getConfiguration()) && google_protobuf_any_pb.Any.toObject(includeInstance, f),
    failOpen: jspb.Message.getFieldWithDefault(msg, 5, false)
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
 * @return {!proto.envoy.extensions.wasm.v3.PluginConfig}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.extensions.wasm.v3.PluginConfig;
  return proto.envoy.extensions.wasm.v3.PluginConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.extensions.wasm.v3.PluginConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.extensions.wasm.v3.PluginConfig}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRootId(value);
      break;
    case 3:
      var value = new proto.envoy.extensions.wasm.v3.VmConfig;
      reader.readMessage(value,proto.envoy.extensions.wasm.v3.VmConfig.deserializeBinaryFromReader);
      msg.setVmConfig(value);
      break;
    case 4:
      var value = new google_protobuf_any_pb.Any;
      reader.readMessage(value,google_protobuf_any_pb.Any.deserializeBinaryFromReader);
      msg.setConfiguration(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setFailOpen(value);
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
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.extensions.wasm.v3.PluginConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.extensions.wasm.v3.PluginConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.extensions.wasm.v3.PluginConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRootId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getVmConfig();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.envoy.extensions.wasm.v3.VmConfig.serializeBinaryToWriter
    );
  }
  f = message.getConfiguration();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_any_pb.Any.serializeBinaryToWriter
    );
  }
  f = message.getFailOpen();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string root_id = 2;
 * @return {string}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.getRootId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.setRootId = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional VmConfig vm_config = 3;
 * @return {?proto.envoy.extensions.wasm.v3.VmConfig}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.getVmConfig = function() {
  return /** @type{?proto.envoy.extensions.wasm.v3.VmConfig} */ (
    jspb.Message.getWrapperField(this, proto.envoy.extensions.wasm.v3.VmConfig, 3));
};


/** @param {?proto.envoy.extensions.wasm.v3.VmConfig|undefined} value */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.setVmConfig = function(value) {
  jspb.Message.setOneofWrapperField(this, 3, proto.envoy.extensions.wasm.v3.PluginConfig.oneofGroups_[0], value);
};


proto.envoy.extensions.wasm.v3.PluginConfig.prototype.clearVmConfig = function() {
  this.setVmConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.hasVmConfig = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Any configuration = 4;
 * @return {?proto.google.protobuf.Any}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.getConfiguration = function() {
  return /** @type{?proto.google.protobuf.Any} */ (
    jspb.Message.getWrapperField(this, google_protobuf_any_pb.Any, 4));
};


/** @param {?proto.google.protobuf.Any|undefined} value */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.setConfiguration = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.envoy.extensions.wasm.v3.PluginConfig.prototype.clearConfiguration = function() {
  this.setConfiguration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.hasConfiguration = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional bool fail_open = 5;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.getFailOpen = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 5, false));
};


/** @param {boolean} value */
proto.envoy.extensions.wasm.v3.PluginConfig.prototype.setFailOpen = function(value) {
  jspb.Message.setProto3BooleanField(this, 5, value);
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
proto.envoy.extensions.wasm.v3.WasmService = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.extensions.wasm.v3.WasmService, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.extensions.wasm.v3.WasmService.displayName = 'proto.envoy.extensions.wasm.v3.WasmService';
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
proto.envoy.extensions.wasm.v3.WasmService.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.extensions.wasm.v3.WasmService.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.extensions.wasm.v3.WasmService} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.extensions.wasm.v3.WasmService.toObject = function(includeInstance, msg) {
  var f, obj = {
    config: (f = msg.getConfig()) && proto.envoy.extensions.wasm.v3.PluginConfig.toObject(includeInstance, f),
    singleton: jspb.Message.getFieldWithDefault(msg, 2, false)
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
 * @return {!proto.envoy.extensions.wasm.v3.WasmService}
 */
proto.envoy.extensions.wasm.v3.WasmService.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.extensions.wasm.v3.WasmService;
  return proto.envoy.extensions.wasm.v3.WasmService.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.extensions.wasm.v3.WasmService} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.extensions.wasm.v3.WasmService}
 */
proto.envoy.extensions.wasm.v3.WasmService.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.envoy.extensions.wasm.v3.PluginConfig;
      reader.readMessage(value,proto.envoy.extensions.wasm.v3.PluginConfig.deserializeBinaryFromReader);
      msg.setConfig(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSingleton(value);
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
proto.envoy.extensions.wasm.v3.WasmService.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.extensions.wasm.v3.WasmService.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.extensions.wasm.v3.WasmService} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.extensions.wasm.v3.WasmService.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getConfig();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.envoy.extensions.wasm.v3.PluginConfig.serializeBinaryToWriter
    );
  }
  f = message.getSingleton();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
};


/**
 * optional PluginConfig config = 1;
 * @return {?proto.envoy.extensions.wasm.v3.PluginConfig}
 */
proto.envoy.extensions.wasm.v3.WasmService.prototype.getConfig = function() {
  return /** @type{?proto.envoy.extensions.wasm.v3.PluginConfig} */ (
    jspb.Message.getWrapperField(this, proto.envoy.extensions.wasm.v3.PluginConfig, 1));
};


/** @param {?proto.envoy.extensions.wasm.v3.PluginConfig|undefined} value */
proto.envoy.extensions.wasm.v3.WasmService.prototype.setConfig = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.envoy.extensions.wasm.v3.WasmService.prototype.clearConfig = function() {
  this.setConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.extensions.wasm.v3.WasmService.prototype.hasConfig = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional bool singleton = 2;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.envoy.extensions.wasm.v3.WasmService.prototype.getSingleton = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 2, false));
};


/** @param {boolean} value */
proto.envoy.extensions.wasm.v3.WasmService.prototype.setSingleton = function(value) {
  jspb.Message.setProto3BooleanField(this, 2, value);
};


goog.object.extend(exports, proto.envoy.extensions.wasm.v3);
