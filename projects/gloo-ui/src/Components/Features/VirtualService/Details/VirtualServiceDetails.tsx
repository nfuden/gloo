import styled from '@emotion/styled';
import { Spin } from 'antd';
import { ReactComponent as GlooIcon } from 'assets/GlooEE.svg';
import { Breadcrumb } from 'Components/Common/Breadcrumb';
import { ConfigDisplayer } from 'Components/Common/DisplayOnly/ConfigDisplayer';
import { FileDownloadLink } from 'Components/Common/FileDownloadLink';
import { SectionCard } from 'Components/Common/SectionCard';
import { Route } from 'proto/github.com/solo-io/gloo/projects/gateway/api/v1/virtual_service_pb';
import { OAuth } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/extauth/v1/extauth_pb';
import { IngressRateLimit } from 'proto/github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/ratelimit/ratelimit_pb';
import { EditedResourceYaml } from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/types_pb';
import * as React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RouteComponentProps } from 'react-router';
import { AppState } from 'store';
import {
  updateVirtualService,
  updateVirtualServiceYaml
} from 'store/virtualServices/actions';
import { colors, healthConstants, soloConstants } from 'Styles';
import { Domains } from './Domains';
import { Routes } from './Routes';
import { ExtAuth } from './ExtAuth';
import { RateLimit } from './RateLimit';

const ConfigContainer = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  background: ${colors.januaryGrey};
  height: 80%;
  border-radius: ${soloConstants.smallRadius}px;
`;

const ConfigItem = styled.div`
  margin: 20px;
  padding: 10px;
  justify-items: center;
  background: white;
`;

type DetailsContentProps = { configurationShowing?: boolean };
const DetailsContent = styled.div`
  position: relative;
  display: grid;
  grid-template-rows: ${(props: DetailsContentProps) =>
      props.configurationShowing ? 'auto' : ''} auto 1fr 1fr;
  grid-template-columns: 100%;
  grid-column-gap: 30px;
`;

const YamlLink = styled.div`
  position: absolute;
  top: 10px;
  right: 0;
  display: flex;
`;
const ConfigurationToggle = styled.div`
  cursor: pointer;
  color: ${colors.seaBlue};
  font-size: 14px;
  margin-right: 8px;
`;

const DetailsSection = styled.div`
  width: 100%;
`;
export const DetailsSectionTitle = styled.div`
  font-size: 18px;
  font-weight: bold;
  color: ${colors.novemberGrey};
  margin-top: 10px;
  margin-bottom: 10px;
`;

interface Props
  extends RouteComponentProps<{
    virtualservicename: string;
    virtualservicenamespace: string;
  }> {}

export const VirtualServiceDetails = (props: Props) => {
  const { match, history, location } = props;
  let { virtualservicename, virtualservicenamespace } = match.params;

  const [showConfiguration, setShowConfiguration] = React.useState(false);

  const virtualServicesList = useSelector(
    (state: AppState) => state.virtualServices.virtualServicesList
  );
  const yamlError = useSelector(
    (state: AppState) => state.virtualServices.yamlParseError
  );

  const dispatch = useDispatch();

  let virtualServiceDetails = virtualServicesList.find(
    vsD => vsD && vsD.virtualService!.metadata!.name === virtualservicename
  )!;
  if (
    !virtualServicesList.length ||
    !virtualServiceDetails ||
    !virtualServiceDetails.virtualService ||
    !virtualServiceDetails.virtualService.virtualHost
  ) {
    return (
      <React.Fragment>
        <Breadcrumb />

        <Spin size='large' />
      </React.Fragment>
    );
  }

  let { virtualService, raw, plugins } = virtualServiceDetails!;

  let {
    routesList,
    virtualHostPlugins,
    domainsList
  } = virtualServiceDetails.virtualService!.virtualHost!;

  let rateLimits;
  let externalAuth;

  if (plugins) {
    rateLimits = plugins.rateLimit;
    externalAuth = plugins.extAuth;
  }

  const handleUpdateVirtualService = (newInfo: {
    newDomainsList?: string[];
    newRoutesList?: Route.AsObject[];
    newRateLimits?: IngressRateLimit.AsObject;
    newOAuth?: OAuth.AsObject;
  }) => {
    dispatch(
      updateVirtualService({
        inputV2: {
          domains: { valuesList: newInfo.newDomainsList as string[] },
          rateLimitConfig: { value: newInfo.newRateLimits },
          ref: { name: virtualservicename, namespace: virtualservicenamespace },
          routes: {
            valuesList: newInfo.newRoutesList!
          }
        }
      })
    );
  };

  const saveYamlChange = (newYaml: string) => {
    dispatch(
      updateVirtualServiceYaml({
        editedYamlData: {
          editedYaml: newYaml,
          ref: {
            name: virtualService!.metadata!.name,
            namespace: virtualService!.metadata!.namespace
          }
        }
      })
    );
  };

  const headerInfo = [
    {
      title: 'namespace',
      value: virtualservicenamespace
    }
  ];
  return (
    <React.Fragment>
      <Breadcrumb />

      <SectionCard
        cardName={
          virtualService!.displayName.length
            ? virtualService!.displayName
            : match.params
            ? match.params.virtualservicename
            : 'Error'
        }
        logoIcon={<GlooIcon />}
        health={
          virtualService!.status
            ? virtualService!.status!.state
            : healthConstants.Pending.value
        }
        headerSecondaryInformation={headerInfo}
        healthMessage={
          virtualService!.status && virtualService!.status!.reason.length
            ? virtualService!.status!.reason
            : 'Service Status'
        }
        onClose={() => history.push(`/virtualservices/`)}>
        <DetailsContent configurationShowing={showConfiguration}>
          {!!raw && (
            <YamlLink>
              <ConfigurationToggle
                onClick={() => setShowConfiguration(s => !s)}>
                {showConfiguration ? 'Hide' : 'View'} Raw Configuration
              </ConfigurationToggle>
              <FileDownloadLink
                fileContent={raw.content}
                fileName={raw.fileName}
              />
            </YamlLink>
          )}
          {showConfiguration && (
            <DetailsSection>
              <DetailsSectionTitle>Raw Configuration</DetailsSectionTitle>
              <ConfigDisplayer
                content={raw ? raw.content : ''}
                asEditor
                yamlError={yamlError}
                saveEdits={saveYamlChange}
              />
            </DetailsSection>
          )}
          <DetailsSection>
            <Domains
              domains={domainsList}
              vsRef={{
                name: virtualservicename,
                namespace: virtualservicenamespace
              }}
            />
          </DetailsSection>
          <DetailsSection>
            <Routes routes={routesList} virtualService={virtualService!} />
          </DetailsSection>
          <DetailsSection>
            <React.Fragment>
              <DetailsSectionTitle>Configuration</DetailsSectionTitle>
              <ConfigContainer>
                <ConfigItem>
                  <ExtAuth externalAuth={externalAuth} />
                </ConfigItem>
                <ConfigItem>
                  <RateLimit rateLimits={rateLimits} />
                </ConfigItem>
              </ConfigContainer>
            </React.Fragment>
          </DetailsSection>
        </DetailsContent>
      </SectionCard>
    </React.Fragment>
  );
};
