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
var github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/enterprise.gloo/v1/auth_config_pb.js');
goog.exportSymbol('proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec', null, global);
goog.exportSymbol('proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template', null, global);
goog.exportSymbol('proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus', null, global);

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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.displayName = 'proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec';
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.toObject = function(includeInstance, msg) {
  var f, obj = {
    template: (f = msg.getTemplate()) && proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.toObject(includeInstance, f),
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
 * @return {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec;
  return proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template;
      reader.readMessage(value,proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.deserializeBinaryFromReader);
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTemplate();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.serializeBinaryToWriter
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.displayName = 'proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template';
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.toObject = function(includeInstance, msg) {
  var f, obj = {
    spec: (f = msg.getSpec()) && github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.AuthConfigSpec.toObject(includeInstance, f),
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
 * @return {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template;
  return proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.AuthConfigSpec;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.AuthConfigSpec.deserializeBinaryFromReader);
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSpec();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.AuthConfigSpec.serializeBinaryToWriter
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
 * optional enterprise.gloo.solo.io.AuthConfigSpec spec = 1;
 * @return {?proto.enterprise.gloo.solo.io.AuthConfigSpec}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.getSpec = function() {
  return /** @type{?proto.enterprise.gloo.solo.io.AuthConfigSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_enterprise_gloo_v1_auth_config_pb.AuthConfigSpec, 1));
};


/** @param {?proto.enterprise.gloo.solo.io.AuthConfigSpec|undefined} value */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.setSpec = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.clearSpec = function() {
  this.setSpec(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.hasSpec = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional core.fed.solo.io.TemplateMetadata metadata = 2;
 * @return {?proto.core.fed.solo.io.TemplateMetadata}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.getMetadata = function() {
  return /** @type{?proto.core.fed.solo.io.TemplateMetadata} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.TemplateMetadata, 2));
};


/** @param {?proto.core.fed.solo.io.TemplateMetadata|undefined} value */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.setMetadata = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.clearMetadata = function() {
  this.setMetadata(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template.prototype.hasMetadata = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Template template = 1;
 * @return {?proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.getTemplate = function() {
  return /** @type{?proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template} */ (
    jspb.Message.getWrapperField(this, proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template, 1));
};


/** @param {?proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.Template|undefined} value */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.setTemplate = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.clearTemplate = function() {
  this.setTemplate(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.hasTemplate = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional multicluster.solo.io.Placement placement = 2;
 * @return {?proto.multicluster.solo.io.Placement}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.getPlacement = function() {
  return /** @type{?proto.multicluster.solo.io.Placement} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_multicluster_v1alpha1_multicluster_pb.Placement, 2));
};


/** @param {?proto.multicluster.solo.io.Placement|undefined} value */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.setPlacement = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.clearPlacement = function() {
  this.setPlacement(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigSpec.prototype.hasPlacement = function() {
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.displayName = 'proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus';
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.toObject = function(includeInstance, msg) {
  var f, obj = {
    placementStatus: (f = msg.getPlacementStatus()) && github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus.toObject(includeInstance, f),
    namespacedPlacementStatusesMap: (f = msg.getNamespacedPlacementStatusesMap()) ? f.toObject(includeInstance, proto.core.fed.solo.io.PlacementStatus.toObject) : []
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
 * @return {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus;
  return proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.deserializeBinaryFromReader = function(msg, reader) {
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
    case 2:
      var value = msg.getNamespacedPlacementStatusesMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.core.fed.solo.io.PlacementStatus.deserializeBinaryFromReader, "");
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
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPlacementStatus();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus.serializeBinaryToWriter
    );
  }
  f = message.getNamespacedPlacementStatusesMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(2, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.core.fed.solo.io.PlacementStatus.serializeBinaryToWriter);
  }
};


/**
 * optional core.fed.solo.io.PlacementStatus placement_status = 1;
 * @return {?proto.core.fed.solo.io.PlacementStatus}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.getPlacementStatus = function() {
  return /** @type{?proto.core.fed.solo.io.PlacementStatus} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_core_v1_placement_pb.PlacementStatus, 1));
};


/** @param {?proto.core.fed.solo.io.PlacementStatus|undefined} value */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.setPlacementStatus = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.clearPlacementStatus = function() {
  this.setPlacementStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.hasPlacementStatus = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * map<string, core.fed.solo.io.PlacementStatus> namespaced_placement_statuses = 2;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.core.fed.solo.io.PlacementStatus>}
 */
proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.getNamespacedPlacementStatusesMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.core.fed.solo.io.PlacementStatus>} */ (
      jspb.Message.getMapField(this, 2, opt_noLazyCreate,
      proto.core.fed.solo.io.PlacementStatus));
};


proto.fed.enterprise.gloo.solo.io.FederatedAuthConfigStatus.prototype.clearNamespacedPlacementStatusesMap = function() {
  this.getNamespacedPlacementStatusesMap().clear();
};


goog.object.extend(exports, proto.fed.enterprise.gloo.solo.io);