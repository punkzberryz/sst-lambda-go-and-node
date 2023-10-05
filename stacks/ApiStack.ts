import { StackContext, Api, EventBus } from "sst/constructs";

export function NodeAPI({ stack }: StackContext) {
  const api = new Api(stack, "node-api", {
    defaults: {},
    routes: {
      "GET /": "packages/lambda/node/src/lambda.handler",
    },
  });

  stack.addOutputs({
    NodeApiEndpoint: api.url,
  });
}

export function GoAPI({ stack, app }: StackContext) {
  const api = new Api(stack, "go-api", {
    routes: {
      "GET /": "packages/lambda/go/hello/main.go",
      "GET /notes": "packages/lambda/go/notes/getAllNotes/main.go",
      "GET /notes/{id}": "packages/lambda/go/notes/getNoteById/main.go",
    },
  });

  stack.addOutputs({
    GoApiEndpoint: api.url,
  });
}
