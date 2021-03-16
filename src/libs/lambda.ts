import middy from '@middy/core';

// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
export const middyfy = (handler) => {
  return middy(handler);
};
