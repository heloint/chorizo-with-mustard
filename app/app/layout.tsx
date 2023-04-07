import './globals.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootswatch/dist/flatly/bootstrap.min.css'
import Navbar from './navbar'

import Head from "next/head";
import Script from "next/script";

export const metadata = {
    title: 'Create Next App',
    description: 'Generated by create next app',
}

export default function RootLayout({
    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
            <Head>
                <meta name="viewport" content="width=device-width, initial-scale=1" />
            </Head>
            <Script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/js/bootstrap.bundle.min.js"
                integrity="sha384-ENjdO4Dr2bkBIFxQpeoTz1HIcje39Wm4jDKdf19U8gI4ddQ3GYNS7NTKfAdVQSZe"
                crossOrigin="anonymous"
            />
            <body>
                <Navbar />
                {children}
            </body>
        </html>
    )
}
