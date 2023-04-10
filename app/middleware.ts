import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

type User = {
    Id: number
    Role: string
    Password: string
    Email: string
    Username: string
    Firstname: string
    Lastname: string
    RegistrationDate: string
}

type SessionVar = {
    IsLoggedIn: boolean
    User: User
}

async function GetSession(request: NextRequest) {

    let headers: HeadersInit = {
        "Access-Control-Allow-Origin": "*",
        "Content-Type": "text/plain",
    }

    const jwtToken: string | undefined = request.cookies.get('jwt')?.value

    if (jwtToken) {
        headers['Authorization'] = `Bearer ${jwtToken}`
    }

    const res = await fetch('http://localhost:8000/profile', {
        cache: 'no-store',
        headers: headers,
        credentials: 'include'
    })

    return res.json()
}

export async function middleware(request: NextRequest) {

    const session: SessionVar = await GetSession(request);

    // LOGIN
    if (request.nextUrl.pathname === '/login' && session.IsLoggedIn) {
            return NextResponse.redirect(new URL('/', request.url))
    }

    // REGISTER
    if (request.nextUrl.pathname === '/register' && session.IsLoggedIn) {
            return NextResponse.redirect(new URL('/', request.url))
    }
    
    // LOGOUT
    if (request.nextUrl.pathname === '/logout' && !session.IsLoggedIn) {
            return NextResponse.redirect(new URL('/login', request.url))
    }

}


// See "Matching Paths" below to learn more
export const config = {
    matcher: [
        '/login:path*', 
        '/logout:path*'
    ]
}
