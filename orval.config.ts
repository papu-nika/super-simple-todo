import { defineConfig } from "orval";

export default defineConfig({
  cloud_cost: {
    input: "./openapi/openapi.yaml",

    output: {
      mode: "tags-split",
      target: "frontend/app/client/",
      schemas: "frontend/app/model",
      client: "react-query",
    },
  },
});
