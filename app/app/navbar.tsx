import Link from "next/link";
import { cookies } from 'next/headers';

export async function getServerSideProps(context: any) {
  const cookies = context.req.headers.cookie['jwt'];
  console.log(cookies);
  return {
    props: {},
  };
}

export default function Navbar() {
    const cookieStore = cookies();
    console.log(cookieStore.getAll());
    return (
        <nav className="navbar navbar-expand-lg navbar-dark bg-primary">
            <div className="container-fluid">
                <Link href="/" className="navbar-brand">Home</Link>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarColor01">
                    <ul className="navbar-nav me-auto">
                        <li className="nav-item">
                            <Link href="/login" className="nav-link" >Login</Link>
                        </li>
                        <li className="nav-item">
                            <Link href="/register" className="nav-link" >Register</Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    );
}
