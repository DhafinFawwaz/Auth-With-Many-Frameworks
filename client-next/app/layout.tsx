import type { Metadata } from 'next'
import { League_Spartan } from 'next/font/google'
import './globals.css'
import { NextAuthProvider } from './providers'
import { getServerSession } from 'next-auth'

const leagueSpartan = League_Spartan({ 
  subsets: ['latin'],
  weight: ["400", "600", "700", "800"],
  fallback: ["Roboto","Poppins", "sans-serif"],
  // variable: "--league-spartan",
  display: "swap",
})

export const metadata: Metadata = {
  title: 'Next Django Auth Template',
  description: 'For template',
}

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {

  return (
    <html lang="en" className='dark'>
      <body className={leagueSpartan.className}>
        <NextAuthProvider>
          {children}
        </NextAuthProvider>
      </body>
    </html>
  )
}
