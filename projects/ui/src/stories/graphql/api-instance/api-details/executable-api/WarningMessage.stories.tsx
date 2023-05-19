import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import WarningMessage from '../../../../../Components/Features/Graphql/api-instance/api-details/executable-api/WarningMessage';

export default {
  title:
    'Graphql / api-instance / api-details / executable-api / WarningMessage',
  component: WarningMessage,
} as ComponentMeta<typeof WarningMessage>;

const Template: ComponentStory<typeof WarningMessage> = args => (
  <WarningMessage {...args} />
);

export const Primary = Template.bind({});
// @ts-ignore
Primary.args = {
  message: 'Some warning message',
  className: 'p-2 mb-3 mt-3',
} as Partial<typeof WarningMessage>;