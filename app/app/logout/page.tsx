'use client'

const logout = async () => {
    console.log(process.env.NEXT_PUBLIC_GO_API);
    // const res = await fetch('http://localhost:8000/logout', {
    const res = await fetch(`${process.env.NEXT_PUBLIC_GO_API}/logout`, {
        method: "GET",
        credentials: 'include',
    });

    if (res.ok) {
        window.location.href = '/';
    }
}

export default function Logout() {
    logout();
}
