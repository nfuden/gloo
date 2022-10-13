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
var github_com_solo$io_solo$projects_projects_gloo$fed_api_multicluster_v1alpha1_multicluster_pb = require('../../../../../../../../github.com/solo-io/solo-projects/projects/gloo-fed/api/multicluster/v1alpha1/multicluster_pb.js');
var github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb = require('../../../../../../../../github.com/solo-io/solo-projects/projects/gloo-fed/api/fed/core/v1/placement_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_settings_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/settings_pb.js');
goog.exportSymbol('proto.fed.gloo.solo.io.FederatedSettingsSpec', null, global);
goog.exportSymbol('proto.fed.gloo.solo.io.FederatedSettingsSpec.Template', null, global);
goog.exportSymbol('proto.fed.gloo.solo.io.FederatedSettingsStatus', null, global);

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
proto.fed.gloo.solo.io.FederatedSettingsSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.gloo.solo.io.FederatedSettingsSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.gloo.solo.io.FederatedSettingsSpec.displayName = 'proto.fed.gloo.solo.io.FederatedSettingsSpec';
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
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.gloo.solo.io.FederatedSettingsSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    template: (f = msg.getTemplate()) && proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.toObject(includeInstance, f),
    placement: (f = msg.getPlacement()) && github_com_solo$io_solo$projects_projects_gloo$fed_api_multicluster_v1alpha1_multicluster_pb.Placement.toObject(includeInstance, f)
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
 * @return {!proto.fed.gloo.solo.io.FederatedSettingsSpec}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.gloo.solo.io.FederatedSettingsSpec;
  return proto.fed.gloo.solo.io.FederatedSettingsSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.gloo.solo.io.FederatedSettingsSpec}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.fed.gloo.solo.io.FederatedSettingsSpec.Template;
      reader.readMessage(value,proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.deserializeBinaryFromReader);
      msg.setTemplate(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$projects_projects_gloo$fed_api_multicluster_v1alpha1_multicluster_pb.Placement;
      reader.readMessage(value,github_com_solo$io_solo$projects_projects_gloo$fed_api_multicluster_v1alpha1_multicluster_pb.Placement.deserializeBinaryFromReader);
      msg.setPlacement(value);
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
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.gloo.solo.io.FederatedSettingsSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTemplate();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.serializeBinaryToWriter
    );
  }
  f = message.getPlacement();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$projects_projects_gloo$fed_api_multicluster_v1alpha1_multicluster_pb.Placement.serializeBinaryToWriter
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
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.gloo.solo.io.FederatedSettingsSpec.Template, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.displayName = 'proto.fed.gloo.solo.io.FederatedSettingsSpec.Template';
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
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsSpec.Template} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.toObject = function(includeInstance, msg) {
  var f, obj = {
    spec: (f = msg.getSpec()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_settings_pb.SettingsSpec.toObject(includeInstance, f),
    metadata: (f = msg.getMetadata()) && github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.TemplateMetadata.toObject(includeInstance, f)
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
 * @return {!proto.fed.gloo.solo.io.FederatedSettingsSpec.Template}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.gloo.solo.io.FederatedSettingsSpec.Template;
  return proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsSpec.Template} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.gloo.solo.io.FederatedSettingsSpec.Template}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_settings_pb.SettingsSpec;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_settings_pb.SettingsSpec.deserializeBinaryFromReader);
      msg.setSpec(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.TemplateMetadata;
      reader.readMessage(value,github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.TemplateMetadata.deserializeBinaryFromReader);
      msg.setMetadata(value);
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
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsSpec.Template} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSpec();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_settings_pb.SettingsSpec.serializeBinaryToWriter
    );
  }
  f = message.getMetadata();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.TemplateMetadata.serializeBinaryToWriter
    );
  }
};


/**
 * optional gloo.solo.io.SettingsSpec spec = 1;
 * @return {?proto.gloo.solo.io.SettingsSpec}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.getSpec = function() {
  return /** @type{?proto.gloo.solo.io.SettingsSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_settings_pb.SettingsSpec, 1));
};


/** @param {?proto.gloo.solo.io.SettingsSpec|undefined} value */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.setSpec = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.clearSpec = function() {
  this.setSpec(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.hasSpec = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional core.fed.solo.io.TemplateMetadata metadata = 2;
 * @return {?proto.core.fed.solo.io.TemplateMetadata}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.getMetadata = function() {
  return /** @type{?proto.core.fed.solo.io.TemplateMetadata} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.TemplateMetadata, 2));
};


/** @param {?proto.core.fed.solo.io.TemplateMetadata|undefined} value */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.setMetadata = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.clearMetadata = function() {
  this.setMetadata(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.Template.prototype.hasMetadata = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Template template = 1;
 * @return {?proto.fed.gloo.solo.io.FederatedSettingsSpec.Template}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.getTemplate = function() {
  return /** @type{?proto.fed.gloo.solo.io.FederatedSettingsSpec.Template} */ (
    jspb.Message.getWrapperField(this, proto.fed.gloo.solo.io.FederatedSettingsSpec.Template, 1));
};


/** @param {?proto.fed.gloo.solo.io.FederatedSettingsSpec.Template|undefined} value */
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.setTemplate = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.clearTemplate = function() {
  this.setTemplate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.hasTemplate = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional multicluster.solo.io.Placement placement = 2;
 * @return {?proto.multicluster.solo.io.Placement}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.getPlacement = function() {
  return /** @type{?proto.multicluster.solo.io.Placement} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_multicluster_v1alpha1_multicluster_pb.Placement, 2));
};


/** @param {?proto.multicluster.solo.io.Placement|undefined} value */
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.setPlacement = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.clearPlacement = function() {
  this.setPlacement(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.gloo.solo.io.FederatedSettingsSpec.prototype.hasPlacement = function() {
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
proto.fed.gloo.solo.io.FederatedSettingsStatus = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.gloo.solo.io.FederatedSettingsStatus, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.gloo.solo.io.FederatedSettingsStatus.displayName = 'proto.fed.gloo.solo.io.FederatedSettingsStatus';
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
proto.fed.gloo.solo.io.FederatedSettingsStatus.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.gloo.solo.io.FederatedSettingsStatus.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsStatus} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.gloo.solo.io.FederatedSettingsStatus.toObject = function(includeInstance, msg) {
  var f, obj = {
    placementStatus: (f = msg.getPlacementStatus()) && github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus.toObject(includeInstance, f)
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
 * @return {!proto.fed.gloo.solo.io.FederatedSettingsStatus}
 */
proto.fed.gloo.solo.io.FederatedSettingsStatus.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.gloo.solo.io.FederatedSettingsStatus;
  return proto.fed.gloo.solo.io.FederatedSettingsStatus.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsStatus} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.gloo.solo.io.FederatedSettingsStatus}
 */
proto.fed.gloo.solo.io.FederatedSettingsStatus.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus;
      reader.readMessage(value,github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus.deserializeBinaryFromReader);
      msg.setPlacementStatus(value);
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
proto.fed.gloo.solo.io.FederatedSettingsStatus.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.gloo.solo.io.FederatedSettingsStatus.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.gloo.solo.io.FederatedSettingsStatus} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.gloo.solo.io.FederatedSettingsStatus.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPlacementStatus();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus.serializeBinaryToWriter
    );
  }
};


/**
 * optional core.fed.solo.io.PlacementStatus placement_status = 1;
 * @return {?proto.core.fed.solo.io.PlacementStatus}
 */
proto.fed.gloo.solo.io.FederatedSettingsStatus.prototype.getPlacementStatus = function() {
  return /** @type{?proto.core.fed.solo.io.PlacementStatus} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus, 1));
};


/** @param {?proto.core.fed.solo.io.PlacementStatus|undefined} value */
proto.fed.gloo.solo.io.FederatedSettingsStatus.prototype.setPlacementStatus = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.gloo.solo.io.FederatedSettingsStatus.prototype.clearPlacementStatus = function() {
  this.setPlacementStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.gloo.solo.io.FederatedSettingsStatus.prototype.hasPlacementStatus = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.fed.gloo.solo.io);
