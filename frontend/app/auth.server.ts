import { Authenticator } from "remix-auth";
import { GoogleStrategy } from "remix-auth-google";
import { sessionStorage } from "./session.server";

export const authenticator = new Authenticator(sessionStorage);

const googleStrategy = new GoogleStrategy(
  {
    clientID: process.env.GOOGLE_CLIENT_ID!,
    clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
    callbackURL: "http://localhost:5173/auth/google/callback",
  },
  async ({ profile }) => {
    // ユーザー情報を保存または取得するロジックをここに記述
    return profile;
  }
);

authenticator.use(googleStrategy);

export async function getUser(request: Request) {
  return await authenticator.isAuthenticated(request);
}
