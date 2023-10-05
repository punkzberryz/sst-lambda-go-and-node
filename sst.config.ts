import { SSTConfig } from "sst";
import { GoAPI, NodeAPI } from "./stacks/ApiStack";

export default {
  config(_input) {
    return {
      name: "sst-lamda-go-and-node",
      region: "ap-southeast-1",
    };
  },
  stacks(app) {
    app.stack(NodeAPI);
    app.setDefaultFunctionProps({
      runtime: "go",
      timeout: 30,
    });
    app.stack(GoAPI);
  },
} satisfies SSTConfig;
