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

var validate_validate_pb = require('../../../../../../../../../../validate/validate_pb.js');
var github_com_solo$io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_jwt_authn_v3_config_pb = require('../../../../../../../../../../github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/filters/http/jwt_authn/v3/config_pb.js');
goog.exportSymbol('proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage', null, global);
goog.exportSymbol('proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute', null, global);
goog.exportSymbol('proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader', null, global);
goog.exportSymbol('proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders', null, global);
goog.exportSymbol('proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute', null, global);

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
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.displayName = 'proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage';
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.toObject = function(includeInstance, msg) {
  var f, obj = {
    jwtAuthn: (f = msg.getJwtAuthn()) && github_com_solo$io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_jwt_authn_v3_config_pb.JwtAuthentication.toObject(includeInstance, f),
    stage: jspb.Message.getFieldWithDefault(msg, 2, 0)
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
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage;
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_jwt_authn_v3_config_pb.JwtAuthentication;
      reader.readMessage(value,github_com_solo$io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_jwt_authn_v3_config_pb.JwtAuthentication.deserializeBinaryFromReader);
      msg.setJwtAuthn(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setStage(value);
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getJwtAuthn();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_jwt_authn_v3_config_pb.JwtAuthentication.serializeBinaryToWriter
    );
  }
  f = message.getStage();
  if (f !== 0) {
    writer.writeUint32(
      2,
      f
    );
  }
};


/**
 * optional solo.io.envoy.extensions.filters.http.jwt_authn.v3.JwtAuthentication jwt_authn = 1;
 * @return {?proto.solo.io.envoy.extensions.filters.http.jwt_authn.v3.JwtAuthentication}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.getJwtAuthn = function() {
  return /** @type{?proto.solo.io.envoy.extensions.filters.http.jwt_authn.v3.JwtAuthentication} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_gloo_projects_gloo_api_external_envoy_extensions_filters_http_jwt_authn_v3_config_pb.JwtAuthentication, 1));
};


/** @param {?proto.solo.io.envoy.extensions.filters.http.jwt_authn.v3.JwtAuthentication|undefined} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.setJwtAuthn = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.clearJwtAuthn = function() {
  this.setJwtAuthn(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.hasJwtAuthn = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint32 stage = 2;
 * @return {number}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.getStage = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.JwtWithStage.prototype.setStage = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.displayName = 'proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute';
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.toObject = function(includeInstance, msg) {
  var f, obj = {
    requirement: jspb.Message.getFieldWithDefault(msg, 1, ""),
    claimsToHeadersMap: (f = msg.getClaimsToHeadersMap()) ? f.toObject(includeInstance, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.toObject) : [],
    clearRouteCache: jspb.Message.getFieldWithDefault(msg, 3, false),
    payloadInMetadata: jspb.Message.getFieldWithDefault(msg, 4, "")
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
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute;
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRequirement(value);
      break;
    case 2:
      var value = msg.getClaimsToHeadersMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.deserializeBinaryFromReader, "");
         });
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setClearRouteCache(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setPayloadInMetadata(value);
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequirement();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getClaimsToHeadersMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(2, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.serializeBinaryToWriter);
  }
  f = message.getClearRouteCache();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getPayloadInMetadata();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.displayName = 'proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader';
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.toObject = function(includeInstance, msg) {
  var f, obj = {
    claim: jspb.Message.getFieldWithDefault(msg, 1, ""),
    header: jspb.Message.getFieldWithDefault(msg, 2, ""),
    append: jspb.Message.getFieldWithDefault(msg, 3, false)
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
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader;
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setClaim(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setHeader(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClaim();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getHeader();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getAppend();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
};


/**
 * optional string claim = 1;
 * @return {string}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.getClaim = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.setClaim = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string header = 2;
 * @return {string}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.getHeader = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.setHeader = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bool append = 3;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.getAppend = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 3, false));
};


/** @param {boolean} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.prototype.setAppend = function(value) {
  jspb.Message.setProto3BooleanField(this, 3, value);
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.repeatedFields_, null);
};
goog.inherits(proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.displayName = 'proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.repeatedFields_ = [1];



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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.toObject = function(includeInstance, msg) {
  var f, obj = {
    claimsList: jspb.Message.toObjectList(msg.getClaimsList(),
    proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.toObject, includeInstance)
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
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders;
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader;
      reader.readMessage(value,proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.deserializeBinaryFromReader);
      msg.addClaims(value);
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClaimsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader.serializeBinaryToWriter
    );
  }
};


/**
 * repeated ClaimToHeader claims = 1;
 * @return {!Array<!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader>}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.prototype.getClaimsList = function() {
  return /** @type{!Array<!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader, 1));
};


/** @param {!Array<!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader>} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.prototype.setClaimsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader=} opt_value
 * @param {number=} opt_index
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.prototype.addClaims = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader, opt_index);
};


proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders.prototype.clearClaimsList = function() {
  this.setClaimsList([]);
};


/**
 * optional string requirement = 1;
 * @return {string}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.getRequirement = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.setRequirement = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * map<string, ClaimToHeaders> claims_to_headers = 2;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders>}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.getClaimsToHeadersMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders>} */ (
      jspb.Message.getMapField(this, 2, opt_noLazyCreate,
      proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders));
};


proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.clearClaimsToHeadersMap = function() {
  this.getClaimsToHeadersMap().clear();
};


/**
 * optional bool clear_route_cache = 3;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.getClearRouteCache = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 3, false));
};


/** @param {boolean} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.setClearRouteCache = function(value) {
  jspb.Message.setProto3BooleanField(this, 3, value);
};


/**
 * optional string payload_in_metadata = 4;
 * @return {string}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.getPayloadInMetadata = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.prototype.setPayloadInMetadata = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.displayName = 'proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute';
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.toObject = function(includeInstance, msg) {
  var f, obj = {
    jwtConfigsMap: (f = msg.getJwtConfigsMap()) ? f.toObject(includeInstance, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.toObject) : []
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
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute;
  return proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 5:
      var value = msg.getJwtConfigsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readUint32, jspb.BinaryReader.prototype.readMessage, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.deserializeBinaryFromReader, 0);
         });
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
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getJwtConfigsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeUint32, jspb.BinaryWriter.prototype.writeMessage, proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.serializeBinaryToWriter);
  }
};


/**
 * map<uint32, SoloJwtAuthnPerRoute> jwt_configs = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<number,!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute>}
 */
proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.prototype.getJwtConfigsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<number,!proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      proto.envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute));
};


proto.envoy.config.filter.http.solo_jwt_authn.v2.StagedJwtAuthnPerRoute.prototype.clearJwtConfigsMap = function() {
  this.getJwtConfigsMap().clear();
};


goog.object.extend(exports, proto.envoy.config.filter.http.solo_jwt_authn.v2);
