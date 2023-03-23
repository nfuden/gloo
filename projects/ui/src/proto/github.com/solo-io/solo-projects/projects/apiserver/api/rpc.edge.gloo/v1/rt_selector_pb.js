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
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/core/matchers/matchers_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/proxy_pb.js');
var github_com_solo$io_solo$apis_api_gloo_gateway_v1_virtual_service_pb = require('../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gateway/v1/virtual_service_pb.js');
var github_com_solo$io_skv2_api_core_v1_core_pb = require('../../../../../../../../github.com/solo-io/skv2/api/core/v1/core_pb.js');
goog.exportSymbol('proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest', null, global);
goog.exportSymbol('proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse', null, global);
goog.exportSymbol('proto.rpc.edge.gloo.solo.io.SubRouteTableRow', null, global);

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
proto.rpc.edge.gloo.solo.io.SubRouteTableRow = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.rpc.edge.gloo.solo.io.SubRouteTableRow.repeatedFields_, proto.rpc.edge.gloo.solo.io.SubRouteTableRow.oneofGroups_);
};
goog.inherits(proto.rpc.edge.gloo.solo.io.SubRouteTableRow, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.rpc.edge.gloo.solo.io.SubRouteTableRow.displayName = 'proto.rpc.edge.gloo.solo.io.SubRouteTableRow';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.repeatedFields_ = [7,8,9,10];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.oneofGroups_ = [[1,2,3,4]];

/**
 * @enum {number}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.ActionCase = {
  ACTION_NOT_SET: 0,
  ROUTE_ACTION: 1,
  REDIRECT_ACTION: 2,
  DIRECT_RESPONSE_ACTION: 3,
  DELEGATE_ACTION: 4
};

/**
 * @return {proto.rpc.edge.gloo.solo.io.SubRouteTableRow.ActionCase}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getActionCase = function() {
  return /** @type {proto.rpc.edge.gloo.solo.io.SubRouteTableRow.ActionCase} */(jspb.Message.computeOneofCase(this, proto.rpc.edge.gloo.solo.io.SubRouteTableRow.oneofGroups_[0]));
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
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.toObject = function(opt_includeInstance) {
  return proto.rpc.edge.gloo.solo.io.SubRouteTableRow.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.toObject = function(includeInstance, msg) {
  var f, obj = {
    routeAction: (f = msg.getRouteAction()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RouteAction.toObject(includeInstance, f),
    redirectAction: (f = msg.getRedirectAction()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RedirectAction.toObject(includeInstance, f),
    directResponseAction: (f = msg.getDirectResponseAction()) && github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.DirectResponseAction.toObject(includeInstance, f),
    delegateAction: (f = msg.getDelegateAction()) && github_com_solo$io_solo$apis_api_gloo_gateway_v1_virtual_service_pb.DelegateAction.toObject(includeInstance, f),
    matcher: jspb.Message.getFieldWithDefault(msg, 5, ""),
    matchType: jspb.Message.getFieldWithDefault(msg, 6, ""),
    methodsList: jspb.Message.getRepeatedField(msg, 7),
    headersList: jspb.Message.toObjectList(msg.getHeadersList(),
    github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.HeaderMatcher.toObject, includeInstance),
    queryParametersList: jspb.Message.toObjectList(msg.getQueryParametersList(),
    github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.QueryParameterMatcher.toObject, includeInstance),
    rtRoutesList: jspb.Message.toObjectList(msg.getRtRoutesList(),
    proto.rpc.edge.gloo.solo.io.SubRouteTableRow.toObject, includeInstance)
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
 * @return {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.rpc.edge.gloo.solo.io.SubRouteTableRow;
  return proto.rpc.edge.gloo.solo.io.SubRouteTableRow.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RouteAction;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RouteAction.deserializeBinaryFromReader);
      msg.setRouteAction(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RedirectAction;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RedirectAction.deserializeBinaryFromReader);
      msg.setRedirectAction(value);
      break;
    case 3:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.DirectResponseAction;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.DirectResponseAction.deserializeBinaryFromReader);
      msg.setDirectResponseAction(value);
      break;
    case 4:
      var value = new github_com_solo$io_solo$apis_api_gloo_gateway_v1_virtual_service_pb.DelegateAction;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gateway_v1_virtual_service_pb.DelegateAction.deserializeBinaryFromReader);
      msg.setDelegateAction(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setMatcher(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setMatchType(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.addMethods(value);
      break;
    case 8:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.HeaderMatcher;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.HeaderMatcher.deserializeBinaryFromReader);
      msg.addHeaders(value);
      break;
    case 9:
      var value = new github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.QueryParameterMatcher;
      reader.readMessage(value,github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.QueryParameterMatcher.deserializeBinaryFromReader);
      msg.addQueryParameters(value);
      break;
    case 10:
      var value = new proto.rpc.edge.gloo.solo.io.SubRouteTableRow;
      reader.readMessage(value,proto.rpc.edge.gloo.solo.io.SubRouteTableRow.deserializeBinaryFromReader);
      msg.addRtRoutes(value);
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
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.rpc.edge.gloo.solo.io.SubRouteTableRow.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRouteAction();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RouteAction.serializeBinaryToWriter
    );
  }
  f = message.getRedirectAction();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RedirectAction.serializeBinaryToWriter
    );
  }
  f = message.getDirectResponseAction();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.DirectResponseAction.serializeBinaryToWriter
    );
  }
  f = message.getDelegateAction();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      github_com_solo$io_solo$apis_api_gloo_gateway_v1_virtual_service_pb.DelegateAction.serializeBinaryToWriter
    );
  }
  f = message.getMatcher();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getMatchType();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getMethodsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      7,
      f
    );
  }
  f = message.getHeadersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      8,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.HeaderMatcher.serializeBinaryToWriter
    );
  }
  f = message.getQueryParametersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      9,
      f,
      github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.QueryParameterMatcher.serializeBinaryToWriter
    );
  }
  f = message.getRtRoutesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      10,
      f,
      proto.rpc.edge.gloo.solo.io.SubRouteTableRow.serializeBinaryToWriter
    );
  }
};


/**
 * optional gloo.solo.io.RouteAction route_action = 1;
 * @return {?proto.gloo.solo.io.RouteAction}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getRouteAction = function() {
  return /** @type{?proto.gloo.solo.io.RouteAction} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RouteAction, 1));
};


/** @param {?proto.gloo.solo.io.RouteAction|undefined} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setRouteAction = function(value) {
  jspb.Message.setOneofWrapperField(this, 1, proto.rpc.edge.gloo.solo.io.SubRouteTableRow.oneofGroups_[0], value);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearRouteAction = function() {
  this.setRouteAction(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.hasRouteAction = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional gloo.solo.io.RedirectAction redirect_action = 2;
 * @return {?proto.gloo.solo.io.RedirectAction}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getRedirectAction = function() {
  return /** @type{?proto.gloo.solo.io.RedirectAction} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.RedirectAction, 2));
};


/** @param {?proto.gloo.solo.io.RedirectAction|undefined} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setRedirectAction = function(value) {
  jspb.Message.setOneofWrapperField(this, 2, proto.rpc.edge.gloo.solo.io.SubRouteTableRow.oneofGroups_[0], value);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearRedirectAction = function() {
  this.setRedirectAction(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.hasRedirectAction = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional gloo.solo.io.DirectResponseAction direct_response_action = 3;
 * @return {?proto.gloo.solo.io.DirectResponseAction}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getDirectResponseAction = function() {
  return /** @type{?proto.gloo.solo.io.DirectResponseAction} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_proxy_pb.DirectResponseAction, 3));
};


/** @param {?proto.gloo.solo.io.DirectResponseAction|undefined} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setDirectResponseAction = function(value) {
  jspb.Message.setOneofWrapperField(this, 3, proto.rpc.edge.gloo.solo.io.SubRouteTableRow.oneofGroups_[0], value);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearDirectResponseAction = function() {
  this.setDirectResponseAction(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.hasDirectResponseAction = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional gateway.solo.io.DelegateAction delegate_action = 4;
 * @return {?proto.gateway.solo.io.DelegateAction}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getDelegateAction = function() {
  return /** @type{?proto.gateway.solo.io.DelegateAction} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gateway_v1_virtual_service_pb.DelegateAction, 4));
};


/** @param {?proto.gateway.solo.io.DelegateAction|undefined} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setDelegateAction = function(value) {
  jspb.Message.setOneofWrapperField(this, 4, proto.rpc.edge.gloo.solo.io.SubRouteTableRow.oneofGroups_[0], value);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearDelegateAction = function() {
  this.setDelegateAction(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.hasDelegateAction = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional string matcher = 5;
 * @return {string}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getMatcher = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setMatcher = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string match_type = 6;
 * @return {string}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getMatchType = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setMatchType = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * repeated string methods = 7;
 * @return {!Array<string>}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getMethodsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 7));
};


/** @param {!Array<string>} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setMethodsList = function(value) {
  jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {!string} value
 * @param {number=} opt_index
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.addMethods = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearMethodsList = function() {
  this.setMethodsList([]);
};


/**
 * repeated matchers.core.gloo.solo.io.HeaderMatcher headers = 8;
 * @return {!Array<!proto.matchers.core.gloo.solo.io.HeaderMatcher>}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getHeadersList = function() {
  return /** @type{!Array<!proto.matchers.core.gloo.solo.io.HeaderMatcher>} */ (
    jspb.Message.getRepeatedWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.HeaderMatcher, 8));
};


/** @param {!Array<!proto.matchers.core.gloo.solo.io.HeaderMatcher>} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setHeadersList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 8, value);
};


/**
 * @param {!proto.matchers.core.gloo.solo.io.HeaderMatcher=} opt_value
 * @param {number=} opt_index
 * @return {!proto.matchers.core.gloo.solo.io.HeaderMatcher}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.addHeaders = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 8, opt_value, proto.matchers.core.gloo.solo.io.HeaderMatcher, opt_index);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearHeadersList = function() {
  this.setHeadersList([]);
};


/**
 * repeated matchers.core.gloo.solo.io.QueryParameterMatcher query_parameters = 9;
 * @return {!Array<!proto.matchers.core.gloo.solo.io.QueryParameterMatcher>}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getQueryParametersList = function() {
  return /** @type{!Array<!proto.matchers.core.gloo.solo.io.QueryParameterMatcher>} */ (
    jspb.Message.getRepeatedWrapperField(this, github_com_solo$io_solo$apis_api_gloo_gloo_v1_core_matchers_matchers_pb.QueryParameterMatcher, 9));
};


/** @param {!Array<!proto.matchers.core.gloo.solo.io.QueryParameterMatcher>} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setQueryParametersList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 9, value);
};


/**
 * @param {!proto.matchers.core.gloo.solo.io.QueryParameterMatcher=} opt_value
 * @param {number=} opt_index
 * @return {!proto.matchers.core.gloo.solo.io.QueryParameterMatcher}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.addQueryParameters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 9, opt_value, proto.matchers.core.gloo.solo.io.QueryParameterMatcher, opt_index);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearQueryParametersList = function() {
  this.setQueryParametersList([]);
};


/**
 * repeated SubRouteTableRow rt_routes = 10;
 * @return {!Array<!proto.rpc.edge.gloo.solo.io.SubRouteTableRow>}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.getRtRoutesList = function() {
  return /** @type{!Array<!proto.rpc.edge.gloo.solo.io.SubRouteTableRow>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.rpc.edge.gloo.solo.io.SubRouteTableRow, 10));
};


/** @param {!Array<!proto.rpc.edge.gloo.solo.io.SubRouteTableRow>} value */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.setRtRoutesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 10, value);
};


/**
 * @param {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow=} opt_value
 * @param {number=} opt_index
 * @return {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow}
 */
proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.addRtRoutes = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 10, opt_value, proto.rpc.edge.gloo.solo.io.SubRouteTableRow, opt_index);
};


proto.rpc.edge.gloo.solo.io.SubRouteTableRow.prototype.clearRtRoutesList = function() {
  this.setRtRoutesList([]);
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
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.displayName = 'proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest';
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
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    virtualServiceRef: (f = msg.getVirtualServiceRef()) && github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef.toObject(includeInstance, f)
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
 * @return {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest;
  return proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef;
      reader.readMessage(value,github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef.deserializeBinaryFromReader);
      msg.setVirtualServiceRef(value);
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
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVirtualServiceRef();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef.serializeBinaryToWriter
    );
  }
};


/**
 * optional core.skv2.solo.io.ClusterObjectRef virtual_service_ref = 1;
 * @return {?proto.core.skv2.solo.io.ClusterObjectRef}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.prototype.getVirtualServiceRef = function() {
  return /** @type{?proto.core.skv2.solo.io.ClusterObjectRef} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef, 1));
};


/** @param {?proto.core.skv2.solo.io.ClusterObjectRef|undefined} value */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.prototype.setVirtualServiceRef = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.prototype.clearVirtualServiceRef = function() {
  this.setVirtualServiceRef(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesRequest.prototype.hasVirtualServiceRef = function() {
  return jspb.Message.getField(this, 1) != null;
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
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.repeatedFields_, null);
};
goog.inherits(proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.displayName = 'proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.repeatedFields_ = [1];



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
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    vsRoutesList: jspb.Message.toObjectList(msg.getVsRoutesList(),
    proto.rpc.edge.gloo.solo.io.SubRouteTableRow.toObject, includeInstance)
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
 * @return {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse;
  return proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.rpc.edge.gloo.solo.io.SubRouteTableRow;
      reader.readMessage(value,proto.rpc.edge.gloo.solo.io.SubRouteTableRow.deserializeBinaryFromReader);
      msg.addVsRoutes(value);
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
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVsRoutesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.rpc.edge.gloo.solo.io.SubRouteTableRow.serializeBinaryToWriter
    );
  }
};


/**
 * repeated SubRouteTableRow vs_routes = 1;
 * @return {!Array<!proto.rpc.edge.gloo.solo.io.SubRouteTableRow>}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.prototype.getVsRoutesList = function() {
  return /** @type{!Array<!proto.rpc.edge.gloo.solo.io.SubRouteTableRow>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.rpc.edge.gloo.solo.io.SubRouteTableRow, 1));
};


/** @param {!Array<!proto.rpc.edge.gloo.solo.io.SubRouteTableRow>} value */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.prototype.setVsRoutesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow=} opt_value
 * @param {number=} opt_index
 * @return {!proto.rpc.edge.gloo.solo.io.SubRouteTableRow}
 */
proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.prototype.addVsRoutes = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.rpc.edge.gloo.solo.io.SubRouteTableRow, opt_index);
};


proto.rpc.edge.gloo.solo.io.GetVirtualServiceRoutesResponse.prototype.clearVsRoutesList = function() {
  this.setVsRoutesList([]);
};


goog.object.extend(exports, proto.rpc.edge.gloo.solo.io);