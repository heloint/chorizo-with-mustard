import './globals.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootswatch/dist/flatly/bootstrap.min.css'
import Navbar from './navbar'
import GetUserProfile from './(auth)/profile';
import Head from "next/head";
import Script from "next/script";
import Carousel from 'react-bootstrap/Carousel'

export const metadata = {
    title: 'Create Next App',
    description: 'Generated by create next app',
}

export default async function RootLayout({
    children,
}: {
    children: React.ReactNode       
}) {

        const user = await GetUserProfile();

    return (
        <html lang="en">
            <Head>
                <meta name="viewport" content="width=device-width, initial-scale=1" />
                <Script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/js/bootstrap.bundle.min.js" defer/>
            </Head>
            <body>
                <Navbar {...user}/>
                {children}
            </body>
        </html>
    )
}
