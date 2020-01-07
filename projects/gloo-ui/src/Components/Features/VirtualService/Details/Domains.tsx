import styled from '@emotion/styled';
import { StringCardsList } from 'Components/Common/StringCardsList';
import { isEqual } from 'lodash';
import * as React from 'react';
import { colors, soloConstants } from 'Styles';
import { DetailsSectionTitle } from './VirtualServiceDetails';
import { useDispatch } from 'react-redux';
import { updateDomains } from 'store/virtualServices/actions';
import { ResourceRef } from 'proto/solo-kit/api/v1/ref_pb';

const DomainsContainer = styled.div`
  background: ${colors.januaryGrey};
  display: flex;
  justify-content: flex-start;
  align-content: center;
  align-items: center;
  padding: 7px 7px 3px;
  border-radius: ${soloConstants.smallRadius}px;
`;

interface Props {
  domains: string[];
  vsRef: ResourceRef.AsObject;
}

function equivalentProps(
  oldProps: Readonly<Props>,
  nextProps: Readonly<Props>
): boolean {
  return isEqual(oldProps.domains, nextProps.domains);
}

export const Domains: React.FC<Props> = React.memo(props => {
  const dispatch = useDispatch();
  const [domains, setDomains] = React.useState(props.domains);

  React.useEffect(() => {
    if (!isEqual(props.domains, domains)) {
      setDomains(props.domains);
    }
  }, [props.domains]);

  const addDomain = (domain: string) => {
    setDomains([...domains, domain]);
    dispatch(
      updateDomains({ ref: props.vsRef, domains: [...domains, domain] })
    );
  };

  const removeDomain = (removeIndex: number) => {
    let newList = [...domains];
    newList.splice(removeIndex, 1);

    dispatch(updateDomains({ ref: props.vsRef, domains: newList }));
    setDomains(newList);
  };

  const isValidDomain = (domain: string) => {
    if (domain === '*') {
      return true;
    }

    return RegExp(
      '(?:[a-z0-9*](?:[a-z0-9-]{0,61}[a-z0-9])?.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]'
    ).test(domain);
  };

  return (
    <>
      <DetailsSectionTitle>Domains</DetailsSectionTitle>
      <DomainsContainer>
        <StringCardsList
          values={domains}
          valueDeleted={removeDomain}
          createNewPromptText='new.domain.com'
          createNew={addDomain}
          valueIsValid={isValidDomain}
        />
      </DomainsContainer>
    </>
  );
}, equivalentProps);
