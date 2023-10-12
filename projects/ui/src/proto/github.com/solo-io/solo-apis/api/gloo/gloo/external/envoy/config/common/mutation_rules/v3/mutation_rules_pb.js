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

var github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_config_core_v3_base_pb = require('../../../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/config/core/v3/base_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb = require('../../../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/type/matcher/v3/regex_pb.js');
var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js');
var udpa_annotations_status_pb = require('../../../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/external/udpa/annotations/status_pb.js');
var validate_validate_pb = require('../../../../../../../../../../../../validate/validate_pb.js');
var extproto_ext_pb = require('../../../../../../../../../../../../extproto/ext_pb.js');
goog.exportSymbol('proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation', null, global);
goog.exportSymbol('proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules', null, global);

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
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.displayName = 'proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules';
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
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.toObject = function(opt_includeInstance) {
  return proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.toObject = function(includeInstance, msg) {
  var f, obj = {
    allowAllRouting: (f = msg.getAllowAllRouting()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    allowEnvoy: (f = msg.getAllowEnvoy()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    disallowSystem: (f = msg.getDisallowSystem()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    disallowAll: (f = msg.getDisallowAll()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f),
    allowExpression: (f = msg.getAllowExpression()) && github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.toObject(includeInstance, f),
    disallowExpression: (f = msg.getDisallowExpression()) && github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.toObject(includeInstance, f),
    disallowIsError: (f = msg.getDisallowIsError()) && google_protobuf_wrappers_pb.BoolValue.toObject(includeInstance, f)
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
 * @return {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules;
  return proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setAllowAllRouting(value);
      break;
    case 2:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setAllowEnvoy(value);
      break;
    case 3:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setDisallowSystem(value);
      break;
    case 4:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setDisallowAll(value);
      break;
    case 5:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.deserializeBinaryFromReader);
      msg.setAllowExpression(value);
      break;
    case 6:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.deserializeBinaryFromReader);
      msg.setDisallowExpression(value);
      break;
    case 7:
      var value = new google_protobuf_wrappers_pb.BoolValue;
      reader.readMessage(value,google_protobuf_wrappers_pb.BoolValue.deserializeBinaryFromReader);
      msg.setDisallowIsError(value);
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
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAllowAllRouting();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getAllowEnvoy();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getDisallowSystem();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getDisallowAll();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
  f = message.getAllowExpression();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.serializeBinaryToWriter
    );
  }
  f = message.getDisallowExpression();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher.serializeBinaryToWriter
    );
  }
  f = message.getDisallowIsError();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      google_protobuf_wrappers_pb.BoolValue.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.BoolValue allow_all_routing = 1;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.getAllowAllRouting = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 1));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.setAllowAllRouting = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.clearAllowAllRouting = function() {
  this.setAllowAllRouting(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.hasAllowAllRouting = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.BoolValue allow_envoy = 2;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.getAllowEnvoy = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 2));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.setAllowEnvoy = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.clearAllowEnvoy = function() {
  this.setAllowEnvoy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.hasAllowEnvoy = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.BoolValue disallow_system = 3;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.getDisallowSystem = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 3));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.setDisallowSystem = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.clearDisallowSystem = function() {
  this.setDisallowSystem(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.hasDisallowSystem = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.BoolValue disallow_all = 4;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.getDisallowAll = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 4));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.setDisallowAll = function(value) {
  jspb.Message.setWrapperField(this, 4, value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.clearDisallowAll = function() {
  this.setDisallowAll(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.hasDisallowAll = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional solo.io.envoy.type.matcher.v3.RegexMatcher allow_expression = 5;
 * @return {?proto.solo.io.envoy.type.matcher.v3.RegexMatcher}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.getAllowExpression = function() {
  return /** @type{?proto.solo.io.envoy.type.matcher.v3.RegexMatcher} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher, 5));
};


/** @param {?proto.solo.io.envoy.type.matcher.v3.RegexMatcher|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.setAllowExpression = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.clearAllowExpression = function() {
  this.setAllowExpression(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.hasAllowExpression = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional solo.io.envoy.type.matcher.v3.RegexMatcher disallow_expression = 6;
 * @return {?proto.solo.io.envoy.type.matcher.v3.RegexMatcher}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.getDisallowExpression = function() {
  return /** @type{?proto.solo.io.envoy.type.matcher.v3.RegexMatcher} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_type_matcher_v3_regex_pb.RegexMatcher, 6));
};


/** @param {?proto.solo.io.envoy.type.matcher.v3.RegexMatcher|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.setDisallowExpression = function(value) {
  jspb.Message.setWrapperField(this, 6, value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.clearDisallowExpression = function() {
  this.setDisallowExpression(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.hasDisallowExpression = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional google.protobuf.BoolValue disallow_is_error = 7;
 * @return {?proto.google.protobuf.BoolValue}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.getDisallowIsError = function() {
  return /** @type{?proto.google.protobuf.BoolValue} */ (
    jspb.Message.getWrapperField(this, google_protobuf_wrappers_pb.BoolValue, 7));
};


/** @param {?proto.google.protobuf.BoolValue|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.setDisallowIsError = function(value) {
  jspb.Message.setWrapperField(this, 7, value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.clearDisallowIsError = function() {
  this.setDisallowIsError(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutationRules.prototype.hasDisallowIsError = function() {
  return jspb.Message.getField(this, 7) != null;
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
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.oneofGroups_);
};
goog.inherits(proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.displayName = 'proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation';
}
/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.ActionCase = {
  ACTION_NOT_SET: 0,
  REMOVE: 1,
  APPEND: 2
};

/**
 * @return {proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.ActionCase}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.getActionCase = function() {
  return /** @type {proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.ActionCase} */(jspb.Message.computeOneofCase(this, proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.oneofGroups_[0]));
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
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.toObject = function(opt_includeInstance) {
  return proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.toObject = function(includeInstance, msg) {
  var f, obj = {
    remove: jspb.Message.getFieldWithDefault(msg, 1, ""),
    append: (f = msg.getAppend()) && github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_config_core_v3_base_pb.HeaderValueOption.toObject(includeInstance, f)
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
 * @return {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation;
  return proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRemove(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_config_core_v3_base_pb.HeaderValueOption;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_config_core_v3_base_pb.HeaderValueOption.deserializeBinaryFromReader);
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
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAppend();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_config_core_v3_base_pb.HeaderValueOption.serializeBinaryToWriter
    );
  }
};


/**
 * optional string remove = 1;
 * @return {string}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.getRemove = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.setRemove = function(value) {
  jspb.Message.setOneofField(this, 1, proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.oneofGroups_[0], value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.clearRemove = function() {
  jspb.Message.setOneofField(this, 1, proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.hasRemove = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional solo.io.envoy.config.core.v3.HeaderValueOption append = 2;
 * @return {?proto.solo.io.envoy.config.core.v3.HeaderValueOption}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.getAppend = function() {
  return /** @type{?proto.solo.io.envoy.config.core.v3.HeaderValueOption} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_external_envoy_config_core_v3_base_pb.HeaderValueOption, 2));
};


/** @param {?proto.solo.io.envoy.config.core.v3.HeaderValueOption|undefined} value */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.setAppend = function(value) {
  jspb.Message.setOneofWrapperField(this, 2, proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.oneofGroups_[0], value);
};


proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.clearAppend = function() {
  this.setAppend(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.solo.io.envoy.config.common.mutation_rules.v3.HeaderMutation.prototype.hasAppend = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.solo.io.envoy.config.common.mutation_rules.v3);