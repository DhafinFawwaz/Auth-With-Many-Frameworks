import NextAuth, { DefaultSession, User } from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials";
import { getServerSession } from 'next-auth';
import { redirect } from "next/navigation";

import type { NextAuthOptions } from 'next-auth'
import { ENV } from "../../env";

type UserData = {
    username: string;
    email: string;
    nim: string;
    password: string;
    accessToken: string;
    refreshToken: string;
}
export interface UserSession extends DefaultSession {
    data: UserData;
}

export const authOptions: NextAuthOptions = {
    secret: ENV.AUTH_SECRET,
    session: {
        strategy: "jwt",
    },
    providers: [
        CredentialsProvider({
            name: "Sign in",
            credentials: {
              email: {
                label: "Email",
                type: "email",
                placeholder: "example@example.com",
              },
              username: {
                label: "Username",
                type: "username",
                placeholder: "example",
              },
              password: { label: "Password", type: "password" },
        },
        async authorize(credentials) {

          const url = `${ENV.API_URL}/api/auth/login/`;
          const res = await fetch(url, {
            method: "POST",
            body: JSON.stringify(credentials),
            headers: {
              "Content-Type": "application/json",
            },
          }).catch((err) => {
            throw err;
          });
          const user = await res!.json();
          
          if (res.ok && user) {
            return user;
          }
          return null;
        },
      }),
    ],
  
    callbacks: {
      async jwt({ token, user }) {// user is returned from authorize() above
        if(user) token.user = user;
        return token;
      },
  
      async session({ session, token }) { // token is returned from jwt() above
        const accessToken = (token.user as UserData).accessToken;
        const refreshToken = (token.user as UserData).refreshToken;
        
        const url = `${process.env.API_URL}/api/auth/`;
          const res = await fetch(url, {
            method: "POST",
            body: JSON.stringify(session),
            headers: {
              "Content-Type": "application/json",
              "Authorization": `Bearer ${accessToken}`,
            },
          });
        const user: UserData = await res.json();
        (session as UserSession).data = user; // get session.user by useSession()

        return session;

      },
    },
}
const handler = NextAuth(authOptions);
export { handler as GET, handler as POST };

export async function getSessionAndEnsureAuthenticated(): Promise<UserSession | null>{
  const session: UserSession | null = await getServerSession(authOptions);
  if(!session) redirect('/login');
  return session;
}
export async function ensureNotAuthenticated(){
  const session: UserSession | null = await getServerSession(authOptions);
  if(session) redirect('/profile');
}