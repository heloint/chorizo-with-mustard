import { cookies } from 'next/headers';

export default async function GetUserProfile() {

    let headers: HeadersInit = {
        "Access-Control-Allow-Origin": "*",
        "Content-Type": "text/plain",
    }

    const jwtToken = cookies().get("jwt")?.value;

    if (jwtToken) {
        headers['Authorization'] = `Bearer ${jwtToken}`
    }

    const res = await fetch('http://localhost:8000/profile', {
        cache: 'no-store',
        headers: headers,
        credentials: 'include'
    })

    /* if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data');
    } */

    return res.json()
}
