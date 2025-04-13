import { LoaderFunction } from "@remix-run/node";
import { authenticator } from "~/auth.server";

export let loader: LoaderFunction = async ({ request }) => {
  return authenticator.authenticate("google", request);
};
