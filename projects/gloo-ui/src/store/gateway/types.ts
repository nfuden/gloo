import { GatewayDetails } from 'proto/solo-projects/projects/grpcserver/api/v1/gateway_pb';

export enum GatewayAction {
  GET_GATEWAY = 'GET_GATEWAY',
  LIST_GATEWAYS = 'LIST_GATEWAYS',
  UPDATE_GATEWAY = 'UPDATE_GATEWAY',
  UPDATE_GATEWAY_YAML = 'UPDATE_GATEWAY_YAML',
  UPDATE_GATEWAY_YAML_ERROR = 'UPDATE_GATEWAY_YAML_ERROR'
}

export interface GetGatewayAction {
  type: typeof GatewayAction.GET_GATEWAY;
  payload: GatewayDetails.AsObject;
}
export interface ListGatewaysAction {
  type: typeof GatewayAction.LIST_GATEWAYS;
  payload: GatewayDetails.AsObject[];
}

export interface UpdateGatewayAction {
  type: typeof GatewayAction.UPDATE_GATEWAY;
  payload: GatewayDetails.AsObject;
}

export interface UpdateGatewayYamlAction {
  type: typeof GatewayAction.UPDATE_GATEWAY_YAML;
  payload: GatewayDetails.AsObject;
}
export interface UpdateGatewayYamlErrorAction {
  type: typeof GatewayAction.UPDATE_GATEWAY_YAML_ERROR;
  payload: Error;
}

export type GatewayActionTypes =
  | GetGatewayAction
  | ListGatewaysAction
  | UpdateGatewayAction
  | UpdateGatewayYamlAction
  | UpdateGatewayYamlErrorAction;
