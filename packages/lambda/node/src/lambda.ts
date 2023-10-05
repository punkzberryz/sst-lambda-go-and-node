import { ApiHandler } from "sst/node/api";

export const handler = ApiHandler(async (_evt) => {
  return {
    statusCode: 200,
    body: `Hello world from Node. The time is ${new Date().toISOString()}`,
  };
});
